resource "google_service_account" "service_account" {
  // No roles required for now

  account_id   = "medius-transcoder"
  display_name = "Transcoder API"
  description  = "The service account used by the Cloud Run service for transcoding videos."
}

resource "google_service_account" "pubsub_invoker" {
  account_id   = "medius-transcoder-invoker"
  display_name = "Transcoder Pub/Sub Invoker"
  description  = "To be used by Pub/Sub to invoke the transcoder Cloud Run service"
}

resource "google_cloud_run_service_iam_binding" "binding" {
  location = google_cloud_run_service.transcoder.location
  service  = google_cloud_run_service.transcoder.name
  role     = "roles/run.invoker"
  members  = ["serviceAccount:${google_service_account.pubsub_invoker.email}"]
}
