# Sensu Go Relay Handler

[![Bonsai Asset Badge](https://img.shields.io/badge/Sensu%20Relay%20Handler-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/sensu/sensu-relay-handler) [![Build Status](https://travis-ci.com/sensu/sensu-relay-handler.svg?token=D3sR2y7qtwxXTz3VygZw&branch=master)](https://travis-ci.com/sensu/sensu-relay-handler)

The [Sensu Go][1] Relay handler is a [Sensu Event Handler][2] that relays Events to another Sensu Go installation.

![Relay Handler](https://raw.githubusercontent.com/sensu/sensu-relay-handler/master/images/relay_handler.png)

# Overview

The `sensu-relay-handler` is designed to allow users to forward events to an alternate sensu-go-backend instance. This can be used to forward events from one environment to a separate sensu-go-backend instance for event handling. The `sensu-relay-handler` does not relay back subscriptions and checks. As such, a sensu-go-backend is still required to manage the remote agents, subscriptions and checks.

## Enterprise Plugin

The Sensu Go Relay Handler is an Enterprise plugin that requires a valid Sensu license to run.
Sensu Go >= 5.21 will add the `SENSU_LICENSE_FILE` environment variable to the handler execution.
To run the plugin independently of Sensu (ex. test/dev), you must set the env var:

```
SENSU_LICENSE_FILE=$(sensuctl license info --format json)
```

## Setup

The relay-agent is a sensu-go-agent that is local to your remote site network. It will send events to another site's sensu-go-backend. Events will be queued up as normal between an agent and backend.

### Remote Site (Site 1)

The remote site should be configured as a standalone Sensu-Go installation. The agents will connect to their local sensu-backend. In order for a check's results to be forwarded to another sensu-backend it must have the `sensu-relay-handler` configured as one of its handlers. Agents do not require access to the relay agent, but the sensu-backend does require network access to the relay agent. The relay agent also requires a network connection to the sensu-backend you want to forward events to. This may be over a WAN connection, VPN tunnel, etc.

### Local Site (Site 2)

The local sensu-backend needs to be made available to the relay agent. This may be achieved via a public internet connection, VPN or similar. No further configuration is required of the sensu-backend.

### Recommendations

No extra configuration or checks are required to have a working relay handler and agent. However, it is recommended that you have checks in place to verify that the relay agent is available. This is especially recommended for instances where there may not be a reliable connection between the relay-agent and sensu-backend.

## Relay Handler Configuration

Example Sensu Go handler definition:

```yaml
api_version: core/v2
type: Handler
metadata:
  namespace: default
  name: relay
spec:
  type: pipe
  runtime_assets:
  - sensu-relay-handler
  command: sensu-relay-handler --api-url http://127.0.0.1:3031/events --disable-check-handling
  timeout: 30
```

Example Sensu Go check definition:

**check-dummy-app-healthz.json**

```json
{
    "api_version": "core/v2",
    "type": "CheckConfig",
    "metadata": {
        "namespace": "default",
        "name": "dummy-app-healthz"
    },
    "spec": {
        "command": "check-http -u http://localhost:8080/healthz",
        "subscriptions":[
            "dummy"
        ],
        "publish": true,
        "interval": 10,
        "handlers": [
            "relay"
        ]
    }
}
```

**check-dummy-app-healthz.yml**

```yaml
---
api_version: core/v2
type: CheckConfig
metadata:
  namespace: default
  name: dummy-app-healthz
spec:
  command: check-http -u http://localhost:8080/healthz
  subscriptions:
  - dummy
  publish: true
  interval: 10
  handlers:
  - relay
```



## Asset configuration

### Asset registration

Assets are the best way to make use of this handler. If you're not using an asset, please consider doing so! If you're using sensuctl 5.13 or later, you can use the following command to add the asset: 

`sensuctl asset add sensu/sensu-relay-handler`

If you're using an earlier version of sensuctl, you can download the asset definition from [this project's Bonsai Asset Index page](https://bonsai.sensu.io/assets/sensu/sensu-relay-handler).

## Usage Examples

```
The Sensu Go handler for relaying Events to another Sensu Go installation

Usage:
  sensu-relay-handler [flags]

Flags:
  -a, --api-url string             The Sensu Go Agent Events API URL (default "http://127.0.0.1:3031/events")
  -c, --check-handlers string      The Sensu Go Event Handlers to set in relayed Check Events (replace)
  -d, --disable-check-handling     Disable Event Handling for relayed Check Events
  -D, --disable-metrics-handling   Disable Event Handling for relayed Metrics Events
  -h, --help                       help for sensu-relay-handler
  -m, --metrics-handlers string    The Sensu Go Event Handlers to set in relayed Metrics Events (replace)
  -p, --password string            The Sensu Go Events API user password
  -u, --username string            The Sensu Go Events API username
```


## Installing from source and contributing

The preferred way of installing and deploying this plugin is to use it as an [asset][5]. If you would like to compile and install the plugin from source, or contribute to it, download the latest version of the sensu-relay-handler from [releases][4],
or create an executable script from this source.

From the local path of the relay-handler repository:
```
go build -o /usr/local/bin/sensu-relay-handler main.go
```


[1]: https://github.com/sensu/sensu-go
[2]: https://docs.sensu.io/sensu-go/latest/reference/handlers/#how-do-sensu-handlers-work
[3]: https://bonsai.sensu.io/assets/sensu/sensu-relay-handler
[4]: https://github.com/sensu/sensu-relay-handler/releasesS
[5]: https://docs.sensu.io/sensu-go/latest/guides/install-check-executables-with-assets
