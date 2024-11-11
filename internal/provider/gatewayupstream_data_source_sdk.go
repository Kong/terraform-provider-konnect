// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayUpstreamDataSourceModel) RefreshFromSharedUpstream(resp *shared.Upstream) {
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
