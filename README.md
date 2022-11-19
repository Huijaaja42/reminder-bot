# Reminder Bot

A simple Discord bot for scheduling reminders.

## Time Formats

All inputs must be UTC, but outputs will be formated by Discord to your local time.

### Explicit Format

Any common format should work, excluding those that are ambiguous, i.e. 12/10/2022 will not work.

Prefer

> `yyyy-mm-dd [hh:mm:ss]`

(Without a time it will default to midnight.)

### Relative Format

You can also use the following format to specify a time in relation to the current time.

> `in x [s|min|h|d|w|mon|y]`

> s = seconds
>
> min = minutes
>
> h = hours
>
> d = days
>
> w = weeks
>
> mon = months
>
> y = years

(You can also write seconds, minutes etc., that will still match the correct one.)

## Permissions

The bot needs to have permissions to send messages as well as both `bot` and `applications.commands` scopes.

## Config

The bot will only respond on the specified guild (server).

`scheduleInterval` will determine how often the bot checks for reminders due to be sent. By default, this is set to 60 seconds.

## Running the bot

1. Create a Discord bot in the [Discord Developer Portal](https://discord.com/developers/applications)

2. Make a copy of [config.json.example](config.json.example) and fill in your bot token and guild id.

3. Invite the bot to your server.

4. Start the bot.

The bot will print an invite to stdout when starting. If you want to invite it before that, use this:

> `https://discord.com/api/oauth2/authorize?client_id={BOT ID}&permissions=2048&scope=bot%20applications.commands`

## Demo Server

You can test the bot [here](https://discord.gg/VZx3qRgYDb).

## Built With

* [github.com/objectbox/objectbox-go](https://github.com/objectbox/objectbox-go)
* [github.com/bwmarrin/discordgo](https://github.com/bwmarrin/discordgo)
* [github.com/araddon/dateparse](https://github.com/araddon/dateparse)
* [github.com/spf13/viper](https://github.com/spf13/viper)

## License

This project is licensed under the GNU GPLv3 License - see the [LICENSE](LICENSE) file for details.
