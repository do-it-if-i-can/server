terraform {
  required_version = "~> 1.1.5"
}

module "cloud_run_api" {
  source = "../../gcp/cloud-run"
  project = "qin-todo-341312"
}