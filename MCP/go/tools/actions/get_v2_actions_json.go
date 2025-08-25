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

func Get_v2_actions_jsonHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ids=%v", val))
		}
		if val, ok := args["step_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("step_id=%v", val))
		}
		if val, ok := args["type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("type=%v", val))
		}
		if val, ok := args["due_on"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("due_on=%v", val))
		}
		if val, ok := args["user_guid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("user_guid=%v", val))
		}
		if val, ok := args["person_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("person_id=%v", val))
		}
		if val, ok := args["cadence_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cadence_id=%v", val))
		}
		if val, ok := args["multitouch_group_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("multitouch_group_id=%v", val))
		}
		if val, ok := args["updated_at"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("updated_at=%v", val))
		}
		if val, ok := args["sort_by"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_by=%v", val))
		}
		if val, ok := args["sort_direction"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_direction=%v", val))
		}
		if val, ok := args["per_page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("per_page=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		if val, ok := args["include_paging_counts"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_paging_counts=%v", val))
		}
		if val, ok := args["limit_paging_counts"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit_paging_counts=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v2/actions.json%s", cfg.BaseURL, queryString)
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
		var result []Action
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

func CreateGet_v2_actions_jsonTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v2_actions_json",
		mcp.WithDescription("List actions"),
		mcp.WithArray("ids", mcp.Description("IDs of actions to fetch.")),
		mcp.WithNumber("step_id", mcp.Description("Fetch actions by step ID")),
		mcp.WithString("type", mcp.Description("Filter actions by type")),
		mcp.WithArray("due_on", mcp.Description("Equality filters that are applied to the due_on field. A single filter can be used by itself or combined with other filters to create a range.\n\n---CUSTOM---\n{\"type\":\"object\",\"keys\":[{\"name\":\"gt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"gte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"}]}\n")),
		mcp.WithArray("user_guid", mcp.Description("Filters actions by the user's guid. Multiple user guids can be applied. The user must be a team admin to filter other users' actions")),
		mcp.WithArray("person_id", mcp.Description("Filters actions by person_id. Multiple person ids can be applied")),
		mcp.WithArray("cadence_id", mcp.Description("Filters actions by cadence_id. Multiple cadence ids can be applied")),
		mcp.WithArray("multitouch_group_id", mcp.Description("Filters actions by multitouch_group_id. Multiple multitouch group ids can be applied")),
		mcp.WithArray("updated_at", mcp.Description("Equality filters that are applied to the updated_at field. A single filter can be used by itself or combined with other filters to create a range.\n\n---CUSTOM---\n{\"type\":\"object\",\"keys\":[{\"name\":\"gt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"gte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"}]}\n")),
		mcp.WithString("sort_by", mcp.Description("Key to sort on, must be one of: created_at, updated_at. Defaults to updated_at")),
		mcp.WithString("sort_direction", mcp.Description("Direction to sort in, must be one of: ASC, DESC. Defaults to DESC")),
		mcp.WithNumber("per_page", mcp.Description("How many records to show per page in the range [1, 100]. Defaults to 25")),
		mcp.WithNumber("page", mcp.Description("The current page to fetch results from. Defaults to 1")),
		mcp.WithBoolean("include_paging_counts", mcp.Description("Whether to include total_pages and total_count in the metadata. Defaults to false")),
		mcp.WithBoolean("limit_paging_counts", mcp.Description("Specifies whether the max limit of 10k records should be applied to pagination counts. Affects the total_count and total_pages data")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_v2_actions_jsonHandler(cfg),
	}
}
