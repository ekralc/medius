output "service_url" {
  value = google_cloud_run_service.transcoder.status.0.url
}
