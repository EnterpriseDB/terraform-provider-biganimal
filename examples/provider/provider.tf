# Configure the BigAnimal Provider
provider "biganimal" {
  ba_bearer_token = "<redacted>"
  // edb_tf_access_key: if set, this will be used instead of the ba_bearer_token above.
  // This can also be set as an environment variable. If it is set both here and
  // in an environment variable then the access key set in the environment variable
  // will take priority and be used
  edb_tf_access_key = "<redacted>"
  //ba_api_uri   = "https://portal.biganimal.com/api/v3" // Optional
}
# Manage the resources
