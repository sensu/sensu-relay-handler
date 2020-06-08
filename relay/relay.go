package relay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	corev2 "github.com/sensu/sensu-go/api/core/v2"
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

type Relay struct {
	config     *Config
	httpClient *http.Client
}

// Creates a new Relay
func NewRelay(config *Config) (*Relay, error) {

	relay := &Relay{
		config,
		&http.Client{
			Timeout: time.Second * time.Duration(config.Timeout),
		},
	}

	return relay, nil
}

// Relay an Event
func (relay *Relay) SendEvent(event *corev2.Event) error {
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

	response, err := relay.RelayRequest(http.MethodPost, relay.config.URL, event)

	if err != nil {
		return err
	}

	if response.StatusCode != 201 && response.StatusCode != 202 {
		return fmt.Errorf("could not relay event, http status %d", response.StatusCode)
	}

	return nil
}

func (relay *Relay) RelayRequest(method string, RequestURL string, requestBody interface{}) (*http.Response, error) {
	payload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, RequestURL, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.SetBasicAuth(relay.config.User, relay.config.Password)
	return relay.httpClient.Do(request)
}
