package zendesk

import (
	"context"
	"errors"
	"os"

	"github.com/scalingo/go-scalingo"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*scalingo.Client, error) {
	endpoint := os.Getenv("SCALINGO_ENDPOINT")
	token := os.Getenv("SCALINGO_TOKEN")

	scalingoConfig := GetConfig(d.Connection)
	if &scalingoonfig != nil {
		if scalingokConfig.Endpoint != nil {
			endpoint = *scalingokConfig.Endpoint
		}
		if scalingoConfig.Token != nil {
			token = *scalingoConfig.Token
		}
	}

	if endpoint == "" {
		return nil, errors.New("'endpoint' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	config := scalingo.ClientConfig{
		APIEndpoint: endpoint,
		APIToken: token,
	}
	return scalingo.New(config)
}
