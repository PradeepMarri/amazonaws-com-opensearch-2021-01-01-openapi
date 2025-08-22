package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// PurchaseReservedInstanceOfferingResponse represents the PurchaseReservedInstanceOfferingResponse schema from the OpenAPI specification
type PurchaseReservedInstanceOfferingResponse struct {
	Reservationname interface{} `json:"ReservationName,omitempty"`
	Reservedinstanceid interface{} `json:"ReservedInstanceId,omitempty"`
}

// AutoTuneDetails represents the AutoTuneDetails schema from the OpenAPI specification
type AutoTuneDetails struct {
	Scheduledautotunedetails interface{} `json:"ScheduledAutoTuneDetails,omitempty"`
}

// DescribeDryRunProgressRequest represents the DescribeDryRunProgressRequest schema from the OpenAPI specification
type DescribeDryRunProgressRequest struct {
}

// VPCDerivedInfo represents the VPCDerivedInfo schema from the OpenAPI specification
type VPCDerivedInfo struct {
	Vpcid interface{} `json:"VPCId,omitempty"`
	Availabilityzones interface{} `json:"AvailabilityZones,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
	Subnetids interface{} `json:"SubnetIds,omitempty"`
}

// ListScheduledActionsResponse represents the ListScheduledActionsResponse schema from the OpenAPI specification
type ListScheduledActionsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Scheduledactions interface{} `json:"ScheduledActions,omitempty"`
}

// AssociatePackageResponse represents the AssociatePackageResponse schema from the OpenAPI specification
type AssociatePackageResponse struct {
	Domainpackagedetails interface{} `json:"DomainPackageDetails,omitempty"`
}

// DomainPackageDetails represents the DomainPackageDetails schema from the OpenAPI specification
type DomainPackageDetails struct {
	Packageid interface{} `json:"PackageID,omitempty"`
	Packageversion interface{} `json:"PackageVersion,omitempty"`
	Referencepath interface{} `json:"ReferencePath,omitempty"`
	Packagename interface{} `json:"PackageName,omitempty"`
	Lastupdated interface{} `json:"LastUpdated,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Domainpackagestatus interface{} `json:"DomainPackageStatus,omitempty"`
	Packagetype interface{} `json:"PackageType,omitempty"`
	Errordetails interface{} `json:"ErrorDetails,omitempty"`
}

// DescribeDomainChangeProgressResponse represents the DescribeDomainChangeProgressResponse schema from the OpenAPI specification
type DescribeDomainChangeProgressResponse struct {
	Changeprogressstatus interface{} `json:"ChangeProgressStatus,omitempty"`
}

// CreatePackageRequest represents the CreatePackageRequest schema from the OpenAPI specification
type CreatePackageRequest struct {
	Packagedescription interface{} `json:"PackageDescription,omitempty"`
	Packagename interface{} `json:"PackageName"`
	Packagesource interface{} `json:"PackageSource"`
	Packagetype interface{} `json:"PackageType"`
}

// AutoTuneOptionsInput represents the AutoTuneOptionsInput schema from the OpenAPI specification
type AutoTuneOptionsInput struct {
	Useoffpeakwindow interface{} `json:"UseOffPeakWindow,omitempty"`
	Desiredstate interface{} `json:"DesiredState,omitempty"`
	Maintenanceschedules interface{} `json:"MaintenanceSchedules,omitempty"`
}

