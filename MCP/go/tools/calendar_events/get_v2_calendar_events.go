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

func Get_v2_calendar_eventsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["start_time"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start_time=%v", val))
		}
		if val, ok := args["end_time"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("end_time=%v", val))
		}
		if val, ok := args["user_guid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("user_guid=%v", val))
		}
		if val, ok := args["calendar_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("calendar_id=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v2/calendar/events%s", cfg.BaseURL, queryString)
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
		var result []CalendarEvent
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

func CreateGet_v2_calendar_eventsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v2_calendar_events",
		mcp.WithDescription("List calendar events"),
		mcp.WithNumber("per_page", mcp.Description("How many records to show per page in the range [1, 100]. Defaults to 25")),
		mcp.WithNumber("page", mcp.Description("The current page to fetch results from. Defaults to 1")),
		mcp.WithBoolean("include_paging_counts", mcp.Description("Whether to include total_pages and total_count in the metadata. Defaults to false")),
		mcp.WithString("sort_by", mcp.Description("Key to sort on, must be one of: start_time. Defaults to start_time")),
		mcp.WithString("sort_direction", mcp.Description("Direction to sort in, must be one of: ASC, DESC. Defaults to DESC")),
		mcp.WithString("start_time", mcp.Description("Lower bound (inclusive) for a calendar event's end time to filter by.\nMust be in ISO 8601 format.\n\nExample: `2022-02-14T10:12:59+00:00`.\n")),
		mcp.WithString("end_time", mcp.Description("Upper bound (exclusive) for a calendar event's start time to filter by.\nMust be in ISO 8601 format.\n\nExample: `2022-02-14T10:12:59+00:00`.\n")),
		mcp.WithString("user_guid", mcp.Description("user_guid of the user who created or included as a guest to the event.\n")),
		mcp.WithString("calendar_id", mcp.Description("calendar_id of the user who created or included as a guest to the event.\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_v2_calendar_eventsHandler(cfg),
	}
}
