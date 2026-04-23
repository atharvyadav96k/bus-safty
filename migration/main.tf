terraform {
  backend "gcs" {
    bucket  = "dogzh-bucket"
    prefix  = "cloud-functions/migration"
  }

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.0"
    }
  }
}

variable "project_id" {
  type = string
}

variable "project_number" {
  type = string
}

variable "region" {
  type = string
}

variable "service_account" {
  type = string
}

variable "neon_database_url" {
  type = string
}

provider "google" {
  project = var.project_id
  region  = var.region
}

resource "random_id" "bucket_suffix" {
  byte_length = 4
}

resource "google_storage_bucket" "source_bucket" {
  name                        = "${var.project_id}-migration-function-${random_id.bucket_suffix.hex}"
  location                    = var.region
  uniform_bucket_level_access = true
  force_destroy               = true
}

resource "google_storage_bucket_object" "source_archive" {
  name   = "source-${filesha256("source.zip")}.zip"
  bucket = google_storage_bucket.source_bucket.name
  source = "source.zip"
}

resource "google_cloudfunctions2_function" "migration" {
  name     = "migration"
  location = var.region

  build_config {
    runtime     = "go122"
    entry_point = "Migration"
    service_account = "projects/${var.project_id}/serviceAccounts/${var.service_account}"
    source {
      storage_source {
        bucket = google_storage_bucket.source_bucket.name
        object = google_storage_bucket_object.source_archive.name
      }
    }
  }

  service_config {
    max_instance_count    = 1
    available_memory      = "512M"
    timeout_seconds       = 300
    service_account_email = var.service_account
    ingress_settings      = "ALLOW_ALL"

    environment_variables = {
      GCP_PROJECT_ID = var.project_id
      GCP_NEON_DATABASE_URL = var.neon_database_url
    }
  }
}

resource "google_cloud_run_service_iam_member" "public_access" {
  location = var.region
  service  = google_cloudfunctions2_function.migration.service_config[0].service
  role     = "roles/run.invoker"
  member   = "allUsers"

  depends_on = [
    google_cloudfunctions2_function.migration
  ]
}

output "function_url" {
  value = google_cloudfunctions2_function.migration.service_config[0].uri
}
