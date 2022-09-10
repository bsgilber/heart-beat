# Create the Cloud Run service
resource "google_cloud_run_service" "run_service" {
  name     = "health-check-service"
  location = var.region

  template {
    spec {
      containers {
        image = var.health_check_image
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  # Waits for the Cloud Run API to be enabled from main.tf
  depends_on = [google_project_service.run_api]
}

# Allow unauthenticated users to invoke the service
resource "google_cloud_run_service_iam_member" "run_all_users" {
  service  = google_cloud_run_service.run_service.name
  location = google_cloud_run_service.run_service.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}
