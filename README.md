# alertmanager-bot

![Golang release](https://github.com/Nayls/alertmanager-bot/workflows/Golang%20release/badge.svg?branch=master)
[![DockerHub](https://images.microbadger.com/badges/version/nayls/alertmanager-bot.svg)](https://hub.docker.com/repository/docker/nayls/alertmanager-bot)

## Configurations

Example `config.yaml`

```bash
hostname: "0.0.0.0"
port: "8080"
debug: "false"

discord:
  username: "Discord AlertBot"
  webhook: ""
  avatar_url: ""

grafana:
  internal_url: "https://grafana.localhost:9090"
  external_url: "http://localhost:8080"
  basic_username: ""
  basic_password: ""

severity:
  critical:
    - "<@243687813197856768>" # @nayls
  warning:
    - "<@243687813197856768>" # @nayls
  none:
    - "<@243687813197856768>" # @nayls
  normal:
    - "<@243687813197856768>" # @nayls
```

### Variables

|Name |Default |Descriptions |ENV |
|---|---|---|---|
| hostname             | _"0.0.0.0"_ | Host service | APP_HOSTNAME |
| port                 | _8080_ | Port service | APP_PORT |
| debug                | _false_ | Debug mode | APP_DEBUG |
| discord:username     | _Discord AlertBot_ | Username bot in Discord | APP_DISCORD_USERNAME |
| discord:webhook      | _""_ | Webhook url in Discord | APP_DISCORD_WEBHOOK |
| discord:avatar_url   | _""_ | Image discord bot in message | APP_DISCORD_AVATAR_URL |
| grafana:internal_url | _http://localhost:8080_ | Url bot | APP_GRAFANA_INTERNAL_URL |
| grafana:external_url | _""_ | Url grafana | APP_GRAFANA_EXTERNAL_URL |
| grafana:basic_username | _""_ | Basic auth for grafana | APP_GRAFANA_BASIC_USERNAME |
| grafana:basic_password | _""_ | Basic password for grafana | APP_GRAFANA_BASIC_PASSWORD |
