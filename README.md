# wl-control - mikrotik wl controller

The service gives you an oportunity to add IP adressess to white-lists from Telegram messages.

## How to use
- [Create a telegram bot](https://core.telegram.org/bots/features#botfather)
- Allow the bot read the messages from groups `BotFather -> Mybots -> Pick your bo -> Bot settings -> Group Privacy -> Turn on`
- Add your new bot to a telegram group where you want to use it
- Fill the config file `configs/app.yml`. You can see the example in `configs/app.example.yml`. it has comments for every block
- Start the program, if there is not errors, try to send an IP to the telegram group,the bot will ask you about IP adding


## Info
- After the start - the program try to connect for all devices you set for health check and write errors to a logfile (and stderr)
- Bot does not handle private messages. Only messages from groups
- Bot ignoring local and unspecified ip adressess (like 127.0.0.0, 0.0.0.0, 172.X.X.X, 10.X.X.X and etc)
- Bot allow add subnets (like `31.146.221.0/48`) only admin users or all users and admins chat (configure in config)
