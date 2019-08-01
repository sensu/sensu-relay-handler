# Sensu Go Relay Handler

The [Sensu Go][1] Relay handler is a [Sensu Event Handler][2] that relays Events to another Sensu Go installation.

![Relay Handler](images/relay_handler.png?raw=true"Title")

# Overview

The `sensu-relay-handler` is designed to allow users to forward events to an alternate sensu-go-backend instance. This can be used to forward events from one enviroment and sent to a different sensu-go-backend instance for handling. This will not relay back subscriptions and checks, as such you still require a sensu-go-backend to manage the remote agents, subscriptions and checks.

## Setup

The relay-agent is a sensu-go-agent that will be local to your remote site network. It will send events to another site's sensu-go-backend. Events will be queued up as normal between an agent and backend.

### Remote Site

The remote site should be configured as you would any other standalone Sensu-Go installation and configuration. The agents will connect to their local sensu-backend. The only addition would be of configuring checks to use the relay handler for those checks you want to be forward to another sensu instance. Agents do not have to have access to the relay agent but the sensu-backend does require network acess to the relay agent. The relay agent should also be able to connect to the main sensu-backend you want to send events to. This may be over a WAN connection, VPN tunnel, etc.

### Local Site

Your local Sensu-backend should be made available to your relay agent. Note that you may need to make it available via an internet connection, VPN or something similar. No further configuration is required on

### Recommendations

While no extra configurationa or checks are required to have a working relay handler and agent, it is recommended that you have checks in place to verify that the relay agent is available. This would be especially true for instances where you may not have a reliable connection between your relay-agent and sensu-backend.

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

## Asset configuration

See [Sensu-Relay-Handler](3) at bonsai.sensu.io for asset creation info.

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

[1]: https://github.com/sensu/sensu-go
[2]: https://docs.sensu.io/sensu-go/latest/reference/handlers/#how-do-sensu-handlers-work
[3]: https://bonsai.sensu.io/assets/sensu/sensu-relay-handler