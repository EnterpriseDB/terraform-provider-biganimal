package api

type ClusterCreateDataGroup struct {
	AllowedIpRanges       []ClusterAllowedIpRange        `json:"allowedIpRanges"`
	BackupRetentionPeriod *string                        `json:"backupRetentionPeriod,omitempty"`
	ClusterArchitecture   *ClusterClusterArchitecture    `json:"clusterArchitecture"`
	ClusterType           string                         `json:"clusterType"`
	CspAuth               *bool                          `json:"cspAuth,omitempty"`
	InstanceType          *ClusterInstanceType           `json:"instanceType"`
	MaintenanceWindow     *ClusterMaintenanceWindow      `json:"maintenanceWindow,omitempty"`
	PgConfig              []ArrayOfNameValueObjectsInner `json:"pgConfig"`
	PgType                *ClusterPgType                 `json:"pgType"`
	PgVersion             *ClusterPgVersion              `json:"pgVersion"`
	PrivateNetworking     bool                           `json:"privateNetworking"`
	Provider              *ClusterCloudProvider          `json:"provider"`
	Region                *ClusterRegion                 `json:"region"`
	Storage               *ClusterStorage                `json:"storage"`
}
