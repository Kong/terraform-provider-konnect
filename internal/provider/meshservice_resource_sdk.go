// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshServiceResourceModel) ToSharedMeshServiceItemInput() *shared.MeshServiceItemInput {
	typeVar := shared.MeshServiceItemType(r.Type.ValueString())
	mesh := new(string)
	if !r.Mesh.IsUnknown() && !r.Mesh.IsNull() {
		*mesh = r.Mesh.ValueString()
	} else {
		mesh = nil
	}
	var name string
	name = r.Name.ValueString()

	labels := make(map[string]string)
	for labelsKey, labelsValue := range r.Labels {
		var labelsInst string
		labelsInst = labelsValue.ValueString()

		labels[labelsKey] = labelsInst
	}
	var identities []shared.Identities = []shared.Identities{}
	for _, identitiesItem := range r.Spec.Identities {
		type1 := shared.MeshServiceItemSpecType(identitiesItem.Type.ValueString())
		var value string
		value = identitiesItem.Value.ValueString()

		identities = append(identities, shared.Identities{
			Type:  type1,
			Value: value,
		})
	}
	var ports []shared.MeshServiceItemPorts = []shared.MeshServiceItemPorts{}
	for _, portsItem := range r.Spec.Ports {
		appProtocol := new(string)
		if !portsItem.AppProtocol.IsUnknown() && !portsItem.AppProtocol.IsNull() {
			*appProtocol = portsItem.AppProtocol.ValueString()
		} else {
			appProtocol = nil
		}
		name1 := new(string)
		if !portsItem.Name.IsUnknown() && !portsItem.Name.IsNull() {
			*name1 = portsItem.Name.ValueString()
		} else {
			name1 = nil
		}
		var port int
		port = int(portsItem.Port.ValueInt64())

		var targetPort *shared.TargetPort
		if portsItem.TargetPort != nil {
			integer := new(int64)
			if !portsItem.TargetPort.Integer.IsUnknown() && !portsItem.TargetPort.Integer.IsNull() {
				*integer = portsItem.TargetPort.Integer.ValueInt64()
			} else {
				integer = nil
			}
			if integer != nil {
				targetPort = &shared.TargetPort{
					Integer: integer,
				}
			}
			str := new(string)
			if !portsItem.TargetPort.Str.IsUnknown() && !portsItem.TargetPort.Str.IsNull() {
				*str = portsItem.TargetPort.Str.ValueString()
			} else {
				str = nil
			}
			if str != nil {
				targetPort = &shared.TargetPort{
					Str: str,
				}
			}
		}
		ports = append(ports, shared.MeshServiceItemPorts{
			AppProtocol: appProtocol,
			Name:        name1,
			Port:        port,
			TargetPort:  targetPort,
		})
	}
	var selector *shared.MeshServiceItemSelector
	if r.Spec.Selector != nil {
		var dataplaneRef *shared.DataplaneRef
		if r.Spec.Selector.DataplaneRef != nil {
			name2 := new(string)
			if !r.Spec.Selector.DataplaneRef.Name.IsUnknown() && !r.Spec.Selector.DataplaneRef.Name.IsNull() {
				*name2 = r.Spec.Selector.DataplaneRef.Name.ValueString()
			} else {
				name2 = nil
			}
			dataplaneRef = &shared.DataplaneRef{
				Name: name2,
			}
		}
		dataplaneTags := make(map[string]string)
		for dataplaneTagsKey, dataplaneTagsValue := range r.Spec.Selector.DataplaneTags {
			var dataplaneTagsInst string
			dataplaneTagsInst = dataplaneTagsValue.ValueString()

			dataplaneTags[dataplaneTagsKey] = dataplaneTagsInst
		}
		selector = &shared.MeshServiceItemSelector{
			DataplaneRef:  dataplaneRef,
			DataplaneTags: dataplaneTags,
		}
	}
	state := new(shared.MeshServiceItemState)
	if !r.Spec.State.IsUnknown() && !r.Spec.State.IsNull() {
		*state = shared.MeshServiceItemState(r.Spec.State.ValueString())
	} else {
		state = nil
	}
	spec := shared.MeshServiceItemSpec{
		Identities: identities,
		Ports:      ports,
		Selector:   selector,
		State:      state,
	}
	out := shared.MeshServiceItemInput{
		Type:   typeVar,
		Mesh:   mesh,
		Name:   name,
		Labels: labels,
		Spec:   spec,
	}
	return &out
}

