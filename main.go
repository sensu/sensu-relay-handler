package main

import (
	"fmt"

	"github.com/sensu/sensu-go/types"
	"github.com/sensu/sensu-plugins-go-library/sensu"
	"github.com/sensu/sensu-relay-handler/relay"
)

var (
	relayConfig = relay.Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "sensu-relay-handler",
			Short:    "The Sensu Go handler for relaying Events to another Sensu Go installation",
			Timeout:  10,
			Keyspace: "sensu.io/plugins/relay/config",
		},
	}

	relayConfigOptions = relay.ConfigOptions{
		URL: sensu.PluginConfigOption{
			Path:      "api-url",
			Env:       "RELAY_API_URL",
			Argument:  "api-url",
			Shorthand: "a",
			Default:   "http://127.0.0.1:3031/events",
			Usage:     "The Sensu Go Agent Events API URL",
			Value:     &relayConfig.URL,
		},
		User: sensu.PluginConfigOption{
			Path:      "username",
			Env:       "RELAY_API_USERNAME",
			Argument:  "username",
			Shorthand: "u",
			Default:   "",
			Usage:     "The Sensu Go Events API username",
			Value:     &relayConfig.User,
		},
		Password: sensu.PluginConfigOption{
			Path:      "password",
			Env:       "RELAY_API_PASSWORD",
			Argument:  "password",
			Shorthand: "p",
			Default:   "",
			Usage:     "The Sensu Go Events API user password",
			Value:     &relayConfig.Password,
		},
		DisableCheckHandling: sensu.PluginConfigOption{
			Path:      "disable-check-handling",
			Env:       "RELAY_DISABLE_CHECK_HANDLING",
			Argument:  "disable-check-handling",
			Shorthand: "d",
			Default:   false,
			Usage:     "Disable Event Handling for relayed Check Events",
			Value:     &relayConfig.DisableCheckHandling,
		},
		DisableMetricsHandling: sensu.PluginConfigOption{
			Path:      "disable-metrics-handling",
			Env:       "RELAY_DISABLE_METRICS_HANDLING",
			Argument:  "disable-metrics-handling",
			Shorthand: "D",
			Default:   false,
			Usage:     "Disable Event Handling for relayed Metrics Events",
			Value:     &relayConfig.DisableMetricsHandling,
		},
		CheckHandlers: sensu.PluginConfigOption{
			Path:      "check-handlers",
			Env:       "RELAY_CHECK_HANDLERS",
			Argument:  "check-handlers",
			Shorthand: "c",
			Default:   "",
			Usage:     "The Sensu Go Event Handlers to set in relayed Check Events (replace)",
			Value:     &relayConfig.CheckHandlers,
		},
		MetricsHandlers: sensu.PluginConfigOption{
			Path:      "metrics-handlers",
			Env:       "RELAY_METRICS_HANDLERS",
			Argument:  "metrics-handlers",
			Shorthand: "m",
			Default:   "",
			Usage:     "The Sensu Go Event Handlers to set in relayed Metrics Events (replace)",
			Value:     &relayConfig.MetricsHandlers,
		},
	}

	options = []*sensu.PluginConfigOption{
		&relayConfigOptions.URL,
		&relayConfigOptions.User,
		&relayConfigOptions.Password,
		&relayConfigOptions.DisableCheckHandling,
		&relayConfigOptions.DisableMetricsHandling,
		&relayConfigOptions.CheckHandlers,
		&relayConfigOptions.MetricsHandlers,
	}
)

func main() {
	goHandler := sensu.NewGoHandler(&relayConfig.PluginConfig, options, checkArgs, executeHandler)
	goHandler.Execute()
}

func checkArgs(_ *types.Event) error {
	if len(relayConfig.URL) == 0 || relayConfig.URL == "" {
		return fmt.Errorf("--api-url or RELAY_API_URL environment variable is required")
	}

	return nil
}

func executeHandler(event *types.Event) error {
	relay, err := relay.NewRelay(&relayConfig)
	if err != nil {
		return err
	}

	return relay.SendEvent(event)
}
