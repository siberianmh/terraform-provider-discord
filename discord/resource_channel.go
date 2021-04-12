package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceChannelCreate,
		Read:   resourceChannelRead,
		Update: resourceChannelUpdate,
		Delete: resourceChannelDelete,
		Importer: &schema.ResourceImporter{
			State: resourceChannelImportState,
		},

		Schema: map[string]*schema.Schema{
			"channel_name": {
				Type:         schema.TypeString,
				Description:  "The name of Discord Channel that will be created",
				Required:     true,
				ValidateFunc: validation.IntBetween(2, 100),
			},

			"channel_topic": {
				Type:        schema.TypeString,
				Description: "Sets the topic for a channel",
				Optional:    true,
			},

			"type": {
				Type:        schema.TypeInt,
				Description: "Type of the channel (se)",
				Optional:    true,
				Default:     0,
			},

			"nsfw": {
				Type:        schema.TypeBool,
				Description: "Should this channel be marked as NSFW",
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func resourceChannelCreate(d *schema.ResourceData, meta interface{}) error {
	api, err := discordgo.New("Bot " + meta.(*Config).APIToken)
	if err != nil {
		return err
	}

	data := discordgo.GuildChannelCreateData{
		Name:  d.Get("channe_name").(string),
		Topic: d.Get("channel_topic").(string),
		Type:  discordgo.ChannelType(d.Get("type").(int)),
		NSFW:  d.Get("nsfw").(bool),
	}

	// Create Discord Channel
	channel, err := api.GuildChannelCreateComplex(meta.(*Config).GuildID, data)

	if err != nil {
		return err
	}

	d.SetId(channel.ID)

	return nil
}

func resourceChannelRead(d *schema.ResourceData, meta interface{}) error {
	api, err := discordgo.New("Bot " + meta.(*Config).APIToken)
	if err != nil {
		return err
	}

	// Checks if Discord Channel exists, if not remove resources from state
	_, err = api.Channel(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	return nil
}

func resourceChannelUpdate(d *schema.ResourceData, meta interface{}) error {
	api, err := discordgo.New("Bot " + meta.(Config).APIToken)
	if err != nil {
		return err
	}

	name := d.Get("channel_name").(string)
	if _, err := api.ChannelEdit(meta.(*Config).GuildID, name); err != nil {
		return err
	}
	return nil
}

func resourceChannelDelete(d *schema.ResourceData, meta interface{}) error {
	api, err := discordgo.New("Bot " + meta.(Config).APIToken)
	if err != nil {
		return err
	}

	// Deletes Discord Channel and clears state
	if _, err := api.ChannelDelete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceChannelImportState(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	api, err := discordgo.New("Bot " + meta.(Config).APIToken)
	if err != nil {
		return nil, err
	}

	// Check if channel exists, otherwise remove
	channel, err := api.Channel(d.Id())
	if err != nil {
		d.SetId("")
		return nil, err
	}
	d.Set("channel_name", channel.Name)
	d.Set("channel_topic", channel.Topic)

	return []*schema.ResourceData{d}, nil
}