func (r *MeshServiceResourceModel) RefreshFromSharedMeshServiceCreateOrUpdateSuccessResponse(resp *shared.MeshServiceCreateOrUpdateSuccessResponse) {
	if resp != nil {
		r.Warnings = make([]types.String, 0, len(resp.Warnings))
		for _, v := range resp.Warnings {
			r.Warnings = append(r.Warnings, types.StringValue(v))
		}
	}
}

func (r *MeshServiceResourceModel) RefreshFromSharedMeshServiceItem(resp *shared.MeshServiceItem) {
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
		r.Spec.Identities = []tfTypes.Path{}
		if len(r.Spec.Identities) > len(resp.Spec.Identities) {
			r.Spec.Identities = r.Spec.Identities[:len(resp.Spec.Identities)]
		}
		for identitiesCount, identitiesItem := range resp.Spec.Identities {
			var identities1 tfTypes.Path
			identities1.Type = types.StringValue(string(identitiesItem.Type))
			identities1.Value = types.StringValue(identitiesItem.Value)
			if identitiesCount+1 > len(r.Spec.Identities) {
				r.Spec.Identities = append(r.Spec.Identities, identities1)
			} else {
				r.Spec.Identities[identitiesCount].Type = identities1.Type
				r.Spec.Identities[identitiesCount].Value = identities1.Value
			}
		}
		r.Spec.Ports = []tfTypes.MeshServiceItemPorts{}
		if len(r.Spec.Ports) > len(resp.Spec.Ports) {
			r.Spec.Ports = r.Spec.Ports[:len(resp.Spec.Ports)]
		}
		for portsCount, portsItem := range resp.Spec.Ports {
			var ports1 tfTypes.MeshServiceItemPorts
			ports1.AppProtocol = types.StringPointerValue(portsItem.AppProtocol)
			ports1.Name = types.StringPointerValue(portsItem.Name)
			ports1.Port = types.Int64Value(int64(portsItem.Port))
			if portsItem.TargetPort == nil {
				ports1.TargetPort = nil
			} else {
				ports1.TargetPort = &tfTypes.Mode{}
				if portsItem.TargetPort.Integer != nil {
					ports1.TargetPort.Integer = types.Int64PointerValue(portsItem.TargetPort.Integer)
				}
				if portsItem.TargetPort.Str != nil {
					ports1.TargetPort.Str = types.StringPointerValue(portsItem.TargetPort.Str)
				}
			}
			if portsCount+1 > len(r.Spec.Ports) {
				r.Spec.Ports = append(r.Spec.Ports, ports1)
			} else {
				r.Spec.Ports[portsCount].AppProtocol = ports1.AppProtocol
				r.Spec.Ports[portsCount].Name = ports1.Name
				r.Spec.Ports[portsCount].Port = ports1.Port
				r.Spec.Ports[portsCount].TargetPort = ports1.TargetPort
			}
		}
		if resp.Spec.Selector == nil {
			r.Spec.Selector = nil
		} else {
			r.Spec.Selector = &tfTypes.MeshServiceItemSelector{}
			if resp.Spec.Selector.DataplaneRef == nil {
				r.Spec.Selector.DataplaneRef = nil
			} else {
				r.Spec.Selector.DataplaneRef = &tfTypes.DataplaneRef{}
				r.Spec.Selector.DataplaneRef.Name = types.StringPointerValue(resp.Spec.Selector.DataplaneRef.Name)
			}
			if len(resp.Spec.Selector.DataplaneTags) > 0 {
				r.Spec.Selector.DataplaneTags = make(map[string]types.String)
				for key1, value2 := range resp.Spec.Selector.DataplaneTags {
					r.Spec.Selector.DataplaneTags[key1] = types.StringValue(value2)
				}
			}
		}
		if resp.Spec.State != nil {
			r.Spec.State = types.StringValue(string(*resp.Spec.State))
		} else {
			r.Spec.State = types.StringNull()
		}
		if resp.Status == nil {
			r.Status = nil
		} else {
			r.Status = &tfTypes.MeshServiceItemStatus{}
			r.Status.Addresses = []tfTypes.Addresses{}
			if len(r.Status.Addresses) > len(resp.Status.Addresses) {
				r.Status.Addresses = r.Status.Addresses[:len(resp.Status.Addresses)]
			}
			for addressesCount, addressesItem := range resp.Status.Addresses {
				var addresses1 tfTypes.Addresses
				addresses1.Hostname = types.StringPointerValue(addressesItem.Hostname)
				if addressesItem.HostnameGeneratorRef == nil {
					addresses1.HostnameGeneratorRef = nil
				} else {
					addresses1.HostnameGeneratorRef = &tfTypes.HostnameGeneratorRef{}
					addresses1.HostnameGeneratorRef.CoreName = types.StringValue(addressesItem.HostnameGeneratorRef.CoreName)
				}
				addresses1.Origin = types.StringPointerValue(addressesItem.Origin)
				if addressesCount+1 > len(r.Status.Addresses) {
					r.Status.Addresses = append(r.Status.Addresses, addresses1)
				} else {
					r.Status.Addresses[addressesCount].Hostname = addresses1.Hostname
					r.Status.Addresses[addressesCount].HostnameGeneratorRef = addresses1.HostnameGeneratorRef
					r.Status.Addresses[addressesCount].Origin = addresses1.Origin
				}
			}
			if resp.Status.DataplaneProxies == nil {
				r.Status.DataplaneProxies = nil
			} else {
				r.Status.DataplaneProxies = &tfTypes.DataplaneProxies{}
				r.Status.DataplaneProxies.Connected = types.Int64PointerValue(resp.Status.DataplaneProxies.Connected)
				r.Status.DataplaneProxies.Healthy = types.Int64PointerValue(resp.Status.DataplaneProxies.Healthy)
				r.Status.DataplaneProxies.Total = types.Int64PointerValue(resp.Status.DataplaneProxies.Total)
			}
			r.Status.HostnameGenerators = []tfTypes.HostnameGenerators{}
			if len(r.Status.HostnameGenerators) > len(resp.Status.HostnameGenerators) {
				r.Status.HostnameGenerators = r.Status.HostnameGenerators[:len(resp.Status.HostnameGenerators)]
			}
			for hostnameGeneratorsCount, hostnameGeneratorsItem := range resp.Status.HostnameGenerators {
				var hostnameGenerators1 tfTypes.HostnameGenerators
				hostnameGenerators1.Conditions = []tfTypes.Conditions{}
				for conditionsCount, conditionsItem := range hostnameGeneratorsItem.Conditions {
					var conditions1 tfTypes.Conditions
					conditions1.Message = types.StringValue(conditionsItem.Message)
					conditions1.Reason = types.StringValue(conditionsItem.Reason)
					conditions1.Status = types.StringValue(string(conditionsItem.Status))
					conditions1.Type = types.StringValue(conditionsItem.Type)
					if conditionsCount+1 > len(hostnameGenerators1.Conditions) {
						hostnameGenerators1.Conditions = append(hostnameGenerators1.Conditions, conditions1)
					} else {
						hostnameGenerators1.Conditions[conditionsCount].Message = conditions1.Message
						hostnameGenerators1.Conditions[conditionsCount].Reason = conditions1.Reason
						hostnameGenerators1.Conditions[conditionsCount].Status = conditions1.Status
						hostnameGenerators1.Conditions[conditionsCount].Type = conditions1.Type
					}
				}
				hostnameGenerators1.HostnameGeneratorRef.CoreName = types.StringValue(hostnameGeneratorsItem.HostnameGeneratorRef.CoreName)
				if hostnameGeneratorsCount+1 > len(r.Status.HostnameGenerators) {
					r.Status.HostnameGenerators = append(r.Status.HostnameGenerators, hostnameGenerators1)
				} else {
					r.Status.HostnameGenerators[hostnameGeneratorsCount].Conditions = hostnameGenerators1.Conditions
					r.Status.HostnameGenerators[hostnameGeneratorsCount].HostnameGeneratorRef = hostnameGenerators1.HostnameGeneratorRef
				}
			}
			if resp.Status.TLS == nil {
				r.Status.TLS = nil
			} else {
				r.Status.TLS = &tfTypes.MeshServiceItemTLS{}
				if resp.Status.TLS.Status != nil {
					r.Status.TLS.Status = types.StringValue(string(*resp.Status.TLS.Status))
				} else {
					r.Status.TLS.Status = types.StringNull()
				}
			}
			r.Status.Vips = []tfTypes.Vip{}
			if len(r.Status.Vips) > len(resp.Status.Vips) {
				r.Status.Vips = r.Status.Vips[:len(resp.Status.Vips)]
			}
			for vipsCount, vipsItem := range resp.Status.Vips {
				var vips1 tfTypes.Vip
				vips1.IP = types.StringPointerValue(vipsItem.IP)
				if vipsCount+1 > len(r.Status.Vips) {
					r.Status.Vips = append(r.Status.Vips, vips1)
				} else {
					r.Status.Vips[vipsCount].IP = vips1.IP
				}
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}
}
