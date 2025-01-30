// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/hooks"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/retry"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://global.api.konghq.com",
	"https://us.api.konghq.com",
	"https://eu.api.konghq.com",
	"https://au.api.konghq.com",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Konnect API: The Konnect platform API
//
// https://docs.konghq.com - Documentation for Kong Gateway and its APIs
type Konnect struct {
	ServerlessCloudGateways        *ServerlessCloudGateways
	Mesh                           *Mesh
	APIProducts                    *APIProducts
	APIProductDocumentation        *APIProductDocumentation
	APIProductVersions             *APIProductVersions
	APIProductVersionSpecification *APIProductVersionSpecification
	// Application Auth Strategies are sets of plugin configurations that represent how the gateway will perform authentication and authorization for a Product Version.
	// Called “Auth Strategy” for short in the context of portals/applications.
	// The plugins are synced to any Gateway Service that is currently linked or becomes linked to the Product Version.
	//
	AppAuthStrategies *AppAuthStrategies
	CloudGateways     *CloudGateways
	Networks          *Networks
	ControlPlanes     *ControlPlanes
	// A CA certificate object represents a trusted certificate authority.
	// These objects are used by Kong Gateway to verify the validity of a client or server certificate.
	CACertificates *CACertificates
	// A certificate object represents a public certificate, and can be optionally paired with the corresponding private key. These objects are used by Kong Gateway to handle SSL/TLS termination for encrypted requests, or for use as a trusted CA store when validating peer certificate of client/service.
	// <br><br>
	// Certificates are optionally associated with SNI objects to tie a cert/key pair to one or more hostnames.
	// <br><br>
	// If intermediate certificates are required in addition to the main certificate, they should be concatenated together into one string.
	//
	Certificates *Certificates
	// Consumer groups enable the organization and categorization of consumers (users or applications) within an API ecosystem.
	// By grouping consumers together, you eliminate the need to manage them individually, providing a scalable, efficient approach to managing configurations.
	ConsumerGroups *ConsumerGroups
	// The consumer object represents a consumer - or a user - of a service.
	// You can either rely on Kong Gateway as the primary datastore, or you can map the consumer list with your database to keep consistency between Kong Gateway and your existing primary datastore.
	//
	Consumers            *Consumers
	ACLs                 *ACLs
	BasicAuthCredentials *BasicAuthCredentials
	HMACAuthCredentials  *HMACAuthCredentials
	JWTs                 *JWTs
	APIKeys              *APIKeys
	MTLSAuthCredentials  *MTLSAuthCredentials
	// A JSON Web key set. Key sets are the preferred way to expose keys to plugins because they tell the plugin where to look for keys or have a scoping mechanism to restrict plugins to specific keys.
	//
	KeySets *KeySets
	// A key object holds a representation of asymmetric keys in various formats. When Kong Gateway or a Kong plugin requires a specific public or private key to perform certain operations, it can use this entity.
	//
	Keys *Keys
	// Custom Plugin Schemas
	CustomPluginSchemas *CustomPluginSchemas
	// A plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. Plugins let you add functionality to services that run behind a Kong Gateway instance, like authentication or rate limiting.
	// You can find more information about available plugins and which values each plugin accepts at the [Plugin Hub](https://docs.konghq.com/hub/).
	// <br><br>
	// When adding a plugin configuration to a service, the plugin will run on every request made by a client to that service. If a plugin needs to be tuned to different values for some specific consumers, you can do so by creating a separate plugin instance that specifies both the service and the consumer, through the service and consumer fields.
	// <br><br>
	// Plugins can be both [tagged and filtered by tags](https://docs.konghq.com/gateway/latest/admin-api/#tags).
	//
	Plugins *Plugins
	// Route entities define rules to match client requests. Each route is associated with a service, and a service may have multiple routes associated to it. Every request matching a given route will be proxied to the associated service. You need at least one matching rule that applies to the protocol being matched by the route.
	// <br><br>
	// The combination of routes and services, and the separation of concerns between them, offers a powerful routing mechanism with which it is possible to define fine-grained entrypoints in Kong Gateway leading to different upstream services of your infrastructure.
	// <br><br>
	// Depending on the protocol, one of the following attributes must be set:
	// <br>
	//
	// - `http`: At least one of `methods`, `hosts`, `headers`, or `paths`
	// - `https`: At least one of `methods`, `hosts`, `headers`, `paths`, or `snis`
	// - `tcp`: At least one of `sources` or `destinations`
	// - `tls`: at least one of `sources`, `destinations`, or `snis`
	// - `tls_passthrough`: set `snis`
	// - `grpc`: At least one of `hosts`, `headers`, or `paths`
	// - `grpcs`: At least one of `hosts`, `headers`, `paths`, or `snis`
	// - `ws`: At least one of `hosts`, `headers`, or `paths`
	// - `wss`: At least one of `hosts`, `headers`, `paths`, or `snis`
	//
	//
	//
	//
	//
	//
	//
	//   <br>
	//   A route can't have both `tls` and `tls_passthrough` protocols at same time.
	//   <br><br>
	//   Learn more about the router:
	// - [Configure routes using expressions](https://docs.konghq.com/gateway/latest/key-concepts/routes/expressions)
	// - [Router Expressions language reference](https://docs.konghq.com/gateway/latest/reference/router-expressions-language/)
	//
	Routes *Routes
	// Service entities are abstractions of your microservice interfaces or formal APIs. For example, a service could be a data transformation microservice or a billing API.
	// <br><br>
	// The main attribute of a service is the destination URL for proxying traffic. This URL can be set as a single string or by specifying its protocol, host, port and path individually.
	// <br><br>
	// Services are associated to routes, and a single service can have many routes associated with it. Routes are entrypoints in Kong Gateway which define rules to match client requests. Once a route is matched, Kong Gateway proxies the request to its associated service. See the [Proxy Reference](https://docs.konghq.com/gateway/latest/how-kong-works/routing-traffic/) for a detailed explanation of how Kong proxies traffic.
	// <br><br>
	// Services can be both [tagged and filtered by tags](https://docs.konghq.com/gateway/latest/admin-api/#tags).
	//
	Services *Services
	// An SNI object represents a many-to-one mapping of hostnames to a certificate.
	// <br><br>
	// A certificate object can have many hostnames associated with it. When Kong Gateway receives an SSL request, it uses the SNI field in the Client Hello to look up the certificate object based on the SNI associated with the certificate.
	SNIs *SNIs
	// The upstream object represents a virtual hostname and can be used to load balance incoming requests over multiple services (targets).
	// <br><br>
	// An upstream also includes a [health checker](https://docs.konghq.com/gateway/latest/how-kong-works/health-checks/), which can enable and disable targets based on their ability or inability to serve requests.
	// The configuration for the health checker is stored in the upstream object, and applies to all of its targets.
	Upstreams *Upstreams
	Targets   *Targets
	// Vault objects are used to configure different vault connectors for [managing secrets](https://docs.konghq.com/gateway/latest/kong-enterprise/secrets-management/).
	// Configuring a vault lets you reference secrets from other entities.
	// This allows for a proper separation of secrets and configuration and prevents secret sprawl.
	// <br><br>
	// For example, you could store a certificate and a key in a vault, then reference them from a certificate entity. This way, the certificate and key are not stored in the entity directly and are more secure.
	// <br><br>
	// Secrets rotation can be managed using [TTLs](https://docs.konghq.com/gateway/latest/kong-enterprise/secrets-management/advanced-usage/).
	//
	Vaults *Vaults
	// DP Certificates
	DPCertificates     *DPCertificates
	ControlPlaneGroups *ControlPlaneGroups
	// APIs related to Configuration of Konnect Developer Portals.
	Portals *Portals
	// APIs related to Konnect Developer Portal Appearance Settings.
	PortalAppearance *PortalAppearance
	// APIs related to Konnect Developer Portal Authentication Settings.
	PortalAuthSettings *PortalAuthSettings
	// Portal Product Versions hold metadata that describes how a Product Version is configured for a specific portal.
	//
	//   They contain:
	//   - Lifecyle and deprecation statuses
	//   - Application registration settings like auto approve or whether application registration is enabled
	//   - The authentication strategy that is enabled for Application Registration
	//
	PortalProductVersions *PortalProductVersions
	// APIs related to Konnect Developer Portal Developer Teams.
	PortalTeams                  *PortalTeams
	SystemAccounts               *SystemAccounts
	SystemAccountsAccessTokens   *SystemAccountsAccessTokens
	SystemAccountsRoles          *SystemAccountsRoles
	Teams                        *Teams
	Roles                        *Roles
	SystemAccountsTeamMembership *SystemAccountsTeamMembership
	TeamMembership               *TeamMembership

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Konnect)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Konnect) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Konnect) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Konnect) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Konnect) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Konnect) {
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *Konnect) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *Konnect) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *Konnect) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Konnect {
	sdk := &Konnect{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "2.0.0",
			SDKVersion:        "0.0.1",
			GenVersion:        "2.477.0",
			UserAgent:         "speakeasy-sdk/go 0.0.1 2.477.0 2.0.0 github.com/kong/terraform-provider-konnect/v2/internal/sdk",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.ServerlessCloudGateways = newServerlessCloudGateways(sdk.sdkConfiguration)

	sdk.Mesh = newMesh(sdk.sdkConfiguration)

	sdk.APIProducts = newAPIProducts(sdk.sdkConfiguration)

	sdk.APIProductDocumentation = newAPIProductDocumentation(sdk.sdkConfiguration)

	sdk.APIProductVersions = newAPIProductVersions(sdk.sdkConfiguration)

	sdk.APIProductVersionSpecification = newAPIProductVersionSpecification(sdk.sdkConfiguration)

	sdk.AppAuthStrategies = newAppAuthStrategies(sdk.sdkConfiguration)

	sdk.CloudGateways = newCloudGateways(sdk.sdkConfiguration)

	sdk.Networks = newNetworks(sdk.sdkConfiguration)

	sdk.ControlPlanes = newControlPlanes(sdk.sdkConfiguration)

	sdk.CACertificates = newCACertificates(sdk.sdkConfiguration)

	sdk.Certificates = newCertificates(sdk.sdkConfiguration)

	sdk.ConsumerGroups = newConsumerGroups(sdk.sdkConfiguration)

	sdk.Consumers = newConsumers(sdk.sdkConfiguration)

	sdk.ACLs = newACLs(sdk.sdkConfiguration)

	sdk.BasicAuthCredentials = newBasicAuthCredentials(sdk.sdkConfiguration)

	sdk.HMACAuthCredentials = newHMACAuthCredentials(sdk.sdkConfiguration)

	sdk.JWTs = newJWTs(sdk.sdkConfiguration)

	sdk.APIKeys = newAPIKeys(sdk.sdkConfiguration)

	sdk.MTLSAuthCredentials = newMTLSAuthCredentials(sdk.sdkConfiguration)

	sdk.KeySets = newKeySets(sdk.sdkConfiguration)

	sdk.Keys = newKeys(sdk.sdkConfiguration)

	sdk.CustomPluginSchemas = newCustomPluginSchemas(sdk.sdkConfiguration)

	sdk.Plugins = newPlugins(sdk.sdkConfiguration)

	sdk.Routes = newRoutes(sdk.sdkConfiguration)

	sdk.Services = newServices(sdk.sdkConfiguration)

	sdk.SNIs = newSNIs(sdk.sdkConfiguration)

	sdk.Upstreams = newUpstreams(sdk.sdkConfiguration)

	sdk.Targets = newTargets(sdk.sdkConfiguration)

	sdk.Vaults = newVaults(sdk.sdkConfiguration)

	sdk.DPCertificates = newDPCertificates(sdk.sdkConfiguration)

	sdk.ControlPlaneGroups = newControlPlaneGroups(sdk.sdkConfiguration)

	sdk.Portals = newPortals(sdk.sdkConfiguration)

	sdk.PortalAppearance = newPortalAppearance(sdk.sdkConfiguration)

	sdk.PortalAuthSettings = newPortalAuthSettings(sdk.sdkConfiguration)

	sdk.PortalProductVersions = newPortalProductVersions(sdk.sdkConfiguration)

	sdk.PortalTeams = newPortalTeams(sdk.sdkConfiguration)

	sdk.SystemAccounts = newSystemAccounts(sdk.sdkConfiguration)

	sdk.SystemAccountsAccessTokens = newSystemAccountsAccessTokens(sdk.sdkConfiguration)

	sdk.SystemAccountsRoles = newSystemAccountsRoles(sdk.sdkConfiguration)

	sdk.Teams = newTeams(sdk.sdkConfiguration)

	sdk.Roles = newRoles(sdk.sdkConfiguration)

	sdk.SystemAccountsTeamMembership = newSystemAccountsTeamMembership(sdk.sdkConfiguration)

	sdk.TeamMembership = newTeamMembership(sdk.sdkConfiguration)

	return sdk
}
