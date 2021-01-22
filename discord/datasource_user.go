package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscordUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDiscordUserRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"real_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceDiscordUserRead(d *schema.ResourceData, meta interface{}) error {
	api, err := discordgo.New("Bot " + meta.(*Config).APIToken)

	if err != nil {
		return err
	}

	email := d.Get("email").(string)
	fmt.Printf("[Info] Reading Discord user '%s'", email)

	user, err := api.User("@me")
	if err != nil {
		return err
	}

	if user.Email == email {
		fmt.Printf("[Debug] Discord user: %s", user)
		d.SetId(user.ID)
		d.Set("name", user.Username)
		d.Set("email", user.Email)
		return nil
	}

	return fmt.Errorf("Invalid user with email: '%s'", email)
}
