# top level, create the service account to attach policy to
resource "google_service_account" "service_account" {
  account_id   = "tf-deployment-account-id"
  display_name = "tf-deployment-account"
  project      = var.project_id
}

# create the policy and attach it to the policy above
resource "google_service_account_iam_binding" "admin_account_iam" {
  service_account_id = google_service_account.service_account.name
  role               = "roles/iam.serviceAccountUser"

  members = [
    "user:bsgilber@gmail.com",
  ]
}

# create an account key (to be used to auth with docker login to push to artifact registry)
resource "google_service_account_key" "service_account_key" {
  service_account_id = google_service_account.service_account.name
  public_key_type    = "TYPE_X509_PEM_FILE"
}
