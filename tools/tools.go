//go:build tools
// +build tools

package tools

import (
	// document generation
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
	// automatically read .env file as environment variables
	_ "github.com/joho/godotenv/cmd/godotenv"
)
