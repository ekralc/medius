terraform {
  backend "gcs" {
    bucket = "peter-built-terraform"
    prefix = "dev/apis"
  }
}

provider "google" {
  project = var.gcp_project_id
  region  = "us-east1"
}

module "interactions" {
  source = "../modules/interactions"

  image_url          = var.interactions_api_image
  gcp_region         = "us-east1"
  discord_public_key = var.discord_public_key

  domain = "medius.phatchips.net"
}

module "transcoder" {
  source = "../modules/transcoder"

  image_url  = var.transcoder_api_image
  gcp_region = "us-east1"
}
