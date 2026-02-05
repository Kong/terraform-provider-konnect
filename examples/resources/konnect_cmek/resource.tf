resource "konnect_cmek" "my_cmek" {
  description = "My Key Description"
  id          = "default"
  key_arn     = "arn:aws:kms:us-east-1:123456789012:key/mrk-12345678123412341234123456789012"
  name        = "My KMS Key"
}