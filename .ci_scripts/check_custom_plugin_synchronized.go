package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	frameworkresource "github.com/hashicorp/terraform-plugin-framework/resource"
	frameworkschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	custom "github.com/kong/terraform-provider-konnect/v3/src"
	"github.com/zclconf/go-cty/cty"
	"gopkg.in/yaml.v3"
)

const (
	openAPIPath             = "openapi.yaml"
	schemaName              = "PluginBase"
	customPluginResource    = "konnect_gateway_custom_plugin"
	customPluginFixturePath = "tests/resources/testdata/TestGatewayCustomPlugin/CRUD/main.tf"
)

type openAPIDocument struct {
	Components struct {
		Schemas map[string]*openAPISchema `yaml:"schemas"`
	} `yaml:"components"`
}

type openAPISchema struct {
	Ref        string                    `yaml:"$ref"`
	Properties map[string]*openAPISchema `yaml:"properties"`
	Items      *openAPISchema            `yaml:"items"`
	AllOf      []*openAPISchema          `yaml:"allOf"`
	AnyOf      []*openAPISchema          `yaml:"anyOf"`
	OneOf      []*openAPISchema          `yaml:"oneOf"`
}

type synchronizationResult struct {
	missingFromModel   []string
	missingFromFixture []string
}

func main() {
	result, err := checkSynchronization()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot check custom plugin synchronization: %v\n", err)
		os.Exit(1)
	}

	if len(result.missingFromModel) > 0 {
		printMissingProperties("PluginBase properties missing from CustomPluginResourceModel:", result.missingFromModel)
	}
	if len(result.missingFromFixture) > 0 {
		printMissingProperties("Configurable PluginBase properties missing from "+customPluginFixturePath+":", result.missingFromFixture)
	}
	if len(result.missingFromModel) > 0 || len(result.missingFromFixture) > 0 {
		os.Exit(1)
	}

	fmt.Printf(
		"CustomPluginResourceModel and %s contain all required properties from #/components/schemas/%s.\n",
		customPluginFixturePath,
		schemaName,
	)
}

func checkSynchronization() (synchronizationResult, error) {
	openAPIProperties, err := pluginBaseProperties()
	if err != nil {
		return synchronizationResult{}, err
	}

	modelProperties := make(map[string]struct{})
	collectTerraformProperties(
		reflect.TypeOf(custom.CustomPluginResourceModel{}),
		"",
		make(map[reflect.Type]bool),
		modelProperties,
	)

	fixtureProperties, err := customPluginFixtureProperties()
	if err != nil {
		return synchronizationResult{}, err
	}

	computedOnlyProperties, err := customPluginComputedOnlyProperties()
	if err != nil {
		return synchronizationResult{}, err
	}

	result := synchronizationResult{
		missingFromModel: missingProperties(openAPIProperties, modelProperties, nil),
		missingFromFixture: missingProperties(
			openAPIProperties,
			fixtureProperties,
			computedOnlyProperties,
		),
	}

	return result, nil
}

func pluginBaseProperties() (map[string]struct{}, error) {
	contents, err := os.ReadFile(openAPIPath)
	if err != nil {
		return nil, fmt.Errorf("read OpenAPI document %s: %w", openAPIPath, err)
	}

	var document openAPIDocument
	if err := yaml.Unmarshal(contents, &document); err != nil {
		return nil, fmt.Errorf("parse OpenAPI document %s: %w", openAPIPath, err)
	}

	pluginBase, ok := document.Components.Schemas[schemaName]
	if !ok {
		return nil, fmt.Errorf("schema #/components/schemas/%s does not exist", schemaName)
	}

	properties := make(map[string]struct{})
	if err := collectOpenAPIProperties(&document, pluginBase, "", make(map[*openAPISchema]bool), properties); err != nil {
		return nil, err
	}

	return properties, nil
}

func missingProperties(
	expected map[string]struct{},
	actual map[string]struct{},
	excluded map[string]struct{},
) []string {
	missing := make([]string, 0)
	for property := range expected {
		if hasPropertyOrAncestor(excluded, property) {
			continue
		}
		if _, ok := actual[property]; !ok {
			missing = append(missing, property)
		}
	}
	sort.Strings(missing)

	return missing
}

