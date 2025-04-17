// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	custom "github.com/kong/terraform-provider-konnect/v2/src"
	"net/http"
	"os"
)

var _ provider.Provider = &KonnectProvider{}

type KonnectProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// KonnectProviderModel describes the provider data model.
type KonnectProviderModel struct {
	KonnectAccessToken       types.String `tfsdk:"konnect_access_token"`
	PersonalAccessToken      types.String `tfsdk:"personal_access_token"`
	ServerURL                types.String `tfsdk:"server_url"`
	SystemAccountAccessToken types.String `tfsdk:"system_account_access_token"`
}

func (p *KonnectProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "konnect"
	resp.Version = p.version
}

func (p *KonnectProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"konnect_access_token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"personal_access_token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"server_url": schema.StringAttribute{
				Description: `Server URL (defaults to https://global.api.konghq.com)`,
				Optional:    true,
			},
			"system_account_access_token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
		MarkdownDescription: `Konnect API: The Konnect platform API`,
	}
}

func (p *KonnectProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data KonnectProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" && len(os.Getenv("KONNECT_SERVER_URL")) > 0 {
		ServerURL = os.Getenv("KONNECT_SERVER_URL")
	}
	if ServerURL == "" {
		ServerURL = "https://global.api.konghq.com"
	}

	personalAccessToken := new(string)
	if !data.PersonalAccessToken.IsUnknown() && !data.PersonalAccessToken.IsNull() {
		*personalAccessToken = data.PersonalAccessToken.ValueString()
	} else {
		if len(os.Getenv("KONNECT_TOKEN")) > 0 {
			*personalAccessToken = os.Getenv("KONNECT_TOKEN")
		} else {
			personalAccessToken = nil
		}
	}
	systemAccountAccessToken := new(string)
	if !data.SystemAccountAccessToken.IsUnknown() && !data.SystemAccountAccessToken.IsNull() {
		*systemAccountAccessToken = data.SystemAccountAccessToken.ValueString()
	} else {
		if len(os.Getenv("KONNECT_SPAT")) > 0 {
			*systemAccountAccessToken = os.Getenv("KONNECT_SPAT")
		} else {
			systemAccountAccessToken = nil
		}
	}
	konnectAccessToken := new(string)
	if !data.KonnectAccessToken.IsUnknown() && !data.KonnectAccessToken.IsNull() {
		*konnectAccessToken = data.KonnectAccessToken.ValueString()
	} else {
		konnectAccessToken = nil
	}
	security := shared.Security{
		PersonalAccessToken:      personalAccessToken,
		SystemAccountAccessToken: systemAccountAccessToken,
		KonnectAccessToken:       konnectAccessToken,
	}

	providerHTTPTransportOpts := ProviderHTTPTransportOpts{
		SetHeaders: make(map[string]string),
		Transport:  http.DefaultTransport,
	}

	httpClient := http.DefaultClient
	httpClient.Transport = NewProviderHTTPTransport(providerHTTPTransportOpts)

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
		sdk.WithClient(httpClient),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *KonnectProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAPIProductResource,
		NewAPIProductDocumentResource,
		NewAPIProductSpecificationResource,
		NewAPIProductVersionResource,
		NewApplicationAuthStrategyResource,
		NewAuditLogResource,
		NewAuditLogDestinationResource,
		NewCloudGatewayConfigurationResource,
		NewCloudGatewayCustomDomainResource,
		NewCloudGatewayNetworkResource,
		NewCloudGatewayTransitGatewayResource,
		NewGatewayACLResource,
		NewGatewayBasicAuthResource,
		NewGatewayCACertificateResource,
		NewGatewayCertificateResource,
		NewGatewayConfigStoreResource,
		NewGatewayConsumerResource,
		NewGatewayConsumerGroupResource,
		NewGatewayConsumerGroupMemberResource,
		NewGatewayControlPlaneResource,
		NewGatewayControlPlaneMembershipResource,
		NewGatewayCustomPluginSchemaResource,
		NewGatewayCustomPluginStreamingResource,
		NewGatewayDataPlaneClientCertificateResource,
		NewGatewayHMACAuthResource,
		NewGatewayJWTResource,
		NewGatewayKeyResource,
		NewGatewayKeyAuthResource,
		NewGatewayKeySetResource,
		NewGatewayMTLSAuthResource,
		NewGatewayPluginACLResource,
		NewGatewayPluginAcmeResource,
		NewGatewayPluginAiAzureContentSafetyResource,
		NewGatewayPluginAiPromptDecoratorResource,
		NewGatewayPluginAiPromptGuardResource,
		NewGatewayPluginAiPromptTemplateResource,
		NewGatewayPluginAiProxyResource,
		NewGatewayPluginAiProxyAdvancedResource,
		NewGatewayPluginAiRateLimitingAdvancedResource,
		NewGatewayPluginAiRequestTransformerResource,
		NewGatewayPluginAiResponseTransformerResource,
		NewGatewayPluginAiSemanticCacheResource,
		NewGatewayPluginAiSemanticPromptGuardResource,
		NewGatewayPluginAwsLambdaResource,
		NewGatewayPluginAzureFunctionsResource,
		NewGatewayPluginBasicAuthResource,
		NewGatewayPluginBotDetectionResource,
		NewGatewayPluginCanaryResource,
		NewGatewayPluginConfluentResource,
		NewGatewayPluginCorrelationIDResource,
		NewGatewayPluginCorsResource,
		NewGatewayPluginDatadogResource,
		NewGatewayPluginDatadogTracingResource,
		NewGatewayPluginDegraphqlResource,
		NewGatewayPluginExitTransformerResource,
		NewGatewayPluginFileLogResource,
		NewGatewayPluginForwardProxyResource,
		NewGatewayPluginGraphqlProxyCacheAdvancedResource,
		NewGatewayPluginGraphqlRateLimitingAdvancedResource,
		NewGatewayPluginGrpcGatewayResource,
		NewGatewayPluginGrpcWebResource,
		NewGatewayPluginHeaderCertAuthResource,
		NewGatewayPluginHmacAuthResource,
		NewGatewayPluginHTTPLogResource,
		NewGatewayPluginInjectionProtectionResource,
		NewGatewayPluginIPRestrictionResource,
		NewGatewayPluginJqResource,
		NewGatewayPluginJSONThreatProtectionResource,
		NewGatewayPluginJweDecryptResource,
		NewGatewayPluginJwtResource,
		NewGatewayPluginJwtSignerResource,
		NewGatewayPluginKafkaLogResource,
		NewGatewayPluginKafkaUpstreamResource,
		NewGatewayPluginKeyAuthResource,
		NewGatewayPluginKeyAuthEncResource,
		NewGatewayPluginLdapAuthResource,
		NewGatewayPluginLdapAuthAdvancedResource,
		NewGatewayPluginLogglyResource,
		NewGatewayPluginMockingResource,
		NewGatewayPluginMtlsAuthResource,
		NewGatewayPluginOasValidationResource,
		NewGatewayPluginOauth2Resource,
		NewGatewayPluginOauth2IntrospectionResource,
		NewGatewayPluginOpaResource,
		NewGatewayPluginOpenidConnectResource,
		NewGatewayPluginOpentelemetryResource,
		NewGatewayPluginPostFunctionResource,
		NewGatewayPluginPreFunctionResource,
		NewGatewayPluginPrometheusResource,
		NewGatewayPluginProxyCacheResource,
		NewGatewayPluginProxyCacheAdvancedResource,
		NewGatewayPluginRateLimitingResource,
		NewGatewayPluginRateLimitingAdvancedResource,
		NewGatewayPluginRedirectResource,
		NewGatewayPluginRequestSizeLimitingResource,
		NewGatewayPluginRequestTerminationResource,
		NewGatewayPluginRequestTransformerResource,
		NewGatewayPluginRequestTransformerAdvancedResource,
		NewGatewayPluginRequestValidatorResource,
		NewGatewayPluginResponseRatelimitingResource,
		NewGatewayPluginResponseTransformerResource,
		NewGatewayPluginResponseTransformerAdvancedResource,
		NewGatewayPluginRouteByHeaderResource,
		NewGatewayPluginRouteTransformerAdvancedResource,
		NewGatewayPluginSamlResource,
		NewGatewayPluginServiceProtectionResource,
		NewGatewayPluginSessionResource,
		NewGatewayPluginStandardWebhooksResource,
		NewGatewayPluginStatsdResource,
		NewGatewayPluginStatsdAdvancedResource,
		NewGatewayPluginSyslogResource,
		NewGatewayPluginTCPLogResource,
		NewGatewayPluginTLSHandshakeModifierResource,
		NewGatewayPluginTLSMetadataHeadersResource,
		NewGatewayPluginUDPLogResource,
		NewGatewayPluginUpstreamOauthResource,
		NewGatewayPluginUpstreamTimeoutResource,
		NewGatewayPluginVaultAuthResource,
		NewGatewayPluginWebsocketSizeLimitResource,
		NewGatewayPluginWebsocketValidatorResource,
		NewGatewayPluginXMLThreatProtectionResource,
		NewGatewayPluginZipkinResource,
		NewGatewayRouteResource,
		NewGatewayRouteExpressionResource,
		NewGatewayServiceResource,
		NewGatewaySNIResource,
		NewGatewayTargetResource,
		NewGatewayUpstreamResource,
		NewGatewayVaultResource,
		NewMeshControlPlaneResource,
		NewPortalResource,
		NewPortalAppearanceResource,
		NewPortalAuthResource,
		NewPortalProductVersionResource,
		NewPortalTeamResource,
		NewServerlessCloudGatewayResource,
		NewSystemAccountResource,
		NewSystemAccountAccessTokenResource,
		NewSystemAccountRoleResource,
		NewSystemAccountTeamResource,
		NewTeamResource,
		NewTeamRoleResource,
		NewTeamUserResource,
		custom.NewCustomPluginResource,
	}
}

