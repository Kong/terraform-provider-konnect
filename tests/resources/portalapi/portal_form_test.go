package portalapi

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestPortalForm(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		const form = "konnect_portal_form.test_form"
		const dev = "konnect_portal_developer.test_developer"
		const apiForm = "konnect_portal_form.test_api_form"
		const pub = "konnect_api_publication.test_api_publication"
		const reg = "konnect_portal_application_registration.test_registration"
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(form, plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction(dev, plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction(apiForm, plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction(reg, plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet(form, "id"),
						resource.TestCheckResourceAttr(form, "developer_registration.status", "published"),
						resource.TestCheckResourceAttr(form, "developer_registration.fields.2.text.name", "custom"),
						resource.TestCheckResourceAttr(form, "developer_registration.fields.2.text.label", "custom"),
						resource.TestCheckResourceAttr(form, "developer_registration.fields.2.text.required", "true"),
						resource.TestCheckResourceAttrSet(apiForm, "api_registration.id"),
						resource.TestCheckResourceAttr(apiForm, "api_registration.name", "tftest-api-registration"),
						resource.TestCheckResourceAttr(apiForm, "api_registration.status", "published"),
						resource.TestCheckResourceAttrPair(pub, "form_id", apiForm, "api_registration.id"),
						resource.TestCheckResourceAttrSet(reg, "id"),
						resource.TestCheckResourceAttr(reg, "status", "approved"),
						resource.TestCheckResourceAttrWith(reg, "additional_data", func(v string) error {
							var m map[string]any
							if err := json.Unmarshal([]byte(v), &m); err != nil {
								return err
							}
							if m["use_case"] != "tftest use case" {
								return fmt.Errorf("unexpected additional_data.use_case: %v", m["use_case"])
							}
							return nil
						}),
						resource.TestCheckResourceAttrSet(dev, "id"),
						resource.TestCheckResourceAttr(dev, "email", "tftestformdev@example.com"),
						resource.TestCheckResourceAttr(dev, "status", "approved"),
						resource.TestCheckResourceAttrWith(dev, "additional_data", func(v string) error {
							var m map[string]any
							if err := json.Unmarshal([]byte(v), &m); err != nil {
								return err
							}
							if m["custom"] != "tftestcustomvalue" {
								return fmt.Errorf("unexpected additional_data.custom: %v", m["custom"])
							}
							return nil
						}),
					),
				},
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(form, plancheck.ResourceActionUpdate),
							plancheck.ExpectResourceAction(dev, plancheck.ResourceActionNoop),
							plancheck.ExpectResourceAction(apiForm, plancheck.ResourceActionNoop),
							plancheck.ExpectResourceAction(reg, plancheck.ResourceActionNoop),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(form, "developer_registration.fields.2.text.label", "custom updated"),
						resource.TestCheckResourceAttr(form, "developer_registration.fields.2.text.required", "true"),
					),
				},
			},
		})
	})
}
