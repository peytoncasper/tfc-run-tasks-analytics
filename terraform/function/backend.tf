terraform {
  cloud {
    organization = "tfc4b-peyton"

    workspaces {
      name = "run-task-archive-function"
    }
  }
}