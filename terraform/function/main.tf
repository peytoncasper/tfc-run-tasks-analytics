

resource "google_cloudfunctions_function" "function" {
  name        = "function-test"
  description = "My function"
  runtime     = "go116"

  available_memory_mb   = 128
  source_repository     = var.cloud_source_url
  trigger_http          = true
  timeout               = 60
  entry_point           = "HandleTerraformCloudRunTask"

#  environment_variables = {
#    MY_ENV_VAR = "my-env-var-value"
#  }
}