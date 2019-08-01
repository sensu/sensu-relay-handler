# Sensu Go Relay Handler

The [Sensu Go][1] Relay handler is a [Sensu Event Handler][2] that relays Events to another Sensu Go installation.

![Relay Handler](images/relay_handler.png?raw=true"Title")

# Overview

The `sensu-relay-handler` is designed to allow users to forward events to an alternate sensu-go-backend instance. This can be used to forward events from one environment to a separate sensu-go-backend instance for event handling. The `sensu-relay-handler` does not relay back subscriptions and checks. As such, a sensu-go-backend is still required to manage the remote agents, subscriptions and checks.

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