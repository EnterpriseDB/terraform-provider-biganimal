# Examples

This directory contains examples that are mostly used for documentation, but can also be run/tested manually via the Terraform CLI.

The document generation tool looks for files in the following locations by default. All other *.tf files besides the ones mentioned below are ignored by the documentation tool. This is useful for creating examples that can run and/or ar testable even if some parts are not relevant for the documentation.

* **provider/provider.tf** example file for the provider index page
* **data-sources/<full_data_source_name>/data-source.tf** example file for the named data source page
* **resources/<full_resource_name>/resource.tf** example file for the named data source page

## biganimal_cluster resource examples
* [Single node cluster example](./resources/biganimal_cluster/single_node/resource.tf)
  * [Single node cluster example on AWS(Your Cloud Account)](./resources/biganimal_cluster/single_node/aws/resource.tf)
  * [Single node cluster example on AWS(BigAnimal's Cloud Account)](./resources/biganimal_cluster/single_node/bah_aws/resource.tf)
  * [Single node cluster example on Azure(Your Cloud Account)](./resources/biganimal_cluster/single_node/azure/resource.tf)
  * [Single node cluster example on Azure(BigAnimal's Cloud Account)](./resources/biganimal_cluster/single_node/bah_azure/resource.tf)
  * [Single node cluster example on Google Cloud(Your Cloud Account)](./resources/biganimal_cluster/single_node/gcp/resource.tf)
  * [Single node cluster example on Google Cloud(BigAnimal's Cloud Account)](./resources/biganimal_cluster/single_node/bah_gcp/resource.tf)

* [Primary/Standby High availability cluster example](./resources/biganimal_cluster/ha/resource.tf)
* For Distributed High Availability cluster examples, please check [the biganimal_pgd resource examples below](#biganimal_pgd-resource-examples-for-managing-distributed-high-availability-clusters)

## biganimal_pgd resource examples (for managing Distributed High Availability clusters)

* [PGD Azure One Data Group Example](./resources/biganimal_pgd/azure/data_group/resource.tf)
* [PGD Azure Two Data Groups with One Witness Group Example](./resources/biganimal_pgd/azure/data_groups_with_witness_group/resource.tf)
* [PGD Azure BigAnimal's cloud account One Data Group Example](./resources/biganimal_pgd/azure/bah_data_group/resource.tf)
* [PGD Azure BigAnimal's cloud account Two Data Groups with One Witness Group Example](./resources/biganimal_pgd/azure/bah_data_groups_with_witness_group/resource.tf)
* [PGD AWS One Data Group Example](./resources/biganimal_pgd/aws/data_group/resource.tf)
* [PGD AWS Two Data Groups with One Witness Group Example](./resources/biganimal_pgd/aws/data_groups_with_witness_group/resource.tf)
* [PGD AWS BigAnimal's cloud account One Data Group Example](./resources/biganimal_pgd/aws/bah_data_group/resource.tf)
* [PGD AWS BigAnimal's cloud account Two Data Groups with One Witness Group Example](./resources/biganimal_pgd/aws/bah_data_groups_with_witness_group/resource.tf)
* [PGD Google Cloud One Data Group Example](./resources/biganimal_pgd/gcp/data_group/resource.tf)
* [PGD Google Cloud Two Data Groups with One Witness Group Example](./resources/biganimal_pgd/gcp/data_groups_with_witness_group/resource.tf)
* [PGD GCP BigAnimal's cloud account One Data Group Example](./resources/biganimal_pgd/gcp/bah_data_group/resource.tf)
* [PGD GCP BigAnimal's cloud account Two Data Groups with One Witness Group Example](./resources/biganimal_pgd/gcp/bah_data_groups_with_witness_group/resource.tf)
