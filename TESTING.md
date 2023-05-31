# Testing

We have basically 2 types of tests: [Acceptance tests](#acceptance-tests) and [Unit tests](#unit-tests).

## Acceptance Tests:

We're using [terraform-plugin-testing](https://github.com/hashicorp/terraform-plugin-testing) Golang module for writing acceptance tests.

For more information, please refer to the acceptance tests documentation.

In order to run the acceptance tests, several environment variables should be present. `BA_API_URI` and `BA_BEARER_TOKEN` to communicate with the Biganimal API.
We also use the `BA_TF_ACC_VAR_<resource_type>_<variable_name>` environment variables to run the acceptance tests. Example variable names can be found in the [.env.example](.env.example) file.

You can run the acceptance tests with the following command:
```
$> make testacc
```
or if you would like to run an individual test, For example:
```
$> TF_ACC=1 go test -timeout 600s -run ^TestAccResourceCluster_basic$ github.com/EnterpriseDB/terraform-provider-biganimal/pkg/provider
```
For more details, please refer to [Running Acceptance Tests](https://developer.hashicorp.com/terraform/plugin/testing/acceptance-tests#running-acceptance-tests).

### Adding a new env var for your Acceptance tests:

When you want to introduce a new env var for your tests, please make sure to pay attention to those points:

1. The environment variable name is in the form of `BA_TF_ACC_VAR_<resource_type>_<variable_name>`
1. Add the variable to the `acc_env_vars_checklist`.
1. Make sure that the `testAccResourcePreCheck` is called in the PreCheck function of the test. For example:
```
func TestAccBiganimalRegionResource_basic(t *testing.T) {
    var (
		acc_env_vars_checklist = []string{
			"BA_TF_ACC_VAR_region_project_id",
            ...
		}

        ...
		projectID  = os.Getenv("BA_TF_ACC_VAR_region_project_id")
        ...
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccResourcePreCheck(t, "region", acc_env_vars_checklist)
		},
        ...
```
4. Also, add the new env var to the [.env.example](.env.example) file.


## Unit Tests:

Especially for isolated functions, unit tests are essential to ensure that the plugin code works.
Please refer to [Unit Testing](https://developer.hashicorp.com/terraform/plugin/testing/unit-testing) page of the official terraform documentation for further details on the topic.

You can run the acceptance tests with the following command:
```
make test
```

## Other Resources:

* [Official Golang Docs on Testing](https://pkg.go.dev/testing)
* [Testing Patterns](https://developer.hashicorp.com/terraform/plugin/testing/testing-patterns): Official terraform documentation that covers some test patterns that are common and considered a best practice to have when developing and verifying your Terraform plugins.
