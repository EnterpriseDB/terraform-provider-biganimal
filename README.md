# Terraform Provider BigAnimal

A Terraform Provider to manage your workloads on [EDB BigAnimal](https://www.enterprisedb.com/products/biganimal-cloud-postgresql) interacting with the BigAnimal API. The provider is licensed under the [MPL v2](https://www.mozilla.org/en-US/MPL/2.0/).

If you are willing to contribute please read [here](./CONTRIBUTING.md).

Main links:

- [License](./LICENSE)
- [Code of Conduct](./CODE_OF_CONDUCT.md)
- [Security](./SECURITY.md)
- [Contributing](./CONTRIBUTING.md)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- [Go](https://golang.org/doc/install) >= 1.18

## Using the provider

To install the BigAnimal provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    biganimal = {
      source = "EnterpriseDB/biganimal"
      version = "0.1.2"
    }
  }
}

provider "biganimal" {
  # Configuration options
    ba_bearer_token = <redacted> // See Getting an API Token section for details
  // ba_api_uri   = "https://portal.biganimal.com/api/v2" // Optional
}
```

You can also set the `BA_BEARER_TOKEN` and `BA_API_URI` env vars. When those environment variables are present, you don't need to add any configuration options to the provider block described above.

```bash
export BA_BEARER_TOKEN=<redacted>
export BA_API_URI=https://portal.biganimal.com/api/v2
```

### Getting an API Token

In order to access the BigAnimal API, it's necessary to fetch an api bearer token and either export it into your environment or add this token to the provider block.

Please visit [Using the get-token script](https://www.enterprisedb.com/docs/biganimal/latest/reference/api/#using-the-get-token-script) for more details.

## Development

Please make sure to read the [Contributing guideline](./CONTRIBUTING.md) first.

### dev overrides
Terraform can be configured by adding the following to your `~/.terraformrc` file.

```
provider_installation {
  # Use /home/developer/tmp/terraform-provider-biganimal as an overridden package directory
  # for the EnterpriseDB/biganimal provider. This disables the version and checksum
  # verifications for this provider and forces Terraform to look for the
  # null provider plugin in the given directory.
  dev_overrides {
      "registry.terraform.io/EnterpriseDB/biganimal" = "/home/developer/tmp/terraform-provider-biganimal"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

### Debugging the provider

If you're using Vscode, you can use the embedded Golang Debugger. Intro to debugging in Vscode is [here](https://code.visualstudio.com/docs/editor/debugging).

There is already a .vscode/launch.json file, so you can easily run `Debug - Attach External CLI` in the Run and Debug section, which is going to print a `TF_REATTACH_PROVIDERS` env var if your code builds successfully. The env var looks similar to this one:

```bash
$> TF_REATTACH_PROVIDERS='{"registry.terraform.io/EnterpriseDB/biganimal":{"Protocol":"grpc","ProtocolVersion":5,"Pid":14123,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/99/kt3b7rgn7wbcc55jt9zv_rch0000gn/T/plugin608643082"}}}'
```

You can navigate to the folders under `examples/` directory and run your terraform commands with this env var:

```bash
$> TF_REATTACH_PROVIDERS='{"registry.terraform.io/EnterpriseDB/biganimal":{"Protocol":"grpc","ProtocolVersion":5,"Pid":14123,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/99/kt3b7rgn7wbcc55jt9zv_rch0000gn/T/plugin608643082"}}}' terraform plan
```

For more information about Vscode Golang debugging, please refer to [this documentation](https://github.com/golang/vscode-go/blob/master/docs/debugging.md).
