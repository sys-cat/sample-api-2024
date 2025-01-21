resource "google_cloud_run_v2_service" "default" {
  name     = var.name
  location = var.region
  ingress = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      image = var.image_url
    }
  }
}