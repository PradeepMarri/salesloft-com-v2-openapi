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

func Post_v2_meetings_settings_searches_jsonHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["user_guids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("user_guids=%v", val))
		}
		if val, ok := args["updated_at"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("updated_at=%v", val))
		}
		if val, ok := args["calendar_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("calendar_type=%v", val))
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
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v2/meetings/settings/searches.json%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, nil)
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
		var result []MeetingSetting
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

func CreatePost_v2_meetings_settings_searches_jsonTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v2_meetings_settings_searches_json",
		mcp.WithDescription("List meeting settings"),
		mcp.WithArray("user_guids", mcp.Description("Filters meeting settings by array of user_guids")),
		mcp.WithArray("updated_at", mcp.Description("Equality filters that are applied to the updated_at field. A single filter can be used by itself or combined with other filters to create a range.\n\n---CUSTOM---\n{\"type\":\"object\",\"keys\":[{\"name\":\"gt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"gte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"}]}\n")),
		mcp.WithString("calendar_type", mcp.Description("Filters meeting settings by calendar type")),
		mcp.WithNumber("per_page", mcp.Description("How many records to show per page in the range [1, 100]. Defaults to 25")),
		mcp.WithNumber("page", mcp.Description("The current page to fetch results from. Defaults to 1")),
		mcp.WithBoolean("include_paging_counts", mcp.Description("Whether to include total_pages and total_count in the metadata. Defaults to false")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_v2_meetings_settings_searches_jsonHandler(cfg),
	}
}
