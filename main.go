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
			Usage:     "The Sensu Go Agent or Backend Events API URL",
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
	}

	options = []*sensu.HandlerConfigOption{
		&relayConfigOptions.URL,
		&relayConfigOptions.User,
		&relayConfigOptions.Password,
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
	if len(relayConfig.URL) == 0 {
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
