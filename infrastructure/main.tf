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
  project = var.gcp_project_id
  region  = var.gcp_region
}

provider "cloudflare" {
  api_token = data.google_secret_manager_secret_version.secret_variables.secret_data
}

data "google_secret_manager_secret_version" "secret_variables" {
  project = var.gcp_project_id
  secret  = google_secret_manager_secret.cloudflare_api_token.secret_id
  version = "latest"
}

resource "cloudflare_record" "subdomain" {
  name    = var.interactions_domain
  type    = "CNAME"
  value   = "ghs.googlehosted.com"
  zone_id = var.cloudflare_zone_id
}
