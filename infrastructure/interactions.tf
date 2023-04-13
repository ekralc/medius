resource "google_cloud_run_service" "interactions_api" {
  name     = "discord-interactions"
  location = var.gcp_region

  template {
    spec {
      containers {
        image = var.interactions_api_image

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
    }
  }

  metadata {
    namespace = "peter-built"
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
}

resource "google_cloud_run_domain_mapping" "default" {
  name     = var.interactions_domain
  location = var.gcp_region

  spec {
    route_name = google_cloud_run_service.interactions_api.name
  }

  metadata {
    namespace = "peter-built"
  }
}

resource "cloudflare_record" "subdomain" {
  name    = var.interactions_domain
  type    = "CNAME"
  value   = "ghs.googlehosted.com"
  zone_id = var.cloudflare_zone_id
}