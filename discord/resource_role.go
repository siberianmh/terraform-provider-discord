package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoleCreate,
		Read:   resourceRoleRead,
		Update: resourceRoleUpdate,
		Delete: resourceRoleDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"color": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"hoist": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"permissions": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mentionable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"managed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceRoleCreate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*Config).Session
	guild := meta.(*Config).GuildID

	role, err := api.GuildRoleCreate(guild)
	if err != nil {
		return err
	}

	role, err = api.GuildRoleEdit(guild, role.ID, d.Get("name").(string), d.Get("color").(int), d.Get("hoist").(bool), d.Get("permissions").(int64), d.Get("mentionable").(bool))
	if err != nil {
		return err
	}

	d.SetId(role.ID)

	return resourceRoleRead(d, meta)
}

func resourceRoleRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*Config).Session

	roles, err := api.GuildRoles(meta.(*Config).GuildID)
	if err != nil {
		return err
	}

	// Match the role by ID
	for _, role := range roles {
		if role.ID != d.Id() {
			continue
		}

		d.Set("name", role.Name)
		d.Set("color", role.Color)
		d.Set("hoist", role.Hoist)
		d.Set("position", role.Position)
		d.Set("permissions", role.Permissions)
		d.Set("managed", role.Managed)
		d.Set("mentionable", role.Mentionable)

		break
	}

	return nil
}

func resourceRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*Config).Session

	_, err := api.GuildRoleEdit(meta.(*Config).GuildID, d.Id(), d.Get("name").(string), d.Get("color").(int), d.Get("hoist").(bool), d.Get("persmissions").(int64), d.Get("mentionable").(bool))
	if err != nil {
		return err
	}

	if d.HasChange("position") {
		var orderList []*discordgo.Role
		oldPosition, newPosition := d.GetChange("position")
		roles, err := api.GuildRoles(meta.(*Config).GuildID)
		if err != nil {
			return err
		}

		for _, role := range roles {
			if role.Position == newPosition {
				orderList = append(orderList, &discordgo.Role{
					ID:       role.ID,
					Position: oldPosition.(int),
				})
				orderList = append(orderList, &discordgo.Role{
					ID:       d.Id(),
					Position: newPosition.(int),
				})
				break
			}
		}
		_, err = api.GuildRoleReorder(meta.(*Config).GuildID, orderList)
		if err != nil {
			return err
		}
	}

	return resourceRoleRead(d, meta)
}

func resourceRoleDelete(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*Config).Session
	return api.GuildRoleDelete(meta.(*Config).GuildID, d.Id())
}
