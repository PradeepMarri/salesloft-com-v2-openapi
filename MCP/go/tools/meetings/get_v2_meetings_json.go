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

func Get_v2_meetings_jsonHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ids=%v", val))
		}
		if val, ok := args["status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("status=%v", val))
		}
		if val, ok := args["person_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("person_id=%v", val))
		}
		if val, ok := args["account_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("account_id=%v", val))
		}
		if val, ok := args["person_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("person_ids=%v", val))
		}
		if val, ok := args["event_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("event_ids=%v", val))
		}
		if val, ok := args["i_cal_uids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("i_cal_uids=%v", val))
		}
		if val, ok := args["task_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("task_ids=%v", val))
		}
		if val, ok := args["include_meetings_settings"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_meetings_settings=%v", val))
		}
		if val, ok := args["start_time"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start_time=%v", val))
		}
		if val, ok := args["created_at"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("created_at=%v", val))
		}
		if val, ok := args["user_guids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("user_guids=%v", val))
		}
		if val, ok := args["show_deleted"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("show_deleted=%v", val))
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
		url := fmt.Sprintf("%s/v2/meetings.json%s", cfg.BaseURL, queryString)
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
		var result []Meeting
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

func CreateGet_v2_meetings_jsonTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v2_meetings_json",
		mcp.WithDescription("List meetings"),
		mcp.WithArray("ids", mcp.Description("IDs of meetings to fetch. If a record can't be found, that record won't be returned and your request will be successful")),
		mcp.WithString("status", mcp.Description("Filters meetings by status. Possible values are: pending, booked, failed, retry")),
		mcp.WithString("person_id", mcp.Description("Filters meetings by person_id. Multiple person ids can be applied")),
		mcp.WithString("account_id", mcp.Description("Filters meetings by account_id. Multiple account ids can be applied")),
		mcp.WithArray("person_ids", mcp.Description("Filters meetings by person_id. Multiple person ids can be applied")),
		mcp.WithArray("event_ids", mcp.Description("List of event IDs. If both event_ids and i_cal_uids params are passed, this filters will be ORed. If a record can't be found, that record won't be returned and your request will be successful")),
		mcp.WithArray("i_cal_uids", mcp.Description("List of UIDs provided by calendar provider. If both event_ids and i_cal_uids params are passed, this filters will be ORed. If a record can't be found, that record won't be returned and your request will be successful")),
		mcp.WithArray("task_ids", mcp.Description("Filters meetings by task_id. Multiple task ids can be applied")),
		mcp.WithBoolean("include_meetings_settings", mcp.Description("Flag to indicate whether to include owned_by_meetings_settings and booked_by_meetings_settings objects")),
		mcp.WithArray("start_time", mcp.Description("Equality filters that are applied to the start_time field. A single filter can be used by itself or combined with other filters to create a range.\n\n---CUSTOM---\n{\"type\":\"object\",\"keys\":[{\"name\":\"gt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"gte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"}]}\n")),
		mcp.WithArray("created_at", mcp.Description("Equality filters that are applied to the created_at field. A single filter can be used by itself or combined with other filters to create a range.\n\n---CUSTOM---\n{\"type\":\"object\",\"keys\":[{\"name\":\"gt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"gte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"}]}\n")),
		mcp.WithArray("user_guids", mcp.Description("Filters meetings by user_guid. Multiple user guids can be applied")),
		mcp.WithBoolean("show_deleted", mcp.Description("Whether to include deleted events in the result")),
		mcp.WithString("sort_by", mcp.Description("Key to sort on, must be one of: start_time, created_at, updated_at. Defaults to start_time")),
		mcp.WithString("sort_direction", mcp.Description("Direction to sort in, must be one of: ASC, DESC. Defaults to DESC")),
		mcp.WithNumber("per_page", mcp.Description("How many records to show per page in the range [1, 100]. Defaults to 25")),
		mcp.WithNumber("page", mcp.Description("The current page to fetch results from. Defaults to 1")),
		mcp.WithBoolean("include_paging_counts", mcp.Description("Whether to include total_pages and total_count in the metadata. Defaults to false")),
		mcp.WithBoolean("limit_paging_counts", mcp.Description("Specifies whether the max limit of 10k records should be applied to pagination counts. Affects the total_count and total_pages data")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_v2_meetings_jsonHandler(cfg),
	}
}
