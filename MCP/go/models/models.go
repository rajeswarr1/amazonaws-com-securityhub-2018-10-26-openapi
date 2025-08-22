package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// AwsEcsClusterConfigurationExecuteCommandConfigurationLogConfigurationDetails represents the AwsEcsClusterConfigurationExecuteCommandConfigurationLogConfigurationDetails schema from the OpenAPI specification
type AwsEcsClusterConfigurationExecuteCommandConfigurationLogConfigurationDetails struct {
	Cloudwatchencryptionenabled interface{} `json:"CloudWatchEncryptionEnabled,omitempty"`
	Cloudwatchloggroupname interface{} `json:"CloudWatchLogGroupName,omitempty"`
	S3bucketname interface{} `json:"S3BucketName,omitempty"`
	S3encryptionenabled interface{} `json:"S3EncryptionEnabled,omitempty"`
	S3keyprefix interface{} `json:"S3KeyPrefix,omitempty"`
}

// AwsWafRegionalWebAclRulesListActionDetails represents the AwsWafRegionalWebAclRulesListActionDetails schema from the OpenAPI specification
type AwsWafRegionalWebAclRulesListActionDetails struct {
	TypeField interface{} `json:"Type,omitempty"`
}

// AwsWafRulePredicateListDetails represents the AwsWafRulePredicateListDetails schema from the OpenAPI specification
type AwsWafRulePredicateListDetails struct {
	Dataid interface{} `json:"DataId,omitempty"`
	Negated interface{} `json:"Negated,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// AwsS3BucketLoggingConfiguration represents the AwsS3BucketLoggingConfiguration schema from the OpenAPI specification
type AwsS3BucketLoggingConfiguration struct {
	Destinationbucketname interface{} `json:"DestinationBucketName,omitempty"`
	Logfileprefix interface{} `json:"LogFilePrefix,omitempty"`
}

// AwsEcsTaskVolumeDetails represents the AwsEcsTaskVolumeDetails schema from the OpenAPI specification
type AwsEcsTaskVolumeDetails struct {
	Host interface{} `json:"Host,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// AwsEc2LaunchTemplateDataCreditSpecificationDetails represents the AwsEc2LaunchTemplateDataCreditSpecificationDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataCreditSpecificationDetails struct {
	Cpucredits interface{} `json:"CpuCredits,omitempty"`
}

// AwsRdsDbSubnetGroup represents the AwsRdsDbSubnetGroup schema from the OpenAPI specification
type AwsRdsDbSubnetGroup struct {
	Dbsubnetgrouparn interface{} `json:"DbSubnetGroupArn,omitempty"`
	Dbsubnetgroupdescription interface{} `json:"DbSubnetGroupDescription,omitempty"`
	Dbsubnetgroupname interface{} `json:"DbSubnetGroupName,omitempty"`
	Subnetgroupstatus interface{} `json:"SubnetGroupStatus,omitempty"`
	Subnets interface{} `json:"Subnets,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
}

// UpdateFindingAggregatorRequest represents the UpdateFindingAggregatorRequest schema from the OpenAPI specification
type UpdateFindingAggregatorRequest struct {
	Findingaggregatorarn interface{} `json:"FindingAggregatorArn"`
	Regionlinkingmode interface{} `json:"RegionLinkingMode"`
	Regions interface{} `json:"Regions,omitempty"`
}

// AwsStepFunctionStateMachineLoggingConfigurationDestinationsCloudWatchLogsLogGroupDetails represents the AwsStepFunctionStateMachineLoggingConfigurationDestinationsCloudWatchLogsLogGroupDetails schema from the OpenAPI specification
type AwsStepFunctionStateMachineLoggingConfigurationDestinationsCloudWatchLogsLogGroupDetails struct {
	Loggrouparn interface{} `json:"LogGroupArn,omitempty"`
}

// DeleteActionTargetResponse represents the DeleteActionTargetResponse schema from the OpenAPI specification
type DeleteActionTargetResponse struct {
	Actiontargetarn interface{} `json:"ActionTargetArn"`
}

// NumberFilter represents the NumberFilter schema from the OpenAPI specification
type NumberFilter struct {
	Lte interface{} `json:"Lte,omitempty"`
	Eq interface{} `json:"Eq,omitempty"`
	Gte interface{} `json:"Gte,omitempty"`
}

// AwsApiGatewayRestApiDetails represents the AwsApiGatewayRestApiDetails schema from the OpenAPI specification
type AwsApiGatewayRestApiDetails struct {
	Apikeysource interface{} `json:"ApiKeySource,omitempty"`
	Binarymediatypes interface{} `json:"BinaryMediaTypes,omitempty"`
	Minimumcompressionsize interface{} `json:"MinimumCompressionSize,omitempty"`
	Version interface{} `json:"Version,omitempty"`
	Endpointconfiguration interface{} `json:"EndpointConfiguration,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// ClassificationResult represents the ClassificationResult schema from the OpenAPI specification
type ClassificationResult struct {
	Sensitivedata interface{} `json:"SensitiveData,omitempty"`
	Sizeclassified interface{} `json:"SizeClassified,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Additionaloccurrences interface{} `json:"AdditionalOccurrences,omitempty"`
	Customdataidentifiers interface{} `json:"CustomDataIdentifiers,omitempty"`
	Mimetype interface{} `json:"MimeType,omitempty"`
}

// AwsEc2TransitGatewayDetails represents the AwsEc2TransitGatewayDetails schema from the OpenAPI specification
type AwsEc2TransitGatewayDetails struct {
	Propagationdefaultroutetableid interface{} `json:"PropagationDefaultRouteTableId,omitempty"`
	Transitgatewaycidrblocks interface{} `json:"TransitGatewayCidrBlocks,omitempty"`
	Amazonsideasn interface{} `json:"AmazonSideAsn,omitempty"`
	Autoacceptsharedattachments interface{} `json:"AutoAcceptSharedAttachments,omitempty"`
	Associationdefaultroutetableid interface{} `json:"AssociationDefaultRouteTableId,omitempty"`
	Defaultroutetablepropagation interface{} `json:"DefaultRouteTablePropagation,omitempty"`
	Dnssupport interface{} `json:"DnsSupport,omitempty"`
	Multicastsupport interface{} `json:"MulticastSupport,omitempty"`
	Defaultroutetableassociation interface{} `json:"DefaultRouteTableAssociation,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Vpnecmpsupport interface{} `json:"VpnEcmpSupport,omitempty"`
	Id interface{} `json:"Id,omitempty"`
}

// AwsWafRuleGroupRulesDetails represents the AwsWafRuleGroupRulesDetails schema from the OpenAPI specification
type AwsWafRuleGroupRulesDetails struct {
	TypeField interface{} `json:"Type,omitempty"`
	Action interface{} `json:"Action,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
	Ruleid interface{} `json:"RuleId,omitempty"`
}

// AwsCloudFrontDistributionOrigins represents the AwsCloudFrontDistributionOrigins schema from the OpenAPI specification
type AwsCloudFrontDistributionOrigins struct {
	Items interface{} `json:"Items,omitempty"`
}

// AwsOpenSearchServiceDomainLogPublishingOption represents the AwsOpenSearchServiceDomainLogPublishingOption schema from the OpenAPI specification
type AwsOpenSearchServiceDomainLogPublishingOption struct {
	Cloudwatchlogsloggrouparn interface{} `json:"CloudWatchLogsLogGroupArn,omitempty"`
	Enabled interface{} `json:"Enabled,omitempty"`
}

// AwsGuardDutyDetectorDataSourcesDnsLogsDetails represents the AwsGuardDutyDetectorDataSourcesDnsLogsDetails schema from the OpenAPI specification
type AwsGuardDutyDetectorDataSourcesDnsLogsDetails struct {
	Status interface{} `json:"Status,omitempty"`
}

// ListStandardsControlAssociationsRequest represents the ListStandardsControlAssociationsRequest schema from the OpenAPI specification
type ListStandardsControlAssociationsRequest struct {
}

// UpdateStandardsControlResponse represents the UpdateStandardsControlResponse schema from the OpenAPI specification
type UpdateStandardsControlResponse struct {
}

// AwsCertificateManagerCertificateKeyUsage represents the AwsCertificateManagerCertificateKeyUsage schema from the OpenAPI specification
type AwsCertificateManagerCertificateKeyUsage struct {
	Name interface{} `json:"Name,omitempty"`
}

// AwsSnsTopicDetails represents the AwsSnsTopicDetails schema from the OpenAPI specification
type AwsSnsTopicDetails struct {
	Firehosefailurefeedbackrolearn interface{} `json:"FirehoseFailureFeedbackRoleArn,omitempty"`
	Firehosesuccessfeedbackrolearn interface{} `json:"FirehoseSuccessFeedbackRoleArn,omitempty"`
	Kmsmasterkeyid interface{} `json:"KmsMasterKeyId,omitempty"`
	Owner interface{} `json:"Owner,omitempty"`
	Sqssuccessfeedbackrolearn interface{} `json:"SqsSuccessFeedbackRoleArn,omitempty"`
	Subscription interface{} `json:"Subscription,omitempty"`
	Httpfailurefeedbackrolearn interface{} `json:"HttpFailureFeedbackRoleArn,omitempty"`
	Httpsuccessfeedbackrolearn interface{} `json:"HttpSuccessFeedbackRoleArn,omitempty"`
	Sqsfailurefeedbackrolearn interface{} `json:"SqsFailureFeedbackRoleArn,omitempty"`
	Topicname interface{} `json:"TopicName,omitempty"`
	Applicationsuccessfeedbackrolearn interface{} `json:"ApplicationSuccessFeedbackRoleArn,omitempty"`
}

// AwsEc2VpcPeeringConnectionStatusDetails represents the AwsEc2VpcPeeringConnectionStatusDetails schema from the OpenAPI specification
type AwsEc2VpcPeeringConnectionStatusDetails struct {
	Code interface{} `json:"Code,omitempty"`
	Message interface{} `json:"Message,omitempty"`
}

// AwsNetworkFirewallFirewallSubnetMappingsDetails represents the AwsNetworkFirewallFirewallSubnetMappingsDetails schema from the OpenAPI specification
type AwsNetworkFirewallFirewallSubnetMappingsDetails struct {
	Subnetid interface{} `json:"SubnetId,omitempty"`
}

// AwsAutoScalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification represents the AwsAutoScalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification schema from the OpenAPI specification
type AwsAutoScalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification struct {
	Launchtemplateid interface{} `json:"LaunchTemplateId,omitempty"`
	Launchtemplatename interface{} `json:"LaunchTemplateName,omitempty"`
	Version interface{} `json:"Version,omitempty"`
}

// AwsRedshiftClusterEndpoint represents the AwsRedshiftClusterEndpoint schema from the OpenAPI specification
type AwsRedshiftClusterEndpoint struct {
	Port interface{} `json:"Port,omitempty"`
	Address interface{} `json:"Address,omitempty"`
}

// AwsOpenSearchServiceDomainServiceSoftwareOptionsDetails represents the AwsOpenSearchServiceDomainServiceSoftwareOptionsDetails schema from the OpenAPI specification
type AwsOpenSearchServiceDomainServiceSoftwareOptionsDetails struct {
	Automatedupdatedate interface{} `json:"AutomatedUpdateDate,omitempty"`
	Cancellable interface{} `json:"Cancellable,omitempty"`
	Currentversion interface{} `json:"CurrentVersion,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Newversion interface{} `json:"NewVersion,omitempty"`
	Optionaldeployment interface{} `json:"OptionalDeployment,omitempty"`
	Updateavailable interface{} `json:"UpdateAvailable,omitempty"`
	Updatestatus interface{} `json:"UpdateStatus,omitempty"`
}

// AwsEcsServiceCapacityProviderStrategyDetails represents the AwsEcsServiceCapacityProviderStrategyDetails schema from the OpenAPI specification
type AwsEcsServiceCapacityProviderStrategyDetails struct {
	Base interface{} `json:"Base,omitempty"`
	Capacityprovider interface{} `json:"CapacityProvider,omitempty"`
	Weight interface{} `json:"Weight,omitempty"`
}

// AwsWafv2RulesActionCaptchaDetails represents the AwsWafv2RulesActionCaptchaDetails schema from the OpenAPI specification
type AwsWafv2RulesActionCaptchaDetails struct {
	Customrequesthandling interface{} `json:"CustomRequestHandling,omitempty"`
}

// CustomDataIdentifiersDetections represents the CustomDataIdentifiersDetections schema from the OpenAPI specification
type CustomDataIdentifiersDetections struct {
	Arn interface{} `json:"Arn,omitempty"`
	Count interface{} `json:"Count,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Occurrences interface{} `json:"Occurrences,omitempty"`
}

// AwsRedshiftClusterElasticIpStatus represents the AwsRedshiftClusterElasticIpStatus schema from the OpenAPI specification
type AwsRedshiftClusterElasticIpStatus struct {
	Status interface{} `json:"Status,omitempty"`
	Elasticip interface{} `json:"ElasticIp,omitempty"`
}

// RuleGroupSourceStatelessRuleMatchAttributesSourcePorts represents the RuleGroupSourceStatelessRuleMatchAttributesSourcePorts schema from the OpenAPI specification
type RuleGroupSourceStatelessRuleMatchAttributesSourcePorts struct {
	Toport interface{} `json:"ToPort,omitempty"`
	Fromport interface{} `json:"FromPort,omitempty"`
}

// ActionLocalPortDetails represents the ActionLocalPortDetails schema from the OpenAPI specification
type ActionLocalPortDetails struct {
	Portname interface{} `json:"PortName,omitempty"`
	Port interface{} `json:"Port,omitempty"`
}

// AwsElbLoadBalancerConnectionDraining represents the AwsElbLoadBalancerConnectionDraining schema from the OpenAPI specification
type AwsElbLoadBalancerConnectionDraining struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Timeout interface{} `json:"Timeout,omitempty"`
}

// Standard represents the Standard schema from the OpenAPI specification
type Standard struct {
	Standardsmanagedby interface{} `json:"StandardsManagedBy,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Enabledbydefault interface{} `json:"EnabledByDefault,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Standardsarn interface{} `json:"StandardsArn,omitempty"`
}

// AwsBackupRecoveryPointCreatedByDetails represents the AwsBackupRecoveryPointCreatedByDetails schema from the OpenAPI specification
type AwsBackupRecoveryPointCreatedByDetails struct {
	Backupplanarn interface{} `json:"BackupPlanArn,omitempty"`
	Backupplanid interface{} `json:"BackupPlanId,omitempty"`
	Backupplanversion interface{} `json:"BackupPlanVersion,omitempty"`
	Backupruleid interface{} `json:"BackupRuleId,omitempty"`
}

// UpdateAutomationRulesRequestItem represents the UpdateAutomationRulesRequestItem schema from the OpenAPI specification
type UpdateAutomationRulesRequestItem struct {
	Actions interface{} `json:"Actions,omitempty"`
	Criteria interface{} `json:"Criteria,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Isterminal interface{} `json:"IsTerminal,omitempty"`
	Rulearn interface{} `json:"RuleArn"`
	Rulename interface{} `json:"RuleName,omitempty"`
	Ruleorder interface{} `json:"RuleOrder,omitempty"`
	Rulestatus interface{} `json:"RuleStatus,omitempty"`
}

// AwsEcsContainerDetails represents the AwsEcsContainerDetails schema from the OpenAPI specification
type AwsEcsContainerDetails struct {
	Image interface{} `json:"Image,omitempty"`
	Mountpoints interface{} `json:"MountPoints,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Privileged interface{} `json:"Privileged,omitempty"`
}

// AwsOpenSearchServiceDomainNodeToNodeEncryptionOptionsDetails represents the AwsOpenSearchServiceDomainNodeToNodeEncryptionOptionsDetails schema from the OpenAPI specification
type AwsOpenSearchServiceDomainNodeToNodeEncryptionOptionsDetails struct {
	Enabled interface{} `json:"Enabled,omitempty"`
}

// AwsEc2LaunchTemplateDataInstanceRequirementsAcceleratorCountDetails represents the AwsEc2LaunchTemplateDataInstanceRequirementsAcceleratorCountDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataInstanceRequirementsAcceleratorCountDetails struct {
	Max interface{} `json:"Max,omitempty"`
	Min interface{} `json:"Min,omitempty"`
}

// AwsOpenSearchServiceDomainEncryptionAtRestOptionsDetails represents the AwsOpenSearchServiceDomainEncryptionAtRestOptionsDetails schema from the OpenAPI specification
type AwsOpenSearchServiceDomainEncryptionAtRestOptionsDetails struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
}

// AwsDynamoDbTableDetails represents the AwsDynamoDbTableDetails schema from the OpenAPI specification
type AwsDynamoDbTableDetails struct {
	Streamspecification interface{} `json:"StreamSpecification,omitempty"`
	Ssedescription interface{} `json:"SseDescription,omitempty"`
	Tablename interface{} `json:"TableName,omitempty"`
	Tablesizebytes interface{} `json:"TableSizeBytes,omitempty"`
	Attributedefinitions interface{} `json:"AttributeDefinitions,omitempty"`
	Replicas interface{} `json:"Replicas,omitempty"`
	Tablestatus interface{} `json:"TableStatus,omitempty"`
	Billingmodesummary interface{} `json:"BillingModeSummary,omitempty"`
	Tableid interface{} `json:"TableId,omitempty"`
	Lateststreamarn interface{} `json:"LatestStreamArn,omitempty"`
	Restoresummary interface{} `json:"RestoreSummary,omitempty"`
	Itemcount interface{} `json:"ItemCount,omitempty"`
	Globalsecondaryindexes interface{} `json:"GlobalSecondaryIndexes,omitempty"`
	Provisionedthroughput interface{} `json:"ProvisionedThroughput,omitempty"`
	Lateststreamlabel interface{} `json:"LatestStreamLabel,omitempty"`
	Globaltableversion interface{} `json:"GlobalTableVersion,omitempty"`
	Keyschema interface{} `json:"KeySchema,omitempty"`
	Localsecondaryindexes interface{} `json:"LocalSecondaryIndexes,omitempty"`
	Creationdatetime interface{} `json:"CreationDateTime,omitempty"`
}

// CreateAutomationRuleRequest represents the CreateAutomationRuleRequest schema from the OpenAPI specification
type CreateAutomationRuleRequest struct {
	Ruleorder interface{} `json:"RuleOrder"`
	Rulestatus interface{} `json:"RuleStatus,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Actions interface{} `json:"Actions"`
	Criteria interface{} `json:"Criteria"`
	Description interface{} `json:"Description"`
	Isterminal interface{} `json:"IsTerminal,omitempty"`
	Rulename interface{} `json:"RuleName"`
}

// AwsRedshiftClusterRestoreStatus represents the AwsRedshiftClusterRestoreStatus schema from the OpenAPI specification
type AwsRedshiftClusterRestoreStatus struct {
	Elapsedtimeinseconds interface{} `json:"ElapsedTimeInSeconds,omitempty"`
	Estimatedtimetocompletioninseconds interface{} `json:"EstimatedTimeToCompletionInSeconds,omitempty"`
	Progressinmegabytes interface{} `json:"ProgressInMegaBytes,omitempty"`
	Snapshotsizeinmegabytes interface{} `json:"SnapshotSizeInMegaBytes,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Currentrestorerateinmegabytespersecond interface{} `json:"CurrentRestoreRateInMegaBytesPerSecond,omitempty"`
}

// DisableOrganizationAdminAccountRequest represents the DisableOrganizationAdminAccountRequest schema from the OpenAPI specification
type DisableOrganizationAdminAccountRequest struct {
	Adminaccountid interface{} `json:"AdminAccountId"`
}

// Network represents the Network schema from the OpenAPI specification
type Network struct {
	Sourcemac interface{} `json:"SourceMac,omitempty"`
	Destinationipv4 interface{} `json:"DestinationIpV4,omitempty"`
	Destinationipv6 interface{} `json:"DestinationIpV6,omitempty"`
	Openportrange interface{} `json:"OpenPortRange,omitempty"`
	Protocol interface{} `json:"Protocol,omitempty"`
	Sourcedomain interface{} `json:"SourceDomain,omitempty"`
	Destinationdomain interface{} `json:"DestinationDomain,omitempty"`
	Sourceport interface{} `json:"SourcePort,omitempty"`
	Destinationport interface{} `json:"DestinationPort,omitempty"`
	Direction interface{} `json:"Direction,omitempty"`
	Sourceipv4 interface{} `json:"SourceIpV4,omitempty"`
	Sourceipv6 interface{} `json:"SourceIpV6,omitempty"`
}

// AwsAutoScalingLaunchConfigurationMetadataOptions represents the AwsAutoScalingLaunchConfigurationMetadataOptions schema from the OpenAPI specification
type AwsAutoScalingLaunchConfigurationMetadataOptions struct {
	Httpendpoint interface{} `json:"HttpEndpoint,omitempty"`
	Httpputresponsehoplimit interface{} `json:"HttpPutResponseHopLimit,omitempty"`
	Httptokens interface{} `json:"HttpTokens,omitempty"`
}

// AwsOpenSearchServiceDomainDetails represents the AwsOpenSearchServiceDomainDetails schema from the OpenAPI specification
type AwsOpenSearchServiceDomainDetails struct {
	Logpublishingoptions interface{} `json:"LogPublishingOptions,omitempty"`
	Nodetonodeencryptionoptions interface{} `json:"NodeToNodeEncryptionOptions,omitempty"`
	Encryptionatrestoptions interface{} `json:"EncryptionAtRestOptions,omitempty"`
	Domainendpoints interface{} `json:"DomainEndpoints,omitempty"`
	Engineversion interface{} `json:"EngineVersion,omitempty"`
	Advancedsecurityoptions interface{} `json:"AdvancedSecurityOptions,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Domainendpoint interface{} `json:"DomainEndpoint,omitempty"`
	Vpcoptions interface{} `json:"VpcOptions,omitempty"`
	Clusterconfig interface{} `json:"ClusterConfig,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Domainendpointoptions interface{} `json:"DomainEndpointOptions,omitempty"`
	Servicesoftwareoptions interface{} `json:"ServiceSoftwareOptions,omitempty"`
	Accesspolicies interface{} `json:"AccessPolicies,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
}

// ListEnabledProductsForImportResponse represents the ListEnabledProductsForImportResponse schema from the OpenAPI specification
type ListEnabledProductsForImportResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Productsubscriptions interface{} `json:"ProductSubscriptions,omitempty"`
}

// AwsEcsServiceDeploymentConfigurationDeploymentCircuitBreakerDetails represents the AwsEcsServiceDeploymentConfigurationDeploymentCircuitBreakerDetails schema from the OpenAPI specification
type AwsEcsServiceDeploymentConfigurationDeploymentCircuitBreakerDetails struct {
	Rollback interface{} `json:"Rollback,omitempty"`
	Enable interface{} `json:"Enable,omitempty"`
}

// AssociatedStandard represents the AssociatedStandard schema from the OpenAPI specification
type AssociatedStandard struct {
	Standardsid interface{} `json:"StandardsId,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsMountPointsDetails represents the AwsEcsTaskDefinitionContainerDefinitionsMountPointsDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsMountPointsDetails struct {
	Containerpath interface{} `json:"ContainerPath,omitempty"`
	Readonly interface{} `json:"ReadOnly,omitempty"`
	Sourcevolume interface{} `json:"SourceVolume,omitempty"`
}

// AwsCodeBuildProjectLogsConfigDetails represents the AwsCodeBuildProjectLogsConfigDetails schema from the OpenAPI specification
type AwsCodeBuildProjectLogsConfigDetails struct {
	Cloudwatchlogs interface{} `json:"CloudWatchLogs,omitempty"`
	S3logs interface{} `json:"S3Logs,omitempty"`
}

// AwsElasticsearchDomainElasticsearchClusterConfigDetails represents the AwsElasticsearchDomainElasticsearchClusterConfigDetails schema from the OpenAPI specification
type AwsElasticsearchDomainElasticsearchClusterConfigDetails struct {
	Zoneawarenessenabled interface{} `json:"ZoneAwarenessEnabled,omitempty"`
	Dedicatedmastercount interface{} `json:"DedicatedMasterCount,omitempty"`
	Dedicatedmasterenabled interface{} `json:"DedicatedMasterEnabled,omitempty"`
	Dedicatedmastertype interface{} `json:"DedicatedMasterType,omitempty"`
	Instancecount interface{} `json:"InstanceCount,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Zoneawarenessconfig interface{} `json:"ZoneAwarenessConfig,omitempty"`
}

// AwsElbLoadBalancerHealthCheck represents the AwsElbLoadBalancerHealthCheck schema from the OpenAPI specification
type AwsElbLoadBalancerHealthCheck struct {
	Healthythreshold interface{} `json:"HealthyThreshold,omitempty"`
	Interval interface{} `json:"Interval,omitempty"`
	Target interface{} `json:"Target,omitempty"`
	Timeout interface{} `json:"Timeout,omitempty"`
	Unhealthythreshold interface{} `json:"UnhealthyThreshold,omitempty"`
}

// AwsRedshiftClusterPendingModifiedValues represents the AwsRedshiftClusterPendingModifiedValues schema from the OpenAPI specification
type AwsRedshiftClusterPendingModifiedValues struct {
	Encryptiontype interface{} `json:"EncryptionType,omitempty"`
	Clusteridentifier interface{} `json:"ClusterIdentifier,omitempty"`
	Clustertype interface{} `json:"ClusterType,omitempty"`
	Enhancedvpcrouting interface{} `json:"EnhancedVpcRouting,omitempty"`
	Maintenancetrackname interface{} `json:"MaintenanceTrackName,omitempty"`
	Nodetype interface{} `json:"NodeType,omitempty"`
	Numberofnodes interface{} `json:"NumberOfNodes,omitempty"`
	Masteruserpassword interface{} `json:"MasterUserPassword,omitempty"`
	Publiclyaccessible interface{} `json:"PubliclyAccessible,omitempty"`
	Automatedsnapshotretentionperiod interface{} `json:"AutomatedSnapshotRetentionPeriod,omitempty"`
	Clusterversion interface{} `json:"ClusterVersion,omitempty"`
}

// AwsRdsDbClusterSnapshotDbClusterSnapshotAttribute represents the AwsRdsDbClusterSnapshotDbClusterSnapshotAttribute schema from the OpenAPI specification
type AwsRdsDbClusterSnapshotDbClusterSnapshotAttribute struct {
	Attributename interface{} `json:"AttributeName,omitempty"`
	Attributevalues interface{} `json:"AttributeValues,omitempty"`
}

// AwsElbLoadBalancerSourceSecurityGroup represents the AwsElbLoadBalancerSourceSecurityGroup schema from the OpenAPI specification
type AwsElbLoadBalancerSourceSecurityGroup struct {
	Owneralias interface{} `json:"OwnerAlias,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
}

// AwsBackupBackupPlanAdvancedBackupSettingsDetails represents the AwsBackupBackupPlanAdvancedBackupSettingsDetails schema from the OpenAPI specification
type AwsBackupBackupPlanAdvancedBackupSettingsDetails struct {
	Backupoptions interface{} `json:"BackupOptions,omitempty"`
	Resourcetype interface{} `json:"ResourceType,omitempty"`
}

// StandardsControlAssociationId represents the StandardsControlAssociationId schema from the OpenAPI specification
type StandardsControlAssociationId struct {
	Securitycontrolid interface{} `json:"SecurityControlId"`
	Standardsarn interface{} `json:"StandardsArn"`
}

// RouteSetDetails represents the RouteSetDetails schema from the OpenAPI specification
type RouteSetDetails struct {
	Corenetworkarn interface{} `json:"CoreNetworkArn,omitempty"`
	Destinationprefixlistid interface{} `json:"DestinationPrefixListId,omitempty"`
	Transitgatewayid interface{} `json:"TransitGatewayId,omitempty"`
	Carriergatewayid interface{} `json:"CarrierGatewayId,omitempty"`
	Destinationcidrblock interface{} `json:"DestinationCidrBlock,omitempty"`
	Instanceownerid interface{} `json:"InstanceOwnerId,omitempty"`
	Networkinterfaceid interface{} `json:"NetworkInterfaceId,omitempty"`
	Gatewayid interface{} `json:"GatewayId,omitempty"`
	Destinationipv6cidrblock interface{} `json:"DestinationIpv6CidrBlock,omitempty"`
	Egressonlyinternetgatewayid interface{} `json:"EgressOnlyInternetGatewayId,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Natgatewayid interface{} `json:"NatGatewayId,omitempty"`
	Vpcpeeringconnectionid interface{} `json:"VpcPeeringConnectionId,omitempty"`
	Localgatewayid interface{} `json:"LocalGatewayId,omitempty"`
	Origin interface{} `json:"Origin,omitempty"`
	State interface{} `json:"State,omitempty"`
}

// ListMembersResponse represents the ListMembersResponse schema from the OpenAPI specification
type ListMembersResponse struct {
	Members interface{} `json:"Members,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// InsightResultValue represents the InsightResultValue schema from the OpenAPI specification
type InsightResultValue struct {
	Count interface{} `json:"Count"`
	Groupbyattributevalue interface{} `json:"GroupByAttributeValue"`
}

// AwsEc2LaunchTemplateDataInstanceRequirementsDetails represents the AwsEc2LaunchTemplateDataInstanceRequirementsDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataInstanceRequirementsDetails struct {
	Acceleratorcount interface{} `json:"AcceleratorCount,omitempty"`
	Burstableperformance interface{} `json:"BurstablePerformance,omitempty"`
	Localstorage interface{} `json:"LocalStorage,omitempty"`
	Acceleratortypes interface{} `json:"AcceleratorTypes,omitempty"`
	Requirehibernatesupport interface{} `json:"RequireHibernateSupport,omitempty"`
	Localstoragetypes interface{} `json:"LocalStorageTypes,omitempty"`
	Baselineebsbandwidthmbps interface{} `json:"BaselineEbsBandwidthMbps,omitempty"`
	Cpumanufacturers interface{} `json:"CpuManufacturers,omitempty"`
	Instancegenerations interface{} `json:"InstanceGenerations,omitempty"`
	Spotmaxpricepercentageoverlowestprice interface{} `json:"SpotMaxPricePercentageOverLowestPrice,omitempty"`
	Totallocalstoragegb interface{} `json:"TotalLocalStorageGB,omitempty"`
	Acceleratornames interface{} `json:"AcceleratorNames,omitempty"`
	Baremetal interface{} `json:"BareMetal,omitempty"`
	Memorymib interface{} `json:"MemoryMiB,omitempty"`
	Ondemandmaxpricepercentageoverlowestprice interface{} `json:"OnDemandMaxPricePercentageOverLowestPrice,omitempty"`
	Vcpucount interface{} `json:"VCpuCount,omitempty"`
	Excludedinstancetypes interface{} `json:"ExcludedInstanceTypes,omitempty"`
	Memorygibpervcpu interface{} `json:"MemoryGiBPerVCpu,omitempty"`
	Acceleratormanufacturers interface{} `json:"AcceleratorManufacturers,omitempty"`
	Acceleratortotalmemorymib interface{} `json:"AcceleratorTotalMemoryMiB,omitempty"`
	Networkinterfacecount interface{} `json:"NetworkInterfaceCount,omitempty"`
}

// AwsOpenSearchServiceDomainLogPublishingOptionsDetails represents the AwsOpenSearchServiceDomainLogPublishingOptionsDetails schema from the OpenAPI specification
type AwsOpenSearchServiceDomainLogPublishingOptionsDetails struct {
	Indexslowlogs interface{} `json:"IndexSlowLogs,omitempty"`
	Searchslowlogs interface{} `json:"SearchSlowLogs,omitempty"`
	Auditlogs interface{} `json:"AuditLogs,omitempty"`
}

// AwsEfsAccessPointPosixUserDetails represents the AwsEfsAccessPointPosixUserDetails schema from the OpenAPI specification
type AwsEfsAccessPointPosixUserDetails struct {
	Uid interface{} `json:"Uid,omitempty"`
	Gid interface{} `json:"Gid,omitempty"`
	Secondarygids interface{} `json:"SecondaryGids,omitempty"`
}

// ListInvitationsResponse represents the ListInvitationsResponse schema from the OpenAPI specification
type ListInvitationsResponse struct {
	Invitations interface{} `json:"Invitations,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// AwsDynamoDbTableKeySchema represents the AwsDynamoDbTableKeySchema schema from the OpenAPI specification
type AwsDynamoDbTableKeySchema struct {
	Attributename interface{} `json:"AttributeName,omitempty"`
	Keytype interface{} `json:"KeyType,omitempty"`
}

// AwsWafRuleGroupRulesActionDetails represents the AwsWafRuleGroupRulesActionDetails schema from the OpenAPI specification
type AwsWafRuleGroupRulesActionDetails struct {
	TypeField interface{} `json:"Type,omitempty"`
}

// CreateFindingAggregatorRequest represents the CreateFindingAggregatorRequest schema from the OpenAPI specification
type CreateFindingAggregatorRequest struct {
	Regionlinkingmode interface{} `json:"RegionLinkingMode"`
	Regions interface{} `json:"Regions,omitempty"`
}

// AwsRedshiftClusterClusterSnapshotCopyStatus represents the AwsRedshiftClusterClusterSnapshotCopyStatus schema from the OpenAPI specification
type AwsRedshiftClusterClusterSnapshotCopyStatus struct {
	Destinationregion interface{} `json:"DestinationRegion,omitempty"`
	Manualsnapshotretentionperiod interface{} `json:"ManualSnapshotRetentionPeriod,omitempty"`
	Retentionperiod interface{} `json:"RetentionPeriod,omitempty"`
	Snapshotcopygrantname interface{} `json:"SnapshotCopyGrantName,omitempty"`
}

// BatchEnableStandardsRequest represents the BatchEnableStandardsRequest schema from the OpenAPI specification
type BatchEnableStandardsRequest struct {
	Standardssubscriptionrequests interface{} `json:"StandardsSubscriptionRequests"`
}

// AwsCloudFormationStackOutputsDetails represents the AwsCloudFormationStackOutputsDetails schema from the OpenAPI specification
type AwsCloudFormationStackOutputsDetails struct {
	Description interface{} `json:"Description,omitempty"`
	Outputkey interface{} `json:"OutputKey,omitempty"`
	Outputvalue interface{} `json:"OutputValue,omitempty"`
}

// AwsS3BucketBucketLifecycleConfigurationDetails represents the AwsS3BucketBucketLifecycleConfigurationDetails schema from the OpenAPI specification
type AwsS3BucketBucketLifecycleConfigurationDetails struct {
	Rules interface{} `json:"Rules,omitempty"`
}

// EnableSecurityHubRequest represents the EnableSecurityHubRequest schema from the OpenAPI specification
type EnableSecurityHubRequest struct {
	Tags interface{} `json:"Tags,omitempty"`
	Controlfindinggenerator interface{} `json:"ControlFindingGenerator,omitempty"`
	Enabledefaultstandards interface{} `json:"EnableDefaultStandards,omitempty"`
}

// UpdateStandardsControlRequest represents the UpdateStandardsControlRequest schema from the OpenAPI specification
type UpdateStandardsControlRequest struct {
	Controlstatus interface{} `json:"ControlStatus,omitempty"`
	Disabledreason interface{} `json:"DisabledReason,omitempty"`
}

// AwsS3BucketBucketLifecycleConfigurationRulesDetails represents the AwsS3BucketBucketLifecycleConfigurationRulesDetails schema from the OpenAPI specification
type AwsS3BucketBucketLifecycleConfigurationRulesDetails struct {
	Id interface{} `json:"ID,omitempty"`
	Noncurrentversionexpirationindays interface{} `json:"NoncurrentVersionExpirationInDays,omitempty"`
	Prefix interface{} `json:"Prefix,omitempty"`
	Transitions interface{} `json:"Transitions,omitempty"`
	Abortincompletemultipartupload interface{} `json:"AbortIncompleteMultipartUpload,omitempty"`
	Expirationdate interface{} `json:"ExpirationDate,omitempty"`
	Expirationindays interface{} `json:"ExpirationInDays,omitempty"`
	Noncurrentversiontransitions interface{} `json:"NoncurrentVersionTransitions,omitempty"`
	Expiredobjectdeletemarker interface{} `json:"ExpiredObjectDeleteMarker,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Filter interface{} `json:"Filter,omitempty"`
}

// StandardsManagedBy represents the StandardsManagedBy schema from the OpenAPI specification
type StandardsManagedBy struct {
	Company interface{} `json:"Company,omitempty"`
	Product interface{} `json:"Product,omitempty"`
}

// AwsRedshiftClusterDetails represents the AwsRedshiftClusterDetails schema from the OpenAPI specification
type AwsRedshiftClusterDetails struct {
	Masterusername interface{} `json:"MasterUsername,omitempty"`
	Publiclyaccessible interface{} `json:"PubliclyAccessible,omitempty"`
	Clusterversion interface{} `json:"ClusterVersion,omitempty"`
	Nodetype interface{} `json:"NodeType,omitempty"`
	Snapshotschedulestate interface{} `json:"SnapshotScheduleState,omitempty"`
	Automatedsnapshotretentionperiod interface{} `json:"AutomatedSnapshotRetentionPeriod,omitempty"`
	Elasticipstatus interface{} `json:"ElasticIpStatus,omitempty"`
	Deferredmaintenancewindows interface{} `json:"DeferredMaintenanceWindows,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Preferredmaintenancewindow interface{} `json:"PreferredMaintenanceWindow,omitempty"`
	Loggingstatus interface{} `json:"LoggingStatus,omitempty"`
	Pendingactions interface{} `json:"PendingActions,omitempty"`
	Numberofnodes interface{} `json:"NumberOfNodes,omitempty"`
	Allowversionupgrade interface{} `json:"AllowVersionUpgrade,omitempty"`
	Nextmaintenancewindowstarttime interface{} `json:"NextMaintenanceWindowStartTime,omitempty"`
	Clusteridentifier interface{} `json:"ClusterIdentifier,omitempty"`
	Clustersecuritygroups interface{} `json:"ClusterSecurityGroups,omitempty"`
	Manualsnapshotretentionperiod interface{} `json:"ManualSnapshotRetentionPeriod,omitempty"`
	Clusternodes interface{} `json:"ClusterNodes,omitempty"`
	Endpoint interface{} `json:"Endpoint,omitempty"`
	Clustercreatetime interface{} `json:"ClusterCreateTime,omitempty"`
	Expectednextsnapshotscheduletimestatus interface{} `json:"ExpectedNextSnapshotScheduleTimeStatus,omitempty"`
	Clusterpublickey interface{} `json:"ClusterPublicKey,omitempty"`
	Clustersnapshotcopystatus interface{} `json:"ClusterSnapshotCopyStatus,omitempty"`
	Dbname interface{} `json:"DBName,omitempty"`
	Clustersubnetgroupname interface{} `json:"ClusterSubnetGroupName,omitempty"`
	Enhancedvpcrouting interface{} `json:"EnhancedVpcRouting,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Maintenancetrackname interface{} `json:"MaintenanceTrackName,omitempty"`
	Elasticresizenumberofnodeoptions interface{} `json:"ElasticResizeNumberOfNodeOptions,omitempty"`
	Pendingmodifiedvalues interface{} `json:"PendingModifiedValues,omitempty"`
	Clusterrevisionnumber interface{} `json:"ClusterRevisionNumber,omitempty"`
	Clusterstatus interface{} `json:"ClusterStatus,omitempty"`
	Resizeinfo interface{} `json:"ResizeInfo,omitempty"`
	Encrypted interface{} `json:"Encrypted,omitempty"`
	Expectednextsnapshotscheduletime interface{} `json:"ExpectedNextSnapshotScheduleTime,omitempty"`
	Availabilityzone interface{} `json:"AvailabilityZone,omitempty"`
	Iamroles interface{} `json:"IamRoles,omitempty"`
	Vpcsecuritygroups interface{} `json:"VpcSecurityGroups,omitempty"`
	Clusteravailabilitystatus interface{} `json:"ClusterAvailabilityStatus,omitempty"`
	Snapshotscheduleidentifier interface{} `json:"SnapshotScheduleIdentifier,omitempty"`
	Clusterparametergroups interface{} `json:"ClusterParameterGroups,omitempty"`
	Restorestatus interface{} `json:"RestoreStatus,omitempty"`
	Hsmstatus interface{} `json:"HsmStatus,omitempty"`
}

// InviteMembersResponse represents the InviteMembersResponse schema from the OpenAPI specification
type InviteMembersResponse struct {
	Unprocessedaccounts interface{} `json:"UnprocessedAccounts,omitempty"`
}

// MapFilter represents the MapFilter schema from the OpenAPI specification
type MapFilter struct {
	Comparison interface{} `json:"Comparison,omitempty"`
	Key interface{} `json:"Key,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// FilePaths represents the FilePaths schema from the OpenAPI specification
type FilePaths struct {
	Filepath interface{} `json:"FilePath,omitempty"`
	Hash interface{} `json:"Hash,omitempty"`
	Resourceid interface{} `json:"ResourceId,omitempty"`
	Filename interface{} `json:"FileName,omitempty"`
}

// BatchDisableStandardsResponse represents the BatchDisableStandardsResponse schema from the OpenAPI specification
type BatchDisableStandardsResponse struct {
	Standardssubscriptions interface{} `json:"StandardsSubscriptions,omitempty"`
}

// AwsS3BucketServerSideEncryptionByDefault represents the AwsS3BucketServerSideEncryptionByDefault schema from the OpenAPI specification
type AwsS3BucketServerSideEncryptionByDefault struct {
	Kmsmasterkeyid interface{} `json:"KMSMasterKeyID,omitempty"`
	Ssealgorithm interface{} `json:"SSEAlgorithm,omitempty"`
}

// AwsApiGatewayEndpointConfiguration represents the AwsApiGatewayEndpointConfiguration schema from the OpenAPI specification
type AwsApiGatewayEndpointConfiguration struct {
	Types interface{} `json:"Types,omitempty"`
}

// DeleteFindingAggregatorResponse represents the DeleteFindingAggregatorResponse schema from the OpenAPI specification
type DeleteFindingAggregatorResponse struct {
}

// AwsEcsServiceDeploymentControllerDetails represents the AwsEcsServiceDeploymentControllerDetails schema from the OpenAPI specification
type AwsEcsServiceDeploymentControllerDetails struct {
	TypeField interface{} `json:"Type,omitempty"`
}

// AwsApiGatewayAccessLogSettings represents the AwsApiGatewayAccessLogSettings schema from the OpenAPI specification
type AwsApiGatewayAccessLogSettings struct {
	Destinationarn interface{} `json:"DestinationArn,omitempty"`
	Format interface{} `json:"Format,omitempty"`
}

// AwsEc2NetworkInterfaceIpV6AddressDetail represents the AwsEc2NetworkInterfaceIpV6AddressDetail schema from the OpenAPI specification
type AwsEc2NetworkInterfaceIpV6AddressDetail struct {
	Ipv6address interface{} `json:"IpV6Address,omitempty"`
}

// AwsEcsServiceLoadBalancersDetails represents the AwsEcsServiceLoadBalancersDetails schema from the OpenAPI specification
type AwsEcsServiceLoadBalancersDetails struct {
	Containername interface{} `json:"ContainerName,omitempty"`
	Containerport interface{} `json:"ContainerPort,omitempty"`
	Loadbalancername interface{} `json:"LoadBalancerName,omitempty"`
	Targetgrouparn interface{} `json:"TargetGroupArn,omitempty"`
}

// AwsCloudFrontDistributionDefaultCacheBehavior represents the AwsCloudFrontDistributionDefaultCacheBehavior schema from the OpenAPI specification
type AwsCloudFrontDistributionDefaultCacheBehavior struct {
	Viewerprotocolpolicy interface{} `json:"ViewerProtocolPolicy,omitempty"`
}

// AwsGuardDutyDetectorDataSourcesMalwareProtectionScanEc2InstanceWithFindingsDetails represents the AwsGuardDutyDetectorDataSourcesMalwareProtectionScanEc2InstanceWithFindingsDetails schema from the OpenAPI specification
type AwsGuardDutyDetectorDataSourcesMalwareProtectionScanEc2InstanceWithFindingsDetails struct {
	Ebsvolumes interface{} `json:"EbsVolumes,omitempty"`
}

// ListMembersRequest represents the ListMembersRequest schema from the OpenAPI specification
type ListMembersRequest struct {
}

// AwsCertificateManagerCertificateResourceRecord represents the AwsCertificateManagerCertificateResourceRecord schema from the OpenAPI specification
type AwsCertificateManagerCertificateResourceRecord struct {
	Name interface{} `json:"Name,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// AwsDynamoDbTableLocalSecondaryIndex represents the AwsDynamoDbTableLocalSecondaryIndex schema from the OpenAPI specification
type AwsDynamoDbTableLocalSecondaryIndex struct {
	Indexname interface{} `json:"IndexName,omitempty"`
	Keyschema interface{} `json:"KeySchema,omitempty"`
	Projection interface{} `json:"Projection,omitempty"`
	Indexarn interface{} `json:"IndexArn,omitempty"`
}

// AwsRedshiftClusterLoggingStatus represents the AwsRedshiftClusterLoggingStatus schema from the OpenAPI specification
type AwsRedshiftClusterLoggingStatus struct {
	Bucketname interface{} `json:"BucketName,omitempty"`
	Lastfailuremessage interface{} `json:"LastFailureMessage,omitempty"`
	Lastfailuretime interface{} `json:"LastFailureTime,omitempty"`
	Lastsuccessfuldeliverytime interface{} `json:"LastSuccessfulDeliveryTime,omitempty"`
	Loggingenabled interface{} `json:"LoggingEnabled,omitempty"`
	S3keyprefix interface{} `json:"S3KeyPrefix,omitempty"`
}

// AwsEcsServicePlacementConstraintsDetails represents the AwsEcsServicePlacementConstraintsDetails schema from the OpenAPI specification
type AwsEcsServicePlacementConstraintsDetails struct {
	TypeField interface{} `json:"Type,omitempty"`
	Expression interface{} `json:"Expression,omitempty"`
}

// AwsRdsDbInstanceAssociatedRole represents the AwsRdsDbInstanceAssociatedRole schema from the OpenAPI specification
type AwsRdsDbInstanceAssociatedRole struct {
	Featurename interface{} `json:"FeatureName,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// AwsIamGroupPolicy represents the AwsIamGroupPolicy schema from the OpenAPI specification
type AwsIamGroupPolicy struct {
	Policyname interface{} `json:"PolicyName,omitempty"`
}

// AwsElasticBeanstalkEnvironmentTier represents the AwsElasticBeanstalkEnvironmentTier schema from the OpenAPI specification
type AwsElasticBeanstalkEnvironmentTier struct {
	Version interface{} `json:"Version,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// Action represents the Action schema from the OpenAPI specification
type Action struct {
	Portprobeaction interface{} `json:"PortProbeAction,omitempty"`
	Actiontype interface{} `json:"ActionType,omitempty"`
	Awsapicallaction interface{} `json:"AwsApiCallAction,omitempty"`
	Dnsrequestaction interface{} `json:"DnsRequestAction,omitempty"`
	Networkconnectionaction interface{} `json:"NetworkConnectionAction,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsFirelensConfigurationDetails represents the AwsEcsTaskDefinitionContainerDefinitionsFirelensConfigurationDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsFirelensConfigurationDetails struct {
	Options interface{} `json:"Options,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// ActionTarget represents the ActionTarget schema from the OpenAPI specification
type ActionTarget struct {
	Actiontargetarn interface{} `json:"ActionTargetArn"`
	Description interface{} `json:"Description"`
	Name interface{} `json:"Name"`
}

// AwsCertificateManagerCertificateRenewalSummary represents the AwsCertificateManagerCertificateRenewalSummary schema from the OpenAPI specification
type AwsCertificateManagerCertificateRenewalSummary struct {
	Renewalstatusreason interface{} `json:"RenewalStatusReason,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Domainvalidationoptions interface{} `json:"DomainValidationOptions,omitempty"`
	Renewalstatus interface{} `json:"RenewalStatus,omitempty"`
}

// Threat represents the Threat schema from the OpenAPI specification
type Threat struct {
	Itemcount interface{} `json:"ItemCount,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Severity interface{} `json:"Severity,omitempty"`
	Filepaths interface{} `json:"FilePaths,omitempty"`
}

// AwsEc2LaunchTemplateDataPlacementDetails represents the AwsEc2LaunchTemplateDataPlacementDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataPlacementDetails struct {
	Partitionnumber interface{} `json:"PartitionNumber,omitempty"`
	Spreaddomain interface{} `json:"SpreadDomain,omitempty"`
	Tenancy interface{} `json:"Tenancy,omitempty"`
	Affinity interface{} `json:"Affinity,omitempty"`
	Availabilityzone interface{} `json:"AvailabilityZone,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
	Hostid interface{} `json:"HostId,omitempty"`
	Hostresourcegrouparn interface{} `json:"HostResourceGroupArn,omitempty"`
}

// AwsRdsDbClusterSnapshotDetails represents the AwsRdsDbClusterSnapshotDetails schema from the OpenAPI specification
type AwsRdsDbClusterSnapshotDetails struct {
	Storageencrypted interface{} `json:"StorageEncrypted,omitempty"`
	Licensemodel interface{} `json:"LicenseModel,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Dbclusteridentifier interface{} `json:"DbClusterIdentifier,omitempty"`
	Clustercreatetime interface{} `json:"ClusterCreateTime,omitempty"`
	Snapshotcreatetime interface{} `json:"SnapshotCreateTime,omitempty"`
	Availabilityzones interface{} `json:"AvailabilityZones,omitempty"`
	Engine interface{} `json:"Engine,omitempty"`
	Dbclustersnapshotattributes interface{} `json:"DbClusterSnapshotAttributes,omitempty"`
	Engineversion interface{} `json:"EngineVersion,omitempty"`
	Iamdatabaseauthenticationenabled interface{} `json:"IamDatabaseAuthenticationEnabled,omitempty"`
	Port interface{} `json:"Port,omitempty"`
	Masterusername interface{} `json:"MasterUsername,omitempty"`
	Snapshottype interface{} `json:"SnapshotType,omitempty"`
	Allocatedstorage interface{} `json:"AllocatedStorage,omitempty"`
	Dbclustersnapshotidentifier interface{} `json:"DbClusterSnapshotIdentifier,omitempty"`
	Percentprogress interface{} `json:"PercentProgress,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
}

// UpdateActionTargetResponse represents the UpdateActionTargetResponse schema from the OpenAPI specification
type UpdateActionTargetResponse struct {
}

// DeleteMembersResponse represents the DeleteMembersResponse schema from the OpenAPI specification
type DeleteMembersResponse struct {
	Unprocessedaccounts interface{} `json:"UnprocessedAccounts,omitempty"`
}

// AwsCloudFrontDistributionLogging represents the AwsCloudFrontDistributionLogging schema from the OpenAPI specification
type AwsCloudFrontDistributionLogging struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Includecookies interface{} `json:"IncludeCookies,omitempty"`
	Prefix interface{} `json:"Prefix,omitempty"`
	Bucket interface{} `json:"Bucket,omitempty"`
}

// AwsElbLoadBalancerConnectionSettings represents the AwsElbLoadBalancerConnectionSettings schema from the OpenAPI specification
type AwsElbLoadBalancerConnectionSettings struct {
	Idletimeout interface{} `json:"IdleTimeout,omitempty"`
}

// AwsEventSchemasRegistryDetails represents the AwsEventSchemasRegistryDetails schema from the OpenAPI specification
type AwsEventSchemasRegistryDetails struct {
	Registryarn interface{} `json:"RegistryArn,omitempty"`
	Registryname interface{} `json:"RegistryName,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// AwsWafv2RulesActionCountDetails represents the AwsWafv2RulesActionCountDetails schema from the OpenAPI specification
type AwsWafv2RulesActionCountDetails struct {
	Customrequesthandling interface{} `json:"CustomRequestHandling,omitempty"`
}

// DescribeStandardsResponse represents the DescribeStandardsResponse schema from the OpenAPI specification
type DescribeStandardsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Standards interface{} `json:"Standards,omitempty"`
}

// ListEnabledProductsForImportRequest represents the ListEnabledProductsForImportRequest schema from the OpenAPI specification
type ListEnabledProductsForImportRequest struct {
}

// AwsEcsServiceNetworkConfigurationDetails represents the AwsEcsServiceNetworkConfigurationDetails schema from the OpenAPI specification
type AwsEcsServiceNetworkConfigurationDetails struct {
	Awsvpcconfiguration interface{} `json:"AwsVpcConfiguration,omitempty"`
}

// AwsWafv2CustomRequestHandlingDetails represents the AwsWafv2CustomRequestHandlingDetails schema from the OpenAPI specification
type AwsWafv2CustomRequestHandlingDetails struct {
	Insertheaders interface{} `json:"InsertHeaders,omitempty"`
}

// FieldMap represents the FieldMap schema from the OpenAPI specification
type FieldMap struct {
}

// AwsAutoScalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesListDetails represents the AwsAutoScalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesListDetails schema from the OpenAPI specification
type AwsAutoScalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesListDetails struct {
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Weightedcapacity interface{} `json:"WeightedCapacity,omitempty"`
}

// AwsWafv2CustomHttpHeader represents the AwsWafv2CustomHttpHeader schema from the OpenAPI specification
type AwsWafv2CustomHttpHeader struct {
	Name interface{} `json:"Name,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// SecurityControl represents the SecurityControl schema from the OpenAPI specification
type SecurityControl struct {
	Securitycontrolstatus interface{} `json:"SecurityControlStatus"`
	Severityrating interface{} `json:"SeverityRating"`
	Title interface{} `json:"Title"`
	Description interface{} `json:"Description"`
	Remediationurl interface{} `json:"RemediationUrl"`
	Securitycontrolarn interface{} `json:"SecurityControlArn"`
	Securitycontrolid interface{} `json:"SecurityControlId"`
}

// DisableSecurityHubResponse represents the DisableSecurityHubResponse schema from the OpenAPI specification
type DisableSecurityHubResponse struct {
}

// EnableOrganizationAdminAccountResponse represents the EnableOrganizationAdminAccountResponse schema from the OpenAPI specification
type EnableOrganizationAdminAccountResponse struct {
}

// AwsEc2LaunchTemplateDataInstanceRequirementsAcceleratorTotalMemoryMiBDetails represents the AwsEc2LaunchTemplateDataInstanceRequirementsAcceleratorTotalMemoryMiBDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataInstanceRequirementsAcceleratorTotalMemoryMiBDetails struct {
	Max interface{} `json:"Max,omitempty"`
	Min interface{} `json:"Min,omitempty"`
}

// AwsLambdaFunctionEnvironmentError represents the AwsLambdaFunctionEnvironmentError schema from the OpenAPI specification
type AwsLambdaFunctionEnvironmentError struct {
	Message interface{} `json:"Message,omitempty"`
	Errorcode interface{} `json:"ErrorCode,omitempty"`
}

// GetInsightsRequest represents the GetInsightsRequest schema from the OpenAPI specification
type GetInsightsRequest struct {
	Insightarns interface{} `json:"InsightArns,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// RelatedFinding represents the RelatedFinding schema from the OpenAPI specification
type RelatedFinding struct {
	Id interface{} `json:"Id"`
	Productarn interface{} `json:"ProductArn"`
}

// AwsEc2InstanceDetails represents the AwsEc2InstanceDetails schema from the OpenAPI specification
type AwsEc2InstanceDetails struct {
	Networkinterfaces interface{} `json:"NetworkInterfaces,omitempty"`
	Subnetid interface{} `json:"SubnetId,omitempty"`
	Virtualizationtype interface{} `json:"VirtualizationType,omitempty"`
	Keyname interface{} `json:"KeyName,omitempty"`
	Monitoring interface{} `json:"Monitoring,omitempty"`
	Ipv4addresses interface{} `json:"IpV4Addresses,omitempty"`
	Launchedat interface{} `json:"LaunchedAt,omitempty"`
	Imageid interface{} `json:"ImageId,omitempty"`
	Ipv6addresses interface{} `json:"IpV6Addresses,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Iaminstanceprofilearn interface{} `json:"IamInstanceProfileArn,omitempty"`
	Metadataoptions interface{} `json:"MetadataOptions,omitempty"`
}

// AwsElasticsearchDomainServiceSoftwareOptions represents the AwsElasticsearchDomainServiceSoftwareOptions schema from the OpenAPI specification
type AwsElasticsearchDomainServiceSoftwareOptions struct {
	Currentversion interface{} `json:"CurrentVersion,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Newversion interface{} `json:"NewVersion,omitempty"`
	Updateavailable interface{} `json:"UpdateAvailable,omitempty"`
	Updatestatus interface{} `json:"UpdateStatus,omitempty"`
	Automatedupdatedate interface{} `json:"AutomatedUpdateDate,omitempty"`
	Cancellable interface{} `json:"Cancellable,omitempty"`
}

// AwsElbLoadBalancerPolicies represents the AwsElbLoadBalancerPolicies schema from the OpenAPI specification
type AwsElbLoadBalancerPolicies struct {
	Lbcookiestickinesspolicies interface{} `json:"LbCookieStickinessPolicies,omitempty"`
	Otherpolicies interface{} `json:"OtherPolicies,omitempty"`
	Appcookiestickinesspolicies interface{} `json:"AppCookieStickinessPolicies,omitempty"`
}

// PortRangeFromTo represents the PortRangeFromTo schema from the OpenAPI specification
type PortRangeFromTo struct {
	From interface{} `json:"From,omitempty"`
	To interface{} `json:"To,omitempty"`
}

// BatchUpdateStandardsControlAssociationsRequest represents the BatchUpdateStandardsControlAssociationsRequest schema from the OpenAPI specification
type BatchUpdateStandardsControlAssociationsRequest struct {
	Standardscontrolassociationupdates interface{} `json:"StandardsControlAssociationUpdates"`
}

// DeclineInvitationsResponse represents the DeclineInvitationsResponse schema from the OpenAPI specification
type DeclineInvitationsResponse struct {
	Unprocessedaccounts interface{} `json:"UnprocessedAccounts,omitempty"`
}

// AwsEc2VpnConnectionVgwTelemetryDetails represents the AwsEc2VpnConnectionVgwTelemetryDetails schema from the OpenAPI specification
type AwsEc2VpnConnectionVgwTelemetryDetails struct {
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Acceptedroutecount interface{} `json:"AcceptedRouteCount,omitempty"`
	Certificatearn interface{} `json:"CertificateArn,omitempty"`
	Laststatuschange interface{} `json:"LastStatusChange,omitempty"`
	Outsideipaddress interface{} `json:"OutsideIpAddress,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// AwsEc2LaunchTemplateDataIamInstanceProfileDetails represents the AwsEc2LaunchTemplateDataIamInstanceProfileDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataIamInstanceProfileDetails struct {
	Arn interface{} `json:"Arn,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// AwsSecurityFindingFilters represents the AwsSecurityFindingFilters schema from the OpenAPI specification
type AwsSecurityFindingFilters struct {
	Networkdirection interface{} `json:"NetworkDirection,omitempty"`
	Criticality interface{} `json:"Criticality,omitempty"`
	Resourcecontainerimagename interface{} `json:"ResourceContainerImageName,omitempty"`
	Networksourcemac interface{} `json:"NetworkSourceMac,omitempty"`
	Relatedfindingsid interface{} `json:"RelatedFindingsId,omitempty"`
	Resourcecontainerlaunchedat interface{} `json:"ResourceContainerLaunchedAt,omitempty"`
	Resourceawss3bucketownerid interface{} `json:"ResourceAwsS3BucketOwnerId,omitempty"`
	Relatedfindingsproductarn interface{} `json:"RelatedFindingsProductArn,omitempty"`
	Lastobservedat interface{} `json:"LastObservedAt,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Severitynormalized interface{} `json:"SeverityNormalized,omitempty"`
	Resourcedetailsother interface{} `json:"ResourceDetailsOther,omitempty"`
	Resourceregion interface{} `json:"ResourceRegion,omitempty"`
	Resourceawsiamaccesskeystatus interface{} `json:"ResourceAwsIamAccessKeyStatus,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Verificationstate interface{} `json:"VerificationState,omitempty"`
	Complianceassociatedstandardsid interface{} `json:"ComplianceAssociatedStandardsId,omitempty"`
	Processparentpid interface{} `json:"ProcessParentPid,omitempty"`
	Findingproviderfieldsconfidence interface{} `json:"FindingProviderFieldsConfidence,omitempty"`
	Firstobservedat interface{} `json:"FirstObservedAt,omitempty"`
	Networkdestinationport interface{} `json:"NetworkDestinationPort,omitempty"`
	Workflowstate interface{} `json:"WorkflowState,omitempty"`
	Malwaretype interface{} `json:"MalwareType,omitempty"`
	Resourceawsiamaccesskeyprincipalname interface{} `json:"ResourceAwsIamAccessKeyPrincipalName,omitempty"`
	Networksourceipv6 interface{} `json:"NetworkSourceIpV6,omitempty"`
	Resourceawss3bucketownername interface{} `json:"ResourceAwsS3BucketOwnerName,omitempty"`
	Generatorid interface{} `json:"GeneratorId,omitempty"`
	Resourcecontainerimageid interface{} `json:"ResourceContainerImageId,omitempty"`
	Resourcetype interface{} `json:"ResourceType,omitempty"`
	Threatintelindicatorsource interface{} `json:"ThreatIntelIndicatorSource,omitempty"`
	Title interface{} `json:"Title,omitempty"`
	Processpath interface{} `json:"ProcessPath,omitempty"`
	Findingproviderfieldsseverityoriginal interface{} `json:"FindingProviderFieldsSeverityOriginal,omitempty"`
	Resourcepartition interface{} `json:"ResourcePartition,omitempty"`
	Threatintelindicatorvalue interface{} `json:"ThreatIntelIndicatorValue,omitempty"`
	Compliancestatus interface{} `json:"ComplianceStatus,omitempty"`
	Noteupdatedat interface{} `json:"NoteUpdatedAt,omitempty"`
	Processlaunchedat interface{} `json:"ProcessLaunchedAt,omitempty"`
	Findingproviderfieldscriticality interface{} `json:"FindingProviderFieldsCriticality,omitempty"`
	Productname interface{} `json:"ProductName,omitempty"`
	Networksourceipv4 interface{} `json:"NetworkSourceIpV4,omitempty"`
	Resourceawsiamaccesskeyusername interface{} `json:"ResourceAwsIamAccessKeyUserName,omitempty"`
	Resourceawsiamaccesskeycreatedat interface{} `json:"ResourceAwsIamAccessKeyCreatedAt,omitempty"`
	Networksourcedomain interface{} `json:"NetworkSourceDomain,omitempty"`
	Severityproduct interface{} `json:"SeverityProduct,omitempty"`
	Threatintelindicatorlastobservedat interface{} `json:"ThreatIntelIndicatorLastObservedAt,omitempty"`
	Resourceawsec2instancelaunchedat interface{} `json:"ResourceAwsEc2InstanceLaunchedAt,omitempty"`
	Resourcetags interface{} `json:"ResourceTags,omitempty"`
	Resourceid interface{} `json:"ResourceId,omitempty"`
	Findingproviderfieldsseveritylabel interface{} `json:"FindingProviderFieldsSeverityLabel,omitempty"`
	Resourceawsec2instancesubnetid interface{} `json:"ResourceAwsEc2InstanceSubnetId,omitempty"`
	Malwarepath interface{} `json:"MalwarePath,omitempty"`
	Threatintelindicatorsourceurl interface{} `json:"ThreatIntelIndicatorSourceUrl,omitempty"`
	Resourceawsec2instanceimageid interface{} `json:"ResourceAwsEc2InstanceImageId,omitempty"`
	Resourcecontainername interface{} `json:"ResourceContainerName,omitempty"`
	Threatintelindicatortype interface{} `json:"ThreatIntelIndicatorType,omitempty"`
	Companyname interface{} `json:"CompanyName,omitempty"`
	Region interface{} `json:"Region,omitempty"`
	Productfields interface{} `json:"ProductFields,omitempty"`
	Threatintelindicatorcategory interface{} `json:"ThreatIntelIndicatorCategory,omitempty"`
	Networkdestinationipv4 interface{} `json:"NetworkDestinationIpV4,omitempty"`
	Awsaccountid interface{} `json:"AwsAccountId,omitempty"`
	Processterminatedat interface{} `json:"ProcessTerminatedAt,omitempty"`
	Severitylabel interface{} `json:"SeverityLabel,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Resourceawsec2instancevpcid interface{} `json:"ResourceAwsEc2InstanceVpcId,omitempty"`
	Malwarestate interface{} `json:"MalwareState,omitempty"`
	Resourceawsec2instancekeyname interface{} `json:"ResourceAwsEc2InstanceKeyName,omitempty"`
	Findingproviderfieldsrelatedfindingsid interface{} `json:"FindingProviderFieldsRelatedFindingsId,omitempty"`
	Resourceawsec2instanceiaminstanceprofilearn interface{} `json:"ResourceAwsEc2InstanceIamInstanceProfileArn,omitempty"`
	Confidence interface{} `json:"Confidence,omitempty"`
	Notetext interface{} `json:"NoteText,omitempty"`
	Sample interface{} `json:"Sample,omitempty"`
	Workflowstatus interface{} `json:"WorkflowStatus,omitempty"`
	Recordstate interface{} `json:"RecordState,omitempty"`
	Findingproviderfieldsrelatedfindingsproductarn interface{} `json:"FindingProviderFieldsRelatedFindingsProductArn,omitempty"`
	Productarn interface{} `json:"ProductArn,omitempty"`
	Malwarename interface{} `json:"MalwareName,omitempty"`
	Resourceawsec2instanceipv4addresses interface{} `json:"ResourceAwsEc2InstanceIpV4Addresses,omitempty"`
	Compliancesecuritycontrolid interface{} `json:"ComplianceSecurityControlId,omitempty"`
	Processpid interface{} `json:"ProcessPid,omitempty"`
	Resourceawsiamuserusername interface{} `json:"ResourceAwsIamUserUserName,omitempty"`
	Findingproviderfieldstypes interface{} `json:"FindingProviderFieldsTypes,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Userdefinedfields interface{} `json:"UserDefinedFields,omitempty"`
	Noteupdatedby interface{} `json:"NoteUpdatedBy,omitempty"`
	Resourceawsec2instanceipv6addresses interface{} `json:"ResourceAwsEc2InstanceIpV6Addresses,omitempty"`
	Resourceawsec2instancetype interface{} `json:"ResourceAwsEc2InstanceType,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Processname interface{} `json:"ProcessName,omitempty"`
	Networksourceport interface{} `json:"NetworkSourcePort,omitempty"`
	Networkprotocol interface{} `json:"NetworkProtocol,omitempty"`
	Recommendationtext interface{} `json:"RecommendationText,omitempty"`
	Networkdestinationdomain interface{} `json:"NetworkDestinationDomain,omitempty"`
	Sourceurl interface{} `json:"SourceUrl,omitempty"`
	Keyword interface{} `json:"Keyword,omitempty"`
	Networkdestinationipv6 interface{} `json:"NetworkDestinationIpV6,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsSecretsDetails represents the AwsEcsTaskDefinitionContainerDefinitionsSecretsDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsSecretsDetails struct {
	Name interface{} `json:"Name,omitempty"`
	Valuefrom interface{} `json:"ValueFrom,omitempty"`
}

// VulnerabilityVendor represents the VulnerabilityVendor schema from the OpenAPI specification
type VulnerabilityVendor struct {
	Vendorupdatedat interface{} `json:"VendorUpdatedAt,omitempty"`
	Name interface{} `json:"Name"`
	Url interface{} `json:"Url,omitempty"`
	Vendorcreatedat interface{} `json:"VendorCreatedAt,omitempty"`
	Vendorseverity interface{} `json:"VendorSeverity,omitempty"`
}

// AwsWafRegionalRulePredicateListDetails represents the AwsWafRegionalRulePredicateListDetails schema from the OpenAPI specification
type AwsWafRegionalRulePredicateListDetails struct {
	Dataid interface{} `json:"DataId,omitempty"`
	Negated interface{} `json:"Negated,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// SoftwarePackage represents the SoftwarePackage schema from the OpenAPI specification
type SoftwarePackage struct {
	Fixedinversion interface{} `json:"FixedInVersion,omitempty"`
	Sourcelayerarn interface{} `json:"SourceLayerArn,omitempty"`
	Version interface{} `json:"Version,omitempty"`
	Architecture interface{} `json:"Architecture,omitempty"`
	Epoch interface{} `json:"Epoch,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Remediation interface{} `json:"Remediation,omitempty"`
	Sourcelayerhash interface{} `json:"SourceLayerHash,omitempty"`
	Packagemanager interface{} `json:"PackageManager,omitempty"`
	Release interface{} `json:"Release,omitempty"`
	Filepath interface{} `json:"FilePath,omitempty"`
}

// GetEnabledStandardsRequest represents the GetEnabledStandardsRequest schema from the OpenAPI specification
type GetEnabledStandardsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Standardssubscriptionarns interface{} `json:"StandardsSubscriptionArns,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// AwsEc2VpcEndpointServiceServiceTypeDetails represents the AwsEc2VpcEndpointServiceServiceTypeDetails schema from the OpenAPI specification
type AwsEc2VpcEndpointServiceServiceTypeDetails struct {
	Servicetype interface{} `json:"ServiceType,omitempty"`
}

// AwsStepFunctionStateMachineLoggingConfigurationDestinationsDetails represents the AwsStepFunctionStateMachineLoggingConfigurationDestinationsDetails schema from the OpenAPI specification
type AwsStepFunctionStateMachineLoggingConfigurationDestinationsDetails struct {
	Cloudwatchlogsloggroup interface{} `json:"CloudWatchLogsLogGroup,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsRepositoryCredentialsDetails represents the AwsEcsTaskDefinitionContainerDefinitionsRepositoryCredentialsDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsRepositoryCredentialsDetails struct {
	Credentialsparameter interface{} `json:"CredentialsParameter,omitempty"`
}

// AwsSsmPatch represents the AwsSsmPatch schema from the OpenAPI specification
type AwsSsmPatch struct {
	Compliancesummary interface{} `json:"ComplianceSummary,omitempty"`
}

// DisableImportFindingsForProductRequest represents the DisableImportFindingsForProductRequest schema from the OpenAPI specification
type DisableImportFindingsForProductRequest struct {
}

// RuleGroupVariablesPortSetsDetails represents the RuleGroupVariablesPortSetsDetails schema from the OpenAPI specification
type RuleGroupVariablesPortSetsDetails struct {
	Definition interface{} `json:"Definition,omitempty"`
}

// AwsS3BucketBucketLifecycleConfigurationRulesAbortIncompleteMultipartUploadDetails represents the AwsS3BucketBucketLifecycleConfigurationRulesAbortIncompleteMultipartUploadDetails schema from the OpenAPI specification
type AwsS3BucketBucketLifecycleConfigurationRulesAbortIncompleteMultipartUploadDetails struct {
	Daysafterinitiation interface{} `json:"DaysAfterInitiation,omitempty"`
}

// CreateFindingAggregatorResponse represents the CreateFindingAggregatorResponse schema from the OpenAPI specification
type CreateFindingAggregatorResponse struct {
	Regions interface{} `json:"Regions,omitempty"`
	Findingaggregationregion interface{} `json:"FindingAggregationRegion,omitempty"`
	Findingaggregatorarn interface{} `json:"FindingAggregatorArn,omitempty"`
	Regionlinkingmode interface{} `json:"RegionLinkingMode,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsVolumesFromDetails represents the AwsEcsTaskDefinitionContainerDefinitionsVolumesFromDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsVolumesFromDetails struct {
	Sourcecontainer interface{} `json:"SourceContainer,omitempty"`
	Readonly interface{} `json:"ReadOnly,omitempty"`
}

// AwsGuardDutyDetectorDataSourcesFlowLogsDetails represents the AwsGuardDutyDetectorDataSourcesFlowLogsDetails schema from the OpenAPI specification
type AwsGuardDutyDetectorDataSourcesFlowLogsDetails struct {
	Status interface{} `json:"Status,omitempty"`
}

// DisassociateMembersResponse represents the DisassociateMembersResponse schema from the OpenAPI specification
type DisassociateMembersResponse struct {
}

// AwsCloudFrontDistributionCacheBehaviors represents the AwsCloudFrontDistributionCacheBehaviors schema from the OpenAPI specification
type AwsCloudFrontDistributionCacheBehaviors struct {
	Items interface{} `json:"Items,omitempty"`
}

// AwsEc2LaunchTemplateDataHibernationOptionsDetails represents the AwsEc2LaunchTemplateDataHibernationOptionsDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataHibernationOptionsDetails struct {
	Configured interface{} `json:"Configured,omitempty"`
}

// DescribeActionTargetsRequest represents the DescribeActionTargetsRequest schema from the OpenAPI specification
type DescribeActionTargetsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Actiontargetarns interface{} `json:"ActionTargetArns,omitempty"`
}

// AwsNetworkFirewallRuleGroupDetails represents the AwsNetworkFirewallRuleGroupDetails schema from the OpenAPI specification
type AwsNetworkFirewallRuleGroupDetails struct {
	Capacity interface{} `json:"Capacity,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Rulegroup interface{} `json:"RuleGroup,omitempty"`
	Rulegrouparn interface{} `json:"RuleGroupArn,omitempty"`
	Rulegroupid interface{} `json:"RuleGroupId,omitempty"`
	Rulegroupname interface{} `json:"RuleGroupName,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// AwsOpenSearchServiceDomainAdvancedSecurityOptionsDetails represents the AwsOpenSearchServiceDomainAdvancedSecurityOptionsDetails schema from the OpenAPI specification
type AwsOpenSearchServiceDomainAdvancedSecurityOptionsDetails struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Internaluserdatabaseenabled interface{} `json:"InternalUserDatabaseEnabled,omitempty"`
	Masteruseroptions interface{} `json:"MasterUserOptions,omitempty"`
}

// AwsBackupBackupVaultDetails represents the AwsBackupBackupVaultDetails schema from the OpenAPI specification
type AwsBackupBackupVaultDetails struct {
	Accesspolicy interface{} `json:"AccessPolicy,omitempty"`
	Backupvaultarn interface{} `json:"BackupVaultArn,omitempty"`
	Backupvaultname interface{} `json:"BackupVaultName,omitempty"`
	Encryptionkeyarn interface{} `json:"EncryptionKeyArn,omitempty"`
	Notifications interface{} `json:"Notifications,omitempty"`
}

// AwsOpenSearchServiceDomainClusterConfigDetails represents the AwsOpenSearchServiceDomainClusterConfigDetails schema from the OpenAPI specification
type AwsOpenSearchServiceDomainClusterConfigDetails struct {
	Dedicatedmastercount interface{} `json:"DedicatedMasterCount,omitempty"`
	Dedicatedmastertype interface{} `json:"DedicatedMasterType,omitempty"`
	Warmtype interface{} `json:"WarmType,omitempty"`
	Zoneawarenessconfig interface{} `json:"ZoneAwarenessConfig,omitempty"`
	Dedicatedmasterenabled interface{} `json:"DedicatedMasterEnabled,omitempty"`
	Warmenabled interface{} `json:"WarmEnabled,omitempty"`
	Instancecount interface{} `json:"InstanceCount,omitempty"`
	Zoneawarenessenabled interface{} `json:"ZoneAwarenessEnabled,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Warmcount interface{} `json:"WarmCount,omitempty"`
}

// DeleteFindingAggregatorRequest represents the DeleteFindingAggregatorRequest schema from the OpenAPI specification
type DeleteFindingAggregatorRequest struct {
}

// GetInsightResultsResponse represents the GetInsightResultsResponse schema from the OpenAPI specification
type GetInsightResultsResponse struct {
	Insightresults interface{} `json:"InsightResults"`
}

// AwsWafRateBasedRuleMatchPredicate represents the AwsWafRateBasedRuleMatchPredicate schema from the OpenAPI specification
type AwsWafRateBasedRuleMatchPredicate struct {
	Dataid interface{} `json:"DataId,omitempty"`
	Negated interface{} `json:"Negated,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// ListAutomationRulesResponse represents the ListAutomationRulesResponse schema from the OpenAPI specification
type ListAutomationRulesResponse struct {
	Automationrulesmetadata interface{} `json:"AutomationRulesMetadata,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// AwsGuardDutyDetectorDataSourcesCloudTrailDetails represents the AwsGuardDutyDetectorDataSourcesCloudTrailDetails schema from the OpenAPI specification
type AwsGuardDutyDetectorDataSourcesCloudTrailDetails struct {
	Status interface{} `json:"Status,omitempty"`
}

// AwsS3BucketBucketLifecycleConfigurationRulesFilterPredicateOperandsDetails represents the AwsS3BucketBucketLifecycleConfigurationRulesFilterPredicateOperandsDetails schema from the OpenAPI specification
type AwsS3BucketBucketLifecycleConfigurationRulesFilterPredicateOperandsDetails struct {
	Prefix interface{} `json:"Prefix,omitempty"`
	Tag interface{} `json:"Tag,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// AwsEcsClusterConfigurationExecuteCommandConfigurationDetails represents the AwsEcsClusterConfigurationExecuteCommandConfigurationDetails schema from the OpenAPI specification
type AwsEcsClusterConfigurationExecuteCommandConfigurationDetails struct {
	Logconfiguration interface{} `json:"LogConfiguration,omitempty"`
	Logging interface{} `json:"Logging,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
}

// StatusReason represents the StatusReason schema from the OpenAPI specification
type StatusReason struct {
	Description interface{} `json:"Description,omitempty"`
	Reasoncode interface{} `json:"ReasonCode"`
}

// Page represents the Page schema from the OpenAPI specification
type Page struct {
	Linerange interface{} `json:"LineRange,omitempty"`
	Offsetrange interface{} `json:"OffsetRange,omitempty"`
	Pagenumber interface{} `json:"PageNumber,omitempty"`
}

// AwsCloudFrontDistributionOriginCustomOriginConfig represents the AwsCloudFrontDistributionOriginCustomOriginConfig schema from the OpenAPI specification
type AwsCloudFrontDistributionOriginCustomOriginConfig struct {
	Httpport interface{} `json:"HttpPort,omitempty"`
	Httpsport interface{} `json:"HttpsPort,omitempty"`
	Originkeepalivetimeout interface{} `json:"OriginKeepaliveTimeout,omitempty"`
	Originprotocolpolicy interface{} `json:"OriginProtocolPolicy,omitempty"`
	Originreadtimeout interface{} `json:"OriginReadTimeout,omitempty"`
	Originsslprotocols interface{} `json:"OriginSslProtocols,omitempty"`
}

// GeoLocation represents the GeoLocation schema from the OpenAPI specification
type GeoLocation struct {
	Lat interface{} `json:"Lat,omitempty"`
	Lon interface{} `json:"Lon,omitempty"`
}

// AwsEc2LaunchTemplateDataMonitoringDetails represents the AwsEc2LaunchTemplateDataMonitoringDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataMonitoringDetails struct {
	Enabled interface{} `json:"Enabled,omitempty"`
}

// AwsAthenaWorkGroupConfigurationResultConfigurationDetails represents the AwsAthenaWorkGroupConfigurationResultConfigurationDetails schema from the OpenAPI specification
type AwsAthenaWorkGroupConfigurationResultConfigurationDetails struct {
	Encryptionconfiguration interface{} `json:"EncryptionConfiguration,omitempty"`
}

// AwsAutoScalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionDetails represents the AwsAutoScalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionDetails schema from the OpenAPI specification
type AwsAutoScalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionDetails struct {
	Ondemandbasecapacity interface{} `json:"OnDemandBaseCapacity,omitempty"`
	Ondemandpercentageabovebasecapacity interface{} `json:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	Spotallocationstrategy interface{} `json:"SpotAllocationStrategy,omitempty"`
	Spotinstancepools interface{} `json:"SpotInstancePools,omitempty"`
	Spotmaxprice interface{} `json:"SpotMaxPrice,omitempty"`
	Ondemandallocationstrategy interface{} `json:"OnDemandAllocationStrategy,omitempty"`
}

// Result represents the Result schema from the OpenAPI specification
type Result struct {
	Accountid interface{} `json:"AccountId,omitempty"`
	Processingresult interface{} `json:"ProcessingResult,omitempty"`
}

// AwsCloudFormationStackDriftInformationDetails represents the AwsCloudFormationStackDriftInformationDetails schema from the OpenAPI specification
type AwsCloudFormationStackDriftInformationDetails struct {
	Stackdriftstatus interface{} `json:"StackDriftStatus,omitempty"`
}

// AwsAutoScalingAutoScalingGroupAvailabilityZonesListDetails represents the AwsAutoScalingAutoScalingGroupAvailabilityZonesListDetails schema from the OpenAPI specification
type AwsAutoScalingAutoScalingGroupAvailabilityZonesListDetails struct {
	Value interface{} `json:"Value,omitempty"`
}

// CreateActionTargetRequest represents the CreateActionTargetRequest schema from the OpenAPI specification
type CreateActionTargetRequest struct {
	Description interface{} `json:"Description"`
	Id interface{} `json:"Id"`
	Name interface{} `json:"Name"`
}

// ThreatIntelIndicator represents the ThreatIntelIndicator schema from the OpenAPI specification
type ThreatIntelIndicator struct {
	Source interface{} `json:"Source,omitempty"`
	Sourceurl interface{} `json:"SourceUrl,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Value interface{} `json:"Value,omitempty"`
	Category interface{} `json:"Category,omitempty"`
	Lastobservedat interface{} `json:"LastObservedAt,omitempty"`
}

// AwsElbLoadBalancerAccessLog represents the AwsElbLoadBalancerAccessLog schema from the OpenAPI specification
type AwsElbLoadBalancerAccessLog struct {
	S3bucketname interface{} `json:"S3BucketName,omitempty"`
	S3bucketprefix interface{} `json:"S3BucketPrefix,omitempty"`
	Emitinterval interface{} `json:"EmitInterval,omitempty"`
	Enabled interface{} `json:"Enabled,omitempty"`
}

// AwsEc2SubnetDetails represents the AwsEc2SubnetDetails schema from the OpenAPI specification
type AwsEc2SubnetDetails struct {
	Mappubliciponlaunch interface{} `json:"MapPublicIpOnLaunch,omitempty"`
	Ownerid interface{} `json:"OwnerId,omitempty"`
	Cidrblock interface{} `json:"CidrBlock,omitempty"`
	Subnetarn interface{} `json:"SubnetArn,omitempty"`
	Subnetid interface{} `json:"SubnetId,omitempty"`
	Assignipv6addressoncreation interface{} `json:"AssignIpv6AddressOnCreation,omitempty"`
	Availableipaddresscount interface{} `json:"AvailableIpAddressCount,omitempty"`
	State interface{} `json:"State,omitempty"`
	Availabilityzoneid interface{} `json:"AvailabilityZoneId,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Availabilityzone interface{} `json:"AvailabilityZone,omitempty"`
	Defaultforaz interface{} `json:"DefaultForAz,omitempty"`
	Ipv6cidrblockassociationset interface{} `json:"Ipv6CidrBlockAssociationSet,omitempty"`
}

// AwsGuardDutyDetectorDataSourcesDetails represents the AwsGuardDutyDetectorDataSourcesDetails schema from the OpenAPI specification
type AwsGuardDutyDetectorDataSourcesDetails struct {
	Flowlogs interface{} `json:"FlowLogs,omitempty"`
	Kubernetes interface{} `json:"Kubernetes,omitempty"`
	Malwareprotection interface{} `json:"MalwareProtection,omitempty"`
	S3logs interface{} `json:"S3Logs,omitempty"`
	Cloudtrail interface{} `json:"CloudTrail,omitempty"`
	Dnslogs interface{} `json:"DnsLogs,omitempty"`
}

// StandardsControlAssociationSummary represents the StandardsControlAssociationSummary schema from the OpenAPI specification
type StandardsControlAssociationSummary struct {
	Associationstatus interface{} `json:"AssociationStatus"`
	Securitycontrolarn interface{} `json:"SecurityControlArn"`
	Relatedrequirements interface{} `json:"RelatedRequirements,omitempty"`
	Securitycontrolid interface{} `json:"SecurityControlId"`
	Standardsarn interface{} `json:"StandardsArn"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Updatedreason interface{} `json:"UpdatedReason,omitempty"`
	Standardscontroldescription interface{} `json:"StandardsControlDescription,omitempty"`
	Standardscontroltitle interface{} `json:"StandardsControlTitle,omitempty"`
}

// AwsCloudFrontDistributionDetails represents the AwsCloudFrontDistributionDetails schema from the OpenAPI specification
type AwsCloudFrontDistributionDetails struct {
	Logging interface{} `json:"Logging,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Webaclid interface{} `json:"WebAclId,omitempty"`
	Cachebehaviors interface{} `json:"CacheBehaviors,omitempty"`
	Defaultcachebehavior interface{} `json:"DefaultCacheBehavior,omitempty"`
	Defaultrootobject interface{} `json:"DefaultRootObject,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Etag interface{} `json:"ETag,omitempty"`
	Origingroups interface{} `json:"OriginGroups,omitempty"`
	Origins interface{} `json:"Origins,omitempty"`
	Viewercertificate interface{} `json:"ViewerCertificate,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
}

// AwsXrayEncryptionConfigDetails represents the AwsXrayEncryptionConfigDetails schema from the OpenAPI specification
type AwsXrayEncryptionConfigDetails struct {
	Keyid interface{} `json:"KeyId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// RuleGroupSourceStatelessRuleMatchAttributesTcpFlags represents the RuleGroupSourceStatelessRuleMatchAttributesTcpFlags schema from the OpenAPI specification
type RuleGroupSourceStatelessRuleMatchAttributesTcpFlags struct {
	Flags interface{} `json:"Flags,omitempty"`
	Masks interface{} `json:"Masks,omitempty"`
}

// DescribeHubRequest represents the DescribeHubRequest schema from the OpenAPI specification
type DescribeHubRequest struct {
}

// AwsEc2SecurityGroupPrefixListId represents the AwsEc2SecurityGroupPrefixListId schema from the OpenAPI specification
type AwsEc2SecurityGroupPrefixListId struct {
	Prefixlistid interface{} `json:"PrefixListId,omitempty"`
}

// DeleteInsightResponse represents the DeleteInsightResponse schema from the OpenAPI specification
type DeleteInsightResponse struct {
	Insightarn interface{} `json:"InsightArn"`
}

// AwsEc2LaunchTemplateDataNetworkInterfaceSetIpv6PrefixesDetails represents the AwsEc2LaunchTemplateDataNetworkInterfaceSetIpv6PrefixesDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataNetworkInterfaceSetIpv6PrefixesDetails struct {
	Ipv6prefix interface{} `json:"Ipv6Prefix,omitempty"`
}

// AwsApiGatewayV2ApiDetails represents the AwsApiGatewayV2ApiDetails schema from the OpenAPI specification
type AwsApiGatewayV2ApiDetails struct {
	Apiendpoint interface{} `json:"ApiEndpoint,omitempty"`
	Apiid interface{} `json:"ApiId,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Routeselectionexpression interface{} `json:"RouteSelectionExpression,omitempty"`
	Apikeyselectionexpression interface{} `json:"ApiKeySelectionExpression,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Corsconfiguration interface{} `json:"CorsConfiguration,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Protocoltype interface{} `json:"ProtocolType,omitempty"`
	Version interface{} `json:"Version,omitempty"`
}

// AwsEksClusterLoggingClusterLoggingDetails represents the AwsEksClusterLoggingClusterLoggingDetails schema from the OpenAPI specification
type AwsEksClusterLoggingClusterLoggingDetails struct {
	Types interface{} `json:"Types,omitempty"`
	Enabled interface{} `json:"Enabled,omitempty"`
}

// AwsGuardDutyDetectorDataSourcesKubernetesAuditLogsDetails represents the AwsGuardDutyDetectorDataSourcesKubernetesAuditLogsDetails schema from the OpenAPI specification
type AwsGuardDutyDetectorDataSourcesKubernetesAuditLogsDetails struct {
	Status interface{} `json:"Status,omitempty"`
}

// AwsEc2NetworkInterfaceSecurityGroup represents the AwsEc2NetworkInterfaceSecurityGroup schema from the OpenAPI specification
type AwsEc2NetworkInterfaceSecurityGroup struct {
	Groupid interface{} `json:"GroupId,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
}

// RuleGroupSourceListDetails represents the RuleGroupSourceListDetails schema from the OpenAPI specification
type RuleGroupSourceListDetails struct {
	Targettypes interface{} `json:"TargetTypes,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Generatedrulestype interface{} `json:"GeneratedRulesType,omitempty"`
}

// AwsIamRolePolicy represents the AwsIamRolePolicy schema from the OpenAPI specification
type AwsIamRolePolicy struct {
	Policyname interface{} `json:"PolicyName,omitempty"`
}

// ActionLocalIpDetails represents the ActionLocalIpDetails schema from the OpenAPI specification
type ActionLocalIpDetails struct {
	Ipaddressv4 interface{} `json:"IpAddressV4,omitempty"`
}

// AwsRdsDbSubnetGroupSubnet represents the AwsRdsDbSubnetGroupSubnet schema from the OpenAPI specification
type AwsRdsDbSubnetGroupSubnet struct {
	Subnetavailabilityzone interface{} `json:"SubnetAvailabilityZone,omitempty"`
	Subnetidentifier interface{} `json:"SubnetIdentifier,omitempty"`
	Subnetstatus interface{} `json:"SubnetStatus,omitempty"`
}

// BatchEnableStandardsResponse represents the BatchEnableStandardsResponse schema from the OpenAPI specification
type BatchEnableStandardsResponse struct {
	Standardssubscriptions interface{} `json:"StandardsSubscriptions,omitempty"`
}

// GetInvitationsCountResponse represents the GetInvitationsCountResponse schema from the OpenAPI specification
type GetInvitationsCountResponse struct {
	Invitationscount interface{} `json:"InvitationsCount,omitempty"`
}

// AwsEc2VpnConnectionOptionsDetails represents the AwsEc2VpnConnectionOptionsDetails schema from the OpenAPI specification
type AwsEc2VpnConnectionOptionsDetails struct {
	Staticroutesonly interface{} `json:"StaticRoutesOnly,omitempty"`
	Tunneloptions interface{} `json:"TunnelOptions,omitempty"`
}

// AwsEcrContainerImageDetails represents the AwsEcrContainerImageDetails schema from the OpenAPI specification
type AwsEcrContainerImageDetails struct {
	Imagetags interface{} `json:"ImageTags,omitempty"`
	Registryid interface{} `json:"RegistryId,omitempty"`
	Repositoryname interface{} `json:"RepositoryName,omitempty"`
	Architecture interface{} `json:"Architecture,omitempty"`
	Imagedigest interface{} `json:"ImageDigest,omitempty"`
	Imagepublishedat interface{} `json:"ImagePublishedAt,omitempty"`
}

// RuleGroupSourceStatelessRulesDetails represents the RuleGroupSourceStatelessRulesDetails schema from the OpenAPI specification
type RuleGroupSourceStatelessRulesDetails struct {
	Priority interface{} `json:"Priority,omitempty"`
	Ruledefinition interface{} `json:"RuleDefinition,omitempty"`
}

// AwsEc2NetworkInterfaceDetails represents the AwsEc2NetworkInterfaceDetails schema from the OpenAPI specification
type AwsEc2NetworkInterfaceDetails struct {
	Sourcedestcheck interface{} `json:"SourceDestCheck,omitempty"`
	Attachment interface{} `json:"Attachment,omitempty"`
	Ipv6addresses interface{} `json:"IpV6Addresses,omitempty"`
	Networkinterfaceid interface{} `json:"NetworkInterfaceId,omitempty"`
	Privateipaddresses interface{} `json:"PrivateIpAddresses,omitempty"`
	Publicdnsname interface{} `json:"PublicDnsName,omitempty"`
	Publicip interface{} `json:"PublicIp,omitempty"`
	Securitygroups interface{} `json:"SecurityGroups,omitempty"`
}

// EnableOrganizationAdminAccountRequest represents the EnableOrganizationAdminAccountRequest schema from the OpenAPI specification
type EnableOrganizationAdminAccountRequest struct {
	Adminaccountid interface{} `json:"AdminAccountId"`
}

// InsightResults represents the InsightResults schema from the OpenAPI specification
type InsightResults struct {
	Insightarn interface{} `json:"InsightArn"`
	Resultvalues interface{} `json:"ResultValues"`
	Groupbyattribute interface{} `json:"GroupByAttribute"`
}

// AwsRedshiftClusterResizeInfo represents the AwsRedshiftClusterResizeInfo schema from the OpenAPI specification
type AwsRedshiftClusterResizeInfo struct {
	Allowcancelresize interface{} `json:"AllowCancelResize,omitempty"`
	Resizetype interface{} `json:"ResizeType,omitempty"`
}

// SeverityUpdate represents the SeverityUpdate schema from the OpenAPI specification
type SeverityUpdate struct {
	Normalized interface{} `json:"Normalized,omitempty"`
	Product interface{} `json:"Product,omitempty"`
	Label interface{} `json:"Label,omitempty"`
}

// AwsElasticsearchDomainElasticsearchClusterConfigZoneAwarenessConfigDetails represents the AwsElasticsearchDomainElasticsearchClusterConfigZoneAwarenessConfigDetails schema from the OpenAPI specification
type AwsElasticsearchDomainElasticsearchClusterConfigZoneAwarenessConfigDetails struct {
	Availabilityzonecount interface{} `json:"AvailabilityZoneCount,omitempty"`
}

// AwsSecurityFinding represents the AwsSecurityFinding schema from the OpenAPI specification
type AwsSecurityFinding struct {
	Sample interface{} `json:"Sample,omitempty"`
	Threats interface{} `json:"Threats,omitempty"`
	Networkpath interface{} `json:"NetworkPath,omitempty"`
	Severity interface{} `json:"Severity,omitempty"`
	Updatedat interface{} `json:"UpdatedAt"`
	Workflow interface{} `json:"Workflow,omitempty"`
	Title interface{} `json:"Title"`
	Relatedfindings interface{} `json:"RelatedFindings,omitempty"`
	Sourceurl interface{} `json:"SourceUrl,omitempty"`
	Threatintelindicators interface{} `json:"ThreatIntelIndicators,omitempty"`
	Productfields interface{} `json:"ProductFields,omitempty"`
	Productarn interface{} `json:"ProductArn"`
	Region interface{} `json:"Region,omitempty"`
	Userdefinedfields interface{} `json:"UserDefinedFields,omitempty"`
	Firstobservedat interface{} `json:"FirstObservedAt,omitempty"`
	Network interface{} `json:"Network,omitempty"`
	Verificationstate interface{} `json:"VerificationState,omitempty"`
	Action interface{} `json:"Action,omitempty"`
	Companyname interface{} `json:"CompanyName,omitempty"`
	Confidence interface{} `json:"Confidence,omitempty"`
	Generatorid interface{} `json:"GeneratorId"`
	Resources interface{} `json:"Resources"`
	Createdat interface{} `json:"CreatedAt"`
	Compliance interface{} `json:"Compliance,omitempty"`
	Recordstate interface{} `json:"RecordState,omitempty"`
	Lastobservedat interface{} `json:"LastObservedAt,omitempty"`
	Id interface{} `json:"Id"`
	Malware interface{} `json:"Malware,omitempty"`
	Description interface{} `json:"Description"`
	Patchsummary interface{} `json:"PatchSummary,omitempty"`
	Process interface{} `json:"Process,omitempty"`
	Awsaccountid interface{} `json:"AwsAccountId"`
	Vulnerabilities interface{} `json:"Vulnerabilities,omitempty"`
	Remediation interface{} `json:"Remediation,omitempty"`
	Findingproviderfields interface{} `json:"FindingProviderFields,omitempty"`
	Types interface{} `json:"Types,omitempty"`
	Criticality interface{} `json:"Criticality,omitempty"`
	Note interface{} `json:"Note,omitempty"`
	Productname interface{} `json:"ProductName,omitempty"`
	Schemaversion interface{} `json:"SchemaVersion"`
	Workflowstate interface{} `json:"WorkflowState,omitempty"`
}

// AwsIamRoleDetails represents the AwsIamRoleDetails schema from the OpenAPI specification
type AwsIamRoleDetails struct {
	Rolepolicylist interface{} `json:"RolePolicyList,omitempty"`
	Createdate interface{} `json:"CreateDate,omitempty"`
	Instanceprofilelist interface{} `json:"InstanceProfileList,omitempty"`
	Maxsessionduration interface{} `json:"MaxSessionDuration,omitempty"`
	Path interface{} `json:"Path,omitempty"`
	Permissionsboundary AwsIamPermissionsBoundary `json:"PermissionsBoundary,omitempty"` // Information about the policy used to set the permissions boundary for an IAM principal.
	Rolename interface{} `json:"RoleName,omitempty"`
	Assumerolepolicydocument interface{} `json:"AssumeRolePolicyDocument,omitempty"`
	Attachedmanagedpolicies interface{} `json:"AttachedManagedPolicies,omitempty"`
	Roleid interface{} `json:"RoleId,omitempty"`
}

// CreateInsightResponse represents the CreateInsightResponse schema from the OpenAPI specification
type CreateInsightResponse struct {
	Insightarn interface{} `json:"InsightArn"`
}

// AwsEcsTaskDefinitionContainerDefinitionsDetails represents the AwsEcsTaskDefinitionContainerDefinitionsDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsDetails struct {
	Ulimits interface{} `json:"Ulimits,omitempty"`
	Environment interface{} `json:"Environment,omitempty"`
	Environmentfiles interface{} `json:"EnvironmentFiles,omitempty"`
	Repositorycredentials interface{} `json:"RepositoryCredentials,omitempty"`
	Dnssearchdomains interface{} `json:"DnsSearchDomains,omitempty"`
	Dockerlabels interface{} `json:"DockerLabels,omitempty"`
	Hostname interface{} `json:"Hostname,omitempty"`
	Image interface{} `json:"Image,omitempty"`
	Memoryreservation interface{} `json:"MemoryReservation,omitempty"`
	Interactive interface{} `json:"Interactive,omitempty"`
	Volumesfrom interface{} `json:"VolumesFrom,omitempty"`
	Disablenetworking interface{} `json:"DisableNetworking,omitempty"`
	Secrets interface{} `json:"Secrets,omitempty"`
	Linuxparameters interface{} `json:"LinuxParameters,omitempty"`
	Pseudoterminal interface{} `json:"PseudoTerminal,omitempty"`
	Command interface{} `json:"Command,omitempty"`
	Healthcheck interface{} `json:"HealthCheck,omitempty"`
	Logconfiguration interface{} `json:"LogConfiguration,omitempty"`
	User interface{} `json:"User,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Starttimeout interface{} `json:"StartTimeout,omitempty"`
	Dependson interface{} `json:"DependsOn,omitempty"`
	Entrypoint interface{} `json:"EntryPoint,omitempty"`
	Dnsservers interface{} `json:"DnsServers,omitempty"`
	Workingdirectory interface{} `json:"WorkingDirectory,omitempty"`
	Privileged interface{} `json:"Privileged,omitempty"`
	Systemcontrols interface{} `json:"SystemControls,omitempty"`
	Dockersecurityoptions interface{} `json:"DockerSecurityOptions,omitempty"`
	Resourcerequirements interface{} `json:"ResourceRequirements,omitempty"`
	Extrahosts interface{} `json:"ExtraHosts,omitempty"`
	Stoptimeout interface{} `json:"StopTimeout,omitempty"`
	Memory interface{} `json:"Memory,omitempty"`
	Mountpoints interface{} `json:"MountPoints,omitempty"`
	Readonlyrootfilesystem interface{} `json:"ReadonlyRootFilesystem,omitempty"`
	Firelensconfiguration interface{} `json:"FirelensConfiguration,omitempty"`
	Portmappings interface{} `json:"PortMappings,omitempty"`
	Links interface{} `json:"Links,omitempty"`
	Cpu interface{} `json:"Cpu,omitempty"`
	Essential interface{} `json:"Essential,omitempty"`
}

// DisassociateFromAdministratorAccountRequest represents the DisassociateFromAdministratorAccountRequest schema from the OpenAPI specification
type DisassociateFromAdministratorAccountRequest struct {
}

// AwsDynamoDbTableStreamSpecification represents the AwsDynamoDbTableStreamSpecification schema from the OpenAPI specification
type AwsDynamoDbTableStreamSpecification struct {
	Streamenabled interface{} `json:"StreamEnabled,omitempty"`
	Streamviewtype interface{} `json:"StreamViewType,omitempty"`
}

// AwsLambdaFunctionDeadLetterConfig represents the AwsLambdaFunctionDeadLetterConfig schema from the OpenAPI specification
type AwsLambdaFunctionDeadLetterConfig struct {
	Targetarn interface{} `json:"TargetArn,omitempty"`
}

// AwsWafv2VisibilityConfigDetails represents the AwsWafv2VisibilityConfigDetails schema from the OpenAPI specification
type AwsWafv2VisibilityConfigDetails struct {
	Cloudwatchmetricsenabled interface{} `json:"CloudWatchMetricsEnabled,omitempty"`
	Metricname interface{} `json:"MetricName,omitempty"`
	Sampledrequestsenabled interface{} `json:"SampledRequestsEnabled,omitempty"`
}

// AwsIamInstanceProfile represents the AwsIamInstanceProfile schema from the OpenAPI specification
type AwsIamInstanceProfile struct {
	Arn interface{} `json:"Arn,omitempty"`
	Createdate interface{} `json:"CreateDate,omitempty"`
	Instanceprofileid interface{} `json:"InstanceProfileId,omitempty"`
	Instanceprofilename interface{} `json:"InstanceProfileName,omitempty"`
	Path interface{} `json:"Path,omitempty"`
	Roles interface{} `json:"Roles,omitempty"`
}

// UpdateSecurityHubConfigurationResponse represents the UpdateSecurityHubConfigurationResponse schema from the OpenAPI specification
type UpdateSecurityHubConfigurationResponse struct {
}

// CreateMembersResponse represents the CreateMembersResponse schema from the OpenAPI specification
type CreateMembersResponse struct {
	Unprocessedaccounts interface{} `json:"UnprocessedAccounts,omitempty"`
}

// AwsAutoScalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateDetails represents the AwsAutoScalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateDetails schema from the OpenAPI specification
type AwsAutoScalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateDetails struct {
	Overrides interface{} `json:"Overrides,omitempty"`
	Launchtemplatespecification interface{} `json:"LaunchTemplateSpecification,omitempty"`
}

// AwsEcsTaskDefinitionInferenceAcceleratorsDetails represents the AwsEcsTaskDefinitionInferenceAcceleratorsDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionInferenceAcceleratorsDetails struct {
	Devicename interface{} `json:"DeviceName,omitempty"`
	Devicetype interface{} `json:"DeviceType,omitempty"`
}

// EnableSecurityHubResponse represents the EnableSecurityHubResponse schema from the OpenAPI specification
type EnableSecurityHubResponse struct {
}

// UpdateFindingAggregatorResponse represents the UpdateFindingAggregatorResponse schema from the OpenAPI specification
type UpdateFindingAggregatorResponse struct {
	Findingaggregationregion interface{} `json:"FindingAggregationRegion,omitempty"`
	Findingaggregatorarn interface{} `json:"FindingAggregatorArn,omitempty"`
	Regionlinkingmode interface{} `json:"RegionLinkingMode,omitempty"`
	Regions interface{} `json:"Regions,omitempty"`
}

// AwsCodeBuildProjectVpcConfig represents the AwsCodeBuildProjectVpcConfig schema from the OpenAPI specification
type AwsCodeBuildProjectVpcConfig struct {
	Vpcid interface{} `json:"VpcId,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
	Subnets interface{} `json:"Subnets,omitempty"`
}

// Vulnerability represents the Vulnerability schema from the OpenAPI specification
type Vulnerability struct {
	Fixavailable interface{} `json:"FixAvailable,omitempty"`
	Id interface{} `json:"Id"`
	Referenceurls interface{} `json:"ReferenceUrls,omitempty"`
	Relatedvulnerabilities interface{} `json:"RelatedVulnerabilities,omitempty"`
	Vendor interface{} `json:"Vendor,omitempty"`
	Vulnerablepackages interface{} `json:"VulnerablePackages,omitempty"`
	Cvss interface{} `json:"Cvss,omitempty"`
}

// AwsEcsTaskDefinitionVolumesDockerVolumeConfigurationDetails represents the AwsEcsTaskDefinitionVolumesDockerVolumeConfigurationDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionVolumesDockerVolumeConfigurationDetails struct {
	Labels interface{} `json:"Labels,omitempty"`
	Scope interface{} `json:"Scope,omitempty"`
	Autoprovision interface{} `json:"Autoprovision,omitempty"`
	Driver interface{} `json:"Driver,omitempty"`
	Driveropts interface{} `json:"DriverOpts,omitempty"`
}

// AwsSqsQueueDetails represents the AwsSqsQueueDetails schema from the OpenAPI specification
type AwsSqsQueueDetails struct {
	Kmsdatakeyreuseperiodseconds interface{} `json:"KmsDataKeyReusePeriodSeconds,omitempty"`
	Kmsmasterkeyid interface{} `json:"KmsMasterKeyId,omitempty"`
	Queuename interface{} `json:"QueueName,omitempty"`
	Deadlettertargetarn interface{} `json:"DeadLetterTargetArn,omitempty"`
}

// RuleGroupSourceStatefulRulesHeaderDetails represents the RuleGroupSourceStatefulRulesHeaderDetails schema from the OpenAPI specification
type RuleGroupSourceStatefulRulesHeaderDetails struct {
	Source interface{} `json:"Source,omitempty"`
	Sourceport interface{} `json:"SourcePort,omitempty"`
	Destination interface{} `json:"Destination,omitempty"`
	Destinationport interface{} `json:"DestinationPort,omitempty"`
	Direction interface{} `json:"Direction,omitempty"`
	Protocol interface{} `json:"Protocol,omitempty"`
}

// AwsCloudWatchAlarmDimensionsDetails represents the AwsCloudWatchAlarmDimensionsDetails schema from the OpenAPI specification
type AwsCloudWatchAlarmDimensionsDetails struct {
	Name interface{} `json:"Name,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// AwsEfsAccessPointDetails represents the AwsEfsAccessPointDetails schema from the OpenAPI specification
type AwsEfsAccessPointDetails struct {
	Arn interface{} `json:"Arn,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Filesystemid interface{} `json:"FileSystemId,omitempty"`
	Posixuser interface{} `json:"PosixUser,omitempty"`
	Rootdirectory interface{} `json:"RootDirectory,omitempty"`
	Accesspointid interface{} `json:"AccessPointId,omitempty"`
}

// AwsGuardDutyDetectorDataSourcesKubernetesDetails represents the AwsGuardDutyDetectorDataSourcesKubernetesDetails schema from the OpenAPI specification
type AwsGuardDutyDetectorDataSourcesKubernetesDetails struct {
	Auditlogs interface{} `json:"AuditLogs,omitempty"`
}

// AwsElasticsearchDomainVPCOptions represents the AwsElasticsearchDomainVPCOptions schema from the OpenAPI specification
type AwsElasticsearchDomainVPCOptions struct {
	Availabilityzones interface{} `json:"AvailabilityZones,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
	Subnetids interface{} `json:"SubnetIds,omitempty"`
	Vpcid interface{} `json:"VPCId,omitempty"`
}

// AwsDynamoDbTableProvisionedThroughputOverride represents the AwsDynamoDbTableProvisionedThroughputOverride schema from the OpenAPI specification
type AwsDynamoDbTableProvisionedThroughputOverride struct {
	Readcapacityunits interface{} `json:"ReadCapacityUnits,omitempty"`
}

// BatchUpdateFindingsResponse represents the BatchUpdateFindingsResponse schema from the OpenAPI specification
type BatchUpdateFindingsResponse struct {
	Unprocessedfindings interface{} `json:"UnprocessedFindings"`
	Processedfindings interface{} `json:"ProcessedFindings"`
}

// AwsEc2LaunchTemplateDataBlockDeviceMappingSetEbsDetails represents the AwsEc2LaunchTemplateDataBlockDeviceMappingSetEbsDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataBlockDeviceMappingSetEbsDetails struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Snapshotid interface{} `json:"SnapshotId,omitempty"`
	Throughput interface{} `json:"Throughput,omitempty"`
	Volumesize interface{} `json:"VolumeSize,omitempty"`
	Volumetype interface{} `json:"VolumeType,omitempty"`
	Deleteontermination interface{} `json:"DeleteOnTermination,omitempty"`
	Encrypted interface{} `json:"Encrypted,omitempty"`
	Iops interface{} `json:"Iops,omitempty"`
}

// AwsDynamoDbTableRestoreSummary represents the AwsDynamoDbTableRestoreSummary schema from the OpenAPI specification
type AwsDynamoDbTableRestoreSummary struct {
	Sourcebackuparn interface{} `json:"SourceBackupArn,omitempty"`
	Sourcetablearn interface{} `json:"SourceTableArn,omitempty"`
	Restoredatetime interface{} `json:"RestoreDateTime,omitempty"`
	Restoreinprogress interface{} `json:"RestoreInProgress,omitempty"`
}

// Cell represents the Cell schema from the OpenAPI specification
type Cell struct {
	Column interface{} `json:"Column,omitempty"`
	Columnname interface{} `json:"ColumnName,omitempty"`
	Row interface{} `json:"Row,omitempty"`
	Cellreference interface{} `json:"CellReference,omitempty"`
}

// AwsEksClusterResourcesVpcConfigDetails represents the AwsEksClusterResourcesVpcConfigDetails schema from the OpenAPI specification
type AwsEksClusterResourcesVpcConfigDetails struct {
	Endpointpublicaccess interface{} `json:"EndpointPublicAccess,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
	Subnetids interface{} `json:"SubnetIds,omitempty"`
}

// ListAutomationRulesRequest represents the ListAutomationRulesRequest schema from the OpenAPI specification
type ListAutomationRulesRequest struct {
}

// AwsWafv2RuleGroupDetails represents the AwsWafv2RuleGroupDetails schema from the OpenAPI specification
type AwsWafv2RuleGroupDetails struct {
	Capacity interface{} `json:"Capacity,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Rules interface{} `json:"Rules,omitempty"`
	Scope interface{} `json:"Scope,omitempty"`
	Visibilityconfig interface{} `json:"VisibilityConfig,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
}

// AwsApiGatewayCanarySettings represents the AwsApiGatewayCanarySettings schema from the OpenAPI specification
type AwsApiGatewayCanarySettings struct {
	Stagevariableoverrides interface{} `json:"StageVariableOverrides,omitempty"`
	Usestagecache interface{} `json:"UseStageCache,omitempty"`
	Deploymentid interface{} `json:"DeploymentId,omitempty"`
	Percenttraffic interface{} `json:"PercentTraffic,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// InviteMembersRequest represents the InviteMembersRequest schema from the OpenAPI specification
type InviteMembersRequest struct {
	Accountids interface{} `json:"AccountIds"`
}

// AwsLambdaFunctionTracingConfig represents the AwsLambdaFunctionTracingConfig schema from the OpenAPI specification
type AwsLambdaFunctionTracingConfig struct {
	Mode interface{} `json:"Mode,omitempty"`
}

// AwsEc2LaunchTemplateDataElasticInferenceAcceleratorSetDetails represents the AwsEc2LaunchTemplateDataElasticInferenceAcceleratorSetDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataElasticInferenceAcceleratorSetDetails struct {
	Count interface{} `json:"Count,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// AwsRdsDbProcessorFeature represents the AwsRdsDbProcessorFeature schema from the OpenAPI specification
type AwsRdsDbProcessorFeature struct {
	Value interface{} `json:"Value,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsSystemControlsDetails represents the AwsEcsTaskDefinitionContainerDefinitionsSystemControlsDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsSystemControlsDetails struct {
	Value interface{} `json:"Value,omitempty"`
	Namespace interface{} `json:"Namespace,omitempty"`
}

// UpdateInsightRequest represents the UpdateInsightRequest schema from the OpenAPI specification
type UpdateInsightRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Groupbyattribute interface{} `json:"GroupByAttribute,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// AwsEc2LaunchTemplateDataElasticGpuSpecificationSetDetails represents the AwsEc2LaunchTemplateDataElasticGpuSpecificationSetDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataElasticGpuSpecificationSetDetails struct {
	TypeField interface{} `json:"Type,omitempty"`
}

// AwsGuardDutyDetectorDataSourcesMalwareProtectionDetails represents the AwsGuardDutyDetectorDataSourcesMalwareProtectionDetails schema from the OpenAPI specification
type AwsGuardDutyDetectorDataSourcesMalwareProtectionDetails struct {
	Scanec2instancewithfindings interface{} `json:"ScanEc2InstanceWithFindings,omitempty"`
	Servicerole interface{} `json:"ServiceRole,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsUlimitsDetails represents the AwsEcsTaskDefinitionContainerDefinitionsUlimitsDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsUlimitsDetails struct {
	Hardlimit interface{} `json:"HardLimit,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Softlimit interface{} `json:"SoftLimit,omitempty"`
}

// StandardsSubscriptionRequest represents the StandardsSubscriptionRequest schema from the OpenAPI specification
type StandardsSubscriptionRequest struct {
	Standardsarn interface{} `json:"StandardsArn"`
	Standardsinput interface{} `json:"StandardsInput,omitempty"`
}

// UpdateOrganizationConfigurationResponse represents the UpdateOrganizationConfigurationResponse schema from the OpenAPI specification
type UpdateOrganizationConfigurationResponse struct {
}

// AwsEc2NetworkAclEntry represents the AwsEc2NetworkAclEntry schema from the OpenAPI specification
type AwsEc2NetworkAclEntry struct {
	Cidrblock interface{} `json:"CidrBlock,omitempty"`
	Egress interface{} `json:"Egress,omitempty"`
	Icmptypecode interface{} `json:"IcmpTypeCode,omitempty"`
	Ipv6cidrblock interface{} `json:"Ipv6CidrBlock,omitempty"`
	Portrange interface{} `json:"PortRange,omitempty"`
	Protocol interface{} `json:"Protocol,omitempty"`
	Ruleaction interface{} `json:"RuleAction,omitempty"`
	Rulenumber interface{} `json:"RuleNumber,omitempty"`
}

// FirewallPolicyStatelessCustomActionsDetails represents the FirewallPolicyStatelessCustomActionsDetails schema from the OpenAPI specification
type FirewallPolicyStatelessCustomActionsDetails struct {
	Actiondefinition interface{} `json:"ActionDefinition,omitempty"`
	Actionname interface{} `json:"ActionName,omitempty"`
}

// BatchDeleteAutomationRulesRequest represents the BatchDeleteAutomationRulesRequest schema from the OpenAPI specification
type BatchDeleteAutomationRulesRequest struct {
	Automationrulesarns interface{} `json:"AutomationRulesArns"`
}

// AwsAmazonMqBrokerLogsPendingDetails represents the AwsAmazonMqBrokerLogsPendingDetails schema from the OpenAPI specification
type AwsAmazonMqBrokerLogsPendingDetails struct {
	Audit interface{} `json:"Audit,omitempty"`
	General interface{} `json:"General,omitempty"`
}

// RuleGroupSourceStatelessRuleMatchAttributesDestinations represents the RuleGroupSourceStatelessRuleMatchAttributesDestinations schema from the OpenAPI specification
type RuleGroupSourceStatelessRuleMatchAttributesDestinations struct {
	Addressdefinition interface{} `json:"AddressDefinition,omitempty"`
}

// AwsElbLoadBalancerListenerDescription represents the AwsElbLoadBalancerListenerDescription schema from the OpenAPI specification
type AwsElbLoadBalancerListenerDescription struct {
	Listener interface{} `json:"Listener,omitempty"`
	Policynames interface{} `json:"PolicyNames,omitempty"`
}

// AwsEc2VpnConnectionRoutesDetails represents the AwsEc2VpnConnectionRoutesDetails schema from the OpenAPI specification
type AwsEc2VpnConnectionRoutesDetails struct {
	Destinationcidrblock interface{} `json:"DestinationCidrBlock,omitempty"`
	State interface{} `json:"State,omitempty"`
}

// Note represents the Note schema from the OpenAPI specification
type Note struct {
	Text interface{} `json:"Text"`
	Updatedat interface{} `json:"UpdatedAt"`
	Updatedby interface{} `json:"UpdatedBy"`
}

// UpdateActionTargetRequest represents the UpdateActionTargetRequest schema from the OpenAPI specification
type UpdateActionTargetRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// AwsAppSyncGraphQlApiDetails represents the AwsAppSyncGraphQlApiDetails schema from the OpenAPI specification
type AwsAppSyncGraphQlApiDetails struct {
	Id interface{} `json:"Id,omitempty"`
	Logconfig interface{} `json:"LogConfig,omitempty"`
	Wafwebaclarn interface{} `json:"WafWebAclArn,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Lambdaauthorizerconfig interface{} `json:"LambdaAuthorizerConfig,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Apiid interface{} `json:"ApiId,omitempty"`
	Authenticationtype interface{} `json:"AuthenticationType,omitempty"`
	Xrayenabled interface{} `json:"XrayEnabled,omitempty"`
	Openidconnectconfig interface{} `json:"OpenIdConnectConfig,omitempty"`
	Userpoolconfig interface{} `json:"UserPoolConfig,omitempty"`
	Additionalauthenticationproviders interface{} `json:"AdditionalAuthenticationProviders,omitempty"`
}

// AwsRdsDbSecurityGroupEc2SecurityGroup represents the AwsRdsDbSecurityGroupEc2SecurityGroup schema from the OpenAPI specification
type AwsRdsDbSecurityGroupEc2SecurityGroup struct {
	Ec2securitygroupid interface{} `json:"Ec2SecurityGroupId,omitempty"`
	Ec2securitygroupname interface{} `json:"Ec2SecurityGroupName,omitempty"`
	Ec2securitygroupownerid interface{} `json:"Ec2SecurityGroupOwnerId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// PortRange represents the PortRange schema from the OpenAPI specification
type PortRange struct {
	Begin interface{} `json:"Begin,omitempty"`
	End interface{} `json:"End,omitempty"`
}

// AwsS3BucketBucketLifecycleConfigurationRulesFilterPredicateOperandsTagDetails represents the AwsS3BucketBucketLifecycleConfigurationRulesFilterPredicateOperandsTagDetails schema from the OpenAPI specification
type AwsS3BucketBucketLifecycleConfigurationRulesFilterPredicateOperandsTagDetails struct {
	Value interface{} `json:"Value,omitempty"`
	Key interface{} `json:"Key,omitempty"`
}

// BatchDeleteAutomationRulesResponse represents the BatchDeleteAutomationRulesResponse schema from the OpenAPI specification
type BatchDeleteAutomationRulesResponse struct {
	Processedautomationrules interface{} `json:"ProcessedAutomationRules,omitempty"`
	Unprocessedautomationrules interface{} `json:"UnprocessedAutomationRules,omitempty"`
}

// RuleGroupVariables represents the RuleGroupVariables schema from the OpenAPI specification
type RuleGroupVariables struct {
	Ipsets interface{} `json:"IpSets,omitempty"`
	Portsets interface{} `json:"PortSets,omitempty"`
}

// AwsBackupRecoveryPointLifecycleDetails represents the AwsBackupRecoveryPointLifecycleDetails schema from the OpenAPI specification
type AwsBackupRecoveryPointLifecycleDetails struct {
	Deleteafterdays interface{} `json:"DeleteAfterDays,omitempty"`
	Movetocoldstorageafterdays interface{} `json:"MoveToColdStorageAfterDays,omitempty"`
}

// FindingHistoryUpdateSource represents the FindingHistoryUpdateSource schema from the OpenAPI specification
type FindingHistoryUpdateSource struct {
	Identity interface{} `json:"Identity,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// AwsAmazonMqBrokerEncryptionOptionsDetails represents the AwsAmazonMqBrokerEncryptionOptionsDetails schema from the OpenAPI specification
type AwsAmazonMqBrokerEncryptionOptionsDetails struct {
	Useawsownedkey interface{} `json:"UseAwsOwnedKey,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
}

// AwsRdsDbSnapshotDetails represents the AwsRdsDbSnapshotDetails schema from the OpenAPI specification
type AwsRdsDbSnapshotDetails struct {
	Iamdatabaseauthenticationenabled interface{} `json:"IamDatabaseAuthenticationEnabled,omitempty"`
	Iops interface{} `json:"Iops,omitempty"`
	Snapshottype interface{} `json:"SnapshotType,omitempty"`
	Allocatedstorage interface{} `json:"AllocatedStorage,omitempty"`
	Engine interface{} `json:"Engine,omitempty"`
	Masterusername interface{} `json:"MasterUsername,omitempty"`
	Availabilityzone interface{} `json:"AvailabilityZone,omitempty"`
	Encrypted interface{} `json:"Encrypted,omitempty"`
	Sourceregion interface{} `json:"SourceRegion,omitempty"`
	Engineversion interface{} `json:"EngineVersion,omitempty"`
	Percentprogress interface{} `json:"PercentProgress,omitempty"`
	Licensemodel interface{} `json:"LicenseModel,omitempty"`
	Sourcedbsnapshotidentifier interface{} `json:"SourceDbSnapshotIdentifier,omitempty"`
	Dbiresourceid interface{} `json:"DbiResourceId,omitempty"`
	Timezone interface{} `json:"Timezone,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Snapshotcreatetime interface{} `json:"SnapshotCreateTime,omitempty"`
	Tdecredentialarn interface{} `json:"TdeCredentialArn,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Processorfeatures interface{} `json:"ProcessorFeatures,omitempty"`
	Optiongroupname interface{} `json:"OptionGroupName,omitempty"`
	Dbsnapshotidentifier interface{} `json:"DbSnapshotIdentifier,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Instancecreatetime interface{} `json:"InstanceCreateTime,omitempty"`
	Port interface{} `json:"Port,omitempty"`
	Storagetype interface{} `json:"StorageType,omitempty"`
	Dbinstanceidentifier interface{} `json:"DbInstanceIdentifier,omitempty"`
}

// AwsElasticsearchDomainEncryptionAtRestOptions represents the AwsElasticsearchDomainEncryptionAtRestOptions schema from the OpenAPI specification
type AwsElasticsearchDomainEncryptionAtRestOptions struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
}

// RuleGroupSource represents the RuleGroupSource schema from the OpenAPI specification
type RuleGroupSource struct {
	Statefulrules interface{} `json:"StatefulRules,omitempty"`
	Statelessrulesandcustomactions interface{} `json:"StatelessRulesAndCustomActions,omitempty"`
	Rulessourcelist interface{} `json:"RulesSourceList,omitempty"`
	Rulesstring interface{} `json:"RulesString,omitempty"`
}

// AwsRdsDbStatusInfo represents the AwsRdsDbStatusInfo schema from the OpenAPI specification
type AwsRdsDbStatusInfo struct {
	Normal interface{} `json:"Normal,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Statustype interface{} `json:"StatusType,omitempty"`
	Message interface{} `json:"Message,omitempty"`
}

// AwsGuardDutyDetectorDetails represents the AwsGuardDutyDetectorDetails schema from the OpenAPI specification
type AwsGuardDutyDetectorDetails struct {
	Status interface{} `json:"Status,omitempty"`
	Datasources interface{} `json:"DataSources,omitempty"`
	Features interface{} `json:"Features,omitempty"`
	Findingpublishingfrequency interface{} `json:"FindingPublishingFrequency,omitempty"`
	Servicerole interface{} `json:"ServiceRole,omitempty"`
}

// DisassociateMembersRequest represents the DisassociateMembersRequest schema from the OpenAPI specification
type DisassociateMembersRequest struct {
	Accountids interface{} `json:"AccountIds"`
}

// AwsEc2LaunchTemplateDataBlockDeviceMappingSetDetails represents the AwsEc2LaunchTemplateDataBlockDeviceMappingSetDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataBlockDeviceMappingSetDetails struct {
	Ebs interface{} `json:"Ebs,omitempty"`
	Nodevice interface{} `json:"NoDevice,omitempty"`
	Virtualname interface{} `json:"VirtualName,omitempty"`
	Devicename interface{} `json:"DeviceName,omitempty"`
}

// DescribeOrganizationConfigurationRequest represents the DescribeOrganizationConfigurationRequest schema from the OpenAPI specification
type DescribeOrganizationConfigurationRequest struct {
}

// WafAction represents the WafAction schema from the OpenAPI specification
type WafAction struct {
	TypeField interface{} `json:"Type,omitempty"`
}

// PortProbeDetail represents the PortProbeDetail schema from the OpenAPI specification
type PortProbeDetail struct {
	Localipdetails interface{} `json:"LocalIpDetails,omitempty"`
	Localportdetails interface{} `json:"LocalPortDetails,omitempty"`
	Remoteipdetails interface{} `json:"RemoteIpDetails,omitempty"`
}

// AwsWafv2WebAclDetails represents the AwsWafv2WebAclDetails schema from the OpenAPI specification
type AwsWafv2WebAclDetails struct {
	Captchaconfig interface{} `json:"CaptchaConfig,omitempty"`
	Rules interface{} `json:"Rules,omitempty"`
	Capacity interface{} `json:"Capacity,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Defaultaction interface{} `json:"DefaultAction,omitempty"`
	Managedbyfirewallmanager interface{} `json:"ManagedbyFirewallManager,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Visibilityconfig interface{} `json:"VisibilityConfig,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
}

// AwsEc2LaunchTemplateDataDetails represents the AwsEc2LaunchTemplateDataDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataDetails struct {
	Enclaveoptions interface{} `json:"EnclaveOptions,omitempty"`
	Instanceinitiatedshutdownbehavior interface{} `json:"InstanceInitiatedShutdownBehavior,omitempty"`
	Cpuoptions interface{} `json:"CpuOptions,omitempty"`
	Userdata interface{} `json:"UserData,omitempty"`
	Securitygroupidset interface{} `json:"SecurityGroupIdSet,omitempty"`
	Blockdevicemappingset interface{} `json:"BlockDeviceMappingSet,omitempty"`
	Licenseset interface{} `json:"LicenseSet,omitempty"`
	Ebsoptimized interface{} `json:"EbsOptimized,omitempty"`
	Ramdiskid interface{} `json:"RamDiskId,omitempty"`
	Creditspecification interface{} `json:"CreditSpecification,omitempty"`
	Disableapitermination interface{} `json:"DisableApiTermination,omitempty"`
	Instancemarketoptions interface{} `json:"InstanceMarketOptions,omitempty"`
	Imageid interface{} `json:"ImageId,omitempty"`
	Disableapistop interface{} `json:"DisableApiStop,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Networkinterfaceset interface{} `json:"NetworkInterfaceSet,omitempty"`
	Hibernationoptions interface{} `json:"HibernationOptions,omitempty"`
	Instancerequirements interface{} `json:"InstanceRequirements,omitempty"`
	Keyname interface{} `json:"KeyName,omitempty"`
	Monitoring interface{} `json:"Monitoring,omitempty"`
	Elasticgpuspecificationset interface{} `json:"ElasticGpuSpecificationSet,omitempty"`
	Iaminstanceprofile interface{} `json:"IamInstanceProfile,omitempty"`
	Maintenanceoptions interface{} `json:"MaintenanceOptions,omitempty"`
	Privatednsnameoptions interface{} `json:"PrivateDnsNameOptions,omitempty"`
	Elasticinferenceacceleratorset interface{} `json:"ElasticInferenceAcceleratorSet,omitempty"`
	Placement interface{} `json:"Placement,omitempty"`
	Metadataoptions interface{} `json:"MetadataOptions,omitempty"`
	Securitygroupset interface{} `json:"SecurityGroupSet,omitempty"`
	Capacityreservationspecification interface{} `json:"CapacityReservationSpecification,omitempty"`
	Kernelid interface{} `json:"KernelId,omitempty"`
}

// DisableSecurityHubRequest represents the DisableSecurityHubRequest schema from the OpenAPI specification
type DisableSecurityHubRequest struct {
}

// GetMasterAccountRequest represents the GetMasterAccountRequest schema from the OpenAPI specification
type GetMasterAccountRequest struct {
}

// AwsCertificateManagerCertificateExtendedKeyUsage represents the AwsCertificateManagerCertificateExtendedKeyUsage schema from the OpenAPI specification
type AwsCertificateManagerCertificateExtendedKeyUsage struct {
	Name interface{} `json:"Name,omitempty"`
	Oid interface{} `json:"OId,omitempty"`
}

// AwsCertificateManagerCertificateDomainValidationOption represents the AwsCertificateManagerCertificateDomainValidationOption schema from the OpenAPI specification
type AwsCertificateManagerCertificateDomainValidationOption struct {
	Validationstatus interface{} `json:"ValidationStatus,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Resourcerecord interface{} `json:"ResourceRecord,omitempty"`
	Validationdomain interface{} `json:"ValidationDomain,omitempty"`
	Validationemails interface{} `json:"ValidationEmails,omitempty"`
	Validationmethod interface{} `json:"ValidationMethod,omitempty"`
}

// CreateInsightRequest represents the CreateInsightRequest schema from the OpenAPI specification
type CreateInsightRequest struct {
	Filters interface{} `json:"Filters"`
	Groupbyattribute interface{} `json:"GroupByAttribute"`
	Name interface{} `json:"Name"`
}

// AwsEc2SecurityGroupUserIdGroupPair represents the AwsEc2SecurityGroupUserIdGroupPair schema from the OpenAPI specification
type AwsEc2SecurityGroupUserIdGroupPair struct {
	Userid interface{} `json:"UserId,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Vpcpeeringconnectionid interface{} `json:"VpcPeeringConnectionId,omitempty"`
	Groupid interface{} `json:"GroupId,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
	Peeringstatus interface{} `json:"PeeringStatus,omitempty"`
}

// AwsCloudFrontDistributionViewerCertificate represents the AwsCloudFrontDistributionViewerCertificate schema from the OpenAPI specification
type AwsCloudFrontDistributionViewerCertificate struct {
	Minimumprotocolversion interface{} `json:"MinimumProtocolVersion,omitempty"`
	Sslsupportmethod interface{} `json:"SslSupportMethod,omitempty"`
	Acmcertificatearn interface{} `json:"AcmCertificateArn,omitempty"`
	Certificate interface{} `json:"Certificate,omitempty"`
	Certificatesource interface{} `json:"CertificateSource,omitempty"`
	Cloudfrontdefaultcertificate interface{} `json:"CloudFrontDefaultCertificate,omitempty"`
	Iamcertificateid interface{} `json:"IamCertificateId,omitempty"`
}

// City represents the City schema from the OpenAPI specification
type City struct {
	Cityname interface{} `json:"CityName,omitempty"`
}

// AwsOpenSearchServiceDomainDomainEndpointOptionsDetails represents the AwsOpenSearchServiceDomainDomainEndpointOptionsDetails schema from the OpenAPI specification
type AwsOpenSearchServiceDomainDomainEndpointOptionsDetails struct {
	Customendpoint interface{} `json:"CustomEndpoint,omitempty"`
	Customendpointcertificatearn interface{} `json:"CustomEndpointCertificateArn,omitempty"`
	Customendpointenabled interface{} `json:"CustomEndpointEnabled,omitempty"`
	Enforcehttps interface{} `json:"EnforceHTTPS,omitempty"`
	Tlssecuritypolicy interface{} `json:"TLSSecurityPolicy,omitempty"`
}

// CreateMembersRequest represents the CreateMembersRequest schema from the OpenAPI specification
type CreateMembersRequest struct {
	Accountdetails interface{} `json:"AccountDetails"`
}

// BatchUpdateStandardsControlAssociationsResponse represents the BatchUpdateStandardsControlAssociationsResponse schema from the OpenAPI specification
type BatchUpdateStandardsControlAssociationsResponse struct {
	Unprocessedassociationupdates interface{} `json:"UnprocessedAssociationUpdates,omitempty"`
}

// AwsEc2InstanceNetworkInterfacesDetails represents the AwsEc2InstanceNetworkInterfacesDetails schema from the OpenAPI specification
type AwsEc2InstanceNetworkInterfacesDetails struct {
	Networkinterfaceid interface{} `json:"NetworkInterfaceId,omitempty"`
}

// AwsS3BucketObjectLockConfiguration represents the AwsS3BucketObjectLockConfiguration schema from the OpenAPI specification
type AwsS3BucketObjectLockConfiguration struct {
	Objectlockenabled interface{} `json:"ObjectLockEnabled,omitempty"`
	Rule interface{} `json:"Rule,omitempty"`
}

// Ipv6CidrBlockAssociation represents the Ipv6CidrBlockAssociation schema from the OpenAPI specification
type Ipv6CidrBlockAssociation struct {
	Ipv6cidrblock interface{} `json:"Ipv6CidrBlock,omitempty"`
	Associationid interface{} `json:"AssociationId,omitempty"`
	Cidrblockstate interface{} `json:"CidrBlockState,omitempty"`
}

// AwsWafRegionalRuleGroupRulesDetails represents the AwsWafRegionalRuleGroupRulesDetails schema from the OpenAPI specification
type AwsWafRegionalRuleGroupRulesDetails struct {
	Ruleid interface{} `json:"RuleId,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Action interface{} `json:"Action,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
}

// AwsApiCallActionDomainDetails represents the AwsApiCallActionDomainDetails schema from the OpenAPI specification
type AwsApiCallActionDomainDetails struct {
	Domain interface{} `json:"Domain,omitempty"`
}

// AwsAmazonMqBrokerMaintenanceWindowStartTimeDetails represents the AwsAmazonMqBrokerMaintenanceWindowStartTimeDetails schema from the OpenAPI specification
type AwsAmazonMqBrokerMaintenanceWindowStartTimeDetails struct {
	Dayofweek interface{} `json:"DayOfWeek,omitempty"`
	Timeofday interface{} `json:"TimeOfDay,omitempty"`
	Timezone interface{} `json:"TimeZone,omitempty"`
}

// AwsIamInstanceProfileRole represents the AwsIamInstanceProfileRole schema from the OpenAPI specification
type AwsIamInstanceProfileRole struct {
	Path interface{} `json:"Path,omitempty"`
	Roleid interface{} `json:"RoleId,omitempty"`
	Rolename interface{} `json:"RoleName,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Assumerolepolicydocument interface{} `json:"AssumeRolePolicyDocument,omitempty"`
	Createdate interface{} `json:"CreateDate,omitempty"`
}

// ImportFindingsError represents the ImportFindingsError schema from the OpenAPI specification
type ImportFindingsError struct {
	Errorcode interface{} `json:"ErrorCode"`
	Errormessage interface{} `json:"ErrorMessage"`
	Id interface{} `json:"Id"`
}

// DescribeProductsResponse represents the DescribeProductsResponse schema from the OpenAPI specification
type DescribeProductsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Products interface{} `json:"Products"`
}

// AwsWafRuleGroupDetails represents the AwsWafRuleGroupDetails schema from the OpenAPI specification
type AwsWafRuleGroupDetails struct {
	Rulegroupid interface{} `json:"RuleGroupId,omitempty"`
	Rules interface{} `json:"Rules,omitempty"`
	Metricname interface{} `json:"MetricName,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// Remediation represents the Remediation schema from the OpenAPI specification
type Remediation struct {
	Recommendation interface{} `json:"Recommendation,omitempty"`
}

// AwsWafv2ActionAllowDetails represents the AwsWafv2ActionAllowDetails schema from the OpenAPI specification
type AwsWafv2ActionAllowDetails struct {
	Customrequesthandling interface{} `json:"CustomRequestHandling,omitempty"`
}

// Range represents the Range schema from the OpenAPI specification
type Range struct {
	End interface{} `json:"End,omitempty"`
	Start interface{} `json:"Start,omitempty"`
	Startcolumn interface{} `json:"StartColumn,omitempty"`
}

// NetworkPathComponentDetails represents the NetworkPathComponentDetails schema from the OpenAPI specification
type NetworkPathComponentDetails struct {
	Portranges interface{} `json:"PortRanges,omitempty"`
	Address interface{} `json:"Address,omitempty"`
}

// AwsDynamoDbTableGlobalSecondaryIndex represents the AwsDynamoDbTableGlobalSecondaryIndex schema from the OpenAPI specification
type AwsDynamoDbTableGlobalSecondaryIndex struct {
	Indexarn interface{} `json:"IndexArn,omitempty"`
	Indexname interface{} `json:"IndexName,omitempty"`
	Keyschema interface{} `json:"KeySchema,omitempty"`
	Projection interface{} `json:"Projection,omitempty"`
	Provisionedthroughput interface{} `json:"ProvisionedThroughput,omitempty"`
	Indexstatus interface{} `json:"IndexStatus,omitempty"`
	Itemcount interface{} `json:"ItemCount,omitempty"`
	Backfilling interface{} `json:"Backfilling,omitempty"`
	Indexsizebytes interface{} `json:"IndexSizeBytes,omitempty"`
}

// UpdateInsightResponse represents the UpdateInsightResponse schema from the OpenAPI specification
type UpdateInsightResponse struct {
}

// AwsS3BucketNotificationConfigurationDetail represents the AwsS3BucketNotificationConfigurationDetail schema from the OpenAPI specification
type AwsS3BucketNotificationConfigurationDetail struct {
	Events interface{} `json:"Events,omitempty"`
	Filter interface{} `json:"Filter,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Destination interface{} `json:"Destination,omitempty"`
}

// AwsCloudFrontDistributionOriginGroupFailoverStatusCodes represents the AwsCloudFrontDistributionOriginGroupFailoverStatusCodes schema from the OpenAPI specification
type AwsCloudFrontDistributionOriginGroupFailoverStatusCodes struct {
	Items interface{} `json:"Items,omitempty"`
	Quantity interface{} `json:"Quantity,omitempty"`
}

// AwsStepFunctionStateMachineTracingConfigurationDetails represents the AwsStepFunctionStateMachineTracingConfigurationDetails schema from the OpenAPI specification
type AwsStepFunctionStateMachineTracingConfigurationDetails struct {
	Enabled interface{} `json:"Enabled,omitempty"`
}

// GetInsightResultsRequest represents the GetInsightResultsRequest schema from the OpenAPI specification
type GetInsightResultsRequest struct {
}

// DeleteMembersRequest represents the DeleteMembersRequest schema from the OpenAPI specification
type DeleteMembersRequest struct {
	Accountids interface{} `json:"AccountIds"`
}

// AwsCloudFrontDistributionCacheBehavior represents the AwsCloudFrontDistributionCacheBehavior schema from the OpenAPI specification
type AwsCloudFrontDistributionCacheBehavior struct {
	Viewerprotocolpolicy interface{} `json:"ViewerProtocolPolicy,omitempty"`
}

// AwsS3BucketNotificationConfigurationS3KeyFilterRule represents the AwsS3BucketNotificationConfigurationS3KeyFilterRule schema from the OpenAPI specification
type AwsS3BucketNotificationConfigurationS3KeyFilterRule struct {
	Name interface{} `json:"Name,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// AwsRedshiftClusterVpcSecurityGroup represents the AwsRedshiftClusterVpcSecurityGroup schema from the OpenAPI specification
type AwsRedshiftClusterVpcSecurityGroup struct {
	Vpcsecuritygroupid interface{} `json:"VpcSecurityGroupId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// AwsEcsClusterDefaultCapacityProviderStrategyDetails represents the AwsEcsClusterDefaultCapacityProviderStrategyDetails schema from the OpenAPI specification
type AwsEcsClusterDefaultCapacityProviderStrategyDetails struct {
	Base interface{} `json:"Base,omitempty"`
	Capacityprovider interface{} `json:"CapacityProvider,omitempty"`
	Weight interface{} `json:"Weight,omitempty"`
}

// SensitiveDataDetections represents the SensitiveDataDetections schema from the OpenAPI specification
type SensitiveDataDetections struct {
	TypeField interface{} `json:"Type,omitempty"`
	Count interface{} `json:"Count,omitempty"`
	Occurrences interface{} `json:"Occurrences,omitempty"`
}

// AwsAmazonMqBrokerUsersDetails represents the AwsAmazonMqBrokerUsersDetails schema from the OpenAPI specification
type AwsAmazonMqBrokerUsersDetails struct {
	Pendingchange interface{} `json:"PendingChange,omitempty"`
	Username interface{} `json:"Username,omitempty"`
}

// NoteUpdate represents the NoteUpdate schema from the OpenAPI specification
type NoteUpdate struct {
	Text interface{} `json:"Text"`
	Updatedby interface{} `json:"UpdatedBy"`
}

// WorkflowUpdate represents the WorkflowUpdate schema from the OpenAPI specification
type WorkflowUpdate struct {
	Status interface{} `json:"Status,omitempty"`
}

// GetFindingHistoryResponse represents the GetFindingHistoryResponse schema from the OpenAPI specification
type GetFindingHistoryResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Records interface{} `json:"Records,omitempty"`
}

// AwsEcrRepositoryImageScanningConfigurationDetails represents the AwsEcrRepositoryImageScanningConfigurationDetails schema from the OpenAPI specification
type AwsEcrRepositoryImageScanningConfigurationDetails struct {
	Scanonpush interface{} `json:"ScanOnPush,omitempty"`
}

// AwsElbLoadBalancerListener represents the AwsElbLoadBalancerListener schema from the OpenAPI specification
type AwsElbLoadBalancerListener struct {
	Sslcertificateid interface{} `json:"SslCertificateId,omitempty"`
	Instanceport interface{} `json:"InstancePort,omitempty"`
	Instanceprotocol interface{} `json:"InstanceProtocol,omitempty"`
	Loadbalancerport interface{} `json:"LoadBalancerPort,omitempty"`
	Protocol interface{} `json:"Protocol,omitempty"`
}

// AwsRedshiftClusterClusterSecurityGroup represents the AwsRedshiftClusterClusterSecurityGroup schema from the OpenAPI specification
type AwsRedshiftClusterClusterSecurityGroup struct {
	Status interface{} `json:"Status,omitempty"`
	Clustersecuritygroupname interface{} `json:"ClusterSecurityGroupName,omitempty"`
}

// AwsElbLoadBalancerInstance represents the AwsElbLoadBalancerInstance schema from the OpenAPI specification
type AwsElbLoadBalancerInstance struct {
	Instanceid interface{} `json:"InstanceId,omitempty"`
}

// AwsWafRegionalRuleDetails represents the AwsWafRegionalRuleDetails schema from the OpenAPI specification
type AwsWafRegionalRuleDetails struct {
	Metricname interface{} `json:"MetricName,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Predicatelist interface{} `json:"PredicateList,omitempty"`
	Ruleid interface{} `json:"RuleId,omitempty"`
}

// NetworkPathComponent represents the NetworkPathComponent schema from the OpenAPI specification
type NetworkPathComponent struct {
	Componentid interface{} `json:"ComponentId,omitempty"`
	Componenttype interface{} `json:"ComponentType,omitempty"`
	Egress interface{} `json:"Egress,omitempty"`
	Ingress interface{} `json:"Ingress,omitempty"`
}

// AwsAthenaWorkGroupConfigurationDetails represents the AwsAthenaWorkGroupConfigurationDetails schema from the OpenAPI specification
type AwsAthenaWorkGroupConfigurationDetails struct {
	Resultconfiguration interface{} `json:"ResultConfiguration,omitempty"`
}

// AwsEcsServicePlacementStrategiesDetails represents the AwsEcsServicePlacementStrategiesDetails schema from the OpenAPI specification
type AwsEcsServicePlacementStrategiesDetails struct {
	Field interface{} `json:"Field,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// AwsElasticsearchDomainLogPublishingOptionsLogConfig represents the AwsElasticsearchDomainLogPublishingOptionsLogConfig schema from the OpenAPI specification
type AwsElasticsearchDomainLogPublishingOptionsLogConfig struct {
	Cloudwatchlogsloggrouparn interface{} `json:"CloudWatchLogsLogGroupArn,omitempty"`
	Enabled interface{} `json:"Enabled,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsLogConfigurationSecretOptionsDetails represents the AwsEcsTaskDefinitionContainerDefinitionsLogConfigurationSecretOptionsDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsLogConfigurationSecretOptionsDetails struct {
	Name interface{} `json:"Name,omitempty"`
	Valuefrom interface{} `json:"ValueFrom,omitempty"`
}

// AwsWafv2RulesDetails represents the AwsWafv2RulesDetails schema from the OpenAPI specification
type AwsWafv2RulesDetails struct {
	Visibilityconfig interface{} `json:"VisibilityConfig,omitempty"`
	Action interface{} `json:"Action,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Overrideaction interface{} `json:"OverrideAction,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
}

// AvailabilityZone represents the AvailabilityZone schema from the OpenAPI specification
type AvailabilityZone struct {
	Subnetid interface{} `json:"SubnetId,omitempty"`
	Zonename interface{} `json:"ZoneName,omitempty"`
}

// AwsCodeBuildProjectEnvironmentRegistryCredential represents the AwsCodeBuildProjectEnvironmentRegistryCredential schema from the OpenAPI specification
type AwsCodeBuildProjectEnvironmentRegistryCredential struct {
	Credential interface{} `json:"Credential,omitempty"`
	Credentialprovider interface{} `json:"CredentialProvider,omitempty"`
}

// DataClassificationDetails represents the DataClassificationDetails schema from the OpenAPI specification
type DataClassificationDetails struct {
	Detailedresultslocation interface{} `json:"DetailedResultsLocation,omitempty"`
	Result interface{} `json:"Result,omitempty"`
}

// AwsEc2LaunchTemplateDataInstanceMarketOptionsSpotOptionsDetails represents the AwsEc2LaunchTemplateDataInstanceMarketOptionsSpotOptionsDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataInstanceMarketOptionsSpotOptionsDetails struct {
	Spotinstancetype interface{} `json:"SpotInstanceType,omitempty"`
	Validuntil interface{} `json:"ValidUntil,omitempty"`
	Blockdurationminutes interface{} `json:"BlockDurationMinutes,omitempty"`
	Instanceinterruptionbehavior interface{} `json:"InstanceInterruptionBehavior,omitempty"`
	Maxprice interface{} `json:"MaxPrice,omitempty"`
}

// AwsEc2VolumeDetails represents the AwsEc2VolumeDetails schema from the OpenAPI specification
type AwsEc2VolumeDetails struct {
	Volumetype interface{} `json:"VolumeType,omitempty"`
	Volumeid interface{} `json:"VolumeId,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Size interface{} `json:"Size,omitempty"`
	Volumescanstatus interface{} `json:"VolumeScanStatus,omitempty"`
	Devicename interface{} `json:"DeviceName,omitempty"`
	Encrypted interface{} `json:"Encrypted,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Attachments interface{} `json:"Attachments,omitempty"`
	Createtime interface{} `json:"CreateTime,omitempty"`
	Snapshotid interface{} `json:"SnapshotId,omitempty"`
}

// StandardsControl represents the StandardsControl schema from the OpenAPI specification
type StandardsControl struct {
	Controlstatusupdatedat interface{} `json:"ControlStatusUpdatedAt,omitempty"`
	Disabledreason interface{} `json:"DisabledReason,omitempty"`
	Controlid interface{} `json:"ControlId,omitempty"`
	Controlstatus interface{} `json:"ControlStatus,omitempty"`
	Remediationurl interface{} `json:"RemediationUrl,omitempty"`
	Severityrating interface{} `json:"SeverityRating,omitempty"`
	Title interface{} `json:"Title,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Relatedrequirements interface{} `json:"RelatedRequirements,omitempty"`
	Standardscontrolarn interface{} `json:"StandardsControlArn,omitempty"`
}

// AwsAutoScalingLaunchConfigurationBlockDeviceMappingsEbsDetails represents the AwsAutoScalingLaunchConfigurationBlockDeviceMappingsEbsDetails schema from the OpenAPI specification
type AwsAutoScalingLaunchConfigurationBlockDeviceMappingsEbsDetails struct {
	Deleteontermination interface{} `json:"DeleteOnTermination,omitempty"`
	Encrypted interface{} `json:"Encrypted,omitempty"`
	Iops interface{} `json:"Iops,omitempty"`
	Snapshotid interface{} `json:"SnapshotId,omitempty"`
	Volumesize interface{} `json:"VolumeSize,omitempty"`
	Volumetype interface{} `json:"VolumeType,omitempty"`
}

// DeleteInvitationsRequest represents the DeleteInvitationsRequest schema from the OpenAPI specification
type DeleteInvitationsRequest struct {
	Accountids interface{} `json:"AccountIds"`
}

// AwsCodeBuildProjectSource represents the AwsCodeBuildProjectSource schema from the OpenAPI specification
type AwsCodeBuildProjectSource struct {
	Gitclonedepth interface{} `json:"GitCloneDepth,omitempty"`
	Insecuressl interface{} `json:"InsecureSsl,omitempty"`
	Location interface{} `json:"Location,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// AwsAthenaWorkGroupConfigurationResultConfigurationEncryptionConfigurationDetails represents the AwsAthenaWorkGroupConfigurationResultConfigurationEncryptionConfigurationDetails schema from the OpenAPI specification
type AwsAthenaWorkGroupConfigurationResultConfigurationEncryptionConfigurationDetails struct {
	Encryptionoption interface{} `json:"EncryptionOption,omitempty"`
	Kmskey interface{} `json:"KmsKey,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsDependsOnDetails represents the AwsEcsTaskDefinitionContainerDefinitionsDependsOnDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsDependsOnDetails struct {
	Condition interface{} `json:"Condition,omitempty"`
	Containername interface{} `json:"ContainerName,omitempty"`
}

// AwsRdsPendingCloudWatchLogsExports represents the AwsRdsPendingCloudWatchLogsExports schema from the OpenAPI specification
type AwsRdsPendingCloudWatchLogsExports struct {
	Logtypestodisable interface{} `json:"LogTypesToDisable,omitempty"`
	Logtypestoenable interface{} `json:"LogTypesToEnable,omitempty"`
}

// AwsRdsDbParameterGroup represents the AwsRdsDbParameterGroup schema from the OpenAPI specification
type AwsRdsDbParameterGroup struct {
	Dbparametergroupname interface{} `json:"DbParameterGroupName,omitempty"`
	Parameterapplystatus interface{} `json:"ParameterApplyStatus,omitempty"`
}

// AwsElbLoadBalancerBackendServerDescription represents the AwsElbLoadBalancerBackendServerDescription schema from the OpenAPI specification
type AwsElbLoadBalancerBackendServerDescription struct {
	Policynames interface{} `json:"PolicyNames,omitempty"`
	Instanceport interface{} `json:"InstancePort,omitempty"`
}

// AwsEksClusterLoggingDetails represents the AwsEksClusterLoggingDetails schema from the OpenAPI specification
type AwsEksClusterLoggingDetails struct {
	Clusterlogging interface{} `json:"ClusterLogging,omitempty"`
}

// AwsEcsTaskDefinitionProxyConfigurationDetails represents the AwsEcsTaskDefinitionProxyConfigurationDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionProxyConfigurationDetails struct {
	Containername interface{} `json:"ContainerName,omitempty"`
	Proxyconfigurationproperties interface{} `json:"ProxyConfigurationProperties,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// AwsWafv2CustomResponseDetails represents the AwsWafv2CustomResponseDetails schema from the OpenAPI specification
type AwsWafv2CustomResponseDetails struct {
	Responsecode interface{} `json:"ResponseCode,omitempty"`
	Responseheaders interface{} `json:"ResponseHeaders,omitempty"`
	Customresponsebodykey interface{} `json:"CustomResponseBodyKey,omitempty"`
}

// AwsEc2InstanceMetadataOptions represents the AwsEc2InstanceMetadataOptions schema from the OpenAPI specification
type AwsEc2InstanceMetadataOptions struct {
	Httpprotocolipv6 interface{} `json:"HttpProtocolIpv6,omitempty"`
	Httpputresponsehoplimit interface{} `json:"HttpPutResponseHopLimit,omitempty"`
	Httptokens interface{} `json:"HttpTokens,omitempty"`
	Instancemetadatatags interface{} `json:"InstanceMetadataTags,omitempty"`
	Httpendpoint interface{} `json:"HttpEndpoint,omitempty"`
}

// AwsElbAppCookieStickinessPolicy represents the AwsElbAppCookieStickinessPolicy schema from the OpenAPI specification
type AwsElbAppCookieStickinessPolicy struct {
	Cookiename interface{} `json:"CookieName,omitempty"`
	Policyname interface{} `json:"PolicyName,omitempty"`
}

// AwsWafRegionalWebAclRulesListDetails represents the AwsWafRegionalWebAclRulesListDetails schema from the OpenAPI specification
type AwsWafRegionalWebAclRulesListDetails struct {
	Priority interface{} `json:"Priority,omitempty"`
	Ruleid interface{} `json:"RuleId,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Action interface{} `json:"Action,omitempty"`
	Overrideaction interface{} `json:"OverrideAction,omitempty"`
}

// BatchUpdateAutomationRulesResponse represents the BatchUpdateAutomationRulesResponse schema from the OpenAPI specification
type BatchUpdateAutomationRulesResponse struct {
	Unprocessedautomationrules interface{} `json:"UnprocessedAutomationRules,omitempty"`
	Processedautomationrules interface{} `json:"ProcessedAutomationRules,omitempty"`
}

// SensitiveDataResult represents the SensitiveDataResult schema from the OpenAPI specification
type SensitiveDataResult struct {
	Category interface{} `json:"Category,omitempty"`
	Detections interface{} `json:"Detections,omitempty"`
	Totalcount interface{} `json:"TotalCount,omitempty"`
}

// CidrBlockAssociation represents the CidrBlockAssociation schema from the OpenAPI specification
type CidrBlockAssociation struct {
	Cidrblockstate interface{} `json:"CidrBlockState,omitempty"`
	Associationid interface{} `json:"AssociationId,omitempty"`
	Cidrblock interface{} `json:"CidrBlock,omitempty"`
}

// AwsElasticsearchDomainNodeToNodeEncryptionOptions represents the AwsElasticsearchDomainNodeToNodeEncryptionOptions schema from the OpenAPI specification
type AwsElasticsearchDomainNodeToNodeEncryptionOptions struct {
	Enabled interface{} `json:"Enabled,omitempty"`
}

// AwsEc2LaunchTemplateDataNetworkInterfaceSetIpv4PrefixesDetails represents the AwsEc2LaunchTemplateDataNetworkInterfaceSetIpv4PrefixesDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataNetworkInterfaceSetIpv4PrefixesDetails struct {
	Ipv4prefix interface{} `json:"Ipv4Prefix,omitempty"`
}

// AwsDynamoDbTableProjection represents the AwsDynamoDbTableProjection schema from the OpenAPI specification
type AwsDynamoDbTableProjection struct {
	Nonkeyattributes interface{} `json:"NonKeyAttributes,omitempty"`
	Projectiontype interface{} `json:"ProjectionType,omitempty"`
}

// AwsDynamoDbTableBillingModeSummary represents the AwsDynamoDbTableBillingModeSummary schema from the OpenAPI specification
type AwsDynamoDbTableBillingModeSummary struct {
	Lastupdatetopayperrequestdatetime interface{} `json:"LastUpdateToPayPerRequestDateTime,omitempty"`
	Billingmode interface{} `json:"BillingMode,omitempty"`
}

// AwsWafv2WebAclCaptchaConfigDetails represents the AwsWafv2WebAclCaptchaConfigDetails schema from the OpenAPI specification
type AwsWafv2WebAclCaptchaConfigDetails struct {
	Immunitytimeproperty interface{} `json:"ImmunityTimeProperty,omitempty"`
}

// AwsAmazonMqBrokerLogsDetails represents the AwsAmazonMqBrokerLogsDetails schema from the OpenAPI specification
type AwsAmazonMqBrokerLogsDetails struct {
	Pending interface{} `json:"Pending,omitempty"`
	Audit interface{} `json:"Audit,omitempty"`
	Auditloggroup interface{} `json:"AuditLogGroup,omitempty"`
	General interface{} `json:"General,omitempty"`
	Generalloggroup interface{} `json:"GeneralLogGroup,omitempty"`
}

// AwsRdsDbDomainMembership represents the AwsRdsDbDomainMembership schema from the OpenAPI specification
type AwsRdsDbDomainMembership struct {
	Domain interface{} `json:"Domain,omitempty"`
	Fqdn interface{} `json:"Fqdn,omitempty"`
	Iamrolename interface{} `json:"IamRoleName,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// Workflow represents the Workflow schema from the OpenAPI specification
type Workflow struct {
	Status interface{} `json:"Status,omitempty"`
}

// AwsAutoScalingAutoScalingGroupMixedInstancesPolicyDetails represents the AwsAutoScalingAutoScalingGroupMixedInstancesPolicyDetails schema from the OpenAPI specification
type AwsAutoScalingAutoScalingGroupMixedInstancesPolicyDetails struct {
	Instancesdistribution interface{} `json:"InstancesDistribution,omitempty"`
	Launchtemplate interface{} `json:"LaunchTemplate,omitempty"`
}

// AwsEc2LaunchTemplateDetails represents the AwsEc2LaunchTemplateDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDetails struct {
	Latestversionnumber interface{} `json:"LatestVersionNumber,omitempty"`
	Launchtemplatedata interface{} `json:"LaunchTemplateData,omitempty"`
	Launchtemplatename interface{} `json:"LaunchTemplateName,omitempty"`
	Defaultversionnumber interface{} `json:"DefaultVersionNumber,omitempty"`
	Id interface{} `json:"Id,omitempty"`
}

// AwsEcsTaskDefinitionProxyConfigurationProxyConfigurationPropertiesDetails represents the AwsEcsTaskDefinitionProxyConfigurationProxyConfigurationPropertiesDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionProxyConfigurationProxyConfigurationPropertiesDetails struct {
	Value interface{} `json:"Value,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// AwsWafRegionalWebAclDetails represents the AwsWafRegionalWebAclDetails schema from the OpenAPI specification
type AwsWafRegionalWebAclDetails struct {
	Defaultaction interface{} `json:"DefaultAction,omitempty"`
	Metricname interface{} `json:"MetricName,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Ruleslist interface{} `json:"RulesList,omitempty"`
	Webaclid interface{} `json:"WebAclId,omitempty"`
}

// UpdateSecurityHubConfigurationRequest represents the UpdateSecurityHubConfigurationRequest schema from the OpenAPI specification
type UpdateSecurityHubConfigurationRequest struct {
	Autoenablecontrols interface{} `json:"AutoEnableControls,omitempty"`
	Controlfindinggenerator interface{} `json:"ControlFindingGenerator,omitempty"`
}

// SortCriterion represents the SortCriterion schema from the OpenAPI specification
type SortCriterion struct {
	Field interface{} `json:"Field,omitempty"`
	Sortorder interface{} `json:"SortOrder,omitempty"`
}

// ActionRemotePortDetails represents the ActionRemotePortDetails schema from the OpenAPI specification
type ActionRemotePortDetails struct {
	Port interface{} `json:"Port,omitempty"`
	Portname interface{} `json:"PortName,omitempty"`
}

// AwsS3BucketServerSideEncryptionRule represents the AwsS3BucketServerSideEncryptionRule schema from the OpenAPI specification
type AwsS3BucketServerSideEncryptionRule struct {
	Applyserversideencryptionbydefault interface{} `json:"ApplyServerSideEncryptionByDefault,omitempty"`
}

// UpdateFindingsResponse represents the UpdateFindingsResponse schema from the OpenAPI specification
type UpdateFindingsResponse struct {
}

// DisassociateFromAdministratorAccountResponse represents the DisassociateFromAdministratorAccountResponse schema from the OpenAPI specification
type DisassociateFromAdministratorAccountResponse struct {
}

// ListFindingAggregatorsResponse represents the ListFindingAggregatorsResponse schema from the OpenAPI specification
type ListFindingAggregatorsResponse struct {
	Findingaggregators interface{} `json:"FindingAggregators,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// AwsSsmPatchComplianceDetails represents the AwsSsmPatchComplianceDetails schema from the OpenAPI specification
type AwsSsmPatchComplianceDetails struct {
	Patch interface{} `json:"Patch,omitempty"`
}

// ProcessDetails represents the ProcessDetails schema from the OpenAPI specification
type ProcessDetails struct {
	Name interface{} `json:"Name,omitempty"`
	Parentpid interface{} `json:"ParentPid,omitempty"`
	Path interface{} `json:"Path,omitempty"`
	Pid interface{} `json:"Pid,omitempty"`
	Terminatedat interface{} `json:"TerminatedAt,omitempty"`
	Launchedat interface{} `json:"LaunchedAt,omitempty"`
}

// AwsRedshiftClusterClusterParameterStatus represents the AwsRedshiftClusterClusterParameterStatus schema from the OpenAPI specification
type AwsRedshiftClusterClusterParameterStatus struct {
	Parameterapplyerrordescription interface{} `json:"ParameterApplyErrorDescription,omitempty"`
	Parameterapplystatus interface{} `json:"ParameterApplyStatus,omitempty"`
	Parametername interface{} `json:"ParameterName,omitempty"`
}

// AwsElasticBeanstalkEnvironmentEnvironmentLink represents the AwsElasticBeanstalkEnvironmentEnvironmentLink schema from the OpenAPI specification
type AwsElasticBeanstalkEnvironmentEnvironmentLink struct {
	Environmentname interface{} `json:"EnvironmentName,omitempty"`
	Linkname interface{} `json:"LinkName,omitempty"`
}

// PatchSummary represents the PatchSummary schema from the OpenAPI specification
type PatchSummary struct {
	Failedcount interface{} `json:"FailedCount,omitempty"`
	Operationendtime interface{} `json:"OperationEndTime,omitempty"`
	Operationstarttime interface{} `json:"OperationStartTime,omitempty"`
	Rebootoption interface{} `json:"RebootOption,omitempty"`
	Id interface{} `json:"Id"`
	Installedpendingreboot interface{} `json:"InstalledPendingReboot,omitempty"`
	Missingcount interface{} `json:"MissingCount,omitempty"`
	Operation interface{} `json:"Operation,omitempty"`
	Installedcount interface{} `json:"InstalledCount,omitempty"`
	Installedothercount interface{} `json:"InstalledOtherCount,omitempty"`
	Installedrejectedcount interface{} `json:"InstalledRejectedCount,omitempty"`
}

// AwsAutoScalingLaunchConfigurationBlockDeviceMappingsDetails represents the AwsAutoScalingLaunchConfigurationBlockDeviceMappingsDetails schema from the OpenAPI specification
type AwsAutoScalingLaunchConfigurationBlockDeviceMappingsDetails struct {
	Ebs interface{} `json:"Ebs,omitempty"`
	Nodevice interface{} `json:"NoDevice,omitempty"`
	Virtualname interface{} `json:"VirtualName,omitempty"`
	Devicename interface{} `json:"DeviceName,omitempty"`
}

// RuleGroupSourceStatelessRuleDefinition represents the RuleGroupSourceStatelessRuleDefinition schema from the OpenAPI specification
type RuleGroupSourceStatelessRuleDefinition struct {
	Actions interface{} `json:"Actions,omitempty"`
	Matchattributes interface{} `json:"MatchAttributes,omitempty"`
}

// AwsKinesisStreamStreamEncryptionDetails represents the AwsKinesisStreamStreamEncryptionDetails schema from the OpenAPI specification
type AwsKinesisStreamStreamEncryptionDetails struct {
	Encryptiontype interface{} `json:"EncryptionType,omitempty"`
	Keyid interface{} `json:"KeyId,omitempty"`
}

// AwsAutoScalingLaunchConfigurationDetails represents the AwsAutoScalingLaunchConfigurationDetails schema from the OpenAPI specification
type AwsAutoScalingLaunchConfigurationDetails struct {
	Instancemonitoring interface{} `json:"InstanceMonitoring,omitempty"`
	Securitygroups interface{} `json:"SecurityGroups,omitempty"`
	Classiclinkvpcsecuritygroups interface{} `json:"ClassicLinkVpcSecurityGroups,omitempty"`
	Placementtenancy interface{} `json:"PlacementTenancy,omitempty"`
	Userdata interface{} `json:"UserData,omitempty"`
	Imageid interface{} `json:"ImageId,omitempty"`
	Iaminstanceprofile interface{} `json:"IamInstanceProfile,omitempty"`
	Classiclinkvpcid interface{} `json:"ClassicLinkVpcId,omitempty"`
	Ramdiskid interface{} `json:"RamdiskId,omitempty"`
	Ebsoptimized interface{} `json:"EbsOptimized,omitempty"`
	Launchconfigurationname interface{} `json:"LaunchConfigurationName,omitempty"`
	Associatepublicipaddress interface{} `json:"AssociatePublicIpAddress,omitempty"`
	Createdtime interface{} `json:"CreatedTime,omitempty"`
	Keyname interface{} `json:"KeyName,omitempty"`
	Blockdevicemappings interface{} `json:"BlockDeviceMappings,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Spotprice interface{} `json:"SpotPrice,omitempty"`
	Kernelid interface{} `json:"KernelId,omitempty"`
	Metadataoptions interface{} `json:"MetadataOptions,omitempty"`
}

// AcceptInvitationResponse represents the AcceptInvitationResponse schema from the OpenAPI specification
type AcceptInvitationResponse struct {
}

// AwsSecretsManagerSecretDetails represents the AwsSecretsManagerSecretDetails schema from the OpenAPI specification
type AwsSecretsManagerSecretDetails struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Rotationenabled interface{} `json:"RotationEnabled,omitempty"`
	Rotationlambdaarn interface{} `json:"RotationLambdaArn,omitempty"`
	Rotationoccurredwithinfrequency interface{} `json:"RotationOccurredWithinFrequency,omitempty"`
	Rotationrules interface{} `json:"RotationRules,omitempty"`
	Deleted interface{} `json:"Deleted,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// AwsRdsDbOptionGroupMembership represents the AwsRdsDbOptionGroupMembership schema from the OpenAPI specification
type AwsRdsDbOptionGroupMembership struct {
	Optiongroupname interface{} `json:"OptionGroupName,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// AwsCloudTrailTrailDetails represents the AwsCloudTrailTrailDetails schema from the OpenAPI specification
type AwsCloudTrailTrailDetails struct {
	Snstopicname interface{} `json:"SnsTopicName,omitempty"`
	S3bucketname interface{} `json:"S3BucketName,omitempty"`
	Cloudwatchlogsloggrouparn interface{} `json:"CloudWatchLogsLogGroupArn,omitempty"`
	Hascustomeventselectors interface{} `json:"HasCustomEventSelectors,omitempty"`
	Ismultiregiontrail interface{} `json:"IsMultiRegionTrail,omitempty"`
	Logfilevalidationenabled interface{} `json:"LogFileValidationEnabled,omitempty"`
	S3keyprefix interface{} `json:"S3KeyPrefix,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Isorganizationtrail interface{} `json:"IsOrganizationTrail,omitempty"`
	Includeglobalserviceevents interface{} `json:"IncludeGlobalServiceEvents,omitempty"`
	Snstopicarn interface{} `json:"SnsTopicArn,omitempty"`
	Trailarn interface{} `json:"TrailArn,omitempty"`
	Cloudwatchlogsrolearn interface{} `json:"CloudWatchLogsRoleArn,omitempty"`
	Homeregion interface{} `json:"HomeRegion,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// GetFindingAggregatorResponse represents the GetFindingAggregatorResponse schema from the OpenAPI specification
type GetFindingAggregatorResponse struct {
	Regions interface{} `json:"Regions,omitempty"`
	Findingaggregationregion interface{} `json:"FindingAggregationRegion,omitempty"`
	Findingaggregatorarn interface{} `json:"FindingAggregatorArn,omitempty"`
	Regionlinkingmode interface{} `json:"RegionLinkingMode,omitempty"`
}

// DeclineInvitationsRequest represents the DeclineInvitationsRequest schema from the OpenAPI specification
type DeclineInvitationsRequest struct {
	Accountids interface{} `json:"AccountIds"`
}

// GetMembersResponse represents the GetMembersResponse schema from the OpenAPI specification
type GetMembersResponse struct {
	Members interface{} `json:"Members,omitempty"`
	Unprocessedaccounts interface{} `json:"UnprocessedAccounts,omitempty"`
}

// AwsStepFunctionStateMachineLoggingConfigurationDetails represents the AwsStepFunctionStateMachineLoggingConfigurationDetails schema from the OpenAPI specification
type AwsStepFunctionStateMachineLoggingConfigurationDetails struct {
	Destinations interface{} `json:"Destinations,omitempty"`
	Includeexecutiondata interface{} `json:"IncludeExecutionData,omitempty"`
	Level interface{} `json:"Level,omitempty"`
}

// AwsDynamoDbTableAttributeDefinition represents the AwsDynamoDbTableAttributeDefinition schema from the OpenAPI specification
type AwsDynamoDbTableAttributeDefinition struct {
	Attributetype interface{} `json:"AttributeType,omitempty"`
	Attributename interface{} `json:"AttributeName,omitempty"`
}

// AwsS3BucketWebsiteConfigurationRedirectTo represents the AwsS3BucketWebsiteConfigurationRedirectTo schema from the OpenAPI specification
type AwsS3BucketWebsiteConfigurationRedirectTo struct {
	Hostname interface{} `json:"Hostname,omitempty"`
	Protocol interface{} `json:"Protocol,omitempty"`
}

// AwsSecretsManagerSecretRotationRules represents the AwsSecretsManagerSecretRotationRules schema from the OpenAPI specification
type AwsSecretsManagerSecretRotationRules struct {
	Automaticallyafterdays interface{} `json:"AutomaticallyAfterDays,omitempty"`
}

// GetFindingHistoryRequest represents the GetFindingHistoryRequest schema from the OpenAPI specification
type GetFindingHistoryRequest struct {
	Findingidentifier AwsSecurityFindingIdentifier `json:"FindingIdentifier"` // Identifies which finding to get the finding history for.
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
}

// AwsElasticsearchDomainDetails represents the AwsElasticsearchDomainDetails schema from the OpenAPI specification
type AwsElasticsearchDomainDetails struct {
	Elasticsearchversion interface{} `json:"ElasticsearchVersion,omitempty"`
	Endpoints interface{} `json:"Endpoints,omitempty"`
	Servicesoftwareoptions interface{} `json:"ServiceSoftwareOptions,omitempty"`
	Vpcoptions interface{} `json:"VPCOptions,omitempty"`
	Endpoint interface{} `json:"Endpoint,omitempty"`
	Encryptionatrestoptions interface{} `json:"EncryptionAtRestOptions,omitempty"`
	Logpublishingoptions interface{} `json:"LogPublishingOptions,omitempty"`
	Nodetonodeencryptionoptions interface{} `json:"NodeToNodeEncryptionOptions,omitempty"`
	Accesspolicies interface{} `json:"AccessPolicies,omitempty"`
	Domainendpointoptions interface{} `json:"DomainEndpointOptions,omitempty"`
	Domainid interface{} `json:"DomainId,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Elasticsearchclusterconfig interface{} `json:"ElasticsearchClusterConfig,omitempty"`
}

// BatchGetAutomationRulesRequest represents the BatchGetAutomationRulesRequest schema from the OpenAPI specification
type BatchGetAutomationRulesRequest struct {
	Automationrulesarns interface{} `json:"AutomationRulesArns"`
}

// AwsIamAccessKeySessionContext represents the AwsIamAccessKeySessionContext schema from the OpenAPI specification
type AwsIamAccessKeySessionContext struct {
	Attributes interface{} `json:"Attributes,omitempty"`
	Sessionissuer interface{} `json:"SessionIssuer,omitempty"`
}

// AwsIamUserPolicy represents the AwsIamUserPolicy schema from the OpenAPI specification
type AwsIamUserPolicy struct {
	Policyname interface{} `json:"PolicyName,omitempty"`
}

// BatchGetSecurityControlsRequest represents the BatchGetSecurityControlsRequest schema from the OpenAPI specification
type BatchGetSecurityControlsRequest struct {
	Securitycontrolids interface{} `json:"SecurityControlIds"`
}

// GetInvitationsCountRequest represents the GetInvitationsCountRequest schema from the OpenAPI specification
type GetInvitationsCountRequest struct {
}

// AwsIamPermissionsBoundary represents the AwsIamPermissionsBoundary schema from the OpenAPI specification
type AwsIamPermissionsBoundary struct {
	Permissionsboundaryarn interface{} `json:"PermissionsBoundaryArn,omitempty"`
	Permissionsboundarytype interface{} `json:"PermissionsBoundaryType,omitempty"`
}

// UpdateOrganizationConfigurationRequest represents the UpdateOrganizationConfigurationRequest schema from the OpenAPI specification
type UpdateOrganizationConfigurationRequest struct {
	Autoenable interface{} `json:"AutoEnable"`
	Autoenablestandards interface{} `json:"AutoEnableStandards,omitempty"`
}

// RuleGroupVariablesIpSetsDetails represents the RuleGroupVariablesIpSetsDetails schema from the OpenAPI specification
type RuleGroupVariablesIpSetsDetails struct {
	Definition interface{} `json:"Definition,omitempty"`
}

// RuleGroupSourceStatefulRulesDetails represents the RuleGroupSourceStatefulRulesDetails schema from the OpenAPI specification
type RuleGroupSourceStatefulRulesDetails struct {
	Action interface{} `json:"Action,omitempty"`
	Header interface{} `json:"Header,omitempty"`
	Ruleoptions interface{} `json:"RuleOptions,omitempty"`
}

// AwsLambdaFunctionVpcConfig represents the AwsLambdaFunctionVpcConfig schema from the OpenAPI specification
type AwsLambdaFunctionVpcConfig struct {
	Vpcid interface{} `json:"VpcId,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
	Subnetids interface{} `json:"SubnetIds,omitempty"`
}

// StandardsStatusReason represents the StandardsStatusReason schema from the OpenAPI specification
type StandardsStatusReason struct {
	Statusreasoncode interface{} `json:"StatusReasonCode"`
}

// AwsCloudFrontDistributionOriginSslProtocols represents the AwsCloudFrontDistributionOriginSslProtocols schema from the OpenAPI specification
type AwsCloudFrontDistributionOriginSslProtocols struct {
	Quantity interface{} `json:"Quantity,omitempty"`
	Items interface{} `json:"Items,omitempty"`
}

// AwsDynamoDbTableSseDescription represents the AwsDynamoDbTableSseDescription schema from the OpenAPI specification
type AwsDynamoDbTableSseDescription struct {
	Ssetype interface{} `json:"SseType,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Inaccessibleencryptiondatetime interface{} `json:"InaccessibleEncryptionDateTime,omitempty"`
	Kmsmasterkeyarn interface{} `json:"KmsMasterKeyArn,omitempty"`
}

// GetInsightsResponse represents the GetInsightsResponse schema from the OpenAPI specification
type GetInsightsResponse struct {
	Insights interface{} `json:"Insights"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// AwsEc2LaunchTemplateDataNetworkInterfaceSetIpv6AddressesDetails represents the AwsEc2LaunchTemplateDataNetworkInterfaceSetIpv6AddressesDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataNetworkInterfaceSetIpv6AddressesDetails struct {
	Ipv6address interface{} `json:"Ipv6Address,omitempty"`
}

// AwsWafv2RulesActionDetails represents the AwsWafv2RulesActionDetails schema from the OpenAPI specification
type AwsWafv2RulesActionDetails struct {
	Block interface{} `json:"Block,omitempty"`
	Captcha interface{} `json:"Captcha,omitempty"`
	Count interface{} `json:"Count,omitempty"`
	Allow interface{} `json:"Allow,omitempty"`
}

// AwsElasticsearchDomainLogPublishingOptions represents the AwsElasticsearchDomainLogPublishingOptions schema from the OpenAPI specification
type AwsElasticsearchDomainLogPublishingOptions struct {
	Searchslowlogs interface{} `json:"SearchSlowLogs,omitempty"`
	Auditlogs AwsElasticsearchDomainLogPublishingOptionsLogConfig `json:"AuditLogs,omitempty"` // The log configuration.
	Indexslowlogs interface{} `json:"IndexSlowLogs,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsEnvironmentFilesDetails represents the AwsEcsTaskDefinitionContainerDefinitionsEnvironmentFilesDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsEnvironmentFilesDetails struct {
	TypeField interface{} `json:"Type,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// AcceptAdministratorInvitationResponse represents the AcceptAdministratorInvitationResponse schema from the OpenAPI specification
type AcceptAdministratorInvitationResponse struct {
}

// AwsWafRuleDetails represents the AwsWafRuleDetails schema from the OpenAPI specification
type AwsWafRuleDetails struct {
	Metricname interface{} `json:"MetricName,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Predicatelist interface{} `json:"PredicateList,omitempty"`
	Ruleid interface{} `json:"RuleId,omitempty"`
}

// VolumeMount represents the VolumeMount schema from the OpenAPI specification
type VolumeMount struct {
	Mountpath interface{} `json:"MountPath,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// AwsAmazonMqBrokerDetails represents the AwsAmazonMqBrokerDetails schema from the OpenAPI specification
type AwsAmazonMqBrokerDetails struct {
	Engineversion interface{} `json:"EngineVersion,omitempty"`
	Ldapservermetadata interface{} `json:"LdapServerMetadata,omitempty"`
	Autominorversionupgrade interface{} `json:"AutoMinorVersionUpgrade,omitempty"`
	Brokerarn interface{} `json:"BrokerArn,omitempty"`
	Enginetype interface{} `json:"EngineType,omitempty"`
	Maintenancewindowstarttime interface{} `json:"MaintenanceWindowStartTime,omitempty"`
	Storagetype interface{} `json:"StorageType,omitempty"`
	Hostinstancetype interface{} `json:"HostInstanceType,omitempty"`
	Subnetids interface{} `json:"SubnetIds,omitempty"`
	Users interface{} `json:"Users,omitempty"`
	Deploymentmode interface{} `json:"DeploymentMode,omitempty"`
	Publiclyaccessible interface{} `json:"PubliclyAccessible,omitempty"`
	Encryptionoptions interface{} `json:"EncryptionOptions,omitempty"`
	Logs interface{} `json:"Logs,omitempty"`
	Securitygroups interface{} `json:"SecurityGroups,omitempty"`
	Authenticationstrategy interface{} `json:"AuthenticationStrategy,omitempty"`
	Brokerid interface{} `json:"BrokerId,omitempty"`
	Brokername interface{} `json:"BrokerName,omitempty"`
}

// AutomationRulesFindingFilters represents the AutomationRulesFindingFilters schema from the OpenAPI specification
type AutomationRulesFindingFilters struct {
	Complianceassociatedstandardsid interface{} `json:"ComplianceAssociatedStandardsId,omitempty"`
	Noteupdatedby interface{} `json:"NoteUpdatedBy,omitempty"`
	Productarn interface{} `json:"ProductArn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Verificationstate interface{} `json:"VerificationState,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Confidence interface{} `json:"Confidence,omitempty"`
	Recordstate interface{} `json:"RecordState,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Relatedfindingsid interface{} `json:"RelatedFindingsId,omitempty"`
	Title interface{} `json:"Title,omitempty"`
	Criticality interface{} `json:"Criticality,omitempty"`
	Resourcepartition interface{} `json:"ResourcePartition,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Compliancesecuritycontrolid interface{} `json:"ComplianceSecurityControlId,omitempty"`
	Resourcetags interface{} `json:"ResourceTags,omitempty"`
	Firstobservedat interface{} `json:"FirstObservedAt,omitempty"`
	Lastobservedat interface{} `json:"LastObservedAt,omitempty"`
	Companyname interface{} `json:"CompanyName,omitempty"`
	Awsaccountid interface{} `json:"AwsAccountId,omitempty"`
	Relatedfindingsproductarn interface{} `json:"RelatedFindingsProductArn,omitempty"`
	Userdefinedfields interface{} `json:"UserDefinedFields,omitempty"`
	Sourceurl interface{} `json:"SourceUrl,omitempty"`
	Resourceid interface{} `json:"ResourceId,omitempty"`
	Compliancestatus interface{} `json:"ComplianceStatus,omitempty"`
	Productname interface{} `json:"ProductName,omitempty"`
	Generatorid interface{} `json:"GeneratorId,omitempty"`
	Severitylabel interface{} `json:"SeverityLabel,omitempty"`
	Notetext interface{} `json:"NoteText,omitempty"`
	Resourcetype interface{} `json:"ResourceType,omitempty"`
	Resourceregion interface{} `json:"ResourceRegion,omitempty"`
	Workflowstatus interface{} `json:"WorkflowStatus,omitempty"`
	Resourcedetailsother interface{} `json:"ResourceDetailsOther,omitempty"`
	Noteupdatedat interface{} `json:"NoteUpdatedAt,omitempty"`
}

// AwsEc2LaunchTemplateDataInstanceRequirementsTotalLocalStorageGBDetails represents the AwsEc2LaunchTemplateDataInstanceRequirementsTotalLocalStorageGBDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataInstanceRequirementsTotalLocalStorageGBDetails struct {
	Max interface{} `json:"Max,omitempty"`
	Min interface{} `json:"Min,omitempty"`
}

// AwsIamAccessKeySessionContextAttributes represents the AwsIamAccessKeySessionContextAttributes schema from the OpenAPI specification
type AwsIamAccessKeySessionContextAttributes struct {
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Mfaauthenticated interface{} `json:"MfaAuthenticated,omitempty"`
}

// DescribeOrganizationConfigurationResponse represents the DescribeOrganizationConfigurationResponse schema from the OpenAPI specification
type DescribeOrganizationConfigurationResponse struct {
	Autoenable interface{} `json:"AutoEnable,omitempty"`
	Autoenablestandards interface{} `json:"AutoEnableStandards,omitempty"`
	Memberaccountlimitreached interface{} `json:"MemberAccountLimitReached,omitempty"`
}

// ListOrganizationAdminAccountsResponse represents the ListOrganizationAdminAccountsResponse schema from the OpenAPI specification
type ListOrganizationAdminAccountsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Adminaccounts interface{} `json:"AdminAccounts,omitempty"`
}

// FindingAggregator represents the FindingAggregator schema from the OpenAPI specification
type FindingAggregator struct {
	Findingaggregatorarn interface{} `json:"FindingAggregatorArn,omitempty"`
}

// AutomationRulesFindingFieldsUpdate represents the AutomationRulesFindingFieldsUpdate schema from the OpenAPI specification
type AutomationRulesFindingFieldsUpdate struct {
	Criticality interface{} `json:"Criticality,omitempty"`
	Note NoteUpdate `json:"Note,omitempty"` // The updated note.
	Severity SeverityUpdate `json:"Severity,omitempty"` // Updates to the severity information for a finding.
	Userdefinedfields interface{} `json:"UserDefinedFields,omitempty"`
	Confidence interface{} `json:"Confidence,omitempty"`
	Relatedfindings interface{} `json:"RelatedFindings,omitempty"`
	Verificationstate interface{} `json:"VerificationState,omitempty"`
	Types interface{} `json:"Types,omitempty"`
	Workflow WorkflowUpdate `json:"Workflow,omitempty"` // Used to update information about the investigation into the finding.
}

// StandardsInputParameterMap represents the StandardsInputParameterMap schema from the OpenAPI specification
type StandardsInputParameterMap struct {
}

// AwsCertificateManagerCertificateDetails represents the AwsCertificateManagerCertificateDetails schema from the OpenAPI specification
type AwsCertificateManagerCertificateDetails struct {
	Notbefore interface{} `json:"NotBefore,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Importedat interface{} `json:"ImportedAt,omitempty"`
	Issuer interface{} `json:"Issuer,omitempty"`
	Keyalgorithm interface{} `json:"KeyAlgorithm,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Options interface{} `json:"Options,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Keyusages interface{} `json:"KeyUsages,omitempty"`
	Notafter interface{} `json:"NotAfter,omitempty"`
	Domainvalidationoptions interface{} `json:"DomainValidationOptions,omitempty"`
	Inuseby interface{} `json:"InUseBy,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Renewaleligibility interface{} `json:"RenewalEligibility,omitempty"`
	Renewalsummary interface{} `json:"RenewalSummary,omitempty"`
	Issuedat interface{} `json:"IssuedAt,omitempty"`
	Subjectalternativenames interface{} `json:"SubjectAlternativeNames,omitempty"`
	Signaturealgorithm interface{} `json:"SignatureAlgorithm,omitempty"`
	Subject interface{} `json:"Subject,omitempty"`
	Certificateauthorityarn interface{} `json:"CertificateAuthorityArn,omitempty"`
	Serial interface{} `json:"Serial,omitempty"`
	Extendedkeyusages interface{} `json:"ExtendedKeyUsages,omitempty"`
}

// AwsCloudFrontDistributionOriginS3OriginConfig represents the AwsCloudFrontDistributionOriginS3OriginConfig schema from the OpenAPI specification
type AwsCloudFrontDistributionOriginS3OriginConfig struct {
	Originaccessidentity interface{} `json:"OriginAccessIdentity,omitempty"`
}

// AwsRedshiftClusterHsmStatus represents the AwsRedshiftClusterHsmStatus schema from the OpenAPI specification
type AwsRedshiftClusterHsmStatus struct {
	Hsmclientcertificateidentifier interface{} `json:"HsmClientCertificateIdentifier,omitempty"`
	Hsmconfigurationidentifier interface{} `json:"HsmConfigurationIdentifier,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// NetworkConnectionAction represents the NetworkConnectionAction schema from the OpenAPI specification
type NetworkConnectionAction struct {
	Connectiondirection interface{} `json:"ConnectionDirection,omitempty"`
	Localportdetails interface{} `json:"LocalPortDetails,omitempty"`
	Protocol interface{} `json:"Protocol,omitempty"`
	Remoteipdetails interface{} `json:"RemoteIpDetails,omitempty"`
	Remoteportdetails interface{} `json:"RemotePortDetails,omitempty"`
	Blocked interface{} `json:"Blocked,omitempty"`
}

// PortProbeAction represents the PortProbeAction schema from the OpenAPI specification
type PortProbeAction struct {
	Blocked interface{} `json:"Blocked,omitempty"`
	Portprobedetails interface{} `json:"PortProbeDetails,omitempty"`
}

// AwsEc2LaunchTemplateDataInstanceRequirementsVCpuCountDetails represents the AwsEc2LaunchTemplateDataInstanceRequirementsVCpuCountDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataInstanceRequirementsVCpuCountDetails struct {
	Max interface{} `json:"Max,omitempty"`
	Min interface{} `json:"Min,omitempty"`
}

// AwsS3BucketWebsiteConfigurationRoutingRule represents the AwsS3BucketWebsiteConfigurationRoutingRule schema from the OpenAPI specification
type AwsS3BucketWebsiteConfigurationRoutingRule struct {
	Condition interface{} `json:"Condition,omitempty"`
	Redirect interface{} `json:"Redirect,omitempty"`
}

// BatchUpdateFindingsUnprocessedFinding represents the BatchUpdateFindingsUnprocessedFinding schema from the OpenAPI specification
type BatchUpdateFindingsUnprocessedFinding struct {
	Errormessage interface{} `json:"ErrorMessage"`
	Findingidentifier interface{} `json:"FindingIdentifier"`
	Errorcode interface{} `json:"ErrorCode"`
}

// GetMasterAccountResponse represents the GetMasterAccountResponse schema from the OpenAPI specification
type GetMasterAccountResponse struct {
	Master interface{} `json:"Master,omitempty"`
}

// AwsWafRateBasedRuleDetails represents the AwsWafRateBasedRuleDetails schema from the OpenAPI specification
type AwsWafRateBasedRuleDetails struct {
	Name interface{} `json:"Name,omitempty"`
	Ratekey interface{} `json:"RateKey,omitempty"`
	Ratelimit interface{} `json:"RateLimit,omitempty"`
	Ruleid interface{} `json:"RuleId,omitempty"`
	Matchpredicates interface{} `json:"MatchPredicates,omitempty"`
	Metricname interface{} `json:"MetricName,omitempty"`
}

// AwsEc2LaunchTemplateDataCpuOptionsDetails represents the AwsEc2LaunchTemplateDataCpuOptionsDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataCpuOptionsDetails struct {
	Threadspercore interface{} `json:"ThreadsPerCore,omitempty"`
	Corecount interface{} `json:"CoreCount,omitempty"`
}

// AwsNetworkFirewallFirewallDetails represents the AwsNetworkFirewallFirewallDetails schema from the OpenAPI specification
type AwsNetworkFirewallFirewallDetails struct {
	Deleteprotection interface{} `json:"DeleteProtection,omitempty"`
	Firewallpolicyarn interface{} `json:"FirewallPolicyArn,omitempty"`
	Firewallarn interface{} `json:"FirewallArn,omitempty"`
	Firewallname interface{} `json:"FirewallName,omitempty"`
	Subnetchangeprotection interface{} `json:"SubnetChangeProtection,omitempty"`
	Firewallpolicychangeprotection interface{} `json:"FirewallPolicyChangeProtection,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Firewallid interface{} `json:"FirewallId,omitempty"`
	Subnetmappings interface{} `json:"SubnetMappings,omitempty"`
}

// Malware represents the Malware schema from the OpenAPI specification
type Malware struct {
	Name interface{} `json:"Name"`
	Path interface{} `json:"Path,omitempty"`
	State interface{} `json:"State,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"Tags"`
}

// BatchUpdateFindingsRequest represents the BatchUpdateFindingsRequest schema from the OpenAPI specification
type BatchUpdateFindingsRequest struct {
	Confidence interface{} `json:"Confidence,omitempty"`
	Note NoteUpdate `json:"Note,omitempty"` // The updated note.
	Userdefinedfields interface{} `json:"UserDefinedFields,omitempty"`
	Workflow interface{} `json:"Workflow,omitempty"`
	Criticality interface{} `json:"Criticality,omitempty"`
	Relatedfindings interface{} `json:"RelatedFindings,omitempty"`
	Severity interface{} `json:"Severity,omitempty"`
	Types interface{} `json:"Types,omitempty"`
	Verificationstate interface{} `json:"VerificationState,omitempty"`
	Findingidentifiers interface{} `json:"FindingIdentifiers"`
}

// DescribeStandardsControlsRequest represents the DescribeStandardsControlsRequest schema from the OpenAPI specification
type DescribeStandardsControlsRequest struct {
}

// DeleteInvitationsResponse represents the DeleteInvitationsResponse schema from the OpenAPI specification
type DeleteInvitationsResponse struct {
	Unprocessedaccounts interface{} `json:"UnprocessedAccounts,omitempty"`
}

// AwsSageMakerNotebookInstanceDetails represents the AwsSageMakerNotebookInstanceDetails schema from the OpenAPI specification
type AwsSageMakerNotebookInstanceDetails struct {
	Platformidentifier interface{} `json:"PlatformIdentifier,omitempty"`
	Rootaccess interface{} `json:"RootAccess,omitempty"`
	Instancemetadataserviceconfiguration interface{} `json:"InstanceMetadataServiceConfiguration,omitempty"`
	Networkinterfaceid interface{} `json:"NetworkInterfaceId,omitempty"`
	Acceleratortypes interface{} `json:"AcceleratorTypes,omitempty"`
	Additionalcoderepositories interface{} `json:"AdditionalCodeRepositories,omitempty"`
	Notebookinstancename interface{} `json:"NotebookInstanceName,omitempty"`
	Securitygroups interface{} `json:"SecurityGroups,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Notebookinstancestatus interface{} `json:"NotebookInstanceStatus,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Notebookinstancearn interface{} `json:"NotebookInstanceArn,omitempty"`
	Url interface{} `json:"Url,omitempty"`
	Volumesizeingb interface{} `json:"VolumeSizeInGB,omitempty"`
	Subnetid interface{} `json:"SubnetId,omitempty"`
	Defaultcoderepository interface{} `json:"DefaultCodeRepository,omitempty"`
	Directinternetaccess interface{} `json:"DirectInternetAccess,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Notebookinstancelifecycleconfigname interface{} `json:"NotebookInstanceLifecycleConfigName,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
}

// Member represents the Member schema from the OpenAPI specification
type Member struct {
	Accountid interface{} `json:"AccountId,omitempty"`
	Administratorid interface{} `json:"AdministratorId,omitempty"`
	Email interface{} `json:"Email,omitempty"`
	Invitedat interface{} `json:"InvitedAt,omitempty"`
	Masterid interface{} `json:"MasterId,omitempty"`
	Memberstatus interface{} `json:"MemberStatus,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
}

// AwsElbLbCookieStickinessPolicy represents the AwsElbLbCookieStickinessPolicy schema from the OpenAPI specification
type AwsElbLbCookieStickinessPolicy struct {
	Cookieexpirationperiod interface{} `json:"CookieExpirationPeriod,omitempty"`
	Policyname interface{} `json:"PolicyName,omitempty"`
}

// AwsWafRegionalRuleGroupDetails represents the AwsWafRegionalRuleGroupDetails schema from the OpenAPI specification
type AwsWafRegionalRuleGroupDetails struct {
	Metricname interface{} `json:"MetricName,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Rulegroupid interface{} `json:"RuleGroupId,omitempty"`
	Rules interface{} `json:"Rules,omitempty"`
}

// AwsS3BucketBucketLifecycleConfigurationRulesTransitionsDetails represents the AwsS3BucketBucketLifecycleConfigurationRulesTransitionsDetails schema from the OpenAPI specification
type AwsS3BucketBucketLifecycleConfigurationRulesTransitionsDetails struct {
	Date interface{} `json:"Date,omitempty"`
	Days interface{} `json:"Days,omitempty"`
	Storageclass interface{} `json:"StorageClass,omitempty"`
}

// DisableOrganizationAdminAccountResponse represents the DisableOrganizationAdminAccountResponse schema from the OpenAPI specification
type DisableOrganizationAdminAccountResponse struct {
}

// AwsCloudFrontDistributionOriginItem represents the AwsCloudFrontDistributionOriginItem schema from the OpenAPI specification
type AwsCloudFrontDistributionOriginItem struct {
	Customoriginconfig interface{} `json:"CustomOriginConfig,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Originpath interface{} `json:"OriginPath,omitempty"`
	S3originconfig interface{} `json:"S3OriginConfig,omitempty"`
}

// AwsEcsServiceDetails represents the AwsEcsServiceDetails schema from the OpenAPI specification
type AwsEcsServiceDetails struct {
	Servicename interface{} `json:"ServiceName,omitempty"`
	Desiredcount interface{} `json:"DesiredCount,omitempty"`
	Healthcheckgraceperiodseconds interface{} `json:"HealthCheckGracePeriodSeconds,omitempty"`
	Launchtype interface{} `json:"LaunchType,omitempty"`
	Loadbalancers interface{} `json:"LoadBalancers,omitempty"`
	Placementstrategies interface{} `json:"PlacementStrategies,omitempty"`
	Enableecsmanagedtags interface{} `json:"EnableEcsManagedTags,omitempty"`
	Enableexecutecommand interface{} `json:"EnableExecuteCommand,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Schedulingstrategy interface{} `json:"SchedulingStrategy,omitempty"`
	Capacityproviderstrategy interface{} `json:"CapacityProviderStrategy,omitempty"`
	Cluster interface{} `json:"Cluster,omitempty"`
	Serviceregistries interface{} `json:"ServiceRegistries,omitempty"`
	Placementconstraints interface{} `json:"PlacementConstraints,omitempty"`
	Platformversion interface{} `json:"PlatformVersion,omitempty"`
	Propagatetags interface{} `json:"PropagateTags,omitempty"`
	Deploymentcontroller interface{} `json:"DeploymentController,omitempty"`
	Role interface{} `json:"Role,omitempty"`
	Taskdefinition interface{} `json:"TaskDefinition,omitempty"`
	Deploymentconfiguration interface{} `json:"DeploymentConfiguration,omitempty"`
	Networkconfiguration interface{} `json:"NetworkConfiguration,omitempty"`
	Servicearn interface{} `json:"ServiceArn,omitempty"`
}

// AwsLambdaFunctionLayer represents the AwsLambdaFunctionLayer schema from the OpenAPI specification
type AwsLambdaFunctionLayer struct {
	Arn interface{} `json:"Arn,omitempty"`
	Codesize interface{} `json:"CodeSize,omitempty"`
}

// AwsEc2VpnConnectionOptionsTunnelOptionsDetails represents the AwsEc2VpnConnectionOptionsTunnelOptionsDetails schema from the OpenAPI specification
type AwsEc2VpnConnectionOptionsTunnelOptionsDetails struct {
	Phase2lifetimeseconds interface{} `json:"Phase2LifetimeSeconds,omitempty"`
	Rekeymargintimeseconds interface{} `json:"RekeyMarginTimeSeconds,omitempty"`
	Outsideipaddress interface{} `json:"OutsideIpAddress,omitempty"`
	Phase1dhgroupnumbers interface{} `json:"Phase1DhGroupNumbers,omitempty"`
	Rekeyfuzzpercentage interface{} `json:"RekeyFuzzPercentage,omitempty"`
	Dpdtimeoutseconds interface{} `json:"DpdTimeoutSeconds,omitempty"`
	Ikeversions interface{} `json:"IkeVersions,omitempty"`
	Phase1integrityalgorithms interface{} `json:"Phase1IntegrityAlgorithms,omitempty"`
	Presharedkey interface{} `json:"PreSharedKey,omitempty"`
	Tunnelinsidecidr interface{} `json:"TunnelInsideCidr,omitempty"`
	Replaywindowsize interface{} `json:"ReplayWindowSize,omitempty"`
	Phase2dhgroupnumbers interface{} `json:"Phase2DhGroupNumbers,omitempty"`
	Phase1lifetimeseconds interface{} `json:"Phase1LifetimeSeconds,omitempty"`
	Phase2encryptionalgorithms interface{} `json:"Phase2EncryptionAlgorithms,omitempty"`
	Phase1encryptionalgorithms interface{} `json:"Phase1EncryptionAlgorithms,omitempty"`
	Phase2integrityalgorithms interface{} `json:"Phase2IntegrityAlgorithms,omitempty"`
}

// AwsS3BucketNotificationConfigurationS3KeyFilter represents the AwsS3BucketNotificationConfigurationS3KeyFilter schema from the OpenAPI specification
type AwsS3BucketNotificationConfigurationS3KeyFilter struct {
	Filterrules interface{} `json:"FilterRules,omitempty"`
}

// GetFindingsRequest represents the GetFindingsRequest schema from the OpenAPI specification
type GetFindingsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sortcriteria interface{} `json:"SortCriteria,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
}

// DeleteInsightRequest represents the DeleteInsightRequest schema from the OpenAPI specification
type DeleteInsightRequest struct {
}

// AwsEc2LaunchTemplateDataLicenseSetDetails represents the AwsEc2LaunchTemplateDataLicenseSetDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataLicenseSetDetails struct {
	Licenseconfigurationarn interface{} `json:"LicenseConfigurationArn,omitempty"`
}

// AwsElbv2LoadBalancerDetails represents the AwsElbv2LoadBalancerDetails schema from the OpenAPI specification
type AwsElbv2LoadBalancerDetails struct {
	State interface{} `json:"State,omitempty"`
	Createdtime interface{} `json:"CreatedTime,omitempty"`
	Loadbalancerattributes interface{} `json:"LoadBalancerAttributes,omitempty"`
	Securitygroups interface{} `json:"SecurityGroups,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Canonicalhostedzoneid interface{} `json:"CanonicalHostedZoneId,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Ipaddresstype interface{} `json:"IpAddressType,omitempty"`
	Availabilityzones interface{} `json:"AvailabilityZones,omitempty"`
	Dnsname interface{} `json:"DNSName,omitempty"`
	Scheme interface{} `json:"Scheme,omitempty"`
}

// AwsOpenSearchServiceDomainClusterConfigZoneAwarenessConfigDetails represents the AwsOpenSearchServiceDomainClusterConfigZoneAwarenessConfigDetails schema from the OpenAPI specification
type AwsOpenSearchServiceDomainClusterConfigZoneAwarenessConfigDetails struct {
	Availabilityzonecount interface{} `json:"AvailabilityZoneCount,omitempty"`
}

// CreateActionTargetResponse represents the CreateActionTargetResponse schema from the OpenAPI specification
type CreateActionTargetResponse struct {
	Actiontargetarn interface{} `json:"ActionTargetArn"`
}

// BatchGetStandardsControlAssociationsResponse represents the BatchGetStandardsControlAssociationsResponse schema from the OpenAPI specification
type BatchGetStandardsControlAssociationsResponse struct {
	Standardscontrolassociationdetails interface{} `json:"StandardsControlAssociationDetails"`
	Unprocessedassociations interface{} `json:"UnprocessedAssociations,omitempty"`
}

// GetFindingsResponse represents the GetFindingsResponse schema from the OpenAPI specification
type GetFindingsResponse struct {
	Findings interface{} `json:"Findings"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// VpcInfoCidrBlockSetDetails represents the VpcInfoCidrBlockSetDetails schema from the OpenAPI specification
type VpcInfoCidrBlockSetDetails struct {
	Cidrblock interface{} `json:"CidrBlock,omitempty"`
}

// AwsEcsTaskDefinitionVolumesHostDetails represents the AwsEcsTaskDefinitionVolumesHostDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionVolumesHostDetails struct {
	Sourcepath interface{} `json:"SourcePath,omitempty"`
}

// AwsElbLoadBalancerAttributes represents the AwsElbLoadBalancerAttributes schema from the OpenAPI specification
type AwsElbLoadBalancerAttributes struct {
	Additionalattributes interface{} `json:"AdditionalAttributes,omitempty"`
	Connectiondraining interface{} `json:"ConnectionDraining,omitempty"`
	Connectionsettings interface{} `json:"ConnectionSettings,omitempty"`
	Crosszoneloadbalancing interface{} `json:"CrossZoneLoadBalancing,omitempty"`
	Accesslog interface{} `json:"AccessLog,omitempty"`
}

// AccountDetails represents the AccountDetails schema from the OpenAPI specification
type AccountDetails struct {
	Accountid interface{} `json:"AccountId"`
	Email interface{} `json:"Email,omitempty"`
}

// AwsCloudFormationStackDetails represents the AwsCloudFormationStackDetails schema from the OpenAPI specification
type AwsCloudFormationStackDetails struct {
	Outputs interface{} `json:"Outputs,omitempty"`
	Disablerollback interface{} `json:"DisableRollback,omitempty"`
	Timeoutinminutes interface{} `json:"TimeoutInMinutes,omitempty"`
	Stackid interface{} `json:"StackId,omitempty"`
	Stackstatusreason interface{} `json:"StackStatusReason,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Driftinformation interface{} `json:"DriftInformation,omitempty"`
	Notificationarns interface{} `json:"NotificationArns,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Stackname interface{} `json:"StackName,omitempty"`
	Capabilities interface{} `json:"Capabilities,omitempty"`
	Enableterminationprotection interface{} `json:"EnableTerminationProtection,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Stackstatus interface{} `json:"StackStatus,omitempty"`
	Lastupdatedtime interface{} `json:"LastUpdatedTime,omitempty"`
}

// AwsIamAttachedManagedPolicy represents the AwsIamAttachedManagedPolicy schema from the OpenAPI specification
type AwsIamAttachedManagedPolicy struct {
	Policyarn interface{} `json:"PolicyArn,omitempty"`
	Policyname interface{} `json:"PolicyName,omitempty"`
}

// DateFilter represents the DateFilter schema from the OpenAPI specification
type DateFilter struct {
	End interface{} `json:"End,omitempty"`
	Start interface{} `json:"Start,omitempty"`
	Daterange interface{} `json:"DateRange,omitempty"`
}

// AwsEc2LaunchTemplateDataInstanceRequirementsNetworkInterfaceCountDetails represents the AwsEc2LaunchTemplateDataInstanceRequirementsNetworkInterfaceCountDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataInstanceRequirementsNetworkInterfaceCountDetails struct {
	Max interface{} `json:"Max,omitempty"`
	Min interface{} `json:"Min,omitempty"`
}

// AwsBackupBackupPlanRuleCopyActionsDetails represents the AwsBackupBackupPlanRuleCopyActionsDetails schema from the OpenAPI specification
type AwsBackupBackupPlanRuleCopyActionsDetails struct {
	Lifecycle interface{} `json:"Lifecycle,omitempty"`
	Destinationbackupvaultarn interface{} `json:"DestinationBackupVaultArn,omitempty"`
}

// Cvss represents the Cvss schema from the OpenAPI specification
type Cvss struct {
	Source interface{} `json:"Source,omitempty"`
	Version interface{} `json:"Version,omitempty"`
	Adjustments interface{} `json:"Adjustments,omitempty"`
	Basescore interface{} `json:"BaseScore,omitempty"`
	Basevector interface{} `json:"BaseVector,omitempty"`
}

// AwsEcsClusterConfigurationDetails represents the AwsEcsClusterConfigurationDetails schema from the OpenAPI specification
type AwsEcsClusterConfigurationDetails struct {
	Executecommandconfiguration interface{} `json:"ExecuteCommandConfiguration,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsHealthCheckDetails represents the AwsEcsTaskDefinitionContainerDefinitionsHealthCheckDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsHealthCheckDetails struct {
	Command interface{} `json:"Command,omitempty"`
	Interval interface{} `json:"Interval,omitempty"`
	Retries interface{} `json:"Retries,omitempty"`
	Startperiod interface{} `json:"StartPeriod,omitempty"`
	Timeout interface{} `json:"Timeout,omitempty"`
}

// AwsS3BucketNotificationConfiguration represents the AwsS3BucketNotificationConfiguration schema from the OpenAPI specification
type AwsS3BucketNotificationConfiguration struct {
	Configurations interface{} `json:"Configurations,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsLinuxParametersDevicesDetails represents the AwsEcsTaskDefinitionContainerDefinitionsLinuxParametersDevicesDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsLinuxParametersDevicesDetails struct {
	Containerpath interface{} `json:"ContainerPath,omitempty"`
	Hostpath interface{} `json:"HostPath,omitempty"`
	Permissions interface{} `json:"Permissions,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsExtraHostsDetails represents the AwsEcsTaskDefinitionContainerDefinitionsExtraHostsDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsExtraHostsDetails struct {
	Hostname interface{} `json:"Hostname,omitempty"`
	Ipaddress interface{} `json:"IpAddress,omitempty"`
}

// AwsRdsDbPendingModifiedValues represents the AwsRdsDbPendingModifiedValues schema from the OpenAPI specification
type AwsRdsDbPendingModifiedValues struct {
	Licensemodel interface{} `json:"LicenseModel,omitempty"`
	Dbsubnetgroupname interface{} `json:"DbSubnetGroupName,omitempty"`
	Pendingcloudwatchlogsexports interface{} `json:"PendingCloudWatchLogsExports,omitempty"`
	Dbinstanceclass interface{} `json:"DbInstanceClass,omitempty"`
	Cacertificateidentifier interface{} `json:"CaCertificateIdentifier,omitempty"`
	Dbinstanceidentifier interface{} `json:"DbInstanceIdentifier,omitempty"`
	Engineversion interface{} `json:"EngineVersion,omitempty"`
	Multiaz interface{} `json:"MultiAZ,omitempty"`
	Backupretentionperiod interface{} `json:"BackupRetentionPeriod,omitempty"`
	Processorfeatures interface{} `json:"ProcessorFeatures,omitempty"`
	Storagetype interface{} `json:"StorageType,omitempty"`
	Iops interface{} `json:"Iops,omitempty"`
	Allocatedstorage interface{} `json:"AllocatedStorage,omitempty"`
	Masteruserpassword interface{} `json:"MasterUserPassword,omitempty"`
	Port interface{} `json:"Port,omitempty"`
}

// AwsEcsClusterDetails represents the AwsEcsClusterDetails schema from the OpenAPI specification
type AwsEcsClusterDetails struct {
	Clustersettings interface{} `json:"ClusterSettings,omitempty"`
	Defaultcapacityproviderstrategy interface{} `json:"DefaultCapacityProviderStrategy,omitempty"`
	Runningtaskscount interface{} `json:"RunningTasksCount,omitempty"`
	Clustername interface{} `json:"ClusterName,omitempty"`
	Configuration interface{} `json:"Configuration,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Activeservicescount interface{} `json:"ActiveServicesCount,omitempty"`
	Capacityproviders interface{} `json:"CapacityProviders,omitempty"`
	Clusterarn interface{} `json:"ClusterArn,omitempty"`
	Registeredcontainerinstancescount interface{} `json:"RegisteredContainerInstancesCount,omitempty"`
}

// AwsAthenaWorkGroupDetails represents the AwsAthenaWorkGroupDetails schema from the OpenAPI specification
type AwsAthenaWorkGroupDetails struct {
	Configuration interface{} `json:"Configuration,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	State interface{} `json:"State,omitempty"`
}

// AcceptAdministratorInvitationRequest represents the AcceptAdministratorInvitationRequest schema from the OpenAPI specification
type AcceptAdministratorInvitationRequest struct {
	Administratorid interface{} `json:"AdministratorId"`
	Invitationid interface{} `json:"InvitationId"`
}

// KeywordFilter represents the KeywordFilter schema from the OpenAPI specification
type KeywordFilter struct {
	Value interface{} `json:"Value,omitempty"`
}

// AwsCloudFrontDistributionOriginGroups represents the AwsCloudFrontDistributionOriginGroups schema from the OpenAPI specification
type AwsCloudFrontDistributionOriginGroups struct {
	Items interface{} `json:"Items,omitempty"`
}

// ListOrganizationAdminAccountsRequest represents the ListOrganizationAdminAccountsRequest schema from the OpenAPI specification
type ListOrganizationAdminAccountsRequest struct {
}

// AwsWafRegionalRateBasedRuleMatchPredicate represents the AwsWafRegionalRateBasedRuleMatchPredicate schema from the OpenAPI specification
type AwsWafRegionalRateBasedRuleMatchPredicate struct {
	Dataid interface{} `json:"DataId,omitempty"`
	Negated interface{} `json:"Negated,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// AwsAppSyncGraphQlApiLogConfigDetails represents the AwsAppSyncGraphQlApiLogConfigDetails schema from the OpenAPI specification
type AwsAppSyncGraphQlApiLogConfigDetails struct {
	Cloudwatchlogsrolearn interface{} `json:"CloudWatchLogsRoleArn,omitempty"`
	Excludeverbosecontent interface{} `json:"ExcludeVerboseContent,omitempty"`
	Fieldloglevel interface{} `json:"FieldLogLevel,omitempty"`
}

// AwsCorsConfiguration represents the AwsCorsConfiguration schema from the OpenAPI specification
type AwsCorsConfiguration struct {
	Alloworigins interface{} `json:"AllowOrigins,omitempty"`
	Exposeheaders interface{} `json:"ExposeHeaders,omitempty"`
	Maxage interface{} `json:"MaxAge,omitempty"`
	Allowcredentials interface{} `json:"AllowCredentials,omitempty"`
	Allowheaders interface{} `json:"AllowHeaders,omitempty"`
	Allowmethods interface{} `json:"AllowMethods,omitempty"`
}

// AwsEcsServiceDeploymentConfigurationDetails represents the AwsEcsServiceDeploymentConfigurationDetails schema from the OpenAPI specification
type AwsEcsServiceDeploymentConfigurationDetails struct {
	Deploymentcircuitbreaker interface{} `json:"DeploymentCircuitBreaker,omitempty"`
	Maximumpercent interface{} `json:"MaximumPercent,omitempty"`
	Minimumhealthypercent interface{} `json:"MinimumHealthyPercent,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsLogConfigurationDetails represents the AwsEcsTaskDefinitionContainerDefinitionsLogConfigurationDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsLogConfigurationDetails struct {
	Logdriver interface{} `json:"LogDriver,omitempty"`
	Options interface{} `json:"Options,omitempty"`
	Secretoptions interface{} `json:"SecretOptions,omitempty"`
}

// UpdateFindingsRequest represents the UpdateFindingsRequest schema from the OpenAPI specification
type UpdateFindingsRequest struct {
	Recordstate interface{} `json:"RecordState,omitempty"`
	Filters interface{} `json:"Filters"`
	Note interface{} `json:"Note,omitempty"`
}

// AwsIamPolicyDetails represents the AwsIamPolicyDetails schema from the OpenAPI specification
type AwsIamPolicyDetails struct {
	Isattachable interface{} `json:"IsAttachable,omitempty"`
	Defaultversionid interface{} `json:"DefaultVersionId,omitempty"`
	Path interface{} `json:"Path,omitempty"`
	Permissionsboundaryusagecount interface{} `json:"PermissionsBoundaryUsageCount,omitempty"`
	Policyid interface{} `json:"PolicyId,omitempty"`
	Attachmentcount interface{} `json:"AttachmentCount,omitempty"`
	Policyname interface{} `json:"PolicyName,omitempty"`
	Policyversionlist interface{} `json:"PolicyVersionList,omitempty"`
	Updatedate interface{} `json:"UpdateDate,omitempty"`
	Createdate interface{} `json:"CreateDate,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// AwsEc2VpnConnectionDetails represents the AwsEc2VpnConnectionDetails schema from the OpenAPI specification
type AwsEc2VpnConnectionDetails struct {
	Customergatewayid interface{} `json:"CustomerGatewayId,omitempty"`
	Routes interface{} `json:"Routes,omitempty"`
	Transitgatewayid interface{} `json:"TransitGatewayId,omitempty"`
	Vgwtelemetry interface{} `json:"VgwTelemetry,omitempty"`
	Category interface{} `json:"Category,omitempty"`
	Customergatewayconfiguration interface{} `json:"CustomerGatewayConfiguration,omitempty"`
	Options interface{} `json:"Options,omitempty"`
	Vpngatewayid interface{} `json:"VpnGatewayId,omitempty"`
	State interface{} `json:"State,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Vpnconnectionid interface{} `json:"VpnConnectionId,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsLinuxParametersCapabilitiesDetails represents the AwsEcsTaskDefinitionContainerDefinitionsLinuxParametersCapabilitiesDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsLinuxParametersCapabilitiesDetails struct {
	Add interface{} `json:"Add,omitempty"`
	Drop interface{} `json:"Drop,omitempty"`
}

// AwsDynamoDbTableReplicaGlobalSecondaryIndex represents the AwsDynamoDbTableReplicaGlobalSecondaryIndex schema from the OpenAPI specification
type AwsDynamoDbTableReplicaGlobalSecondaryIndex struct {
	Indexname interface{} `json:"IndexName,omitempty"`
	Provisionedthroughputoverride interface{} `json:"ProvisionedThroughputOverride,omitempty"`
}

// AwsS3BucketDetails represents the AwsS3BucketDetails schema from the OpenAPI specification
type AwsS3BucketDetails struct {
	Owneraccountid interface{} `json:"OwnerAccountId,omitempty"`
	Ownerid interface{} `json:"OwnerId,omitempty"`
	Serversideencryptionconfiguration interface{} `json:"ServerSideEncryptionConfiguration,omitempty"`
	Bucketnotificationconfiguration interface{} `json:"BucketNotificationConfiguration,omitempty"`
	Ownername interface{} `json:"OwnerName,omitempty"`
	Accesscontrollist interface{} `json:"AccessControlList,omitempty"`
	Bucketlifecycleconfiguration interface{} `json:"BucketLifecycleConfiguration,omitempty"`
	Bucketloggingconfiguration interface{} `json:"BucketLoggingConfiguration,omitempty"`
	Bucketversioningconfiguration interface{} `json:"BucketVersioningConfiguration,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Publicaccessblockconfiguration interface{} `json:"PublicAccessBlockConfiguration,omitempty"`
	Bucketwebsiteconfiguration interface{} `json:"BucketWebsiteConfiguration,omitempty"`
	Objectlockconfiguration interface{} `json:"ObjectLockConfiguration,omitempty"`
}

// AwsLambdaFunctionCode represents the AwsLambdaFunctionCode schema from the OpenAPI specification
type AwsLambdaFunctionCode struct {
	S3bucket interface{} `json:"S3Bucket,omitempty"`
	S3key interface{} `json:"S3Key,omitempty"`
	S3objectversion interface{} `json:"S3ObjectVersion,omitempty"`
	Zipfile interface{} `json:"ZipFile,omitempty"`
}

// BatchGetAutomationRulesResponse represents the BatchGetAutomationRulesResponse schema from the OpenAPI specification
type BatchGetAutomationRulesResponse struct {
	Rules interface{} `json:"Rules,omitempty"`
	Unprocessedautomationrules interface{} `json:"UnprocessedAutomationRules,omitempty"`
}

// AwsEc2NetworkInterfacePrivateIpAddressDetail represents the AwsEc2NetworkInterfacePrivateIpAddressDetail schema from the OpenAPI specification
type AwsEc2NetworkInterfacePrivateIpAddressDetail struct {
	Privatednsname interface{} `json:"PrivateDnsName,omitempty"`
	Privateipaddress interface{} `json:"PrivateIpAddress,omitempty"`
}

// AwsEc2LaunchTemplateDataCapacityReservationSpecificationDetails represents the AwsEc2LaunchTemplateDataCapacityReservationSpecificationDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataCapacityReservationSpecificationDetails struct {
	Capacityreservationpreference interface{} `json:"CapacityReservationPreference,omitempty"`
	Capacityreservationtarget interface{} `json:"CapacityReservationTarget,omitempty"`
}

// Severity represents the Severity schema from the OpenAPI specification
type Severity struct {
	Normalized interface{} `json:"Normalized,omitempty"`
	Original interface{} `json:"Original,omitempty"`
	Product interface{} `json:"Product,omitempty"`
	Label interface{} `json:"Label,omitempty"`
}

// AwsEc2VpcPeeringConnectionVpcInfoDetails represents the AwsEc2VpcPeeringConnectionVpcInfoDetails schema from the OpenAPI specification
type AwsEc2VpcPeeringConnectionVpcInfoDetails struct {
	Peeringoptions interface{} `json:"PeeringOptions,omitempty"`
	Region interface{} `json:"Region,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Cidrblock interface{} `json:"CidrBlock,omitempty"`
	Cidrblockset interface{} `json:"CidrBlockSet,omitempty"`
	Ipv6cidrblockset interface{} `json:"Ipv6CidrBlockSet,omitempty"`
	Ownerid interface{} `json:"OwnerId,omitempty"`
}

// AwsKinesisStreamDetails represents the AwsKinesisStreamDetails schema from the OpenAPI specification
type AwsKinesisStreamDetails struct {
	Arn interface{} `json:"Arn,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Retentionperiodhours interface{} `json:"RetentionPeriodHours,omitempty"`
	Shardcount interface{} `json:"ShardCount,omitempty"`
	Streamencryption interface{} `json:"StreamEncryption,omitempty"`
}

// AwsBackupBackupPlanLifecycleDetails represents the AwsBackupBackupPlanLifecycleDetails schema from the OpenAPI specification
type AwsBackupBackupPlanLifecycleDetails struct {
	Deleteafterdays interface{} `json:"DeleteAfterDays,omitempty"`
	Movetocoldstorageafterdays interface{} `json:"MoveToColdStorageAfterDays,omitempty"`
}

// AwsEc2NetworkAclDetails represents the AwsEc2NetworkAclDetails schema from the OpenAPI specification
type AwsEc2NetworkAclDetails struct {
	Associations interface{} `json:"Associations,omitempty"`
	Entries interface{} `json:"Entries,omitempty"`
	Isdefault interface{} `json:"IsDefault,omitempty"`
	Networkaclid interface{} `json:"NetworkAclId,omitempty"`
	Ownerid interface{} `json:"OwnerId,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
}

// AwsApiGatewayV2StageDetails represents the AwsApiGatewayV2StageDetails schema from the OpenAPI specification
type AwsApiGatewayV2StageDetails struct {
	Accesslogsettings interface{} `json:"AccessLogSettings,omitempty"`
	Apigatewaymanaged interface{} `json:"ApiGatewayManaged,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Clientcertificateid interface{} `json:"ClientCertificateId,omitempty"`
	Autodeploy interface{} `json:"AutoDeploy,omitempty"`
	Lastdeploymentstatusmessage interface{} `json:"LastDeploymentStatusMessage,omitempty"`
	Defaultroutesettings interface{} `json:"DefaultRouteSettings,omitempty"`
	Deploymentid interface{} `json:"DeploymentId,omitempty"`
	Lastupdateddate interface{} `json:"LastUpdatedDate,omitempty"`
	Routesettings interface{} `json:"RouteSettings,omitempty"`
	Stagename interface{} `json:"StageName,omitempty"`
	Stagevariables interface{} `json:"StageVariables,omitempty"`
}

// AwsAppSyncGraphQlApiUserPoolConfigDetails represents the AwsAppSyncGraphQlApiUserPoolConfigDetails schema from the OpenAPI specification
type AwsAppSyncGraphQlApiUserPoolConfigDetails struct {
	Appidclientregex interface{} `json:"AppIdClientRegex,omitempty"`
	Awsregion interface{} `json:"AwsRegion,omitempty"`
	Defaultaction interface{} `json:"DefaultAction,omitempty"`
	Userpoolid interface{} `json:"UserPoolId,omitempty"`
}

// AwsWafRegionalRateBasedRuleDetails represents the AwsWafRegionalRateBasedRuleDetails schema from the OpenAPI specification
type AwsWafRegionalRateBasedRuleDetails struct {
	Ruleid interface{} `json:"RuleId,omitempty"`
	Matchpredicates interface{} `json:"MatchPredicates,omitempty"`
	Metricname interface{} `json:"MetricName,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Ratekey interface{} `json:"RateKey,omitempty"`
	Ratelimit interface{} `json:"RateLimit,omitempty"`
}

// AwsCertificateManagerCertificateOptions represents the AwsCertificateManagerCertificateOptions schema from the OpenAPI specification
type AwsCertificateManagerCertificateOptions struct {
	Certificatetransparencyloggingpreference interface{} `json:"CertificateTransparencyLoggingPreference,omitempty"`
}

// RuleGroupSourceStatefulRulesOptionsDetails represents the RuleGroupSourceStatefulRulesOptionsDetails schema from the OpenAPI specification
type RuleGroupSourceStatefulRulesOptionsDetails struct {
	Keyword interface{} `json:"Keyword,omitempty"`
	Settings interface{} `json:"Settings,omitempty"`
}

// AwsSnsTopicSubscription represents the AwsSnsTopicSubscription schema from the OpenAPI specification
type AwsSnsTopicSubscription struct {
	Endpoint interface{} `json:"Endpoint,omitempty"`
	Protocol interface{} `json:"Protocol,omitempty"`
}

// BatchUpdateAutomationRulesRequest represents the BatchUpdateAutomationRulesRequest schema from the OpenAPI specification
type BatchUpdateAutomationRulesRequest struct {
	Updateautomationrulesrequestitems interface{} `json:"UpdateAutomationRulesRequestItems"`
}

// AwsAppSyncGraphQlApiAdditionalAuthenticationProvidersDetails represents the AwsAppSyncGraphQlApiAdditionalAuthenticationProvidersDetails schema from the OpenAPI specification
type AwsAppSyncGraphQlApiAdditionalAuthenticationProvidersDetails struct {
	Authenticationtype interface{} `json:"AuthenticationType,omitempty"`
	Lambdaauthorizerconfig interface{} `json:"LambdaAuthorizerConfig,omitempty"`
	Openidconnectconfig interface{} `json:"OpenIdConnectConfig,omitempty"`
	Userpoolconfig interface{} `json:"UserPoolConfig,omitempty"`
}

// AwsEcsTaskDefinitionDetails represents the AwsEcsTaskDefinitionDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionDetails struct {
	Taskrolearn interface{} `json:"TaskRoleArn,omitempty"`
	Containerdefinitions interface{} `json:"ContainerDefinitions,omitempty"`
	Cpu interface{} `json:"Cpu,omitempty"`
	Networkmode interface{} `json:"NetworkMode,omitempty"`
	Requirescompatibilities interface{} `json:"RequiresCompatibilities,omitempty"`
	Memory interface{} `json:"Memory,omitempty"`
	Placementconstraints interface{} `json:"PlacementConstraints,omitempty"`
	Volumes interface{} `json:"Volumes,omitempty"`
	Executionrolearn interface{} `json:"ExecutionRoleArn,omitempty"`
	Family interface{} `json:"Family,omitempty"`
	Ipcmode interface{} `json:"IpcMode,omitempty"`
	Proxyconfiguration interface{} `json:"ProxyConfiguration,omitempty"`
	Inferenceaccelerators interface{} `json:"InferenceAccelerators,omitempty"`
	Pidmode interface{} `json:"PidMode,omitempty"`
}

// Adjustment represents the Adjustment schema from the OpenAPI specification
type Adjustment struct {
	Metric interface{} `json:"Metric,omitempty"`
	Reason interface{} `json:"Reason,omitempty"`
}

// AwsElasticBeanstalkEnvironmentDetails represents the AwsElasticBeanstalkEnvironmentDetails schema from the OpenAPI specification
type AwsElasticBeanstalkEnvironmentDetails struct {
	Tier interface{} `json:"Tier,omitempty"`
	Versionlabel interface{} `json:"VersionLabel,omitempty"`
	Optionsettings interface{} `json:"OptionSettings,omitempty"`
	Platformarn interface{} `json:"PlatformArn,omitempty"`
	Solutionstackname interface{} `json:"SolutionStackName,omitempty"`
	Applicationname interface{} `json:"ApplicationName,omitempty"`
	Environmentlinks interface{} `json:"EnvironmentLinks,omitempty"`
	Cname interface{} `json:"Cname,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Endpointurl interface{} `json:"EndpointUrl,omitempty"`
	Dateupdated interface{} `json:"DateUpdated,omitempty"`
	Environmentname interface{} `json:"EnvironmentName,omitempty"`
	Datecreated interface{} `json:"DateCreated,omitempty"`
	Environmentarn interface{} `json:"EnvironmentArn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Environmentid interface{} `json:"EnvironmentId,omitempty"`
}

// AwsEc2VolumeAttachment represents the AwsEc2VolumeAttachment schema from the OpenAPI specification
type AwsEc2VolumeAttachment struct {
	Attachtime interface{} `json:"AttachTime,omitempty"`
	Deleteontermination interface{} `json:"DeleteOnTermination,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// AwsStepFunctionStateMachineDetails represents the AwsStepFunctionStateMachineDetails schema from the OpenAPI specification
type AwsStepFunctionStateMachineDetails struct {
	Loggingconfiguration interface{} `json:"LoggingConfiguration,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Statemachinearn interface{} `json:"StateMachineArn,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Tracingconfiguration interface{} `json:"TracingConfiguration,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Label interface{} `json:"Label,omitempty"`
}

// UnprocessedStandardsControlAssociation represents the UnprocessedStandardsControlAssociation schema from the OpenAPI specification
type UnprocessedStandardsControlAssociation struct {
	Standardscontrolassociationid interface{} `json:"StandardsControlAssociationId"`
	Errorcode interface{} `json:"ErrorCode"`
	Errorreason interface{} `json:"ErrorReason,omitempty"`
}

// AwsWafWebAclRule represents the AwsWafWebAclRule schema from the OpenAPI specification
type AwsWafWebAclRule struct {
	Excludedrules interface{} `json:"ExcludedRules,omitempty"`
	Overrideaction interface{} `json:"OverrideAction,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
	Ruleid interface{} `json:"RuleId,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Action interface{} `json:"Action,omitempty"`
}

// ListFindingAggregatorsRequest represents the ListFindingAggregatorsRequest schema from the OpenAPI specification
type ListFindingAggregatorsRequest struct {
}

// GetFindingAggregatorRequest represents the GetFindingAggregatorRequest schema from the OpenAPI specification
type GetFindingAggregatorRequest struct {
}

// AwsRedshiftClusterClusterNode represents the AwsRedshiftClusterClusterNode schema from the OpenAPI specification
type AwsRedshiftClusterClusterNode struct {
	Publicipaddress interface{} `json:"PublicIpAddress,omitempty"`
	Noderole interface{} `json:"NodeRole,omitempty"`
	Privateipaddress interface{} `json:"PrivateIpAddress,omitempty"`
}

// AwsAppSyncGraphQlApiLambdaAuthorizerConfigDetails represents the AwsAppSyncGraphQlApiLambdaAuthorizerConfigDetails schema from the OpenAPI specification
type AwsAppSyncGraphQlApiLambdaAuthorizerConfigDetails struct {
	Authorizeruri interface{} `json:"AuthorizerUri,omitempty"`
	Identityvalidationexpression interface{} `json:"IdentityValidationExpression,omitempty"`
	Authorizerresultttlinseconds interface{} `json:"AuthorizerResultTtlInSeconds,omitempty"`
}

// AwsDynamoDbTableProvisionedThroughput represents the AwsDynamoDbTableProvisionedThroughput schema from the OpenAPI specification
type AwsDynamoDbTableProvisionedThroughput struct {
	Lastincreasedatetime interface{} `json:"LastIncreaseDateTime,omitempty"`
	Numberofdecreasestoday interface{} `json:"NumberOfDecreasesToday,omitempty"`
	Readcapacityunits interface{} `json:"ReadCapacityUnits,omitempty"`
	Writecapacityunits interface{} `json:"WriteCapacityUnits,omitempty"`
	Lastdecreasedatetime interface{} `json:"LastDecreaseDateTime,omitempty"`
}

// StatelessCustomPublishMetricActionDimension represents the StatelessCustomPublishMetricActionDimension schema from the OpenAPI specification
type StatelessCustomPublishMetricActionDimension struct {
	Value interface{} `json:"Value,omitempty"`
}

// AwsCodeBuildProjectLogsConfigCloudWatchLogsDetails represents the AwsCodeBuildProjectLogsConfigCloudWatchLogsDetails schema from the OpenAPI specification
type AwsCodeBuildProjectLogsConfigCloudWatchLogsDetails struct {
	Groupname interface{} `json:"GroupName,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Streamname interface{} `json:"StreamName,omitempty"`
}

// AwsEc2LaunchTemplateDataCapacityReservationSpecificationCapacityReservationTargetDetails represents the AwsEc2LaunchTemplateDataCapacityReservationSpecificationCapacityReservationTargetDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataCapacityReservationSpecificationCapacityReservationTargetDetails struct {
	Capacityreservationid interface{} `json:"CapacityReservationId,omitempty"`
	Capacityreservationresourcegrouparn interface{} `json:"CapacityReservationResourceGroupArn,omitempty"`
}

// StatelessCustomActionDefinition represents the StatelessCustomActionDefinition schema from the OpenAPI specification
type StatelessCustomActionDefinition struct {
	Publishmetricaction interface{} `json:"PublishMetricAction,omitempty"`
}

// UnprocessedSecurityControl represents the UnprocessedSecurityControl schema from the OpenAPI specification
type UnprocessedSecurityControl struct {
	Errorcode interface{} `json:"ErrorCode"`
	Errorreason interface{} `json:"ErrorReason,omitempty"`
	Securitycontrolid interface{} `json:"SecurityControlId"`
}

// RuleGroupSourceStatelessRuleMatchAttributes represents the RuleGroupSourceStatelessRuleMatchAttributes schema from the OpenAPI specification
type RuleGroupSourceStatelessRuleMatchAttributes struct {
	Destinationports interface{} `json:"DestinationPorts,omitempty"`
	Destinations interface{} `json:"Destinations,omitempty"`
	Protocols interface{} `json:"Protocols,omitempty"`
	Sourceports interface{} `json:"SourcePorts,omitempty"`
	Sources interface{} `json:"Sources,omitempty"`
	Tcpflags interface{} `json:"TcpFlags,omitempty"`
}

// AwsCodeBuildProjectLogsConfigS3LogsDetails represents the AwsCodeBuildProjectLogsConfigS3LogsDetails schema from the OpenAPI specification
type AwsCodeBuildProjectLogsConfigS3LogsDetails struct {
	Encryptiondisabled interface{} `json:"EncryptionDisabled,omitempty"`
	Location interface{} `json:"Location,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// AwsApiCallAction represents the AwsApiCallAction schema from the OpenAPI specification
type AwsApiCallAction struct {
	Api interface{} `json:"Api,omitempty"`
	Callertype interface{} `json:"CallerType,omitempty"`
	Domaindetails interface{} `json:"DomainDetails,omitempty"`
	Firstseen interface{} `json:"FirstSeen,omitempty"`
	Lastseen interface{} `json:"LastSeen,omitempty"`
	Remoteipdetails interface{} `json:"RemoteIpDetails,omitempty"`
	Servicename interface{} `json:"ServiceName,omitempty"`
	Affectedresources interface{} `json:"AffectedResources,omitempty"`
}

// AwsEcrRepositoryLifecyclePolicyDetails represents the AwsEcrRepositoryLifecyclePolicyDetails schema from the OpenAPI specification
type AwsEcrRepositoryLifecyclePolicyDetails struct {
	Lifecyclepolicytext interface{} `json:"LifecyclePolicyText,omitempty"`
	Registryid interface{} `json:"RegistryId,omitempty"`
}

// AwsS3BucketObjectLockConfigurationRuleDetails represents the AwsS3BucketObjectLockConfigurationRuleDetails schema from the OpenAPI specification
type AwsS3BucketObjectLockConfigurationRuleDetails struct {
	Defaultretention interface{} `json:"DefaultRetention,omitempty"`
}

// AwsEc2LaunchTemplateDataNetworkInterfaceSetPrivateIpAddressesDetails represents the AwsEc2LaunchTemplateDataNetworkInterfaceSetPrivateIpAddressesDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataNetworkInterfaceSetPrivateIpAddressesDetails struct {
	Primary interface{} `json:"Primary,omitempty"`
	Privateipaddress interface{} `json:"PrivateIpAddress,omitempty"`
}

// StatelessCustomPublishMetricAction represents the StatelessCustomPublishMetricAction schema from the OpenAPI specification
type StatelessCustomPublishMetricAction struct {
	Dimensions interface{} `json:"Dimensions,omitempty"`
}

// AwsElbLoadBalancerCrossZoneLoadBalancing represents the AwsElbLoadBalancerCrossZoneLoadBalancing schema from the OpenAPI specification
type AwsElbLoadBalancerCrossZoneLoadBalancing struct {
	Enabled interface{} `json:"Enabled,omitempty"`
}

// StringFilter represents the StringFilter schema from the OpenAPI specification
type StringFilter struct {
	Value interface{} `json:"Value,omitempty"`
	Comparison interface{} `json:"Comparison,omitempty"`
}

// EnableImportFindingsForProductResponse represents the EnableImportFindingsForProductResponse schema from the OpenAPI specification
type EnableImportFindingsForProductResponse struct {
	Productsubscriptionarn interface{} `json:"ProductSubscriptionArn,omitempty"`
}

// AwsEc2LaunchTemplateDataNetworkInterfaceSetDetails represents the AwsEc2LaunchTemplateDataNetworkInterfaceSetDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataNetworkInterfaceSetDetails struct {
	Networkinterfaceid interface{} `json:"NetworkInterfaceId,omitempty"`
	Privateipaddress interface{} `json:"PrivateIpAddress,omitempty"`
	Associatecarrieripaddress interface{} `json:"AssociateCarrierIpAddress,omitempty"`
	Ipv4prefixes interface{} `json:"Ipv4Prefixes,omitempty"`
	Ipv6prefixes interface{} `json:"Ipv6Prefixes,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Ipv6addresscount interface{} `json:"Ipv6AddressCount,omitempty"`
	Ipv6addresses interface{} `json:"Ipv6Addresses,omitempty"`
	Subnetid interface{} `json:"SubnetId,omitempty"`
	Secondaryprivateipaddresscount interface{} `json:"SecondaryPrivateIpAddressCount,omitempty"`
	Deviceindex interface{} `json:"DeviceIndex,omitempty"`
	Deleteontermination interface{} `json:"DeleteOnTermination,omitempty"`
	Networkcardindex interface{} `json:"NetworkCardIndex,omitempty"`
	Privateipaddresses interface{} `json:"PrivateIpAddresses,omitempty"`
	Ipv6prefixcount interface{} `json:"Ipv6PrefixCount,omitempty"`
	Associatepublicipaddress interface{} `json:"AssociatePublicIpAddress,omitempty"`
	Interfacetype interface{} `json:"InterfaceType,omitempty"`
	Groups interface{} `json:"Groups,omitempty"`
	Ipv4prefixcount interface{} `json:"Ipv4PrefixCount,omitempty"`
}

// AwsAmazonMqBrokerLdapServerMetadataDetails represents the AwsAmazonMqBrokerLdapServerMetadataDetails schema from the OpenAPI specification
type AwsAmazonMqBrokerLdapServerMetadataDetails struct {
	Userbase interface{} `json:"UserBase,omitempty"`
	Hosts interface{} `json:"Hosts,omitempty"`
	Rolesearchmatching interface{} `json:"RoleSearchMatching,omitempty"`
	Serviceaccountusername interface{} `json:"ServiceAccountUsername,omitempty"`
	Usersearchsubtree interface{} `json:"UserSearchSubtree,omitempty"`
	Rolename interface{} `json:"RoleName,omitempty"`
	Rolesearchsubtree interface{} `json:"RoleSearchSubtree,omitempty"`
	Usersearchmatching interface{} `json:"UserSearchMatching,omitempty"`
	Rolebase interface{} `json:"RoleBase,omitempty"`
	Userrolename interface{} `json:"UserRoleName,omitempty"`
}

// FindingProviderSeverity represents the FindingProviderSeverity schema from the OpenAPI specification
type FindingProviderSeverity struct {
	Original interface{} `json:"Original,omitempty"`
	Label interface{} `json:"Label,omitempty"`
}

// AwsRedshiftClusterClusterParameterGroup represents the AwsRedshiftClusterClusterParameterGroup schema from the OpenAPI specification
type AwsRedshiftClusterClusterParameterGroup struct {
	Parametergroupname interface{} `json:"ParameterGroupName,omitempty"`
	Clusterparameterstatuslist interface{} `json:"ClusterParameterStatusList,omitempty"`
	Parameterapplystatus interface{} `json:"ParameterApplyStatus,omitempty"`
}

// DescribeHubResponse represents the DescribeHubResponse schema from the OpenAPI specification
type DescribeHubResponse struct {
	Hubarn interface{} `json:"HubArn,omitempty"`
	Subscribedat interface{} `json:"SubscribedAt,omitempty"`
	Autoenablecontrols interface{} `json:"AutoEnableControls,omitempty"`
	Controlfindinggenerator interface{} `json:"ControlFindingGenerator,omitempty"`
}

// AwsSsmComplianceSummary represents the AwsSsmComplianceSummary schema from the OpenAPI specification
type AwsSsmComplianceSummary struct {
	Compliantunspecifiedcount interface{} `json:"CompliantUnspecifiedCount,omitempty"`
	Overallseverity interface{} `json:"OverallSeverity,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Compliantcriticalcount interface{} `json:"CompliantCriticalCount,omitempty"`
	Noncompliantlowcount interface{} `json:"NonCompliantLowCount,omitempty"`
	Noncompliantunspecifiedcount interface{} `json:"NonCompliantUnspecifiedCount,omitempty"`
	Compliancetype interface{} `json:"ComplianceType,omitempty"`
	Noncompliantinformationalcount interface{} `json:"NonCompliantInformationalCount,omitempty"`
	Noncomplianthighcount interface{} `json:"NonCompliantHighCount,omitempty"`
	Noncompliantcriticalcount interface{} `json:"NonCompliantCriticalCount,omitempty"`
	Patchgroup interface{} `json:"PatchGroup,omitempty"`
	Compliantinformationalcount interface{} `json:"CompliantInformationalCount,omitempty"`
	Compliantmediumcount interface{} `json:"CompliantMediumCount,omitempty"`
	Noncompliantmediumcount interface{} `json:"NonCompliantMediumCount,omitempty"`
	Patchbaselineid interface{} `json:"PatchBaselineId,omitempty"`
	Complianthighcount interface{} `json:"CompliantHighCount,omitempty"`
	Executiontype interface{} `json:"ExecutionType,omitempty"`
	Compliantlowcount interface{} `json:"CompliantLowCount,omitempty"`
}

// NetworkHeader represents the NetworkHeader schema from the OpenAPI specification
type NetworkHeader struct {
	Destination interface{} `json:"Destination,omitempty"`
	Protocol interface{} `json:"Protocol,omitempty"`
	Source interface{} `json:"Source,omitempty"`
}

// AssociationSetDetails represents the AssociationSetDetails schema from the OpenAPI specification
type AssociationSetDetails struct {
	Routetableid interface{} `json:"RouteTableId,omitempty"`
	Subnetid interface{} `json:"SubnetId,omitempty"`
	Associationstate interface{} `json:"AssociationState,omitempty"`
	Gatewayid interface{} `json:"GatewayId,omitempty"`
	Main interface{} `json:"Main,omitempty"`
	Routetableassociationid interface{} `json:"RouteTableAssociationId,omitempty"`
}

// AwsRdsDbClusterOptionGroupMembership represents the AwsRdsDbClusterOptionGroupMembership schema from the OpenAPI specification
type AwsRdsDbClusterOptionGroupMembership struct {
	Dbclusteroptiongroupname interface{} `json:"DbClusterOptionGroupName,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// AwsGuardDutyDetectorDataSourcesS3LogsDetails represents the AwsGuardDutyDetectorDataSourcesS3LogsDetails schema from the OpenAPI specification
type AwsGuardDutyDetectorDataSourcesS3LogsDetails struct {
	Status interface{} `json:"Status,omitempty"`
}

// AwsApiGatewayV2RouteSettings represents the AwsApiGatewayV2RouteSettings schema from the OpenAPI specification
type AwsApiGatewayV2RouteSettings struct {
	Detailedmetricsenabled interface{} `json:"DetailedMetricsEnabled,omitempty"`
	Logginglevel interface{} `json:"LoggingLevel,omitempty"`
	Throttlingburstlimit interface{} `json:"ThrottlingBurstLimit,omitempty"`
	Throttlingratelimit interface{} `json:"ThrottlingRateLimit,omitempty"`
	Datatraceenabled interface{} `json:"DataTraceEnabled,omitempty"`
}

// VpcInfoPeeringOptionsDetails represents the VpcInfoPeeringOptionsDetails schema from the OpenAPI specification
type VpcInfoPeeringOptionsDetails struct {
	Allowdnsresolutionfromremotevpc interface{} `json:"AllowDnsResolutionFromRemoteVpc,omitempty"`
	Allowegressfromlocalclassiclinktoremotevpc interface{} `json:"AllowEgressFromLocalClassicLinkToRemoteVpc,omitempty"`
	Allowegressfromlocalvpctoremoteclassiclink interface{} `json:"AllowEgressFromLocalVpcToRemoteClassicLink,omitempty"`
}

// AwsSageMakerNotebookInstanceMetadataServiceConfigurationDetails represents the AwsSageMakerNotebookInstanceMetadataServiceConfigurationDetails schema from the OpenAPI specification
type AwsSageMakerNotebookInstanceMetadataServiceConfigurationDetails struct {
	Minimuminstancemetadataserviceversion interface{} `json:"MinimumInstanceMetadataServiceVersion,omitempty"`
}

// AwsOpenSearchServiceDomainVpcOptionsDetails represents the AwsOpenSearchServiceDomainVpcOptionsDetails schema from the OpenAPI specification
type AwsOpenSearchServiceDomainVpcOptionsDetails struct {
	Subnetids interface{} `json:"SubnetIds,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
}

// AwsS3BucketBucketLifecycleConfigurationRulesFilterPredicateDetails represents the AwsS3BucketBucketLifecycleConfigurationRulesFilterPredicateDetails schema from the OpenAPI specification
type AwsS3BucketBucketLifecycleConfigurationRulesFilterPredicateDetails struct {
	Operands interface{} `json:"Operands,omitempty"`
	Prefix interface{} `json:"Prefix,omitempty"`
	Tag interface{} `json:"Tag,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// FirewallPolicyStatelessRuleGroupReferencesDetails represents the FirewallPolicyStatelessRuleGroupReferencesDetails schema from the OpenAPI specification
type FirewallPolicyStatelessRuleGroupReferencesDetails struct {
	Priority interface{} `json:"Priority,omitempty"`
	Resourcearn interface{} `json:"ResourceArn,omitempty"`
}

// RuleGroupSourceStatelessRulesAndCustomActionsDetails represents the RuleGroupSourceStatelessRulesAndCustomActionsDetails schema from the OpenAPI specification
type RuleGroupSourceStatelessRulesAndCustomActionsDetails struct {
	Customactions interface{} `json:"CustomActions,omitempty"`
	Statelessrules interface{} `json:"StatelessRules,omitempty"`
}

// AwsEcsTaskDefinitionPlacementConstraintsDetails represents the AwsEcsTaskDefinitionPlacementConstraintsDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionPlacementConstraintsDetails struct {
	Expression interface{} `json:"Expression,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// CustomDataIdentifiersResult represents the CustomDataIdentifiersResult schema from the OpenAPI specification
type CustomDataIdentifiersResult struct {
	Detections interface{} `json:"Detections,omitempty"`
	Totalcount interface{} `json:"TotalCount,omitempty"`
}

// DisassociateFromMasterAccountResponse represents the DisassociateFromMasterAccountResponse schema from the OpenAPI specification
type DisassociateFromMasterAccountResponse struct {
}

// DeleteActionTargetRequest represents the DeleteActionTargetRequest schema from the OpenAPI specification
type DeleteActionTargetRequest struct {
}

// StandardsControlAssociationDetail represents the StandardsControlAssociationDetail schema from the OpenAPI specification
type StandardsControlAssociationDetail struct {
	Associationstatus interface{} `json:"AssociationStatus"`
	Standardscontrolarns interface{} `json:"StandardsControlArns,omitempty"`
	Updatedreason interface{} `json:"UpdatedReason,omitempty"`
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Securitycontrolid interface{} `json:"SecurityControlId"`
	Standardsarn interface{} `json:"StandardsArn"`
	Standardscontroldescription interface{} `json:"StandardsControlDescription,omitempty"`
	Standardscontroltitle interface{} `json:"StandardsControlTitle,omitempty"`
	Relatedrequirements interface{} `json:"RelatedRequirements,omitempty"`
	Securitycontrolarn interface{} `json:"SecurityControlArn"`
}

// AwsCodeBuildProjectDetails represents the AwsCodeBuildProjectDetails schema from the OpenAPI specification
type AwsCodeBuildProjectDetails struct {
	Source interface{} `json:"Source,omitempty"`
	Vpcconfig interface{} `json:"VpcConfig,omitempty"`
	Artifacts interface{} `json:"Artifacts,omitempty"`
	Logsconfig interface{} `json:"LogsConfig,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Encryptionkey interface{} `json:"EncryptionKey,omitempty"`
	Secondaryartifacts interface{} `json:"SecondaryArtifacts,omitempty"`
	Environment interface{} `json:"Environment,omitempty"`
	Servicerole interface{} `json:"ServiceRole,omitempty"`
}

// FirewallPolicyDetails represents the FirewallPolicyDetails schema from the OpenAPI specification
type FirewallPolicyDetails struct {
	Statefulrulegroupreferences interface{} `json:"StatefulRuleGroupReferences,omitempty"`
	Statelesscustomactions interface{} `json:"StatelessCustomActions,omitempty"`
	Statelessdefaultactions interface{} `json:"StatelessDefaultActions,omitempty"`
	Statelessfragmentdefaultactions interface{} `json:"StatelessFragmentDefaultActions,omitempty"`
	Statelessrulegroupreferences interface{} `json:"StatelessRuleGroupReferences,omitempty"`
}

// AwsSecurityFindingIdentifier represents the AwsSecurityFindingIdentifier schema from the OpenAPI specification
type AwsSecurityFindingIdentifier struct {
	Id interface{} `json:"Id"`
	Productarn interface{} `json:"ProductArn"`
}

// FindingProviderFields represents the FindingProviderFields schema from the OpenAPI specification
type FindingProviderFields struct {
	Criticality interface{} `json:"Criticality,omitempty"`
	Relatedfindings interface{} `json:"RelatedFindings,omitempty"`
	Severity interface{} `json:"Severity,omitempty"`
	Types interface{} `json:"Types,omitempty"`
	Confidence interface{} `json:"Confidence,omitempty"`
}

// UnprocessedAutomationRule represents the UnprocessedAutomationRule schema from the OpenAPI specification
type UnprocessedAutomationRule struct {
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Rulearn interface{} `json:"RuleArn,omitempty"`
}

// AwsLambdaLayerVersionDetails represents the AwsLambdaLayerVersionDetails schema from the OpenAPI specification
type AwsLambdaLayerVersionDetails struct {
	Compatibleruntimes interface{} `json:"CompatibleRuntimes,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Version interface{} `json:"Version,omitempty"`
}

// AwsS3BucketWebsiteConfigurationRoutingRuleRedirect represents the AwsS3BucketWebsiteConfigurationRoutingRuleRedirect schema from the OpenAPI specification
type AwsS3BucketWebsiteConfigurationRoutingRuleRedirect struct {
	Protocol interface{} `json:"Protocol,omitempty"`
	Replacekeyprefixwith interface{} `json:"ReplaceKeyPrefixWith,omitempty"`
	Replacekeywith interface{} `json:"ReplaceKeyWith,omitempty"`
	Hostname interface{} `json:"Hostname,omitempty"`
	Httpredirectcode interface{} `json:"HttpRedirectCode,omitempty"`
}

// BatchDisableStandardsRequest represents the BatchDisableStandardsRequest schema from the OpenAPI specification
type BatchDisableStandardsRequest struct {
	Standardssubscriptionarns interface{} `json:"StandardsSubscriptionArns"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// AwsElbLoadBalancerDetails represents the AwsElbLoadBalancerDetails schema from the OpenAPI specification
type AwsElbLoadBalancerDetails struct {
	Canonicalhostedzonenameid interface{} `json:"CanonicalHostedZoneNameID,omitempty"`
	Healthcheck interface{} `json:"HealthCheck,omitempty"`
	Createdtime interface{} `json:"CreatedTime,omitempty"`
	Subnets interface{} `json:"Subnets,omitempty"`
	Loadbalancerattributes interface{} `json:"LoadBalancerAttributes,omitempty"`
	Instances interface{} `json:"Instances,omitempty"`
	Canonicalhostedzonename interface{} `json:"CanonicalHostedZoneName,omitempty"`
	Listenerdescriptions interface{} `json:"ListenerDescriptions,omitempty"`
	Scheme interface{} `json:"Scheme,omitempty"`
	Policies interface{} `json:"Policies,omitempty"`
	Securitygroups interface{} `json:"SecurityGroups,omitempty"`
	Sourcesecuritygroup interface{} `json:"SourceSecurityGroup,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Backendserverdescriptions interface{} `json:"BackendServerDescriptions,omitempty"`
	Loadbalancername interface{} `json:"LoadBalancerName,omitempty"`
	Availabilityzones interface{} `json:"AvailabilityZones,omitempty"`
	Dnsname interface{} `json:"DnsName,omitempty"`
}

// AwsRedshiftClusterIamRole represents the AwsRedshiftClusterIamRole schema from the OpenAPI specification
type AwsRedshiftClusterIamRole struct {
	Iamrolearn interface{} `json:"IamRoleArn,omitempty"`
	Applystatus interface{} `json:"ApplyStatus,omitempty"`
}

// Product represents the Product schema from the OpenAPI specification
type Product struct {
	Categories interface{} `json:"Categories,omitempty"`
	Companyname interface{} `json:"CompanyName,omitempty"`
	Marketplaceurl interface{} `json:"MarketplaceUrl,omitempty"`
	Productname interface{} `json:"ProductName,omitempty"`
	Integrationtypes interface{} `json:"IntegrationTypes,omitempty"`
	Productarn interface{} `json:"ProductArn"`
	Productsubscriptionresourcepolicy interface{} `json:"ProductSubscriptionResourcePolicy,omitempty"`
	Activationurl interface{} `json:"ActivationUrl,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// AwsGuardDutyDetectorDataSourcesMalwareProtectionScanEc2InstanceWithFindingsEbsVolumesDetails represents the AwsGuardDutyDetectorDataSourcesMalwareProtectionScanEc2InstanceWithFindingsEbsVolumesDetails schema from the OpenAPI specification
type AwsGuardDutyDetectorDataSourcesMalwareProtectionScanEc2InstanceWithFindingsEbsVolumesDetails struct {
	Status interface{} `json:"Status,omitempty"`
	Reason interface{} `json:"Reason,omitempty"`
}

// AwsEc2NetworkInterfaceAttachment represents the AwsEc2NetworkInterfaceAttachment schema from the OpenAPI specification
type AwsEc2NetworkInterfaceAttachment struct {
	Deleteontermination interface{} `json:"DeleteOnTermination,omitempty"`
	Deviceindex interface{} `json:"DeviceIndex,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Instanceownerid interface{} `json:"InstanceOwnerId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Attachtime interface{} `json:"AttachTime,omitempty"`
	Attachmentid interface{} `json:"AttachmentId,omitempty"`
}

// AwsLambdaFunctionDetails represents the AwsLambdaFunctionDetails schema from the OpenAPI specification
type AwsLambdaFunctionDetails struct {
	Environment interface{} `json:"Environment,omitempty"`
	Code interface{} `json:"Code,omitempty"`
	Handler interface{} `json:"Handler,omitempty"`
	Memorysize interface{} `json:"MemorySize,omitempty"`
	Runtime interface{} `json:"Runtime,omitempty"`
	Tracingconfig interface{} `json:"TracingConfig,omitempty"`
	Masterarn interface{} `json:"MasterArn,omitempty"`
	Packagetype interface{} `json:"PackageType,omitempty"`
	Timeout interface{} `json:"Timeout,omitempty"`
	Functionname interface{} `json:"FunctionName,omitempty"`
	Revisionid interface{} `json:"RevisionId,omitempty"`
	Lastmodified interface{} `json:"LastModified,omitempty"`
	Version interface{} `json:"Version,omitempty"`
	Layers interface{} `json:"Layers,omitempty"`
	Architectures interface{} `json:"Architectures,omitempty"`
	Kmskeyarn interface{} `json:"KmsKeyArn,omitempty"`
	Deadletterconfig interface{} `json:"DeadLetterConfig,omitempty"`
	Role interface{} `json:"Role,omitempty"`
	Vpcconfig interface{} `json:"VpcConfig,omitempty"`
	Codesha256 interface{} `json:"CodeSha256,omitempty"`
}

// AwsEc2LaunchTemplateDataInstanceRequirementsMemoryMiBDetails represents the AwsEc2LaunchTemplateDataInstanceRequirementsMemoryMiBDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataInstanceRequirementsMemoryMiBDetails struct {
	Max interface{} `json:"Max,omitempty"`
	Min interface{} `json:"Min,omitempty"`
}

// AwsAutoScalingAutoScalingGroupDetails represents the AwsAutoScalingAutoScalingGroupDetails schema from the OpenAPI specification
type AwsAutoScalingAutoScalingGroupDetails struct {
	Launchtemplate interface{} `json:"LaunchTemplate,omitempty"`
	Availabilityzones interface{} `json:"AvailabilityZones,omitempty"`
	Capacityrebalance interface{} `json:"CapacityRebalance,omitempty"`
	Healthchecktype interface{} `json:"HealthCheckType,omitempty"`
	Launchconfigurationname interface{} `json:"LaunchConfigurationName,omitempty"`
	Loadbalancernames interface{} `json:"LoadBalancerNames,omitempty"`
	Createdtime interface{} `json:"CreatedTime,omitempty"`
	Healthcheckgraceperiod interface{} `json:"HealthCheckGracePeriod,omitempty"`
	Mixedinstancespolicy interface{} `json:"MixedInstancesPolicy,omitempty"`
}

// AutomationRulesConfig represents the AutomationRulesConfig schema from the OpenAPI specification
type AutomationRulesConfig struct {
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Actions interface{} `json:"Actions,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Isterminal interface{} `json:"IsTerminal,omitempty"`
	Rulestatus interface{} `json:"RuleStatus,omitempty"`
	Criteria interface{} `json:"Criteria,omitempty"`
	Rulename interface{} `json:"RuleName,omitempty"`
	Ruleorder interface{} `json:"RuleOrder,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Createdby interface{} `json:"CreatedBy,omitempty"`
	Rulearn interface{} `json:"RuleArn,omitempty"`
}

// AwsElbv2LoadBalancerAttribute represents the AwsElbv2LoadBalancerAttribute schema from the OpenAPI specification
type AwsElbv2LoadBalancerAttribute struct {
	Key interface{} `json:"Key,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// AwsEcsTaskDefinitionVolumesEfsVolumeConfigurationAuthorizationConfigDetails represents the AwsEcsTaskDefinitionVolumesEfsVolumeConfigurationAuthorizationConfigDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionVolumesEfsVolumeConfigurationAuthorizationConfigDetails struct {
	Accesspointid interface{} `json:"AccessPointId,omitempty"`
	Iam interface{} `json:"Iam,omitempty"`
}

// BatchGetSecurityControlsResponse represents the BatchGetSecurityControlsResponse schema from the OpenAPI specification
type BatchGetSecurityControlsResponse struct {
	Securitycontrols interface{} `json:"SecurityControls"`
	Unprocessedids interface{} `json:"UnprocessedIds,omitempty"`
}

// AwsEc2SecurityGroupIpPermission represents the AwsEc2SecurityGroupIpPermission schema from the OpenAPI specification
type AwsEc2SecurityGroupIpPermission struct {
	Ipv6ranges interface{} `json:"Ipv6Ranges,omitempty"`
	Prefixlistids interface{} `json:"PrefixListIds,omitempty"`
	Toport interface{} `json:"ToPort,omitempty"`
	Useridgrouppairs interface{} `json:"UserIdGroupPairs,omitempty"`
	Fromport interface{} `json:"FromPort,omitempty"`
	Ipprotocol interface{} `json:"IpProtocol,omitempty"`
	Ipranges interface{} `json:"IpRanges,omitempty"`
}

// IpOrganizationDetails represents the IpOrganizationDetails schema from the OpenAPI specification
type IpOrganizationDetails struct {
	Org interface{} `json:"Org,omitempty"`
	Asn interface{} `json:"Asn,omitempty"`
	Asnorg interface{} `json:"AsnOrg,omitempty"`
	Isp interface{} `json:"Isp,omitempty"`
}

// AwsCloudFrontDistributionOriginGroupFailover represents the AwsCloudFrontDistributionOriginGroupFailover schema from the OpenAPI specification
type AwsCloudFrontDistributionOriginGroupFailover struct {
	Statuscodes interface{} `json:"StatusCodes,omitempty"`
}

// AwsElasticsearchDomainDomainEndpointOptions represents the AwsElasticsearchDomainDomainEndpointOptions schema from the OpenAPI specification
type AwsElasticsearchDomainDomainEndpointOptions struct {
	Tlssecuritypolicy interface{} `json:"TLSSecurityPolicy,omitempty"`
	Enforcehttps interface{} `json:"EnforceHTTPS,omitempty"`
}

// AwsEc2SecurityGroupIpRange represents the AwsEc2SecurityGroupIpRange schema from the OpenAPI specification
type AwsEc2SecurityGroupIpRange struct {
	Cidrip interface{} `json:"CidrIp,omitempty"`
}

// WafExcludedRule represents the WafExcludedRule schema from the OpenAPI specification
type WafExcludedRule struct {
	Ruleid interface{} `json:"RuleId,omitempty"`
}

// ListInvitationsRequest represents the ListInvitationsRequest schema from the OpenAPI specification
type ListInvitationsRequest struct {
}

// AwsRdsDbSubnetGroupSubnetAvailabilityZone represents the AwsRdsDbSubnetGroupSubnetAvailabilityZone schema from the OpenAPI specification
type AwsRdsDbSubnetGroupSubnetAvailabilityZone struct {
	Name interface{} `json:"Name,omitempty"`
}

// GetAdministratorAccountRequest represents the GetAdministratorAccountRequest schema from the OpenAPI specification
type GetAdministratorAccountRequest struct {
}

// AwsKmsKeyDetails represents the AwsKmsKeyDetails schema from the OpenAPI specification
type AwsKmsKeyDetails struct {
	Keymanager interface{} `json:"KeyManager,omitempty"`
	Keyrotationstatus interface{} `json:"KeyRotationStatus,omitempty"`
	Keystate interface{} `json:"KeyState,omitempty"`
	Origin interface{} `json:"Origin,omitempty"`
	Awsaccountid interface{} `json:"AWSAccountId,omitempty"`
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Keyid interface{} `json:"KeyId,omitempty"`
}

// IcmpTypeCode represents the IcmpTypeCode schema from the OpenAPI specification
type IcmpTypeCode struct {
	Code interface{} `json:"Code,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// GetMembersRequest represents the GetMembersRequest schema from the OpenAPI specification
type GetMembersRequest struct {
	Accountids interface{} `json:"AccountIds"`
}

// AdminAccount represents the AdminAccount schema from the OpenAPI specification
type AdminAccount struct {
	Accountid interface{} `json:"AccountId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// AwsAppSyncGraphQlApiOpenIdConnectConfigDetails represents the AwsAppSyncGraphQlApiOpenIdConnectConfigDetails schema from the OpenAPI specification
type AwsAppSyncGraphQlApiOpenIdConnectConfigDetails struct {
	Issuer interface{} `json:"Issuer,omitempty"`
	Authttl interface{} `json:"AuthTtL,omitempty"`
	Clientid interface{} `json:"ClientId,omitempty"`
	Iatttl interface{} `json:"IatTtL,omitempty"`
}

// AwsEcsTaskDetails represents the AwsEcsTaskDetails schema from the OpenAPI specification
type AwsEcsTaskDetails struct {
	Startedby interface{} `json:"StartedBy,omitempty"`
	Taskdefinitionarn interface{} `json:"TaskDefinitionArn,omitempty"`
	Group interface{} `json:"Group,omitempty"`
	Version interface{} `json:"Version,omitempty"`
	Volumes interface{} `json:"Volumes,omitempty"`
	Clusterarn interface{} `json:"ClusterArn,omitempty"`
	Startedat interface{} `json:"StartedAt,omitempty"`
	Containers interface{} `json:"Containers,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
}

// AwsRedshiftClusterDeferredMaintenanceWindow represents the AwsRedshiftClusterDeferredMaintenanceWindow schema from the OpenAPI specification
type AwsRedshiftClusterDeferredMaintenanceWindow struct {
	Defermaintenanceidentifier interface{} `json:"DeferMaintenanceIdentifier,omitempty"`
	Defermaintenancestarttime interface{} `json:"DeferMaintenanceStartTime,omitempty"`
	Defermaintenanceendtime interface{} `json:"DeferMaintenanceEndTime,omitempty"`
}

// AwsEc2EipDetails represents the AwsEc2EipDetails schema from the OpenAPI specification
type AwsEc2EipDetails struct {
	Allocationid interface{} `json:"AllocationId,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Networkbordergroup interface{} `json:"NetworkBorderGroup,omitempty"`
	Networkinterfaceid interface{} `json:"NetworkInterfaceId,omitempty"`
	Networkinterfaceownerid interface{} `json:"NetworkInterfaceOwnerId,omitempty"`
	Privateipaddress interface{} `json:"PrivateIpAddress,omitempty"`
	Publicip interface{} `json:"PublicIp,omitempty"`
	Associationid interface{} `json:"AssociationId,omitempty"`
	Publicipv4pool interface{} `json:"PublicIpv4Pool,omitempty"`
	Domain interface{} `json:"Domain,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// AwsS3BucketWebsiteConfigurationRoutingRuleCondition represents the AwsS3BucketWebsiteConfigurationRoutingRuleCondition schema from the OpenAPI specification
type AwsS3BucketWebsiteConfigurationRoutingRuleCondition struct {
	Httperrorcodereturnedequals interface{} `json:"HttpErrorCodeReturnedEquals,omitempty"`
	Keyprefixequals interface{} `json:"KeyPrefixEquals,omitempty"`
}

// AwsEc2LaunchTemplateDataInstanceMarketOptionsDetails represents the AwsEc2LaunchTemplateDataInstanceMarketOptionsDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataInstanceMarketOptionsDetails struct {
	Spotoptions interface{} `json:"SpotOptions,omitempty"`
	Markettype interface{} `json:"MarketType,omitempty"`
}

// CreateAutomationRuleResponse represents the CreateAutomationRuleResponse schema from the OpenAPI specification
type CreateAutomationRuleResponse struct {
	Rulearn interface{} `json:"RuleArn,omitempty"`
}

// AwsEfsAccessPointRootDirectoryCreationInfoDetails represents the AwsEfsAccessPointRootDirectoryCreationInfoDetails schema from the OpenAPI specification
type AwsEfsAccessPointRootDirectoryCreationInfoDetails struct {
	Owneruid interface{} `json:"OwnerUid,omitempty"`
	Permissions interface{} `json:"Permissions,omitempty"`
	Ownergid interface{} `json:"OwnerGid,omitempty"`
}

// UnprocessedStandardsControlAssociationUpdate represents the UnprocessedStandardsControlAssociationUpdate schema from the OpenAPI specification
type UnprocessedStandardsControlAssociationUpdate struct {
	Errorcode interface{} `json:"ErrorCode"`
	Errorreason interface{} `json:"ErrorReason,omitempty"`
	Standardscontrolassociationupdate interface{} `json:"StandardsControlAssociationUpdate"`
}

// GetEnabledStandardsResponse represents the GetEnabledStandardsResponse schema from the OpenAPI specification
type GetEnabledStandardsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Standardssubscriptions interface{} `json:"StandardsSubscriptions,omitempty"`
}

// AwsIamAccessKeySessionContextSessionIssuer represents the AwsIamAccessKeySessionContextSessionIssuer schema from the OpenAPI specification
type AwsIamAccessKeySessionContextSessionIssuer struct {
	Username interface{} `json:"UserName,omitempty"`
	Accountid interface{} `json:"AccountId,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Principalid interface{} `json:"PrincipalId,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// AwsIamAccessKeyDetails represents the AwsIamAccessKeyDetails schema from the OpenAPI specification
type AwsIamAccessKeyDetails struct {
	Principalid interface{} `json:"PrincipalId,omitempty"`
	Accesskeyid interface{} `json:"AccessKeyId,omitempty"`
	Principaltype interface{} `json:"PrincipalType,omitempty"`
	Username interface{} `json:"UserName,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Principalname interface{} `json:"PrincipalName,omitempty"`
	Sessioncontext interface{} `json:"SessionContext,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Accountid interface{} `json:"AccountId,omitempty"`
}

// AwsS3ObjectDetails represents the AwsS3ObjectDetails schema from the OpenAPI specification
type AwsS3ObjectDetails struct {
	Lastmodified interface{} `json:"LastModified,omitempty"`
	Ssekmskeyid interface{} `json:"SSEKMSKeyId,omitempty"`
	Serversideencryption interface{} `json:"ServerSideEncryption,omitempty"`
	Versionid interface{} `json:"VersionId,omitempty"`
	Contenttype interface{} `json:"ContentType,omitempty"`
	Etag interface{} `json:"ETag,omitempty"`
}

// AwsRdsDbClusterDetails represents the AwsRdsDbClusterDetails schema from the OpenAPI specification
type AwsRdsDbClusterDetails struct {
	Associatedroles interface{} `json:"AssociatedRoles,omitempty"`
	Domainmemberships interface{} `json:"DomainMemberships,omitempty"`
	Httpendpointenabled interface{} `json:"HttpEndpointEnabled,omitempty"`
	Deletionprotection interface{} `json:"DeletionProtection,omitempty"`
	Port interface{} `json:"Port,omitempty"`
	Endpoint interface{} `json:"Endpoint,omitempty"`
	Availabilityzones interface{} `json:"AvailabilityZones,omitempty"`
	Dbsubnetgroup interface{} `json:"DbSubnetGroup,omitempty"`
	Readerendpoint interface{} `json:"ReaderEndpoint,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Crossaccountclone interface{} `json:"CrossAccountClone,omitempty"`
	Multiaz interface{} `json:"MultiAz,omitempty"`
	Vpcsecuritygroups interface{} `json:"VpcSecurityGroups,omitempty"`
	Activitystreamstatus interface{} `json:"ActivityStreamStatus,omitempty"`
	Enginemode interface{} `json:"EngineMode,omitempty"`
	Engine interface{} `json:"Engine,omitempty"`
	Engineversion interface{} `json:"EngineVersion,omitempty"`
	Readreplicaidentifiers interface{} `json:"ReadReplicaIdentifiers,omitempty"`
	Dbclusteridentifier interface{} `json:"DbClusterIdentifier,omitempty"`
	Clustercreatetime interface{} `json:"ClusterCreateTime,omitempty"`
	Dbclusterresourceid interface{} `json:"DbClusterResourceId,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Allocatedstorage interface{} `json:"AllocatedStorage,omitempty"`
	Storageencrypted interface{} `json:"StorageEncrypted,omitempty"`
	Enabledcloudwatchlogsexports interface{} `json:"EnabledCloudWatchLogsExports,omitempty"`
	Copytagstosnapshot interface{} `json:"CopyTagsToSnapshot,omitempty"`
	Hostedzoneid interface{} `json:"HostedZoneId,omitempty"`
	Masterusername interface{} `json:"MasterUsername,omitempty"`
	Customendpoints interface{} `json:"CustomEndpoints,omitempty"`
	Preferredmaintenancewindow interface{} `json:"PreferredMaintenanceWindow,omitempty"`
	Dbclustermembers interface{} `json:"DbClusterMembers,omitempty"`
	Dbclusterparametergroup interface{} `json:"DbClusterParameterGroup,omitempty"`
	Backupretentionperiod interface{} `json:"BackupRetentionPeriod,omitempty"`
	Dbclusteroptiongroupmemberships interface{} `json:"DbClusterOptionGroupMemberships,omitempty"`
	Databasename interface{} `json:"DatabaseName,omitempty"`
	Preferredbackupwindow interface{} `json:"PreferredBackupWindow,omitempty"`
	Iamdatabaseauthenticationenabled interface{} `json:"IamDatabaseAuthenticationEnabled,omitempty"`
}

// BatchImportFindingsRequest represents the BatchImportFindingsRequest schema from the OpenAPI specification
type BatchImportFindingsRequest struct {
	Findings interface{} `json:"Findings"`
}

// AwsS3BucketBucketLifecycleConfigurationRulesNoncurrentVersionTransitionsDetails represents the AwsS3BucketBucketLifecycleConfigurationRulesNoncurrentVersionTransitionsDetails schema from the OpenAPI specification
type AwsS3BucketBucketLifecycleConfigurationRulesNoncurrentVersionTransitionsDetails struct {
	Storageclass interface{} `json:"StorageClass,omitempty"`
	Days interface{} `json:"Days,omitempty"`
}

// AwsApiGatewayStageDetails represents the AwsApiGatewayStageDetails schema from the OpenAPI specification
type AwsApiGatewayStageDetails struct {
	Variables interface{} `json:"Variables,omitempty"`
	Webaclarn interface{} `json:"WebAclArn,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Clientcertificateid interface{} `json:"ClientCertificateId,omitempty"`
	Accesslogsettings interface{} `json:"AccessLogSettings,omitempty"`
	Lastupdateddate interface{} `json:"LastUpdatedDate,omitempty"`
	Stagename interface{} `json:"StageName,omitempty"`
	Cacheclusterenabled interface{} `json:"CacheClusterEnabled,omitempty"`
	Cacheclustersize interface{} `json:"CacheClusterSize,omitempty"`
	Deploymentid interface{} `json:"DeploymentId,omitempty"`
	Tracingenabled interface{} `json:"TracingEnabled,omitempty"`
	Canarysettings interface{} `json:"CanarySettings,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Cacheclusterstatus interface{} `json:"CacheClusterStatus,omitempty"`
	Methodsettings interface{} `json:"MethodSettings,omitempty"`
	Documentationversion interface{} `json:"DocumentationVersion,omitempty"`
}

// AwsRdsDbClusterMember represents the AwsRdsDbClusterMember schema from the OpenAPI specification
type AwsRdsDbClusterMember struct {
	Dbclusterparametergroupstatus interface{} `json:"DbClusterParameterGroupStatus,omitempty"`
	Dbinstanceidentifier interface{} `json:"DbInstanceIdentifier,omitempty"`
	Isclusterwriter interface{} `json:"IsClusterWriter,omitempty"`
	Promotiontier interface{} `json:"PromotionTier,omitempty"`
}

// AwsCloudFrontDistributionOriginGroup represents the AwsCloudFrontDistributionOriginGroup schema from the OpenAPI specification
type AwsCloudFrontDistributionOriginGroup struct {
	Failovercriteria interface{} `json:"FailoverCriteria,omitempty"`
}

// DescribeActionTargetsResponse represents the DescribeActionTargetsResponse schema from the OpenAPI specification
type DescribeActionTargetsResponse struct {
	Actiontargets interface{} `json:"ActionTargets"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// AwsEc2VpcEndpointServiceDetails represents the AwsEc2VpcEndpointServiceDetails schema from the OpenAPI specification
type AwsEc2VpcEndpointServiceDetails struct {
	Privatednsname interface{} `json:"PrivateDnsName,omitempty"`
	Servicename interface{} `json:"ServiceName,omitempty"`
	Baseendpointdnsnames interface{} `json:"BaseEndpointDnsNames,omitempty"`
	Managesvpcendpoints interface{} `json:"ManagesVpcEndpoints,omitempty"`
	Networkloadbalancerarns interface{} `json:"NetworkLoadBalancerArns,omitempty"`
	Acceptancerequired interface{} `json:"AcceptanceRequired,omitempty"`
	Availabilityzones interface{} `json:"AvailabilityZones,omitempty"`
	Gatewayloadbalancerarns interface{} `json:"GatewayLoadBalancerArns,omitempty"`
	Serviceid interface{} `json:"ServiceId,omitempty"`
	Servicestate interface{} `json:"ServiceState,omitempty"`
	Servicetype interface{} `json:"ServiceType,omitempty"`
}

// BooleanFilter represents the BooleanFilter schema from the OpenAPI specification
type BooleanFilter struct {
	Value interface{} `json:"Value,omitempty"`
}

// AwsRdsDbSecurityGroupDetails represents the AwsRdsDbSecurityGroupDetails schema from the OpenAPI specification
type AwsRdsDbSecurityGroupDetails struct {
	Ipranges interface{} `json:"IpRanges,omitempty"`
	Ownerid interface{} `json:"OwnerId,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Dbsecuritygrouparn interface{} `json:"DbSecurityGroupArn,omitempty"`
	Dbsecuritygroupdescription interface{} `json:"DbSecurityGroupDescription,omitempty"`
	Dbsecuritygroupname interface{} `json:"DbSecurityGroupName,omitempty"`
	Ec2securitygroups interface{} `json:"Ec2SecurityGroups,omitempty"`
}

// Resource represents the Resource schema from the OpenAPI specification
type Resource struct {
	Tags interface{} `json:"Tags,omitempty"`
	TypeField interface{} `json:"Type"`
	Dataclassification interface{} `json:"DataClassification,omitempty"`
	Details interface{} `json:"Details,omitempty"`
	Id interface{} `json:"Id"`
	Partition interface{} `json:"Partition,omitempty"`
	Region interface{} `json:"Region,omitempty"`
	Resourcerole interface{} `json:"ResourceRole,omitempty"`
}

// RuleGroupSourceCustomActionsDetails represents the RuleGroupSourceCustomActionsDetails schema from the OpenAPI specification
type RuleGroupSourceCustomActionsDetails struct {
	Actiondefinition interface{} `json:"ActionDefinition,omitempty"`
	Actionname interface{} `json:"ActionName,omitempty"`
}

// AwsEc2LaunchTemplateDataPrivateDnsNameOptionsDetails represents the AwsEc2LaunchTemplateDataPrivateDnsNameOptionsDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataPrivateDnsNameOptionsDetails struct {
	Hostnametype interface{} `json:"HostnameType,omitempty"`
	Enableresourcenamednsaaaarecord interface{} `json:"EnableResourceNameDnsAAAARecord,omitempty"`
	Enableresourcenamednsarecord interface{} `json:"EnableResourceNameDnsARecord,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsResourceRequirementsDetails represents the AwsEcsTaskDefinitionContainerDefinitionsResourceRequirementsDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsResourceRequirementsDetails struct {
	TypeField interface{} `json:"Type,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// WafOverrideAction represents the WafOverrideAction schema from the OpenAPI specification
type WafOverrideAction struct {
	TypeField interface{} `json:"Type,omitempty"`
}

// AwsEc2VpcDetails represents the AwsEc2VpcDetails schema from the OpenAPI specification
type AwsEc2VpcDetails struct {
	Cidrblockassociationset interface{} `json:"CidrBlockAssociationSet,omitempty"`
	Dhcpoptionsid interface{} `json:"DhcpOptionsId,omitempty"`
	Ipv6cidrblockassociationset interface{} `json:"Ipv6CidrBlockAssociationSet,omitempty"`
	State interface{} `json:"State,omitempty"`
}

// StandardsSubscription represents the StandardsSubscription schema from the OpenAPI specification
type StandardsSubscription struct {
	Standardsstatus interface{} `json:"StandardsStatus"`
	Standardsstatusreason interface{} `json:"StandardsStatusReason,omitempty"`
	Standardssubscriptionarn interface{} `json:"StandardsSubscriptionArn"`
	Standardsarn interface{} `json:"StandardsArn"`
	Standardsinput interface{} `json:"StandardsInput"`
}

// AwsDynamoDbTableReplica represents the AwsDynamoDbTableReplica schema from the OpenAPI specification
type AwsDynamoDbTableReplica struct {
	Replicastatusdescription interface{} `json:"ReplicaStatusDescription,omitempty"`
	Globalsecondaryindexes interface{} `json:"GlobalSecondaryIndexes,omitempty"`
	Kmsmasterkeyid interface{} `json:"KmsMasterKeyId,omitempty"`
	Provisionedthroughputoverride interface{} `json:"ProvisionedThroughputOverride,omitempty"`
	Regionname interface{} `json:"RegionName,omitempty"`
	Replicastatus interface{} `json:"ReplicaStatus,omitempty"`
}

// AwsS3BucketWebsiteConfiguration represents the AwsS3BucketWebsiteConfiguration schema from the OpenAPI specification
type AwsS3BucketWebsiteConfiguration struct {
	Redirectallrequeststo interface{} `json:"RedirectAllRequestsTo,omitempty"`
	Routingrules interface{} `json:"RoutingRules,omitempty"`
	Errordocument interface{} `json:"ErrorDocument,omitempty"`
	Indexdocumentsuffix interface{} `json:"IndexDocumentSuffix,omitempty"`
}

// RuleGroupSourceStatelessRuleMatchAttributesDestinationPorts represents the RuleGroupSourceStatelessRuleMatchAttributesDestinationPorts schema from the OpenAPI specification
type RuleGroupSourceStatelessRuleMatchAttributesDestinationPorts struct {
	Fromport interface{} `json:"FromPort,omitempty"`
	Toport interface{} `json:"ToPort,omitempty"`
}

// RuleGroupSourceStatelessRuleMatchAttributesSources represents the RuleGroupSourceStatelessRuleMatchAttributesSources schema from the OpenAPI specification
type RuleGroupSourceStatelessRuleMatchAttributesSources struct {
	Addressdefinition interface{} `json:"AddressDefinition,omitempty"`
}

// AwsNetworkFirewallFirewallPolicyDetails represents the AwsNetworkFirewallFirewallPolicyDetails schema from the OpenAPI specification
type AwsNetworkFirewallFirewallPolicyDetails struct {
	Firewallpolicyid interface{} `json:"FirewallPolicyId,omitempty"`
	Firewallpolicyname interface{} `json:"FirewallPolicyName,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Firewallpolicy interface{} `json:"FirewallPolicy,omitempty"`
	Firewallpolicyarn interface{} `json:"FirewallPolicyArn,omitempty"`
}

// AwsBackupBackupPlanRuleDetails represents the AwsBackupBackupPlanRuleDetails schema from the OpenAPI specification
type AwsBackupBackupPlanRuleDetails struct {
	Targetbackupvault interface{} `json:"TargetBackupVault,omitempty"`
	Ruleid interface{} `json:"RuleId,omitempty"`
	Rulename interface{} `json:"RuleName,omitempty"`
	Startwindowminutes interface{} `json:"StartWindowMinutes,omitempty"`
	Completionwindowminutes interface{} `json:"CompletionWindowMinutes,omitempty"`
	Enablecontinuousbackup interface{} `json:"EnableContinuousBackup,omitempty"`
	Lifecycle interface{} `json:"Lifecycle,omitempty"`
	Copyactions interface{} `json:"CopyActions,omitempty"`
	Scheduleexpression interface{} `json:"ScheduleExpression,omitempty"`
}

// AwsIamGroupDetails represents the AwsIamGroupDetails schema from the OpenAPI specification
type AwsIamGroupDetails struct {
	Path interface{} `json:"Path,omitempty"`
	Attachedmanagedpolicies interface{} `json:"AttachedManagedPolicies,omitempty"`
	Createdate interface{} `json:"CreateDate,omitempty"`
	Groupid interface{} `json:"GroupId,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
	Grouppolicylist interface{} `json:"GroupPolicyList,omitempty"`
}

// AwsOpenSearchServiceDomainMasterUserOptionsDetails represents the AwsOpenSearchServiceDomainMasterUserOptionsDetails schema from the OpenAPI specification
type AwsOpenSearchServiceDomainMasterUserOptionsDetails struct {
	Masteruserpassword interface{} `json:"MasterUserPassword,omitempty"`
	Masteruserarn interface{} `json:"MasterUserArn,omitempty"`
	Masterusername interface{} `json:"MasterUserName,omitempty"`
}

// AwsLambdaFunctionEnvironment represents the AwsLambdaFunctionEnvironment schema from the OpenAPI specification
type AwsLambdaFunctionEnvironment struct {
	Variables interface{} `json:"Variables,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
}

// AwsS3BucketNotificationConfigurationFilter represents the AwsS3BucketNotificationConfigurationFilter schema from the OpenAPI specification
type AwsS3BucketNotificationConfigurationFilter struct {
	S3keyfilter interface{} `json:"S3KeyFilter,omitempty"`
}

// DescribeProductsRequest represents the DescribeProductsRequest schema from the OpenAPI specification
type DescribeProductsRequest struct {
}

// Occurrences represents the Occurrences schema from the OpenAPI specification
type Occurrences struct {
	Offsetranges interface{} `json:"OffsetRanges,omitempty"`
	Pages interface{} `json:"Pages,omitempty"`
	Records interface{} `json:"Records,omitempty"`
	Cells interface{} `json:"Cells,omitempty"`
	Lineranges interface{} `json:"LineRanges,omitempty"`
}

// AwsEcsClusterClusterSettingsDetails represents the AwsEcsClusterClusterSettingsDetails schema from the OpenAPI specification
type AwsEcsClusterClusterSettingsDetails struct {
	Name interface{} `json:"Name,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// AwsBackupBackupPlanDetails represents the AwsBackupBackupPlanDetails schema from the OpenAPI specification
type AwsBackupBackupPlanDetails struct {
	Backupplan interface{} `json:"BackupPlan,omitempty"`
	Backupplanarn interface{} `json:"BackupPlanArn,omitempty"`
	Backupplanid interface{} `json:"BackupPlanId,omitempty"`
	Versionid interface{} `json:"VersionId,omitempty"`
}

// AssociationStateDetails represents the AssociationStateDetails schema from the OpenAPI specification
type AssociationStateDetails struct {
	State interface{} `json:"State,omitempty"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
}

// AwsEc2LaunchTemplateDataInstanceRequirementsBaselineEbsBandwidthMbpsDetails represents the AwsEc2LaunchTemplateDataInstanceRequirementsBaselineEbsBandwidthMbpsDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataInstanceRequirementsBaselineEbsBandwidthMbpsDetails struct {
	Max interface{} `json:"Max,omitempty"`
	Min interface{} `json:"Min,omitempty"`
}

// FirewallPolicyStatefulRuleGroupReferencesDetails represents the FirewallPolicyStatefulRuleGroupReferencesDetails schema from the OpenAPI specification
type FirewallPolicyStatefulRuleGroupReferencesDetails struct {
	Resourcearn interface{} `json:"ResourceArn,omitempty"`
}

// PropagatingVgwSetDetails represents the PropagatingVgwSetDetails schema from the OpenAPI specification
type PropagatingVgwSetDetails struct {
	Gatewayid interface{} `json:"GatewayId,omitempty"`
}

// DescribeStandardsControlsResponse represents the DescribeStandardsControlsResponse schema from the OpenAPI specification
type DescribeStandardsControlsResponse struct {
	Controls interface{} `json:"Controls,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// AwsEc2InstanceMonitoringDetails represents the AwsEc2InstanceMonitoringDetails schema from the OpenAPI specification
type AwsEc2InstanceMonitoringDetails struct {
	State interface{} `json:"State,omitempty"`
}

// AwsCodeBuildProjectEnvironmentEnvironmentVariablesDetails represents the AwsCodeBuildProjectEnvironmentEnvironmentVariablesDetails schema from the OpenAPI specification
type AwsCodeBuildProjectEnvironmentEnvironmentVariablesDetails struct {
	Name interface{} `json:"Name,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// AwsBackupRecoveryPointDetails represents the AwsBackupRecoveryPointDetails schema from the OpenAPI specification
type AwsBackupRecoveryPointDetails struct {
	Recoverypointarn interface{} `json:"RecoveryPointArn,omitempty"`
	Statusmessage interface{} `json:"StatusMessage,omitempty"`
	Backupsizeinbytes interface{} `json:"BackupSizeInBytes,omitempty"`
	Calculatedlifecycle interface{} `json:"CalculatedLifecycle,omitempty"`
	Sourcebackupvaultarn interface{} `json:"SourceBackupVaultArn,omitempty"`
	Backupvaultname interface{} `json:"BackupVaultName,omitempty"`
	Completiondate interface{} `json:"CompletionDate,omitempty"`
	Encryptionkeyarn interface{} `json:"EncryptionKeyArn,omitempty"`
	Iamrolearn interface{} `json:"IamRoleArn,omitempty"`
	Resourcearn interface{} `json:"ResourceArn,omitempty"`
	Lifecycle interface{} `json:"Lifecycle,omitempty"`
	Createdby interface{} `json:"CreatedBy,omitempty"`
	Lastrestoretime interface{} `json:"LastRestoreTime,omitempty"`
	Resourcetype interface{} `json:"ResourceType,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Storageclass interface{} `json:"StorageClass,omitempty"`
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Isencrypted interface{} `json:"IsEncrypted,omitempty"`
	Backupvaultarn interface{} `json:"BackupVaultArn,omitempty"`
}

// AwsBackupBackupVaultNotificationsDetails represents the AwsBackupBackupVaultNotificationsDetails schema from the OpenAPI specification
type AwsBackupBackupVaultNotificationsDetails struct {
	Backupvaultevents interface{} `json:"BackupVaultEvents,omitempty"`
	Snstopicarn interface{} `json:"SnsTopicArn,omitempty"`
}

// AwsS3BucketServerSideEncryptionConfiguration represents the AwsS3BucketServerSideEncryptionConfiguration schema from the OpenAPI specification
type AwsS3BucketServerSideEncryptionConfiguration struct {
	Rules interface{} `json:"Rules,omitempty"`
}

// ContainerDetails represents the ContainerDetails schema from the OpenAPI specification
type ContainerDetails struct {
	Name interface{} `json:"Name,omitempty"`
	Privileged interface{} `json:"Privileged,omitempty"`
	Volumemounts interface{} `json:"VolumeMounts,omitempty"`
	Containerruntime interface{} `json:"ContainerRuntime,omitempty"`
	Imageid interface{} `json:"ImageId,omitempty"`
	Imagename interface{} `json:"ImageName,omitempty"`
	Launchedat interface{} `json:"LaunchedAt,omitempty"`
}

// Recommendation represents the Recommendation schema from the OpenAPI specification
type Recommendation struct {
	Text interface{} `json:"Text,omitempty"`
	Url interface{} `json:"Url,omitempty"`
}

// AwsRdsEventSubscriptionDetails represents the AwsRdsEventSubscriptionDetails schema from the OpenAPI specification
type AwsRdsEventSubscriptionDetails struct {
	Snstopicarn interface{} `json:"SnsTopicArn,omitempty"`
	Sourceidslist interface{} `json:"SourceIdsList,omitempty"`
	Customerawsid interface{} `json:"CustomerAwsId,omitempty"`
	Sourcetype interface{} `json:"SourceType,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Custsubscriptionid interface{} `json:"CustSubscriptionId,omitempty"`
	Eventsubscriptionarn interface{} `json:"EventSubscriptionArn,omitempty"`
	Subscriptioncreationtime interface{} `json:"SubscriptionCreationTime,omitempty"`
	Enabled interface{} `json:"Enabled,omitempty"`
	Eventcategorieslist interface{} `json:"EventCategoriesList,omitempty"`
}

// AwsBackupBackupPlanBackupPlanDetails represents the AwsBackupBackupPlanBackupPlanDetails schema from the OpenAPI specification
type AwsBackupBackupPlanBackupPlanDetails struct {
	Advancedbackupsettings interface{} `json:"AdvancedBackupSettings,omitempty"`
	Backupplanname interface{} `json:"BackupPlanName,omitempty"`
	Backupplanrule interface{} `json:"BackupPlanRule,omitempty"`
}

// AwsEcsServiceServiceRegistriesDetails represents the AwsEcsServiceServiceRegistriesDetails schema from the OpenAPI specification
type AwsEcsServiceServiceRegistriesDetails struct {
	Containername interface{} `json:"ContainerName,omitempty"`
	Containerport interface{} `json:"ContainerPort,omitempty"`
	Port interface{} `json:"Port,omitempty"`
	Registryarn interface{} `json:"RegistryArn,omitempty"`
}

// AwsEksClusterDetails represents the AwsEksClusterDetails schema from the OpenAPI specification
type AwsEksClusterDetails struct {
	Logging interface{} `json:"Logging,omitempty"`
	Certificateauthoritydata interface{} `json:"CertificateAuthorityData,omitempty"`
	Clusterstatus interface{} `json:"ClusterStatus,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Endpoint interface{} `json:"Endpoint,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Version interface{} `json:"Version,omitempty"`
	Resourcesvpcconfig interface{} `json:"ResourcesVpcConfig,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
}

// AwsEc2LaunchTemplateDataMaintenanceOptionsDetails represents the AwsEc2LaunchTemplateDataMaintenanceOptionsDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataMaintenanceOptionsDetails struct {
	Autorecovery interface{} `json:"AutoRecovery,omitempty"`
}

// AwsRdsDbInstanceEndpoint represents the AwsRdsDbInstanceEndpoint schema from the OpenAPI specification
type AwsRdsDbInstanceEndpoint struct {
	Address interface{} `json:"Address,omitempty"`
	Hostedzoneid interface{} `json:"HostedZoneId,omitempty"`
	Port interface{} `json:"Port,omitempty"`
}

// AwsMountPoint represents the AwsMountPoint schema from the OpenAPI specification
type AwsMountPoint struct {
	Sourcevolume interface{} `json:"SourceVolume,omitempty"`
	Containerpath interface{} `json:"ContainerPath,omitempty"`
}

// AwsIamUserDetails represents the AwsIamUserDetails schema from the OpenAPI specification
type AwsIamUserDetails struct {
	Userpolicylist interface{} `json:"UserPolicyList,omitempty"`
	Attachedmanagedpolicies interface{} `json:"AttachedManagedPolicies,omitempty"`
	Createdate interface{} `json:"CreateDate,omitempty"`
	Grouplist interface{} `json:"GroupList,omitempty"`
	Path interface{} `json:"Path,omitempty"`
	Permissionsboundary interface{} `json:"PermissionsBoundary,omitempty"`
	Userid interface{} `json:"UserId,omitempty"`
	Username interface{} `json:"UserName,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsEnvironmentDetails represents the AwsEcsTaskDefinitionContainerDefinitionsEnvironmentDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsEnvironmentDetails struct {
	Name interface{} `json:"Name,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// AwsCodeBuildProjectEnvironment represents the AwsCodeBuildProjectEnvironment schema from the OpenAPI specification
type AwsCodeBuildProjectEnvironment struct {
	Privilegedmode interface{} `json:"PrivilegedMode,omitempty"`
	Registrycredential interface{} `json:"RegistryCredential,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Certificate interface{} `json:"Certificate,omitempty"`
	Environmentvariables interface{} `json:"EnvironmentVariables,omitempty"`
	Imagepullcredentialstype interface{} `json:"ImagePullCredentialsType,omitempty"`
}

// AwsEfsAccessPointRootDirectoryDetails represents the AwsEfsAccessPointRootDirectoryDetails schema from the OpenAPI specification
type AwsEfsAccessPointRootDirectoryDetails struct {
	Creationinfo interface{} `json:"CreationInfo,omitempty"`
	Path interface{} `json:"Path,omitempty"`
}

// AwsWafRegionalRuleGroupRulesActionDetails represents the AwsWafRegionalRuleGroupRulesActionDetails schema from the OpenAPI specification
type AwsWafRegionalRuleGroupRulesActionDetails struct {
	TypeField interface{} `json:"Type,omitempty"`
}

// AwsAutoScalingAutoScalingGroupLaunchTemplateLaunchTemplateSpecification represents the AwsAutoScalingAutoScalingGroupLaunchTemplateLaunchTemplateSpecification schema from the OpenAPI specification
type AwsAutoScalingAutoScalingGroupLaunchTemplateLaunchTemplateSpecification struct {
	Launchtemplateid interface{} `json:"LaunchTemplateId,omitempty"`
	Launchtemplatename interface{} `json:"LaunchTemplateName,omitempty"`
	Version interface{} `json:"Version,omitempty"`
}

// AwsRdsDbInstanceDetails represents the AwsRdsDbInstanceDetails schema from the OpenAPI specification
type AwsRdsDbInstanceDetails struct {
	Dbinstanceclass interface{} `json:"DBInstanceClass,omitempty"`
	Backupretentionperiod interface{} `json:"BackupRetentionPeriod,omitempty"`
	Performanceinsightskmskeyid interface{} `json:"PerformanceInsightsKmsKeyId,omitempty"`
	Dbclusteridentifier interface{} `json:"DBClusterIdentifier,omitempty"`
	Dbinstancestatus interface{} `json:"DbInstanceStatus,omitempty"`
	Licensemodel interface{} `json:"LicenseModel,omitempty"`
	Dbinstanceidentifier interface{} `json:"DBInstanceIdentifier,omitempty"`
	Dbparametergroups interface{} `json:"DbParameterGroups,omitempty"`
	Iamdatabaseauthenticationenabled interface{} `json:"IAMDatabaseAuthenticationEnabled,omitempty"`
	Domainmemberships interface{} `json:"DomainMemberships,omitempty"`
	Pendingmodifiedvalues interface{} `json:"PendingModifiedValues,omitempty"`
	Dbname interface{} `json:"DBName,omitempty"`
	Engineversion interface{} `json:"EngineVersion,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Latestrestorabletime interface{} `json:"LatestRestorableTime,omitempty"`
	Promotiontier interface{} `json:"PromotionTier,omitempty"`
	Endpoint interface{} `json:"Endpoint,omitempty"`
	Optiongroupmemberships interface{} `json:"OptionGroupMemberships,omitempty"`
	Instancecreatetime interface{} `json:"InstanceCreateTime,omitempty"`
	Associatedroles interface{} `json:"AssociatedRoles,omitempty"`
	Listenerendpoint AwsRdsDbInstanceEndpoint `json:"ListenerEndpoint,omitempty"` // Specifies the connection endpoint.
	Multiaz interface{} `json:"MultiAz,omitempty"`
	Cacertificateidentifier interface{} `json:"CACertificateIdentifier,omitempty"`
	Availabilityzone interface{} `json:"AvailabilityZone,omitempty"`
	Dbsecuritygroups interface{} `json:"DbSecurityGroups,omitempty"`
	Autominorversionupgrade interface{} `json:"AutoMinorVersionUpgrade,omitempty"`
	Performanceinsightsenabled interface{} `json:"PerformanceInsightsEnabled,omitempty"`
	Dbsubnetgroup interface{} `json:"DbSubnetGroup,omitempty"`
	Dbiresourceid interface{} `json:"DbiResourceId,omitempty"`
	Readreplicadbclusteridentifiers interface{} `json:"ReadReplicaDBClusterIdentifiers,omitempty"`
	Monitoringrolearn interface{} `json:"MonitoringRoleArn,omitempty"`
	Preferredbackupwindow interface{} `json:"PreferredBackupWindow,omitempty"`
	Deletionprotection interface{} `json:"DeletionProtection,omitempty"`
	Masterusername interface{} `json:"MasterUsername,omitempty"`
	Charactersetname interface{} `json:"CharacterSetName,omitempty"`
	Enhancedmonitoringresourcearn interface{} `json:"EnhancedMonitoringResourceArn,omitempty"`
	Performanceinsightsretentionperiod interface{} `json:"PerformanceInsightsRetentionPeriod,omitempty"`
	Dbinstanceport interface{} `json:"DbInstancePort,omitempty"`
	Maxallocatedstorage interface{} `json:"MaxAllocatedStorage,omitempty"`
	Copytagstosnapshot interface{} `json:"CopyTagsToSnapshot,omitempty"`
	Secondaryavailabilityzone interface{} `json:"SecondaryAvailabilityZone,omitempty"`
	Preferredmaintenancewindow interface{} `json:"PreferredMaintenanceWindow,omitempty"`
	Storagetype interface{} `json:"StorageType,omitempty"`
	Enabledcloudwatchlogsexports interface{} `json:"EnabledCloudWatchLogsExports,omitempty"`
	Tdecredentialarn interface{} `json:"TdeCredentialArn,omitempty"`
	Iops interface{} `json:"Iops,omitempty"`
	Statusinfos interface{} `json:"StatusInfos,omitempty"`
	Allocatedstorage interface{} `json:"AllocatedStorage,omitempty"`
	Publiclyaccessible interface{} `json:"PubliclyAccessible,omitempty"`
	Readreplicasourcedbinstanceidentifier interface{} `json:"ReadReplicaSourceDBInstanceIdentifier,omitempty"`
	Timezone interface{} `json:"Timezone,omitempty"`
	Processorfeatures interface{} `json:"ProcessorFeatures,omitempty"`
	Monitoringinterval interface{} `json:"MonitoringInterval,omitempty"`
	Readreplicadbinstanceidentifiers interface{} `json:"ReadReplicaDBInstanceIdentifiers,omitempty"`
	Engine interface{} `json:"Engine,omitempty"`
	Storageencrypted interface{} `json:"StorageEncrypted,omitempty"`
	Vpcsecuritygroups interface{} `json:"VpcSecurityGroups,omitempty"`
}

// AcceptInvitationRequest represents the AcceptInvitationRequest schema from the OpenAPI specification
type AcceptInvitationRequest struct {
	Invitationid interface{} `json:"InvitationId"`
	Masterid interface{} `json:"MasterId"`
}

// AwsEc2NetworkAclAssociation represents the AwsEc2NetworkAclAssociation schema from the OpenAPI specification
type AwsEc2NetworkAclAssociation struct {
	Networkaclassociationid interface{} `json:"NetworkAclAssociationId,omitempty"`
	Networkaclid interface{} `json:"NetworkAclId,omitempty"`
	Subnetid interface{} `json:"SubnetId,omitempty"`
}

// AwsEcrRepositoryDetails represents the AwsEcrRepositoryDetails schema from the OpenAPI specification
type AwsEcrRepositoryDetails struct {
	Imagescanningconfiguration interface{} `json:"ImageScanningConfiguration,omitempty"`
	Imagetagmutability interface{} `json:"ImageTagMutability,omitempty"`
	Lifecyclepolicy interface{} `json:"LifecyclePolicy,omitempty"`
	Repositoryname interface{} `json:"RepositoryName,omitempty"`
	Repositorypolicytext interface{} `json:"RepositoryPolicyText,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
}

// Country represents the Country schema from the OpenAPI specification
type Country struct {
	Countrycode interface{} `json:"CountryCode,omitempty"`
	Countryname interface{} `json:"CountryName,omitempty"`
}

// AwsCloudWatchAlarmDetails represents the AwsCloudWatchAlarmDetails schema from the OpenAPI specification
type AwsCloudWatchAlarmDetails struct {
	Threshold interface{} `json:"Threshold,omitempty"`
	Actionsenabled interface{} `json:"ActionsEnabled,omitempty"`
	Okactions interface{} `json:"OkActions,omitempty"`
	Statistic interface{} `json:"Statistic,omitempty"`
	Alarmactions interface{} `json:"AlarmActions,omitempty"`
	Namespace interface{} `json:"Namespace,omitempty"`
	Evaluationperiods interface{} `json:"EvaluationPeriods,omitempty"`
	Thresholdmetricid interface{} `json:"ThresholdMetricId,omitempty"`
	Treatmissingdata interface{} `json:"TreatMissingData,omitempty"`
	Alarmconfigurationupdatedtimestamp interface{} `json:"AlarmConfigurationUpdatedTimestamp,omitempty"`
	Datapointstoalarm interface{} `json:"DatapointsToAlarm,omitempty"`
	Dimensions interface{} `json:"Dimensions,omitempty"`
	Unit interface{} `json:"Unit,omitempty"`
	Alarmname interface{} `json:"AlarmName,omitempty"`
	Evaluatelowsamplecountpercentile interface{} `json:"EvaluateLowSampleCountPercentile,omitempty"`
	Insufficientdataactions interface{} `json:"InsufficientDataActions,omitempty"`
	Alarmarn interface{} `json:"AlarmArn,omitempty"`
	Metricname interface{} `json:"MetricName,omitempty"`
	Comparisonoperator interface{} `json:"ComparisonOperator,omitempty"`
	Extendedstatistic interface{} `json:"ExtendedStatistic,omitempty"`
	Alarmdescription interface{} `json:"AlarmDescription,omitempty"`
	Period interface{} `json:"Period,omitempty"`
}

// AwsEcsTaskDefinitionVolumesEfsVolumeConfigurationDetails represents the AwsEcsTaskDefinitionVolumesEfsVolumeConfigurationDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionVolumesEfsVolumeConfigurationDetails struct {
	Transitencryptionport interface{} `json:"TransitEncryptionPort,omitempty"`
	Authorizationconfig interface{} `json:"AuthorizationConfig,omitempty"`
	Filesystemid interface{} `json:"FilesystemId,omitempty"`
	Rootdirectory interface{} `json:"RootDirectory,omitempty"`
	Transitencryption interface{} `json:"TransitEncryption,omitempty"`
}

// AwsEc2SecurityGroupDetails represents the AwsEc2SecurityGroupDetails schema from the OpenAPI specification
type AwsEc2SecurityGroupDetails struct {
	Vpcid interface{} `json:"VpcId,omitempty"`
	Groupid interface{} `json:"GroupId,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
	Ippermissions interface{} `json:"IpPermissions,omitempty"`
	Ippermissionsegress interface{} `json:"IpPermissionsEgress,omitempty"`
	Ownerid interface{} `json:"OwnerId,omitempty"`
}

// AwsEc2LaunchTemplateDataEnclaveOptionsDetails represents the AwsEc2LaunchTemplateDataEnclaveOptionsDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataEnclaveOptionsDetails struct {
	Enabled interface{} `json:"Enabled,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsPortMappingsDetails represents the AwsEcsTaskDefinitionContainerDefinitionsPortMappingsDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsPortMappingsDetails struct {
	Containerport interface{} `json:"ContainerPort,omitempty"`
	Hostport interface{} `json:"HostPort,omitempty"`
	Protocol interface{} `json:"Protocol,omitempty"`
}

// IpFilter represents the IpFilter schema from the OpenAPI specification
type IpFilter struct {
	Cidr interface{} `json:"Cidr,omitempty"`
}

// GetAdministratorAccountResponse represents the GetAdministratorAccountResponse schema from the OpenAPI specification
type GetAdministratorAccountResponse struct {
	Administrator Invitation `json:"Administrator,omitempty"` // Details about an invitation.
}

// FindingHistoryUpdate represents the FindingHistoryUpdate schema from the OpenAPI specification
type FindingHistoryUpdate struct {
	Newvalue interface{} `json:"NewValue,omitempty"`
	Oldvalue interface{} `json:"OldValue,omitempty"`
	Updatedfield interface{} `json:"UpdatedField,omitempty"`
}

// AwsEc2SecurityGroupIpv6Range represents the AwsEc2SecurityGroupIpv6Range schema from the OpenAPI specification
type AwsEc2SecurityGroupIpv6Range struct {
	Cidripv6 interface{} `json:"CidrIpv6,omitempty"`
}

// DisableImportFindingsForProductResponse represents the DisableImportFindingsForProductResponse schema from the OpenAPI specification
type DisableImportFindingsForProductResponse struct {
}

// AwsEcsTaskDefinitionVolumesDetails represents the AwsEcsTaskDefinitionVolumesDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionVolumesDetails struct {
	Efsvolumeconfiguration interface{} `json:"EfsVolumeConfiguration,omitempty"`
	Host interface{} `json:"Host,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Dockervolumeconfiguration interface{} `json:"DockerVolumeConfiguration,omitempty"`
}

// AwsWafWebAclDetails represents the AwsWafWebAclDetails schema from the OpenAPI specification
type AwsWafWebAclDetails struct {
	Rules interface{} `json:"Rules,omitempty"`
	Webaclid interface{} `json:"WebAclId,omitempty"`
	Defaultaction interface{} `json:"DefaultAction,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// SecurityControlDefinition represents the SecurityControlDefinition schema from the OpenAPI specification
type SecurityControlDefinition struct {
	Description interface{} `json:"Description"`
	Remediationurl interface{} `json:"RemediationUrl"`
	Securitycontrolid interface{} `json:"SecurityControlId"`
	Severityrating interface{} `json:"SeverityRating"`
	Title interface{} `json:"Title"`
	Currentregionavailability interface{} `json:"CurrentRegionAvailability"`
}

// AwsCodeBuildProjectArtifactsDetails represents the AwsCodeBuildProjectArtifactsDetails schema from the OpenAPI specification
type AwsCodeBuildProjectArtifactsDetails struct {
	Encryptiondisabled interface{} `json:"EncryptionDisabled,omitempty"`
	Path interface{} `json:"Path,omitempty"`
	Artifactidentifier interface{} `json:"ArtifactIdentifier,omitempty"`
	Location interface{} `json:"Location,omitempty"`
	Namespacetype interface{} `json:"NamespaceType,omitempty"`
	Overrideartifactname interface{} `json:"OverrideArtifactName,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Packaging interface{} `json:"Packaging,omitempty"`
}

// AwsS3BucketBucketLifecycleConfigurationRulesFilterDetails represents the AwsS3BucketBucketLifecycleConfigurationRulesFilterDetails schema from the OpenAPI specification
type AwsS3BucketBucketLifecycleConfigurationRulesFilterDetails struct {
	Predicate interface{} `json:"Predicate,omitempty"`
}

// RuleGroupDetails represents the RuleGroupDetails schema from the OpenAPI specification
type RuleGroupDetails struct {
	Rulevariables interface{} `json:"RuleVariables,omitempty"`
	Rulessource interface{} `json:"RulesSource,omitempty"`
}

// DescribeStandardsRequest represents the DescribeStandardsRequest schema from the OpenAPI specification
type DescribeStandardsRequest struct {
}

// LoadBalancerState represents the LoadBalancerState schema from the OpenAPI specification
type LoadBalancerState struct {
	Code interface{} `json:"Code,omitempty"`
	Reason interface{} `json:"Reason,omitempty"`
}

// AwsIamPolicyVersion represents the AwsIamPolicyVersion schema from the OpenAPI specification
type AwsIamPolicyVersion struct {
	Isdefaultversion interface{} `json:"IsDefaultVersion,omitempty"`
	Versionid interface{} `json:"VersionId,omitempty"`
	Createdate interface{} `json:"CreateDate,omitempty"`
}

// AwsElbLoadBalancerAdditionalAttribute represents the AwsElbLoadBalancerAdditionalAttribute schema from the OpenAPI specification
type AwsElbLoadBalancerAdditionalAttribute struct {
	Key interface{} `json:"Key,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// DnsRequestAction represents the DnsRequestAction schema from the OpenAPI specification
type DnsRequestAction struct {
	Domain interface{} `json:"Domain,omitempty"`
	Protocol interface{} `json:"Protocol,omitempty"`
	Blocked interface{} `json:"Blocked,omitempty"`
}

// ResourceDetails represents the ResourceDetails schema from the OpenAPI specification
type ResourceDetails struct {
	Awslambdafunction interface{} `json:"AwsLambdaFunction,omitempty"`
	Awselasticsearchdomain interface{} `json:"AwsElasticsearchDomain,omitempty"`
	Awselbv2loadbalancer interface{} `json:"AwsElbv2LoadBalancer,omitempty"`
	Awsec2vpnconnection interface{} `json:"AwsEc2VpnConnection,omitempty"`
	Awsiamgroup interface{} `json:"AwsIamGroup,omitempty"`
	Awsiampolicy interface{} `json:"AwsIamPolicy,omitempty"`
	Awswafrulegroup interface{} `json:"AwsWafRuleGroup,omitempty"`
	Awsautoscalingautoscalinggroup interface{} `json:"AwsAutoScalingAutoScalingGroup,omitempty"`
	Awsamazonmqbroker interface{} `json:"AwsAmazonMqBroker,omitempty"`
	Awsekscluster interface{} `json:"AwsEksCluster,omitempty"`
	Awsecstask interface{} `json:"AwsEcsTask,omitempty"`
	Awsnetworkfirewallfirewallpolicy interface{} `json:"AwsNetworkFirewallFirewallPolicy,omitempty"`
	Awsec2instance interface{} `json:"AwsEc2Instance,omitempty"`
	Awswafregionalratebasedrule interface{} `json:"AwsWafRegionalRateBasedRule,omitempty"`
	Awskmskey interface{} `json:"AwsKmsKey,omitempty"`
	Awsopensearchservicedomain interface{} `json:"AwsOpenSearchServiceDomain,omitempty"`
	Awsxrayencryptionconfig interface{} `json:"AwsXrayEncryptionConfig,omitempty"`
	Awss3object interface{} `json:"AwsS3Object,omitempty"`
	Awswafregionalwebacl interface{} `json:"AwsWafRegionalWebAcl,omitempty"`
	Awswafwebacl interface{} `json:"AwsWafWebAcl,omitempty"`
	Awsssmpatchcompliance interface{} `json:"AwsSsmPatchCompliance,omitempty"`
	Awsec2launchtemplate AwsEc2LaunchTemplateDetails `json:"AwsEc2LaunchTemplate,omitempty"` // Specifies the properties for creating an Amazon Elastic Compute Cloud (Amazon EC2) launch template.
	Awsredshiftcluster interface{} `json:"AwsRedshiftCluster,omitempty"`
	Awsrdsdbsnapshot interface{} `json:"AwsRdsDbSnapshot,omitempty"`
	Container interface{} `json:"Container,omitempty"`
	Awsiamuser interface{} `json:"AwsIamUser,omitempty"`
	Awscloudfrontdistribution interface{} `json:"AwsCloudFrontDistribution,omitempty"`
	Awsbackuprecoverypoint interface{} `json:"AwsBackupRecoveryPoint,omitempty"`
	Awsec2transitgateway interface{} `json:"AwsEc2TransitGateway,omitempty"`
	Awselasticbeanstalkenvironment interface{} `json:"AwsElasticBeanstalkEnvironment,omitempty"`
	Awssagemakernotebookinstance AwsSageMakerNotebookInstanceDetails `json:"AwsSageMakerNotebookInstance,omitempty"` // Provides details about an Amazon SageMaker notebook instance.
	Awss3bucket interface{} `json:"AwsS3Bucket,omitempty"`
	Awsapigatewaystage interface{} `json:"AwsApiGatewayStage,omitempty"`
	Awsbackupbackupplan interface{} `json:"AwsBackupBackupPlan,omitempty"`
	Awsec2vpcendpointservice interface{} `json:"AwsEc2VpcEndpointService,omitempty"`
	Awsiamaccesskey interface{} `json:"AwsIamAccessKey,omitempty"`
	Awss3accountpublicaccessblock interface{} `json:"AwsS3AccountPublicAccessBlock,omitempty"`
	Awswafv2webacl AwsWafv2WebAclDetails `json:"AwsWafv2WebAcl,omitempty"` // Details about an WAFv2 web Access Control List (ACL).
	Awswafv2rulegroup AwsWafv2RuleGroupDetails `json:"AwsWafv2RuleGroup,omitempty"` // Details about an WAFv2 rule group.
	Other interface{} `json:"Other,omitempty"`
	Awsguarddutydetector interface{} `json:"AwsGuardDutyDetector,omitempty"`
	Awsappsyncgraphqlapi interface{} `json:"AwsAppSyncGraphQlApi,omitempty"`
	Awsec2securitygroup interface{} `json:"AwsEc2SecurityGroup,omitempty"`
	Awsapigatewayv2api interface{} `json:"AwsApiGatewayV2Api,omitempty"`
	Awsecrcontainerimage interface{} `json:"AwsEcrContainerImage,omitempty"`
	Awskinesisstream interface{} `json:"AwsKinesisStream,omitempty"`
	Awsrdsdbcluster interface{} `json:"AwsRdsDbCluster,omitempty"`
	Awsecrrepository interface{} `json:"AwsEcrRepository,omitempty"`
	Awsecsservice interface{} `json:"AwsEcsService,omitempty"`
	Awsec2vpcpeeringconnection interface{} `json:"AwsEc2VpcPeeringConnection,omitempty"`
	Awswafrule interface{} `json:"AwsWafRule,omitempty"`
	Awswafregionalrulegroup interface{} `json:"AwsWafRegionalRuleGroup,omitempty"`
	Awsec2subnet interface{} `json:"AwsEc2Subnet,omitempty"`
	Awsefsaccesspoint interface{} `json:"AwsEfsAccessPoint,omitempty"`
	Awsrdsdbsecuritygroup interface{} `json:"AwsRdsDbSecurityGroup,omitempty"`
	Awscloudtrailtrail interface{} `json:"AwsCloudTrailTrail,omitempty"`
	Awsrdseventsubscription interface{} `json:"AwsRdsEventSubscription,omitempty"`
	Awsec2networkinterface interface{} `json:"AwsEc2NetworkInterface,omitempty"`
	Awslambdalayerversion interface{} `json:"AwsLambdaLayerVersion,omitempty"`
	Awsautoscalinglaunchconfiguration interface{} `json:"AwsAutoScalingLaunchConfiguration,omitempty"`
	Awsdynamodbtable interface{} `json:"AwsDynamoDbTable,omitempty"`
	Awsec2eip interface{} `json:"AwsEc2Eip,omitempty"`
	Awsec2routetable interface{} `json:"AwsEc2RouteTable,omitempty"`
	Awssqsqueue interface{} `json:"AwsSqsQueue,omitempty"`
	Awselbloadbalancer interface{} `json:"AwsElbLoadBalancer,omitempty"`
	Awssnstopic interface{} `json:"AwsSnsTopic,omitempty"`
	Awscodebuildproject interface{} `json:"AwsCodeBuildProject,omitempty"`
	Awsec2volume interface{} `json:"AwsEc2Volume,omitempty"`
	Awsecstaskdefinition interface{} `json:"AwsEcsTaskDefinition,omitempty"`
	Awsstepfunctionstatemachine interface{} `json:"AwsStepFunctionStateMachine,omitempty"`
	Awsnetworkfirewallfirewall interface{} `json:"AwsNetworkFirewallFirewall,omitempty"`
	Awsbackupbackupvault interface{} `json:"AwsBackupBackupVault,omitempty"`
	Awsec2vpc interface{} `json:"AwsEc2Vpc,omitempty"`
	Awscloudwatchalarm interface{} `json:"AwsCloudWatchAlarm,omitempty"`
	Awssecretsmanagersecret interface{} `json:"AwsSecretsManagerSecret,omitempty"`
	Awsapigatewayv2stage interface{} `json:"AwsApiGatewayV2Stage,omitempty"`
	Awsrdsdbclustersnapshot interface{} `json:"AwsRdsDbClusterSnapshot,omitempty"`
	Awsnetworkfirewallrulegroup interface{} `json:"AwsNetworkFirewallRuleGroup,omitempty"`
	Awsapigatewayrestapi interface{} `json:"AwsApiGatewayRestApi,omitempty"`
	Awswafratebasedrule interface{} `json:"AwsWafRateBasedRule,omitempty"`
	Awseventschemasregistry interface{} `json:"AwsEventSchemasRegistry,omitempty"`
	Awscloudformationstack interface{} `json:"AwsCloudFormationStack,omitempty"`
	Awsecscontainer interface{} `json:"AwsEcsContainer,omitempty"`
	Awswafregionalrule interface{} `json:"AwsWafRegionalRule,omitempty"`
	Awsec2networkacl interface{} `json:"AwsEc2NetworkAcl,omitempty"`
	Awsecscluster interface{} `json:"AwsEcsCluster,omitempty"`
	Awsiamrole interface{} `json:"AwsIamRole,omitempty"`
	Awscertificatemanagercertificate interface{} `json:"AwsCertificateManagerCertificate,omitempty"`
	Awsathenaworkgroup interface{} `json:"AwsAthenaWorkGroup,omitempty"`
	Awsrdsdbinstance interface{} `json:"AwsRdsDbInstance,omitempty"`
}

// AwsEc2LaunchTemplateDataMetadataOptionsDetails represents the AwsEc2LaunchTemplateDataMetadataOptionsDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataMetadataOptionsDetails struct {
	Httpputresponsehoplimit interface{} `json:"HttpPutResponseHopLimit,omitempty"`
	Httptokens interface{} `json:"HttpTokens,omitempty"`
	Instancemetadatatags interface{} `json:"InstanceMetadataTags,omitempty"`
	Httpendpoint interface{} `json:"HttpEndpoint,omitempty"`
	Httpprotocolipv6 interface{} `json:"HttpProtocolIpv6,omitempty"`
}

// EnableImportFindingsForProductRequest represents the EnableImportFindingsForProductRequest schema from the OpenAPI specification
type EnableImportFindingsForProductRequest struct {
	Productarn interface{} `json:"ProductArn"`
}

// DateRange represents the DateRange schema from the OpenAPI specification
type DateRange struct {
	Unit interface{} `json:"Unit,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// AwsBackupRecoveryPointCalculatedLifecycleDetails represents the AwsBackupRecoveryPointCalculatedLifecycleDetails schema from the OpenAPI specification
type AwsBackupRecoveryPointCalculatedLifecycleDetails struct {
	Deleteat interface{} `json:"DeleteAt,omitempty"`
	Movetocoldstorageat interface{} `json:"MoveToColdStorageAt,omitempty"`
}

// StandardsControlAssociationUpdate represents the StandardsControlAssociationUpdate schema from the OpenAPI specification
type StandardsControlAssociationUpdate struct {
	Associationstatus interface{} `json:"AssociationStatus"`
	Securitycontrolid interface{} `json:"SecurityControlId"`
	Standardsarn interface{} `json:"StandardsArn"`
	Updatedreason interface{} `json:"UpdatedReason,omitempty"`
}

// ListStandardsControlAssociationsResponse represents the ListStandardsControlAssociationsResponse schema from the OpenAPI specification
type ListStandardsControlAssociationsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Standardscontrolassociationsummaries interface{} `json:"StandardsControlAssociationSummaries"`
}

// AwsRdsDbInstanceVpcSecurityGroup represents the AwsRdsDbInstanceVpcSecurityGroup schema from the OpenAPI specification
type AwsRdsDbInstanceVpcSecurityGroup struct {
	Status interface{} `json:"Status,omitempty"`
	Vpcsecuritygroupid interface{} `json:"VpcSecurityGroupId,omitempty"`
}

// AwsWafv2WebAclActionDetails represents the AwsWafv2WebAclActionDetails schema from the OpenAPI specification
type AwsWafv2WebAclActionDetails struct {
	Allow interface{} `json:"Allow,omitempty"`
	Block interface{} `json:"Block,omitempty"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// BatchGetStandardsControlAssociationsRequest represents the BatchGetStandardsControlAssociationsRequest schema from the OpenAPI specification
type BatchGetStandardsControlAssociationsRequest struct {
	Standardscontrolassociationids interface{} `json:"StandardsControlAssociationIds"`
}

// Compliance represents the Compliance schema from the OpenAPI specification
type Compliance struct {
	Status interface{} `json:"Status,omitempty"`
	Statusreasons interface{} `json:"StatusReasons,omitempty"`
	Associatedstandards interface{} `json:"AssociatedStandards,omitempty"`
	Relatedrequirements interface{} `json:"RelatedRequirements,omitempty"`
	Securitycontrolid interface{} `json:"SecurityControlId,omitempty"`
}

// ClassificationStatus represents the ClassificationStatus schema from the OpenAPI specification
type ClassificationStatus struct {
	Reason interface{} `json:"Reason,omitempty"`
	Code interface{} `json:"Code,omitempty"`
}

// AwsGuardDutyDetectorFeaturesDetails represents the AwsGuardDutyDetectorFeaturesDetails schema from the OpenAPI specification
type AwsGuardDutyDetectorFeaturesDetails struct {
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// Invitation represents the Invitation schema from the OpenAPI specification
type Invitation struct {
	Accountid interface{} `json:"AccountId,omitempty"`
	Invitationid interface{} `json:"InvitationId,omitempty"`
	Invitedat interface{} `json:"InvitedAt,omitempty"`
	Memberstatus interface{} `json:"MemberStatus,omitempty"`
}

// ActionRemoteIpDetails represents the ActionRemoteIpDetails schema from the OpenAPI specification
type ActionRemoteIpDetails struct {
	Country interface{} `json:"Country,omitempty"`
	Geolocation interface{} `json:"GeoLocation,omitempty"`
	Ipaddressv4 interface{} `json:"IpAddressV4,omitempty"`
	Organization interface{} `json:"Organization,omitempty"`
	City interface{} `json:"City,omitempty"`
}

// AwsEcsServiceNetworkConfigurationAwsVpcConfigurationDetails represents the AwsEcsServiceNetworkConfigurationAwsVpcConfigurationDetails schema from the OpenAPI specification
type AwsEcsServiceNetworkConfigurationAwsVpcConfigurationDetails struct {
	Assignpublicip interface{} `json:"AssignPublicIp,omitempty"`
	Securitygroups interface{} `json:"SecurityGroups,omitempty"`
	Subnets interface{} `json:"Subnets,omitempty"`
}

// AwsEc2RouteTableDetails represents the AwsEc2RouteTableDetails schema from the OpenAPI specification
type AwsEc2RouteTableDetails struct {
	Vpcid interface{} `json:"VpcId,omitempty"`
	Associationset interface{} `json:"AssociationSet,omitempty"`
	Ownerid interface{} `json:"OwnerId,omitempty"`
	Propagatingvgwset interface{} `json:"PropagatingVgwSet,omitempty"`
	Routeset interface{} `json:"RouteSet,omitempty"`
	Routetableid interface{} `json:"RouteTableId,omitempty"`
}

// AwsElasticBeanstalkEnvironmentOptionSetting represents the AwsElasticBeanstalkEnvironmentOptionSetting schema from the OpenAPI specification
type AwsElasticBeanstalkEnvironmentOptionSetting struct {
	Namespace interface{} `json:"Namespace,omitempty"`
	Optionname interface{} `json:"OptionName,omitempty"`
	Resourcename interface{} `json:"ResourceName,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// DisassociateFromMasterAccountRequest represents the DisassociateFromMasterAccountRequest schema from the OpenAPI specification
type DisassociateFromMasterAccountRequest struct {
}

// AwsWafRegionalWebAclRulesListOverrideActionDetails represents the AwsWafRegionalWebAclRulesListOverrideActionDetails schema from the OpenAPI specification
type AwsWafRegionalWebAclRulesListOverrideActionDetails struct {
	TypeField interface{} `json:"Type,omitempty"`
}

// AwsS3BucketBucketLifecycleConfigurationRulesFilterPredicateTagDetails represents the AwsS3BucketBucketLifecycleConfigurationRulesFilterPredicateTagDetails schema from the OpenAPI specification
type AwsS3BucketBucketLifecycleConfigurationRulesFilterPredicateTagDetails struct {
	Value interface{} `json:"Value,omitempty"`
	Key interface{} `json:"Key,omitempty"`
}

// ListSecurityControlDefinitionsResponse represents the ListSecurityControlDefinitionsResponse schema from the OpenAPI specification
type ListSecurityControlDefinitionsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Securitycontroldefinitions interface{} `json:"SecurityControlDefinitions"`
}

// AwsEc2VpcPeeringConnectionDetails represents the AwsEc2VpcPeeringConnectionDetails schema from the OpenAPI specification
type AwsEc2VpcPeeringConnectionDetails struct {
	Requestervpcinfo interface{} `json:"RequesterVpcInfo,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Vpcpeeringconnectionid interface{} `json:"VpcPeeringConnectionId,omitempty"`
	Acceptervpcinfo interface{} `json:"AccepterVpcInfo,omitempty"`
	Expirationtime interface{} `json:"ExpirationTime,omitempty"`
}

// AwsAutoScalingLaunchConfigurationInstanceMonitoringDetails represents the AwsAutoScalingLaunchConfigurationInstanceMonitoringDetails schema from the OpenAPI specification
type AwsAutoScalingLaunchConfigurationInstanceMonitoringDetails struct {
	Enabled interface{} `json:"Enabled,omitempty"`
}

// Insight represents the Insight schema from the OpenAPI specification
type Insight struct {
	Insightarn interface{} `json:"InsightArn"`
	Name interface{} `json:"Name"`
	Filters interface{} `json:"Filters"`
	Groupbyattribute interface{} `json:"GroupByAttribute"`
}

// AwsRdsDbSecurityGroupIpRange represents the AwsRdsDbSecurityGroupIpRange schema from the OpenAPI specification
type AwsRdsDbSecurityGroupIpRange struct {
	Status interface{} `json:"Status,omitempty"`
	Cidrip interface{} `json:"CidrIp,omitempty"`
}

// FindingHistoryRecord represents the FindingHistoryRecord schema from the OpenAPI specification
type FindingHistoryRecord struct {
	Findingcreated interface{} `json:"FindingCreated,omitempty"`
	Findingidentifier AwsSecurityFindingIdentifier `json:"FindingIdentifier,omitempty"` // Identifies which finding to get the finding history for.
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Updatesource interface{} `json:"UpdateSource,omitempty"`
	Updatetime interface{} `json:"UpdateTime,omitempty"`
	Updates interface{} `json:"Updates,omitempty"`
}

// VpcInfoIpv6CidrBlockSetDetails represents the VpcInfoIpv6CidrBlockSetDetails schema from the OpenAPI specification
type VpcInfoIpv6CidrBlockSetDetails struct {
	Ipv6cidrblock interface{} `json:"Ipv6CidrBlock,omitempty"`
}

// AwsS3AccountPublicAccessBlockDetails represents the AwsS3AccountPublicAccessBlockDetails schema from the OpenAPI specification
type AwsS3AccountPublicAccessBlockDetails struct {
	Restrictpublicbuckets interface{} `json:"RestrictPublicBuckets,omitempty"`
	Blockpublicacls interface{} `json:"BlockPublicAcls,omitempty"`
	Blockpublicpolicy interface{} `json:"BlockPublicPolicy,omitempty"`
	Ignorepublicacls interface{} `json:"IgnorePublicAcls,omitempty"`
}

// AwsApiGatewayMethodSettings represents the AwsApiGatewayMethodSettings schema from the OpenAPI specification
type AwsApiGatewayMethodSettings struct {
	Httpmethod interface{} `json:"HttpMethod,omitempty"`
	Resourcepath interface{} `json:"ResourcePath,omitempty"`
	Cachedataencrypted interface{} `json:"CacheDataEncrypted,omitempty"`
	Throttlingburstlimit interface{} `json:"ThrottlingBurstLimit,omitempty"`
	Throttlingratelimit interface{} `json:"ThrottlingRateLimit,omitempty"`
	Cachettlinseconds interface{} `json:"CacheTtlInSeconds,omitempty"`
	Datatraceenabled interface{} `json:"DataTraceEnabled,omitempty"`
	Requireauthorizationforcachecontrol interface{} `json:"RequireAuthorizationForCacheControl,omitempty"`
	Logginglevel interface{} `json:"LoggingLevel,omitempty"`
	Metricsenabled interface{} `json:"MetricsEnabled,omitempty"`
	Unauthorizedcachecontrolheaderstrategy interface{} `json:"UnauthorizedCacheControlHeaderStrategy,omitempty"`
	Cachingenabled interface{} `json:"CachingEnabled,omitempty"`
}

// BatchImportFindingsResponse represents the BatchImportFindingsResponse schema from the OpenAPI specification
type BatchImportFindingsResponse struct {
	Failedfindings interface{} `json:"FailedFindings,omitempty"`
	Successcount interface{} `json:"SuccessCount"`
	Failedcount interface{} `json:"FailedCount"`
}

// AutomationRulesAction represents the AutomationRulesAction schema from the OpenAPI specification
type AutomationRulesAction struct {
	TypeField interface{} `json:"Type,omitempty"`
	Findingfieldsupdate interface{} `json:"FindingFieldsUpdate,omitempty"`
}

// AwsS3BucketObjectLockConfigurationRuleDefaultRetentionDetails represents the AwsS3BucketObjectLockConfigurationRuleDefaultRetentionDetails schema from the OpenAPI specification
type AwsS3BucketObjectLockConfigurationRuleDefaultRetentionDetails struct {
	Days interface{} `json:"Days,omitempty"`
	Mode interface{} `json:"Mode,omitempty"`
	Years interface{} `json:"Years,omitempty"`
}

// AutomationRulesMetadata represents the AutomationRulesMetadata schema from the OpenAPI specification
type AutomationRulesMetadata struct {
	Updatedat interface{} `json:"UpdatedAt,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Createdby interface{} `json:"CreatedBy,omitempty"`
	Isterminal interface{} `json:"IsTerminal,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Rulearn interface{} `json:"RuleArn,omitempty"`
	Ruleorder interface{} `json:"RuleOrder,omitempty"`
	Rulestatus interface{} `json:"RuleStatus,omitempty"`
	Rulename interface{} `json:"RuleName,omitempty"`
}

// AwsWafv2ActionBlockDetails represents the AwsWafv2ActionBlockDetails schema from the OpenAPI specification
type AwsWafv2ActionBlockDetails struct {
	Customresponse interface{} `json:"CustomResponse,omitempty"`
}

// Record represents the Record schema from the OpenAPI specification
type Record struct {
	Jsonpath interface{} `json:"JsonPath,omitempty"`
	Recordindex interface{} `json:"RecordIndex,omitempty"`
}

// ListSecurityControlDefinitionsRequest represents the ListSecurityControlDefinitionsRequest schema from the OpenAPI specification
type ListSecurityControlDefinitionsRequest struct {
}

// AwsRdsDbClusterAssociatedRole represents the AwsRdsDbClusterAssociatedRole schema from the OpenAPI specification
type AwsRdsDbClusterAssociatedRole struct {
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// AwsWafv2WebAclCaptchaConfigImmunityTimePropertyDetails represents the AwsWafv2WebAclCaptchaConfigImmunityTimePropertyDetails schema from the OpenAPI specification
type AwsWafv2WebAclCaptchaConfigImmunityTimePropertyDetails struct {
	Immunitytime interface{} `json:"ImmunityTime,omitempty"`
}

// AwsEcsTaskDefinitionContainerDefinitionsLinuxParametersTmpfsDetails represents the AwsEcsTaskDefinitionContainerDefinitionsLinuxParametersTmpfsDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsLinuxParametersTmpfsDetails struct {
	Containerpath interface{} `json:"ContainerPath,omitempty"`
	Mountoptions interface{} `json:"MountOptions,omitempty"`
	Size interface{} `json:"Size,omitempty"`
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// AwsEcsTaskDefinitionContainerDefinitionsLinuxParametersDetails represents the AwsEcsTaskDefinitionContainerDefinitionsLinuxParametersDetails schema from the OpenAPI specification
type AwsEcsTaskDefinitionContainerDefinitionsLinuxParametersDetails struct {
	Swappiness interface{} `json:"Swappiness,omitempty"`
	Tmpfs interface{} `json:"Tmpfs,omitempty"`
	Capabilities interface{} `json:"Capabilities,omitempty"`
	Devices interface{} `json:"Devices,omitempty"`
	Initprocessenabled interface{} `json:"InitProcessEnabled,omitempty"`
	Maxswap interface{} `json:"MaxSwap,omitempty"`
	Sharedmemorysize interface{} `json:"SharedMemorySize,omitempty"`
}

// AwsS3BucketBucketVersioningConfiguration represents the AwsS3BucketBucketVersioningConfiguration schema from the OpenAPI specification
type AwsS3BucketBucketVersioningConfiguration struct {
	Status interface{} `json:"Status,omitempty"`
	Ismfadeleteenabled interface{} `json:"IsMfaDeleteEnabled,omitempty"`
}

// AwsEc2LaunchTemplateDataInstanceRequirementsMemoryGiBPerVCpuDetails represents the AwsEc2LaunchTemplateDataInstanceRequirementsMemoryGiBPerVCpuDetails schema from the OpenAPI specification
type AwsEc2LaunchTemplateDataInstanceRequirementsMemoryGiBPerVCpuDetails struct {
	Max interface{} `json:"Max,omitempty"`
	Min interface{} `json:"Min,omitempty"`
}

// AwsEcsTaskVolumeHostDetails represents the AwsEcsTaskVolumeHostDetails schema from the OpenAPI specification
type AwsEcsTaskVolumeHostDetails struct {
	Sourcepath interface{} `json:"SourcePath,omitempty"`
}
