# Terceira ponte status bot

#### [Add this bot to your contacts list](https://t.me/terceira_ponte_bot)

### Commands
#### /now
Get most recent images

### Running your own instance

The only thing needed to run it yourself is an telegram bot token and docker installed. The token should be stored in a `secret.yml` file like the example, just rename it and edit with your token.

After setting that up just run the following commands:

`$ git clone https://github.com/guilhermelimak/Terceira-ponte-bot.git`
`$ cd Terceira-ponte-bot`
`$ docker build . -t terceira_ponte_bot`
`$ docker run terceira_ponte_bot`

If everything went right you should now see the `Authorized on account...` message on your terminal.
