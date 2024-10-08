---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_cloud_gateway_transit_gateway Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  CloudGatewayTransitGateway Resource
---

# konnect_cloud_gateway_transit_gateway (Resource)

CloudGatewayTransitGateway Resource

## Example Usage

```terraform
resource "konnect_cloud_gateway_transit_gateway" "my_cloudgatewaytransitgateway" {
  cidr_blocks = [
    "...",
  ]
  name       = "us-east-2 transit gateway"
  network_id = "36ae63d3-efd1-4bec-b246-62aa5d3f5695"
  transit_gateway_attachment_config = {
    aws_transit_gateway_attachment_config = {
      kind               = "aws-transit-gateway-attachment"
      ram_share_arn      = "...my_ram_share_arn..."
      transit_gateway_id = "...my_transit_gateway_id..."
    }
  }
  transit_gateway_id = "0850820b-d153-4a2a-b9be-7d2204779139"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cidr_blocks` (List of String) CIDR blocks for constructing a route table for the transit gateway, when attaching to the owning
network.

Requires replacement if changed.
- `name` (String) Human-readable name of the transit gateway. Requires replacement if changed.
- `network_id` (String) The network to operate on. Requires replacement if changed.
- `transit_gateway_attachment_config` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--transit_gateway_attachment_config))

### Optional

- `dns_config` (Attributes List) List of mappings from remote DNS server IP address sets to proxied internal domains, for a transit gateway
attachment.

Requires replacement if changed. (see [below for nested schema](#nestedatt--dns_config))

### Read-Only

- `created_at` (String) An RFC-3339 timestamp representation of transit gateway creation date.
- `entity_version` (Number) Monotonically-increasing version count of the transit gateway, to indicate the order of updates to the
transit gateway.
- `id` (String) The ID of this resource.
- `state` (String) State of the transit gateway. must be one of ["created", "initializing", "ready", "terminating", "terminated"]
- `updated_at` (String) An RFC-3339 timestamp representation of transit gateway update date.

<a id="nestedatt--transit_gateway_attachment_config"></a>
### Nested Schema for `transit_gateway_attachment_config`

Optional:

- `aws_transit_gateway_attachment_config` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--transit_gateway_attachment_config--aws_transit_gateway_attachment_config))

<a id="nestedatt--transit_gateway_attachment_config--aws_transit_gateway_attachment_config"></a>
### Nested Schema for `transit_gateway_attachment_config.aws_transit_gateway_attachment_config`

Optional:

- `kind` (String) Requires replacement if changed. ; Not Null; must be one of ["aws-transit-gateway-attachment"]
- `ram_share_arn` (String) Resource Share ARN to verify request to create transit gateway attachment. Requires replacement if changed. ; Not Null
- `transit_gateway_id` (String) AWS Transit Gateway ID to create attachment to. Requires replacement if changed. ; Not Null



<a id="nestedatt--dns_config"></a>
### Nested Schema for `dns_config`

Optional:

- `domain_proxy_list` (List of String) Internal domain names to proxy for DNS resolution from the listed remote DNS server IP addresses,
for a transit gateway.

Requires replacement if changed. ; Not Null
- `remote_dns_server_ip_addresses` (List of String) Remote DNS Server IP Addresses to connect to for resolving internal DNS via a transit gateway. Requires replacement if changed. ; Not Null

## Import

Import is supported using the following syntax:

```shell
terraform import konnect_cloud_gateway_transit_gateway.my_konnect_cloud_gateway_transit_gateway "{ \"network_id\": \"36ae63d3-efd1-4bec-b246-62aa5d3f5695\",  \"transit_gateway_id\": \"0850820b-d153-4a2a-b9be-7d2204779139\"}"
```
