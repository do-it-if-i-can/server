resource "google_cloud_run_service" "default" {
  name = "cloudrun-srv"
  location = "asia-northeast1"
  project = var.project

  template {
    spec {
      containers {
        image = "gcr.io/${var.project}/qin-todo-api:latest"
      }
    }
  }

  traffic {
    percent = 100
    latest_revision = true
  }
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "policy" {
  location    = google_cloud_run_service.default.location
  project     = google_cloud_run_service.default.project
  service     = google_cloud_run_service.default.name

  policy_data = data.google_iam_policy.noauth.policy_data
}