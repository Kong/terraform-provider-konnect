// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayUpstreamResourceModel) ToSharedUpstreamInput() *shared.UpstreamInput {
	algorithm := new(shared.UpstreamAlgorithm)
	if !r.Algorithm.IsUnknown() && !r.Algorithm.IsNull() {
		*algorithm = shared.UpstreamAlgorithm(r.Algorithm.ValueString())
	} else {
		algorithm = nil
	}
	var clientCertificate *shared.UpstreamClientCertificate
	if r.ClientCertificate != nil {
		id := new(string)
		if !r.ClientCertificate.ID.IsUnknown() && !r.ClientCertificate.ID.IsNull() {
			*id = r.ClientCertificate.ID.ValueString()
		} else {
			id = nil
		}
		clientCertificate = &shared.UpstreamClientCertificate{
			ID: id,
		}
	}
	hashFallback := new(shared.HashFallback)
	if !r.HashFallback.IsUnknown() && !r.HashFallback.IsNull() {
		*hashFallback = shared.HashFallback(r.HashFallback.ValueString())
	} else {
		hashFallback = nil
	}
	hashFallbackHeader := new(string)
	if !r.HashFallbackHeader.IsUnknown() && !r.HashFallbackHeader.IsNull() {
		*hashFallbackHeader = r.HashFallbackHeader.ValueString()
	} else {
		hashFallbackHeader = nil
	}
	hashFallbackQueryArg := new(string)
	if !r.HashFallbackQueryArg.IsUnknown() && !r.HashFallbackQueryArg.IsNull() {
		*hashFallbackQueryArg = r.HashFallbackQueryArg.ValueString()
	} else {
		hashFallbackQueryArg = nil
	}
	hashFallbackURICapture := new(string)
	if !r.HashFallbackURICapture.IsUnknown() && !r.HashFallbackURICapture.IsNull() {
		*hashFallbackURICapture = r.HashFallbackURICapture.ValueString()
	} else {
		hashFallbackURICapture = nil
	}
	hashOn := new(shared.HashOn)
	if !r.HashOn.IsUnknown() && !r.HashOn.IsNull() {
		*hashOn = shared.HashOn(r.HashOn.ValueString())
	} else {
		hashOn = nil
	}
	hashOnCookie := new(string)
	if !r.HashOnCookie.IsUnknown() && !r.HashOnCookie.IsNull() {
		*hashOnCookie = r.HashOnCookie.ValueString()
	} else {
		hashOnCookie = nil
	}
	hashOnCookiePath := new(string)
	if !r.HashOnCookiePath.IsUnknown() && !r.HashOnCookiePath.IsNull() {
		*hashOnCookiePath = r.HashOnCookiePath.ValueString()
	} else {
		hashOnCookiePath = nil
	}
	hashOnHeader := new(string)
	if !r.HashOnHeader.IsUnknown() && !r.HashOnHeader.IsNull() {
		*hashOnHeader = r.HashOnHeader.ValueString()
	} else {
		hashOnHeader = nil
	}
	hashOnQueryArg := new(string)
	if !r.HashOnQueryArg.IsUnknown() && !r.HashOnQueryArg.IsNull() {
		*hashOnQueryArg = r.HashOnQueryArg.ValueString()
	} else {
		hashOnQueryArg = nil
	}
	hashOnURICapture := new(string)
	if !r.HashOnURICapture.IsUnknown() && !r.HashOnURICapture.IsNull() {
		*hashOnURICapture = r.HashOnURICapture.ValueString()
	} else {
		hashOnURICapture = nil
	}
	var healthchecks *shared.Healthchecks
	if r.Healthchecks != nil {
		var active *shared.Active
		if r.Healthchecks.Active != nil {
			concurrency := new(int64)
			if !r.Healthchecks.Active.Concurrency.IsUnknown() && !r.Healthchecks.Active.Concurrency.IsNull() {
				*concurrency = r.Healthchecks.Active.Concurrency.ValueInt64()
			} else {
				concurrency = nil
			}
			headers := make(map[string]string)
			for headersKey, headersValue := range r.Healthchecks.Active.Headers {
				var headersInst string
				headersInst = headersValue.ValueString()

				headers[headersKey] = headersInst
			}
			var healthy *shared.Healthy
			if r.Healthchecks.Active.Healthy != nil {
				var httpStatuses []int64 = []int64{}
				for _, httpStatusesItem := range r.Healthchecks.Active.Healthy.HTTPStatuses {
					httpStatuses = append(httpStatuses, httpStatusesItem.ValueInt64())
				}
				interval := new(float64)
				if !r.Healthchecks.Active.Healthy.Interval.IsUnknown() && !r.Healthchecks.Active.Healthy.Interval.IsNull() {
					*interval, _ = r.Healthchecks.Active.Healthy.Interval.ValueBigFloat().Float64()
				} else {
					interval = nil
				}
				successes := new(int64)
				if !r.Healthchecks.Active.Healthy.Successes.IsUnknown() && !r.Healthchecks.Active.Healthy.Successes.IsNull() {
					*successes = r.Healthchecks.Active.Healthy.Successes.ValueInt64()
				} else {
					successes = nil
				}
				healthy = &shared.Healthy{
					HTTPStatuses: httpStatuses,
					Interval:     interval,
					Successes:    successes,
				}
			}
			httpPath := new(string)
			if !r.Healthchecks.Active.HTTPPath.IsUnknown() && !r.Healthchecks.Active.HTTPPath.IsNull() {
				*httpPath = r.Healthchecks.Active.HTTPPath.ValueString()
			} else {
				httpPath = nil
			}
			httpsSni := new(string)
			if !r.Healthchecks.Active.HTTPSSni.IsUnknown() && !r.Healthchecks.Active.HTTPSSni.IsNull() {
				*httpsSni = r.Healthchecks.Active.HTTPSSni.ValueString()
			} else {
				httpsSni = nil
			}
			httpsVerifyCertificate := new(bool)
			if !r.Healthchecks.Active.HTTPSVerifyCertificate.IsUnknown() && !r.Healthchecks.Active.HTTPSVerifyCertificate.IsNull() {
				*httpsVerifyCertificate = r.Healthchecks.Active.HTTPSVerifyCertificate.ValueBool()
			} else {
				httpsVerifyCertificate = nil
			}
			timeout := new(float64)
			if !r.Healthchecks.Active.Timeout.IsUnknown() && !r.Healthchecks.Active.Timeout.IsNull() {
				*timeout, _ = r.Healthchecks.Active.Timeout.ValueBigFloat().Float64()
			} else {
				timeout = nil
			}
			typeVar := new(shared.UpstreamType)
			if !r.Healthchecks.Active.Type.IsUnknown() && !r.Healthchecks.Active.Type.IsNull() {
				*typeVar = shared.UpstreamType(r.Healthchecks.Active.Type.ValueString())
			} else {
				typeVar = nil
			}
			var unhealthy *shared.Unhealthy
			if r.Healthchecks.Active.Unhealthy != nil {
				httpFailures := new(int64)
				if !r.Healthchecks.Active.Unhealthy.HTTPFailures.IsUnknown() && !r.Healthchecks.Active.Unhealthy.HTTPFailures.IsNull() {
					*httpFailures = r.Healthchecks.Active.Unhealthy.HTTPFailures.ValueInt64()
				} else {
					httpFailures = nil
				}
				var httpStatuses1 []int64 = []int64{}
				for _, httpStatusesItem1 := range r.Healthchecks.Active.Unhealthy.HTTPStatuses {
					httpStatuses1 = append(httpStatuses1, httpStatusesItem1.ValueInt64())
				}
				interval1 := new(float64)
				if !r.Healthchecks.Active.Unhealthy.Interval.IsUnknown() && !r.Healthchecks.Active.Unhealthy.Interval.IsNull() {
					*interval1, _ = r.Healthchecks.Active.Unhealthy.Interval.ValueBigFloat().Float64()
				} else {
					interval1 = nil
				}
				tcpFailures := new(int64)
				if !r.Healthchecks.Active.Unhealthy.TCPFailures.IsUnknown() && !r.Healthchecks.Active.Unhealthy.TCPFailures.IsNull() {
					*tcpFailures = r.Healthchecks.Active.Unhealthy.TCPFailures.ValueInt64()
				} else {
					tcpFailures = nil
				}
				timeouts := new(int64)
				if !r.Healthchecks.Active.Unhealthy.Timeouts.IsUnknown() && !r.Healthchecks.Active.Unhealthy.Timeouts.IsNull() {
					*timeouts = r.Healthchecks.Active.Unhealthy.Timeouts.ValueInt64()
				} else {
					timeouts = nil
				}
				unhealthy = &shared.Unhealthy{
					HTTPFailures: httpFailures,
					HTTPStatuses: httpStatuses1,
					Interval:     interval1,
					TCPFailures:  tcpFailures,
					Timeouts:     timeouts,
				}
			}
			active = &shared.Active{
				Concurrency:            concurrency,
				Headers:                headers,
				Healthy:                healthy,
				HTTPPath:               httpPath,
				HTTPSSni:               httpsSni,
				HTTPSVerifyCertificate: httpsVerifyCertificate,
				Timeout:                timeout,
				Type:                   typeVar,
				Unhealthy:              unhealthy,
			}
		}
		var passive *shared.Passive
		if r.Healthchecks.Passive != nil {
			var healthy1 *shared.UpstreamHealthy
			if r.Healthchecks.Passive.Healthy != nil {
				var httpStatuses2 []int64 = []int64{}
				for _, httpStatusesItem2 := range r.Healthchecks.Passive.Healthy.HTTPStatuses {
					httpStatuses2 = append(httpStatuses2, httpStatusesItem2.ValueInt64())
				}
				successes1 := new(int64)
				if !r.Healthchecks.Passive.Healthy.Successes.IsUnknown() && !r.Healthchecks.Passive.Healthy.Successes.IsNull() {
					*successes1 = r.Healthchecks.Passive.Healthy.Successes.ValueInt64()
				} else {
					successes1 = nil
				}
				healthy1 = &shared.UpstreamHealthy{
					HTTPStatuses: httpStatuses2,
					Successes:    successes1,
				}
			}
			typeVar1 := new(shared.UpstreamHealthchecksType)
			if !r.Healthchecks.Passive.Type.IsUnknown() && !r.Healthchecks.Passive.Type.IsNull() {
				*typeVar1 = shared.UpstreamHealthchecksType(r.Healthchecks.Passive.Type.ValueString())
			} else {
				typeVar1 = nil
			}
			var unhealthy1 *shared.UpstreamUnhealthy
			if r.Healthchecks.Passive.Unhealthy != nil {
				httpFailures1 := new(int64)
				if !r.Healthchecks.Passive.Unhealthy.HTTPFailures.IsUnknown() && !r.Healthchecks.Passive.Unhealthy.HTTPFailures.IsNull() {
					*httpFailures1 = r.Healthchecks.Passive.Unhealthy.HTTPFailures.ValueInt64()
				} else {
					httpFailures1 = nil
				}
				var httpStatuses3 []int64 = []int64{}
				for _, httpStatusesItem3 := range r.Healthchecks.Passive.Unhealthy.HTTPStatuses {
					httpStatuses3 = append(httpStatuses3, httpStatusesItem3.ValueInt64())
				}
				tcpFailures1 := new(int64)
				if !r.Healthchecks.Passive.Unhealthy.TCPFailures.IsUnknown() && !r.Healthchecks.Passive.Unhealthy.TCPFailures.IsNull() {
					*tcpFailures1 = r.Healthchecks.Passive.Unhealthy.TCPFailures.ValueInt64()
				} else {
					tcpFailures1 = nil
				}
				timeouts1 := new(int64)
				if !r.Healthchecks.Passive.Unhealthy.Timeouts.IsUnknown() && !r.Healthchecks.Passive.Unhealthy.Timeouts.IsNull() {
					*timeouts1 = r.Healthchecks.Passive.Unhealthy.Timeouts.ValueInt64()
				} else {
					timeouts1 = nil
				}
				unhealthy1 = &shared.UpstreamUnhealthy{
					HTTPFailures: httpFailures1,
					HTTPStatuses: httpStatuses3,
					TCPFailures:  tcpFailures1,
					Timeouts:     timeouts1,
				}
			}
			passive = &shared.Passive{
				Healthy:   healthy1,
				Type:      typeVar1,
				Unhealthy: unhealthy1,
			}
		}
		threshold := new(float64)
		if !r.Healthchecks.Threshold.IsUnknown() && !r.Healthchecks.Threshold.IsNull() {
			*threshold, _ = r.Healthchecks.Threshold.ValueBigFloat().Float64()
		} else {
			threshold = nil
		}
		healthchecks = &shared.Healthchecks{
			Active:    active,
			Passive:   passive,
			Threshold: threshold,
		}
	}
	hostHeader := new(string)
	if !r.HostHeader.IsUnknown() && !r.HostHeader.IsNull() {
		*hostHeader = r.HostHeader.ValueString()
	} else {
		hostHeader = nil
	}
	id1 := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id1 = r.ID.ValueString()
	} else {
		id1 = nil
	}
	var name string
	name = r.Name.ValueString()

	slots := new(int64)
	if !r.Slots.IsUnknown() && !r.Slots.IsNull() {
		*slots = r.Slots.ValueInt64()
	} else {
		slots = nil
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	useSrvName := new(bool)
	if !r.UseSrvName.IsUnknown() && !r.UseSrvName.IsNull() {
		*useSrvName = r.UseSrvName.ValueBool()
	} else {
		useSrvName = nil
	}
	out := shared.UpstreamInput{
		Algorithm:              algorithm,
		ClientCertificate:      clientCertificate,
		HashFallback:           hashFallback,
		HashFallbackHeader:     hashFallbackHeader,
		HashFallbackQueryArg:   hashFallbackQueryArg,
		HashFallbackURICapture: hashFallbackURICapture,
		HashOn:                 hashOn,
		HashOnCookie:           hashOnCookie,
		HashOnCookiePath:       hashOnCookiePath,
		HashOnHeader:           hashOnHeader,
		HashOnQueryArg:         hashOnQueryArg,
		HashOnURICapture:       hashOnURICapture,
		Healthchecks:           healthchecks,
		HostHeader:             hostHeader,
		ID:                     id1,
		Name:                   name,
		Slots:                  slots,
		Tags:                   tags,
		UseSrvName:             useSrvName,
	}
	return &out
}

