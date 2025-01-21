// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *MeshListDataSourceModel) RefreshFromSharedMeshList(resp *shared.MeshList) {
	if resp != nil {
		r.Items = []tfTypes.MeshItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items1 tfTypes.MeshItem
			if itemsItem.Constraints == nil {
				items1.Constraints = nil
			} else {
				items1.Constraints = &tfTypes.Constraints{}
				if itemsItem.Constraints.DataplaneProxy == nil {
					items1.Constraints.DataplaneProxy = nil
				} else {
					items1.Constraints.DataplaneProxy = &tfTypes.DataplaneProxy{}
					items1.Constraints.DataplaneProxy.Requirements = []tfTypes.Requirements{}
					for requirementsCount, requirementsItem := range itemsItem.Constraints.DataplaneProxy.Requirements {
						var requirements1 tfTypes.Requirements
						if len(requirementsItem.Tags) > 0 {
							requirements1.Tags = make(map[string]types.String)
							for key, value := range requirementsItem.Tags {
								requirements1.Tags[key] = types.StringValue(value)
							}
						}
						if requirementsCount+1 > len(items1.Constraints.DataplaneProxy.Requirements) {
							items1.Constraints.DataplaneProxy.Requirements = append(items1.Constraints.DataplaneProxy.Requirements, requirements1)
						} else {
							items1.Constraints.DataplaneProxy.Requirements[requirementsCount].Tags = requirements1.Tags
						}
					}
					items1.Constraints.DataplaneProxy.Restrictions = []tfTypes.Requirements{}
					for restrictionsCount, restrictionsItem := range itemsItem.Constraints.DataplaneProxy.Restrictions {
						var restrictions1 tfTypes.Requirements
						if len(restrictionsItem.Tags) > 0 {
							restrictions1.Tags = make(map[string]types.String)
							for key1, value1 := range restrictionsItem.Tags {
								restrictions1.Tags[key1] = types.StringValue(value1)
							}
						}
						if restrictionsCount+1 > len(items1.Constraints.DataplaneProxy.Restrictions) {
							items1.Constraints.DataplaneProxy.Restrictions = append(items1.Constraints.DataplaneProxy.Restrictions, restrictions1)
						} else {
							items1.Constraints.DataplaneProxy.Restrictions[restrictionsCount].Tags = restrictions1.Tags
						}
					}
				}
			}
			if len(itemsItem.Labels) > 0 {
				items1.Labels = make(map[string]types.String)
				for key2, value2 := range itemsItem.Labels {
					items1.Labels[key2] = types.StringValue(value2)
				}
			}
			if itemsItem.Logging == nil {
				items1.Logging = nil
			} else {
				items1.Logging = &tfTypes.Logging{}
				items1.Logging.Backends = []tfTypes.Backends{}
				for backendsCount, backendsItem := range itemsItem.Logging.Backends {
					var backends1 tfTypes.Backends
					if backendsItem.Conf == nil {
						backends1.Conf = nil
					} else {
						backends1.Conf = &tfTypes.Metadata{}
					}
					backends1.Format = types.StringPointerValue(backendsItem.Format)
					backends1.Name = types.StringPointerValue(backendsItem.Name)
					backends1.Type = types.StringPointerValue(backendsItem.Type)
					if backendsCount+1 > len(items1.Logging.Backends) {
						items1.Logging.Backends = append(items1.Logging.Backends, backends1)
					} else {
						items1.Logging.Backends[backendsCount].Conf = backends1.Conf
						items1.Logging.Backends[backendsCount].Format = backends1.Format
						items1.Logging.Backends[backendsCount].Name = backends1.Name
						items1.Logging.Backends[backendsCount].Type = backends1.Type
					}
				}
				items1.Logging.DefaultBackend = types.StringPointerValue(itemsItem.Logging.DefaultBackend)
			}
			if itemsItem.MeshServices == nil {
				items1.MeshServices = nil
			} else {
				items1.MeshServices = &tfTypes.MeshServices{}
				if itemsItem.MeshServices.Mode == nil {
					items1.MeshServices.Mode = nil
				} else {
					items1.MeshServices.Mode = &tfTypes.Mode{}
					if itemsItem.MeshServices.Mode.Str != nil {
						items1.MeshServices.Mode.Str = types.StringPointerValue(itemsItem.MeshServices.Mode.Str)
					}
					if itemsItem.MeshServices.Mode.Integer != nil {
						items1.MeshServices.Mode.Integer = types.Int64PointerValue(itemsItem.MeshServices.Mode.Integer)
					}
				}
			}
			if itemsItem.Metrics == nil {
				items1.Metrics = nil
			} else {
				items1.Metrics = &tfTypes.Metrics{}
				items1.Metrics.Backends = []tfTypes.MeshItemBackends{}
				for backendsCount1, backendsItem1 := range itemsItem.Metrics.Backends {
					var backends3 tfTypes.MeshItemBackends
					if backendsItem1.Conf == nil {
						backends3.Conf = nil
					} else {
						backends3.Conf = &tfTypes.Metadata{}
					}
					backends3.Name = types.StringPointerValue(backendsItem1.Name)
					backends3.Type = types.StringPointerValue(backendsItem1.Type)
					if backendsCount1+1 > len(items1.Metrics.Backends) {
						items1.Metrics.Backends = append(items1.Metrics.Backends, backends3)
					} else {
						items1.Metrics.Backends[backendsCount1].Conf = backends3.Conf
						items1.Metrics.Backends[backendsCount1].Name = backends3.Name
						items1.Metrics.Backends[backendsCount1].Type = backends3.Type
					}
				}
				items1.Metrics.EnabledBackend = types.StringPointerValue(itemsItem.Metrics.EnabledBackend)
			}
			if itemsItem.Mtls == nil {
				items1.Mtls = nil
			} else {
				items1.Mtls = &tfTypes.Mtls{}
				items1.Mtls.Backends = []tfTypes.MeshItemMtlsBackends{}
				for backendsCount2, backendsItem2 := range itemsItem.Mtls.Backends {
					var backends5 tfTypes.MeshItemMtlsBackends
					if backendsItem2.Conf == nil {
						backends5.Conf = nil
					} else {
						backends5.Conf = &tfTypes.Metadata{}
					}
					if backendsItem2.DpCert == nil {
						backends5.DpCert = nil
					} else {
						backends5.DpCert = &tfTypes.DpCert{}
						if backendsItem2.DpCert.RequestTimeout == nil {
							backends5.DpCert.RequestTimeout = nil
						} else {
							backends5.DpCert.RequestTimeout = &tfTypes.RequestTimeout{}
							backends5.DpCert.RequestTimeout.Nanos = types.Int64PointerValue(backendsItem2.DpCert.RequestTimeout.Nanos)
							backends5.DpCert.RequestTimeout.Seconds = types.Int64PointerValue(backendsItem2.DpCert.RequestTimeout.Seconds)
						}
						if backendsItem2.DpCert.Rotation == nil {
							backends5.DpCert.Rotation = nil
						} else {
							backends5.DpCert.Rotation = &tfTypes.Rotation{}
							backends5.DpCert.Rotation.Expiration = types.StringPointerValue(backendsItem2.DpCert.Rotation.Expiration)
						}
					}
					if backendsItem2.Mode == nil {
						backends5.Mode = nil
					} else {
						backends5.Mode = &tfTypes.Mode{}
						if backendsItem2.Mode.Str != nil {
							backends5.Mode.Str = types.StringPointerValue(backendsItem2.Mode.Str)
						}
						if backendsItem2.Mode.Integer != nil {
							backends5.Mode.Integer = types.Int64PointerValue(backendsItem2.Mode.Integer)
						}
					}
					backends5.Name = types.StringPointerValue(backendsItem2.Name)
					if backendsItem2.RootChain == nil {
						backends5.RootChain = nil
					} else {
						backends5.RootChain = &tfTypes.RootChain{}
						if backendsItem2.RootChain.RequestTimeout == nil {
							backends5.RootChain.RequestTimeout = nil
						} else {
							backends5.RootChain.RequestTimeout = &tfTypes.RequestTimeout{}
							backends5.RootChain.RequestTimeout.Nanos = types.Int64PointerValue(backendsItem2.RootChain.RequestTimeout.Nanos)
							backends5.RootChain.RequestTimeout.Seconds = types.Int64PointerValue(backendsItem2.RootChain.RequestTimeout.Seconds)
						}
					}
					backends5.Type = types.StringPointerValue(backendsItem2.Type)
					if backendsCount2+1 > len(items1.Mtls.Backends) {
						items1.Mtls.Backends = append(items1.Mtls.Backends, backends5)
					} else {
						items1.Mtls.Backends[backendsCount2].Conf = backends5.Conf
						items1.Mtls.Backends[backendsCount2].DpCert = backends5.DpCert
						items1.Mtls.Backends[backendsCount2].Mode = backends5.Mode
						items1.Mtls.Backends[backendsCount2].Name = backends5.Name
						items1.Mtls.Backends[backendsCount2].RootChain = backends5.RootChain
						items1.Mtls.Backends[backendsCount2].Type = backends5.Type
					}
				}
				items1.Mtls.EnabledBackend = types.StringPointerValue(itemsItem.Mtls.EnabledBackend)
				items1.Mtls.SkipValidation = types.BoolPointerValue(itemsItem.Mtls.SkipValidation)
			}
			items1.Name = types.StringValue(itemsItem.Name)
			if itemsItem.Networking == nil {
				items1.Networking = nil
			} else {
				items1.Networking = &tfTypes.Networking{}
				if itemsItem.Networking.Outbound == nil {
					items1.Networking.Outbound = nil
				} else {
					items1.Networking.Outbound = &tfTypes.Outbound{}
					if itemsItem.Networking.Outbound.Passthrough == nil {
						items1.Networking.Outbound.Passthrough = nil
					} else {
						items1.Networking.Outbound.Passthrough = &tfTypes.Passthrough{}
						items1.Networking.Outbound.Passthrough.Value = types.BoolPointerValue(itemsItem.Networking.Outbound.Passthrough.Value)
					}
				}
			}
			if itemsItem.Routing == nil {
				items1.Routing = nil
			} else {
				items1.Routing = &tfTypes.Routing{}
				items1.Routing.DefaultForbidMeshExternalServiceAccess = types.BoolPointerValue(itemsItem.Routing.DefaultForbidMeshExternalServiceAccess)
				items1.Routing.LocalityAwareLoadBalancing = types.BoolPointerValue(itemsItem.Routing.LocalityAwareLoadBalancing)
				items1.Routing.ZoneEgress = types.BoolPointerValue(itemsItem.Routing.ZoneEgress)
			}
			items1.SkipCreatingInitialPolicies = []types.String{}
			for _, v := range itemsItem.SkipCreatingInitialPolicies {
				items1.SkipCreatingInitialPolicies = append(items1.SkipCreatingInitialPolicies, types.StringValue(v))
			}
			if itemsItem.Tracing == nil {
				items1.Tracing = nil
			} else {
				items1.Tracing = &tfTypes.Tracing{}
				items1.Tracing.Backends = []tfTypes.MeshItemTracingBackends{}
				for backendsCount3, backendsItem3 := range itemsItem.Tracing.Backends {
					var backends7 tfTypes.MeshItemTracingBackends
					if backendsItem3.Conf == nil {
						backends7.Conf = nil
					} else {
						backends7.Conf = &tfTypes.Metadata{}
					}
					backends7.Name = types.StringPointerValue(backendsItem3.Name)
					if backendsItem3.Sampling == nil {
						backends7.Sampling = nil
					} else {
						backends7.Sampling = &tfTypes.Sampling{}
						if backendsItem3.Sampling.Value != nil {
							backends7.Sampling.Value = types.NumberValue(big.NewFloat(float64(*backendsItem3.Sampling.Value)))
						} else {
							backends7.Sampling.Value = types.NumberNull()
						}
					}
					backends7.Type = types.StringPointerValue(backendsItem3.Type)
					if backendsCount3+1 > len(items1.Tracing.Backends) {
						items1.Tracing.Backends = append(items1.Tracing.Backends, backends7)
					} else {
						items1.Tracing.Backends[backendsCount3].Conf = backends7.Conf
						items1.Tracing.Backends[backendsCount3].Name = backends7.Name
						items1.Tracing.Backends[backendsCount3].Sampling = backends7.Sampling
						items1.Tracing.Backends[backendsCount3].Type = backends7.Type
					}
				}
				items1.Tracing.DefaultBackend = types.StringPointerValue(itemsItem.Tracing.DefaultBackend)
			}
			items1.Type = types.StringValue(itemsItem.Type)
			if itemsCount+1 > len(r.Items) {
				r.Items = append(r.Items, items1)
			} else {
				r.Items[itemsCount].Constraints = items1.Constraints
				r.Items[itemsCount].Labels = items1.Labels
				r.Items[itemsCount].Logging = items1.Logging
				r.Items[itemsCount].MeshServices = items1.MeshServices
				r.Items[itemsCount].Metrics = items1.Metrics
				r.Items[itemsCount].Mtls = items1.Mtls
				r.Items[itemsCount].Name = items1.Name
				r.Items[itemsCount].Networking = items1.Networking
				r.Items[itemsCount].Routing = items1.Routing
				r.Items[itemsCount].SkipCreatingInitialPolicies = items1.SkipCreatingInitialPolicies
				r.Items[itemsCount].Tracing = items1.Tracing
				r.Items[itemsCount].Type = items1.Type
			}
		}
		r.Next = types.StringPointerValue(resp.Next)
		if resp.Total != nil {
			r.Total = types.NumberValue(big.NewFloat(float64(*resp.Total)))
		} else {
			r.Total = types.NumberNull()
		}
	}
}
