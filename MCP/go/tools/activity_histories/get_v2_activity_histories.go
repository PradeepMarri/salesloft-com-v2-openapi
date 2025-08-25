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

func Get_v2_activity_historiesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["per_page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("per_page=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		if val, ok := args["include_paging_counts"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_paging_counts=%v", val))
		}
		if val, ok := args["sort_by"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_by=%v", val))
		}
		if val, ok := args["sort_direction"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_direction=%v", val))
		}
		if val, ok := args["type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("type=%v", val))
		}
		if val, ok := args["_resource"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("_resource=%v", val))
		}
		if val, ok := args["occurred_at"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("occurred_at=%v", val))
		}
		if val, ok := args["pinned"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pinned=%v", val))
		}
		if val, ok := args["resource_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("resource_type=%v", val))
		}
		if val, ok := args["resource_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("resource_id=%v", val))
		}
		if val, ok := args["updated_at"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("updated_at=%v", val))
		}
		if val, ok := args["user_guid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("user_guid=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v2/activity_histories%s", cfg.BaseURL, queryString)
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
		var result models.ActivityHistory
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

func CreateGet_v2_activity_historiesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v2_activity_histories",
		mcp.WithDescription("List Past Activities"),
		mcp.WithNumber("per_page", mcp.Description("How many records to show per page in the range [1, 100]. Defaults to 25")),
		mcp.WithNumber("page", mcp.Description("The current page to fetch results from. Defaults to 1")),
		mcp.WithBoolean("include_paging_counts", mcp.Description("Whether to include total_pages and total_count in the metadata. Defaults to false")),
		mcp.WithString("sort_by", mcp.Description("Key to sort on, must be one of: occurred_at, updated_at. Defaults to occurred_at")),
		mcp.WithString("sort_direction", mcp.Description("Direction to sort in, must be one of: ASC, DESC. Defaults to DESC")),
		mcp.WithString("type", mcp.Description("Filter by the type of activity. Must be one of: added_to_cadence, completed_action, call, requested_email, sent_email, received_email, email_reply, note, success, dnc_event, residency_change, meeting, meeting_held, message_conversation, task, voicemail, opportunity_stage_change, opportunity_amount_change, opportunity_close_date_change. Can be provided as an array, or as an object of type[resource_type][]=type")),
		mcp.WithString("_resource", mcp.Description("For internal use only. This field does not comply with our backwards compatibility policies. This filter is for authenticated users of Salesloft only and will not work for OAuth Applications. Filter by the {resource_type, resource_id} of activity. Provide this in the format resource[]=person,1234")),
		mcp.WithObject("occurred_at", mcp.Description("Equality filters that are applied to the occurred_at field. A single filter can be used by itself or combined with other filters to create a range.\n---CUSTOM---\n{\"keys\":[{\"description\":\"Returns all matching records that are greater than the provided iso8601 timestamp. The comparison is done using microsecond precision.\",\"name\":\"gt\",\"type\":\"iso8601 string\"},{\"description\":\"Returns all matching records that are greater than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\",\"name\":\"gte\",\"type\":\"iso8601 string\"},{\"description\":\"Returns all matching records that are less than the provided iso8601 timestamp. The comparison is done using microsecond precision.\",\"name\":\"lt\",\"type\":\"iso8601 string\"},{\"description\":\"Returns all matching records that are less than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\",\"name\":\"lte\",\"type\":\"iso8601 string\"}],\"type\":\"object\"}\n")),
		mcp.WithBoolean("pinned", mcp.Description("Filter by the pinned status of activity. Must be 'true' or 'false'")),
		mcp.WithString("resource_type", mcp.Description("Filter by the resource type. A resource is a Salesloft object that the activity is attributed to. A valid resource types must be one of person, account, crm_opportunity. Can be provided as an array")),
		mcp.WithArray("resource_id", mcp.Description("Filter by the resource id. \"resource_type\" filter is required to use this filter.")),
		mcp.WithObject("updated_at", mcp.Description("Equality filters that are applied to the updated_at field. A single filter can be used by itself or combined with other filters to create a range.\n---CUSTOM---\n{\"keys\":[{\"description\":\"Returns all matching records that are greater than the provided iso8601 timestamp. The comparison is done using microsecond precision.\",\"name\":\"gt\",\"type\":\"iso8601 string\"},{\"description\":\"Returns all matching records that are greater than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\",\"name\":\"gte\",\"type\":\"iso8601 string\"},{\"description\":\"Returns all matching records that are less than the provided iso8601 timestamp. The comparison is done using microsecond precision.\",\"name\":\"lt\",\"type\":\"iso8601 string\"},{\"description\":\"Returns all matching records that are less than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\",\"name\":\"lte\",\"type\":\"iso8601 string\"}],\"type\":\"object\"}\n")),
		mcp.WithString("user_guid", mcp.Description("Filter activities by a user's guid.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_v2_activity_historiesHandler(cfg),
	}
}
