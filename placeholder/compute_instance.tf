# Create a single Compute Engine instance
resource "google_compute_instance" "default" {
  name         = "health-check-instance"
  machine_type = "f1-micro"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = google_compute_network.vpc_network.name

    access_config {
      # Include this section to give the VM an external IP address
    }
  }
}
