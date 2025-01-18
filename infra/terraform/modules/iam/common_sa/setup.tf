resource "google_service_account" "create_sa" {
  account_id = var.account_id
  display_name = var.display_name
}

# まとめない理由はどのリソースにアクセスするためのRole設定かを差別化するため
resource "google_project_iam_member" "artifact_registry_access" {
  project = var.project_id
  role = "roles/artifactregistry.reader"
  member = "serviceAccount:${google_service_account.create_sa.email}"
}
resource "google_project_iam_member" "cloud_build_access" {
  project = var.project_id
  role    = "roles/cloudbuild.builds.builder"
  member = "serviceAccount:${google_service_account.create_sa.email}"
}
resource "google_project_iam_member" "cloud_storage_access" {
  project = var.project_id
  role = "roles/storage.objectUser"
  member = "serviceAccount:${google_service_account.create_sa.email}"
}
resource "google_project_iam_member" "cloud_run_access" {
  project = var.project_id
  role = "roles/run.admin"
  member = "serviceAccount:${google_service_account.create_sa.email}"
}
