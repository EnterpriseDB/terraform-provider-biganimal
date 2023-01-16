# Examples

This directory contains examples that are mostly used for documentation, but can also be run/tested manually via the Terraform CLI.

The document generation tool looks for files in the following locations by default. All other *.tf files besides the ones mentioned below are ignored by the documentation tool. This is useful for creating examples that can run and/or ar testable even if some parts are not relevant for the documentation.

* **provider/provider.tf** example file for the provider index page
* **data-sources/<full_data_source_name>/data-source.tf** example file for the named data source page
* **resources/<full_resource_name>/resource.tf** example file for the named data source page

## biganimal_cluster resource examples
* [Single node cluster example](./resources/biganimal_cluster/single_node/resource.tf)
  * [Single node cluster example on AWS](./resources/biganimal_cluster/single_node/aws/resource.tf)
  * [Single node cluster example on Azure](./resources/biganimal_cluster/single_node/azure/resource.tf)
* [High availability cluster example](./resources/biganimal_cluster/ha/resource.tf)
* [Extreme high availability cluster example](./resources/biganimal_cluster/eha/resource.tf)
