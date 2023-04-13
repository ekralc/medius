variable "cloudflare_zone_id" {}

variable "interactions_domain" {
  description = "The domain name pointing to the interactions API"
}

variable "interactions_api_image" {
  description = "The URL of the image for the interactions API"
}

variable "transcoder_api_image" {
  description = "The URL of the image for the transcoder API"
}

variable "discord_public_key" {
  description = "The public key for the Discord application"
}

variable "gcp_region" {
  description = "The GCP region to use"
}