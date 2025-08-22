package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/aws-securityhub/mcp-server/config"
	"github.com/aws-securityhub/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetfindinghistoryHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
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
		url := fmt.Sprintf("%s/findingHistory/get%s", cfg.BaseURL, queryString)
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
		var result models.GetFindingHistoryResponse
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

func CreateGetfindinghistoryTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_findingHistory_get",
		mcp.WithDescription(" Returns history for a Security Hub finding in the last 90 days. The history includes changes made to any fields in the Amazon Web Services Security Finding Format (ASFF). "),
		mcp.WithString("MaxResults", mcp.Description("Pagination limit")),
		mcp.WithString("NextToken", mcp.Description("Pagination token")),
		mcp.WithObject("FindingIdentifier", mcp.Required(), mcp.Description("Input parameter: Identifies which finding to get the finding history for.")),
		mcp.WithNumber("MaxResults", mcp.Description("Input parameter:  The maximum number of results to be returned. If you don’t provide it, Security Hub returns up to 100 results of finding history. ")),
		mcp.WithString("NextToken", mcp.Description("Input parameter:  A token for pagination purposes. Provide <code>NULL</code> as the initial value. In subsequent requests, provide the token included in the response to get up to an additional 100 results of finding history. If you don’t provide <code>NextToken</code>, Security Hub returns up to 100 results of finding history for each request. ")),
		mcp.WithString("StartTime", mcp.Description("Input parameter: <p> An ISO 8601-formatted timestamp that indicates the start time of the requested finding history. A correctly formatted example is <code>2020-05-21T20:16:34.724Z</code>. The value cannot contain spaces, and date and time should be separated by <code>T</code>. For more information, see <a href=\"https://www.rfc-editor.org/rfc/rfc3339#section-5.6\">RFC 3339 section 5.6, Internet Date/Time Format</a>.</p> <p>If you provide values for both <code>StartTime</code> and <code>EndTime</code>, Security Hub returns finding history for the specified time period. If you provide a value for <code>StartTime</code> but not for <code>EndTime</code>, Security Hub returns finding history from the <code>StartTime</code> to the time at which the API is called. If you provide a value for <code>EndTime</code> but not for <code>StartTime</code>, Security Hub returns finding history from the <a href=\"https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_AwsSecurityFindingFilters.html#securityhub-Type-AwsSecurityFindingFilters-CreatedAt\">CreatedAt</a> timestamp of the finding to the <code>EndTime</code>. If you provide neither <code>StartTime</code> nor <code>EndTime</code>, Security Hub returns finding history from the CreatedAt timestamp of the finding to the time at which the API is called. In all of these scenarios, the response is limited to 100 results, and the maximum time period is limited to 90 days. </p>")),
		mcp.WithString("EndTime", mcp.Description("Input parameter: <p> An ISO 8601-formatted timestamp that indicates the end time of the requested finding history. A correctly formatted example is <code>2020-05-21T20:16:34.724Z</code>. The value cannot contain spaces, and date and time should be separated by <code>T</code>. For more information, see <a href=\"https://www.rfc-editor.org/rfc/rfc3339#section-5.6\">RFC 3339 section 5.6, Internet Date/Time Format</a>.</p> <p>If you provide values for both <code>StartTime</code> and <code>EndTime</code>, Security Hub returns finding history for the specified time period. If you provide a value for <code>StartTime</code> but not for <code>EndTime</code>, Security Hub returns finding history from the <code>StartTime</code> to the time at which the API is called. If you provide a value for <code>EndTime</code> but not for <code>StartTime</code>, Security Hub returns finding history from the <a href=\"https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_AwsSecurityFindingFilters.html#securityhub-Type-AwsSecurityFindingFilters-CreatedAt\">CreatedAt</a> timestamp of the finding to the <code>EndTime</code>. If you provide neither <code>StartTime</code> nor <code>EndTime</code>, Security Hub returns finding history from the CreatedAt timestamp of the finding to the time at which the API is called. In all of these scenarios, the response is limited to 100 results, and the maximum time period is limited to 90 days.</p>")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetfindinghistoryHandler(cfg),
	}
}
