// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshTraceDataSourceModel) RefreshFromSharedMeshTraceItem(resp *shared.MeshTraceItem) {
	if resp != nil {
		if resp.CreationTime != nil {
			r.CreationTime = types.StringValue(resp.CreationTime.Format(time.RFC3339Nano))
		} else {
			r.CreationTime = types.StringNull()
		}
		if len(resp.Labels) > 0 {
			r.Labels = make(map[string]types.String)
			for key, value := range resp.Labels {
				r.Labels[key] = types.StringValue(value)
			}
		}
		r.Mesh = types.StringPointerValue(resp.Mesh)
		if resp.ModificationTime != nil {
			r.ModificationTime = types.StringValue(resp.ModificationTime.Format(time.RFC3339Nano))
		} else {
			r.ModificationTime = types.StringNull()
		}
		r.Name = types.StringValue(resp.Name)
		if resp.Spec.Default == nil {
			r.Spec.Default = nil
		} else {
			r.Spec.Default = &tfTypes.MeshTraceItemDefault{}
			r.Spec.Default.Backends = []tfTypes.MeshTraceItemBackends{}
			if len(r.Spec.Default.Backends) > len(resp.Spec.Default.Backends) {
				r.Spec.Default.Backends = r.Spec.Default.Backends[:len(resp.Spec.Default.Backends)]
			}
			for backendsCount, backendsItem := range resp.Spec.Default.Backends {
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
				if backendsCount+1 > len(r.Spec.Default.Backends) {
					r.Spec.Default.Backends = append(r.Spec.Default.Backends, backends1)
				} else {
					r.Spec.Default.Backends[backendsCount].Datadog = backends1.Datadog
					r.Spec.Default.Backends[backendsCount].OpenTelemetry = backends1.OpenTelemetry
					r.Spec.Default.Backends[backendsCount].Type = backends1.Type
					r.Spec.Default.Backends[backendsCount].Zipkin = backends1.Zipkin
				}
			}
			if resp.Spec.Default.Sampling == nil {
				r.Spec.Default.Sampling = nil
			} else {
				r.Spec.Default.Sampling = &tfTypes.MeshTraceItemSampling{}
				if resp.Spec.Default.Sampling.Client == nil {
					r.Spec.Default.Sampling.Client = nil
				} else {
					r.Spec.Default.Sampling.Client = &tfTypes.Mode{}
					if resp.Spec.Default.Sampling.Client.Integer != nil {
						r.Spec.Default.Sampling.Client.Integer = types.Int64PointerValue(resp.Spec.Default.Sampling.Client.Integer)
					}
					if resp.Spec.Default.Sampling.Client.Str != nil {
						r.Spec.Default.Sampling.Client.Str = types.StringPointerValue(resp.Spec.Default.Sampling.Client.Str)
					}
				}
				if resp.Spec.Default.Sampling.Overall == nil {
					r.Spec.Default.Sampling.Overall = nil
				} else {
					r.Spec.Default.Sampling.Overall = &tfTypes.Mode{}
					if resp.Spec.Default.Sampling.Overall.Integer != nil {
						r.Spec.Default.Sampling.Overall.Integer = types.Int64PointerValue(resp.Spec.Default.Sampling.Overall.Integer)
					}
					if resp.Spec.Default.Sampling.Overall.Str != nil {
						r.Spec.Default.Sampling.Overall.Str = types.StringPointerValue(resp.Spec.Default.Sampling.Overall.Str)
					}
				}
				if resp.Spec.Default.Sampling.Random == nil {
					r.Spec.Default.Sampling.Random = nil
				} else {
					r.Spec.Default.Sampling.Random = &tfTypes.Mode{}
					if resp.Spec.Default.Sampling.Random.Integer != nil {
						r.Spec.Default.Sampling.Random.Integer = types.Int64PointerValue(resp.Spec.Default.Sampling.Random.Integer)
					}
					if resp.Spec.Default.Sampling.Random.Str != nil {
						r.Spec.Default.Sampling.Random.Str = types.StringPointerValue(resp.Spec.Default.Sampling.Random.Str)
					}
				}
			}
			r.Spec.Default.Tags = []tfTypes.Tags{}
			if len(r.Spec.Default.Tags) > len(resp.Spec.Default.Tags) {
				r.Spec.Default.Tags = r.Spec.Default.Tags[:len(resp.Spec.Default.Tags)]
			}
			for tagsCount, tagsItem := range resp.Spec.Default.Tags {
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
				if tagsCount+1 > len(r.Spec.Default.Tags) {
					r.Spec.Default.Tags = append(r.Spec.Default.Tags, tags1)
				} else {
					r.Spec.Default.Tags[tagsCount].Header = tags1.Header
					r.Spec.Default.Tags[tagsCount].Literal = tags1.Literal
					r.Spec.Default.Tags[tagsCount].Name = tags1.Name
				}
			}
		}
		if resp.Spec.TargetRef == nil {
			r.Spec.TargetRef = nil
		} else {
			r.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
			if resp.Spec.TargetRef.Kind != nil {
				r.Spec.TargetRef.Kind = types.StringValue(string(*resp.Spec.TargetRef.Kind))
			} else {
				r.Spec.TargetRef.Kind = types.StringNull()
			}
			if len(resp.Spec.TargetRef.Labels) > 0 {
				r.Spec.TargetRef.Labels = make(map[string]types.String)
				for key1, value1 := range resp.Spec.TargetRef.Labels {
					r.Spec.TargetRef.Labels[key1] = types.StringValue(value1)
				}
			}
			r.Spec.TargetRef.Mesh = types.StringPointerValue(resp.Spec.TargetRef.Mesh)
			r.Spec.TargetRef.Name = types.StringPointerValue(resp.Spec.TargetRef.Name)
			r.Spec.TargetRef.Namespace = types.StringPointerValue(resp.Spec.TargetRef.Namespace)
			r.Spec.TargetRef.ProxyTypes = []types.String{}
			for _, v := range resp.Spec.TargetRef.ProxyTypes {
				r.Spec.TargetRef.ProxyTypes = append(r.Spec.TargetRef.ProxyTypes, types.StringValue(string(v)))
			}
			r.Spec.TargetRef.SectionName = types.StringPointerValue(resp.Spec.TargetRef.SectionName)
			if len(resp.Spec.TargetRef.Tags) > 0 {
				r.Spec.TargetRef.Tags = make(map[string]types.String)
				for key2, value2 := range resp.Spec.TargetRef.Tags {
					r.Spec.TargetRef.Tags[key2] = types.StringValue(value2)
				}
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}
}