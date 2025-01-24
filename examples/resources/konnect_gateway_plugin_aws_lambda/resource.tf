resource "konnect_gateway_plugin_aws_lambda" "my_gatewaypluginawslambda" {
  config = {
    aws_assume_role_arn       = "...my_aws_assume_role_arn..."
    aws_imds_protocol_version = "v1"
    aws_key                   = "...my_aws_key..."
    aws_region                = "...my_aws_region..."
    aws_role_session_name     = "...my_aws_role_session_name..."
    aws_secret                = "...my_aws_secret..."
    aws_sts_endpoint_url      = "...my_aws_sts_endpoint_url..."
    awsgateway_compatible     = true
    base64_encode_body        = false
    disable_https             = false
    empty_arrays_mode         = "correct"
    forward_request_body      = false
    forward_request_headers   = true
    forward_request_method    = true
    forward_request_uri       = true
    function_name             = "...my_function_name..."
    host                      = "...my_host..."
    invocation_type           = "RequestResponse"
    is_proxy_integration      = false
    keepalive                 = 6.97
    log_type                  = "Tail"
    port                      = 25235
    proxy_url                 = "...my_proxy_url..."
    qualifier                 = "...my_qualifier..."
    skip_large_bodies         = false
    timeout                   = 4.31
    unhandled_status          = 115
  }
  consumer = {
    id = "...my_id..."
  }
  consumer_group = {
    id = "...my_id..."
  }
  control_plane_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"
  enabled          = false
  id               = "...my_id..."
  instance_name    = "...my_instance_name..."
  ordering = {
    after = {
      access = [
        "..."
      ]
    }
    before = {
      access = [
        "..."
      ]
    }
  }
  protocols = [
    "tls_passthrough"
  ]
  route = {
    id = "...my_id..."
  }
  service = {
    id = "...my_id..."
  }
  tags = [
    "..."
  ]
}