# Terraform Provider BigAnimal

A Terraform Provider to manage your workloads
on [EDB BigAnimal](https://www.enterprisedb.com/products/biganimal-cloud-postgresql) interacting with the BigAnimal API.
The provider is licensed under the [MPL v2](https://www.mozilla.org/en-US/MPL/2.0/).

If you are willing to contribute please read [here](./CONTRIBUTING.md).

Main links:

- [License](./LICENSE)
- [Code of Conduct](./CODE_OF_CONDUCT.md)
- [Security](./SECURITY.md)
- [Contributing](./CONTRIBUTING.md)
- [Development](./DEVELOPMENT.md)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- [Go](https://golang.org/doc/install) >= 1.18

## Using the provider

To install the BigAnimal provider, copy and paste this code into your Terraform configuration. Then,
run `terraform init`.

```hcl
terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "0.4.0"
    }
  }
}

provider "biganimal" {
  # Configuration options
  ba_bearer_token = <redacted> // See Getting an API Token section for details
  // ba_api_uri   = "https://portal.biganimal.com/api/v3" // Optional
}
```

You can also set the `BA_BEARER_TOKEN` and `BA_API_URI` env vars. When those environment variables are present, you
don't need to add any configuration options to the provider block described above.

```bash
export BA_BEARER_TOKEN=<redacted>
export BA_API_URI=https://portal.biganimal.com/api/v3
```

### Getting an API Token

In order to access the BigAnimal API, it's necessary to fetch an api bearer token and either export it into your
environment or add this token to the provider block.

Please
visit [Using the get-token script](https://www.enterprisedb.com/docs/biganimal/latest/reference/api/#using-the-get-token-script)
for more details.

## Development

Please visit the [DEVELOPMENT.md](./DEVELOPMENT.md) page for further details about development and testing.
