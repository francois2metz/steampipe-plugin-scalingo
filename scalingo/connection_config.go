package scalingo

import (
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/schema"
)

type scalingoConfig struct {
	Token   *string   `cty:"token"`
	Region  *string   `cty:"region"`
	Regions *[]string `cty:"regions"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
	"region": {
		Type: schema.TypeString,
	},
	"regions": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
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
