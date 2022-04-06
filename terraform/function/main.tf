data "archive_file" "function" {
  type        = "zip"
  source_dir = "${path.module}/archive-function"
  output_path = "${path.module}/archive-function.zip"
}

resource "random_id" "id" {
  byte_length = 8

  keepers = {
    hash = data.archive_file.function.output_md5
  }
}

resource "google_storage_bucket" "run_task_func_bucket" {
  name     = "run-task-func-bucket-${random_id.id.hex}"
  location = "US"
}

resource "google_storage_bucket_object" "run_task_func_object" {
  name   = "index.zip"
  bucket = google_storage_bucket.run_task_func_bucket.name
  source = "${path.module}/archive-function.zip"
}

resource "google_cloudfunctions_function" "function" {
  name        = "function-test"
  description = "My function"
  runtime     = "go116"
  region      = "us-east1"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.run_task_func_bucket.name
  source_archive_object = google_storage_bucket_object.run_task_func_object.name
  trigger_http          = true
  timeout               = 60
  entry_point           = "HandleTerraformCloudRunTask"

#  environment_variables = {
#    MY_ENV_VAR = "my-env-var-value"
#  }
}

resource "google_cloudfunctions_function_iam_member" "invoker" {
  region      = "us-east1"
  cloud_function = google_cloudfunctions_function.function.name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"
}