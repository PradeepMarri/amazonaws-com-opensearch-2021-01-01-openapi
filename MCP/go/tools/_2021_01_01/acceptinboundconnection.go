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

func AcceptinboundconnectionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		ConnectionIdVal, ok := args["ConnectionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: ConnectionId"), nil
		}
		ConnectionId, ok := ConnectionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: ConnectionId"), nil
		}
		url := fmt.Sprintf("%s/2021-01-01/opensearch/cc/inboundConnection/%s/accept", cfg.BaseURL, ConnectionId)
		req, err := http.NewRequest("PUT", url, nil)
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
		var result models.AcceptInboundConnectionResponse
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

func CreateAcceptinboundconnectionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_2021-01-01_opensearch_cc_inboundConnection_ConnectionId_accept",
		mcp.WithDescription("Allows the destination Amazon OpenSearch Service domain owner to accept an inbound cross-cluster search connection request. For more information, see <a href="https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cross-cluster-search.html">Cross-cluster search for Amazon OpenSearch Service</a>."),
		mcp.WithString("ConnectionId", mcp.Required(), mcp.Description("The ID of the inbound connection to accept.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AcceptinboundconnectionHandler(cfg),
	}
}
