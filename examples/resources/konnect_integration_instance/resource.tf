resource "konnect_integration_instance" "my_integrationinstance" {
  config           = "{ \"see\": \"documentation\" }"
  description      = "...my_description..."
  display_name     = "AWS (prod)"
  integration_name = "aws-lambda"
  labels = {
    key = "value"
  }
  name = "aws-lambda-prod"
}