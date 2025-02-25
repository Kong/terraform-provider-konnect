// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
	"time"
)

func (r *APIProductDataSourceModel) RefreshFromSharedAPIProduct(resp *shared.APIProduct) {
	if resp != nil {
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.Description = types.StringPointerValue(resp.Description)
		r.ID = types.StringValue(resp.ID)
		if len(resp.Labels) > 0 {
			r.Labels = make(map[string]types.String, len(resp.Labels))
			for key, value := range resp.Labels {
				r.Labels[key] = types.StringValue(value)
			}
		}
		r.Name = types.StringValue(resp.Name)
		r.PortalIds = make([]types.String, 0, len(resp.PortalIds))
		for _, v := range resp.PortalIds {
			r.PortalIds = append(r.PortalIds, types.StringValue(v))
		}
		r.Portals = []tfTypes.APIProductPortal{}
		if len(r.Portals) > len(resp.Portals) {
			r.Portals = r.Portals[:len(resp.Portals)]
		}
		for portalsCount, portalsItem := range resp.Portals {
			var portals1 tfTypes.APIProductPortal
			portals1.PortalID = types.StringValue(portalsItem.PortalID)
			portals1.PortalName = types.StringValue(portalsItem.PortalName)
			if portalsCount+1 > len(r.Portals) {
				r.Portals = append(r.Portals, portals1)
			} else {
				r.Portals[portalsCount].PortalID = portals1.PortalID
				r.Portals[portalsCount].PortalName = portals1.PortalName
			}
		}
		if len(resp.PublicLabels) > 0 {
			r.PublicLabels = make(map[string]types.String, len(resp.PublicLabels))
			for key1, value1 := range resp.PublicLabels {
				r.PublicLabels[key1] = types.StringValue(value1)
			}
		}
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
		r.VersionCount = types.NumberValue(big.NewFloat(float64(resp.VersionCount)))
	}
}
