# Terraform Provider Biganimal PoC (Terraform Plugin SDK)

This repository is manually created from the [Terraform Plugin Scaffolding](https://github.com/hashicorp/terraform-provider-scaffolding) repository.
It's in experimental phase, and intends to be the playground for the Biganimal Terraform Provider.

Biganimal Resource and Datasource definitions are under (`biganimal/provider`)

For now, we're using the `examples/provider/provider.tf` for development purposes. You can run `terraform plan` in that folder.

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.18


## Debugging the provider

If you're using Vscode, you can use the embedded Golang Debugger. Intro to debugging in Vscode is [here](https://code.visualstudio.com/docs/editor/debugging).

There is already a .vscode/launch.json file, so you can easily run `Debug - Attach External CLI` in the Run and Debug section, which is going to print a `TF_REATTACH_PROVIDERS` env var if your code builds successfully. The env var looks similar to this one:
```
$> TF_REATTACH_PROVIDERS='{"registry.terraform.io/hashicorp/biganimal":{"Protocol":"grpc","ProtocolVersion":5,"Pid":14123,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/99/kt3b7rgn7wbcc55jt9zv_rch0000gn/T/plugin608643082"}}}'
```

You can navigate to `examples/provider` directory and run your terraform commands with this env var:
```
$> TF_REATTACH_PROVIDERS='{"registry.terraform.io/hashicorp/biganimal":{"Protocol":"grpc","ProtocolVersion":5,"Pid":14123,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/99/kt3b7rgn7wbcc55jt9zv_rch0000gn/T/plugin608643082"}}}' terraform plan
```

For more information about Vscode Golang debugging, please refer to [this documentation](https://github.com/golang/vscode-go/blob/master/docs/debugging.md).

## Next steps, what can you do?

* [biganimal/provider/data_source_cluster.go](./biganimal/provider/data_source_cluster.go) is a good starting point. You can investigate the contents of the Json response and how it's [unmarshalled to clusters struct](https://github.com/EnterpriseDB/terraform-provider-biganimal/blob/b96336c0cb06d481854cdf0db0688de694e02a71/biganimal/provider/data_source_cluster.go#L103).
  * Add/remove new fields, experiment with how to consume those values, like using them in the resource, or printing their value in the outputs.
* <i>**[[API client]]**</i> You can set up an apiClient inside [biganimal/provider/provider.go](./biganimal/provider/provider.go) and configure it, so that this client can be used in data sources and resources. There are code crumbs in that provider.go file, they don't work for now 😅
* You can start experimenting with [biganimal/provider/resource_cluster.go](biganimal/provider/resource_cluster.go). It has the blueprints for the CRUD operations. Feel free to start writing them.
* While experimenting with the API calls, feel free to propose any schema designs. That would be very helpful for the RFC.
* <i>**[[API client]]**</i> [biganimal/openapi directory](./biganimal/openapi/)  is a direct copy of [the generated API Golang SDK v2 from the upm-cli repository](https://github.com/EnterpriseDB/upm-cli/tree/yh-8468-sdk-v2/generated/apiv2) (See [this PR](https://github.com/EnterpriseDB/upm-cli/pull/242) for details.) This is a temporary situation, It's not decided yet where to reside this library.
* Tests generate is failing after each commit in Github(See [the GH Actions](https://github.com/EnterpriseDB/terraform-provider-biganimal/actions)). If you can fix/silence it, that would be nice 😊


---
# Clean this part if you don't need anymore -- Terraform Provider Scaffolding (Terraform Plugin SDK)

_This template repository is built on the [Terraform Plugin SDK](https://github.com/hashicorp/terraform-plugin-sdk). The template repository built on the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework) can be found at [terraform-provider-scaffolding-framework](https://github.com/hashicorp/terraform-provider-scaffolding-framework). See [Which SDK Should I Use?](https://www.terraform.io/docs/plugin/which-sdk.html) in the Terraform documentation for additional information._

This repository is a *template* for a [Terraform](https://www.terraform.io) provider. It is intended as a starting point for creating Terraform providers, containing:

 - A resource, and a data source (`internal/provider/`),
 - Examples (`examples/`) and generated documentation (`docs/`),
 - Miscellaneous meta files.

These files contain boilerplate code that you will need to edit to create your own Terraform provider. Tutorials for creating Terraform providers can be found on the [HashiCorp Learn](https://learn.hashicorp.com/collections/terraform/providers) platform.

Please see the [GitHub template repository documentation](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template) for how to create a new repository from this template on GitHub.

Once you've written your provider, you'll want to [publish it on the Terraform Registry](https://www.terraform.io/docs/registry/providers/publishing.html) so that others can use it.


## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.18

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:
```sh
$ go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

Fill this in for each provider

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
