package main

import (
	"github.com/siberianmh/terraform-provider-discord/discord"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: discord.Provider,
	})
}
