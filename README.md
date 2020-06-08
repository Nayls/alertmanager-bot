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

