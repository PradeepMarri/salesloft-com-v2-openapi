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

func Get_v2_accounts_jsonHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ids=%v", val))
		}
		if val, ok := args["crm_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("crm_id=%v", val))
		}
		if val, ok := args["tag"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tag=%v", val))
		}
		if val, ok := args["tag_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tag_id=%v", val))
		}
		if val, ok := args["created_at"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("created_at=%v", val))
		}
		if val, ok := args["updated_at"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("updated_at=%v", val))
		}
		if val, ok := args["domain"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("domain=%v", val))
		}
		if val, ok := args["website"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("website=%v", val))
		}
		if val, ok := args["archived"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("archived=%v", val))
		}
		if val, ok := args["name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("name=%v", val))
		}
		if val, ok := args["account_stage_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("account_stage_id=%v", val))
		}
		if val, ok := args["account_tier_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("account_tier_id=%v", val))
		}
		if val, ok := args["owner_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("owner_id=%v", val))
		}
		if val, ok := args["owner_is_active"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("owner_is_active=%v", val))
		}
		if val, ok := args["last_contacted"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("last_contacted=%v", val))
		}
		if val, ok := args["custom_fields"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("custom_fields=%v", val))
		}
		if val, ok := args["industry"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("industry=%v", val))
		}
		if val, ok := args["country"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("country=%v", val))
		}
		if val, ok := args["state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("state=%v", val))
		}
		if val, ok := args["city"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("city=%v", val))
		}
		if val, ok := args["owner_crm_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("owner_crm_id=%v", val))
		}
		if val, ok := args["locales"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("locales=%v", val))
		}
		if val, ok := args["user_relationships"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("user_relationships=%v", val))
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
		url := fmt.Sprintf("%s/v2/accounts.json%s", cfg.BaseURL, queryString)
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
		var result []Account
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

func CreateGet_v2_accounts_jsonTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v2_accounts_json",
		mcp.WithDescription("List accounts"),
		mcp.WithArray("ids", mcp.Description("IDs of accounts to fetch. If a record can't be found, that record won't be returned and your request will be successful")),
		mcp.WithArray("crm_id", mcp.Description("Filters accounts by crm_id. Multiple crm ids can be applied")),
		mcp.WithArray("tag", mcp.Description("Filters accounts by the tags applied to the account. Multiple tags can be applied")),
		mcp.WithArray("tag_id", mcp.Description("Filters accounts by the tag id's applied to the account. Multiple tag id's can be applied")),
		mcp.WithArray("created_at", mcp.Description("Equality filters that are applied to the created_at field. A single filter can be used by itself or combined with other filters to create a range.\n\n---CUSTOM---\n{\"type\":\"object\",\"keys\":[{\"name\":\"gt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"gte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"}]}\n")),
		mcp.WithArray("updated_at", mcp.Description("Equality filters that are applied to the updated_at field. A single filter can be used by itself or combined with other filters to create a range.\n\n---CUSTOM---\n{\"type\":\"object\",\"keys\":[{\"name\":\"gt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"gte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"}]}\n")),
		mcp.WithString("domain", mcp.Description("Domain of the accounts to fetch. Domains are unique and lowercase")),
		mcp.WithArray("website", mcp.Description("Filters accounts by website. Multiple websites can be applied. An additional value of \"_is_null\" can be passed to filter accounts that do not have a website.")),
		mcp.WithBoolean("archived", mcp.Description("Filters accounts by archived_at status. Returns only accounts where archived_at is not null if this field is true. Returns only accounts where archived_at is null if this field is false. Do not pass this parameter to return both archived and unarchived accounts. This filter is not applied if any value other than \"true\" or \"false\" is passed.")),
		mcp.WithArray("name", mcp.Description("Names of accounts to fetch. Name matches are exact and case sensitive. Multiple names can be fetched.")),
		mcp.WithArray("account_stage_id", mcp.Description("Filters accounts by account_stage_id. Multiple account_stage_ids can be applied. An additional value of \"_is_null\" can be passed to filter accounts that do not have account_stage_id")),
		mcp.WithArray("account_tier_id", mcp.Description("Filters accounts by account_tier_id. Multiple account tier ids can be applied")),
		mcp.WithArray("owner_id", mcp.Description("Filters accounts by owner_id. Multiple owner_ids can be applied. An additional value of \"_is_null\" can be passed to filter accounts that are unowned")),
		mcp.WithBoolean("owner_is_active", mcp.Description("Filters accounts by whether the owner is active or not.")),
		mcp.WithObject("last_contacted", mcp.Description("Equality filters that are applied to the last_contacted field. A single filter can be used by itself or combined with other filters to create a range.\nAdditional values of \"_is_null\" or \"_is_not_null\" can be passed to filter records that either have no timestamp value or any timestamp value.\n---CUSTOM---\n{\"type\":\"object\",\"keys\":[{\"name\":\"gt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"gte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"}]}\n")),
		mcp.WithObject("custom_fields", mcp.Description("Filters by accounts matching all given custom fields. The custom field names are case-sensitive, but the provided values are case-insensitive. Example: v2/accounts?custom_fields[custom_field_name]=custom_field_value")),
		mcp.WithArray("industry", mcp.Description("Filters accounts by industry by exact match. Supports partial matching")),
		mcp.WithArray("country", mcp.Description("Filters accounts by country by exact match. Supports partial matching")),
		mcp.WithArray("state", mcp.Description("Filters accounts by state by exact match. Supports partial matching")),
		mcp.WithArray("city", mcp.Description("Filters accounts by city by exact match. Supports partial matching")),
		mcp.WithArray("owner_crm_id", mcp.Description("Filters accounts by owner_crm_id. Multiple owner_crm_ids can be applied. An additional value of \"_is_null\" can be passed to filter accounts that are unowned. A \"_not_in\" modifier can be used to exclude specific owner_crm_ids. Example: v2/accounts?owner_crm_id[_not_in]=id")),
		mcp.WithArray("locales", mcp.Description("Filters accounts by locale. Multiple locales are allowed")),
		mcp.WithObject("user_relationships", mcp.Description("Filters by accounts matching all given user relationship fields, _is_null or _unmapped can be passed to filter accounts with null or unmapped user relationship values. Example: v2/accounts?user_relationships[name]=value")),
		mcp.WithString("sort_by", mcp.Description("Key to sort on, must be one of: created_at, updated_at, last_contacted_at, account_stage, account_stage_name, account_tier, account_tier_name, name, counts_people. Defaults to updated_at")),
		mcp.WithString("sort_direction", mcp.Description("Direction to sort in, must be one of: ASC, DESC. Defaults to DESC")),
		mcp.WithNumber("per_page", mcp.Description("How many records to show per page in the range [1, 100]. Defaults to 25")),
		mcp.WithNumber("page", mcp.Description("The current page to fetch results from. Defaults to 1")),
		mcp.WithBoolean("include_paging_counts", mcp.Description("Whether to include total_pages and total_count in the metadata. Defaults to false")),
		mcp.WithBoolean("limit_paging_counts", mcp.Description("Specifies whether the max limit of 10k records should be applied to pagination counts. Affects the total_count and total_pages data")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_v2_accounts_jsonHandler(cfg),
	}
}
