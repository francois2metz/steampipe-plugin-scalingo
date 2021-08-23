package scalingo

import (
	"context"
	"errors"
	"os"

	"github.com/Scalingo/go-scalingo/v4"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

const defaultEndpointUrl = "https://api.osc-fr1.scalingo.com"

func connect(ctx context.Context, d *plugin.QueryData) (*scalingo.Client, error) {
	endpoint := os.Getenv("SCALINGO_ENDPOINT")
	token := os.Getenv("SCALINGO_TOKEN")

	scalingoConfig := GetConfig(d.Connection)
	if &scalingoConfig != nil {
		if scalingoConfig.Endpoint != nil {
			endpoint = *scalingoConfig.Endpoint
		}
		if scalingoConfig.Token != nil {
			token = *scalingoConfig.Token
		}
	}

	if endpoint == "" {
		endpoint = defaultEndpointUrl
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	config := scalingo.ClientConfig{
		APIEndpoint: endpoint,
		APIToken:    token,
	}
	return scalingo.New(config)
}

func appNameQual(_ context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	return d.KeyColumnQuals["app_name"].GetStringValue(), nil
}
