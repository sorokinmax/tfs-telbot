# tfs-telbot

## Prerequisites
I didn't find a working solution for forwarding [Azure DevOps Webhooks](https://learn.microsoft.com/en-us/azure/devops/service-hooks/services/webhooks?view=azure-devops "Azure DevOps Webhooks") to Telegram, so I wrote my own.

## Description
tfs-telbot has endpoints for catching Azure DevOps Webhooks. Upon receipt, the event is converted to an "edible" format for the Telegram API and sent to the specified notification channel.

**This project is under development. **
Now the application can handle the following webhooks:
- Build completed
- Release deployment completed

I plan to add processing of workitems-related webhooks soon.



## Usage
1. Compile source code or use [ready image from DockerHub](httphttps://hub.docker.com/repository/docker/sorokinmx/tfs-telbot:// "ready image from DockerHub")
2. Configure the config.yml file using the example.
