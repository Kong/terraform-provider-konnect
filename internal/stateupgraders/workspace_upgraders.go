// Per-resource `workspace` backfill upgraders, consolidated into one file
// instead of Speakeasy's usual one-per-entity layout. See .genignore.
package stateupgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func GatewayaclStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_acl", req, resp, defaultWorkspaceIfMissing)
}

func GatewaybasicauthStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_basic_auth", req, resp, defaultWorkspaceIfMissing)
}

func GatewaycacertificateStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_ca_certificate", req, resp, defaultWorkspaceIfMissing)
}

func GatewaycertificateStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_certificate", req, resp, defaultWorkspaceIfMissing)
}

func GatewayconsumerStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_consumer", req, resp, defaultWorkspaceIfMissing)
}

func GatewayconsumergroupStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_consumer_group", req, resp, defaultWorkspaceIfMissing)
}

func GatewayconsumergroupmemberStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_consumer_group_member", req, resp, defaultWorkspaceIfMissing)
}

func GatewayhmacauthStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_hmac_auth", req, resp, defaultWorkspaceIfMissing)
}

func GatewayjwtStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_jwt", req, resp, defaultWorkspaceIfMissing)
}

func GatewaykeyStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_key", req, resp, defaultWorkspaceIfMissing)
}

func GatewaykeyauthStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_key_auth", req, resp, defaultWorkspaceIfMissing)
}

func GatewaykeysetStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_key_set", req, resp, defaultWorkspaceIfMissing)
}

func GatewaymtlsauthStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_mtls_auth", req, resp, defaultWorkspaceIfMissing)
}

func GatewaypartialStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_partial", req, resp, defaultWorkspaceIfMissing)
}

func GatewayrouteStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_route", req, resp, defaultWorkspaceIfMissing)
}

func GatewayrouteexpressionStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_route_expression", req, resp, defaultWorkspaceIfMissing)
}

func GatewayserviceStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_service", req, resp, defaultWorkspaceIfMissing)
}

func GatewaysniStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_sni", req, resp, defaultWorkspaceIfMissing)
}

func GatewaytargetStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_target", req, resp, defaultWorkspaceIfMissing)
}

func GatewayupstreamStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_upstream", req, resp, defaultWorkspaceIfMissing)
}

func GatewayvaultStateUpgraderV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	upgradeToCurrentSchema("konnect_gateway_vault", req, resp, defaultWorkspaceIfMissing)
}
