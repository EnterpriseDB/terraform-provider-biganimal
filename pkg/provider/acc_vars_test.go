package provider_test

import (
	"fmt"
	"os"
	"strings"
)

const acc_default_externalID = "" // TODO:fill the value
const acc_default_projectID = ""  // TODO:fill the value
const acc_default_provider = ""   // TODO:fill the value
const acc_default_region = ""     // TODO:fill the value
const acc_default_roleARN = ""    // TODO:fill the value

func getResourceVarOrDefault(resourceName, varName, default_val string) string {
	if value, ok := os.LookupEnv(fmt.Sprintf("BA_TF_ACC_RESOURCE_%s_%s",
		strings.ToUpper(resourceName),
		strings.ToUpper(varName))); ok {
		return value
	}
	return default_val
}

func getDataSourceVarOrDefault(resourceName, varName, default_val string) string {
	if value, ok := os.LookupEnv(fmt.Sprintf("BA_TF_ACC_DATASOURCE_%s_%s",
		strings.ToUpper(resourceName),
		strings.ToUpper(varName))); ok {
		return value
	}
	return default_val
}
