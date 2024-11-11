# terraform-provider-konnect

This repository contains a Terraform provider for [Kong Konnect](https://konghq.com/products/kong-konnect/register?utm_source=github&utm_campaign=terraform&utm_content=terraform-provider-konnect).

## Capabilities

This provider can manage the following resources:

### Gateway Manager

- Gateway Hybrid Control Plane
- Gateway KIC Control Plane
- Gateway Control Plane Group
  - Control Plane Group Memberships
- Gateway Data Plane Client Certificates

### Cloud Gateways

- Cloud Gateway Configuration
- Custom Domain
- Network
- Transit Gateway
- Serverless Cloud Gateway

### Kong Gateway Entities

- Service
- Route
- Plugins
- Custom Plugins
- Consumers
- Consumer Groups
  - Consumer Group Members
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
- Custom Plugin Schema

### Mesh

- Mesh Control Plane

### API Products

- API Product
- API Product Version
- API Product Specification
- API Product Documentation

### Dev Portal

- Portal
- Portal Authentication
- Application Auth Strategy
- Portal Product Version
- Portal Appearance

### Organization

- System Account
- System Account Access Tokens
- System Account Roles
- System Account Teams
- Team
- Team Roles
- Team Users

## Usage

The provider can be installed from the Terraform registry. Before using the provider, you must configure a `personal_access_token`. If you are running in a non-US region, you must also set the `server_url` configuration option.

```hcl
terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
    }
  }
}

provider "konnect" {
  personal_access_token = "kpat_YOUR_PAT"
  server_url            = "https://us.api.konghq.com"
}
```

You may also configure the `provider` block using environment variables:

- `personal_access_token` = `KONNECT_TOKEN`
- `system_account_access_token` = `KONNECT_SPAT`
- `server_url` = `KONNECT_SERVER_URL`

e.g. `KONNECT_TOKEN=kpat_YOUR_PAT terraform apply`

## Examples

The examples directory contains sample usage for all supported resources. For a full list of supported parameters for each resource, see the [provider documentation](https://registry.terraform.io/providers/Kong/konnect/latest/docs).

The examples in this repo reference resources that are expected to exist in your manifests e.g. `konnect_gateway_control_plane.tfdemo.id`. Update the references to match the names that you have given your resources.

## Data Providers

We are currently working on official data source support (see [#10](https://github.com/Kong/terraform-provider-konnect/issues/10)). Until this is ready, you can use the `http` data source to call the Konnect API directly.

```hcl
data "http" "production_cp" {
  url    = "https://us.api.konghq.com/v2/control-planes?page%5Bsize%5D=1&page%5Bnumber%5D=1&filter%5Bname%5D%5Beq%5D=production"
  request_headers = {
    Accept           = "application/json"
    Authorization    = "Bearer ${var.konnect_personal_access_token}"
  }
}
```

<!-- No SDK Installation -->
<!-- No SDK Example Usage -->
<!-- No SDK Available Operations -->
<!-- Placeholder for Future Speakeasy SDK Sections -->
