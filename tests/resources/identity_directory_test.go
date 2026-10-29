package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Identity directories are limited per org, so every resource that hangs off a
// directory is exercised against a single directory in one test.
func TestIdentityDirectory(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_identity_directory.my_identitydirectory", "name", "tf-test-identity-directory"),
						resource.TestCheckResourceAttr("konnect_identity_group.my_identitygroup", "name", "tftestgroup"),
						resource.TestCheckResourceAttr("konnect_identity_principal.my_identityprincipal", "display_name", "TF Test Principal"),
						resource.TestCheckResourceAttrPair(
							"konnect_identity_group_membership.my_identitygroupmembership", "member.id",
							"konnect_identity_principal.my_identityprincipal", "id",
						),
						resource.TestCheckResourceAttrSet("konnect_identity_group_api_key.my_identitygroupapikey", "id"),
						resource.TestCheckResourceAttrSet("konnect_identity_principal_api_key.my_identityprincipalapikey", "id"),
						resource.TestCheckResourceAttr("konnect_identity_principal_basic_auth.my_identityprincipalbasicauth", "username", "tftest-user"),
						resource.TestCheckResourceAttrSet("konnect_identity_principal_basic_auth_password.my_identityprincipalbasicauthpassword", "id"),
						resource.TestCheckResourceAttr("konnect_identity_principal_identity.my_identityprincipalidentity", "custom.key", "tftest-key"),
						resource.TestCheckResourceAttr("konnect_identity_trusted_idp.my_identitytrustedidp", "issuer_url", "https://tftest.example.com/oauth2/default"),
						resource.TestCheckResourceAttrSet("konnect_identity_vault_key.my_identityvaultkey", "kid"),
						resource.TestCheckResourceAttr("konnect_identity_vault_provider.my_identityvaultprovider", "name", "tftest-vault-provider"),
						resource.TestCheckResourceAttr("konnect_identity_vault_provider_credential.my_identityvaultprovidercredential", "status", "active"),
					),
				},
				{
					// Update description / display_name on the directory, group, principal and trusted IdP
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_identity_directory.my_identitydirectory", "description", "TF test directory updated"),
						resource.TestCheckResourceAttr("konnect_identity_group.my_identitygroup", "description", "TF test group updated"),
						resource.TestCheckResourceAttr("konnect_identity_principal.my_identityprincipal", "description", "TF test principal updated"),
						resource.TestCheckResourceAttr("konnect_identity_trusted_idp.my_identitytrustedidp", "display_name", "TF Test Trusted IdP Updated"),
					),
				},
			},
		})
	})
}
