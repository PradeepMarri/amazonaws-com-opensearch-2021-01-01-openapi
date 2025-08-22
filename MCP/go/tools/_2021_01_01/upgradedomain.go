package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-opensearch-service/mcp-server/config"
	"github.com/amazon-opensearch-service/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func UpgradedomainHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/2021-01-01/opensearch/upgradeDomain", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.UpgradeDomainResponse
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

func CreateUpgradedomainTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_2021-01-01_opensearch_upgradeDomain",
		mcp.WithDescription("Allows you to either upgrade your Amazon OpenSearch Service domain or perform an upgrade eligibility check to a compatible version of OpenSearch or Elasticsearch."),
		mcp.WithObject("AdvancedOptions", mcp.Description("Input parameter: <p>Exposes native OpenSearch configuration values from <code>opensearch.yml</code>. The following advanced options are available: </p> <ul> <li> <p>Allows references to indexes in an HTTP request body. Must be <code>false</code> when configuring access to individual sub-resources. Default is <code>true</code>.</p> </li> <li> <p>Specifies the percentage of heap space allocated to field data. Default is unbounded.</p> </li> </ul> <p>For more information, see <a href=\"https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options\">Advanced cluster parameters</a>.</p>")),
		mcp.WithString("DomainName", mcp.Required(), mcp.Description("Input parameter: The name of an OpenSearch Service domain. Domain names are unique across the domains owned by an account within an Amazon Web Services Region.")),
		mcp.WithBoolean("PerformCheckOnly", mcp.Description("Input parameter: When true, indicates that an upgrade eligibility check needs to be performed. Does not actually perform the upgrade.")),
		mcp.WithString("TargetVersion", mcp.Required(), mcp.Description("Input parameter: OpenSearch or Elasticsearch version to which you want to upgrade, in the format Opensearch_X.Y or Elasticsearch_X.Y.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpgradedomainHandler(cfg),
	}
}
