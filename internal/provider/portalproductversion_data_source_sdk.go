// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"time"
)

func (r *PortalProductVersionDataSourceModel) RefreshFromSharedPortalProductVersion(resp *shared.PortalProductVersion) {
	if resp != nil {
		r.ApplicationRegistrationEnabled = types.BoolValue(resp.ApplicationRegistrationEnabled)
		r.AuthStrategies = []tfTypes.AuthStrategy{}
		if len(r.AuthStrategies) > len(resp.AuthStrategies) {
			r.AuthStrategies = r.AuthStrategies[:len(resp.AuthStrategies)]
		}
		for authStrategiesCount, authStrategiesItem := range resp.AuthStrategies {
			var authStrategies1 tfTypes.AuthStrategy
			if authStrategiesItem.AuthStrategyClientCredentials != nil {
				authStrategies1.ClientCredentials = &tfTypes.AuthStrategyClientCredentials{}
				authStrategies1.ClientCredentials.AuthMethods = []types.String{}
				for _, v := range authStrategiesItem.AuthStrategyClientCredentials.AuthMethods {
					authStrategies1.ClientCredentials.AuthMethods = append(authStrategies1.ClientCredentials.AuthMethods, types.StringValue(v))
				}
				authStrategies1.ClientCredentials.CredentialType = types.StringValue(string(authStrategiesItem.AuthStrategyClientCredentials.CredentialType))
				authStrategies1.ClientCredentials.ID = types.StringValue(authStrategiesItem.AuthStrategyClientCredentials.ID)
				authStrategies1.ClientCredentials.Name = types.StringValue(authStrategiesItem.AuthStrategyClientCredentials.Name)
			}
			if authStrategiesItem.AuthStrategyKeyAuth != nil {
				authStrategies1.KeyAuth = &tfTypes.AuthStrategyKeyAuth{}
				authStrategies1.KeyAuth.CredentialType = types.StringValue(string(authStrategiesItem.AuthStrategyKeyAuth.CredentialType))
				authStrategies1.KeyAuth.ID = types.StringValue(authStrategiesItem.AuthStrategyKeyAuth.ID)
				authStrategies1.KeyAuth.Name = types.StringValue(authStrategiesItem.AuthStrategyKeyAuth.Name)
			}
			if authStrategiesCount+1 > len(r.AuthStrategies) {
				r.AuthStrategies = append(r.AuthStrategies, authStrategies1)
			} else {
				r.AuthStrategies[authStrategiesCount].ClientCredentials = authStrategies1.ClientCredentials
				r.AuthStrategies[authStrategiesCount].KeyAuth = authStrategies1.KeyAuth
			}
		}
		r.AutoApproveRegistration = types.BoolValue(resp.AutoApproveRegistration)
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.Deprecated = types.BoolValue(resp.Deprecated)
		r.ID = types.StringValue(resp.ID)
		r.ProductVersionID = types.StringValue(resp.ProductVersionID)
		r.PublishStatus = types.StringValue(string(resp.PublishStatus))
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	}
}
