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

func Get_v2_people_jsonHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["email_addresses"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("email_addresses=%v", val))
		}
		if val, ok := args["owned_by_guid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("owned_by_guid=%v", val))
		}
		if val, ok := args["person_stage_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("person_stage_id=%v", val))
		}
		if val, ok := args["crm_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("crm_id=%v", val))
		}
		if val, ok := args["owner_crm_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("owner_crm_id=%v", val))
		}
		if val, ok := args["do_not_contact"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("do_not_contact=%v", val))
		}
		if val, ok := args["can_email"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("can_email=%v", val))
		}
		if val, ok := args["can_call"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("can_call=%v", val))
		}
		if val, ok := args["can_text"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("can_text=%v", val))
		}
		if val, ok := args["account_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("account_id=%v", val))
		}
		if val, ok := args["custom_fields"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("custom_fields=%v", val))
		}
		if val, ok := args["import_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("import_id=%v", val))
		}
		if val, ok := args["job_seniority"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("job_seniority=%v", val))
		}
		if val, ok := args["tag_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tag_id=%v", val))
		}
		if val, ok := args["owner_is_active"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("owner_is_active=%v", val))
		}
		if val, ok := args["cadence_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cadence_id=%v", val))
		}
		if val, ok := args["starred_by_guid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("starred_by_guid=%v", val))
		}
		if val, ok := args["replied"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("replied=%v", val))
		}
		if val, ok := args["bounced"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("bounced=%v", val))
		}
		if val, ok := args["success"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("success=%v", val))
		}
		if val, ok := args["eu_resident"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("eu_resident=%v", val))
		}
		if val, ok := args["title"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("title=%v", val))
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
		if val, ok := args["last_contacted"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("last_contacted=%v", val))
		}
		if val, ok := args["created_at"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("created_at=%v", val))
		}
		if val, ok := args["new"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("new=%v", val))
		}
		if val, ok := args["phone_number"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("phone_number=%v", val))
		}
		if val, ok := args["locales"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("locales=%v", val))
		}
		if val, ok := args["owner_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("owner_id=%v", val))
		}
		if val, ok := args["_query"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("_query=%v", val))
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
		url := fmt.Sprintf("%s/v2/people.json%s", cfg.BaseURL, queryString)
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
		var result []Person
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

