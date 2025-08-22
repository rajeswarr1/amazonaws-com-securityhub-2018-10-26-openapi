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

func EnablesecurityhubHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/accounts", cfg.BaseURL)
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
		var result models.EnableSecurityHubResponse
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

func CreateEnablesecurityhubTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_accounts",
		mcp.WithDescription("<p>Enables Security Hub for your account in the current Region or the Region you specify in the request.</p> <p>When you enable Security Hub, you grant to Security Hub the permissions necessary to gather findings from other services that are integrated with Security Hub.</p> <p>When you use the <code>EnableSecurityHub</code> operation to enable Security Hub, you also automatically enable the following standards:</p> <ul> <li> <p>Center for Internet Security (CIS) Amazon Web Services Foundations Benchmark v1.2.0</p> </li> <li> <p>Amazon Web Services Foundational Security Best Practices</p> </li> </ul> <p>Other standards are not automatically enabled. </p> <p>To opt out of automatically enabled standards, set <code>EnableDefaultStandards</code> to <code>false</code>.</p> <p>After you enable Security Hub, to enable a standard, use the <code>BatchEnableStandards</code> operation. To disable a standard, use the <code>BatchDisableStandards</code> operation.</p> <p>To learn more, see the <a href="https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-settingup.html">setup information</a> in the <i>Security Hub User Guide</i>.</p>"),
		mcp.WithString("ControlFindingGenerator", mcp.Description("Input parameter: <p>This field, used when enabling Security Hub, specifies whether the calling account has consolidated control findings turned on. If the value for this field is set to <code>SECURITY_CONTROL</code>, Security Hub generates a single finding for a control check even when the check applies to multiple enabled standards.</p> <p>If the value for this field is set to <code>STANDARD_CONTROL</code>, Security Hub generates separate findings for a control check when the check applies to multiple enabled standards.</p> <p>The value for this field in a member account matches the value in the administrator account. For accounts that aren't part of an organization, the default value of this field is <code>SECURITY_CONTROL</code> if you enabled Security Hub on or after February 23, 2023.</p>")),
		mcp.WithBoolean("EnableDefaultStandards", mcp.Description("Input parameter: Whether to enable the security standards that Security Hub has designated as automatically enabled. If you do not provide a value for <code>EnableDefaultStandards</code>, it is set to <code>true</code>. To not enable the automatically enabled standards, set <code>EnableDefaultStandards</code> to <code>false</code>.")),
		mcp.WithObject("Tags", mcp.Description("Input parameter: The tags to add to the hub resource when you enable Security Hub.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    EnablesecurityhubHandler(cfg),
	}
}
