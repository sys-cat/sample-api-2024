provider "google" {
  project = var.project_id
  region = var.region
}

module "setup_artifactregistry" {
  source = "../../modules/artifact_registry/setup"
  project_id = var.project_id
  region = var.region
  artifact_registry_repository_id = var.artifact_registry_repository_id
}

module "setup_service_account" {
  source = "../../modules/iam/common_sa"
  project_id = var.project_id
  region = var.region
  account_id = var.sa_account_id
  display_name = var.sa_display_name
}