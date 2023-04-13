variable "service_name" {
  description = "The name of the Cloud Run service"
  default     = "transcoder"
}

variable "gcp_region" {
  description = "The GCP region to run the Cloud Run service, default to us-east1 as Discord also runs here."
  default     = "us-east1"
}

variable "image_url" {
  description = "The image URL for a version of the transcoder API"
}
