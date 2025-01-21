provider "google" {
  project = var.project_id
  region = var.region
}

module "create_clour_run" {
  source = "../../../modules/cloud_run/setup"
  project_id = var.project_id
  region = var.region
  name = var.run_name
  image_url = var.artifact_registry_repository_url
}