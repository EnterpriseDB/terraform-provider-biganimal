## v3.0.0 (July 3. 2025)
Features:
* Display tag connected resources in `biganimal_tag` resource.

Enhancements:
* Removed tag color from `biganimal_cluster`, `biganimal_analytics`, `biganimal_faraway` and `biganimal_pgd` resources. Tag color can only be updated/changed in `biganimal_tag` resource.

Bug Fixes:
* Tag name change in `biganimal_tag` resource bug fix.
* Bug fixes for wal_storage.
* Bug fixes when private networking is true.

## v2.0.0 (March 18. 2025)
Features:
* private_link_service_alias and private_link_service_name field is now supported and displayed in `biganimal_cluster`, `biganimal_analytics` and `biganimal_faraway` resources and data sources.

Enhancements:
* Deprecated field cluster_architecture.name for `biganimal_cluster` resource and data source is now removed.

Bug Fixes:
* Tags color validation and assigning tags bug fixes for `biganimal_cluster`, `biganimal_analytics`, `biganimal_faraway`, `biganimal_pgd`, `biganimal_project`, `biganimal_region`, `biganimal_tag` resources and data sources. tag_id field is now removed.

## v1.2.1 (January 06. 2025)
Bug Fixes:
* Fixed cluster_architecture.name field for `biganimal_cluster` resource and Data Source. It is now a deprecated and hidden field

## v1.2.0 (November 29. 2024)
Features:
* Support for Write-Ahead Logs (WAL) Storage in `biganimal_cluster`, `biganimal_faraway_replica`, and `biganimal_pgd` resources
* Support for backup schedule time in `biganimal_cluster`, `biganimal_analytics_cluster`, `biganimal_faraway_replica`, and `biganimal_pgd` resources

Enhancements:
* Validation checks to not allow pe_allowed_principal_ids and service_account_ids if using your cloud account

Bug Fixes:
* Fixed planned allowed_ip_ranges.description when using private_networking = true

## v1.1.1 (October 29. 2024)
Bug Fixes:
* Fixed Data Source `biganimal_cluster` cloud_provider not working with your cloud account
* Fixed Data Source `biganimal_projects` conversion error
* Fixed Data Source `biganimal_region` conversion error

## v1.1.0 (October 21. 2024)
Features:
* New Resource and Data Source to manage tags: `biganimal_tag`
* New Resource and Data Source to manage csp tags: `biganimal_csp_tag`
* Support to assign tags in `biganimal_cluster`, `biganimal_analytics_cluster`, `biganimal_faraway_replica`, `biganimal_pgd`, `biganimal_projects` and `biganimal_region` resources
* Support for read-only connections in `biganimal_pgd` resources
* Support service_name in `biganimal_cluster` resources

Enhancements:
* Updated authentication information in docs
* Updated AWS examples to use series 6 instance types by default

Bug Fixes:
* Fixed cluster_architecture.name not computing when changing cluster_architecture.id in `biganimal_cluster` resources

## v1.0.0 (August 07. 2024)
Features:
* Transparent Data Encryption (TDE) is now supported in `biganimal_cluster` and `biganimal_faraway_replica` resources
* Volume Snapshots are now supported in `biganimal_cluster` and `biganimal_faraway_replica` resources
* (Breaking change) `biganimal_cluster` and `biganimal_faraway_replica` datasources now use cluster ID instead of cluster name

Enhancements:
* (Breaking change) data groups in `biganimal_pgd` resources now use lists instead of sets
* (Breaking change) blocks are migrated to terraform plugin framework attributes in `biganimal_cluster` resources
* (Breaking change) `biganimal_faraway_replica` resources are migrated to terraform plugin framework attributes
* Updated examples

## v0.11.2 (July 31. 2024)
Bug Fixes:
* fixed pg bouncer settings = null will always show changes on update

## v0.11.0 (June 20. 2024)
Features:
* New Resource to manage Analytical clusters: `biganimal_analytics_cluster`
* New Data Source: `biganimal_analytics_cluster`

## v0.10.0 (May 13. 2024)
Features:
* PostGIS support for `biganimal_cluster` resources
* PostGIS and Pgvector support for `biganimal_faraway_replica` resources

## v0.9.0 (March 27. 2024)
Features:
* Added support to pause and resume a cluster for `biganimal_pgd` and `biganimal_cluster` resources

Bug Fixes:
* Fixed maintenance window plan inconsistent with response

## v0.8.1 (February 29. 2024)
Features:
* Updated access key requirements documentation

## v0.8.0 (February 29. 2024)
Features:
* New access keys authorisation support using provider resource or environment variable

## v0.7.1 (February 07. 2024)
Bug Fixes:
* Fixed pg config to only use user custom config values in plan and not include default config values
* Small bug fixes on create/update operation for `biganimal_pgd` and `biganimal_cluster` resources

Enhancements:
* Updated default storage size from 4 Gi to 32 Gi in examples for `biganimal_pgd` resources

## v0.7.0 (January 12. 2024)
Features:
* Pgvector support for `biganimal_cluster` resources
* PgBouncer support for `biganimal_cluster` resources
* BigAnimal Terraform Provider now can output password for `biganimal_pgd` resources

Bug Fixes:
* Private networking and allowed IP ranges mismatching plan bug fix
* Other bug fixes

Enhancements:
* Dependency updates
* Improved warnings for `biganimal_pgd` resources
* Updated example links in the examples README

## v0.6.1 (October 18. 2023)
Features:
* BigAnimal Terraform Provider now supports BigAnimal's cloud account AWS, Azure and Google Cloud provider for `biganimal_cluster` and `biganimal_pgd` resources
* Custom maintenance window support for `biganimal_cluster` resources
* Custom maintenance window support for the Witness groups in the `biganimal_pgd` resources

Bug Fixes:
* Removed PGD upscale not supported note in `biganimal_pgd` resource docs
* Small bug fixes in `biganimal_pgd` resources

Enhancements:
* Dependency updates
* PG Config values warnings support for `biganimal_pgd` and `biganimal_cluster` resources
* Changed default data nodes from 2 to 3 for `biganimal_pgd` resource
* Add and change validations for `biganimal_cluster` resources

## v0.6.0 (August 28. 2023)
Features:
* Added examples for PGD resources on Google Cloud.
* Cross-CSP support for the Witness groups in the PGD resources
* Default PG Config values for `biganimal_pgd` and `biganimal_cluster` resources
* Cluster resources can now be imported.

Enhancements:
* Dependency updates
* Cluster architecture names are updated
* Terraform and Golang models for PGD implementation has been updated.

## v0.5.1 (July 27. 2023)
Features:
* BigAnimal Terraform Provider now supports GCP. Examples and documentation are updated.

Bug Fixes:
* `biganimal_pgd` resource: For immutable fields(cloud provider, pg version or pg type), when user tries to change those fields, provider throws an error.

Enhancements:
* Dependency updates
* Examples now use EPAS-15 as the pre-configured postgresql version

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
