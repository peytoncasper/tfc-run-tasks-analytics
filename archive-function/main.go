package archive_function

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)
type RunTaskEvent struct {
	PayloadVersion int `json:"payload_version"`
	AccessToken string `json:"access_token"`
	Stage string `json:"stage"`
	IsSpeculative bool `json:"is_speculative"`
	TaskResultId string `json:"task_result_id"`
	TaskResultEnforcementLevel string `json:"task_result_enforcement_level"`
	TaskResultCallbackUrl string `json:"task_result_callback_url"`
	RunAppUrl string `json:"run_app_url"`
	RunId string `json:"run_id"`
	RunMessage string `json:"run_message"`
	RunCreatedAt time.Time `json:"run_created_at"`
	RunCreatedBy string `json:"run_created_by"`
	WorkspaceId string `json:"workspace_id"`
	WorkspaceName string `json:"workspace_name"`
	WorkspaceAppUrl string `json:"workspace_app_url"`
	OrganizationName string `json:"organization_name"`
	PlanJsonApiUrl string `json:"plan_json_api_url"`
	VcsRepoUrl string `json:"vcs_repo_url"`
	VcsBranch string `json:"vcs_branch"`
	VcsPullRequestUrl string `json:"vcs_pull_request_url"`
	VcsCommitUrl string `json:"vcs_commit_url"`
}

// GOOGLE_CLOUD_PROJECT is a user-set environment variable.
var projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")

func init() {

}


func HandleTerraformCloudRunTask(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the run task event
	event := RunTaskEvent{}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Printf("json.NewDecoder: %v", err)
		http.Error(w, "Error parsing request", http.StatusBadRequest)
		return
	}

	log.Printf(event.TaskResultId)

	w.WriteHeader(http.StatusOK)
	return
}