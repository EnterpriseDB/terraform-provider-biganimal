# biganimal_pgd (Resource)

The PGD cluster data source describes a BigAnimal cluster. The data source requires your PGD cluster name.



## PGD Azure One Data Group Example
```terraform
terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "1.2.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.6.0"
    }
  }
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

variable "cluster_name" {
  type        = string
  description = "The name of the cluster."
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

resource "biganimal_pgd" "pgd_cluster" {
  cluster_name = var.cluster_name
  project_id   = var.project_id
  password     = resource.random_password.password.result
  pause        = false
  #tags = [
  #  {
  #     tag_name  = "<ex_tag_name_1>"
  #     color = "blue"
  #  },
  #  {
  #     tag_name  = "<ex_tag_name_2>"
  #  },
  #]
  data_groups = [
    {
      allowed_ip_ranges = [
        {
          cidr_block  = "127.0.0.1/32"
          description = "localhost"
        },
        {
          cidr_block  = "192.168.0.1/32"
          description = "description!"
        },
      ]
      backup_retention_period = "6d"
      #  backup_schedule_time = "0 5 1 * * *" //24 hour format cron expression e.g. "0 5 1 * * *" is 01:05
      cluster_architecture = {
        cluster_architecture_id = "pgd"
        nodes                   = 3
      }
      csp_auth = false
      instance_type = {
        instance_type_id = "azure:Standard_D2s_v3"
      }
      pg_config = [
        {
          name  = "application_name"
          value = "created through terraform"
        },
        {
          name  = "array_nulls"
          value = "off"
        },
      ]
      storage = {
        volume_type       = "azurepremiumstorage"
        volume_properties = "P2"
        size              = "32 Gi"
      }
      storage = {
        volume_type       = "azurepremiumstorage"
        volume_properties = "P2"
        size              = "32 Gi"
      }
      #      wal_storage = {
      #        volume_type       = "azurepremiumstorage"
      #        volume_properties = "P2"
      #        size              = "32 Gi"
      #      }
      pg_type = {
        pg_type_id = "epas" #valid values ["epas", "pgextended", "postgres]"
      }
      pg_version = {
        pg_version_id = "15"
      }
      private_networking = false
      cloud_provider = {
        cloud_provider_id = "bah:azure" // "bah:azure" uses BigAnimal's cloud account Azure, use "azure" for your cloud account
      }
      region = {
        region_id = "northeurope"
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 1
        start_time = "13:00"
      }
      read_only_connections = false

      # pe_allowed_principal_ids = [
      #   <example_value> # ex: "9334e5e6-7f47-aE61-5A4F-ee067daeEf4A"
      # ]
    },
  ]
}

output "password" {
  sensitive = true
  value     = resource.biganimal_pgd.pgd_cluster.password
}
```