// DomainInfo represents the DomainInfo schema from the OpenAPI specification
type DomainInfo struct {
	Enginetype interface{} `json:"EngineType,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
}

// DeleteOutboundConnectionRequest represents the DeleteOutboundConnectionRequest schema from the OpenAPI specification
type DeleteOutboundConnectionRequest struct {
}

// OffPeakWindowOptionsStatus represents the OffPeakWindowOptionsStatus schema from the OpenAPI specification
type OffPeakWindowOptionsStatus struct {
	Options interface{} `json:"Options,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// DeleteDomainRequest represents the DeleteDomainRequest schema from the OpenAPI specification
type DeleteDomainRequest struct {
}

// GetUpgradeStatusRequest represents the GetUpgradeStatusRequest schema from the OpenAPI specification
type GetUpgradeStatusRequest struct {
}

// OutboundConnectionStatus represents the OutboundConnectionStatus schema from the OpenAPI specification
type OutboundConnectionStatus struct {
	Statuscode interface{} `json:"StatusCode,omitempty"`
	Message interface{} `json:"Message,omitempty"`
}

// LimitsByRole represents the LimitsByRole schema from the OpenAPI specification
type LimitsByRole struct {
}

// InstanceLimits represents the InstanceLimits schema from the OpenAPI specification
type InstanceLimits struct {
	Instancecountlimits interface{} `json:"InstanceCountLimits,omitempty"`
}

// PackageSource represents the PackageSource schema from the OpenAPI specification
type PackageSource struct {
	S3key interface{} `json:"S3Key,omitempty"`
	S3bucketname interface{} `json:"S3BucketName,omitempty"`
}

// DescribeDomainsRequest represents the DescribeDomainsRequest schema from the OpenAPI specification
type DescribeDomainsRequest struct {
	Domainnames interface{} `json:"DomainNames"`
}

// DescribeVpcEndpointsResponse represents the DescribeVpcEndpointsResponse schema from the OpenAPI specification
type DescribeVpcEndpointsResponse struct {
	Vpcendpointerrors interface{} `json:"VpcEndpointErrors"`
	Vpcendpoints interface{} `json:"VpcEndpoints"`
}

// DescribeInstanceTypeLimitsRequest represents the DescribeInstanceTypeLimitsRequest schema from the OpenAPI specification
type DescribeInstanceTypeLimitsRequest struct {
}

// ServiceSoftwareOptions represents the ServiceSoftwareOptions schema from the OpenAPI specification
type ServiceSoftwareOptions struct {
	Updateavailable interface{} `json:"UpdateAvailable,omitempty"`
	Updatestatus interface{} `json:"UpdateStatus,omitempty"`
	Automatedupdatedate interface{} `json:"AutomatedUpdateDate,omitempty"`
	Cancellable interface{} `json:"Cancellable,omitempty"`
	Currentversion interface{} `json:"CurrentVersion,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Newversion interface{} `json:"NewVersion,omitempty"`
	Optionaldeployment interface{} `json:"OptionalDeployment,omitempty"`
}

// DescribeDomainResponse represents the DescribeDomainResponse schema from the OpenAPI specification
type DescribeDomainResponse struct {
	Domainstatus interface{} `json:"DomainStatus"`
}

// DeleteVpcEndpointRequest represents the DeleteVpcEndpointRequest schema from the OpenAPI specification
type DeleteVpcEndpointRequest struct {
}

// NodeToNodeEncryptionOptionsStatus represents the NodeToNodeEncryptionOptionsStatus schema from the OpenAPI specification
type NodeToNodeEncryptionOptionsStatus struct {
	Status interface{} `json:"Status"`
	Options interface{} `json:"Options"`
}

// DescribeDomainChangeProgressRequest represents the DescribeDomainChangeProgressRequest schema from the OpenAPI specification
type DescribeDomainChangeProgressRequest struct {
}

// EBSOptionsStatus represents the EBSOptionsStatus schema from the OpenAPI specification
type EBSOptionsStatus struct {
	Options interface{} `json:"Options"`
	Status interface{} `json:"Status"`
}

// StorageTypeLimit represents the StorageTypeLimit schema from the OpenAPI specification
type StorageTypeLimit struct {
	Limitname interface{} `json:"LimitName,omitempty"`
	Limitvalues interface{} `json:"LimitValues,omitempty"`
}

// RevokeVpcEndpointAccessRequest represents the RevokeVpcEndpointAccessRequest schema from the OpenAPI specification
type RevokeVpcEndpointAccessRequest struct {
	Account interface{} `json:"Account"`
}

// DeleteDomainResponse represents the DeleteDomainResponse schema from the OpenAPI specification
type DeleteDomainResponse struct {
	Domainstatus interface{} `json:"DomainStatus,omitempty"`
}

// DescribeOutboundConnectionsRequest represents the DescribeOutboundConnectionsRequest schema from the OpenAPI specification
type DescribeOutboundConnectionsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// CancelServiceSoftwareUpdateRequest represents the CancelServiceSoftwareUpdateRequest schema from the OpenAPI specification
type CancelServiceSoftwareUpdateRequest struct {
	Domainname interface{} `json:"DomainName"`
}

// DescribePackagesRequest represents the DescribePackagesRequest schema from the OpenAPI specification
type DescribePackagesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
}

// GetPackageVersionHistoryResponse represents the GetPackageVersionHistoryResponse schema from the OpenAPI specification
type GetPackageVersionHistoryResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Packageid interface{} `json:"PackageID,omitempty"`
	Packageversionhistorylist interface{} `json:"PackageVersionHistoryList,omitempty"`
}

// UpdatePackageRequest represents the UpdatePackageRequest schema from the OpenAPI specification
type UpdatePackageRequest struct {
	Packagesource interface{} `json:"PackageSource"`
	Commitmessage interface{} `json:"CommitMessage,omitempty"`
	Packagedescription interface{} `json:"PackageDescription,omitempty"`
	Packageid interface{} `json:"PackageID"`
}

// DeleteVpcEndpointResponse represents the DeleteVpcEndpointResponse schema from the OpenAPI specification
type DeleteVpcEndpointResponse struct {
	Vpcendpointsummary interface{} `json:"VpcEndpointSummary"`
}

// SnapshotOptions represents the SnapshotOptions schema from the OpenAPI specification
type SnapshotOptions struct {
	Automatedsnapshotstarthour interface{} `json:"AutomatedSnapshotStartHour,omitempty"`
}

// AutoTuneOptionsOutput represents the AutoTuneOptionsOutput schema from the OpenAPI specification
type AutoTuneOptionsOutput struct {
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	State interface{} `json:"State,omitempty"`
	Useoffpeakwindow interface{} `json:"UseOffPeakWindow,omitempty"`
}

// PackageDetails represents the PackageDetails schema from the OpenAPI specification
type PackageDetails struct {
	Availablepackageversion interface{} `json:"AvailablePackageVersion,omitempty"`
	Lastupdatedat interface{} `json:"LastUpdatedAt,omitempty"`
	Errordetails interface{} `json:"ErrorDetails,omitempty"`
	Packagedescription interface{} `json:"PackageDescription,omitempty"`
	Packageid interface{} `json:"PackageID,omitempty"`
	Packagestatus interface{} `json:"PackageStatus,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Packagetype interface{} `json:"PackageType,omitempty"`
	Packagename interface{} `json:"PackageName,omitempty"`
}

// EndpointsMap represents the EndpointsMap schema from the OpenAPI specification
type EndpointsMap struct {
}

// ChangeProgressDetails represents the ChangeProgressDetails schema from the OpenAPI specification
type ChangeProgressDetails struct {
	Message interface{} `json:"Message,omitempty"`
	Changeid interface{} `json:"ChangeId,omitempty"`
}

// AutoTune represents the AutoTune schema from the OpenAPI specification
type AutoTune struct {
	Autotunedetails interface{} `json:"AutoTuneDetails,omitempty"`
	Autotunetype interface{} `json:"AutoTuneType,omitempty"`
}

// CreateOutboundConnectionRequest represents the CreateOutboundConnectionRequest schema from the OpenAPI specification
type CreateOutboundConnectionRequest struct {
	Connectionmode interface{} `json:"ConnectionMode,omitempty"`
	Connectionproperties interface{} `json:"ConnectionProperties,omitempty"`
	Localdomaininfo interface{} `json:"LocalDomainInfo"`
	Remotedomaininfo interface{} `json:"RemoteDomainInfo"`
	Connectionalias interface{} `json:"ConnectionAlias"`
}

// ClusterConfig represents the ClusterConfig schema from the OpenAPI specification
type ClusterConfig struct {
	Coldstorageoptions interface{} `json:"ColdStorageOptions,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Warmcount interface{} `json:"WarmCount,omitempty"`
	Zoneawarenessconfig interface{} `json:"ZoneAwarenessConfig,omitempty"`
	Dedicatedmastercount interface{} `json:"DedicatedMasterCount,omitempty"`
	Warmenabled interface{} `json:"WarmEnabled,omitempty"`
	Warmtype interface{} `json:"WarmType,omitempty"`
	Zoneawarenessenabled interface{} `json:"ZoneAwarenessEnabled,omitempty"`
	Dedicatedmasterenabled interface{} `json:"DedicatedMasterEnabled,omitempty"`
	Instancecount interface{} `json:"InstanceCount,omitempty"`
	Dedicatedmastertype interface{} `json:"DedicatedMasterType,omitempty"`
	Multiazwithstandbyenabled interface{} `json:"MultiAZWithStandbyEnabled,omitempty"`
}

// CreateOutboundConnectionResponse represents the CreateOutboundConnectionResponse schema from the OpenAPI specification
type CreateOutboundConnectionResponse struct {
	Connectionstatus interface{} `json:"ConnectionStatus,omitempty"`
	Localdomaininfo interface{} `json:"LocalDomainInfo,omitempty"`
	Remotedomaininfo interface{} `json:"RemoteDomainInfo,omitempty"`
	Connectionalias interface{} `json:"ConnectionAlias,omitempty"`
	Connectionid interface{} `json:"ConnectionId,omitempty"`
	Connectionmode interface{} `json:"ConnectionMode,omitempty"`
	Connectionproperties interface{} `json:"ConnectionProperties,omitempty"`
}

// ListVpcEndpointAccessResponse represents the ListVpcEndpointAccessResponse schema from the OpenAPI specification
type ListVpcEndpointAccessResponse struct {
	Authorizedprincipallist interface{} `json:"AuthorizedPrincipalList"`
	Nexttoken interface{} `json:"NextToken"`
}

// VPCOptions represents the VPCOptions schema from the OpenAPI specification
type VPCOptions struct {
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
	Subnetids interface{} `json:"SubnetIds,omitempty"`
}

// ListTagsResponse represents the ListTagsResponse schema from the OpenAPI specification
type ListTagsResponse struct {
	Taglist interface{} `json:"TagList,omitempty"`
}

// ListDomainsForPackageRequest represents the ListDomainsForPackageRequest schema from the OpenAPI specification
type ListDomainsForPackageRequest struct {
}

// ValidationFailure represents the ValidationFailure schema from the OpenAPI specification
type ValidationFailure struct {
	Code interface{} `json:"Code,omitempty"`
	Message interface{} `json:"Message,omitempty"`
}

// AdditionalLimit represents the AdditionalLimit schema from the OpenAPI specification
type AdditionalLimit struct {
	Limitvalues interface{} `json:"LimitValues,omitempty"`
	Limitname interface{} `json:"LimitName,omitempty"`
}

// ListDomainNamesResponse represents the ListDomainNamesResponse schema from the OpenAPI specification
type ListDomainNamesResponse struct {
	Domainnames interface{} `json:"DomainNames,omitempty"`
}

// RevokeVpcEndpointAccessResponse represents the RevokeVpcEndpointAccessResponse schema from the OpenAPI specification
type RevokeVpcEndpointAccessResponse struct {
}

// ReservedInstance represents the ReservedInstance schema from the OpenAPI specification
type ReservedInstance struct {
	Reservedinstanceid interface{} `json:"ReservedInstanceId,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Currencycode interface{} `json:"CurrencyCode,omitempty"`
	Recurringcharges interface{} `json:"RecurringCharges,omitempty"`
	State interface{} `json:"State,omitempty"`
	Usageprice interface{} `json:"UsagePrice,omitempty"`
	Duration interface{} `json:"Duration,omitempty"`
	Instancecount interface{} `json:"InstanceCount,omitempty"`
	Reservedinstanceofferingid interface{} `json:"ReservedInstanceOfferingId,omitempty"`
	Billingsubscriptionid interface{} `json:"BillingSubscriptionId,omitempty"`
	Fixedprice interface{} `json:"FixedPrice,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Paymentoption interface{} `json:"PaymentOption,omitempty"`
	Reservationname interface{} `json:"ReservationName,omitempty"`
}

// DescribeDomainAutoTunesResponse represents the DescribeDomainAutoTunesResponse schema from the OpenAPI specification
type DescribeDomainAutoTunesResponse struct {
	Autotunes interface{} `json:"AutoTunes,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeletePackageResponse represents the DeletePackageResponse schema from the OpenAPI specification
type DeletePackageResponse struct {
	Packagedetails interface{} `json:"PackageDetails,omitempty"`
}

// InstanceCountLimits represents the InstanceCountLimits schema from the OpenAPI specification
type InstanceCountLimits struct {
	Maximuminstancecount interface{} `json:"MaximumInstanceCount,omitempty"`
	Minimuminstancecount interface{} `json:"MinimumInstanceCount,omitempty"`
}

// ListVpcEndpointsRequest represents the ListVpcEndpointsRequest schema from the OpenAPI specification
type ListVpcEndpointsRequest struct {
}

// DomainEndpointOptions represents the DomainEndpointOptions schema from the OpenAPI specification
type DomainEndpointOptions struct {
	Tlssecuritypolicy interface{} `json:"TLSSecurityPolicy,omitempty"`
	Customendpoint interface{} `json:"CustomEndpoint,omitempty"`
	Customendpointcertificatearn interface{} `json:"CustomEndpointCertificateArn,omitempty"`
	Customendpointenabled interface{} `json:"CustomEndpointEnabled,omitempty"`
	Enforcehttps interface{} `json:"EnforceHTTPS,omitempty"`
}

// CreateDomainRequest represents the CreateDomainRequest schema from the OpenAPI specification
type CreateDomainRequest struct {
	Encryptionatrestoptions interface{} `json:"EncryptionAtRestOptions,omitempty"`
	Logpublishingoptions interface{} `json:"LogPublishingOptions,omitempty"`
	Advancedsecurityoptions interface{} `json:"AdvancedSecurityOptions,omitempty"`
	Autotuneoptions interface{} `json:"AutoTuneOptions,omitempty"`
	Accesspolicies interface{} `json:"AccessPolicies,omitempty"`
	Advancedoptions interface{} `json:"AdvancedOptions,omitempty"`
	Clusterconfig interface{} `json:"ClusterConfig,omitempty"`
	Engineversion interface{} `json:"EngineVersion,omitempty"`
	Offpeakwindowoptions interface{} `json:"OffPeakWindowOptions,omitempty"`
	Cognitooptions interface{} `json:"CognitoOptions,omitempty"`
	Nodetonodeencryptionoptions interface{} `json:"NodeToNodeEncryptionOptions,omitempty"`
	Taglist interface{} `json:"TagList,omitempty"`
	Domainname interface{} `json:"DomainName"`
	Domainendpointoptions interface{} `json:"DomainEndpointOptions,omitempty"`
	Vpcoptions interface{} `json:"VPCOptions,omitempty"`
	Snapshotoptions interface{} `json:"SnapshotOptions,omitempty"`
	Softwareupdateoptions interface{} `json:"SoftwareUpdateOptions,omitempty"`
	Ebsoptions interface{} `json:"EBSOptions,omitempty"`
}

// SAMLOptionsInput represents the SAMLOptionsInput schema from the OpenAPI specification
type SAMLOptionsInput struct {
	Sessiontimeoutminutes interface{} `json:"SessionTimeoutMinutes,omitempty"`
	Subjectkey interface{} `json:"SubjectKey,omitempty"`
	Enabled interface{} `json:"Enabled,omitempty"`
	Idp interface{} `json:"Idp,omitempty"`
	Masterbackendrole interface{} `json:"MasterBackendRole,omitempty"`
	Masterusername interface{} `json:"MasterUserName,omitempty"`
	Roleskey interface{} `json:"RolesKey,omitempty"`
}

// ListVpcEndpointsForDomainResponse represents the ListVpcEndpointsForDomainResponse schema from the OpenAPI specification
type ListVpcEndpointsForDomainResponse struct {
	Nexttoken interface{} `json:"NextToken"`
	Vpcendpointsummarylist interface{} `json:"VpcEndpointSummaryList"`
}

// DescribeDomainNodesRequest represents the DescribeDomainNodesRequest schema from the OpenAPI specification
type DescribeDomainNodesRequest struct {
}

// DissociatePackageRequest represents the DissociatePackageRequest schema from the OpenAPI specification
type DissociatePackageRequest struct {
}

// NodeToNodeEncryptionOptions represents the NodeToNodeEncryptionOptions schema from the OpenAPI specification
type NodeToNodeEncryptionOptions struct {
	Enabled interface{} `json:"Enabled,omitempty"`
}

// UpgradeStepItem represents the UpgradeStepItem schema from the OpenAPI specification
type UpgradeStepItem struct {
	Issues interface{} `json:"Issues,omitempty"`
	Progresspercent interface{} `json:"ProgressPercent,omitempty"`
	Upgradestep interface{} `json:"UpgradeStep,omitempty"`
	Upgradestepstatus interface{} `json:"UpgradeStepStatus,omitempty"`
}

// DescribeDryRunProgressResponse represents the DescribeDryRunProgressResponse schema from the OpenAPI specification
type DescribeDryRunProgressResponse struct {
	Dryrunconfig interface{} `json:"DryRunConfig,omitempty"`
	Dryrunprogressstatus interface{} `json:"DryRunProgressStatus,omitempty"`
	Dryrunresults interface{} `json:"DryRunResults,omitempty"`
}

// CreateDomainResponse represents the CreateDomainResponse schema from the OpenAPI specification
type CreateDomainResponse struct {
	Domainstatus interface{} `json:"DomainStatus,omitempty"`
}

// EBSOptions represents the EBSOptions schema from the OpenAPI specification
type EBSOptions struct {
	Volumetype interface{} `json:"VolumeType,omitempty"`
	Ebsenabled interface{} `json:"EBSEnabled,omitempty"`
	Iops interface{} `json:"Iops,omitempty"`
	Throughput interface{} `json:"Throughput,omitempty"`
	Volumesize interface{} `json:"VolumeSize,omitempty"`
}

// EncryptionAtRestOptions represents the EncryptionAtRestOptions schema from the OpenAPI specification
type EncryptionAtRestOptions struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
}

// DescribePackagesFilter represents the DescribePackagesFilter schema from the OpenAPI specification
type DescribePackagesFilter struct {
	Name interface{} `json:"Name,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// VersionStatus represents the VersionStatus schema from the OpenAPI specification
type VersionStatus struct {
	Options interface{} `json:"Options"`
	Status interface{} `json:"Status"`
}

// ListPackagesForDomainRequest represents the ListPackagesForDomainRequest schema from the OpenAPI specification
type ListPackagesForDomainRequest struct {
}

// ListTagsRequest represents the ListTagsRequest schema from the OpenAPI specification
type ListTagsRequest struct {
}

// OptionStatus represents the OptionStatus schema from the OpenAPI specification
type OptionStatus struct {
	Creationdate interface{} `json:"CreationDate"`
	Pendingdeletion interface{} `json:"PendingDeletion,omitempty"`
	State interface{} `json:"State"`
	Updatedate interface{} `json:"UpdateDate"`
	Updateversion interface{} `json:"UpdateVersion,omitempty"`
}

// OffPeakWindowOptions represents the OffPeakWindowOptions schema from the OpenAPI specification
type OffPeakWindowOptions struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Offpeakwindow interface{} `json:"OffPeakWindow,omitempty"`
}

// WindowStartTime represents the WindowStartTime schema from the OpenAPI specification
type WindowStartTime struct {
	Minutes interface{} `json:"Minutes"`
	Hours interface{} `json:"Hours"`
}

// DescribeInboundConnectionsRequest represents the DescribeInboundConnectionsRequest schema from the OpenAPI specification
type DescribeInboundConnectionsRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// AddTagsRequest represents the AddTagsRequest schema from the OpenAPI specification
type AddTagsRequest struct {
	Arn interface{} `json:"ARN"`
	Taglist interface{} `json:"TagList"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// AdvancedSecurityOptionsStatus represents the AdvancedSecurityOptionsStatus schema from the OpenAPI specification
type AdvancedSecurityOptionsStatus struct {
	Options interface{} `json:"Options"`
	Status interface{} `json:"Status"`
}

// UpdateDomainConfigRequest represents the UpdateDomainConfigRequest schema from the OpenAPI specification
type UpdateDomainConfigRequest struct {
	Accesspolicies interface{} `json:"AccessPolicies,omitempty"`
	Advancedoptions interface{} `json:"AdvancedOptions,omitempty"`
	Advancedsecurityoptions interface{} `json:"AdvancedSecurityOptions,omitempty"`
	Ebsoptions interface{} `json:"EBSOptions,omitempty"`
	Dryrun interface{} `json:"DryRun,omitempty"`
	Autotuneoptions interface{} `json:"AutoTuneOptions,omitempty"`
	Cognitooptions interface{} `json:"CognitoOptions,omitempty"`
	Dryrunmode interface{} `json:"DryRunMode,omitempty"`
	Clusterconfig interface{} `json:"ClusterConfig,omitempty"`
	Encryptionatrestoptions interface{} `json:"EncryptionAtRestOptions,omitempty"`
	Logpublishingoptions interface{} `json:"LogPublishingOptions,omitempty"`
	Snapshotoptions interface{} `json:"SnapshotOptions,omitempty"`
	Offpeakwindowoptions interface{} `json:"OffPeakWindowOptions,omitempty"`
	Vpcoptions interface{} `json:"VPCOptions,omitempty"`
	Softwareupdateoptions interface{} `json:"SoftwareUpdateOptions,omitempty"`
	Domainendpointoptions interface{} `json:"DomainEndpointOptions,omitempty"`
	Nodetonodeencryptionoptions interface{} `json:"NodeToNodeEncryptionOptions,omitempty"`
}

// EnvironmentInfo represents the EnvironmentInfo schema from the OpenAPI specification
type EnvironmentInfo struct {
	Availabilityzoneinformation interface{} `json:"AvailabilityZoneInformation,omitempty"`
}

// AssociatePackageRequest represents the AssociatePackageRequest schema from the OpenAPI specification
type AssociatePackageRequest struct {
}

// ListPackagesForDomainResponse represents the ListPackagesForDomainResponse schema from the OpenAPI specification
type ListPackagesForDomainResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Domainpackagedetailslist interface{} `json:"DomainPackageDetailsList,omitempty"`
}

// DeleteInboundConnectionResponse represents the DeleteInboundConnectionResponse schema from the OpenAPI specification
type DeleteInboundConnectionResponse struct {
	Connection interface{} `json:"Connection,omitempty"`
}

// PackageVersionHistory represents the PackageVersionHistory schema from the OpenAPI specification
type PackageVersionHistory struct {
	Packageversion interface{} `json:"PackageVersion,omitempty"`
	Commitmessage interface{} `json:"CommitMessage,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
}

// AdvancedSecurityOptions represents the AdvancedSecurityOptions schema from the OpenAPI specification
type AdvancedSecurityOptions struct {
	Anonymousauthdisabledate interface{} `json:"AnonymousAuthDisableDate,omitempty"`
	Anonymousauthenabled interface{} `json:"AnonymousAuthEnabled,omitempty"`
	Enabled interface{} `json:"Enabled,omitempty"`
	Internaluserdatabaseenabled interface{} `json:"InternalUserDatabaseEnabled,omitempty"`
	Samloptions interface{} `json:"SAMLOptions,omitempty"`
}

// RemoveTagsRequest represents the RemoveTagsRequest schema from the OpenAPI specification
type RemoveTagsRequest struct {
	Arn interface{} `json:"ARN"`
	Tagkeys interface{} `json:"TagKeys"`
}

// DeleteOutboundConnectionResponse represents the DeleteOutboundConnectionResponse schema from the OpenAPI specification
type DeleteOutboundConnectionResponse struct {
	Connection interface{} `json:"Connection,omitempty"`
}

// AcceptInboundConnectionResponse represents the AcceptInboundConnectionResponse schema from the OpenAPI specification
type AcceptInboundConnectionResponse struct {
	Connection interface{} `json:"Connection,omitempty"`
}

// ListInstanceTypeDetailsRequest represents the ListInstanceTypeDetailsRequest schema from the OpenAPI specification
type ListInstanceTypeDetailsRequest struct {
}

// ClusterConfigStatus represents the ClusterConfigStatus schema from the OpenAPI specification
type ClusterConfigStatus struct {
	Options interface{} `json:"Options"`
	Status interface{} `json:"Status"`
}

// CancelServiceSoftwareUpdateResponse represents the CancelServiceSoftwareUpdateResponse schema from the OpenAPI specification
type CancelServiceSoftwareUpdateResponse struct {
	Servicesoftwareoptions interface{} `json:"ServiceSoftwareOptions,omitempty"`
}

// ListVersionsRequest represents the ListVersionsRequest schema from the OpenAPI specification
type ListVersionsRequest struct {
}

// CreatePackageResponse represents the CreatePackageResponse schema from the OpenAPI specification
type CreatePackageResponse struct {
	Packagedetails interface{} `json:"PackageDetails,omitempty"`
}

// CreateVpcEndpointRequest represents the CreateVpcEndpointRequest schema from the OpenAPI specification
type CreateVpcEndpointRequest struct {
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Domainarn interface{} `json:"DomainArn"`
	Vpcoptions interface{} `json:"VpcOptions"`
}

// DescribeDomainAutoTunesRequest represents the DescribeDomainAutoTunesRequest schema from the OpenAPI specification
type DescribeDomainAutoTunesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeReservedInstancesRequest represents the DescribeReservedInstancesRequest schema from the OpenAPI specification
type DescribeReservedInstancesRequest struct {
}

// UpdatePackageResponse represents the UpdatePackageResponse schema from the OpenAPI specification
type UpdatePackageResponse struct {
	Packagedetails interface{} `json:"PackageDetails,omitempty"`
}

// AuthorizedPrincipal represents the AuthorizedPrincipal schema from the OpenAPI specification
type AuthorizedPrincipal struct {
	Principal interface{} `json:"Principal,omitempty"`
	Principaltype interface{} `json:"PrincipalType,omitempty"`
}

// GetCompatibleVersionsRequest represents the GetCompatibleVersionsRequest schema from the OpenAPI specification
type GetCompatibleVersionsRequest struct {
}

// SAMLIdp represents the SAMLIdp schema from the OpenAPI specification
type SAMLIdp struct {
	Metadatacontent interface{} `json:"MetadataContent"`
	Entityid interface{} `json:"EntityId"`
}

// UpgradeDomainRequest represents the UpgradeDomainRequest schema from the OpenAPI specification
type UpgradeDomainRequest struct {
	Advancedoptions interface{} `json:"AdvancedOptions,omitempty"`
	Domainname interface{} `json:"DomainName"`
	Performcheckonly interface{} `json:"PerformCheckOnly,omitempty"`
	Targetversion interface{} `json:"TargetVersion"`
}

// EncryptionAtRestOptionsStatus represents the EncryptionAtRestOptionsStatus schema from the OpenAPI specification
type EncryptionAtRestOptionsStatus struct {
	Status interface{} `json:"Status"`
	Options interface{} `json:"Options"`
}

// ErrorDetails represents the ErrorDetails schema from the OpenAPI specification
type ErrorDetails struct {
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Errortype interface{} `json:"ErrorType,omitempty"`
}

// DryRunResults represents the DryRunResults schema from the OpenAPI specification
type DryRunResults struct {
	Message interface{} `json:"Message,omitempty"`
	Deploymenttype interface{} `json:"DeploymentType,omitempty"`
}

// AutoTuneOptionsStatus represents the AutoTuneOptionsStatus schema from the OpenAPI specification
type AutoTuneOptionsStatus struct {
	Options interface{} `json:"Options,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// DescribeDomainNodesResponse represents the DescribeDomainNodesResponse schema from the OpenAPI specification
type DescribeDomainNodesResponse struct {
	Domainnodesstatuslist interface{} `json:"DomainNodesStatusList,omitempty"`
}

// DescribeReservedInstanceOfferingsRequest represents the DescribeReservedInstanceOfferingsRequest schema from the OpenAPI specification
type DescribeReservedInstanceOfferingsRequest struct {
}

// DescribeDomainConfigResponse represents the DescribeDomainConfigResponse schema from the OpenAPI specification
type DescribeDomainConfigResponse struct {
	Domainconfig interface{} `json:"DomainConfig"`
}

// DomainInformationContainer represents the DomainInformationContainer schema from the OpenAPI specification
type DomainInformationContainer struct {
	Awsdomaininformation interface{} `json:"AWSDomainInformation,omitempty"`
}

// SAMLOptionsOutput represents the SAMLOptionsOutput schema from the OpenAPI specification
type SAMLOptionsOutput struct {
	Subjectkey interface{} `json:"SubjectKey,omitempty"`
	Enabled interface{} `json:"Enabled,omitempty"`
	Idp interface{} `json:"Idp,omitempty"`
	Roleskey interface{} `json:"RolesKey,omitempty"`
	Sessiontimeoutminutes interface{} `json:"SessionTimeoutMinutes,omitempty"`
}

// AutoTuneMaintenanceSchedule represents the AutoTuneMaintenanceSchedule schema from the OpenAPI specification
type AutoTuneMaintenanceSchedule struct {
	Cronexpressionforrecurrence interface{} `json:"CronExpressionForRecurrence,omitempty"`
	Duration interface{} `json:"Duration,omitempty"`
	Startat interface{} `json:"StartAt,omitempty"`
}

// DescribePackagesResponse represents the DescribePackagesResponse schema from the OpenAPI specification
type DescribePackagesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Packagedetailslist interface{} `json:"PackageDetailsList,omitempty"`
}

// AdvancedSecurityOptionsInput represents the AdvancedSecurityOptionsInput schema from the OpenAPI specification
type AdvancedSecurityOptionsInput struct {
	Enabled interface{} `json:"Enabled,omitempty"`
	Internaluserdatabaseenabled interface{} `json:"InternalUserDatabaseEnabled,omitempty"`
	Masteruseroptions interface{} `json:"MasterUserOptions,omitempty"`
	Samloptions interface{} `json:"SAMLOptions,omitempty"`
	Anonymousauthenabled interface{} `json:"AnonymousAuthEnabled,omitempty"`
}

// AccessPoliciesStatus represents the AccessPoliciesStatus schema from the OpenAPI specification
type AccessPoliciesStatus struct {
	Options interface{} `json:"Options"`
	Status interface{} `json:"Status"`
}

// ScheduledAutoTuneDetails represents the ScheduledAutoTuneDetails schema from the OpenAPI specification
type ScheduledAutoTuneDetails struct {
	Actiontype interface{} `json:"ActionType,omitempty"`
	Date interface{} `json:"Date,omitempty"`
	Severity interface{} `json:"Severity,omitempty"`
	Action interface{} `json:"Action,omitempty"`
}

// StorageType represents the StorageType schema from the OpenAPI specification
type StorageType struct {
	Storagesubtypename interface{} `json:"StorageSubTypeName,omitempty"`
	Storagetypelimits interface{} `json:"StorageTypeLimits,omitempty"`
	Storagetypename interface{} `json:"StorageTypeName,omitempty"`
}

// ListDomainNamesRequest represents the ListDomainNamesRequest schema from the OpenAPI specification
type ListDomainNamesRequest struct {
}

// CrossClusterSearchConnectionProperties represents the CrossClusterSearchConnectionProperties schema from the OpenAPI specification
type CrossClusterSearchConnectionProperties struct {
	Skipunavailable interface{} `json:"SkipUnavailable,omitempty"`
}

// InboundConnectionStatus represents the InboundConnectionStatus schema from the OpenAPI specification
type InboundConnectionStatus struct {
	Message interface{} `json:"Message,omitempty"`
	Statuscode interface{} `json:"StatusCode,omitempty"`
}

// AutoTuneStatus represents the AutoTuneStatus schema from the OpenAPI specification
type AutoTuneStatus struct {
	Creationdate interface{} `json:"CreationDate"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Pendingdeletion interface{} `json:"PendingDeletion,omitempty"`
	State interface{} `json:"State"`
	Updatedate interface{} `json:"UpdateDate"`
	Updateversion interface{} `json:"UpdateVersion,omitempty"`
}

// ColdStorageOptions represents the ColdStorageOptions schema from the OpenAPI specification
type ColdStorageOptions struct {
	Enabled interface{} `json:"Enabled"`
}

// DescribeDomainConfigRequest represents the DescribeDomainConfigRequest schema from the OpenAPI specification
type DescribeDomainConfigRequest struct {
}

// VpcEndpointError represents the VpcEndpointError schema from the OpenAPI specification
type VpcEndpointError struct {
	Vpcendpointid interface{} `json:"VpcEndpointId,omitempty"`
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
}

// AWSDomainInformation represents the AWSDomainInformation schema from the OpenAPI specification
type AWSDomainInformation struct {
	Region interface{} `json:"Region,omitempty"`
	Domainname interface{} `json:"DomainName"`
	Ownerid interface{} `json:"OwnerId,omitempty"`
}

// VPCDerivedInfoStatus represents the VPCDerivedInfoStatus schema from the OpenAPI specification
type VPCDerivedInfoStatus struct {
	Options interface{} `json:"Options"`
	Status interface{} `json:"Status"`
}

// UpgradeHistory represents the UpgradeHistory schema from the OpenAPI specification
type UpgradeHistory struct {
	Upgradestatus interface{} `json:"UpgradeStatus,omitempty"`
	Starttimestamp interface{} `json:"StartTimestamp,omitempty"`
	Stepslist interface{} `json:"StepsList,omitempty"`
	Upgradename interface{} `json:"UpgradeName,omitempty"`
}

// LogPublishingOptions represents the LogPublishingOptions schema from the OpenAPI specification
type LogPublishingOptions struct {
}

// DescribeDomainHealthRequest represents the DescribeDomainHealthRequest schema from the OpenAPI specification
type DescribeDomainHealthRequest struct {
}

// DryRunProgressStatus represents the DryRunProgressStatus schema from the OpenAPI specification
type DryRunProgressStatus struct {
	Validationfailures interface{} `json:"ValidationFailures,omitempty"`
	Creationdate interface{} `json:"CreationDate"`
	Dryrunid interface{} `json:"DryRunId"`
	Dryrunstatus interface{} `json:"DryRunStatus"`
	Updatedate interface{} `json:"UpdateDate"`
}

// UpgradeDomainResponse represents the UpgradeDomainResponse schema from the OpenAPI specification
type UpgradeDomainResponse struct {
	Changeprogressdetails interface{} `json:"ChangeProgressDetails,omitempty"`
	Domainname interface{} `json:"DomainName,omitempty"`
	Performcheckonly interface{} `json:"PerformCheckOnly,omitempty"`
	Targetversion interface{} `json:"TargetVersion,omitempty"`
	Upgradeid interface{} `json:"UpgradeId,omitempty"`
	Advancedoptions interface{} `json:"AdvancedOptions,omitempty"`
}

// UpdateScheduledActionResponse represents the UpdateScheduledActionResponse schema from the OpenAPI specification
type UpdateScheduledActionResponse struct {
	Scheduledaction interface{} `json:"ScheduledAction,omitempty"`
}

// GetUpgradeHistoryRequest represents the GetUpgradeHistoryRequest schema from the OpenAPI specification
type GetUpgradeHistoryRequest struct {
}

// ConnectionProperties represents the ConnectionProperties schema from the OpenAPI specification
type ConnectionProperties struct {
	Crossclustersearch interface{} `json:"CrossClusterSearch,omitempty"`
	Endpoint interface{} `json:"Endpoint,omitempty"`
}

// RejectInboundConnectionRequest represents the RejectInboundConnectionRequest schema from the OpenAPI specification
type RejectInboundConnectionRequest struct {
}

// InstanceTypeDetails represents the InstanceTypeDetails schema from the OpenAPI specification
type InstanceTypeDetails struct {
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Warmenabled interface{} `json:"WarmEnabled,omitempty"`
	Advancedsecurityenabled interface{} `json:"AdvancedSecurityEnabled,omitempty"`
	Applogsenabled interface{} `json:"AppLogsEnabled,omitempty"`
	Availabilityzones interface{} `json:"AvailabilityZones,omitempty"`
	Cognitoenabled interface{} `json:"CognitoEnabled,omitempty"`
	Encryptionenabled interface{} `json:"EncryptionEnabled,omitempty"`
	Instancerole interface{} `json:"InstanceRole,omitempty"`
}

// DomainConfig represents the DomainConfig schema from the OpenAPI specification
type DomainConfig struct {
	Nodetonodeencryptionoptions interface{} `json:"NodeToNodeEncryptionOptions,omitempty"`
	Offpeakwindowoptions interface{} `json:"OffPeakWindowOptions,omitempty"`
	Softwareupdateoptions interface{} `json:"SoftwareUpdateOptions,omitempty"`
	Changeprogressdetails interface{} `json:"ChangeProgressDetails,omitempty"`
	Engineversion interface{} `json:"EngineVersion,omitempty"`
	Advancedoptions interface{} `json:"AdvancedOptions,omitempty"`
	Encryptionatrestoptions interface{} `json:"EncryptionAtRestOptions,omitempty"`
	Vpcoptions interface{} `json:"VPCOptions,omitempty"`
	Snapshotoptions interface{} `json:"SnapshotOptions,omitempty"`
	Autotuneoptions interface{} `json:"AutoTuneOptions,omitempty"`
	Cognitooptions interface{} `json:"CognitoOptions,omitempty"`
	Domainendpointoptions interface{} `json:"DomainEndpointOptions,omitempty"`
	Logpublishingoptions interface{} `json:"LogPublishingOptions,omitempty"`
	Advancedsecurityoptions interface{} `json:"AdvancedSecurityOptions,omitempty"`
	Ebsoptions interface{} `json:"EBSOptions,omitempty"`
	Accesspolicies interface{} `json:"AccessPolicies,omitempty"`
	Clusterconfig interface{} `json:"ClusterConfig,omitempty"`
}

// AvailabilityZoneInfo represents the AvailabilityZoneInfo schema from the OpenAPI specification
type AvailabilityZoneInfo struct {
	Availabledatanodecount interface{} `json:"AvailableDataNodeCount,omitempty"`
	Configureddatanodecount interface{} `json:"ConfiguredDataNodeCount,omitempty"`
	Totalshards interface{} `json:"TotalShards,omitempty"`
	Totalunassignedshards interface{} `json:"TotalUnAssignedShards,omitempty"`
	Zonestatus interface{} `json:"ZoneStatus,omitempty"`
	Availabilityzonename interface{} `json:"AvailabilityZoneName,omitempty"`
}

// ListScheduledActionsRequest represents the ListScheduledActionsRequest schema from the OpenAPI specification
type ListScheduledActionsRequest struct {
}

// AcceptInboundConnectionRequest represents the AcceptInboundConnectionRequest schema from the OpenAPI specification
type AcceptInboundConnectionRequest struct {
}

// MasterUserOptions represents the MasterUserOptions schema from the OpenAPI specification
type MasterUserOptions struct {
	Masteruserpassword interface{} `json:"MasterUserPassword,omitempty"`
	Masteruserarn interface{} `json:"MasterUserARN,omitempty"`
	Masterusername interface{} `json:"MasterUserName,omitempty"`
}

// Duration represents the Duration schema from the OpenAPI specification
type Duration struct {
	Unit interface{} `json:"Unit,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// Limits represents the Limits schema from the OpenAPI specification
type Limits struct {
	Additionallimits interface{} `json:"AdditionalLimits,omitempty"`
	Instancelimits interface{} `json:"InstanceLimits,omitempty"`
	Storagetypes interface{} `json:"StorageTypes,omitempty"`
}

// UpdateDomainConfigResponse represents the UpdateDomainConfigResponse schema from the OpenAPI specification
type UpdateDomainConfigResponse struct {
	Dryrunresults interface{} `json:"DryRunResults,omitempty"`
	Domainconfig interface{} `json:"DomainConfig"`
	Dryrunprogressstatus interface{} `json:"DryRunProgressStatus,omitempty"`
}

// ListVersionsResponse represents the ListVersionsResponse schema from the OpenAPI specification
type ListVersionsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Versions interface{} `json:"Versions,omitempty"`
}

// GetUpgradeStatusResponse represents the GetUpgradeStatusResponse schema from the OpenAPI specification
type GetUpgradeStatusResponse struct {
	Stepstatus interface{} `json:"StepStatus,omitempty"`
	Upgradename interface{} `json:"UpgradeName,omitempty"`
	Upgradestep interface{} `json:"UpgradeStep,omitempty"`
}

// StartServiceSoftwareUpdateRequest represents the StartServiceSoftwareUpdateRequest schema from the OpenAPI specification
type StartServiceSoftwareUpdateRequest struct {
	Desiredstarttime interface{} `json:"DesiredStartTime,omitempty"`
	Domainname interface{} `json:"DomainName"`
	Scheduleat interface{} `json:"ScheduleAt,omitempty"`
}

// AutoTuneOptions represents the AutoTuneOptions schema from the OpenAPI specification
type AutoTuneOptions struct {
	Desiredstate interface{} `json:"DesiredState,omitempty"`
	Maintenanceschedules interface{} `json:"MaintenanceSchedules,omitempty"`
	Rollbackondisable interface{} `json:"RollbackOnDisable,omitempty"`
	Useoffpeakwindow interface{} `json:"UseOffPeakWindow,omitempty"`
}

// GetCompatibleVersionsResponse represents the GetCompatibleVersionsResponse schema from the OpenAPI specification
type GetCompatibleVersionsResponse struct {
	Compatibleversions interface{} `json:"CompatibleVersions,omitempty"`
}

// DomainStatus represents the DomainStatus schema from the OpenAPI specification
type DomainStatus struct {
	Logpublishingoptions interface{} `json:"LogPublishingOptions,omitempty"`
	Accesspolicies interface{} `json:"AccessPolicies,omitempty"`
	Cognitooptions interface{} `json:"CognitoOptions,omitempty"`
	Created interface{} `json:"Created,omitempty"`
	Servicesoftwareoptions interface{} `json:"ServiceSoftwareOptions,omitempty"`
	Snapshotoptions interface{} `json:"SnapshotOptions,omitempty"`
	Autotuneoptions interface{} `json:"AutoTuneOptions,omitempty"`
	Clusterconfig interface{} `json:"ClusterConfig"`
	Encryptionatrestoptions interface{} `json:"EncryptionAtRestOptions,omitempty"`
	Ebsoptions interface{} `json:"EBSOptions,omitempty"`
	Offpeakwindowoptions interface{} `json:"OffPeakWindowOptions,omitempty"`
	Nodetonodeencryptionoptions interface{} `json:"NodeToNodeEncryptionOptions,omitempty"`
	Arn interface{} `json:"ARN"`
	Domainname interface{} `json:"DomainName"`
	Domainendpointoptions interface{} `json:"DomainEndpointOptions,omitempty"`
	Endpoint interface{} `json:"Endpoint,omitempty"`
	Changeprogressdetails interface{} `json:"ChangeProgressDetails,omitempty"`
	Vpcoptions interface{} `json:"VPCOptions,omitempty"`
	Deleted interface{} `json:"Deleted,omitempty"`
	Engineversion interface{} `json:"EngineVersion,omitempty"`
	Advancedoptions interface{} `json:"AdvancedOptions,omitempty"`
	Advancedsecurityoptions interface{} `json:"AdvancedSecurityOptions,omitempty"`
	Endpoints interface{} `json:"Endpoints,omitempty"`
	Processing interface{} `json:"Processing,omitempty"`
	Softwareupdateoptions interface{} `json:"SoftwareUpdateOptions,omitempty"`
	Upgradeprocessing interface{} `json:"UpgradeProcessing,omitempty"`
	Domainid interface{} `json:"DomainId"`
}

// OffPeakWindow represents the OffPeakWindow schema from the OpenAPI specification
type OffPeakWindow struct {
	Windowstarttime interface{} `json:"WindowStartTime,omitempty"`
}

// DomainEndpointOptionsStatus represents the DomainEndpointOptionsStatus schema from the OpenAPI specification
type DomainEndpointOptionsStatus struct {
	Status interface{} `json:"Status"`
	Options interface{} `json:"Options"`
}

// SnapshotOptionsStatus represents the SnapshotOptionsStatus schema from the OpenAPI specification
type SnapshotOptionsStatus struct {
	Options interface{} `json:"Options"`
	Status interface{} `json:"Status"`
}

// DeletePackageRequest represents the DeletePackageRequest schema from the OpenAPI specification
type DeletePackageRequest struct {
}

// PurchaseReservedInstanceOfferingRequest represents the PurchaseReservedInstanceOfferingRequest schema from the OpenAPI specification
type PurchaseReservedInstanceOfferingRequest struct {
	Reservedinstanceofferingid interface{} `json:"ReservedInstanceOfferingId"`
	Instancecount interface{} `json:"InstanceCount,omitempty"`
	Reservationname interface{} `json:"ReservationName"`
}

// ListInstanceTypeDetailsResponse represents the ListInstanceTypeDetailsResponse schema from the OpenAPI specification
type ListInstanceTypeDetailsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Instancetypedetails interface{} `json:"InstanceTypeDetails,omitempty"`
}

