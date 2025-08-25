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

func Get_v2_users_jsonHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ids=%v", val))
		}
		if val, ok := args["guid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("guid=%v", val))
		}
		if val, ok := args["group_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("group_id=%v", val))
		}
		if val, ok := args["role_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("role_id=%v", val))
		}
		if val, ok := args["search"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("search=%v", val))
		}
		if val, ok := args["active"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("active=%v", val))
		}
		if val, ok := args["visible_only"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("visible_only=%v", val))
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
		if val, ok := args["has_crm_user"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("has_crm_user=%v", val))
		}
		if val, ok := args["work_country"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("work_country=%v", val))
		}
		if val, ok := args["sort_by"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_by=%v", val))
		}
		if val, ok := args["sort_direction"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_direction=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v2/users.json%s", cfg.BaseURL, queryString)
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
		var result []User
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

func CreateGet_v2_users_jsonTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v2_users_json",
		mcp.WithDescription("List users"),
		mcp.WithArray("ids", mcp.Description("IDs of users to fetch. If a record can't be found, that record won't be returned and your request will be successful")),
		mcp.WithArray("guid", mcp.Description("Filters list to only include guids")),
		mcp.WithArray("group_id", mcp.Description("Filters users by group_id.  An additional value of \"_is_null\" can be passed to filter users that are not in a group")),
		mcp.WithArray("role_id", mcp.Description("Filters users by role_id")),
		mcp.WithString("search", mcp.Description("Space-separated list of keywords used to find case-insensitive substring matches against First Name, Last Name, or Email")),
		mcp.WithBoolean("active", mcp.Description("Filters users based on active attribute. Defaults to not applied")),
		mcp.WithBoolean("visible_only", mcp.Description("Defaults to true.\n\nWhen true, only shows users that are actionable based on the team's privacy settings.\nWhen false, shows all users on the team, even if you can't take action on that user. Deactivated users are also included when false.\n")),
		mcp.WithNumber("per_page", mcp.Description("How many users to show per page in the range [1, 100]. Defaults to 25.  Results are only paginated if the page parameter is defined")),
		mcp.WithNumber("page", mcp.Description("The current page to fetch users from. Defaults to returning all users")),
		mcp.WithBoolean("include_paging_counts", mcp.Description("Whether to include total_pages and total_count in the metadata. Defaults to false")),
		mcp.WithBoolean("has_crm_user", mcp.Description("Filters users based on if they have a crm user mapped or not.")),
		mcp.WithArray("work_country", mcp.Description("Filters users based on assigned work_country.")),
		mcp.WithString("sort_by", mcp.Description("Key to sort on, must be one of: id, email, name, group, role. Defaults to id")),
		mcp.WithString("sort_direction", mcp.Description("Direction to sort in, must be one of: ASC, DESC. Defaults to DESC")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_v2_users_jsonHandler(cfg),
	}
}
