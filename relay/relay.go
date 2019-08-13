package relay

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sensu/sensu-go/types"
	sensuhttp "github.com/sensu/sensu-plugins-go-library/http"
	"github.com/sensu/sensu-plugins-go-library/sensu"
)

type Config struct {
	sensu.PluginConfig
	URL                    string
	User                   string
	Password               string
	DisableCheckHandling   bool
	DisableMetricsHandling bool
	CheckHandlers          string
	MetricsHandlers        string
}

// The Relay parameters required to connect to the Rest API
type ConfigOptions struct {
	URL                    sensu.PluginConfigOption
	User                   sensu.PluginConfigOption
	Password               sensu.PluginConfigOption
	DisableCheckHandling   sensu.PluginConfigOption
	DisableMetricsHandling sensu.PluginConfigOption
	CheckHandlers          sensu.PluginConfigOption
	MetricsHandlers        sensu.PluginConfigOption
}

type Relay struct {
	config      Config
	httpWrapper sensuhttp.HttpWrapper
}

// Creates a new Relay
func NewRelay(config *Config) (*Relay, error) {

	httpWrapper, err := sensuhttp.NewHttpWrapper(uint64(config.Timeout), "", config.User, config.Password)
	if err != nil {
		return nil, fmt.Errorf("could not create http wrapper: %s", err.Error())
	}

	return &Relay{
		*config,
		*httpWrapper,
	}, nil
}

// Relay an Event
func (relay *Relay) SendEvent(event *types.Event) error {
	event.Entity.EntityClass = "proxy"

	if event.HasCheck() {
		if relay.config.CheckHandlers != "" {
			event.Check.Handlers = strings.Split(relay.config.CheckHandlers, ",")
		}

		if relay.config.DisableCheckHandling {
			event.Check.Handlers = []string{}
		}
	}

	if event.HasMetrics() {
		if relay.config.MetricsHandlers != "" {
			event.Metrics.Handlers = strings.Split(relay.config.MetricsHandlers, ",")
		}

		if relay.config.DisableMetricsHandling {
			event.Metrics.Handlers = []string{}
		}
	}

	statusCode, _, err := relay.httpWrapper.ExecuteRequest(http.MethodPost, relay.config.URL, event, nil)

	if err != nil {
		return err
	}

	if statusCode != 201 && statusCode != 202 {
		return fmt.Errorf("could not relay event")
	}

	return nil
}
