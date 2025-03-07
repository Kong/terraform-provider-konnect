resource "konnect_serverless_cloud_gateway" "my_serverlesscloudgateway" {
  cluster_cert     = "-----BEGIN CERTIFICATE----- MIICRDCCAa2gAwIBAgIBADANBgkqhkiG9w0BAQ0FADA/MQswCQYDVQQGEwJ1czEL MAkGA1UECAwCVFgxDTALBgNVBAoMBFRlc3QxFDASBgNVBAMMC2V4YW1wbGUuY29t MB4XDTI0MDQyNjA5NTA1OVoXDTI1MDQyNjA5NTA1OVowPzELMAkGA1UEBhMCdXMx CzAJBgNVBAgMAlRYMQ0wCwYDVQQKDARUZXN0MRQwEgYDVQQDDAtleGFtcGxlLmNv bTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA8FaJZmPsthBT1XkTyqUJiuQV 4p4KaLpNMioNQtIdeGKFXelmZlVfW0bfLGBgFmkwo19KIzFtOlITUjb0Qqlub2Dn TIPoDs7rXA8aw6umJu73Z6647U3+alxNCpwTuYOY2CJQ+HWEIuOuwAdtQkndEp9r 7ZWA2xLatQKBYEvEmykCAwEAAaNQME4wHQYDVR0OBBYEFGUznNeZK74vlA4bqKHb 706tyMwcMB8GA1UdIwQYMBaAFGUznNeZK74vlA4bqKHb706tyMwcMAwGA1UdEwQF MAMBAf8wDQYJKoZIhvcNAQENBQADgYEARmnu/2vUcmJYLlg86MN0prXGC3CGXsem fDtPF4SBPxfchdG7HJKywTloIiCBKGEQALkCHiJcQJNcSHmzH3/Qk+SrOJNH01gt HsKA4SNFJZR5fCRpT6USCukyE2Wlr+PWPscrFCWbLXhK4Ql/t0oog1255B10HqKk 1qDkNrzCd/o= -----END CERTIFICATE-----\n"
  cluster_cert_key = "-----BEGIN PRIVATE KEY----- MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBAPBWiWZj7LYQU9V5 E8qlCYrkFeKeCmi6TTIqDULSHXhihV3pZmZVX1tG3yxgYBZpMKNfSiMxbTpSE1I2 9EKpbm9g50yD6A7O61wPGsOrpibu92euuO1N/mpcTQqcE7mDmNgiUPh1hCLjrsAH bUJJ3RKfa+2VgNsS2rUCgWBLxJspAgMBAAECgYEAvA7qqozL/1ZdUu/P1cQ36E86 9L03ZeVJXFRdVgj2eGqW8vob3z00RUb6gE3VQhQDNALvDwSw9G6eoblQfgz31Hju sb+j6bGOm2BqzYrx6rpcgme7k9ScV0tEbtiBNX0E/ToHvNywHtdOBvDocN2wh42Z 6bS9um51H+SXR036mgUCQQD4T7WrJHL97Hj8TtHnTw895xWKaGn94H7ZQa2lo1nk 7CQ4Oi8rFX5tDdyV7UU6fekBWuhpmIhSGJhyHD7UThBjAkEA98ef9ey2Qx+j+R8S tgpgJAF3LVNJJicEHCS/Vltgc84X/vidVAMa2+TYPxPrrUjxBr0STCeB5wZhvvsB D8cOAwJBAJ5JqaQPUx1dDe7Ai/vooO20Dj4xu0c0QYha3sfU7qwIgDo7lO/g/ruj 93a3TscvlkXf3oHZ0ySKOzual86ciMMCQQDGOLgaWHVy+4QFTzt70I8bHuUFqKRT VlEuZqN/ZXijDFQcES5jwFwjYE8zHy+ioEDaIDXcIJsGhA98Zndx9M+bAkA4IFdx 4YIDhuk1MJAYPqVQs5szEF/0BGymLNVYlIox48bZg+TH3uXwTVRVySxvpRa8dd3O 0gHs3EIV6GFUl7ev -----END PRIVATE KEY-----\n"
  control_plane = {
    id     = "d32d905a-ed33-46a3-a093-d8f536af9a8a"
    prefix = "518da50bf7"
    region = "us"
  }
  labels = {
    key = "value",
  }
}