func (p *KonnectProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAPIProductDataSource,
		NewAPIProductDocumentDataSource,
		NewAPIProductSpecificationDataSource,
		NewAPIProductVersionDataSource,
		NewApplicationAuthStrategyDataSource,
		NewAuditLogDataSource,
		NewAuditLogDestinationDataSource,
		NewCloudGatewayConfigurationDataSource,
		NewCloudGatewayCustomDomainDataSource,
		NewCloudGatewayNetworkDataSource,
		NewCloudGatewayProviderAccountListDataSource,
		NewCloudGatewayTransitGatewayDataSource,
		NewGatewayACLDataSource,
		NewGatewayBasicAuthDataSource,
		NewGatewayCACertificateDataSource,
		NewGatewayCertificateDataSource,
		NewGatewayConfigStoreDataSource,
		NewGatewayConsumerDataSource,
		NewGatewayConsumerGroupDataSource,
		NewGatewayControlPlaneDataSource,
		NewGatewayControlPlaneListDataSource,
		NewGatewayCustomPluginSchemaDataSource,
		NewGatewayCustomPluginStreamingDataSource,
		NewGatewayDataPlaneClientCertificateDataSource,
		NewGatewayHMACAuthDataSource,
		NewGatewayJWTDataSource,
		NewGatewayKeyDataSource,
		NewGatewayKeyAuthDataSource,
		NewGatewayKeySetDataSource,
		NewGatewayMTLSAuthDataSource,
		NewGatewayPluginACLDataSource,
		NewGatewayPluginAcmeDataSource,
		NewGatewayPluginAiAzureContentSafetyDataSource,
		NewGatewayPluginAiPromptDecoratorDataSource,
		NewGatewayPluginAiPromptGuardDataSource,
		NewGatewayPluginAiPromptTemplateDataSource,
		NewGatewayPluginAiProxyDataSource,
		NewGatewayPluginAiProxyAdvancedDataSource,
		NewGatewayPluginAiRateLimitingAdvancedDataSource,
		NewGatewayPluginAiRequestTransformerDataSource,
		NewGatewayPluginAiResponseTransformerDataSource,
		NewGatewayPluginAiSemanticCacheDataSource,
		NewGatewayPluginAiSemanticPromptGuardDataSource,
		NewGatewayPluginAwsLambdaDataSource,
		NewGatewayPluginAzureFunctionsDataSource,
		NewGatewayPluginBasicAuthDataSource,
		NewGatewayPluginBotDetectionDataSource,
		NewGatewayPluginCanaryDataSource,
		NewGatewayPluginConfluentDataSource,
		NewGatewayPluginCorrelationIDDataSource,
		NewGatewayPluginCorsDataSource,
		NewGatewayPluginDatadogDataSource,
		NewGatewayPluginDatadogTracingDataSource,
		NewGatewayPluginDegraphqlDataSource,
		NewGatewayPluginExitTransformerDataSource,
		NewGatewayPluginFileLogDataSource,
		NewGatewayPluginForwardProxyDataSource,
		NewGatewayPluginGraphqlProxyCacheAdvancedDataSource,
		NewGatewayPluginGraphqlRateLimitingAdvancedDataSource,
		NewGatewayPluginGrpcGatewayDataSource,
		NewGatewayPluginGrpcWebDataSource,
		NewGatewayPluginHeaderCertAuthDataSource,
		NewGatewayPluginHmacAuthDataSource,
		NewGatewayPluginHTTPLogDataSource,
		NewGatewayPluginInjectionProtectionDataSource,
		NewGatewayPluginIPRestrictionDataSource,
		NewGatewayPluginJqDataSource,
		NewGatewayPluginJSONThreatProtectionDataSource,
		NewGatewayPluginJweDecryptDataSource,
		NewGatewayPluginJwtDataSource,
		NewGatewayPluginJwtSignerDataSource,
		NewGatewayPluginKafkaLogDataSource,
		NewGatewayPluginKafkaUpstreamDataSource,
		NewGatewayPluginKeyAuthDataSource,
		NewGatewayPluginKeyAuthEncDataSource,
		NewGatewayPluginLdapAuthDataSource,
		NewGatewayPluginLdapAuthAdvancedDataSource,
		NewGatewayPluginLogglyDataSource,
		NewGatewayPluginMockingDataSource,
		NewGatewayPluginMtlsAuthDataSource,
		NewGatewayPluginOasValidationDataSource,
		NewGatewayPluginOauth2DataSource,
		NewGatewayPluginOauth2IntrospectionDataSource,
		NewGatewayPluginOpaDataSource,
		NewGatewayPluginOpenidConnectDataSource,
		NewGatewayPluginOpentelemetryDataSource,
		NewGatewayPluginPostFunctionDataSource,
		NewGatewayPluginPreFunctionDataSource,
		NewGatewayPluginPrometheusDataSource,
		NewGatewayPluginProxyCacheDataSource,
		NewGatewayPluginProxyCacheAdvancedDataSource,
		NewGatewayPluginRateLimitingDataSource,
		NewGatewayPluginRateLimitingAdvancedDataSource,
		NewGatewayPluginRedirectDataSource,
		NewGatewayPluginRequestSizeLimitingDataSource,
		NewGatewayPluginRequestTerminationDataSource,
		NewGatewayPluginRequestTransformerDataSource,
		NewGatewayPluginRequestTransformerAdvancedDataSource,
		NewGatewayPluginRequestValidatorDataSource,
		NewGatewayPluginResponseRatelimitingDataSource,
		NewGatewayPluginResponseTransformerDataSource,
		NewGatewayPluginResponseTransformerAdvancedDataSource,
		NewGatewayPluginRouteByHeaderDataSource,
		NewGatewayPluginRouteTransformerAdvancedDataSource,
		NewGatewayPluginSamlDataSource,
		NewGatewayPluginServiceProtectionDataSource,
		NewGatewayPluginSessionDataSource,
		NewGatewayPluginStandardWebhooksDataSource,
		NewGatewayPluginStatsdDataSource,
		NewGatewayPluginStatsdAdvancedDataSource,
		NewGatewayPluginSyslogDataSource,
		NewGatewayPluginTCPLogDataSource,
		NewGatewayPluginTLSHandshakeModifierDataSource,
		NewGatewayPluginTLSMetadataHeadersDataSource,
		NewGatewayPluginUDPLogDataSource,
		NewGatewayPluginUpstreamOauthDataSource,
		NewGatewayPluginUpstreamTimeoutDataSource,
		NewGatewayPluginVaultAuthDataSource,
		NewGatewayPluginWebsocketSizeLimitDataSource,
		NewGatewayPluginWebsocketValidatorDataSource,
		NewGatewayPluginXMLThreatProtectionDataSource,
		NewGatewayPluginZipkinDataSource,
		NewGatewayRouteDataSource,
		NewGatewayRouteExpressionDataSource,
		NewGatewayServiceDataSource,
		NewGatewaySNIDataSource,
		NewGatewayTargetDataSource,
		NewGatewayUpstreamDataSource,
		NewGatewayVaultDataSource,
		NewMeshControlPlaneDataSource,
		NewPortalAppearanceDataSource,
		NewPortalAuthDataSource,
		NewPortalListDataSource,
		NewPortalProductVersionDataSource,
		NewPortalTeamDataSource,
		NewServerlessCloudGatewayDataSource,
		NewSystemAccountDataSource,
		NewSystemAccountAccessTokenDataSource,
		NewTeamDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &KonnectProvider{
			version: version,
		}
	}
}
