package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/aws-securityhub/mcp-server/config"
	"github.com/aws-securityhub/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreateautomationruleHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/automationrules/create", cfg.BaseURL)
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
		var result models.CreateAutomationRuleResponse
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

func CreateCreateautomationruleTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_automationrules_create",
		mcp.WithDescription(" Creates an automation rule based on input parameters. "),
		mcp.WithNumber("RuleOrder", mcp.Required(), mcp.Description("Input parameter: An integer ranging from 1 to 1000 that represents the order in which the rule action is applied to findings. Security Hub applies rules with lower values for this parameter first. ")),
		mcp.WithString("RuleStatus", mcp.Description("Input parameter:  Whether the rule is active after it is created. If this parameter is equal to <code>ENABLED</code>, Security Hub starts applying the rule to findings and finding updates after the rule is created. To change the value of this parameter after creating a rule, use <a href=\"https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_BatchUpdateAutomationRules.html\"> <code>BatchUpdateAutomationRules</code> </a>. ")),
		mcp.WithObject("Tags", mcp.Description("Input parameter:  User-defined tags that help you label the purpose of a rule. ")),
		mcp.WithArray("Actions", mcp.Required(), mcp.Description("Input parameter:  One or more actions to update finding fields if a finding matches the conditions specified in <code>Criteria</code>. ")),
		mcp.WithObject("Criteria", mcp.Required(), mcp.Description("Input parameter:  The criteria that determine which findings a rule applies to. ")),
		mcp.WithString("Description", mcp.Required(), mcp.Description("Input parameter:  A description of the rule. ")),
		mcp.WithBoolean("IsTerminal", mcp.Description("Input parameter: Specifies whether a rule is the last to be applied with respect to a finding that matches the rule criteria. This is useful when a finding matches the criteria for multiple rules, and each rule has different actions. If a rule is terminal, Security Hub applies the rule action to a finding that matches the rule criteria and doesn't evaluate other rules for the finding. By default, a rule isn't terminal. ")),
		mcp.WithString("RuleName", mcp.Required(), mcp.Description("Input parameter:  The name of the rule. ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateautomationruleHandler(cfg),
	}
}
