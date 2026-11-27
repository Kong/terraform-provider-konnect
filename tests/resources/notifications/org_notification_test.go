package notifications

import (
	"fmt"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"

	"github.com/kong/terraform-provider-konnect/v3/tests/internal/acc"
)

func TestOrgNotificationDestinationAndSubscription(t *testing.T) {
	serverHost, serverPort, serverScheme := acc.ProviderConfigFromEnv()
	builder := hclbuilder.NewWithProvider(
		hclbuilder.Konnect,
		fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort),
	)
	builder.ProviderProperty = hclbuilder.Konnect

	destination, err := hclbuilder.FromString(`resource "konnect_org_notification_destination" "example" {
  http = {
    name = "tf-acceptance-example-com"
    type = "http"
    url  = "https://example.com"
  }
}
`)
	require.NoError(t, err)

	subscription, err := hclbuilder.FromString(`resource "konnect_org_notification_subscription" "kai" {
  event_type = "kai.feature.enablement"
  name = "tf-acceptance-kai-subscription"
  destination_ids = [konnect_org_notification_destination.example.id]
  enabled = true
  entities = ["*"]
  regions = ["*"]
}
`)
	require.NoError(t, err)

	config := builder.Upsert(destination).Upsert(subscription).Build()
	updatedConfig := builder.Upsert(destination.AddAttribute("http.name", `"tf-acceptance-example-com-updated"`)).Upsert(subscription).Build()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acc.ProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("konnect_org_notification_destination.example", plancheck.ResourceActionCreate),
						plancheck.ExpectResourceAction("konnect_org_notification_subscription.kai", plancheck.ResourceActionCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("konnect_org_notification_destination.example", "http.url", "https://example.com"),
					resource.TestCheckResourceAttrSet("konnect_org_notification_destination.example", "id"),
					resource.TestCheckResourceAttr("konnect_org_notification_subscription.kai", "event_type", "kai.feature.enablement"),
					resource.TestCheckResourceAttr("konnect_org_notification_subscription.kai", "enabled", "true"),
					resource.TestCheckResourceAttr("konnect_org_notification_subscription.kai", "destination_ids.#", "1"),
					resource.TestCheckResourceAttrPair(
						"konnect_org_notification_subscription.kai", "destination_ids.0",
						"konnect_org_notification_destination.example", "id",
					),
				),
			},
			{
				Config: config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
			{
				Config: updatedConfig,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("konnect_org_notification_destination.example", plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("konnect_org_notification_destination.example", "http.name", "tf-acceptance-example-com-updated"),
				),
			},
		},
	})
}
