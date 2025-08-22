package main

import (
	"github.com/aws-securityhub/mcp-server/config"
	"github.com/aws-securityhub/mcp-server/models"
	tools_master "github.com/aws-securityhub/mcp-server/tools/master"
	tools_members "github.com/aws-securityhub/mcp-server/tools/members"
	tools_administrator "github.com/aws-securityhub/mcp-server/tools/administrator"
	tools_standards "github.com/aws-securityhub/mcp-server/tools/standards"
	tools_organization "github.com/aws-securityhub/mcp-server/tools/organization"
	tools_findings "github.com/aws-securityhub/mcp-server/tools/findings"
	tools_automationrules "github.com/aws-securityhub/mcp-server/tools/automationrules"
	tools_actiontargets "github.com/aws-securityhub/mcp-server/tools/actiontargets"
	tools_findingaggregator "github.com/aws-securityhub/mcp-server/tools/findingaggregator"
	tools_accounts "github.com/aws-securityhub/mcp-server/tools/accounts"
	tools_associations "github.com/aws-securityhub/mcp-server/tools/associations"
	tools_productsubscriptions "github.com/aws-securityhub/mcp-server/tools/productsubscriptions"
	tools_securitycontrols "github.com/aws-securityhub/mcp-server/tools/securitycontrols"
	tools_invitations "github.com/aws-securityhub/mcp-server/tools/invitations"
	tools_tags "github.com/aws-securityhub/mcp-server/tools/tags"
	tools_insights "github.com/aws-securityhub/mcp-server/tools/insights"
	tools_associations_securitycontrolid "github.com/aws-securityhub/mcp-server/tools/associations_securitycontrolid"
	tools_products "github.com/aws-securityhub/mcp-server/tools/products"
	tools_findinghistory "github.com/aws-securityhub/mcp-server/tools/findinghistory"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_master.CreateDisassociatefrommasteraccountTool(cfg),
		tools_members.CreateDisassociatemembersTool(cfg),
		tools_administrator.CreateDisassociatefromadministratoraccountTool(cfg),
		tools_standards.CreateBatchdisablestandardsTool(cfg),
		tools_organization.CreateDescribeorganizationconfigurationTool(cfg),
		tools_organization.CreateUpdateorganizationconfigurationTool(cfg),
		tools_standards.CreateBatchenablestandardsTool(cfg),
		tools_findings.CreateBatchimportfindingsTool(cfg),
		tools_members.CreateGetmembersTool(cfg),
		tools_automationrules.CreateBatchupdateautomationrulesTool(cfg),
		tools_actiontargets.CreateDescribeactiontargetsTool(cfg),
		tools_organization.CreateDisableorganizationadminaccountTool(cfg),
		tools_standards.CreateGetenabledstandardsTool(cfg),
		tools_master.CreateGetmasteraccountTool(cfg),
		tools_master.CreateAcceptinvitationTool(cfg),
		tools_findingaggregator.CreateListfindingaggregatorsTool(cfg),
		tools_accounts.CreateEnablesecurityhubTool(cfg),
		tools_accounts.CreateDisablesecurityhubTool(cfg),
		tools_accounts.CreateDescribehubTool(cfg),
		tools_accounts.CreateUpdatesecurityhubconfigurationTool(cfg),
		tools_associations.CreateBatchgetstandardscontrolassociationsTool(cfg),
		tools_findings.CreateBatchupdatefindingsTool(cfg),
		tools_productsubscriptions.CreateListenabledproductsforimportTool(cfg),
		tools_productsubscriptions.CreateEnableimportfindingsforproductTool(cfg),
		tools_securitycontrols.CreateListsecuritycontroldefinitionsTool(cfg),
		tools_invitations.CreateListinvitationsTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_insights.CreateDeleteinsightTool(cfg),
		tools_insights.CreateUpdateinsightTool(cfg),
		tools_invitations.CreateDeclineinvitationsTool(cfg),
		tools_organization.CreateListorganizationadminaccountsTool(cfg),
		tools_actiontargets.CreateCreateactiontargetTool(cfg),
		tools_findings.CreateUpdatefindingsTool(cfg),
		tools_findings.CreateGetfindingsTool(cfg),
		tools_findingaggregator.CreateCreatefindingaggregatorTool(cfg),
		tools_findingaggregator.CreateGetfindingaggregatorTool(cfg),
		tools_associations.CreateBatchupdatestandardscontrolassociationsTool(cfg),
		tools_findingaggregator.CreateDeletefindingaggregatorTool(cfg),
		tools_administrator.CreateAcceptadministratorinvitationTool(cfg),
		tools_administrator.CreateGetadministratoraccountTool(cfg),
		tools_standards.CreateDescribestandardscontrolsTool(cfg),
		tools_organization.CreateEnableorganizationadminaccountTool(cfg),
		tools_insights.CreateGetinsightsTool(cfg),
		tools_members.CreateDeletemembersTool(cfg),
		tools_standards.CreateDescribestandardsTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_associations_securitycontrolid.CreateListstandardscontrolassociationsTool(cfg),
		tools_securitycontrols.CreateBatchgetsecuritycontrolsTool(cfg),
		tools_actiontargets.CreateDeleteactiontargetTool(cfg),
		tools_actiontargets.CreateUpdateactiontargetTool(cfg),
		tools_automationrules.CreateCreateautomationruleTool(cfg),
		tools_findingaggregator.CreateUpdatefindingaggregatorTool(cfg),
		tools_standards.CreateUpdatestandardscontrolTool(cfg),
		tools_products.CreateDescribeproductsTool(cfg),
		tools_members.CreateListmembersTool(cfg),
		tools_members.CreateCreatemembersTool(cfg),
		tools_productsubscriptions.CreateDisableimportfindingsforproductTool(cfg),
		tools_automationrules.CreateBatchdeleteautomationrulesTool(cfg),
		tools_invitations.CreateDeleteinvitationsTool(cfg),
		tools_findinghistory.CreateGetfindinghistoryTool(cfg),
		tools_automationrules.CreateListautomationrulesTool(cfg),
		tools_invitations.CreateGetinvitationscountTool(cfg),
		tools_automationrules.CreateBatchgetautomationrulesTool(cfg),
		tools_insights.CreateCreateinsightTool(cfg),
		tools_insights.CreateGetinsightresultsTool(cfg),
		tools_members.CreateInvitemembersTool(cfg),
	}
}
