package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aws-securityhub/mcp-server/config"
	"github.com/aws-securityhub/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DeletefindingaggregatorHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		FindingAggregatorArnVal, ok := args["FindingAggregatorArn"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: FindingAggregatorArn"), nil
		}
		FindingAggregatorArn, ok := FindingAggregatorArnVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: FindingAggregatorArn"), nil
		}
		url := fmt.Sprintf("%s/findingAggregator/delete/%s", cfg.BaseURL, FindingAggregatorArn)
		req, err := http.NewRequest("DELETE", url, nil)
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
		var result models.DeleteFindingAggregatorResponse
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

func CreateDeletefindingaggregatorTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_findingAggregator_delete_FindingAggregatorArn",
		mcp.WithDescription("<p>Deletes a finding aggregator. When you delete the finding aggregator, you stop finding aggregation.</p> <p>When you stop finding aggregation, findings that were already aggregated to the aggregation Region are still visible from the aggregation Region. New findings and finding updates are not aggregated. </p>"),
		mcp.WithString("FindingAggregatorArn", mcp.Required(), mcp.Description("The ARN of the finding aggregator to delete. To obtain the ARN, use <code>ListFindingAggregators</code>.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletefindingaggregatorHandler(cfg),
	}
}
