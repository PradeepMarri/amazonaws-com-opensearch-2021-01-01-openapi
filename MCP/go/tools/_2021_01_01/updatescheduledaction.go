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

func UpdatescheduledactionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/2021-01-01/opensearch/domain/%s/scheduledAction/update", cfg.BaseURL, DomainName)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
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
		var result models.UpdateScheduledActionResponse
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

func CreateUpdatescheduledactionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_2021-01-01_opensearch_domain_DomainName_scheduledAction_update",
		mcp.WithDescription("Reschedules a planned domain configuration change for a later time. This change can be a scheduled <a href="https://docs.aws.amazon.com/opensearch-service/latest/developerguide/service-software.html">service software update</a> or a <a href="https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html#auto-tune-types">blue/green Auto-Tune enhancement</a>."),
		mcp.WithString("DomainName", mcp.Required(), mcp.Description("The name of the domain to reschedule an action for.")),
		mcp.WithString("ActionType", mcp.Required(), mcp.Description("Input parameter: The type of action to reschedule. Can be one of <code>SERVICE_SOFTWARE_UPDATE</code>, <code>JVM_HEAP_SIZE_TUNING</code>, or <code>JVM_YOUNG_GEN_TUNING</code>. To retrieve this value, send a <a href=\"https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_ListScheduledActions.html\">ListScheduledActions</a> request.")),
		mcp.WithNumber("DesiredStartTime", mcp.Description("Input parameter: The time to implement the change, in Coordinated Universal Time (UTC). Only specify this parameter if you set <code>ScheduleAt</code> to <code>TIMESTAMP</code>.")),
		mcp.WithString("ScheduleAt", mcp.Required(), mcp.Description("Input parameter: <p>When to schedule the action.</p> <ul> <li> <p> <code>NOW</code> - Immediately schedules the update to happen in the current hour if there's capacity available.</p> </li> <li> <p> <code>TIMESTAMP</code> - Lets you specify a custom date and time to apply the update. If you specify this value, you must also provide a value for <code>DesiredStartTime</code>.</p> </li> <li> <p> <code>OFF_PEAK_WINDOW</code> - Marks the action to be picked up during an upcoming off-peak window. There's no guarantee that the change will be implemented during the next immediate window. Depending on capacity, it might happen in subsequent days.</p> </li> </ul>")),
		mcp.WithString("ActionID", mcp.Required(), mcp.Description("Input parameter: The unique identifier of the action to reschedule. To retrieve this ID, send a <a href=\"https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_ListScheduledActions.html\">ListScheduledActions</a> request.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdatescheduledactionHandler(cfg),
	}
}
