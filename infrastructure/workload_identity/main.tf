terraform {
  backend "gcs" {
    bucket = "peter-built-terraform"
    prefix = "dev/workload_identity"
  }
}

provider "google" {
  project = var.project_id
}

module "service_accounts" {
  source     = "terraform-google-modules/service-accounts/google"
  version    = "~> 3.0"
  project_id = var.project_id

  names        = ["github-actions"]
  display_name = "GitHub Actions"
  description  = "Terraform managed account used by GitHub Actions pipelines"
  project_roles = [
    "${var.project_id}=>roles/storage.objectViewer",
  ]
}

module "gh_oidc" {
  source      = "terraform-google-modules/github-actions-runners/google//modules/gh-oidc"
  project_id  = var.project_id
  pool_id     = "medius-github-wif-pool"
  provider_id = "github-wif"
  sa_mapping = {
    "github-actions" = {
      sa_name   = "projects/${var.project_id}/serviceAccounts/github-actions@${var.project_id}.iam.gserviceaccount.com"
      attribute = "attribute.repository/${var.github_repository}"
    }
  }
}
