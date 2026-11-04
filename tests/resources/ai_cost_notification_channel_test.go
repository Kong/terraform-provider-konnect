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
	aiCostEmailChannel = `
		resource "konnect_ai_cost_notification_channel" "test_email_channel" {
			email = {
				name = "tf-test-email-channel"

				recipients = [
					"finops@example.com",
					"platform-team@example.com",
				]

				labels = {
					team = "finops"
				}
			}
		}
	`

	aiCostEmailChannelUpdated = `
		resource "konnect_ai_cost_notification_channel" "test_email_channel" {
			email = {
				name = "tf-test-email-channel-updated"

				recipients = [
					"finops@example.com",
					"platform-team@example.com",
				]

				labels = {
					team = "finops"
				}
			}
		}
	`

	aiCostSlackChannel = `
		resource "konnect_ai_cost_notification_channel" "test_slack_channel" {
			slack = {
				name = "tf-test-slack-channel"

				webhook_url = "https://hooks.slack.com/services/T0000000/B0000000/XXXXXXXXXXXXXXXXXXXXXXXX"
			}
		}
	`

	aiCostSlackChannelUpdated = `
		resource "konnect_ai_cost_notification_channel" "test_slack_channel" {
			slack = {
				name = "tf-test-slack-channel-updated"

				webhook_url = "https://hooks.slack.com/services/T0000000/B0000000/XXXXXXXXXXXXXXXXXXXXXXXX"
			}
		}
	`
)

func TestAICostNotificationChannel(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()

	newConfig := func(t *testing.T, resources ...string) string {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf("%s://%s:%d", serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		for _, content := range resources {
			resource, err := hclbuilder.FromString(content)
			require.NoError(t, err)
			builder.Upsert(resource)
		}

		return builder.Build()
	}

	t.Run("Email Channel", func(t *testing.T) {
		const address = "konnect_ai_cost_notification_channel.test_email_channel"

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: newConfig(t, aiCostEmailChannel),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(address, plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet(address, "id"),
						resource.TestCheckResourceAttr(address, "name", "tf-test-email-channel"),
						resource.TestCheckResourceAttr(address, "email.name", "tf-test-email-channel"),
						resource.TestCheckResourceAttr(address, "email.recipients.#", "2"),
						resource.TestCheckResourceAttr(address, "email.recipients.0", "finops@example.com"),
						resource.TestCheckResourceAttr(address, "email.labels.team", "finops"),
					),
				},
				{
					Config: newConfig(t, aiCostEmailChannel),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: newConfig(t, aiCostEmailChannelUpdated),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(address, plancheck.ResourceActionUpdate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						// The hoisted top-level `name` keeps its previous value on an update,
						// so the nested attribute is the one to assert.
						resource.TestCheckResourceAttr(address, "email.name", "tf-test-email-channel-updated"),
					),
				},
			},
		})
	})

	t.Run("Slack Channel", func(t *testing.T) {
		const address = "konnect_ai_cost_notification_channel.test_slack_channel"

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: newConfig(t, aiCostSlackChannel),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(address, plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet(address, "id"),
						resource.TestCheckResourceAttr(address, "name", "tf-test-slack-channel"),
						resource.TestCheckResourceAttr(address, "slack.name", "tf-test-slack-channel"),
						resource.TestCheckResourceAttrSet(address, "slack.webhook_url"),
					),
				},
				{
					// The docs say reads return the webhook URL with its token replaced by
					// `…`. This step fails if that ever reaches the state, because the
					// configured URL would then differ from it on every plan.
					Config: newConfig(t, aiCostSlackChannel),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(address, "slack.webhook_url", "https://hooks.slack.com/services/T0000000/B0000000/XXXXXXXXXXXXXXXXXXXXXXXX"),
					),
				},
				{
					Config: newConfig(t, aiCostSlackChannelUpdated),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(address, "slack.name", "tf-test-slack-channel-updated"),
					),
				},
			},
		})
	})
}
