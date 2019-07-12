# Mattermost Bot Sample

## Overview

This Bot is made mainly to make rekord employees laugh :)

This Bot is using Go v.1.12.7

## Setup

1 - Make sure the server is running, check its IP and port.

2 - Create the user the bot will run as.
```
user create --email="bot@example.com" --password="bot_password" --username="bot_username"
```

3 - Add the bot user to the desired team
```
team add team_name bot_username
```

4 - Verify bot's e-mail address.
```
user verify samplebot
```
5 - Ask colleagues when they want the English Day to be or just choose a random one.

5 - modify the [configuration file](config.json) following the template below
```
{
	"Server": "Mattermost_server_IP",
	"Protocol": "http/https/mattermost",
	"Name": "bot_username",
	"Password": "bot_password", 
	"TeamName": "team_name",
	"EnglishDay": "englishday_weekday"
}
```

6 - Run the bot and enjoy it!