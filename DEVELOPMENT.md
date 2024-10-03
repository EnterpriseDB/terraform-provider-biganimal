# Development

Please make sure to read the [Contributing guideline](./CONTRIBUTING.md) first.

Another common practice is using the [Biganimal CLI](https://cli.biganimal.com/).

### Using BA CLI to help initializing Provider credentials

1. [Install the BA CLI v2.0.0 or later](https://www.enterprisedb.com/docs/biganimal/latest/reference/cli/#installing-the-cli) and [jq - Command Line JSON Processor ](https://stedolan.github.io/jq/).
1. [Authenticate as a valid user and create a credential](https://www.enterprisedb.com/docs/biganimal/latest/reference/cli/#installing-the-cli). This command will direct you to your browser.
```shell
biganimal credential create \
  --name "ba-user1"
```
3. Add the following bash functions to your shellrc file (For example: `.bashrc` if you're using bash, `.zshrc` if you're using ZSH) and start a new shell.
```bash
ba_api_get_call () {
	endpoint=$1
	curl -s --request GET --header "content-type: application/json" --header "authorization: Bearer $BA_BEARER_TOKEN" --url "$BA_API_URI$endpoint"
}

ba_get_default_proj_id () {
	echo $(ba_api_get_call "/user-info" | jq -r ".data.organizationId" | cut -d"_" -f2)
}

export_BA_env_vars () {
	cred_name="${1:-ba-user1}" ## Replace "ba-user1" with your credential name, if you're using something different
	if ! biganimal cluster show -c $cred_name > /dev/null
	then
		echo "!!! Running the credential reset command now !!!"
		biganimal credential reset $cred_name
	fi
	biganimal cluster show -c $cred_name >&/dev/null
	export BA_API_URI="https://"$(biganimal credential show -o json | jq -r --arg CREDNAME "$cred_name" '.[]|select(.name==$CREDNAME).address')/api/v3
	export BA_CRED_NAME="$cred_name"
	echo "$cred_name BA_API_URI is exported."
	export TF_VAR_project_id="prj_$(ba_get_default_proj_id)"
	echo "TF_VAR_project_id terraform variable is also exported. Value is $TF_VAR_project_id"
}
```
4. Now, you can use `export_BA_env_vars` command to manage your BA_API_URI environment variable, as well as TF_VAR_project_id terraform environment variable.
```console
$> export_BA_env_vars ba-user1
ba-user1 BA_API_URI is exported.
TF_VAR_project_id terraform variable is also exported. Value is prj_0123456789abcdef
```
5. If you would like to enable bash completion for the `export_BA_env_vars` command, you can add the following lines to your shellrc file:
```
# Bash Completion for export_BA_env_vars
_export_BA_env_vars_completions()
{
  COMPREPLY=($(compgen -W "$(biganimal credential show -o json | jq '.[].name')" -- "${COMP_WORDS[1]}"))
}

complete -F _export_BA_env_vars_completions export_BA_env_vars
```
## How to manage the BigAnimal Access Key?

To manage and get an access key please refer to the readme here: [Readme](./README.md#using-the-provider) file.

## How to test your local copy of the provider?

You can run `make install` at the root of the repository to install your local copy of the provider under your `terraform.d/` folder. After that, you can configure your `~/.terraformrc` file to use either dev overrides or filesystem mirroring.

Another common practice is starting a `delve` debugging session to export the  `TF_REATTACH_PROVIDERS` environment variable in your terminal.

Please check the following sections for details.


### dev overrides

Terraform can be configured by adding the following to your `~/.terraformrc` file.

```
provider_installation {
  # Use /home/developer/tmp/terraform-provider-biganimal as an overridden package directory
  # for the EnterpriseDB/biganimal provider. This disables the version and checksum
  # verifications for this provider and forces Terraform to look for the
  # null provider plugin in the given directory.
  dev_overrides {
      "registry.terraform.io/EnterpriseDB/biganimal" = "/home/<YOUR_HOME>/tmp/terraform-provider-biganimal"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```


### filesystem mirror

Another way to test the local code is using the filesystem_mirror declaration in your `~/.terraformrc` file.
After you compile and install your local copy with the `make install` command, the following ~/.terraformrc`
configuration allows you to declare installation path explicitly.

```
provider_installation {

  direct {
    exclude = ["registry.terraform.io/EnterpriseDB/biganimal"]
  }
  filesystem_mirror {
    path    = "/Users/<YOUR_HOME>/.terraform.d/plugins"
    include = ["registry.terraform.io/EnterpriseDB/biganimal"]
  }
}
```

For more info about `filesystem_mirror`, please
visit [Terraform Documentation of `Explicit Installation Method Configuration`](https://developer.hashicorp.com/terraform/cli/config/config-file#explicit-installation-method-configuration)


### Debugging the provider

If you're using Vscode, you can use the embedded Golang Debugger. Intro to debugging in Vscode
is [here](https://code.visualstudio.com/docs/editor/debugging).

There is already a .vscode/launch.json file, so you can easily run `Debug - Attach External CLI` in the Run and Debug
section, which is going to print a `TF_REATTACH_PROVIDERS` env var if your code builds successfully. The env var looks
similar to this one:

```bash
$> TF_REATTACH_PROVIDERS='{"registry.terraform.io/EnterpriseDB/biganimal":{"Protocol":"grpc","ProtocolVersion":5,"Pid":14123,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/99/kt3b7rgn7wbcc55jt9zv_rch0000gn/T/plugin608643082"}}}'
```

You can navigate to the folders under `examples/` directory and run your terraform commands with this env var:

```bash
$> TF_REATTACH_PROVIDERS='{"registry.terraform.io/EnterpriseDB/biganimal":{"Protocol":"grpc","ProtocolVersion":5,"Pid":14123,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/99/kt3b7rgn7wbcc55jt9zv_rch0000gn/T/plugin608643082"}}}' terraform plan
```

For more information about Vscode Golang debugging, please refer
to [this documentation](https://github.com/golang/vscode-go/blob/master/docs/debugging.md).

## Automated Testing of the provider

In order to test the provider, you can run `make testacc` or `make test`. `make test ` is for running unit
test. `make testacc` is for running acceptance test, to get more detailed information please refer
to [Acceptance Tests](https://developer.hashicorp.com/terraform/plugin/sdkv2/testing/acceptance-tests)

### Running Acceptance Tests

It's necessary to get API token for running acceptance test successfully

To run the test, you must create a copy of [.env.example](.env.example) with name `.env` and fill it with the appropriate environment variables.  Please refer
to .env for the specific values.
