## v0.5.0 (July 24. 2023)
Features:
* New Data Source: `biganimal_pgd`
* New Resource to manage BigAnimal Extreme High Availability clusters: `biganimal_pgd`
* Projects, Regions and PGD resources can now be imported.

Enhancements:
* Dependency updates.
* `biganimal_project` and `biganimal_region` resources and `biganimal_projects` and `biganimal_region` data sources are migrated to use the new Terraform Plugin Framework Library.
* Various CI improvements

## v0.4.2 (June 6. 2023)
Bug Fixes:
* Regression fix for Faraway Replicas
* Fix for Projects data source

Enhancements:
* Acceptance tests are introduced.
* Dependencies updated.
* Development documents improved.

## v0.4.1 (April 24. 2023)
Bug Fixes:
* Throughput field in the `biganimal_cluster` and `biganimal_faraway_replica` resources is now configurable.
* Big fixes on drift-detection of allowedIpRanges and pgConfig fields of the `biganimal_cluster` and `biganimal_faraway_replica` resources.

Enhancements:
* Dependency updates
* Extended developer documentation


## v0.4.0 (April 05. 2023)
Features:
* New Data Source: `biganimal_aws_connection`
* New Data Source: `biganimal_faraway_replica`
* New Resource: `biganimal_aws_connection`
* New Resource: `biganimal_azure_connection`
* New Resource: `biganimal_faraway_replica`

Enhancements:
* data-source/biganimal_cluster: `faraway_replica_ids` and `cluster_type` fields are added.
* resource/biganimal_cluster: `faraway_replica_ids` and `cluster_type` fields are added.
* Initial skeleton for the acceptance tests implemented.
* Switched to terraform-plugin-testing module
* Dependencies updated.
* Various CI improvements.

## v0.3.0 (February 15. 2023)

Enhancements:
* `most_recent` field for biganimal_cluster data-source
* Improvements in drift detection
* New `biganimal_project` resource and `biganimal_projects` data-source
* Timeout increase in the region client.
* Various dependency updates

## v0.2.0 (January 16. 2023)

BigAnimal provider now uses the BigAnimal API v3.

## v0.1.2 (January 10, 2023)

Enhancements:
* New fields are added to biganimal_cluster resource and data-sources
 - csp_auth field for IAM authentication in AWS
 - logs_url and metrics_url fields
* Various dependency updates
* GitHub Issue Templates are updated.

## 0.1.1 (December 9, 2022)

Enhancements:
* Better Error Handling: Now we give more details about the errors
* Better documentation and minor fixes
* Include terraform version in the User Agent string

## 0.1.0 (November 29, 2022)

Initial version of the terraform provider that includes `biganimal_cluster` and `biganimal_region` data source
and resources


BACKWARDS INCOMPATIBILITIES / NOTES:
