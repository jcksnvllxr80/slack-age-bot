# Slack bot in Golang to return age based on given birthdate in format YYYY-MM-DD

## follow tutorial @2:20:55 in this video to configure the slack app

>- https://www.youtube.com/watch?v=jFfo23yIWac&t=22s

## Installation Notes:

1. mkdir slack-age-bot; cd slack-age-bot
2. go mod init github.com/<your_github_id>/slack-age-bot
3. go get "github.com/shomali11/slacker"
4. go mod tidy
5. change the name of the file in slack-age-bot/secrets/ directory from 'envvars.go.example' to 'envvars.go'
6. change the app token and the bot token in that file to be your values given by the sslack app configured
7. change the go mod init to your github as well as the package called in main.go for the secrets module