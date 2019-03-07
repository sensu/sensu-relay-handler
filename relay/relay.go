package relay

import (
	"fmt"
	"net/http"

	sensuhttp "github.com/sensu/sensu-enterprise-go-plugin/http"
	"github.com/sensu/sensu-enterprise-go-plugin/sensu"
	"github.com/sensu/sensu-go/types"
)

type Config struct {
	sensu.HandlerConfig
	URL      string
	User     string
	Password string
}

// The Relay parameters required to connect to the Rest API
type ConfigOptions struct {
	URL      sensu.HandlerConfigOption
	User     sensu.HandlerConfigOption
	Password sensu.HandlerConfigOption
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

	statusCode, _, err := relay.httpWrapper.ExecuteRequest(http.MethodPost, relay.config.URL, event, nil)

	if err != nil {
		return err
	}

	if statusCode != 201 && statusCode != 202 {
		return fmt.Errorf("could not relay event")
	}

	return nil
}
