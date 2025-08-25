package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/salesloft-platform/mcp-server/config"
	"github.com/salesloft-platform/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_v2_bulk_jobs_bulk_jobs_id_job_dataHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		bulk_jobs_idVal, ok := args["bulk_jobs_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: bulk_jobs_id"), nil
		}
		bulk_jobs_id, ok := bulk_jobs_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: bulk_jobs_id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("status=%v", val))
		}
		if val, ok := args["id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("id=%v", val))
		}
		if val, ok := args["per_page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("per_page=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v2/bulk_jobs/%s/job_data%s", cfg.BaseURL, bulk_jobs_id, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result []BulkJobResult
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGet_v2_bulk_jobs_bulk_jobs_id_job_dataTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v2_bulk_jobs_bulk_jobs_id_job_data",
		mcp.WithDescription("List job data for a bulk job"),
		mcp.WithNumber("bulk_jobs_id", mcp.Required(), mcp.Description("The id for the bulk job to which the job data relates")),
		mcp.WithArray("status", mcp.Description("Filter by result status. Accepts multiple statuses. Each status must be one of pending, success, error, retrying")),
		mcp.WithObject("id", mcp.Description("Filter by id using comparison operators. Only supports greater than (gt) comparison (i.e. id[gt]=123)")),
		mcp.WithNumber("per_page", mcp.Description("How many records to show per page in the range [1, 100]. Defaults to 25")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_v2_bulk_jobs_bulk_jobs_id_job_dataHandler(cfg),
	}
}
