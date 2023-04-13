resource "google_service_account" "service_account" {
  // No roles required for now

  account_id   = "medius-interactions"
  display_name = "Interactions API"
  description  = "The service account used by the Cloud Run service for handling Discord interactions."
}

resource "google_project_iam_binding" "service_account" {
  project = data.google_project.project.project_id
  role    = "roles/pubsub.publisher"

  members = [
    "serviceAccount:${google_service_account.service_account.email}"
  ]
}
