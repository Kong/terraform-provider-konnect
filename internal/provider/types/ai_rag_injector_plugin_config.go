// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AiRagInjectorPluginConfig struct {
	Embeddings        *AiRagInjectorPluginEmbeddings `tfsdk:"embeddings"`
	FetchChunksCount  types.Float64                  `tfsdk:"fetch_chunks_count"`
	InjectAsRole      types.String                   `tfsdk:"inject_as_role"`
	InjectTemplate    types.String                   `tfsdk:"inject_template"`
	StopOnFailure     types.Bool                     `tfsdk:"stop_on_failure"`
	Vectordb          *AiRagInjectorPluginVectordb   `tfsdk:"vectordb"`
	VectordbNamespace types.String                   `tfsdk:"vectordb_namespace"`
}
