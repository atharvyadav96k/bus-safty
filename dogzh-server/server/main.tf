variable "project_id" {
  description = "The GCP Project ID"
  type        = string
}

variable "region" {
  description = "The GCP region to deploy to"
  type        = string
  default     = "us-central1"
}

variable "service_name" {
  description = "The name of the Cloud Run service"
  type        = string
  default     = "dogzh-server"
}


variable "container_image" {
  description = "The URI of the container image in Artifact Registry or GCR"
  type        = string
}

variable "port" {
  description = "Server port"
  type        = number
}

provider "google" {
  project = var.project_id
  region  = var.region
}

resource "google_project_service" "enabled_apis" {
  for_each = toset([
    "run.googleapis.com",
    "firestore.googleapis.com",
    "iam.googleapis.com"
  ])
  service            = each.key
  disable_on_destroy = false
}

data "google_service_account" "run_sa" {
  account_id = "${var.service_name}-sa"
}

resource "google_project_iam_member" "firestore_access" {
  project = var.project_id
  role    = "roles/datastore.user"
  member  = "serviceAccount:${data.google_service_account.run_sa.email}"
}

resource "google_cloud_run_v2_service" "server" {
  name     = var.service_name
  location = var.region
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    service_account = data.google_service_account.run_sa.email

    containers {
      image = var.container_image

      ports {
        container_port = var.port
      }

      resources {
        limits = {
          cpu    = "1"
          memory = "512Mi"
        }
      }

      env {
        name  = "GCP_PROJECT_ID"
        value = var.project_id
      }
      
      env {
        name  = "PORT"
        value = tostring(var.port)
      }
    }
  }

  depends_on = [google_project_service.enabled_apis]
}

resource "google_cloud_run_v2_service_iam_member" "public_access" {
  location = google_cloud_run_v2_service.server.location
  name     = google_cloud_run_v2_service.server.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}

output "service_url" {
  description = "The URL where your Go server is live"
  value       = google_cloud_run_v2_service.server.uri
}