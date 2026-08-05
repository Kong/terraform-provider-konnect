package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestDcrProvider(t *testing.T) {
	t.Parallel()

	t.Run("kong_identity", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "kong_identity.name", "my-dcr-provider"),
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "kong_identity.issuer", "https://issuer.example.com"),
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "kong_identity.labels.team", "platform"),
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "kong_identity.active", "false"),
						// The root attributes are hoisted from the kong_identity block.
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "name", "my-dcr-provider"),
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "issuer", "https://issuer.example.com"),
					),
				},
				{
					// Link an auth strategy: non-empty plan by design. active is still
					// false in state here, as this step's refresh predates the strategy.
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
				},
				{
					// Identical config to step 2: the first refresh to see active flip,
					// so it checks the new value lands in state without causing a diff.
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "kong_identity.active", "true"),
					),
				},
			},
		})
	})

	t.Run("http update changes name and labels in place", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "http.name", "my-dcr-provider"),
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "http.labels.team", "platform"),
						// api_key is write-only, so check the config value is kept in state.
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "http.dcr_config.api_key", "tfacctestapikey"),
					),
				},
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("konnect_dcr_provider.my_dcrprovider", plancheck.ResourceActionUpdate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "http.name", "my-dcr-provider-renamed"),
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "http.labels.team", "identity"),
						// root name is stale after an in-place update, since
						// UseHoistedValue only fires when the plan value is unknown (create).
						// The next refresh fixes it, see step 3. If the modifier is fixed,
						// change this to expect "my-dcr-provider-renamed".
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "name", "my-dcr-provider"),
					),
				},
				{
					// Identical config to step 2: catches a perpetual diff on api_key and
					// checks the refresh brings the hoisted root name up to date.
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "name", "my-dcr-provider-renamed"),
					),
				},
			},
		})
	})

	t.Run("kong_identity update to any field requires replacement", func(t *testing.T) {
		replacePlan := resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("konnect_dcr_provider.my_dcrprovider", plancheck.ResourceActionDestroyBeforeCreate),
			},
		}

		// Fails if the id repeats between consecutive steps, proving each replacement
		// really produced a new provider. AddStateValue must be called once per step:
		// CompareValue indexes what it appends by step number.
		idChanged := statecheck.CompareValue(compare.ValuesDiffer())
		idState := func() []statecheck.StateCheck {
			return []statecheck.StateCheck{
				idChanged.AddStateValue("konnect_dcr_provider.my_dcrprovider", tfjsonpath.New("id")),
			}
		}

		// Each step changes exactly one field of the previous step's config.
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:            providerConfigUs,
					ConfigDirectory:   config.TestStepDirectory(),
					ConfigStateChecks: idState(),
				},
				{
					// issuer
					Config:            providerConfigUs,
					ConfigDirectory:   config.TestStepDirectory(),
					ConfigPlanChecks:  replacePlan,
					ConfigStateChecks: idState(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "kong_identity.issuer", "https://issuer-changed.example.com"),
					),
				},
				{
					// name
					Config:            providerConfigUs,
					ConfigDirectory:   config.TestStepDirectory(),
					ConfigPlanChecks:  replacePlan,
					ConfigStateChecks: idState(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "kong_identity.name", "my-dcr-provider-renamed"),
					),
				},
				{
					// labels
					Config:            providerConfigUs,
					ConfigDirectory:   config.TestStepDirectory(),
					ConfigPlanChecks:  replacePlan,
					ConfigStateChecks: idState(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("konnect_dcr_provider.my_dcrprovider", "kong_identity.labels.team", "identity"),
					),
				},
			},
		})
	})
}
