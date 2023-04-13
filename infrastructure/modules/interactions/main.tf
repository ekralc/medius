resource "google_cloud_run_service" "interactions_api" {
  name     = var.service_name
  location = var.gcp_region

  template {
    spec {
      containers {
        image = var.image_url

        env {
          name  = "DISCORD_PUBLIC_KEY"
          value = var.discord_public_key
        }

        env {
          name  = "GIN_MODE"
          value = "release"
        }

        resources {
          limits = {
            "memory" : "512Mi",
            "cpu" : "1000m"
          }
        }
      }

      container_concurrency = 100

      service_account_name = google_service_account.service_account.email
    }
  }

  traffic {
    latest_revision = true
    percent         = 100
  }
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers"
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location = google_cloud_run_service.interactions_api.location
  project  = google_cloud_run_service.interactions_api.project
  service  = google_cloud_run_service.interactions_api.name

  policy_data = data.google_iam_policy.noauth.policy_data

  depends_on = [
    google_cloud_run_service.interactions_api
  ]
}

data "google_project" "project" {}

resource "google_cloud_run_domain_mapping" "default" {
  count = (var.domain != null) ? 1 : 0

  name     = var.domain
  location = var.gcp_region

  metadata {
    namespace = data.google_project.project.project_id
  }

  spec {
    route_name = google_cloud_run_service.interactions_api.name
  }
}
