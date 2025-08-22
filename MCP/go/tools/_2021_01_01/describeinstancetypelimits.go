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

func DescribeinstancetypelimitsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		InstanceTypeVal, ok := args["InstanceType"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: InstanceType"), nil
		}
		InstanceType, ok := InstanceTypeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: InstanceType"), nil
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
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/2021-01-01/opensearch/instanceTypeLimits/%s/%s%s", cfg.BaseURL, InstanceType, EngineVersion, queryString)
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
		var result models.DescribeInstanceTypeLimitsResponse
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

func CreateDescribeinstancetypelimitsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_2021-01-01_opensearch_instanceTypeLimits_EngineVersion_InstanceType",
		mcp.WithDescription("Describes the instance count, storage, and master node limits for a given OpenSearch or Elasticsearch version and instance type."),
		mcp.WithString("domainName", mcp.Description("The name of the domain. Only specify if you need the limits for an existing domain.")),
		mcp.WithString("InstanceType", mcp.Required(), mcp.Description("The OpenSearch Service instance type for which you need limit information.")),
		mcp.WithString("EngineVersion", mcp.Required(), mcp.Description("Version of OpenSearch or Elasticsearch, in the format Elasticsearch_X.Y or OpenSearch_X.Y. Defaults to the latest version of OpenSearch.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DescribeinstancetypelimitsHandler(cfg),
	}
}
