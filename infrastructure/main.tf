terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 3.0"
    }
  }

  backend "gcs" {
    bucket = "peter-built-terraform"
    prefix = "dev/main"
  }
}

provider "google" {
  project = "peter-built"
  region  = var.gcp_region
}

data "google_secret_manager_secret_version" "secret_variables" {
  project = "peter-built"
  secret  = google_secret_manager_secret.cloudflare_api_token.secret_id
  version = "latest"
}

provider "cloudflare" {
  api_token = data.google_secret_manager_secret_version.secret_variables.secret_data
}