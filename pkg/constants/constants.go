package constants

const (
	_TDE_INFO_LINK     = "https://www.enterprisedb.com/docs/biganimal/latest/getting_started/creating_a_cluster/#security"
	TDE_KEY_AWS_ACTION = "\nPlease update your KMS key policy to grant this cluster encrypt and decrypt permissions on your key." +
		"\nYour cluster will activate automatically once key permissions are granted." +
		"\nPrincipal: %v" +
		"\nAction: 'KMS:Encrypt', 'KMS:Decrypt'" +
		"\nPlease see below page for more information.\n" + _TDE_INFO_LINK
	TDE_KEY_GCP_ACTION = "Action Required: grant key permissions to activate the cluster." +
		"\nGrant the following permission to the service account: %v" +
		"\n - cloudkms.cryptoKeyVersions.useToEncrypt" +
		"\n - cloudkms.cryptoKeyVersions.useToDecrypt" +
		"\nYour cluster will activate automatically once key permissions are granted." +
		"\nPlease see below page for more information.\n" + _TDE_INFO_LINK
	TDE_KEY_AZURE_ACTION = "Action Required: grant key permissions to activate the cluster." +
		"\nGrant the following permission to the MSI workload identity: %v" +
		"\n - Microsoft.KeyVault/vaults/keys/encrypt/action" +
		"\n - Microsoft.KeyVault/vaults/keys/decrypt/action" +
		"\nThis can be done by modifying the access policy of the key vault." +
		"\nYour cluster will activate automatically once key permissions are granted." +
		"\nPlease see below page for more information.\n" + _TDE_INFO_LINK
	TDE_KEY_NO_ACTION                     = "No Action Required"
	TDE_KEY_ACTION_UNKNOWN_PROVIDER_ERROR = "Error: Transparent data encryption provider: %v is not supported."
	TDE_CHECK_ACTION                      = "Transparent data encryption is enabled. Please check transparent_data_encryption_action after applying the plan for actions required."
)

type PhaseState int

const (
	NOT_HEALTHY PhaseState = iota
	HEALTHY
	WAITING_FOR_ACCESS_TO_ENCRYPTION_KEY
)

const (
	CONDITION_DEPLOYED                         = "biganimal.com/deployed"
	PHASE_HEALTHY                              = "Cluster in healthy state"
	PHASE_PAUSED                               = "Cluster has been paused"
	PHASE_WAITING_FOR_ACCESS_TO_ENCRYPTION_KEY = "Waiting for TDE key to get reachable"
)