// DomainNodesStatus represents the DomainNodesStatus schema from the OpenAPI specification
type DomainNodesStatus struct {
	Storagetype interface{} `json:"StorageType,omitempty"`
	Storagevolumetype interface{} `json:"StorageVolumeType,omitempty"`
	Availabilityzone interface{} `json:"AvailabilityZone,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Nodeid interface{} `json:"NodeId,omitempty"`
	Nodestatus interface{} `json:"NodeStatus,omitempty"`
	Nodetype interface{} `json:"NodeType,omitempty"`
	Storagesize interface{} `json:"StorageSize,omitempty"`
}

// RejectInboundConnectionResponse represents the RejectInboundConnectionResponse schema from the OpenAPI specification
type RejectInboundConnectionResponse struct {
	Connection interface{} `json:"Connection,omitempty"`
}

// DescribeDomainHealthResponse represents the DescribeDomainHealthResponse schema from the OpenAPI specification
type DescribeDomainHealthResponse struct {
	Dedicatedmaster interface{} `json:"DedicatedMaster,omitempty"`
	Totalshards interface{} `json:"TotalShards,omitempty"`
	Totalunassignedshards interface{} `json:"TotalUnAssignedShards,omitempty"`
	Datanodecount interface{} `json:"DataNodeCount,omitempty"`
	Standbyavailabilityzonecount interface{} `json:"StandByAvailabilityZoneCount,omitempty"`
	Activeavailabilityzonecount interface{} `json:"ActiveAvailabilityZoneCount,omitempty"`
	Environmentinformation interface{} `json:"EnvironmentInformation,omitempty"`
	Masternode interface{} `json:"MasterNode,omitempty"`
	Warmnodecount interface{} `json:"WarmNodeCount,omitempty"`
	Clusterhealth interface{} `json:"ClusterHealth,omitempty"`
	Domainstate interface{} `json:"DomainState,omitempty"`
	Mastereligiblenodecount interface{} `json:"MasterEligibleNodeCount,omitempty"`
	Availabilityzonecount interface{} `json:"AvailabilityZoneCount,omitempty"`
}

// Filter represents the Filter schema from the OpenAPI specification
type Filter struct {
	Name interface{} `json:"Name,omitempty"`
	Values interface{} `json:"Values,omitempty"`
}

// GetPackageVersionHistoryRequest represents the GetPackageVersionHistoryRequest schema from the OpenAPI specification
type GetPackageVersionHistoryRequest struct {
}

// ListVpcEndpointAccessRequest represents the ListVpcEndpointAccessRequest schema from the OpenAPI specification
type ListVpcEndpointAccessRequest struct {
}

// UpdateScheduledActionRequest represents the UpdateScheduledActionRequest schema from the OpenAPI specification
type UpdateScheduledActionRequest struct {
	Actiontype interface{} `json:"ActionType"`
	Desiredstarttime interface{} `json:"DesiredStartTime,omitempty"`
	Scheduleat interface{} `json:"ScheduleAt"`
	Actionid interface{} `json:"ActionID"`
}

// ListVpcEndpointsForDomainRequest represents the ListVpcEndpointsForDomainRequest schema from the OpenAPI specification
type ListVpcEndpointsForDomainRequest struct {
}

// DescribeReservedInstanceOfferingsResponse represents the DescribeReservedInstanceOfferingsResponse schema from the OpenAPI specification
type DescribeReservedInstanceOfferingsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Reservedinstanceofferings interface{} `json:"ReservedInstanceOfferings,omitempty"`
}

// GetUpgradeHistoryResponse represents the GetUpgradeHistoryResponse schema from the OpenAPI specification
type GetUpgradeHistoryResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Upgradehistories interface{} `json:"UpgradeHistories,omitempty"`
}

// DescribeDomainsResponse represents the DescribeDomainsResponse schema from the OpenAPI specification
type DescribeDomainsResponse struct {
	Domainstatuslist interface{} `json:"DomainStatusList"`
}

// DissociatePackageResponse represents the DissociatePackageResponse schema from the OpenAPI specification
type DissociatePackageResponse struct {
	Domainpackagedetails interface{} `json:"DomainPackageDetails,omitempty"`
}

// DescribeInboundConnectionsResponse represents the DescribeInboundConnectionsResponse schema from the OpenAPI specification
type DescribeInboundConnectionsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Connections interface{} `json:"Connections,omitempty"`
}

// CognitoOptionsStatus represents the CognitoOptionsStatus schema from the OpenAPI specification
type CognitoOptionsStatus struct {
	Options interface{} `json:"Options"`
	Status interface{} `json:"Status"`
}

// SoftwareUpdateOptionsStatus represents the SoftwareUpdateOptionsStatus schema from the OpenAPI specification
type SoftwareUpdateOptionsStatus struct {
	Options interface{} `json:"Options,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// DescribeDomainRequest represents the DescribeDomainRequest schema from the OpenAPI specification
type DescribeDomainRequest struct {
}

// UpdateVpcEndpointRequest represents the UpdateVpcEndpointRequest schema from the OpenAPI specification
type UpdateVpcEndpointRequest struct {
	Vpcendpointid interface{} `json:"VpcEndpointId"`
	Vpcoptions interface{} `json:"VpcOptions"`
}

// ChangeProgressStage represents the ChangeProgressStage schema from the OpenAPI specification
type ChangeProgressStage struct {
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Lastupdated interface{} `json:"LastUpdated,omitempty"`
}

// ScheduledAction represents the ScheduledAction schema from the OpenAPI specification
type ScheduledAction struct {
	TypeField interface{} `json:"Type"`
	Cancellable interface{} `json:"Cancellable,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Id interface{} `json:"Id"`
	Scheduledby interface{} `json:"ScheduledBy,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Mandatory interface{} `json:"Mandatory,omitempty"`
	Scheduledtime interface{} `json:"ScheduledTime"`
	Severity interface{} `json:"Severity"`
}

// DescribeReservedInstancesResponse represents the DescribeReservedInstancesResponse schema from the OpenAPI specification
type DescribeReservedInstancesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Reservedinstances interface{} `json:"ReservedInstances,omitempty"`
}

