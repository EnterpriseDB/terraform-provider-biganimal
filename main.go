package main

import (
	"context"
	"flag"
	"log"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

// these will be set by the goreleaser configuration
// to appropriate values for the compiled binary
var version string = "dev" // goreleaser can also pass the specific commit if you want
// commit  string = ""

func main() {
	var debugMode bool
	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	ctx := context.Background()
	upgradedSdkServer, err := tf5to6server.UpgradeServer(
		ctx,
		func() tfprotov5.ProviderServer {
			return schema.NewGRPCProviderServer(provider.NewSDKProvider(version)())
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	providers := []func() tfprotov6.ProviderServer{
		providerserver.NewProtocol6(provider.NewFrameworkProvider(version)()),
		func() tfprotov6.ProviderServer {
			return upgradedSdkServer
		},
	}

	muxServer, err := tf6muxserver.NewMuxServer(ctx, providers...)
	if err != nil {
		log.Fatal(err)
	}

	var serveOpts []tf6server.ServeOpt
	if debugMode {
		serveOpts = append(serveOpts, tf6server.WithManagedDebug())
	}

	err = tf6server.Serve(
		"registry.terraform.io/EnterpriseDB/biganimal",
		muxServer.ProviderServer,
		serveOpts...,
	)
	if err != nil {
		log.Fatal(err)
	}
}
