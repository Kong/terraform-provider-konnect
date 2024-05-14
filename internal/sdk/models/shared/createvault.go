// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateVault - Vault entities are used to configure different Vault connectors. Examples of Vaults are Environment Variables, Hashicorp Vault and AWS Secrets Manager. Configuring a Vault allows referencing the secrets with other entities. For example a certificate entity can store a reference to a certificate and key, stored in a vault, instead of storing the certificate and key within the entity. This allows a proper separation of secrets and configuration and prevents secret sprawl.
type CreateVault struct {
	// The configuration properties for the Vault which can be found on the vaults' documentation page.
	Config any `json:"config,omitempty"`
	// The description of the Vault entity.
	Description *string `json:"description,omitempty"`
	// The name of the Vault that's going to be added. Currently, the Vault implementation must be installed in every Kong instance.
	Name *string `json:"name,omitempty"`
	// The unique prefix (or identifier) for this Vault configuration. The prefix is used to load the right Vault configuration and implementation when referencing secrets with the other entities.
	Prefix *string `json:"prefix,omitempty"`
	// An optional set of strings associated with the Vault for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (o *CreateVault) GetConfig() any {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateVault) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateVault) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateVault) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *CreateVault) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