func CreateGet_v2_people_jsonTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v2_people_json",
		mcp.WithDescription("List people"),
		mcp.WithArray("ids", mcp.Description("IDs of people to fetch. If a record can't be found, that record won't be returned and your request will be successful")),
		mcp.WithArray("updated_at", mcp.Description("Equality filters that are applied to the updated_at field. A single filter can be used by itself or combined with other filters to create a range.\n\n---CUSTOM---\n{\"type\":\"object\",\"keys\":[{\"name\":\"gt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"gte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"}]}\n")),
		mcp.WithArray("email_addresses", mcp.Description("Filters people by email address. Multiple emails can be applied. An additional value of \"_is_null\" can be passed to filter people that do not have an email address.")),
		mcp.WithArray("owned_by_guid", mcp.Description("Filters people by the owner's guid. Multiple owner guids can be applied")),
		mcp.WithArray("person_stage_id", mcp.Description("Includes people that have a given person_stage. Multiple person stage ids can be applied. An additional value of \"_is_null\" can be passed to filter people that do not have a stage set.")),
		mcp.WithArray("crm_id", mcp.Description("Filters people by crm_id. Multiple crm ids can be applied")),
		mcp.WithArray("owner_crm_id", mcp.Description("Filters people by owner_crm_id. Multiple owner_crm_ids can be applied. An additional value of \"_is_null\" can be passed to filter people that are unowned. A \"_not_in\" modifier can be used to exclude specific owner_crm_ids. Example: v2/people?owner_crm_id[_not_in]=id")),
		mcp.WithBoolean("do_not_contact", mcp.Description("Includes people that have a given do_not_contact property")),
		mcp.WithBoolean("can_email", mcp.Description("Includes people that can be emailed given do_not_contact and contact_restrictions property")),
		mcp.WithBoolean("can_call", mcp.Description("Includes people that can be called given do_not_contact and contact_restrictions property")),
		mcp.WithBoolean("can_text", mcp.Description("Includes people that can be sent a text message given do_not_contact and contact_restrictions property")),
		mcp.WithArray("account_id", mcp.Description("Filters people by the account they are linked to. Multiple account ids can be applied")),
		mcp.WithObject("custom_fields", mcp.Description("Filters by people matching all given custom fields. The custom field names are case-sensitive, but the provided values are case-insensitive. Example: v2/people?custom_fields[custom_field_name]=custom_field_value")),
		mcp.WithArray("import_id", mcp.Description("Filters people that were imported by the given import ids. Multiple import ids can be applied. An additional value of \"_is_null\" can be passed to filter people that were not imported.")),
		mcp.WithArray("job_seniority", mcp.Description("Filters people by job seniorty. Multiple job seniorities can be applied. An additional value of \"_is_null\" can be passed to filter people do not have a job_seniority.")),
		mcp.WithArray("tag_id", mcp.Description("Filters people by the tag ids applied to the person. Multiple tag ids can be applied.")),
		mcp.WithBoolean("owner_is_active", mcp.Description("Filters people by whether the owner is active or not.")),
		mcp.WithArray("cadence_id", mcp.Description("Filters people by the cadence that they are currently on. Multiple cadence_ids can be applied. An additional value of \"_is_null\" can be passed to filter people that are not on a cadence.")),
		mcp.WithArray("starred_by_guid", mcp.Description("Filters people who have been starred by the user guids given.")),
		mcp.WithBoolean("replied", mcp.Description("Filters people by whether or not they have replied to an email or not.")),
		mcp.WithBoolean("bounced", mcp.Description("Filters people by whether an email that was sent to them bounced or not.")),
		mcp.WithBoolean("success", mcp.Description("Filters people by whether or not they have been marked as a success or not.")),
		mcp.WithBoolean("eu_resident", mcp.Description("Filters people by whether or not they are marked as an European Union Resident or not.")),
		mcp.WithArray("title", mcp.Description("Filters people by their title by exact match. Supports partial matching")),
		mcp.WithArray("country", mcp.Description("Filters people by their country by exact match. Supports partial matching")),
		mcp.WithArray("state", mcp.Description("Filters people by their state by exact match. Supports partial matching")),
		mcp.WithArray("city", mcp.Description("Filters people by their city by exact match. Supports partial matching")),
		mcp.WithObject("last_contacted", mcp.Description("Equality filters that are applied to the last_contacted field. A single filter can be used by itself or combined with other filters to create a range.\nAdditional values of \"_is_null\" or \"_is_not_null\" can be passed to filter records that either have no timestamp value or any timestamp value.\n---CUSTOM---\n{\"type\":\"object\",\"keys\":[{\"name\":\"gt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"gte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"}]}\n")),
		mcp.WithObject("created_at", mcp.Description("Equality filters that are applied to the last_contacted field. A single filter can be used by itself or combined with other filters to create a range.\n\n---CUSTOM---\n{\"type\":\"object\",\"keys\":[{\"name\":\"gt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"gte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are greater than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lt\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than the provided iso8601 timestamp. The comparison is done using microsecond precision.\"},{\"name\":\"lte\",\"type\":\"iso8601 string\",\"description\":\"Returns all matching records that are less than or equal to the provided iso8601 timestamp. The comparison is done using microsecond precision.\"}]}\n")),
		mcp.WithBoolean("new", mcp.Description("Filters people by whether or not that person is on a cadence or if they have been contacted in any way.")),
		mcp.WithBoolean("phone_number", mcp.Description("Filter people by whether or not they have a phone number or not")),
		mcp.WithArray("locales", mcp.Description("Filters people by locales. Multiple locales can be applied. An additional value of \"Null\" can be passed to filter people that do not have a locale.")),
		mcp.WithArray("owner_id", mcp.Description("Filters people by owner_id. Multiple owner_ids can be applied.")),
		mcp.WithString("_query", mcp.Description("For internal use only. This field does not comply with our backwards compatibility policies. This filter is for authenticated users of Salesloft only and will not work for OAuth Applications. Filters people by the string provided. Can search and filter by name, title, industry, email_address and linked account name.")),
		mcp.WithString("sort_by", mcp.Description("Key to sort on, must be one of: created_at, updated_at, last_contacted_at, name, title, job_seniority, call_count, sent_emails, clicked_emails, replied_emails, viewed_emails, account, cadence_stage_name. Defaults to updated_at")),
		mcp.WithString("sort_direction", mcp.Description("Direction to sort in, must be one of: ASC, DESC. Defaults to DESC")),
		mcp.WithNumber("per_page", mcp.Description("How many records to show per page in the range [1, 100]. Defaults to 25")),
		mcp.WithNumber("page", mcp.Description("The current page to fetch results from. Defaults to 1")),
		mcp.WithBoolean("include_paging_counts", mcp.Description("Whether to include total_pages and total_count in the metadata. Defaults to false")),
		mcp.WithBoolean("limit_paging_counts", mcp.Description("Specifies whether the max limit of 10k records should be applied to pagination counts. Affects the total_count and total_pages data")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_v2_people_jsonHandler(cfg),
	}
}