## PGD Azure Two Data Groups with One Witness Group Example
```terraform
terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "1.2.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.6.0"
    }
  }
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

variable "cluster_name" {
  type        = string
  description = "The name of the cluster."
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

resource "biganimal_pgd" "pgd_cluster" {
  cluster_name = var.cluster_name
  project_id   = var.project_id
  password     = resource.random_password.password.result
  pause        = false
  #tags = [
  #  {
  #     tag_name  = "<ex_tag_name_1>"
  #     color = "blue"
  #  },
  #  {
  #     tag_name  = "<ex_tag_name_2>"
  #  },
  #]
  data_groups = [
    {
      allowed_ip_ranges = [
        {
          cidr_block  = "127.0.0.1/32"
          description = "localhost"
        },
        {
          cidr_block  = "192.168.0.1/32"
          description = "description!"
        },
      ]
      backup_retention_period = "6d"
      #  backup_schedule_time = "0 5 1 * * *" //24 hour format cron expression e.g. "0 5 1 * * *" is 01:05
      cluster_architecture = {
        cluster_architecture_id = "pgd"
        nodes                   = 3
      }
      csp_auth = false
      instance_type = {
        instance_type_id = "azure:Standard_D2s_v3"
      }
      pg_config = [
        {
          name  = "application_name"
          value = "created through terraform"
        },
        {
          name  = "array_nulls"
          value = "off"
        },
      ]
      storage = {
        volume_type       = "azurepremiumstorage"
        volume_properties = "P2"
        size              = "32 Gi"
      }
      #      wal_storage = {
      #        volume_type       = "azurepremiumstorage"
      #        volume_properties = "P2"
      #        size              = "32 Gi"
      #      }
      pg_type = {
        pg_type_id = "epas" #valid values ["epas", "pgextended", "postgres]"
      }
      pg_version = {
        pg_version_id = "15"
      }
      private_networking = false
      cloud_provider = {
        cloud_provider_id = "bah:azure" // "bah:azure" uses BigAnimal's cloud account Azure, use "azure" for your cloud account
      }
      region = {
        region_id = "northeurope"
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 1
        start_time = "13:00"
      }
      # pe_allowed_principal_ids = [
      #   <example_value> # ex: "9334e5e6-7f47-aE61-5A4F-ee067daeEf4A"
      # ]
    },
    {
      allowed_ip_ranges = [
        {
          cidr_block  = "127.0.0.1/32"
          description = "localhost"
        },
        {
          cidr_block  = "192.168.0.1/32"
          description = "description!"
        },
      ]
      backup_retention_period = "6d"
      #  backup_schedule_time = "0 5 1 * * *" //24 hour format cron expression e.g. "0 5 1 * * *" is 01:05
      cluster_architecture = {
        cluster_architecture_id = "pgd"
        nodes                   = 3
      }
      csp_auth = false
      instance_type = {
        instance_type_id = "azure:Standard_D2s_v3"
      }
      pg_config = [
        {
          name  = "application_name"
          value = "created through terraform"
        },
        {
          name  = "array_nulls"
          value = "off"
        },
      ]
      storage = {
        volume_type       = "azurepremiumstorage"
        volume_properties = "P2"
        size              = "32 Gi"
      }
      #      wal_storage = {
      #        volume_type       = "azurepremiumstorage"
      #        volume_properties = "P2"
      #        size              = "32 Gi"
      #      }
      pg_type = {
        pg_type_id = "epas" #valid values ["epas", "pgextended", "postgres]"
      }
      pg_version = {
        pg_version_id = "15"
      }
      private_networking = false
      cloud_provider = {
        cloud_provider_id = "bah:azure" // "bah:azure" uses BigAnimal's cloud account Azure, use "azure" for your cloud account
      }
      region = {
        region_id = "eastus"
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 2
        start_time = "15:00"
      }
      # pe_allowed_principal_ids = [
      #   <example_value> # ex: "9334e5e6-7f47-aE61-5A4F-ee067daeEf4A"
      # ]
    }
  ]
  witness_groups = [
    {
      region = {
        region_id = "canadacentral"
      }
      cloud_provider = {
        cloud_provider_id = "bah:azure" // "bah:azure" uses BigAnimal's cloud account Azure, use "azure" for your cloud account
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 3
        start_time = "03:00"
      }
    }
  ]
}

output "password" {
  sensitive = true
  value     = resource.biganimal_pgd.pgd_cluster.password
}
```

