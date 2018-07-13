resource "random_id" "suffix" {
  byte_length = 4
}

resource "random_id" "cloud_storage_account" {
  byte_length = 4
}

resource "google_service_account" "cloud_storage" {
  display_name = "Cloud Storage"
  account_id   = "cs-${random_id.cloud_storage_account.hex}"
}

resource "google_service_account_key" "cloud_storage" {
  service_account_id = "${google_service_account.cloud_storage.id}"
}

resource "google_project_iam_member" "cloud_storage" {
  project = "${var.project}"
  role    = "projects/${var.project}/roles/${google_project_iam_custom_role.cloud_storage_role.role_id}"
  member  = "serviceAccount:${google_service_account.cloud_storage.email}",
}

resource "google_storage_bucket" "director" {
  name          = "${var.env_name}-director-${random_id.suffix.hex}"
  storage_class = "REGIONAL"
  location = "${var.region}"
  force_destroy = true
}

resource "google_storage_bucket" "buildpacks" {
  name          = "${var.env_name}-buildpacks-${random_id.suffix.hex}"
  storage_class = "REGIONAL"
  location = "${var.region}"
  force_destroy = true
}

resource "google_storage_bucket" "droplets" {
  name          = "${var.env_name}-droplets-${random_id.suffix.hex}"
  storage_class = "REGIONAL"
  location = "${var.region}"
  force_destroy = true
}

resource "google_storage_bucket" "packages" {
  name          = "${var.env_name}-packages-${random_id.suffix.hex}"
  storage_class = "REGIONAL"
  location = "${var.region}"
  force_destroy = true
}

resource "google_storage_bucket" "resources" {
  name          = "${var.env_name}-resources-${random_id.suffix.hex}"
  storage_class = "REGIONAL"
  location = "${var.region}"
  force_destroy = true
}

resource "google_project_iam_custom_role" "cloud_storage_role" {
  role_id     = "cloud_storage_role"
  title       = "cloud storage"
  description = "Cloud Storage Role"
  permissions = [
    "storage.buckets.get"
  ]
}

resource "google_storage_bucket_iam_member" "buildpacks" {
  bucket = "${google_storage_bucket.buildpacks.name}"
  role        = "roles/storage.objectAdmin"
  member      = "serviceAccount:${google_service_account.cloud_storage.email}"
}

resource "google_storage_bucket_iam_member" "resources" {
  bucket = "${google_storage_bucket.resources.name}"
  role        = "roles/storage.objectAdmin"
  member      = "serviceAccount:${google_service_account.cloud_storage.email}"
}

resource "google_storage_bucket_iam_member" "packages" {
  bucket = "${google_storage_bucket.packages.name}"
  role        = "roles/storage.objectAdmin"
  member      = "serviceAccount:${google_service_account.cloud_storage.email}"
}

resource "google_storage_bucket_iam_member" "droplets" {
  bucket = "${google_storage_bucket.droplets.name}"
  role        = "roles/storage.objectAdmin"
  member      = "serviceAccount:${google_service_account.cloud_storage.email}"
}

resource "google_storage_bucket_iam_member" "director" {
  bucket = "${google_storage_bucket.director.name}"
  role        = "roles/storage.objectAdmin"
  member      = "serviceAccount:${google_service_account.cloud_storage.email}"
}
