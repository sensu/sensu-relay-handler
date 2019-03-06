# Sensu Go Relay Handler

The [Sensu Go][1] Relay handler is a [Sensu Event Handler][2] that relays Events to another Sensu Go installation.

## Configuration

Example Sensu Go handler definition:

```json
{
    "api_version": "core/v2",
    "type": "Handler",
    "metadata": {
        "namespace": "default",
        "name": "relay"
    },
    "spec": {
        "type": "pipe",
        "command": "sensu-relay-handler -api-url http://127.0.0.1:3031/events -t 10",
        "timeout": 12
    }
}
```

## Usage Examples

```
The Sensu Go handler for relaying Events to another Sensu Go installation

Usage:
  sensu-relay-handler [flags]

Flags:
  -a, --api-url string    The Sensu Go Agent or Backend Events API URL (default "http://127.0.0.1:3031/events")
  -h, --help              help for sensu-relay-handler
  -p, --password string   The Sensu Go Events API user password
  -u, --username string   The Sensu Go Events API username
```

[1]: https://github.com/sensu/sensu-go
[2]: https://docs.sensu.io/sensu-go/5.0/reference/handlers/#how-do-sensu-handlers-work
