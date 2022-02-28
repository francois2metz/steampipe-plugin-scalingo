package scalingo

import (
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/schema"
)

type scalingoConfig struct {
	Endpoint *string `cty:"endpoint"`
	Token    *string `cty:"token"`
	Region   *string `cty:"region"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"endpoint": {
		Type: schema.TypeString,
	},
	"token": {
		Type: schema.TypeString,
	},
	"region": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &scalingoConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) scalingoConfig {
	if connection == nil || connection.Config == nil {
		return scalingoConfig{}
	}
	config, _ := connection.Config.(scalingoConfig)
	return config
}