func (r *GatewayUpstreamResourceModel) RefreshFromSharedUpstream(resp *shared.Upstream) {
	if resp != nil {
		if resp.Algorithm != nil {
			r.Algorithm = types.StringValue(string(*resp.Algorithm))
		} else {
			r.Algorithm = types.StringNull()
		}
		if resp.ClientCertificate == nil {
			r.ClientCertificate = nil
		} else {
			r.ClientCertificate = &tfTypes.ACLConsumer{}
			r.ClientCertificate.ID = types.StringPointerValue(resp.ClientCertificate.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		if resp.HashFallback != nil {
			r.HashFallback = types.StringValue(string(*resp.HashFallback))
		} else {
			r.HashFallback = types.StringNull()
		}
		r.HashFallbackHeader = types.StringPointerValue(resp.HashFallbackHeader)
		r.HashFallbackQueryArg = types.StringPointerValue(resp.HashFallbackQueryArg)
		r.HashFallbackURICapture = types.StringPointerValue(resp.HashFallbackURICapture)
		if resp.HashOn != nil {
			r.HashOn = types.StringValue(string(*resp.HashOn))
		} else {
			r.HashOn = types.StringNull()
		}
		r.HashOnCookie = types.StringPointerValue(resp.HashOnCookie)
		r.HashOnCookiePath = types.StringPointerValue(resp.HashOnCookiePath)
		r.HashOnHeader = types.StringPointerValue(resp.HashOnHeader)
		r.HashOnQueryArg = types.StringPointerValue(resp.HashOnQueryArg)
		r.HashOnURICapture = types.StringPointerValue(resp.HashOnURICapture)
		if resp.Healthchecks == nil {
			r.Healthchecks = nil
		} else {
			r.Healthchecks = &tfTypes.Healthchecks{}
			if resp.Healthchecks.Active == nil {
				r.Healthchecks.Active = nil
			} else {
				r.Healthchecks.Active = &tfTypes.Active{}
				r.Healthchecks.Active.Concurrency = types.Int64PointerValue(resp.Healthchecks.Active.Concurrency)
				if len(resp.Healthchecks.Active.Headers) > 0 {
					r.Healthchecks.Active.Headers = make(map[string]types.String)
					for key, value := range resp.Healthchecks.Active.Headers {
						r.Healthchecks.Active.Headers[key] = types.StringValue(value)
					}
				}
				if resp.Healthchecks.Active.Healthy == nil {
					r.Healthchecks.Active.Healthy = nil
				} else {
					r.Healthchecks.Active.Healthy = &tfTypes.Healthy{}
					r.Healthchecks.Active.Healthy.HTTPStatuses = []types.Int64{}
					for _, v := range resp.Healthchecks.Active.Healthy.HTTPStatuses {
						r.Healthchecks.Active.Healthy.HTTPStatuses = append(r.Healthchecks.Active.Healthy.HTTPStatuses, types.Int64Value(v))
					}
					if resp.Healthchecks.Active.Healthy.Interval != nil {
						r.Healthchecks.Active.Healthy.Interval = types.NumberValue(big.NewFloat(float64(*resp.Healthchecks.Active.Healthy.Interval)))
					} else {
						r.Healthchecks.Active.Healthy.Interval = types.NumberNull()
					}
					r.Healthchecks.Active.Healthy.Successes = types.Int64PointerValue(resp.Healthchecks.Active.Healthy.Successes)
				}
				r.Healthchecks.Active.HTTPPath = types.StringPointerValue(resp.Healthchecks.Active.HTTPPath)
				r.Healthchecks.Active.HTTPSSni = types.StringPointerValue(resp.Healthchecks.Active.HTTPSSni)
				r.Healthchecks.Active.HTTPSVerifyCertificate = types.BoolPointerValue(resp.Healthchecks.Active.HTTPSVerifyCertificate)
				if resp.Healthchecks.Active.Timeout != nil {
					r.Healthchecks.Active.Timeout = types.NumberValue(big.NewFloat(float64(*resp.Healthchecks.Active.Timeout)))
				} else {
					r.Healthchecks.Active.Timeout = types.NumberNull()
				}
				if resp.Healthchecks.Active.Type != nil {
					r.Healthchecks.Active.Type = types.StringValue(string(*resp.Healthchecks.Active.Type))
				} else {
					r.Healthchecks.Active.Type = types.StringNull()
				}
				if resp.Healthchecks.Active.Unhealthy == nil {
					r.Healthchecks.Active.Unhealthy = nil
				} else {
					r.Healthchecks.Active.Unhealthy = &tfTypes.Unhealthy{}
					r.Healthchecks.Active.Unhealthy.HTTPFailures = types.Int64PointerValue(resp.Healthchecks.Active.Unhealthy.HTTPFailures)
					r.Healthchecks.Active.Unhealthy.HTTPStatuses = []types.Int64{}
					for _, v := range resp.Healthchecks.Active.Unhealthy.HTTPStatuses {
						r.Healthchecks.Active.Unhealthy.HTTPStatuses = append(r.Healthchecks.Active.Unhealthy.HTTPStatuses, types.Int64Value(v))
					}
					if resp.Healthchecks.Active.Unhealthy.Interval != nil {
						r.Healthchecks.Active.Unhealthy.Interval = types.NumberValue(big.NewFloat(float64(*resp.Healthchecks.Active.Unhealthy.Interval)))
					} else {
						r.Healthchecks.Active.Unhealthy.Interval = types.NumberNull()
					}
					r.Healthchecks.Active.Unhealthy.TCPFailures = types.Int64PointerValue(resp.Healthchecks.Active.Unhealthy.TCPFailures)
					r.Healthchecks.Active.Unhealthy.Timeouts = types.Int64PointerValue(resp.Healthchecks.Active.Unhealthy.Timeouts)
				}
			}
			if resp.Healthchecks.Passive == nil {
				r.Healthchecks.Passive = nil
			} else {
				r.Healthchecks.Passive = &tfTypes.Passive{}
				if resp.Healthchecks.Passive.Healthy == nil {
					r.Healthchecks.Passive.Healthy = nil
				} else {
					r.Healthchecks.Passive.Healthy = &tfTypes.UpstreamHealthy{}
					r.Healthchecks.Passive.Healthy.HTTPStatuses = []types.Int64{}
					for _, v := range resp.Healthchecks.Passive.Healthy.HTTPStatuses {
						r.Healthchecks.Passive.Healthy.HTTPStatuses = append(r.Healthchecks.Passive.Healthy.HTTPStatuses, types.Int64Value(v))
					}
					r.Healthchecks.Passive.Healthy.Successes = types.Int64PointerValue(resp.Healthchecks.Passive.Healthy.Successes)
				}
				if resp.Healthchecks.Passive.Type != nil {
					r.Healthchecks.Passive.Type = types.StringValue(string(*resp.Healthchecks.Passive.Type))
				} else {
					r.Healthchecks.Passive.Type = types.StringNull()
				}
				if resp.Healthchecks.Passive.Unhealthy == nil {
					r.Healthchecks.Passive.Unhealthy = nil
				} else {
					r.Healthchecks.Passive.Unhealthy = &tfTypes.UpstreamUnhealthy{}
					r.Healthchecks.Passive.Unhealthy.HTTPFailures = types.Int64PointerValue(resp.Healthchecks.Passive.Unhealthy.HTTPFailures)
					r.Healthchecks.Passive.Unhealthy.HTTPStatuses = []types.Int64{}
					for _, v := range resp.Healthchecks.Passive.Unhealthy.HTTPStatuses {
						r.Healthchecks.Passive.Unhealthy.HTTPStatuses = append(r.Healthchecks.Passive.Unhealthy.HTTPStatuses, types.Int64Value(v))
					}
					r.Healthchecks.Passive.Unhealthy.TCPFailures = types.Int64PointerValue(resp.Healthchecks.Passive.Unhealthy.TCPFailures)
					r.Healthchecks.Passive.Unhealthy.Timeouts = types.Int64PointerValue(resp.Healthchecks.Passive.Unhealthy.Timeouts)
				}
			}
			if resp.Healthchecks.Threshold != nil {
				r.Healthchecks.Threshold = types.NumberValue(big.NewFloat(float64(*resp.Healthchecks.Threshold)))
			} else {
				r.Healthchecks.Threshold = types.NumberNull()
			}
		}
		r.HostHeader = types.StringPointerValue(resp.HostHeader)
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringValue(resp.Name)
		r.Slots = types.Int64PointerValue(resp.Slots)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
		r.UseSrvName = types.BoolPointerValue(resp.UseSrvName)
	}
}
