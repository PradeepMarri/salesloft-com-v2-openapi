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

func Get_v2_email_templates_jsonHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ids=%v", val))
		}
		if val, ok := args["updated_at"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("updated_at=%v", val))
		}
		if val, ok := args["linked_to_team_template"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("linked_to_team_template=%v", val))
		}
		if val, ok := args["search"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("search=%v", val))
		}
		if val, ok := args["tag_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tag_ids=%v", val))
		}
		if val, ok := args["tag"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tag=%v", val))
		}
		if val, ok := args["filter_by_owner"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter_by_owner=%v", val))
		}
		if val, ok := args["group_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("group_id=%v", val))
		}
		if val, ok := args["include_cadence_templates"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_cadence_templates=%v", val))
		}
		if val, ok := args["include_archived_templates"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_archived_templates=%v", val))
		}
		if val, ok := args["cadence_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cadence_id=%v", val))
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
		url := fmt.Sprintf("%s/v2/email_templates.json%s", cfg.BaseURL, queryString)
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
		var result []EmailTemplate
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

func CreateGet_v2_email_templates_jsonTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v2_email_templates_json",
		mcp.WithDescription("List email templates"),
		mcp.WithArray("ids", mcp.Description("IDs of email templates to fetch. If a record can't be found, that record won't be returned and your request will be successful")),
		mcp.WithArray("updated_at", mcp.Description("Equality filters that are applied to the updated_at field. A single filter can be used by itself or combined with other filters to create a range.\n\n---CUSTOM---\n{\"type\":\"object\",\"keys\":[{\"name\":\"gt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"gte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"}]}\n")),
		mcp.WithBoolean("linked_to_team_template", mcp.Description("Filters email templates by whether they are linked to a team template or not")),
		mcp.WithString("search", mcp.Description("Filters email templates by title or subject")),
		mcp.WithArray("tag_ids", mcp.Description("Filters email templates by tags applied to the template by tag ID, not to exceed 100 IDs")),
		mcp.WithArray("tag", mcp.Description("Filters email templates by tags applied to the template, not to exceed 100 tags")),
		mcp.WithBoolean("filter_by_owner", mcp.Description("Filters email templates by current authenticated user")),
		mcp.WithArray("group_id", mcp.Description("Filters email templates by groups applied to the template by group ID. Not to exceed 500 IDs. Returns templates that are assigned to any of the group ids.")),
		mcp.WithBoolean("include_cadence_templates", mcp.Description("Filters email templates based on whether or not the template has been used on a cadence")),
		mcp.WithBoolean("include_archived_templates", mcp.Description("Filters email templates to include archived templates or not")),
		mcp.WithArray("cadence_id", mcp.Description("Filters email templates to those belonging to the cadence. Not to exceed 100 IDs. If a record can't be found, that record won't be returned and your request will be successful")),
		mcp.WithString("sort_by", mcp.Description("Key to sort on, must be one of: created_at, updated_at, last_used_at. Defaults to updated_at")),
		mcp.WithString("sort_direction", mcp.Description("Direction to sort in, must be one of: ASC, DESC. Defaults to DESC")),
		mcp.WithNumber("per_page", mcp.Description("How many records to show per page in the range [1, 100]. Defaults to 25")),
		mcp.WithNumber("page", mcp.Description("The current page to fetch results from. Defaults to 1")),
		mcp.WithBoolean("include_paging_counts", mcp.Description("Whether to include total_pages and total_count in the metadata. Defaults to false")),
		mcp.WithBoolean("limit_paging_counts", mcp.Description("Specifies whether the max limit of 10k records should be applied to pagination counts. Affects the total_count and total_pages data")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_v2_email_templates_jsonHandler(cfg),
	}
}
