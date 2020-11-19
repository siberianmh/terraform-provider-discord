package discord

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"bot_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DISCORD_API_TOKEN", nil),
				Description: "Discord Authentication Token for discord.com/api",
			},
			"guild_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Guild ID.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"discord_channel": resourceChannel(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		APIToken: d.Get("bot_token").(string),
		GuildID:  d.Get("guild_id").(string),
	}
	return config, nil
}
