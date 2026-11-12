resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - custom-policy"
  name         = "tf-test-aigw-custom-policy"
}

resource "konnect_ai_gateway_custom_policy" "my_aigatewaycustompolicy" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  installed = {
    display_name = "TF Test Custom Policy - Installed"
    name         = "tf-test-installed-custom-policy"
    schema       = "return {\n  name = \"tf-test-installed-custom-policy\",\n  fields = {\n    { protocols = require(\"kong.db.schema.typedefs\").protocols_http },\n    {\n      config = {\n        type = \"record\",\n        fields = {\n          { name = { description = \"The name of the header to set for test.\", type = \"string\", required = true } },\n          { value = { description = \"The value for the header.\", type = \"string\", required = true } }\n        }\n      }\n    }\n  }\n}"

  }
}
