package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/siberianmh/terraform-provider-discord/discord"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: discord.Provider,
	})
}
