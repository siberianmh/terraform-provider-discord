# Terraform Provider for Discord

## Installation

### Requirements

`terraform-provider-discord` based on Terraform, this means you need

- [Terraform](https://www.terraform.io/) >=0.13

## Usage

```terraform
provider "discord" {
  bot_token = "DISCORD_BOT_TOKEN"
  guild_id = "ID_OF_THE_GUILD"
}

resource "discord_channel" "proj-party" {
  channel_name = "proj-party"
}
```

## LICENSE

[MIT](LICENSE.md)
