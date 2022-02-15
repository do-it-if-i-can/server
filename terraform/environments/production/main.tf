terraform {
  required_version = "0.14.3"
}

module "cloud_run_api" {
  source = "../../gcp/cloud_run"
  project = var.project
}