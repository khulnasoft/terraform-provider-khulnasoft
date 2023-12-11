package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/khulnasoft/terraform-provider-khulnasoft/khulnasoft"
)

var version string

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return khulnasoft.Provider(version)
		},
	})
}
