package terraform

import (
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd/api"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DataGroup struct {
	GroupId               types.String              `tfsdk:"group_id"`
	AllowedIpRanges       types.Set                 `tfsdk:"allowed_ip_ranges"`
	BackupRetentionPeriod *string                   `tfsdk:"backup_retention_period"`
	ClusterArchitecture   *ClusterArchitecture      `tfsdk:"cluster_architecture"`
	ClusterName           types.String              `tfsdk:"cluster_name"`
	ClusterType           types.String              `tfsdk:"cluster_type"`
	Conditions            types.Set                 `tfsdk:"conditions"`
	Connection            types.String              `tfsdk:"connection_uri"`
	CreatedAt             types.String              `tfsdk:"created_at"`
	CspAuth               *bool                     `tfsdk:"csp_auth"`
	InstanceType          *api.InstanceType         `tfsdk:"instance_type"`
	LogsUrl               types.String              `tfsdk:"logs_url"`
	MetricsUrl            types.String              `tfsdk:"metrics_url"`
	PgConfig              *[]models.KeyValue        `tfsdk:"pg_config"`
	PgType                *api.PgType               `tfsdk:"pg_type"`
	PgVersion             *api.PgVersion            `tfsdk:"pg_version"`
	Phase                 types.String              `tfsdk:"phase"`
	PrivateNetworking     *bool                     `tfsdk:"private_networking"`
	Provider              *api.CloudProvider        `tfsdk:"cloud_provider"`
	Region                *api.Region               `tfsdk:"region"`
	ResizingPvc           types.Set                 `tfsdk:"resizing_pvc"`
	Storage               *Storage                  `tfsdk:"storage"`
	MaintenanceWindow     *models.MaintenanceWindow `tfsdk:"maintenance_window"`
	ServiceAccountIds     types.Set                 `tfsdk:"service_account_ids"`
	PeAllowedPrincipalIds types.Set                 `tfsdk:"pe_allowed_principal_ids"`
}
