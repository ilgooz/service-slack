# Slack

Send a notification through Slack's API

```
mesg-core service deploy https://github.com/mesg-foundation/service-slack.git
```

## Tasks

### notify

Task key: **notify**



#### Inputs

| **key** | **type** | **description** |
| --- | --- | --- |
| **endpoint** | `String` | Create a webhook here https://slack.com/apps/new/A0F7XDUAZ-incoming-webhooks |
| **icon_emoji** | `String` | Emoji for the bot picture. Find the emoji in your Slack |
| **text** | `String` | Content of the text you want to send |
| **username** | `String` | Name of the user that will post the message |


#### Outputs

##### error

Output key: **error**

Result when an error happend

| **key** | **type** | **description** |
| --- | --- | --- |
| **message** | `String` | The description of the error |

##### success

Output key: **success**

Result when everything is as expected

| **key** | **type** | **description** |
| --- | --- | --- |