## PGD AWS One Data Group Example
```terraform
terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "1.2.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.6.0"
    }
  }
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

variable "cluster_name" {
  type        = string
  description = "The name of the cluster."
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

resource "biganimal_pgd" "pgd_cluster" {
  cluster_name = var.cluster_name
  project_id   = var.project_id
  password     = resource.random_password.password.result
  pause        = false
  #tags = [
  #  {
  #     tag_name  = "<ex_tag_name_1>"
  #     color = "blue"
  #  },
  #  {
  #     tag_name  = "<ex_tag_name_2>"
  #  },
  #]
  data_groups = [
    {
      allowed_ip_ranges = [
        {
          cidr_block  = "127.0.0.1/32"
          description = "localhost"
        },
        {
          cidr_block  = "192.168.0.1/32"
          description = "description!"
        },
      ]
      backup_retention_period = "6d"
      #  backup_schedule_time = "0 5 1 * * *" //24 hour format cron expression e.g. "0 5 1 * * *" is 01:05
      cluster_architecture = {
        cluster_architecture_id = "pgd"
        nodes                   = 3
      }
      csp_auth = false
      instance_type = {
        instance_type_id = "aws:m6i.large"
      }
      pg_config = [
        {
          name  = "application_name"
          value = "created through terraform"
        },
        {
          name  = "array_nulls"
          value = "off"
        },
      ]
      storage = {
        volume_type       = "gp3"
        volume_properties = "gp3"
        size              = "32 Gi"
      }
      #      wal_storage = {
      #        volume_type       = "gp3"
      #        volume_properties = "gp3"
      #        size              = "32 Gi"
      #      }
      pg_type = {
        pg_type_id = "epas" #valid values ["epas", "pgextended", "postgres]"
      }
      pg_version = {
        pg_version_id = "15"
      }
      private_networking = false
      cloud_provider = {
        cloud_provider_id = "bah:aws" // "bah:aws" uses BigAnimal's cloud account AWS, use "aws" for your cloud account
      }
      region = {
        region_id = "eu-central-1"
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 6
        start_time = "13:00"
      }
      read_only_connections = false

      # pe_allowed_principal_ids = [
      #   <example_value> # ex: 123456789012
      # ]
    }
  ]
}

output "password" {
  sensitive = true
  value     = resource.biganimal_pgd.pgd_cluster.password
}
```

