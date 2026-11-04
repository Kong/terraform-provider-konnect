resource "konnect_ai_gateway" "my_aigateway" {
  display_name = "TF Test AIGW - datastore vectordb"
  name         = "tf-test-aigw-datastore-vectordb"
}

resource "konnect_ai_gateway_datastore" "my_aigatewaydatastore_vectordb" {
  gateway_id = konnect_ai_gateway.my_aigateway.id
  vectordb = {
    name         = "tf-test-vectordb"
    display_name = "TF Test VectorDB Updated"
    config = {
      database = "kong-pgvector"
      host     = "127.0.0.1"
    }
  }
}
