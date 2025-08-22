package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/amazon-opensearch-service/mcp-server/config"
	"github.com/amazon-opensearch-service/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DissociatepackageHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		PackageIDVal, ok := args["PackageID"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: PackageID"), nil
		}
		PackageID, ok := PackageIDVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: PackageID"), nil
		}
		DomainNameVal, ok := args["DomainName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: DomainName"), nil
		}
		DomainName, ok := DomainNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: DomainName"), nil
		}
		url := fmt.Sprintf("%s/2021-01-01/packages/dissociate/%s/%s", cfg.BaseURL, PackageID, DomainName)
		req, err := http.NewRequest("POST", url, nil)
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
		var result models.DissociatePackageResponse
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

func CreateDissociatepackageTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_2021-01-01_packages_dissociate_PackageID_DomainName",
		mcp.WithDescription("Removes a package from the specified Amazon OpenSearch Service domain. The package can't be in use with any OpenSearch index for the dissociation to succeed. The package is still available in OpenSearch Service for association later. For more information, see <a href="https://docs.aws.amazon.com/opensearch-service/latest/developerguide/custom-packages.html">Custom packages for Amazon OpenSearch Service</a>."),
		mcp.WithString("PackageID", mcp.Required(), mcp.Description("Internal ID of the package to dissociate from the domain. Use <code>ListPackagesForDomain</code> to find this value.")),
		mcp.WithString("DomainName", mcp.Required(), mcp.Description("Name of the domain to dissociate the package from.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DissociatepackageHandler(cfg),
	}
}