## PGD AWS Two Data Groups with One Witness Group Example
```terraform
terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "1.2.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.6.0"
    }
  }
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

variable "cluster_name" {
  type        = string
  description = "The name of the cluster."
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

resource "biganimal_pgd" "pgd_cluster" {
  cluster_name = var.cluster_name
  project_id   = var.project_id
  password     = resource.random_password.password.result
  pause        = false
  #tags = [
  #  {
  #     tag_name  = "<ex_tag_name_1>"
  #     color = "blue"
  #  },
  #  {
  #     tag_name  = "<ex_tag_name_2>"
  #  },
  #]
  data_groups = [
    {
      allowed_ip_ranges = [
        {
          cidr_block  = "127.0.0.1/32"
          description = "localhost"
        },
        {
          cidr_block  = "192.168.0.1/32"
          description = "description!"
        },
      ]
      backup_retention_period = "6d"
      #  backup_schedule_time = "0 5 1 * * *" //24 hour format cron expression e.g. "0 5 1 * * *" is 01:05
      cluster_architecture = {
        cluster_architecture_id = "pgd"
        nodes                   = 3
      }
      csp_auth = false
      instance_type = {
        instance_type_id = "aws:m6i.large"
      }
      pg_config = [
        {
          name  = "application_name"
          value = "created through terraform"
        },
        {
          name  = "array_nulls"
          value = "off"
        },
      ]
      storage = {
        volume_type       = "gp3"
        volume_properties = "gp3"
        size              = "32 Gi"
      }
      #      wal_storage = {
      #        volume_type       = "gp3"
      #        volume_properties = "gp3"
      #        size              = "32 Gi"
      #      }
      pg_type = {
        pg_type_id = "epas" #valid values ["epas", "pgextended", "postgres]"
      }
      pg_version = {
        pg_version_id = "15"
      }
      private_networking = false
      cloud_provider = {
        cloud_provider_id = "bah:aws" // "bah:aws" uses BigAnimal's cloud account AWS, use "aws" for your cloud account
      }
      region = {
        region_id = "eu-west-1"
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 1
        start_time = "13:00"
      }
      # pe_allowed_principal_ids = [
      #   <example_value> # ex: 123456789012
      # ]
    },
    {
      allowed_ip_ranges = [
        {
          cidr_block  = "127.0.0.1/32"
          description = "localhost"
        },
        {
          cidr_block  = "192.168.0.1/32"
          description = "description!"
        },
      ]
      backup_retention_period = "6d"
      #  backup_schedule_time = "0 5 1 * * *" //24 hour format cron expression e.g. "0 5 1 * * *" is 01:05
      cluster_architecture = {
        cluster_architecture_id = "pgd"
        nodes                   = 3
      }
      csp_auth = false
      instance_type = {
        instance_type_id = "aws:m6i.large"
      }
      pg_config = [
        {
          name  = "application_name"
          value = "created through terraform"
        },
        {
          name  = "array_nulls"
          value = "off"
        },
      ]
      storage = {
        volume_type       = "gp3"
        volume_properties = "gp3"
        size              = "32 Gi"
      }
      #      wal_storage = {
      #        volume_type       = "gp3"
      #        volume_properties = "gp3"
      #        size              = "32 Gi"
      #      }
      pg_type = {
        pg_type_id = "epas" #valid values ["epas", "pgextended", "postgres]"
      }
      pg_version = {
        pg_version_id = "15"
      }
      private_networking = false
      cloud_provider = {
        cloud_provider_id = "bah:aws" // "bah:aws" uses BigAnimal's cloud account AWS, use "aws" for your cloud account
      }
      region = {
        region_id = "eu-west-2"
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 2
        start_time = "15:00"
      }
      # pe_allowed_principal_ids = [
      #   <example_value> # ex: 123456789012
      # ]
    }
  ]
  witness_groups = [
    {
      region = {
        region_id = "us-east-1"
      }
      cloud_provider = {
        cloud_provider_id = "bah:aws" // "bah:aws" uses BigAnimal's cloud account AWS, use "aws" for your cloud account
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 3
        start_time = "03:00"
      }
    }
  ]
}

output "password" {
  sensitive = true
  value     = resource.biganimal_pgd.pgd_cluster.password
}
```

## PGD GCP One Data Group Example
```terraform
terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "1.2.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.6.0"
    }
  }
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

variable "cluster_name" {
  type        = string
  description = "The name of the cluster."
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

resource "biganimal_pgd" "pgd_cluster" {
  cluster_name = var.cluster_name
  project_id   = var.project_id
  password     = resource.random_password.password.result
  pause        = false
  #tags = [
  #  {
  #     tag_name  = "<ex_tag_name_1>"
  #     color = "blue"
  #  },
  #  {
  #     tag_name  = "<ex_tag_name_2>"
  #  },
  #]
  data_groups = [
    {
      allowed_ip_ranges = [
        {
          cidr_block  = "127.0.0.1/32"
          description = "localhost"
        },
        {
          cidr_block  = "192.168.0.1/32"
          description = "description!"
        },
      ]
      backup_retention_period = "6d"
      #  backup_schedule_time = "0 5 1 * * *" //24 hour format cron expression e.g. "0 5 1 * * *" is 01:05
      cluster_architecture = {
        cluster_architecture_id = "pgd"
        nodes                   = 3
      }
      csp_auth = false
      instance_type = {
        instance_type_id = "gcp:e2-highcpu-4"
      }
      pg_config = [
        {
          name  = "application_name"
          value = "created through terraform"
        },
        {
          name  = "array_nulls"
          value = "off"
        },
      ]
      storage = {
        volume_type       = "pd-ssd"
        volume_properties = "pd-ssd"
        size              = "32 Gi"
      }
      #      wal_storage = {
      #        volume_type       = "pd-ssd"
      #        volume_properties = "pd-ssd"
      #        size              = "32 Gi"
      #      }
      pg_type = {
        pg_type_id = "epas" #valid values ["epas", "pgextended", "postgres]"
      }
      pg_version = {
        pg_version_id = "15"
      }
      private_networking = false
      cloud_provider = {
        cloud_provider_id = "bah:gcp" // "bah:gpc" uses BigAnimal's cloud account Google Cloud provider, use "gcp" for your cloud account
      }
      region = {
        region_id = "us-east1"
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 6
        start_time = "13:00"
      }
      read_only_connections = false

      # pe_allowed_principal_ids = [
      #   <example_value> # ex: "development-data-123456"
      # ]

      # service_account_ids = [
      #   <only_needed_for_bah:gcp_clusters> # ex: "test@development-data-123456.iam.gserviceaccount.com"
      # ]
    }
  ]
}

output "password" {
  sensitive = true
  value     = resource.biganimal_pgd.pgd_cluster.password
}
```

