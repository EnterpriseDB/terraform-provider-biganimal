package tf

type DataGroupData struct {
	GroupId                    string                              `tfsdk:"group_id"`
	AllowedIpRanges            []ClusterAllowedIpRange             `tfsdk:"allowed_ip_ranges"`
	BackupRetentionPeriod      string                              `tfsdk:"backup_retention_period"`
	ClusterArchitecture        *ClusterClusterArchitectureResponse `tfsdk:"cluster_architecture"`
	ClusterName                string                              `tfsdk:"cluster_name"`
	ClusterType                string                              `tfsdk:"cluster_type"`
	CreatedAt                  *PointInTime                        `tfsdk:"created_at"`
	DeletedAt                  *PointInTime                        `tfsdk:"deleted_at"`
	ExpiredAt                  *PointInTime                        `tfsdk:"expired_at"`
	FirstRecoverabilityPointAt *PointInTime                        `tfsdk:"first_recoverability_point_at"`
	InstanceType               *InstanceType                       `tfsdk:"instance_type"`
	LogsUrl                    *string                             `tfsdk:"logs_url"`
	MetricsUrl                 *string                             `tfsdk:"metrics_url"`
	Connection                 *ClusterConnection                  `tfsdk:"connection_uri"`
	PgConfig                   *[]ArrayOfNameValueObjectsInner     `tfsdk:"pg_config"`
	PgType                     *PgType                             `tfsdk:"pg_type"`
	PgVersion                  *PgVersion                          `tfsdk:"pg_version"`
	Phase                      string                              `tfsdk:"phase"`
	PrivateNetworking          bool                                `tfsdk:"private_networking"`
	Provider                   *CloudProvider                      `tfsdk:"cloud_provider"`
	CspAuth                    *bool                               `tfsdk:"csp_auth"`
	Region                     *string                             `tfsdk:"region"`
	ResizingPvc                *[]string                           `tfsdk:"resizing_pvc"`
	Storage                    *ClusterStorageResponse             `tfsdk:"storage"`
}

type WitnessGroupData struct {
	Region *CloudProviderRegion `tfsdk:"region"`
}

type PGDDataSourceData struct {
	ProjectID     string             `tfsdk:"project_id"`
	ClusterID     *string            `tfsdk:"cluster_id"`
	ClusterName   string             `tfsdk:"cluster_name"`
	MostRecent    *bool              `tfsdk:"most_recent"`
	DataGroups    []DataGroupData    `tfsdk:"data_groups"`
	WitnessGroups []WitnessGroupData `tfsdk:"witness_groups"`
}
