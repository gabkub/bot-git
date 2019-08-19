# Mattermost Bot Sample

## Overview

This Bot is made mainly to make Rekord SI employees laugh :)

Go version: 1.12.7


## Setup

1 - Make sure the server is running, check its IP and port.

2 - Create the user the bot will run as.
```
mattermost user create --email="bot@example.com" --password="bot_password" --username="bot_username"
```

3 - Add the bot user to the desired team
```
mattermost team add team_name bot_username
```

4 - Verify bot's e-mail address.
```
mattermost user verify bot_username
```
5 - Ask colleagues when they want the English Day to be or just choose a random one.

6 - modify the [configuration file](bin/windows/config.json) following the template below
```
{
    "BotConfig": {
            "Server": "Mattermost_server",
            "Port": "Mattermost_server_port", (HTTP=80/HTTPS=443/default=8065)
            "BotName": "bot_username",
            "Password": "bot_password_encrypted", 
            "TeamName": "team_name",
            "EnglishDay": "englishday_weekday"
            },
    "DbConfig": {
            "Name": "database_name",
            "Server": "database_server",
            "Port": database_port,
            "User": "database_username",
            "Password": "database_user_password_encrypted",
            "Connections_warning": 50,
            "Connections_check_cron": "@every 1h",
            "Connections_log_cron": "@every 4h"
            } 
}
```

7 - Run the bot and enjoy it! Use the template:
`@bot_username <command>`