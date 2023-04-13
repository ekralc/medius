resource "google_pubsub_topic" "videos" {
  name = "webm-videos"
}

resource "google_pubsub_subscription" "push_subscription" {
  name  = "cloud-run-transcode-subscription"
  topic = google_pubsub_topic.videos.name

  ack_deadline_seconds = 300

  push_config {
    push_endpoint = "${google_cloud_run_service.transcoder.status.0.url}/pubsub"

    oidc_token {
      service_account_email = google_service_account.pubsub_invoker.email
    }
  }
}

resource "google_cloud_run_service" "transcoder" {
  name     = var.service_name
  location = var.gcp_region

  template {
    spec {
      containers {
        image = var.image_url

        resources {
          limits = {
            "memory" : "1Gi",
            "cpu" : "2000m"
          }
        }

      }

      container_concurrency = 5

      service_account_name = google_service_account.service_account.email
    }

    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale" = 10
      }
    }
  }

  metadata {
    annotations = {
      "run.googleapis.com/ingress" = "internal"
    }
  }

  traffic {
    latest_revision = true
    percent         = 100
  }
}
