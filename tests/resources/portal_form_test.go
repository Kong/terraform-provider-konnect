package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestPortalForm(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		const form = "konnect_portal_form.test_form"
		const dev = "konnect_portal_developer.test_developer"
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestNameDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(form, plancheck.ResourceActionCreate),
							plancheck.ExpectResourceAction(dev, plancheck.ResourceActionCreate),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet(form, "id"),
						resource.TestCheckResourceAttr(form, "developer_registration.status", "published"),
						resource.TestCheckResourceAttr(form, "developer_registration.fields.2.text.name", "custom"),
						resource.TestCheckResourceAttr(form, "developer_registration.fields.2.text.label", "custom"),
						resource.TestCheckResourceAttr(form, "developer_registration.fields.2.text.required", "true"),
						resource.TestCheckResourceAttrSet(dev, "id"),
						resource.TestCheckResourceAttr(dev, "email", "tftestformdev@example.com"),
						resource.TestCheckResourceAttr(dev, "status", "approved"),
						resource.TestCheckResourceAttrWith(dev, "additional_data", func(v string) error {
							var m map[string]any
							if err := json.Unmarshal([]byte(v), &m); err != nil {
								return err
							}
							if m["custom"] != "tftestcustomvalue" {
								return fmt.Errorf("unexpected additional_data.custom: %v", m["custom"])
							}
							return nil
						}),
					),
				},
				{
					Config:          providerConfigUs,
					ConfigDirectory: config.TestStepDirectory(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(form, plancheck.ResourceActionUpdate),
							plancheck.ExpectResourceAction(dev, plancheck.ResourceActionNoop),
						},
					},
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(form, "developer_registration.fields.2.text.label", "custom updated"),
						resource.TestCheckResourceAttr(form, "developer_registration.fields.2.text.required", "true"),
					),
				},
			},
		})
	})
}
