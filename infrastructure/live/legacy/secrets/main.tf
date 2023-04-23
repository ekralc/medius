terraform {
  backend "gcs" {
    bucket = "peter-built-terraform"
    prefix = "dev/secrets"
  }
}

provider "google" {
  project = var.project_id
}

resource "google_project_service" "secrets_manager" {
  service = "secretmanager.googleapis.com"
}

resource "google_secret_manager_secret" "discord_bot_token" {
  secret_id = "discord_bot_token"

  replication {
    user_managed {
      replicas {
        location = "us-east1"
      }
    }
  }

  depends_on = [
    google_project_service.secrets_manager
  ]
}

resource "google_secret_manager_secret" "cloudflare_api_token" {
  secret_id = "cloudflare_api_token"

  replication {
    user_managed {
      replicas {
        location = "us-east1"
      }
    }
  }

  depends_on = [
    google_project_service.secrets_manager
  ]
}