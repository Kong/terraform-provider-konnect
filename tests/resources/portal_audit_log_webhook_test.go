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
	portalAuditLogWebhookPortal = `
		resource "konnect_portal" "my_portal" {
		  force_destroy = "true"
		  name          = "tf-test-portal-audit-log-webhook"
		}
	`

	portalAuditLogDestinationPrimary = `
		resource "konnect_audit_log_destination" "primary" {
		  authorization         = "Bearer sometoken"
		  endpoint              = "https://example.com/audit-logs-primary"
		  log_format            = "cef"
		  name                  = "tf-test-portal-audit-log-destination-primary"
		  skip_ssl_verification = false
		}
	`

	portalAuditLogDestinationSecondary = `
		resource "konnect_audit_log_destination" "secondary" {
		  authorization         = "Bearer sometoken"
		  endpoint              = "https://example.com/audit-logs-secondary"
		  log_format            = "json"
		  name                  = "tf-test-portal-audit-log-destination-secondary"
		  skip_ssl_verification = false
		}
	`

	portalAuditLogWebhook = `
		resource "konnect_portal_audit_log_webhook" "my_portal_audit_log_webhook" {
		  audit_log_destination_id = konnect_audit_log_destination.primary.id
		  enabled                  = true
		  portal_id                = konnect_portal.my_portal.id
		}
	`
)

func TestPortalAuditLogWebhook(t *testing.T) {
	serverHost, serverPort, serverScheme := providerConfigFromEnv()
	providerConfigTemplate := "%s://%s:%d"

	t.Run("CRUD", func(t *testing.T) {
		builder := hclbuilder.NewWithProvider(
			hclbuilder.Konnect,
			fmt.Sprintf(providerConfigTemplate, serverScheme, serverHost, serverPort),
		)
		builder.ProviderProperty = hclbuilder.Konnect

		portal, err := hclbuilder.FromString(portalAuditLogWebhookPortal)
		require.NoError(t, err)

		primaryDestination, err := hclbuilder.FromString(portalAuditLogDestinationPrimary)
		require.NoError(t, err)

		secondaryDestination, err := hclbuilder.FromString(portalAuditLogDestinationSecondary)
		require.NoError(t, err)

		webhookInitial, err := hclbuilder.FromString(portalAuditLogWebhook)
		require.NoError(t, err)

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.
						Upsert(portal).
						Upsert(primaryDestination).
						Upsert(secondaryDestination).
						Upsert(webhookInitial).
						Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_portal_audit_log_webhook.my_portal_audit_log_webhook",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_portal_audit_log_webhook.my_portal_audit_log_webhook",
							"enabled",
							"true",
						),
						resource.TestCheckResourceAttrPair(
							"konnect_portal_audit_log_webhook.my_portal_audit_log_webhook",
							"portal_id",
							"konnect_portal.my_portal",
							"id",
						),
						resource.TestCheckResourceAttrPair(
							"konnect_portal_audit_log_webhook.my_portal_audit_log_webhook",
							"audit_log_destination_id",
							"konnect_audit_log_destination.primary",
							"id",
						),
					),
				},
				{
					Config: builder.
						Upsert(portal).
						Upsert(primaryDestination).
						Upsert(secondaryDestination).
						Upsert(webhookInitial).
						Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: builder.
						Upsert(portal).
						Upsert(primaryDestination).
						Upsert(secondaryDestination).
						Upsert(
							webhookInitial.
								AddAttribute("audit_log_destination_id", secondaryDestination.ResourcePath()+".id").
								AddAttribute("enabled", "false"),
						).
						Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_portal_audit_log_webhook.my_portal_audit_log_webhook",
								plancheck.ResourceActionUpdate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_portal_audit_log_webhook.my_portal_audit_log_webhook",
							"enabled",
							"false",
						),
						resource.TestCheckResourceAttrPair(
							"konnect_portal_audit_log_webhook.my_portal_audit_log_webhook",
							"portal_id",
							"konnect_portal.my_portal",
							"id",
						),
						resource.TestCheckResourceAttrPair(
							"konnect_portal_audit_log_webhook.my_portal_audit_log_webhook",
							"audit_log_destination_id",
							"konnect_audit_log_destination.secondary",
							"id",
						),
					),
				},
				{
					Config: builder.
						Upsert(portal).
						Upsert(primaryDestination).
						Upsert(secondaryDestination).
						Upsert(webhookInitial).
						Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})
}
