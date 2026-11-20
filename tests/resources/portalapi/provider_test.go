package portalapi

import "github.com/kong/terraform-provider-konnect/v3/tests/internal/acc"

// Shared provider setup, see tests/internal/acc.
var (
	providerFactory       = acc.ProviderFactory
	providerConfigUs      = acc.ProviderConfigUs
	providerConfigFromEnv = acc.ProviderConfigFromEnv
)
