package tests

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/kong/terraform-provider-konnect/v3/internal/provider"
)

var (
	providerFactory = map[string]func() (tfprotov6.ProviderServer, error){
		"konnect": providerserver.NewProtocol6WithError(provider.New("999.99.9")()),
	}

	providerConfigUs = `provider "konnect" {
  server_url = "https://us.api.konghq.com"
}
`
)
