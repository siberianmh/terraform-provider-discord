package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider retruns a Terraform Resource Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"bot_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DISCORD_API_TOKEN", nil),
				Description: "Discord Authentication Token for discord.com/api",
			},
			"guild_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Guild ID.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"discord_user": dataSourceDiscordUser(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"discord_channel": resourceChannel(),
			"discord_role":    resourceRole(),
		},

		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	api, err := discordgo.New("Bot ", d.Get("bot_token").(string))
	if err != nil {
		return nil, err
	}

	config := &Config{
		APIToken: d.Get("bot_token").(string),
		GuildID:  d.Get("guild_id").(string),
		Session:  api,
	}

	return config, nil
}
