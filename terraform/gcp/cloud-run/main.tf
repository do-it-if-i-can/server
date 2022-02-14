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
    parcent = 100
    latest_revision = true
  }
}

data "google_iam_policy" "admin" {
    binding {
    role = "roles/editor"
    members = [
      "user:s.kawabe2281@gmail.com",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "policy" {
  location    = google_cloud_run_service.default.location
  project     = google_cloud_run_service.default.project
  service     = google_cloud_run_service.default.name

  policy_data = data.google_iam_policy.admin.policy_data
}