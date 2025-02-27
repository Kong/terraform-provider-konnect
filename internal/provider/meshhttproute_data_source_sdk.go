// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshHTTPRouteDataSourceModel) RefreshFromSharedMeshHTTPRouteItem(resp *shared.MeshHTTPRouteItem) {
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
			r.Spec.TargetRef.ProxyTypes = make([]types.String, 0, len(resp.Spec.TargetRef.ProxyTypes))
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
		r.Spec.To = []tfTypes.MeshHTTPRouteItemTo{}
		if len(r.Spec.To) > len(resp.Spec.To) {
			r.Spec.To = r.Spec.To[:len(resp.Spec.To)]
		}
		for toCount, toItem := range resp.Spec.To {
			var to1 tfTypes.MeshHTTPRouteItemTo
			to1.Hostnames = make([]types.String, 0, len(toItem.Hostnames))
			for _, v := range toItem.Hostnames {
				to1.Hostnames = append(to1.Hostnames, types.StringValue(v))
			}
			to1.Rules = []tfTypes.MeshHTTPRouteItemRules{}
			for rulesCount, rulesItem := range toItem.Rules {
				var rules1 tfTypes.MeshHTTPRouteItemRules
				rules1.Default.BackendRefs = []tfTypes.BackendRefs{}
				for backendRefsCount, backendRefsItem := range rulesItem.Default.BackendRefs {
					var backendRefs1 tfTypes.BackendRefs
					if backendRefsItem.Kind != nil {
						backendRefs1.Kind = types.StringValue(string(*backendRefsItem.Kind))
					} else {
						backendRefs1.Kind = types.StringNull()
					}
					if len(backendRefsItem.Labels) > 0 {
						backendRefs1.Labels = make(map[string]types.String)
						for key3, value3 := range backendRefsItem.Labels {
							backendRefs1.Labels[key3] = types.StringValue(value3)
						}
					}
					backendRefs1.Mesh = types.StringPointerValue(backendRefsItem.Mesh)
					backendRefs1.Name = types.StringPointerValue(backendRefsItem.Name)
					backendRefs1.Namespace = types.StringPointerValue(backendRefsItem.Namespace)
					if backendRefsItem.Port != nil {
						backendRefs1.Port = types.Int64Value(int64(*backendRefsItem.Port))
					} else {
						backendRefs1.Port = types.Int64Null()
					}
					backendRefs1.ProxyTypes = make([]types.String, 0, len(backendRefsItem.ProxyTypes))
					for _, v := range backendRefsItem.ProxyTypes {
						backendRefs1.ProxyTypes = append(backendRefs1.ProxyTypes, types.StringValue(string(v)))
					}
					backendRefs1.SectionName = types.StringPointerValue(backendRefsItem.SectionName)
					if len(backendRefsItem.Tags) > 0 {
						backendRefs1.Tags = make(map[string]types.String)
						for key4, value4 := range backendRefsItem.Tags {
							backendRefs1.Tags[key4] = types.StringValue(value4)
						}
					}
					backendRefs1.Weight = types.Int64PointerValue(backendRefsItem.Weight)
					if backendRefsCount+1 > len(rules1.Default.BackendRefs) {
						rules1.Default.BackendRefs = append(rules1.Default.BackendRefs, backendRefs1)
					} else {
						rules1.Default.BackendRefs[backendRefsCount].Kind = backendRefs1.Kind
						rules1.Default.BackendRefs[backendRefsCount].Labels = backendRefs1.Labels
						rules1.Default.BackendRefs[backendRefsCount].Mesh = backendRefs1.Mesh
						rules1.Default.BackendRefs[backendRefsCount].Name = backendRefs1.Name
						rules1.Default.BackendRefs[backendRefsCount].Namespace = backendRefs1.Namespace
						rules1.Default.BackendRefs[backendRefsCount].Port = backendRefs1.Port
						rules1.Default.BackendRefs[backendRefsCount].ProxyTypes = backendRefs1.ProxyTypes
						rules1.Default.BackendRefs[backendRefsCount].SectionName = backendRefs1.SectionName
						rules1.Default.BackendRefs[backendRefsCount].Tags = backendRefs1.Tags
						rules1.Default.BackendRefs[backendRefsCount].Weight = backendRefs1.Weight
					}
				}
				rules1.Default.Filters = []tfTypes.Filters{}
				for filtersCount, filtersItem := range rulesItem.Default.Filters {
					var filters1 tfTypes.Filters
					if filtersItem.RequestHeaderModifier == nil {
						filters1.RequestHeaderModifier = nil
					} else {
						filters1.RequestHeaderModifier = &tfTypes.RequestHeaderModifier{}
						filters1.RequestHeaderModifier.Add = []tfTypes.ConfigurationDataPlaneGroupEnvironmentField{}
						for addCount, addItem := range filtersItem.RequestHeaderModifier.Add {
							var add1 tfTypes.ConfigurationDataPlaneGroupEnvironmentField
							add1.Name = types.StringValue(addItem.Name)
							add1.Value = types.StringValue(addItem.Value)
							if addCount+1 > len(filters1.RequestHeaderModifier.Add) {
								filters1.RequestHeaderModifier.Add = append(filters1.RequestHeaderModifier.Add, add1)
							} else {
								filters1.RequestHeaderModifier.Add[addCount].Name = add1.Name
								filters1.RequestHeaderModifier.Add[addCount].Value = add1.Value
							}
						}
						filters1.RequestHeaderModifier.Remove = make([]types.String, 0, len(filtersItem.RequestHeaderModifier.Remove))
						for _, v := range filtersItem.RequestHeaderModifier.Remove {
							filters1.RequestHeaderModifier.Remove = append(filters1.RequestHeaderModifier.Remove, types.StringValue(v))
						}
						filters1.RequestHeaderModifier.Set = []tfTypes.ConfigurationDataPlaneGroupEnvironmentField{}
						for setCount, setItem := range filtersItem.RequestHeaderModifier.Set {
							var set1 tfTypes.ConfigurationDataPlaneGroupEnvironmentField
							set1.Name = types.StringValue(setItem.Name)
							set1.Value = types.StringValue(setItem.Value)
							if setCount+1 > len(filters1.RequestHeaderModifier.Set) {
								filters1.RequestHeaderModifier.Set = append(filters1.RequestHeaderModifier.Set, set1)
							} else {
								filters1.RequestHeaderModifier.Set[setCount].Name = set1.Name
								filters1.RequestHeaderModifier.Set[setCount].Value = set1.Value
							}
						}
					}
					if filtersItem.RequestMirror == nil {
						filters1.RequestMirror = nil
					} else {
						filters1.RequestMirror = &tfTypes.RequestMirror{}
						if filtersItem.RequestMirror.BackendRef.Kind != nil {
							filters1.RequestMirror.BackendRef.Kind = types.StringValue(string(*filtersItem.RequestMirror.BackendRef.Kind))
						} else {
							filters1.RequestMirror.BackendRef.Kind = types.StringNull()
						}
						if len(filtersItem.RequestMirror.BackendRef.Labels) > 0 {
							filters1.RequestMirror.BackendRef.Labels = make(map[string]types.String)
							for key5, value7 := range filtersItem.RequestMirror.BackendRef.Labels {
								filters1.RequestMirror.BackendRef.Labels[key5] = types.StringValue(value7)
							}
						}
						filters1.RequestMirror.BackendRef.Mesh = types.StringPointerValue(filtersItem.RequestMirror.BackendRef.Mesh)
						filters1.RequestMirror.BackendRef.Name = types.StringPointerValue(filtersItem.RequestMirror.BackendRef.Name)
						filters1.RequestMirror.BackendRef.Namespace = types.StringPointerValue(filtersItem.RequestMirror.BackendRef.Namespace)
						if filtersItem.RequestMirror.BackendRef.Port != nil {
							filters1.RequestMirror.BackendRef.Port = types.Int64Value(int64(*filtersItem.RequestMirror.BackendRef.Port))
						} else {
							filters1.RequestMirror.BackendRef.Port = types.Int64Null()
						}
						filters1.RequestMirror.BackendRef.ProxyTypes = make([]types.String, 0, len(filtersItem.RequestMirror.BackendRef.ProxyTypes))
						for _, v := range filtersItem.RequestMirror.BackendRef.ProxyTypes {
							filters1.RequestMirror.BackendRef.ProxyTypes = append(filters1.RequestMirror.BackendRef.ProxyTypes, types.StringValue(string(v)))
						}
						filters1.RequestMirror.BackendRef.SectionName = types.StringPointerValue(filtersItem.RequestMirror.BackendRef.SectionName)
						if len(filtersItem.RequestMirror.BackendRef.Tags) > 0 {
							filters1.RequestMirror.BackendRef.Tags = make(map[string]types.String)
							for key6, value8 := range filtersItem.RequestMirror.BackendRef.Tags {
								filters1.RequestMirror.BackendRef.Tags[key6] = types.StringValue(value8)
							}
						}
						filters1.RequestMirror.BackendRef.Weight = types.Int64PointerValue(filtersItem.RequestMirror.BackendRef.Weight)
						if filtersItem.RequestMirror.Percentage == nil {
							filters1.RequestMirror.Percentage = nil
						} else {
							filters1.RequestMirror.Percentage = &tfTypes.Mode{}
							if filtersItem.RequestMirror.Percentage.Integer != nil {
								filters1.RequestMirror.Percentage.Integer = types.Int64PointerValue(filtersItem.RequestMirror.Percentage.Integer)
							}
							if filtersItem.RequestMirror.Percentage.Str != nil {
								filters1.RequestMirror.Percentage.Str = types.StringPointerValue(filtersItem.RequestMirror.Percentage.Str)
							}
						}
					}
					if filtersItem.RequestRedirect == nil {
						filters1.RequestRedirect = nil
					} else {
						filters1.RequestRedirect = &tfTypes.RequestRedirect{}
						filters1.RequestRedirect.Hostname = types.StringPointerValue(filtersItem.RequestRedirect.Hostname)
						if filtersItem.RequestRedirect.Path == nil {
							filters1.RequestRedirect.Path = nil
						} else {
							filters1.RequestRedirect.Path = &tfTypes.MeshHTTPRouteItemSpecPath{}
							filters1.RequestRedirect.Path.ReplaceFullPath = types.StringPointerValue(filtersItem.RequestRedirect.Path.ReplaceFullPath)
							filters1.RequestRedirect.Path.ReplacePrefixMatch = types.StringPointerValue(filtersItem.RequestRedirect.Path.ReplacePrefixMatch)
							filters1.RequestRedirect.Path.Type = types.StringValue(string(filtersItem.RequestRedirect.Path.Type))
						}
						if filtersItem.RequestRedirect.Port != nil {
							filters1.RequestRedirect.Port = types.Int64Value(int64(*filtersItem.RequestRedirect.Port))
						} else {
							filters1.RequestRedirect.Port = types.Int64Null()
						}
						if filtersItem.RequestRedirect.Scheme != nil {
							filters1.RequestRedirect.Scheme = types.StringValue(string(*filtersItem.RequestRedirect.Scheme))
						} else {
							filters1.RequestRedirect.Scheme = types.StringNull()
						}
						if filtersItem.RequestRedirect.StatusCode != nil {
							filters1.RequestRedirect.StatusCode = types.Int64Value(int64(*filtersItem.RequestRedirect.StatusCode))
						} else {
							filters1.RequestRedirect.StatusCode = types.Int64Null()
						}
					}
					if filtersItem.ResponseHeaderModifier == nil {
						filters1.ResponseHeaderModifier = nil
					} else {
						filters1.ResponseHeaderModifier = &tfTypes.RequestHeaderModifier{}
						filters1.ResponseHeaderModifier.Add = []tfTypes.ConfigurationDataPlaneGroupEnvironmentField{}
						for addCount1, addItem1 := range filtersItem.ResponseHeaderModifier.Add {
							var add3 tfTypes.ConfigurationDataPlaneGroupEnvironmentField
							add3.Name = types.StringValue(addItem1.Name)
							add3.Value = types.StringValue(addItem1.Value)
							if addCount1+1 > len(filters1.ResponseHeaderModifier.Add) {
								filters1.ResponseHeaderModifier.Add = append(filters1.ResponseHeaderModifier.Add, add3)
							} else {
								filters1.ResponseHeaderModifier.Add[addCount1].Name = add3.Name
								filters1.ResponseHeaderModifier.Add[addCount1].Value = add3.Value
							}
						}
						filters1.ResponseHeaderModifier.Remove = make([]types.String, 0, len(filtersItem.ResponseHeaderModifier.Remove))
						for _, v := range filtersItem.ResponseHeaderModifier.Remove {
							filters1.ResponseHeaderModifier.Remove = append(filters1.ResponseHeaderModifier.Remove, types.StringValue(v))
						}
						filters1.ResponseHeaderModifier.Set = []tfTypes.ConfigurationDataPlaneGroupEnvironmentField{}
						for setCount1, setItem1 := range filtersItem.ResponseHeaderModifier.Set {
							var set3 tfTypes.ConfigurationDataPlaneGroupEnvironmentField
							set3.Name = types.StringValue(setItem1.Name)
							set3.Value = types.StringValue(setItem1.Value)
							if setCount1+1 > len(filters1.ResponseHeaderModifier.Set) {
								filters1.ResponseHeaderModifier.Set = append(filters1.ResponseHeaderModifier.Set, set3)
							} else {
								filters1.ResponseHeaderModifier.Set[setCount1].Name = set3.Name
								filters1.ResponseHeaderModifier.Set[setCount1].Value = set3.Value
							}
						}
					}
					filters1.Type = types.StringValue(string(filtersItem.Type))
					if filtersItem.URLRewrite == nil {
						filters1.URLRewrite = nil
					} else {
						filters1.URLRewrite = &tfTypes.URLRewrite{}
						filters1.URLRewrite.Hostname = types.StringPointerValue(filtersItem.URLRewrite.Hostname)
						filters1.URLRewrite.HostToBackendHostname = types.BoolPointerValue(filtersItem.URLRewrite.HostToBackendHostname)
						if filtersItem.URLRewrite.Path == nil {
							filters1.URLRewrite.Path = nil
						} else {
							filters1.URLRewrite.Path = &tfTypes.MeshHTTPRouteItemSpecPath{}
							filters1.URLRewrite.Path.ReplaceFullPath = types.StringPointerValue(filtersItem.URLRewrite.Path.ReplaceFullPath)
							filters1.URLRewrite.Path.ReplacePrefixMatch = types.StringPointerValue(filtersItem.URLRewrite.Path.ReplacePrefixMatch)
							filters1.URLRewrite.Path.Type = types.StringValue(string(filtersItem.URLRewrite.Path.Type))
						}
					}
					if filtersCount+1 > len(rules1.Default.Filters) {
						rules1.Default.Filters = append(rules1.Default.Filters, filters1)
					} else {
						rules1.Default.Filters[filtersCount].RequestHeaderModifier = filters1.RequestHeaderModifier
						rules1.Default.Filters[filtersCount].RequestMirror = filters1.RequestMirror
						rules1.Default.Filters[filtersCount].RequestRedirect = filters1.RequestRedirect
						rules1.Default.Filters[filtersCount].ResponseHeaderModifier = filters1.ResponseHeaderModifier
						rules1.Default.Filters[filtersCount].Type = filters1.Type
						rules1.Default.Filters[filtersCount].URLRewrite = filters1.URLRewrite
					}
				}
				rules1.Matches = []tfTypes.Matches{}
				for matchesCount, matchesItem := range rulesItem.Matches {
					var matches1 tfTypes.Matches
					matches1.Headers = []tfTypes.Headers{}
					for headersCount, headersItem := range matchesItem.Headers {
						var headers1 tfTypes.Headers
						headers1.Name = types.StringValue(headersItem.Name)
						if headersItem.Type != nil {
							headers1.Type = types.StringValue(string(*headersItem.Type))
						} else {
							headers1.Type = types.StringNull()
						}
						headers1.Value = types.StringPointerValue(headersItem.Value)
						if headersCount+1 > len(matches1.Headers) {
							matches1.Headers = append(matches1.Headers, headers1)
						} else {
							matches1.Headers[headersCount].Name = headers1.Name
							matches1.Headers[headersCount].Type = headers1.Type
							matches1.Headers[headersCount].Value = headers1.Value
						}
					}
					if matchesItem.Method != nil {
						matches1.Method = types.StringValue(string(*matchesItem.Method))
					} else {
						matches1.Method = types.StringNull()
					}
					if matchesItem.Path == nil {
						matches1.Path = nil
					} else {
						matches1.Path = &tfTypes.Path{}
						matches1.Path.Type = types.StringValue(string(matchesItem.Path.Type))
						matches1.Path.Value = types.StringValue(matchesItem.Path.Value)
					}
					matches1.QueryParams = []tfTypes.QueryParams{}
					for queryParamsCount, queryParamsItem := range matchesItem.QueryParams {
						var queryParams1 tfTypes.QueryParams
						queryParams1.Name = types.StringValue(queryParamsItem.Name)
						queryParams1.Type = types.StringValue(string(queryParamsItem.Type))
						queryParams1.Value = types.StringValue(queryParamsItem.Value)
						if queryParamsCount+1 > len(matches1.QueryParams) {
							matches1.QueryParams = append(matches1.QueryParams, queryParams1)
						} else {
							matches1.QueryParams[queryParamsCount].Name = queryParams1.Name
							matches1.QueryParams[queryParamsCount].Type = queryParams1.Type
							matches1.QueryParams[queryParamsCount].Value = queryParams1.Value
						}
					}
					if matchesCount+1 > len(rules1.Matches) {
						rules1.Matches = append(rules1.Matches, matches1)
					} else {
						rules1.Matches[matchesCount].Headers = matches1.Headers
						rules1.Matches[matchesCount].Method = matches1.Method
						rules1.Matches[matchesCount].Path = matches1.Path
						rules1.Matches[matchesCount].QueryParams = matches1.QueryParams
					}
				}
				if rulesCount+1 > len(to1.Rules) {
					to1.Rules = append(to1.Rules, rules1)
				} else {
					to1.Rules[rulesCount].Default = rules1.Default
					to1.Rules[rulesCount].Matches = rules1.Matches
				}
			}
			if toItem.TargetRef.Kind != nil {
				to1.TargetRef.Kind = types.StringValue(string(*toItem.TargetRef.Kind))
			} else {
				to1.TargetRef.Kind = types.StringNull()
			}
			if len(toItem.TargetRef.Labels) > 0 {
				to1.TargetRef.Labels = make(map[string]types.String)
				for key7, value14 := range toItem.TargetRef.Labels {
					to1.TargetRef.Labels[key7] = types.StringValue(value14)
				}
			}
			to1.TargetRef.Mesh = types.StringPointerValue(toItem.TargetRef.Mesh)
			to1.TargetRef.Name = types.StringPointerValue(toItem.TargetRef.Name)
			to1.TargetRef.Namespace = types.StringPointerValue(toItem.TargetRef.Namespace)
			to1.TargetRef.ProxyTypes = make([]types.String, 0, len(toItem.TargetRef.ProxyTypes))
			for _, v := range toItem.TargetRef.ProxyTypes {
				to1.TargetRef.ProxyTypes = append(to1.TargetRef.ProxyTypes, types.StringValue(string(v)))
			}
			to1.TargetRef.SectionName = types.StringPointerValue(toItem.TargetRef.SectionName)
			if len(toItem.TargetRef.Tags) > 0 {
				to1.TargetRef.Tags = make(map[string]types.String)
				for key8, value15 := range toItem.TargetRef.Tags {
					to1.TargetRef.Tags[key8] = types.StringValue(value15)
				}
			}
			if toCount+1 > len(r.Spec.To) {
				r.Spec.To = append(r.Spec.To, to1)
			} else {
				r.Spec.To[toCount].Hostnames = to1.Hostnames
				r.Spec.To[toCount].Rules = to1.Rules
				r.Spec.To[toCount].TargetRef = to1.TargetRef
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}
}
