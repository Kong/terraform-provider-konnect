// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type APIProductPortal struct {
	PortalID   string `json:"portal_id"`
	PortalName string `json:"portal_name"`
}

func (o *APIProductPortal) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

func (o *APIProductPortal) GetPortalName() string {
	if o == nil {
		return ""
	}
	return o.PortalName
}
