package scalingo

import (
	"context"
	"errors"
	"os"

	"github.com/Scalingo/go-scalingo/v4"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*scalingo.Client, error) {
	// get scalingo client from cache
	cacheKey := "scalingo"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*scalingo.Client), nil
	}

	token := os.Getenv("SCALINGO_TOKEN")
	region := os.Getenv("SCALINGO_REGION")

	scalingoConfig := GetConfig(d.Connection)
	if &scalingoConfig != nil {
		if scalingoConfig.Token != nil {
			token = *scalingoConfig.Token
		}
		if scalingoConfig.Region != nil {
			region = *scalingoConfig.Region
		}
	}

	if region == "" {
		return nil, errors.New("'region' must be set in the connection configuration. Edit your connection configuration file or set the SCALINGO_REGION environment variable and then restart Steampipe")
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file or set the SCALINGO_TOKEN environment variable and then restart Steampipe")
	}

	config := scalingo.ClientConfig{
		APIToken: token,
		Region:   region,
	}
	client, err := scalingo.New(config)

	if err != nil {
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}
