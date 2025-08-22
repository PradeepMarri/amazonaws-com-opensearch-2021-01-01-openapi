package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazon-opensearch-service/mcp-server/config"
	"github.com/amazon-opensearch-service/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ListvpcendpointaccessHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		DomainNameVal, ok := args["DomainName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: DomainName"), nil
		}
		DomainName, ok := DomainNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: DomainName"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["nextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nextToken=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/2021-01-01/opensearch/domain/%s/listVpcEndpointAccess%s", cfg.BaseURL, DomainName, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
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
		var result models.ListVpcEndpointAccessResponse
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

func CreateListvpcendpointaccessTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_2021-01-01_opensearch_domain_DomainName_listVpcEndpointAccess",
		mcp.WithDescription("Retrieves information about each Amazon Web Services principal that is allowed to access a given Amazon OpenSearch Service domain through the use of an interface VPC endpoint."),
		mcp.WithString("DomainName", mcp.Required(), mcp.Description("The name of the OpenSearch Service domain to retrieve access information for.")),
		mcp.WithString("nextToken", mcp.Description("If your initial <code>ListVpcEndpointAccess</code> operation returns a <code>nextToken</code>, you can include the returned <code>nextToken</code> in subsequent <code>ListVpcEndpointAccess</code> operations, which returns results in the next page.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListvpcendpointaccessHandler(cfg),
	}
}
