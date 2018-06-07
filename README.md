# Slack Service

```
mesg-core service deploy https://github.com/mesg-foundation/service-slack.git
```

## Tasks

### Notify

#### inputs
| input | type | description |
| --- | --- | --- |
| endpoint | String | Create a webhook here https://slack.com/apps/new/A0F7XDUAZ-incoming-webhooks |
| text | String | Content of the text you want to send |
| icon_emoji | String | Emoji for the bot picture. Find the emoji in your slack |
| username | String | Name of the user that will post the message |

#### outputs
| ouput | description |
| --- | --- |
| success | Result when everything is fine |
| error | Result when an error happend |