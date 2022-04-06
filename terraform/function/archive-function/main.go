package archive_function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		log.Printf("json.NewDecoder: %v", err)
		fmt.Fprint(w, "error")
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, event.PlanJsonApiUrl, nil)
	req.Header.Set("Authorization", "Bearer " + event.AccessToken)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "error")
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "error")
		return
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "error")
		return
	}
	log.Print(string(bodyBytes))

	req, err = http.NewRequest(http.MethodPatch, event.TaskResultCallbackUrl, bytes.NewBuffer([]byte("{ \"data\": { \"type\": \"task-results\", \"attributes\": { \"status\": \"passed\", \"message\": \"Analytics Task\"} } }")))
	req.Header.Set("Content-Type", "application/vnd.api+json")
	req.Header.Set("Authorization", "Bearer " + event.AccessToken)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "error")
		return
	}

	resp, err = client.Do(req)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "error")
		return
	}
	bodyBytes, err = io.ReadAll(resp.Body)
	log.Println(string(bodyBytes))
	return
}