func hasPropertyOrAncestor(properties map[string]struct{}, property string) bool {
	for candidate := range properties {
		if property == candidate || strings.HasPrefix(property, candidate+".") || strings.HasPrefix(property, candidate+"[]") {
			return true
		}
	}
	return false
}

func printMissingProperties(header string, properties []string) {
	fmt.Fprintln(os.Stderr, header)
	for _, property := range properties {
		fmt.Fprintf(os.Stderr, "  - %s\n", property)
	}
}

func customPluginFixtureProperties() (map[string]struct{}, error) {
	contents, err := os.ReadFile(customPluginFixturePath)
	if err != nil {
		return nil, fmt.Errorf("read Terraform fixture %s: %w", customPluginFixturePath, err)
	}

	file, diagnostics := hclsyntax.ParseConfig(contents, customPluginFixturePath, hcl.InitialPos)
	if diagnostics.HasErrors() {
		return nil, fmt.Errorf("parse Terraform fixture %s: %s", customPluginFixturePath, diagnostics.Error())
	}

	body, ok := file.Body.(*hclsyntax.Body)
	if !ok {
		return nil, fmt.Errorf("parse Terraform fixture %s: unexpected body type %T", customPluginFixturePath, file.Body)
	}

	properties := make(map[string]struct{})
	resourceCount := 0
	for _, block := range body.Blocks {
		if block.Type != "resource" || len(block.Labels) < 2 || block.Labels[0] != customPluginResource {
			continue
		}

		resourceCount++
		for name, attribute := range block.Body.Attributes {
			properties[name] = struct{}{}
			if err := collectHCLExpressionProperties(attribute.Expr, name, properties); err != nil {
				return nil, err
			}
		}
	}

	if resourceCount == 0 {
		return nil, fmt.Errorf("Terraform fixture %s has no %s resources", customPluginFixturePath, customPluginResource)
	}

	return properties, nil
}

func collectHCLExpressionProperties(expression hcl.Expression, prefix string, properties map[string]struct{}) error {
	switch expression := expression.(type) {
	case *hclsyntax.ObjectConsExpr:
		for _, item := range expression.Items {
			key, diagnostics := item.KeyExpr.Value(nil)
			if diagnostics.HasErrors() {
				return fmt.Errorf("read object key at %s: %s", item.KeyExpr.Range(), diagnostics.Error())
			}
			if !key.IsKnown() || key.IsNull() || !key.Type().Equals(cty.String) {
				return fmt.Errorf("object key at %s must be a static string", item.KeyExpr.Range())
			}

			propertyPath := joinPropertyPath(prefix, key.AsString())
			properties[propertyPath] = struct{}{}
			if err := collectHCLExpressionProperties(item.ValueExpr, propertyPath, properties); err != nil {
				return err
			}
		}
	case *hclsyntax.TupleConsExpr:
		for _, item := range expression.Exprs {
			if err := collectHCLExpressionProperties(item, prefix+"[]", properties); err != nil {
				return err
			}
		}
	case *hclsyntax.ParenthesesExpr:
		return collectHCLExpressionProperties(expression.Expression, prefix, properties)
	}

	return nil
}

func customPluginComputedOnlyProperties() (map[string]struct{}, error) {
	var response frameworkresource.SchemaResponse
	custom.NewCustomPluginResource().Schema(
		context.Background(),
		frameworkresource.SchemaRequest{},
		&response,
	)
	if response.Diagnostics.HasError() {
		errors := make([]string, 0, len(response.Diagnostics.Errors()))
		for _, diagnostic := range response.Diagnostics.Errors() {
			errors = append(errors, diagnostic.Summary()+": "+diagnostic.Detail())
		}
		return nil, fmt.Errorf("get custom plugin Terraform schema: %s", strings.Join(errors, "; "))
	}

	properties := make(map[string]struct{})
	collectComputedOnlyProperties(response.Schema.Attributes, "", properties)

	return properties, nil
}