// DescribeInstanceTypeLimitsResponse represents the DescribeInstanceTypeLimitsResponse schema from the OpenAPI specification
type DescribeInstanceTypeLimitsResponse struct {
	Limitsbyrole interface{} `json:"LimitsByRole,omitempty"`
}

// CompatibleVersionsMap represents the CompatibleVersionsMap schema from the OpenAPI specification
type CompatibleVersionsMap struct {
	Sourceversion interface{} `json:"SourceVersion,omitempty"`
	Targetversions interface{} `json:"TargetVersions,omitempty"`
}

// StartServiceSoftwareUpdateResponse represents the StartServiceSoftwareUpdateResponse schema from the OpenAPI specification
type StartServiceSoftwareUpdateResponse struct {
	Servicesoftwareoptions interface{} `json:"ServiceSoftwareOptions,omitempty"`
}

// DescribeOutboundConnectionsResponse represents the DescribeOutboundConnectionsResponse schema from the OpenAPI specification
type DescribeOutboundConnectionsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Connections interface{} `json:"Connections,omitempty"`
}

// ListVpcEndpointsResponse represents the ListVpcEndpointsResponse schema from the OpenAPI specification
type ListVpcEndpointsResponse struct {
	Vpcendpointsummarylist interface{} `json:"VpcEndpointSummaryList"`
	Nexttoken interface{} `json:"NextToken"`
}

