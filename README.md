# Sensu Go Relay Handler

The [Sensu Go][1] Relay handler is a [Sensu Event Handler][2] that relays Events to another Sensu Go installation.

[Site B Agent] -> (Event) -> [Site B Backend] -> [Relay Handler] -> [Site A Agent] -> [Site A Backend]

## Configuration

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
[2]: https://docs.sensu.io/sensu-go/5.0/reference/handlers/#how-do-sensu-handlers-work