## PGD GCP Two Data Groups with One Witness Group Example
```terraform
terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "1.2.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.6.0"
    }
  }
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

variable "cluster_name" {
  type        = string
  description = "The name of the cluster."
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

resource "biganimal_pgd" "pgd_cluster" {
  cluster_name = var.cluster_name
  project_id   = var.project_id
  password     = resource.random_password.password.result
  pause        = false
  #tags = [
  #  {
  #     tag_name  = "<ex_tag_name_1>"
  #     color = "blue"
  #  },
  #  {
  #     tag_name  = "<ex_tag_name_2>"
  #  },
  #]
  data_groups = [
    {
      allowed_ip_ranges = [
        {
          cidr_block  = "127.0.0.1/32"
          description = "localhost"
        },
        {
          cidr_block  = "192.168.0.1/32"
          description = "description!"
        },
      ]
      backup_retention_period = "6d"
      #  backup_schedule_time = "0 5 1 * * *" //24 hour format cron expression e.g. "0 5 1 * * *" is 01:05
      cluster_architecture = {
        cluster_architecture_id = "pgd"
        nodes                   = 3
      }
      csp_auth = false
      instance_type = {
        instance_type_id = "gcp:e2-highcpu-4"
      }
      pg_config = [
        {
          name  = "application_name"
          value = "created through terraform"
        },
        {
          name  = "array_nulls"
          value = "off"
        },
      ]
      storage = {
        volume_type       = "pd-ssd"
        volume_properties = "pd-ssd"
        size              = "32 Gi"
      }
      #      wal_storage = {
      #        volume_type       = "pd-ssd"
      #        volume_properties = "pd-ssd"
      #        size              = "32 Gi"
      #      }
      pg_type = {
        pg_type_id = "epas" #valid values ["epas", "pgextended", "postgres]"
      }
      pg_version = {
        pg_version_id = "15"
      }
      private_networking = false
      cloud_provider = {
        cloud_provider_id = "bah:gcp" // "bah:gpc" uses BigAnimal's cloud account Google Cloud provider, use "gcp" for your cloud account
      }
      region = {
        region_id = "us-east1"
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 6
        start_time = "13:00"
      }
      # pe_allowed_principal_ids = [
      #   <example_value> # ex: "development-data-123456"
      # ]

      # service_account_ids = [
      #   <only_needed_for_bah:gcp_clusters> # ex: "test@development-data-123456.iam.gserviceaccount.com"
      # ]
    },
    {
      allowed_ip_ranges = [
        {
          cidr_block  = "127.0.0.1/32"
          description = "localhost"
        },
        {
          cidr_block  = "192.168.0.1/32"
          description = "description!"
        },
      ]
      backup_retention_period = "6d"
      #  backup_schedule_time = "0 5 1 * * *" //24 hour format cron expression e.g. "0 5 1 * * *" is 01:05
      cluster_architecture = {
        cluster_architecture_id = "pgd"
        nodes                   = 3
      }
      csp_auth = false
      instance_type = {
        instance_type_id = "gcp:e2-highcpu-4"
      }
      pg_config = [
        {
          name  = "application_name"
          value = "created through terraform"
        },
        {
          name  = "array_nulls"
          value = "off"
        },
      ]
      storage = {
        volume_type       = "pd-ssd"
        volume_properties = "pd-ssd"
        size              = "32 Gi"
      }
      #      wal_storage = {
      #        volume_type       = "pd-ssd"
      #        volume_properties = "pd-ssd"
      #        size              = "32 Gi"
      #      }
      pg_type = {
        pg_type_id = "epas" #valid values ["epas", "pgextended", "postgres]"
      }
      pg_version = {
        pg_version_id = "15"
      }
      private_networking = false
      cloud_provider = {
        cloud_provider_id = "bah:gcp" // "bah:gpc" uses BigAnimal's cloud account Google Cloud provider, use "gcp" for your cloud account
      }
      region = {
        region_id = "europe-west1"
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 5
        start_time = "12:00"
      }
      # pe_allowed_principal_ids = [
      #   <example_value> # ex: "development-data-123456"
      # ]

      # service_account_ids = [
      #   <only_needed_for_bah:gcp_clusters> # ex: "test@development-data-123456.iam.gserviceaccount.com"
      # ]
    }
  ]
  witness_groups = [
    {
      region = {
        region_id = "asia-south1"
      }
      cloud_provider = {
        cloud_provider_id = "bah:gcp" // "bah:gpc" uses BigAnimal's cloud account Google Cloud provider, use "gcp" for your cloud account
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 3
        start_time = "03:00"
      }
    }
  ]
}

output "password" {
  sensitive = true
  value     = resource.biganimal_pgd.pgd_cluster.password
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_name` (String) cluster name
- `data_groups` (Attributes List) Cluster data groups. (see [below for nested schema](#nestedatt--data_groups))
- `password` (String, Sensitive) Password for the user edb_admin. It must be 12 characters or more.

### Optional

- `most_recent` (Boolean) Show the most recent cluster when there are multiple clusters with the same name
- `pause` (Boolean) Pause cluster. If true it will put the cluster on pause and set the phase as paused, if false it will resume the cluster and set the phase as healthy. Pausing a cluster allows you to save on compute costs without losing data or cluster configuration settings. While paused, clusters aren't upgraded or patched, but changes are applied when the cluster resumes. Pausing a Postgres Distributed(PGD) cluster shuts down all cluster nodes
- `project_id` (String) BigAnimal Project ID.
- `tags` (Attributes Set) Assign existing tags or create tags to assign to this resource (see [below for nested schema](#nestedatt--tags))
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `witness_groups` (Attributes List) (see [below for nested schema](#nestedatt--witness_groups))

### Read-Only

- `cluster_id` (String) Cluster ID.
- `id` (String) The ID of this resource.

<a id="nestedatt--data_groups"></a>
### Nested Schema for `data_groups`

Required:

- `backup_retention_period` (String) Backup retention period
- `cloud_provider` (Attributes) Cloud provider. (see [below for nested schema](#nestedatt--data_groups--cloud_provider))
- `cluster_architecture` (Attributes) Cluster architecture. (see [below for nested schema](#nestedatt--data_groups--cluster_architecture))
- `csp_auth` (Boolean) Is authentication handled by the cloud service provider.
- `instance_type` (Attributes) Instance type. (see [below for nested schema](#nestedatt--data_groups--instance_type))
- `maintenance_window` (Attributes) Custom maintenance window. (see [below for nested schema](#nestedatt--data_groups--maintenance_window))
- `pg_config` (Attributes Set) Database configuration parameters. (see [below for nested schema](#nestedatt--data_groups--pg_config))
- `pg_type` (Attributes) Postgres type. (see [below for nested schema](#nestedatt--data_groups--pg_type))
- `pg_version` (Attributes) Postgres version. (see [below for nested schema](#nestedatt--data_groups--pg_version))
- `private_networking` (Boolean) Is private networking enabled.
- `region` (Attributes) Region. (see [below for nested schema](#nestedatt--data_groups--region))
- `storage` (Attributes) Storage. (see [below for nested schema](#nestedatt--data_groups--storage))

Optional:

- `allowed_ip_ranges` (Attributes Set) Allowed IP ranges. (see [below for nested schema](#nestedatt--data_groups--allowed_ip_ranges))
- `backup_schedule_time` (String) Backup schedule time in 24 hour cron expression format.
- `cluster_type` (String) Type of the Specified Cluster
- `pe_allowed_principal_ids` (Set of String) Cloud provider subscription/account ID, need to be specified when cluster is deployed on BigAnimal's cloud account.
- `read_only_connections` (Boolean) Is read-only connections enabled.
- `service_account_ids` (Set of String) A Google Cloud Service Account is used for logs. If you leave this blank, then you will be unable to access log details for this cluster. Required when cluster is deployed on BigAnimal's cloud account.
- `wal_storage` (Attributes) Use a separate storage volume for Write-Ahead Logs (Recommended for high write workloads) (see [below for nested schema](#nestedatt--data_groups--wal_storage))

Read-Only:

- `cluster_name` (String) Name of the group.
- `connection_uri` (String) Data group connection URI.
- `created_at` (String) Cluster creation time.
- `group_id` (String) Group ID of the group.
- `logs_url` (String) The URL to find the logs of this cluster.
- `metrics_url` (String) The URL to find the metrics of this cluster.
- `phase` (String) Current phase of the data group.
- `resizing_pvc` (Set of String) Resizing PVC.
- `ro_connection_uri` (String) Cluster read-only connection URI.

<a id="nestedatt--data_groups--cloud_provider"></a>
### Nested Schema for `data_groups.cloud_provider`

Required:

- `cloud_provider_id` (String) Data group cloud provider id.


<a id="nestedatt--data_groups--cluster_architecture"></a>
### Nested Schema for `data_groups.cluster_architecture`

Required:

- `cluster_architecture_id` (String) Cluster architecture ID.
- `nodes` (Number) Node count.

Read-Only:

- `cluster_architecture_name` (String) Cluster architecture name.
- `witness_nodes` (Number) Witness nodes count.


<a id="nestedatt--data_groups--instance_type"></a>
### Nested Schema for `data_groups.instance_type`

Required:

- `instance_type_id` (String) Data group instance type id.


<a id="nestedatt--data_groups--maintenance_window"></a>
### Nested Schema for `data_groups.maintenance_window`

Required:

- `is_enabled` (Boolean) Is maintenance window enabled.
- `start_day` (Number) Start day.
- `start_time` (String) Start time.


<a id="nestedatt--data_groups--pg_config"></a>
### Nested Schema for `data_groups.pg_config`

Required:

- `name` (String) GUC name.
- `value` (String) GUC value.


<a id="nestedatt--data_groups--pg_type"></a>
### Nested Schema for `data_groups.pg_type`

Required:

- `pg_type_id` (String) Data group pg type id.


<a id="nestedatt--data_groups--pg_version"></a>
### Nested Schema for `data_groups.pg_version`

Required:

- `pg_version_id` (String) Data group pg version id.


<a id="nestedatt--data_groups--region"></a>
### Nested Schema for `data_groups.region`

Required:

- `region_id` (String) Data group region id.


<a id="nestedatt--data_groups--storage"></a>
### Nested Schema for `data_groups.storage`

Required:

- `volume_properties` (String) Volume properties.
- `volume_type` (String) Volume type.

Optional:

- `iops` (String) IOPS for the selected volume.
- `size` (String) Size of the volume.
- `throughput` (String) Throughput.


<a id="nestedatt--data_groups--allowed_ip_ranges"></a>
### Nested Schema for `data_groups.allowed_ip_ranges`

Required:

- `cidr_block` (String) CIDR block
- `description` (String) Description of CIDR block


<a id="nestedatt--data_groups--wal_storage"></a>
### Nested Schema for `data_groups.wal_storage`

Required:

- `size` (String) Size of the volume. It can be set to different values depending on your volume type and properties.
- `volume_properties` (String) Volume properties in accordance with the selected volume type.
- `volume_type` (String) Volume type. For Azure: "azurepremiumstorage" or "ultradisk". For AWS: "gp3", "io2", or "io2-block-express". For Google Cloud: only "pd-ssd".

Optional:

- `iops` (String) IOPS for the selected volume. It can be set to different values depending on your volume type and properties.
- `throughput` (String) Throughput is automatically calculated by BigAnimal based on the IOPS input if it's not provided.



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `tag_name` (String)

Optional:

- `color` (String)

Read-Only:

- `tag_id` (String)


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
- `delete` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
- `update` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).


<a id="nestedatt--witness_groups"></a>
### Nested Schema for `witness_groups`

Required:

- `region` (Attributes) Region. (see [below for nested schema](#nestedatt--witness_groups--region))

Optional:

- `cloud_provider` (Attributes) Witness Group cloud provider id. It can be set during creation only and can be different than the cloud provider of the data groups. Once set, cannot be changed. (see [below for nested schema](#nestedatt--witness_groups--cloud_provider))
- `maintenance_window` (Attributes) Custom maintenance window. (see [below for nested schema](#nestedatt--witness_groups--maintenance_window))

Read-Only:

- `cluster_architecture` (Attributes) Cluster architecture. (see [below for nested schema](#nestedatt--witness_groups--cluster_architecture))
- `cluster_type` (String) Type of the Specified Cluster
- `group_id` (String) Group id of witness group.
- `instance_type` (Attributes) Instance type. (see [below for nested schema](#nestedatt--witness_groups--instance_type))
- `phase` (String) Current phase of the witness group.
- `storage` (Attributes) Storage. (see [below for nested schema](#nestedatt--witness_groups--storage))

<a id="nestedatt--witness_groups--region"></a>
### Nested Schema for `witness_groups.region`

Required:

- `region_id` (String) Region id.


<a id="nestedatt--witness_groups--cloud_provider"></a>
### Nested Schema for `witness_groups.cloud_provider`

Optional:

- `cloud_provider_id` (String) Cloud provider id.


<a id="nestedatt--witness_groups--maintenance_window"></a>
### Nested Schema for `witness_groups.maintenance_window`

Required:

- `is_enabled` (Boolean) Is maintenance window enabled.

Optional:

- `start_day` (Number) The day of week, 0 represents Sunday, 1 is Monday, and so on.
- `start_time` (String) Start time. "hh:mm", for example: "23:59".


<a id="nestedatt--witness_groups--cluster_architecture"></a>
### Nested Schema for `witness_groups.cluster_architecture`

Read-Only:

- `cluster_architecture_id` (String) Cluster architecture ID.
- `cluster_architecture_name` (String) Name.
- `nodes` (Number) Nodes.
- `witness_nodes` (Number) Witness nodes count.


<a id="nestedatt--witness_groups--instance_type"></a>
### Nested Schema for `witness_groups.instance_type`

Read-Only:

- `instance_type_id` (String) Witness group instance type id.


<a id="nestedatt--witness_groups--storage"></a>
### Nested Schema for `witness_groups.storage`

Read-Only:

- `iops` (String) IOPS for the selected volume.
- `size` (String) Size of the volume.
- `throughput` (String) Throughput.
- `volume_properties` (String) Volume properties.
- `volume_type` (String) Volume type.

## Import

Import is supported using the following syntax:

```shell
# terraform import biganimal_pgd.<resource_name> <project_id>/<cluster_id>
terraform import biganimal_pgd.pgd_cluster prj_deadbeef01234567/p-abcd123456
```
