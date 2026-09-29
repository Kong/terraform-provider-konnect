resource "konnect_ai_gateway_custom_policy" "my_aigatewaycustompolicy_installed" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  installed = {
    display_name = "Custom Policy - Installed plugin-2"
    name         = "tf-test-installed-custom-policy"
    schema       = "return {\n  name = \"tf-test-installed-custom-policy\",\n  fields = {\n    { protocols = require(\"kong.db.schema.typedefs\").protocols_http },\n    {\n      config = {\n        type = \"record\",\n        fields = {\n          { name = { description = \"The name of the header to set for test.\", type = \"string\", required = true } },\n          { value = { description = \"The value for the header.\", type = \"string\", required = true } }\n        }\n      }\n    }\n  }\n}"
    type         = "installed"
  }
}

resource "konnect_ai_gateway_custom_policy" "my_aigatewaycustompolicy_streaming" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  streaming = {
    display_name = "Custom Policy - Streaming plugin-2"
    handler      = "return {\n  VERSION = \"1.0\",\n  PRIORITY = 1004,\n}\n"
    name         = "tf-test-streaming-custom-policy"
    schema       = "return {\n  name = \"tf-test-streaming-custom-policy\",\n  fields = {\n    { protocols = require(\"kong.db.schema.typedefs\").protocols_http },\n    {\n      config = {\n        type = \"record\",\n        fields = {\n          { name = { description = \"The name of the header to set.\", type = \"string\", required = true } },\n          { value = { description = \"The value for the header.\", type = \"string\", required = true } }\n        }\n      }\n    }\n  }\n}"
    type         = "streaming"
  }
}