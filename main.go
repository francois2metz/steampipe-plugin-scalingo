package main

import (
    "github.com/turbot/steampipe-plugin-sdk/plugin"
    "github.com/francois2metz/steampipe-plugin-scalingo/scalingo"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{PluginFunc: scalingo.Plugin})
}
