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

func UpdatedomainconfigHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/2021-01-01/opensearch/domain/%s/config", cfg.BaseURL, DomainName)
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
		var result models.UpdateDomainConfigResponse
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

func CreateUpdatedomainconfigTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_2021-01-01_opensearch_domain_DomainName_config",
		mcp.WithDescription("Modifies the cluster configuration of the specified Amazon OpenSearch Service domain.sl"),
		mcp.WithString("DomainName", mcp.Required(), mcp.Description("The name of the domain that you're updating.")),
		mcp.WithBoolean("DryRun", mcp.Description("Input parameter: This flag, when set to True, specifies whether the <code>UpdateDomain</code> request should return the results of a dry run analysis without actually applying the change. A dry run determines what type of deployment the update will cause.")),
		mcp.WithObject("EBSOptions", mcp.Description("Input parameter: Container for the parameters required to enable EBS-based storage for an OpenSearch Service domain.")),
		mcp.WithObject("LogPublishingOptions", mcp.Description("Input parameter: Options to publish OpenSearch logs to Amazon CloudWatch Logs.")),
		mcp.WithObject("AdvancedOptions", mcp.Description("Input parameter: <p>Exposes native OpenSearch configuration values from <code>opensearch.yml</code>. The following advanced options are available: </p> <ul> <li> <p>Allows references to indexes in an HTTP request body. Must be <code>false</code> when configuring access to individual sub-resources. Default is <code>true</code>.</p> </li> <li> <p>Specifies the percentage of heap space allocated to field data. Default is unbounded.</p> </li> </ul> <p>For more information, see <a href=\"https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options\">Advanced cluster parameters</a>.</p>")),
		mcp.WithObject("SoftwareUpdateOptions", mcp.Description("Input parameter: Options for configuring service software updates for a domain.")),
		mcp.WithObject("VPCOptions", mcp.Description("Input parameter: Options to specify the subnets and security groups for an Amazon OpenSearch Service VPC endpoint. For more information, see <a href=\"https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html\">Launching your Amazon OpenSearch Service domains using a VPC</a>.")),
		mcp.WithObject("AutoTuneOptions", mcp.Description("Input parameter: Auto-Tune settings when updating a domain. For more information, see <a href=\"https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html\">Auto-Tune for Amazon OpenSearch Service</a>.")),
		mcp.WithObject("DomainEndpointOptions", mcp.Description("Input parameter: Options to configure a custom endpoint for an OpenSearch Service domain.")),
		mcp.WithObject("EncryptionAtRestOptions", mcp.Description("Input parameter: Specifies whether the domain should encrypt data at rest, and if so, the Key Management Service (KMS) key to use. Can be used only to create a new domain, not update an existing one.")),
		mcp.WithObject("NodeToNodeEncryptionOptions", mcp.Description("Input parameter: Enables or disables node-to-node encryption. For more information, see <a href=\"https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ntn.html\">Node-to-node encryption for Amazon OpenSearch Service</a>.")),
		mcp.WithString("AccessPolicies", mcp.Description("Input parameter: Access policy rules for an Amazon OpenSearch Service domain endpoint. For more information, see <a href=\"https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-access-policies\">Configuring access policies</a>. The maximum size of a policy document is 100 KB.")),
		mcp.WithObject("OffPeakWindowOptions", mcp.Description("Input parameter: Options for a domain's <a href=\"https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_OffPeakWindow.html\">off-peak window</a>, during which OpenSearch Service can perform mandatory configuration changes on the domain.")),
		mcp.WithObject("AdvancedSecurityOptions", mcp.Description("Input parameter: Options for enabling and configuring fine-grained access control. For more information, see <a href=\"https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html\">Fine-grained access control in Amazon OpenSearch Service</a>.")),
		mcp.WithObject("CognitoOptions", mcp.Description("Input parameter: Container for the parameters required to enable Cognito authentication for an OpenSearch Service domain. For more information, see <a href=\"https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html\">Configuring Amazon Cognito authentication for OpenSearch Dashboards</a>.")),
		mcp.WithObject("SnapshotOptions", mcp.Description("Input parameter: The time, in UTC format, when OpenSearch Service takes a daily automated snapshot of the specified domain. Default is <code>0</code> hours.")),
		mcp.WithObject("ClusterConfig", mcp.Description("Input parameter: Container for the cluster configuration of an OpenSearch Service domain. For more information, see <a href=\"https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html\">Creating and managing Amazon OpenSearch Service domains</a>.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdatedomainconfigHandler(cfg),
	}
}
