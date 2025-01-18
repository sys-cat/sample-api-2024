resource "google_artifact_registry_repository" "create_repository" {
  repository_id = var.artifact_registry_repository_id
  description = "iamge for sample api "
  location = var.region
  project = var.project_id
  format = "DOCKER"
}