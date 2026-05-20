package tests

import (
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

func TestGatewayDataPlaneClientCertificate(t *testing.T) {
	t.Run("Gateway Data Plane Client Certificate", func(t *testing.T) {
		builder := hclbuilder.New()
		cp, err := hclbuilder.FromString(`
          resource "konnect_gateway_control_plane" "test_cp" {
            name = "tf-test-cp-dp-client-cert"
          }
        `)
		require.NoError(t, err)

		tlsKey, err := hclbuilder.FromString(`
          resource "tls_private_key" "client" {
            algorithm = "RSA"
            rsa_bits  = 2048
          }
        `)
		require.NoError(t, err)

		tlsCert, err := hclbuilder.FromString(`
          resource "tls_self_signed_cert" "client" {
            private_key_pem = tls_private_key.client.private_key_pem

            subject {
              common_name = "client.local"
            }

            validity_period_hours = 2

            allowed_uses = [
              "client_auth",
              "digital_signature",
              "key_encipherment",
            ]
          }
        `)
		require.NoError(t, err)

		cert, err := hclbuilder.FromString(`
          resource "konnect_gateway_data_plane_client_certificate" "my_dp_client_cert" {
            title            = "tf-test-dp-client-cert"
            control_plane_id = konnect_gateway_control_plane.test_cp.id
            cert             = tls_self_signed_cert.client.cert_pem
          }
        `)
		require.NoError(t, err)

		fullConfig := builder.
			Upsert(cp).
			Upsert(tlsKey).
			Upsert(tlsCert).
			Upsert(cert).
			Build()

		updatedConfig := builder.
			Upsert(cp).
			Upsert(tlsKey).
			Upsert(tlsCert).
			Upsert(cert.AddAttribute("title", `"tf-test-dp-client-cert-updated"`)).
			Build()

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			ExternalProviders: map[string]resource.ExternalProvider{
				"tls": {Source: "hashicorp/tls"},
			},
			Steps: []resource.TestStep{
				{
					Config: fullConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_gateway_data_plane_client_certificate.my_dp_client_cert",
								plancheck.ResourceActionCreate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_gateway_data_plane_client_certificate.my_dp_client_cert",
							"title",
							"tf-test-dp-client-cert",
						),
					),
				},
				{
					Config: updatedConfig,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(
								"konnect_gateway_data_plane_client_certificate.my_dp_client_cert",
								plancheck.ResourceActionUpdate,
							),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"konnect_gateway_data_plane_client_certificate.my_dp_client_cert",
							"title",
							"tf-test-dp-client-cert-updated",
						),
					),
				},
			},
		})
	})
}
