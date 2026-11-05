# A vault key encrypts the token vault data in this directory. Generate a key
# pair and upload the public half:
#
#   openssl genpkey -algorithm RSA -pkeyopt rsa_keygen_bits:2048 -out private.pem
#   openssl pkey -in private.pem -pubout -out public.pem
#
# The data plane reads the private half through the vault reference listed in
# the token vault datastore's keys (see ai_gateway_datastore.tf); only the
# public key goes into Konnect.
resource "konnect_identity_vault_key" "my_vault_key" {
  directory_id = konnect_identity_directory.my_directory.id
  name         = "my-vault-key"

  public_key = <<-EOT
    -----BEGIN PUBLIC KEY-----
    MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6eLB+91ZbuIxBLwkVH79
    MGgWDT43PA10ONxcu3k6lBerKJQ9wkO3kHqfa6ASoh1HaVTlxgRckD+KldCIfOOF
    F7upQhpqlasl0rqxT/ahGPtdudVyHpVfprR7qg5izk+VBqqQOKoIEmX/EPGUuy/Z
    ASxnQEWSCFNTL+h6TPc1lKK8kei+CsuIj8r+Zgrh32eJ/cxHMzAE0lokt0F84W/8
    VSNiOqoKblsfBBFaK91CVdS5IHfHbIdBKWVjEvCo+Ve4NKrGoLtEbfemIXn1PrlM
    dAQq/Uw+n39aOprLP/qZAgtzDrXy/6TtigcGrSlSRiFSvZBizbmWfJIysJcEBGmj
    twIDAQAB
    -----END PUBLIC KEY-----
  EOT
}
