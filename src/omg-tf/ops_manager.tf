# Allow HTTP/S access to Ops Manager from the outside world if exposed
resource "google_compute_firewall" "ops-manager-external" {
  name        = "${var.env_name}-ops-manager-external"
  network     = "${google_compute_network.pcf-network.name}"
  target_tags = ["${var.env_name}-ops-manager-external"]

  allow {
    protocol = "tcp"
    ports    = ["443", "80", "22"]
  }

  source_ranges = ["0.0.0.0/0"]
}

resource "google_compute_image" "ops-manager-image" {
  count          = "${var.opsman_image_selflink != "" ? 0 : 1}"
  name           = "${var.env_name}-ops-manager-image"
  create_timeout = 20

  raw_disk {
    source = "${var.opsman_image_url}"
  }
}

resource "google_compute_instance" "ops-manager-internal" {
  count = "${var.opsman_external_ip != "" ? 0 : 1}"

  name           = "${var.env_name}-ops-manager"
  machine_type   = "${var.opsman_machine_type}"
  zone           = "${element(var.zones, 1)}"
  create_timeout = 10
  tags           = ["${var.env_name}-ops-manager", "${var.no_ip_instance_tag}"]

  boot_disk {
    initialize_params {
      image = "${var.opsman_image_selflink != "" ? var.opsman_image_selflink : google_compute_image.ops-manager-image.self_link}"
      size  = 250
      type  = "pd-ssd"
    }
  }

  network_interface {
    subnetwork = "${google_compute_subnetwork.management-subnet.name}"
    address    = "10.0.0.6"
  }

  metadata = {
    ssh-keys               = "${format("ubuntu:%s", var.ssh_public_key)}"
    block-project-ssh-keys = "TRUE"
  }
}

resource "google_compute_address" "ops-manager-external" {
  count = "${var.opsman_external_ip != "" ? 1 : 0}"
  name  = "${var.env_name}-ops-manager"
}

resource "google_compute_instance" "ops-manager-external" {
  count = "${var.opsman_external_ip != "" ? 1 : 0}"

  name           = "${var.env_name}-ops-manager"
  machine_type   = "${var.opsman_machine_type}"
  zone           = "${element(var.zones, 1)}"
  create_timeout = 10
  tags           = ["${var.env_name}-ops-manager", "${var.env_name}-ops-manager-external"]

  boot_disk {
    initialize_params {
      image = "${var.opsman_image_selflink != "" ? var.opsman_image_selflink : google_compute_image.ops-manager-image.self_link}"
      size  = 250
      type  = "pd-ssd"
    }
  }

  network_interface {
    subnetwork = "${google_compute_subnetwork.management-subnet.name}"
    address    = "10.0.0.6"

    access_config {
      nat_ip = "${google_compute_address.ops-manager-external.address}"
    }
  }

  service_account {
    email = "${google_service_account.ops_manager.email}"
    scopes = ["cloud-platform"]
  }

  metadata = {
    ssh-keys               = "${format("ubuntu:%s", var.ssh_public_key)}"
    block-project-ssh-keys = "TRUE"
  }
}

resource "random_id" "ops_manager_password_generator" {
  byte_length = 16
}

resource "random_id" "ops_manager_decryption_phrase_generator" {
  byte_length = 16
}

resource "random_id" "ops_manager_account" {
  byte_length = 4
}

resource "google_service_account" "ops_manager" {
  display_name = "Ops Manager"
  account_id   = "ops-${random_id.ops_manager_account.hex}"
}

resource "google_service_account_key" "ops_manager" {
  service_account_id = "${google_service_account.ops_manager.id}"
}

resource "google_project_iam_custom_role" "opsman_role" {
  role_id     = "opsman_role"
  title       = "opsman"
  description = "Ops Manager Role"
  permissions = [
    "compute.addresses.get",
    "compute.addresses.list",
    "compute.backendServices.get",
    "compute.backendServices.list",
    "compute.diskTypes.get",
    "compute.disks.delete",
    "compute.disks.list",
    "compute.disks.get",
    "compute.disks.createSnapshot",
    "compute.snapshots.create",
    "compute.disks.create",
    "compute.images.useReadOnly",
    "compute.globalOperations.get",
    "compute.images.delete",
    "compute.images.get",
    "compute.images.create",
    "compute.instanceGroups.get",
    "compute.instanceGroups.list",
    "compute.instanceGroups.update",
    "compute.instances.setMetadata",
    "compute.instances.setLabels",
    "compute.instances.setTags",
    "compute.instances.reset",
    "compute.instances.start",
    "compute.instances.list",
    "compute.instances.get",
    "compute.instances.delete",
    "compute.instances.create",
    "compute.subnetworks.use",
    "compute.subnetworks.useExternalIp",
    "compute.instances.detachDisk",
    "compute.instances.attachDisk",
    "compute.disks.use",
    "compute.instances.deleteAccessConfig",
    "compute.instances.addAccessConfig",
    "compute.addresses.use",
    "compute.machineTypes.get",
    "compute.regionOperations.get",
    "compute.zoneOperations.get",
    "compute.networks.get",
    "compute.subnetworks.get",
    "compute.snapshots.delete",
    "compute.snapshots.get",
    "compute.targetPools.list",
    "compute.targetPools.get",
    "compute.targetPools.addInstance",
    "compute.targetPools.removeInstance",
    "compute.instances.use",
    "storage.buckets.create",
    "storage.objects.create",
    "resourcemanager.projects.get",
    "compute.zones.list",
    "compute.subnetworks.get",
    "compute.subnetworks.list",
    "compute.instances.setServiceAccount",
    "compute.regionBackendServices.get",
    "compute.regionBackendServices.list"
  ]
}

resource "google_project_iam_member" "ops_manager" {
  project = "${var.project}"
  # role    = "roles/${google_project_iam_custom_role.opsman_role.role_id}"
  role    = "projects/ps-sding/roles/opsman_role"
  member  = "serviceAccount:${google_service_account.ops_manager.email}"
}

resource "google_project_iam_member" "service_account_user" {
  project = "${var.project}"
  role    = "roles/iam.serviceAccountUser"
  member  = "serviceAccount:${google_service_account.ops_manager.email}"
}