// AdvancedOptions represents the AdvancedOptions schema from the OpenAPI specification
type AdvancedOptions struct {
}

// ChangeProgressStatusDetails represents the ChangeProgressStatusDetails schema from the OpenAPI specification
type ChangeProgressStatusDetails struct {
	Changeprogressstages interface{} `json:"ChangeProgressStages,omitempty"`
	Completedproperties interface{} `json:"CompletedProperties,omitempty"`
	Pendingproperties interface{} `json:"PendingProperties,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Totalnumberofstages interface{} `json:"TotalNumberOfStages,omitempty"`
	Changeid interface{} `json:"ChangeId,omitempty"`
}

// ListDomainsForPackageResponse represents the ListDomainsForPackageResponse schema from the OpenAPI specification
type ListDomainsForPackageResponse struct {
	Domainpackagedetailslist interface{} `json:"DomainPackageDetailsList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// AdvancedOptionsStatus represents the AdvancedOptionsStatus schema from the OpenAPI specification
type AdvancedOptionsStatus struct {
	Status interface{} `json:"Status"`
	Options interface{} `json:"Options"`
}

// ReservedInstanceOffering represents the ReservedInstanceOffering schema from the OpenAPI specification
type ReservedInstanceOffering struct {
	Currencycode interface{} `json:"CurrencyCode,omitempty"`
	Duration interface{} `json:"Duration,omitempty"`
	Fixedprice interface{} `json:"FixedPrice,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Paymentoption interface{} `json:"PaymentOption,omitempty"`
	Recurringcharges interface{} `json:"RecurringCharges,omitempty"`
	Reservedinstanceofferingid interface{} `json:"ReservedInstanceOfferingId,omitempty"`
	Usageprice interface{} `json:"UsagePrice,omitempty"`
}

// DeleteInboundConnectionRequest represents the DeleteInboundConnectionRequest schema from the OpenAPI specification
type DeleteInboundConnectionRequest struct {
}

// AuthorizeVpcEndpointAccessResponse represents the AuthorizeVpcEndpointAccessResponse schema from the OpenAPI specification
type AuthorizeVpcEndpointAccessResponse struct {
	Authorizedprincipal interface{} `json:"AuthorizedPrincipal"`
}

// AuthorizeVpcEndpointAccessRequest represents the AuthorizeVpcEndpointAccessRequest schema from the OpenAPI specification
type AuthorizeVpcEndpointAccessRequest struct {
	Account interface{} `json:"Account"`
}

// OutboundConnection represents the OutboundConnection schema from the OpenAPI specification
type OutboundConnection struct {
	Localdomaininfo interface{} `json:"LocalDomainInfo,omitempty"`
	Remotedomaininfo interface{} `json:"RemoteDomainInfo,omitempty"`
	Connectionalias interface{} `json:"ConnectionAlias,omitempty"`
	Connectionid interface{} `json:"ConnectionId,omitempty"`
	Connectionmode interface{} `json:"ConnectionMode,omitempty"`
	Connectionproperties interface{} `json:"ConnectionProperties,omitempty"`
	Connectionstatus interface{} `json:"ConnectionStatus,omitempty"`
}

// CognitoOptions represents the CognitoOptions schema from the OpenAPI specification
type CognitoOptions struct {
	Identitypoolid interface{} `json:"IdentityPoolId,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Userpoolid interface{} `json:"UserPoolId,omitempty"`
	Enabled interface{} `json:"Enabled,omitempty"`
}

// ZoneAwarenessConfig represents the ZoneAwarenessConfig schema from the OpenAPI specification
type ZoneAwarenessConfig struct {
	Availabilityzonecount interface{} `json:"AvailabilityZoneCount,omitempty"`
}

// InboundConnection represents the InboundConnection schema from the OpenAPI specification
type InboundConnection struct {
	Localdomaininfo interface{} `json:"LocalDomainInfo,omitempty"`
	Remotedomaininfo interface{} `json:"RemoteDomainInfo,omitempty"`
	Connectionid interface{} `json:"ConnectionId,omitempty"`
	Connectionmode interface{} `json:"ConnectionMode,omitempty"`
	Connectionstatus interface{} `json:"ConnectionStatus,omitempty"`
}

// LogPublishingOptionsStatus represents the LogPublishingOptionsStatus schema from the OpenAPI specification
type LogPublishingOptionsStatus struct {
	Options interface{} `json:"Options,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// VpcEndpoint represents the VpcEndpoint schema from the OpenAPI specification
type VpcEndpoint struct {
	Endpoint interface{} `json:"Endpoint,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Vpcendpointid interface{} `json:"VpcEndpointId,omitempty"`
	Vpcendpointowner interface{} `json:"VpcEndpointOwner,omitempty"`
	Vpcoptions interface{} `json:"VpcOptions,omitempty"`
	Domainarn interface{} `json:"DomainArn,omitempty"`
}

// RecurringCharge represents the RecurringCharge schema from the OpenAPI specification
type RecurringCharge struct {
	Recurringchargeamount interface{} `json:"RecurringChargeAmount,omitempty"`
	Recurringchargefrequency interface{} `json:"RecurringChargeFrequency,omitempty"`
}

// LogPublishingOption represents the LogPublishingOption schema from the OpenAPI specification
type LogPublishingOption struct {
	Cloudwatchlogsloggrouparn interface{} `json:"CloudWatchLogsLogGroupArn,omitempty"`
	Enabled interface{} `json:"Enabled,omitempty"`
}

// UpdateVpcEndpointResponse represents the UpdateVpcEndpointResponse schema from the OpenAPI specification
type UpdateVpcEndpointResponse struct {
	Vpcendpoint interface{} `json:"VpcEndpoint"`
}

// VpcEndpointSummary represents the VpcEndpointSummary schema from the OpenAPI specification
type VpcEndpointSummary struct {
	Vpcendpointid interface{} `json:"VpcEndpointId,omitempty"`
	Vpcendpointowner interface{} `json:"VpcEndpointOwner,omitempty"`
	Domainarn interface{} `json:"DomainArn,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// CreateVpcEndpointResponse represents the CreateVpcEndpointResponse schema from the OpenAPI specification
type CreateVpcEndpointResponse struct {
	Vpcendpoint interface{} `json:"VpcEndpoint"`
}

// SoftwareUpdateOptions represents the SoftwareUpdateOptions schema from the OpenAPI specification
type SoftwareUpdateOptions struct {
	Autosoftwareupdateenabled interface{} `json:"AutoSoftwareUpdateEnabled,omitempty"`
}

// DescribeVpcEndpointsRequest represents the DescribeVpcEndpointsRequest schema from the OpenAPI specification
type DescribeVpcEndpointsRequest struct {
	Vpcendpointids interface{} `json:"VpcEndpointIds"`
}
