package scalingo

import (
    "context"

    "github.com/Scalingo/go-scalingo"

    "github.com/turbot/steampipe-plugin-sdk/grpc/proto"
    "github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableScalingoApps() *plugin.Table {
    return &plugin.Table{
        Name:        "scalingo_apps",
        Description: "Scalingo apps",
        List: &plugin.ListConfig{
            Hydrate: listApp,
        },
        Get: &plugin.GetConfig{
            KeyColumns: plugin.SingleColumn("name"),
            Hydrate:    getApp,
        },
        Columns: []*plugin.Column{
            {Name: "id", Type: proto.ColumnType_STRING, Description: "unique id of the appliation"},
            {Name: "name", Type: &proto.ColumnType_STRING, Description: "name of the application"},
            {Name: "created_at", Type: proto.ColumnType_DATETIME, Description: "creation date of the applciation"},
            {Name: "updated_at", Type: proto.ColumnType_DATETIME, Description: "last time the application has been updated"},
            {Name: "git_url", Type: proto.ColumnType_STRING, Description: "URL to the GIT remote to access your application"},
            {Name: "url", Type: proto.ColumnType_STRING, Description: "URL used to access your app"},
            {Name: "base_url", Type: proto.ColumnType_STRING, Description: "URL generated by Scalingo for your app"},

            {Name: "force_https", Type: proto.ColumnType_BOOL, Description: "activation of force HTTPS"},
            {Name: "sticky_session", Type: proto.ColumnType_BOOL, Description: "activation of sticky session"},
            {Name: "router_logs", Type: proto.ColumnType_BOOL, Description: "activation of the router logs in your app logs"},
            {Name: "last_deployed_at", Type: proto.ColumnType_DATETIME, Description: "date of the last deployment attempt"},
            {Name: "last_deployed_by", Type: proto.ColumnType_STRING, Description: "user who attempted the last deployment"},
            {Name: "last_deployment_id", Type: proto.ColumnType_STRING, Description: "id of the last successful deployment"},
            {Name: "stack_id", Type: proto.ColumnType_STRING, Description: "id of the stack used"},
        },
    }
}

func listApp(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx)
	if err != nil {
		return nil, err
	}
	apps, err := client.AppsList()
	if err != nil {
		return nil, err
	}
	for _, app := range apps {
		d.StreamListItem(ctx, app)
	}
	return nil, nil
}

func getApp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    client, err := connect(ctx)
    if err != nil {
        return nil, err
    }
    quals := d.KeyColumnQuals
    plugin.Logger(ctx).Warn("getApp", "quals", quals)
    name := quals["name"].GetInt64Value()
    plugin.Logger(ctx).Warn("getApp", "name", id)
    result, err := client.AppsShow(ctx, name)
    if err != nil {
        return nil, err
    }
    return result, nil
}