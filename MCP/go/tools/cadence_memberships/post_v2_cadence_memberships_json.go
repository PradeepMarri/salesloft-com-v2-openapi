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

func Post_v2_cadence_memberships_jsonHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["person_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("person_id=%v", val))
		}
		if val, ok := args["cadence_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cadence_id=%v", val))
		}
		if val, ok := args["user_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("user_id=%v", val))
		}
		if val, ok := args["step_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("step_id=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v2/cadence_memberships.json%s", cfg.BaseURL, queryString)
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
		var result models.CadenceMembership
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

func CreatePost_v2_cadence_memberships_jsonTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v2_cadence_memberships_json",
		mcp.WithDescription("Create a cadence membership"),
		mcp.WithNumber("person_id", mcp.Required(), mcp.Description("ID of the person to create a cadence membership for")),
		mcp.WithNumber("cadence_id", mcp.Required(), mcp.Description("ID of the cadence to create a cadence membership for")),
		mcp.WithNumber("user_id", mcp.Description("ID of the user to create a cadence membership for. The associated cadence must be owned by the user, or it must be a team cadence")),
		mcp.WithNumber("step_id", mcp.Description("ID of the step on which the person should start the cadence. Start on first step is the default behavior without this parameter.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_v2_cadence_memberships_jsonHandler(cfg),
	}
}
