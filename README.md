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

## Building

Builds are done via make targets.  Running `make` will build and install the provider binary into `~/.terraform.d/plugins/...`

```bash
$ make
go build -o terraform-provider-biganimal
mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/biganimal/0.3.1/darwin_amd64
mv terraform-provider-biganimal ~/.terraform.d/plugins/hashicorp.com/edu/biganimal/0.3.1/darwin_amd64
```

The binary can also be compiled by `go build`, which will output the binary into the current directory.

## Using the provider

Until the provider is accepted into the terraform registry, it's necessary to install the binary into your local terraform cache using `make`, and to configure terraform to look in the right location to find the binary for the BigAnimal terraform provider.

Terraform can be configured by adding the following to your ~/.terraformrc file.

```hcl
provider_installation {
  dev_overrides {
      "registry.terraform.io/hashicorp/biganimal" = "/Users/<YOUR_HOME>/.terraform.d/plugins/hashicorp.com/edu/biganimal/0.3.1/<OS_ARCH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

Now, you can configure your BigAnimal provider. This can be done in one of these two ways:

### Providing BigAnimal configuration in a provider block
```hcl
provider "biganimal" {
  ba_bearer_token = <redacted> // See Getting and API Token section for details
  // ba_api_uri   = "https://portal.biganimal.com/api/v2" // Optional
}
```
### Configuring Environment Variables

You can set `BA_BEARER_TOKEN` and `BA_API_URI` env vars.

```bash
export BA_BEARER_TOKEN=<redacted>
export BA_API_URI=https://portal.biganimal.com/api/v2
```

### Getting an API Token

In order to access the BigAnimal API, it's necessary to fetch an api bearer token and export it into your environment.

This can be done by using the script located [here](https://github.com/EnterpriseDB/cloud-utilities/blob/main/api/get-token.sh) as follows

```bash
sh ~/hackery/edb/cloud-utilities/api/get-token.sh
Please login to https://auth.biganimal.com/activate?user_code=JWPL-RCXL with your BigAnimal account
Have you finished the login successfully? (y/N) y
{
  "access_token": "<REDACTED>",
  "refresh_token": "<REDACTED>",
  "scope": "openid profile email offline_access",
  "expires_in": 86400,
  "token_type": "Bearer"
}
```

Once the `access_token` has been retrieved, you can write it to the provider block as described above.
Alternatively, you can set it in your environment as follows

```bash
export BA_BEARER_TOKEN=<REDACTED>
```

After compiling, configuring the `.terraformrc` and fetching a token, the examples in the `./examples` folder can be run.

## Debugging the provider

If you're using Vscode, you can use the embedded Golang Debugger. Intro to debugging in Vscode is [here](https://code.visualstudio.com/docs/editor/debugging).

There is already a .vscode/launch.json file, so you can easily run `Debug - Attach External CLI` in the Run and Debug section, which is going to print a `TF_REATTACH_PROVIDERS` env var if your code builds successfully. The env var looks similar to this one:

```bash
$> TF_REATTACH_PROVIDERS='{"registry.terraform.io/hashicorp/biganimal":{"Protocol":"grpc","ProtocolVersion":5,"Pid":14123,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/99/kt3b7rgn7wbcc55jt9zv_rch0000gn/T/plugin608643082"}}}'
```

You can navigate to `examples/provider` directory and run your terraform commands with this env var:

```bash
$> TF_REATTACH_PROVIDERS='{"registry.terraform.io/hashicorp/biganimal":{"Protocol":"grpc","ProtocolVersion":5,"Pid":14123,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/99/kt3b7rgn7wbcc55jt9zv_rch0000gn/T/plugin608643082"}}}' terraform plan
```

For more information about Vscode Golang debugging, please refer to [this documentation](https://github.com/golang/vscode-go/blob/master/docs/debugging.md).
