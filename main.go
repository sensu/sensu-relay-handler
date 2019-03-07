package main

import (
	"fmt"

	"github.com/sensu/sensu-enterprise-go-plugin/sensu"
	"github.com/sensu/sensu-go/types"
	"github.com/sensu/sensu-relay-handler/relay"
)

var (
	relayConfig = relay.Config{
		HandlerConfig: sensu.HandlerConfig{
			Name:     "sensu-relay-handler",
			Short:    "The Sensu Go handler for relaying Events to another Sensu Go installation",
			Timeout:  10,
			Keyspace: "sensu.io/plugins/relay/config",
		},
	}

	relayConfigOptions = relay.ConfigOptions{
		URL: sensu.HandlerConfigOption{
			Path:      "api-url",
			Env:       "RELAY_API_URL",
			Argument:  "api-url",
			Shorthand: "a",
			Default:   "http://127.0.0.1:3031/events",
			Usage:     "The Sensu Go Agent Events API URL",
			Value:     &relayConfig.URL,
		},
		User: sensu.HandlerConfigOption{
			Path:      "username",
			Env:       "RELAY_API_USERNAME",
			Argument:  "username",
			Shorthand: "u",
			Default:   "",
			Usage:     "The Sensu Go Events API username",
			Value:     &relayConfig.User,
		},
		Password: sensu.HandlerConfigOption{
			Path:      "password",
			Env:       "RELAY_API_PASSWORD",
			Argument:  "password",
			Shorthand: "p",
			Default:   "",
			Usage:     "The Sensu Go Events API user password",
			Value:     &relayConfig.Password,
		},
		DisableCheckHandling: sensu.HandlerConfigOption{
			Path:      "disable-check-handling",
			Env:       "RELAY_DISABLE_CHECK_HANDLING",
			Argument:  "disable-check-handling",
			Shorthand: "d",
			Default:   false,
			Usage:     "Disable Event Handling for relayed Check Events",
			Value:     &relayConfig.DisableCheckHandling,
		},
		DisableMetricsHandling: sensu.HandlerConfigOption{
			Path:      "disable-metrics-handling",
			Env:       "RELAY_DISABLE_METRICS_HANDLING",
			Argument:  "disable-metrics-handling",
			Shorthand: "D",
			Default:   false,
			Usage:     "Disable Event Handling for relayed Metrics Events",
			Value:     &relayConfig.DisableMetricsHandling,
		},
		CheckHandlers: sensu.HandlerConfigOption{
			Path:      "check-handlers",
			Env:       "RELAY_CHECK_HANDLERS",
			Argument:  "check-handlers",
			Shorthand: "c",
			Default:   "",
			Usage:     "The Sensu Go Event Handlers to set in relayed Check Events (replace)",
			Value:     &relayConfig.CheckHandlers,
		},
		MetricsHandlers: sensu.HandlerConfigOption{
			Path:      "metrics-handlers",
			Env:       "RELAY_METRICS_HANDLERS",
			Argument:  "metrics-handlers",
			Shorthand: "m",
			Default:   "",
			Usage:     "The Sensu Go Event Handlers to set in relayed Metrics Events (replace)",
			Value:     &relayConfig.MetricsHandlers,
		},
	}

	options = []*sensu.HandlerConfigOption{
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
	goHandler := sensu.NewGoHandler(&relayConfig.HandlerConfig, options, checkArgs, executeHandler)
	err := goHandler.Execute()
	if err != nil {
		fmt.Printf("Error executing plugin: %s", err)
	}
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
