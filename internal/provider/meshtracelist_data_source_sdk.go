// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
	"time"
)

func (r *MeshTraceListDataSourceModel) RefreshFromSharedMeshTraceList(resp *shared.MeshTraceList) {
	if resp != nil {
		r.Items = []tfTypes.MeshTraceItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items1 tfTypes.MeshTraceItem
			if itemsItem.CreationTime != nil {
				items1.CreationTime = types.StringValue(itemsItem.CreationTime.Format(time.RFC3339Nano))
			} else {
				items1.CreationTime = types.StringNull()
			}
			if len(itemsItem.Labels) > 0 {
				items1.Labels = make(map[string]types.String)
				for key, value := range itemsItem.Labels {
					items1.Labels[key] = types.StringValue(value)
				}
			}
			items1.Mesh = types.StringPointerValue(itemsItem.Mesh)
			if itemsItem.ModificationTime != nil {
				items1.ModificationTime = types.StringValue(itemsItem.ModificationTime.Format(time.RFC3339Nano))
			} else {
				items1.ModificationTime = types.StringNull()
			}
			items1.Name = types.StringValue(itemsItem.Name)
			if itemsItem.Spec.Default == nil {
				items1.Spec.Default = nil
			} else {
				items1.Spec.Default = &tfTypes.MeshTraceItemDefault{}
				items1.Spec.Default.Backends = []tfTypes.MeshTraceItemBackends{}
				for backendsCount, backendsItem := range itemsItem.Spec.Default.Backends {
					var backends1 tfTypes.MeshTraceItemBackends
					if backendsItem.Datadog == nil {
						backends1.Datadog = nil
					} else {
						backends1.Datadog = &tfTypes.Datadog{}
						backends1.Datadog.SplitService = types.BoolPointerValue(backendsItem.Datadog.SplitService)
						backends1.Datadog.URL = types.StringValue(backendsItem.Datadog.URL)
					}
					if backendsItem.OpenTelemetry == nil {
						backends1.OpenTelemetry = nil
					} else {
						backends1.OpenTelemetry = &tfTypes.MeshTraceItemOpenTelemetry{}
						backends1.OpenTelemetry.Endpoint = types.StringValue(backendsItem.OpenTelemetry.Endpoint)
					}
					backends1.Type = types.StringValue(string(backendsItem.Type))
					if backendsItem.Zipkin == nil {
						backends1.Zipkin = nil
					} else {
						backends1.Zipkin = &tfTypes.Zipkin{}
						if backendsItem.Zipkin.APIVersion != nil {
							backends1.Zipkin.APIVersion = types.StringValue(string(*backendsItem.Zipkin.APIVersion))
						} else {
							backends1.Zipkin.APIVersion = types.StringNull()
						}
						backends1.Zipkin.SharedSpanContext = types.BoolPointerValue(backendsItem.Zipkin.SharedSpanContext)
						backends1.Zipkin.TraceId128bit = types.BoolPointerValue(backendsItem.Zipkin.TraceId128bit)
						backends1.Zipkin.URL = types.StringValue(backendsItem.Zipkin.URL)
					}
					if backendsCount+1 > len(items1.Spec.Default.Backends) {
						items1.Spec.Default.Backends = append(items1.Spec.Default.Backends, backends1)
					} else {
						items1.Spec.Default.Backends[backendsCount].Datadog = backends1.Datadog
						items1.Spec.Default.Backends[backendsCount].OpenTelemetry = backends1.OpenTelemetry
						items1.Spec.Default.Backends[backendsCount].Type = backends1.Type
						items1.Spec.Default.Backends[backendsCount].Zipkin = backends1.Zipkin
					}
				}
				if itemsItem.Spec.Default.Sampling == nil {
					items1.Spec.Default.Sampling = nil
				} else {
					items1.Spec.Default.Sampling = &tfTypes.MeshTraceItemSampling{}
					if itemsItem.Spec.Default.Sampling.Client == nil {
						items1.Spec.Default.Sampling.Client = nil
					} else {
						items1.Spec.Default.Sampling.Client = &tfTypes.Mode{}
						if itemsItem.Spec.Default.Sampling.Client.Integer != nil {
							items1.Spec.Default.Sampling.Client.Integer = types.Int64PointerValue(itemsItem.Spec.Default.Sampling.Client.Integer)
						}
						if itemsItem.Spec.Default.Sampling.Client.Str != nil {
							items1.Spec.Default.Sampling.Client.Str = types.StringPointerValue(itemsItem.Spec.Default.Sampling.Client.Str)
						}
					}
					if itemsItem.Spec.Default.Sampling.Overall == nil {
						items1.Spec.Default.Sampling.Overall = nil
					} else {
						items1.Spec.Default.Sampling.Overall = &tfTypes.Mode{}
						if itemsItem.Spec.Default.Sampling.Overall.Integer != nil {
							items1.Spec.Default.Sampling.Overall.Integer = types.Int64PointerValue(itemsItem.Spec.Default.Sampling.Overall.Integer)
						}
						if itemsItem.Spec.Default.Sampling.Overall.Str != nil {
							items1.Spec.Default.Sampling.Overall.Str = types.StringPointerValue(itemsItem.Spec.Default.Sampling.Overall.Str)
						}
					}
					if itemsItem.Spec.Default.Sampling.Random == nil {
						items1.Spec.Default.Sampling.Random = nil
					} else {
						items1.Spec.Default.Sampling.Random = &tfTypes.Mode{}
						if itemsItem.Spec.Default.Sampling.Random.Integer != nil {
							items1.Spec.Default.Sampling.Random.Integer = types.Int64PointerValue(itemsItem.Spec.Default.Sampling.Random.Integer)
						}
						if itemsItem.Spec.Default.Sampling.Random.Str != nil {
							items1.Spec.Default.Sampling.Random.Str = types.StringPointerValue(itemsItem.Spec.Default.Sampling.Random.Str)
						}
					}
				}
				items1.Spec.Default.Tags = []tfTypes.Tags{}
				for tagsCount, tagsItem := range itemsItem.Spec.Default.Tags {
					var tags1 tfTypes.Tags
					if tagsItem.Header == nil {
						tags1.Header = nil
					} else {
						tags1.Header = &tfTypes.Header{}
						tags1.Header.Default = types.StringPointerValue(tagsItem.Header.Default)
						tags1.Header.Name = types.StringValue(tagsItem.Header.Name)
					}
					tags1.Literal = types.StringPointerValue(tagsItem.Literal)
					tags1.Name = types.StringValue(tagsItem.Name)
					if tagsCount+1 > len(items1.Spec.Default.Tags) {
						items1.Spec.Default.Tags = append(items1.Spec.Default.Tags, tags1)
					} else {
						items1.Spec.Default.Tags[tagsCount].Header = tags1.Header
						items1.Spec.Default.Tags[tagsCount].Literal = tags1.Literal
						items1.Spec.Default.Tags[tagsCount].Name = tags1.Name
					}
				}
			}
			if itemsItem.Spec.TargetRef == nil {
				items1.Spec.TargetRef = nil
			} else {
				items1.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
				if itemsItem.Spec.TargetRef.Kind != nil {
					items1.Spec.TargetRef.Kind = types.StringValue(string(*itemsItem.Spec.TargetRef.Kind))
				} else {
					items1.Spec.TargetRef.Kind = types.StringNull()
				}
				if len(itemsItem.Spec.TargetRef.Labels) > 0 {
					items1.Spec.TargetRef.Labels = make(map[string]types.String)
					for key1, value1 := range itemsItem.Spec.TargetRef.Labels {
						items1.Spec.TargetRef.Labels[key1] = types.StringValue(value1)
					}
				}
				items1.Spec.TargetRef.Mesh = types.StringPointerValue(itemsItem.Spec.TargetRef.Mesh)
				items1.Spec.TargetRef.Name = types.StringPointerValue(itemsItem.Spec.TargetRef.Name)
				items1.Spec.TargetRef.Namespace = types.StringPointerValue(itemsItem.Spec.TargetRef.Namespace)
				items1.Spec.TargetRef.ProxyTypes = []types.String{}
				for _, v := range itemsItem.Spec.TargetRef.ProxyTypes {
					items1.Spec.TargetRef.ProxyTypes = append(items1.Spec.TargetRef.ProxyTypes, types.StringValue(string(v)))
				}
				items1.Spec.TargetRef.SectionName = types.StringPointerValue(itemsItem.Spec.TargetRef.SectionName)
				if len(itemsItem.Spec.TargetRef.Tags) > 0 {
					items1.Spec.TargetRef.Tags = make(map[string]types.String)
					for key2, value2 := range itemsItem.Spec.TargetRef.Tags {
						items1.Spec.TargetRef.Tags[key2] = types.StringValue(value2)
					}
				}
			}
			items1.Type = types.StringValue(string(itemsItem.Type))
			if itemsCount+1 > len(r.Items) {
				r.Items = append(r.Items, items1)
			} else {
				r.Items[itemsCount].CreationTime = items1.CreationTime
				r.Items[itemsCount].Labels = items1.Labels
				r.Items[itemsCount].Mesh = items1.Mesh
				r.Items[itemsCount].ModificationTime = items1.ModificationTime
				r.Items[itemsCount].Name = items1.Name
				r.Items[itemsCount].Spec = items1.Spec
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