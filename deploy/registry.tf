resource "google_artifact_registry_repository" "artifact-repo" {
  location      = var.region
  repository_id = "gilbert-test-repo"
  description   = "example docker repository"
  format        = "DOCKER"
}


