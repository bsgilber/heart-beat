variable "project_id" {
  description = "The project ID to deploy to"
  type        = string
  default     = "gilbert-learning-gcp-113"
}

variable "region" {
  type    = string
  default = "us-central1"
}

variable "health_check_image" {
  type    = string
  default = "us-central1-docker.pkg.dev/gilbert-learning-gcp-113/gilbert-test-repo/health-check:test"
}
