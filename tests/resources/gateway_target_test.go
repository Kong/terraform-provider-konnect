package tests

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

const (
	gatewayControlPlane = `
		resource "konnect_gateway_control_plane" "tf_test_cp" {
			name = "tf-test-gateway-target-cp"
		}
	`

	gatewayUpstream = `
		resource "konnect_gateway_upstream" "tf_test_upstream" {
			name  = "tf-test-upstream"
			slots = 100
		}
	`

	gatewayTarget = `
		resource "konnect_gateway_target" "tf_test_target" {
			target = "one.example.com:9000"
			weight = 100
		}
	`
)

func TestGatewayTarget(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("Gateway Target", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		cp, err := hclbuilder.FromString(gatewayControlPlane)
		require.NoError(t, err)

		upstream, err := hclbuilder.FromString(gatewayUpstream)
		require.NoError(t, err)

		upstream.AddAttribute("control_plane_id", cp.ResourcePath()+".id")

		target, err := hclbuilder.FromString(gatewayTarget)
		require.NoError(t, err)

		target.AddAttribute("upstream_id", upstream.ResourcePath()+".id")
		target.AddAttribute("control_plane_id", cp.ResourcePath()+".id")

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Upsert(cp).Upsert(upstream).Upsert(target).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_gateway_target.tf_test_target", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_target.tf_test_target", "target", "one.example.com:9000"),
						resource.TestCheckResourceAttr("konnect_gateway_target.tf_test_target", "weight", "100"),
					),
				},
				{
					Config: builder.Upsert(cp).Upsert(upstream).Upsert(target).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("Gateway Target - Consecutive Apply Without Changes", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(hclbuilder.Konnect, fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort))
		builder.ProviderProperty = hclbuilder.Konnect

		cp, err := hclbuilder.FromString(gatewayControlPlane)
		require.NoError(t, err)

		upstream, err := hclbuilder.FromString(gatewayUpstream)
		require.NoError(t, err)

		upstream.AddAttribute("control_plane_id", cp.ResourcePath()+".id")

		target, err := hclbuilder.FromString(`
			resource "konnect_gateway_target" "tf_test_target_idempotent" {
				target   = "idempotent.example.com:9000"
				weight   = 100
				tags     = ["test", "idempotent"]
				failover = false
			}
		`)
		require.NoError(t, err)

		target.AddAttribute("upstream_id", upstream.ResourcePath()+".id")
		target.AddAttribute("control_plane_id", cp.ResourcePath()+".id")

		config := builder.Upsert(cp).Upsert(upstream).Upsert(target).Build()

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				// First apply - creates the resource
				{
					Config: config,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_gateway_target.tf_test_target_idempotent", plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_target.tf_test_target_idempotent", "target", "idempotent.example.com:9000"),
						resource.TestCheckResourceAttr("konnect_gateway_target.tf_test_target_idempotent", "weight", "100"),
						resource.TestCheckResourceAttr("konnect_gateway_target.tf_test_target_idempotent", "failover", "false"),
						resource.TestCheckResourceAttr("konnect_gateway_target.tf_test_target_idempotent", "tags.#", "2"),
					),
				},
				// Second apply - should produce no changes (idempotent)
				{
					Config: config,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_target.tf_test_target_idempotent", "target", "idempotent.example.com:9000"),
						resource.TestCheckResourceAttr("konnect_gateway_target.tf_test_target_idempotent", "weight", "100"),
					),
				},
				// Third apply - should still produce no changes (consecutive idempotent applies)
				{
					Config: config,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_gateway_target.tf_test_target_idempotent", "target", "idempotent.example.com:9000"),
						resource.TestCheckResourceAttr("konnect_gateway_target.tf_test_target_idempotent", "weight", "100"),
						resource.TestCheckResourceAttrSet("konnect_gateway_target.tf_test_target_idempotent", "id"),
					),
				},
			},
		})
	})
}
