# terraform-provider-konnect

This repository contains a Terraform provider for [Kong Konnect](https://konghq.com/products/kong-konnect/register?utm_source=github&utm_campaign=terraform&utm_content=terraform-provider-konnect).

## Capabilities

This provider can manage the following resources:

### Gateway Manager

- Gateway Hybrid Control Plane
- Gateway KIC Control Plane
- Gateway Control Plane Group
  - Control Plane Group Memberships

### Kong Gateway Entities

- Service
- Route
- Plugins
  - ACL
  - AI Prompt Decorator
  - AI Prompt Guard
  - AI Prompt Template
  - AI Proxy
  - AWS Lambda
  - Basic Auth
  - Correlation ID
  - CORS
  - File Log
  - IP Restriction
  - JQ
  - JWT
  - JWT Signer
  - KeyAuth
  - OAuth2
  - OpenID Connect
  - OpenTelemetry
  - Pre-function
  - Prometheus
  - Proxy Cache
  - Rate Limiting
  - Rate Limiting Advanced
  - Request Termination
  - Request Transformer
  - Request Transformer Advanced
  - Response Transformer
  - Response Transformer Advanced
  - SAML
- Consumers
- Credentials
  - Basic Auth
  - Key Auth
  - JWT
  - HMAC Auth
  - ACL
- Upstream
- Target
- Vault
- Certificate
- CA Certificate
- SNI
- Key
- Key Set

### Mesh

- Mesh Control Plane

### API Products

- API Product
- API Product Version

### Dev Portal

- Portal
- Portal Authentication
- Application Auth Strategy
- Portal Product Version

## Usage

The provider can be installed from the Terraform registry. Before using the provider, you must configure a `personal_access_token`. If you are running in a non-US region, you must also set the `server_url` configuration option.

```hcl
terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
      version = "TODO_ADD_VERSION"
    }
  }
}

provider "konnect" {
  personal_access_token = "kpat_YOUR_PAT"
  server_url            = "https://us.api.konghq.com/v2"
}
```

## Examples

The examples directory contains sample usage for all supported resources. For a full list of supported parameters for each resource, see the [Konnect API documentation](https://docs.konghq.com/api/).

The examples will reference resources that are expected to exist in your manifests e.g. `konnect_gateway_control_plane.tfdemo.id`. Update the references to match the names that you have given your resources.

### OneOf resources

When a resource has multiple representations the configuration must be placed inside a key that identifies which schema you want to use.

Let's work through a concrete example.

The `konnect_application_auth_strategy` resource supports multiple authentication strategies. The API specification marks `strategy_type` as the discriminator, which is how the API figures out which schema it should use for validation.

This means that the `konnect_application_auth_strategy` resource has a sub key that matches the values in `strategy_type`. If `strategy_type` is `key_auth`, the resource looks like the following:

```hcl
resource "konnect_application_auth_strategy" "my_applicationauthstrategy" {
  key_auth = {
    name          = "my-application-auth-strategy"
    # Other parameters omitted
  }
}
```

## FAQ

### Why do you only support specific plugins?

Kong Gateway's plugin entity supports any JSON blob in the `config` field, and will automatically fill in any default values. This does not work well with Terraform, which requires a predefined schema in order to track default values in the state file.

In order to support idempotent plan and apply operations, this provider has one resource per plugin.

If you require a plugin that is not currently supported, please open an issue.
