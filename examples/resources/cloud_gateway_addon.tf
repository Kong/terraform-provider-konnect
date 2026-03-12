resource "konnect_gateway_control_plane" "test_cp" {
    name         = "CGW Control Plane"
    cloud_gateway = true
}

resource "konnect_cloud_gateway_addon" "my_addon" {
    name     = "My Add-on"

    config = {
        managed_cache = {
            capacity_config = {
                tiered = {
                    tier = "small"
                }
            }
        }
    }
    owner = {
        control_plane = {
            control_plane_geo = "us"
            control_plane_id  = konnect_gateway_control_plane.test_cp.id
        }
    }
}