func collectComputedOnlyProperties(
	attributes map[string]frameworkschema.Attribute,
	prefix string,
	properties map[string]struct{},
) {
	for name, attribute := range attributes {
		propertyPath := joinPropertyPath(prefix, name)
		if attribute.IsComputed() && !attribute.IsOptional() && !attribute.IsRequired() {
			properties[propertyPath] = struct{}{}
			continue
		}

		switch attribute := attribute.(type) {
		case frameworkschema.SingleNestedAttribute:
			collectComputedOnlyProperties(attribute.Attributes, propertyPath, properties)
		case frameworkschema.ListNestedAttribute:
			collectComputedOnlyProperties(attribute.NestedObject.Attributes, propertyPath+"[]", properties)
		case frameworkschema.SetNestedAttribute:
			collectComputedOnlyProperties(attribute.NestedObject.Attributes, propertyPath+"[]", properties)
		case frameworkschema.MapNestedAttribute:
			collectComputedOnlyProperties(attribute.NestedObject.Attributes, propertyPath+"[]", properties)
		}
	}
}

func collectOpenAPIProperties(
	document *openAPIDocument,
	schema *openAPISchema,
	prefix string,
	active map[*openAPISchema]bool,
	properties map[string]struct{},
) error {
	if schema == nil || active[schema] {
		return nil
	}

	active[schema] = true
	defer delete(active, schema)

	if schema.Ref != "" {
		referencedSchema, err := resolveSchemaReference(document, schema.Ref)
		if err != nil {
			return err
		}
		if err := collectOpenAPIProperties(document, referencedSchema, prefix, active, properties); err != nil {
			return err
		}
	}

	for _, composedSchemas := range [][]*openAPISchema{schema.AllOf, schema.AnyOf, schema.OneOf} {
		for _, composedSchema := range composedSchemas {
			if err := collectOpenAPIProperties(document, composedSchema, prefix, active, properties); err != nil {
				return err
			}
		}
	}

	propertyNames := make([]string, 0, len(schema.Properties))
	for name := range schema.Properties {
		propertyNames = append(propertyNames, name)
	}
	sort.Strings(propertyNames)

	for _, name := range propertyNames {
		propertyPath := joinPropertyPath(prefix, name)
		properties[propertyPath] = struct{}{}
		if err := collectOpenAPIProperties(document, schema.Properties[name], propertyPath, active, properties); err != nil {
			return err
		}
	}

	if schema.Items != nil {
		if err := collectOpenAPIProperties(document, schema.Items, prefix+"[]", active, properties); err != nil {
			return err
		}
	}

	return nil
}

func resolveSchemaReference(document *openAPIDocument, ref string) (*openAPISchema, error) {
	const schemaReferencePrefix = "#/components/schemas/"
	if !strings.HasPrefix(ref, schemaReferencePrefix) {
		return nil, fmt.Errorf("unsupported schema reference %q", ref)
	}

	encodedName := strings.TrimPrefix(ref, schemaReferencePrefix)
	if encodedName == "" || strings.Contains(encodedName, "/") {
		return nil, fmt.Errorf("unsupported schema reference %q", ref)
	}

	name := strings.ReplaceAll(encodedName, "~1", "/")
	name = strings.ReplaceAll(name, "~0", "~")

	schema, ok := document.Components.Schemas[name]
	if !ok {
		return nil, fmt.Errorf("schema reference %q does not exist", ref)
	}

	return schema, nil
}

func collectTerraformProperties(
	typeOf reflect.Type,
	prefix string,
	active map[reflect.Type]bool,
	properties map[string]struct{},
) {
	for typeOf.Kind() == reflect.Pointer {
		typeOf = typeOf.Elem()
	}

	for typeOf.Kind() == reflect.Array || typeOf.Kind() == reflect.Slice {
		prefix += "[]"
		typeOf = typeOf.Elem()
		for typeOf.Kind() == reflect.Pointer {
			typeOf = typeOf.Elem()
		}
	}

	if typeOf.Kind() != reflect.Struct || active[typeOf] {
		return
	}

	active[typeOf] = true
	defer delete(active, typeOf)

	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		if !field.IsExported() {
			continue
		}

		tag, ok := field.Tag.Lookup("tfsdk")
		if !ok {
			if field.Anonymous {
				collectTerraformProperties(field.Type, prefix, active, properties)
			}
			continue
		}

		name := strings.Split(tag, ",")[0]
		if name == "" || name == "-" {
			continue
		}

		propertyPath := joinPropertyPath(prefix, name)
		properties[propertyPath] = struct{}{}
		collectTerraformProperties(field.Type, propertyPath, active, properties)
	}
}

func joinPropertyPath(prefix, name string) string {
	if prefix == "" {
		return name
	}
	return prefix + "." + name
}
