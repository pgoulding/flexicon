# Flexicon Slackbot
This is a Slackbot written in GoLang that can be used to define terms from the F3 Lexicon/Exicon. It can be hosted on AWS Lambda and integrated with your Slack workspace to make it easy for your team to look up the meanings of F3 terms.

## Installation
### Local Installation
1. Clone the repository to your local machine using the command:

```bash
git clone https://github.com/pgoulding/flexicon.git
```

2. Navigate to the cloned repository using the command:

```bash
cd flexicon
```

3. Install dependencies using the command:

```go
go mod download
```

4. Rename the `.env.example` file to `.env` and update the values with your Slack API token and signing secret.

5. Start the application using the command:

```go
go run main.go
```

### Slack Installation
1. Create a new Slack app by visiting https://api.slack.com/apps and clicking the "Create New App" button.
2. Navigate to the "Bot" section and add a new bot user.
3. Navigate to the "Event Subscriptions" section and enable events.
4. Enter your Lambda URL in the "Request URL" field, followed by `/events`.
5. Subscribe to the `app_mention` event.
6. Navigate to the "OAuth & Permissions" section and add the following bot token scopes:
   1. `app_mentions:read`
   2. `chat:write`
7. Install the app to your workspace.

### AWS Lambda Installation
Zip the contents of the cloned repository (excluding the `vendor` directory) using the command:

```python
zip f3-lexicon-slackbot.zip *
```

1. Create a new Lambda function by visiting the AWS Lambda console and clicking the "Create Function" button.
2. Select "Author from scratch", give your function a name, and choose Go 1.x as the runtime.
3. Upload the f3-lexicon-slackbot.zip file to the function.
4. Configure the environment variables with your Slack API token and signing secret.
5. Configure the function to use the handler main and set the timeout to at least 10 seconds.
6. Create a new API Gateway trigger for the function and configure it to use the "HTTP API" type.
7. Deploy the API.
8. Update your Slack app's "Request URL" to use the URL of the deployed API.

## Usage
To use the bot, simply mention it in a channel with the term you want to look up. For example:

```
@flexicon define burpee
```

The bot will respond with the definition of the term, if it exists in the F3 Lexicon/Exicon.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.