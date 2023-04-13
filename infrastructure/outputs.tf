output "interactions_api_endpoint" {
  value = google_cloud_run_service.interactions_api.status[0].url
}