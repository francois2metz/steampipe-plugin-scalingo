package scalingo

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Scalingo/go-scalingo/v6"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

const matrixKeyRegion = "region"
const defaultScalingoRegion = "osc-fr1"

func connect(ctx context.Context, d *plugin.QueryData) (*scalingo.Client, error) {
	region := d. EqualsQualString(matrixKeyRegion)
	if region == "" {
		region = defaultScalingoRegion
	}

	// get scalingo client from cache
	cacheKey := fmt.Sprintf("scalingo-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*scalingo.Client), nil
	}

	token := os.Getenv("SCALINGO_TOKEN")

	scalingoConfig := GetConfig(d.Connection)
	if scalingoConfig.Token != nil {
		token = *scalingoConfig.Token
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file or set the SCALINGO_TOKEN environment variable and then restart Steampipe")
	}

	config := scalingo.ClientConfig{
		APIToken: token,
		Region:   region,
	}
	client, err := scalingo.New(ctx, config)

	if err != nil {
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}

func BuildRegionList(ctx context.Context, d *plugin.QueryData) []map[string]interface{} {
	// cache matrix
	cacheKey := "RegionListMatrix"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.([]map[string]interface{})
	}

	var regions []string

	// retrieve regions from connection config
	scalingoConfig := GetConfig(d.Connection)

	// handle compatibility with the old region configuration
	if scalingoConfig.Region != nil {
		regions = append(regions, *scalingoConfig.Region)
	}

	// Get only the regions as required by config file
	if len(*scalingoConfig.Regions) > 0 {
		regions = *scalingoConfig.Regions
	}

	if len(regions) > 0 {
		matrix := make([]map[string]interface{}, len(regions))
		for i, region := range regions {
			matrix[i] = map[string]interface{}{matrixKeyRegion: region}
		}

		// set cache
		d.ConnectionManager.Cache.Set(cacheKey, matrix)
		return matrix
	}

	matrix := []map[string]interface{}{
		{matrixKeyRegion: defaultScalingoRegion},
	}

	// set cache
	d.ConnectionManager.Cache.Set(cacheKey, matrix)
	return matrix
}
