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

func ListinstancetypedetailsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		EngineVersionVal, ok := args["EngineVersion"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: EngineVersion"), nil
		}
		EngineVersion, ok := EngineVersionVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: EngineVersion"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["domainName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("domainName=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["nextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nextToken=%v", val))
		}
		if val, ok := args["retrieveAZs"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("retrieveAZs=%v", val))
		}
		if val, ok := args["instanceType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("instanceType=%v", val))
		}
		if val, ok := args["MaxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("MaxResults=%v", val))
		}
		if val, ok := args["NextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("NextToken=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/2021-01-01/opensearch/instanceTypeDetails/%s%s", cfg.BaseURL, EngineVersion, queryString)
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
		var result models.ListInstanceTypeDetailsResponse
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

func CreateListinstancetypedetailsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_2021-01-01_opensearch_instanceTypeDetails_EngineVersion",
		mcp.WithDescription("Lists all instance types and available features for a given OpenSearch or Elasticsearch version."),
		mcp.WithString("EngineVersion", mcp.Required(), mcp.Description("The version of OpenSearch or Elasticsearch, in the format Elasticsearch_X.Y or OpenSearch_X.Y. Defaults to the latest version of OpenSearch.")),
		mcp.WithString("domainName", mcp.Description("The name of the domain.")),
		mcp.WithNumber("maxResults", mcp.Description("An optional parameter that specifies the maximum number of results to return. You can use <code>nextToken</code> to get the next page of results.")),
		mcp.WithString("nextToken", mcp.Description("If your initial <code>ListInstanceTypeDetails</code> operation returns a <code>nextToken</code>, you can include the returned <code>nextToken</code> in subsequent <code>ListInstanceTypeDetails</code> operations, which returns results in the next page.")),
		mcp.WithBoolean("retrieveAZs", mcp.Description("An optional parameter that specifies the Availability Zones for the domain.")),
		mcp.WithString("instanceType", mcp.Description("An optional parameter that lists information for a given instance type.")),
		mcp.WithString("MaxResults", mcp.Description("Pagination limit")),
		mcp.WithString("NextToken", mcp.Description("Pagination token")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListinstancetypedetailsHandler(cfg),
	}
}
