variable "service_name" {
  description = "The name of the Cloud Run service"
  default     = "discord-interactions"
}

variable "gcp_region" {
  description = "The GCP region to run the Cloud Run service, default to us-east1 as Discord also runs here."
  default     = "us-east1"
}

variable "domain" {
  description = "The domain name for the Cloud Run service (optional)"
  default     = null
}

variable "image_url" {
  description = "The image URL for a version of the interactions API"
}

variable "discord_public_key" {
  description = "The Discord application's public key, for authenticating requests."
}
