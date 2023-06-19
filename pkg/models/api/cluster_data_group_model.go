package api

// Cluster data group
type ClusterDataGroup struct {
	AllowedIpRanges            []ClusterAllowedIpRange             `json:"allowedIpRanges"`
	BackupRetentionPeriod      string                              `json:"backupRetentionPeriod"`
	ClusterArchitecture        *ClusterClusterArchitectureResponse `json:"clusterArchitecture,omitempty"`
	ClusterName                string                              `json:"clusterName"`
	ClusterType                string                              `json:"clusterType"`
	Conditions                 []ClusterConditionsInner            `json:"conditions"`
	Connection                 *ClusterConnection                  `json:"connection,omitempty"`
	CreatedAt                  *PointInTime                        `json:"createdAt,omitempty"`
	CspAuth                    *bool                               `json:"cspAuth,omitempty"`
	DeletedAt                  *PointInTime                        `json:"deletedAt,omitempty"`
	EvaluatedPgConfig          *[]ArrayOfNameValueObjectsInner     `json:"evaluatedPgConfig,omitempty"`
	ExpiredAt                  *PointInTime                        `json:"expiredAt,omitempty"`
	FarawayReplicaIds          *[]string                           `json:"farawayReplicaIds,omitempty"`
	FirstRecoverabilityPointAt *PointInTime                        `json:"firstRecoverabilityPointAt,omitempty"`
	InstanceType               *CloudProviderRegionInstanceType    `json:"instanceType,omitempty"`
	LogsUrl                    *string                             `json:"logsUrl,omitempty"`
	MaintenanceInProgress      *bool                               `json:"maintenanceInProgress,omitempty"`
	MaintenanceWindow          *ClusterMaintenanceWindow           `json:"maintenanceWindow,omitempty"`
	MetricsUrl                 *string                             `json:"metricsUrl,omitempty"`
	PgConfig                   *[]ArrayOfNameValueObjectsInner     `json:"pgConfig,omitempty"`
	PgType                     *PgType                             `json:"pgType,omitempty"`
	PgVersion                  *PgVersion                          `json:"pgVersion,omitempty"`
	Phase                      string                              `json:"phase"`
	PrivateNetworking          bool                                `json:"privateNetworking"`
	Provider                   *CloudProvider                      `json:"provider,omitempty"`
	ReadOnlyConnections        *bool                               `json:"readOnlyConnections,omitempty"`
	Region                     *CloudProviderRegion                `json:"region,omitempty"`
	ReplicaSourceClusterId     *string                             `json:"replicaSourceClusterId,omitempty"`
	ResizingPvc                *[]string                           `json:"resizingPvc,omitempty"`
	Storage                    *ClusterStorageResponse             `json:"storage,omitempty"`
	SuperuserAccess            *bool                               `json:"superuserAccess,omitempty"`
	GroupId                    string                              `json:"groupId"`
}
