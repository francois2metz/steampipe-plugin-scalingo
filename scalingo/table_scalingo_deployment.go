package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableScalingoDeployment() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_deployment",
		Description: "Get deployments from a specific app",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listDeployment,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"app_name", "id"}),
			Hydrate:    getDeployment,
		},
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Hydrate: appNameQual, Transform: transform.FromValue(), Description: "The name of the app"},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "unique ID identifying the event"},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "ID of the application where the event belong"},
			{Name: "created_at", Type: proto.ColumnType_DATETIME, Description: "creation date of the event"},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "status of the deployment"},
			{Name: "git_ref", Type: proto.ColumnType_STRING, Description: "status of the deployment"},
			{Name: "image", Type: proto.ColumnType_STRING, Description: "status of the deployment"},
			{Name: "registry", Type: proto.ColumnType_STRING, Description: "status of the deployment"},
			{Name: "duration", Type: proto.ColumnType_INT, Description: "duration of the deployment"},
			{Name: "user_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.ID"), Description: "unique id of the user"},
			{Name: "user_username", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Username"), Description: "username of the user"},
			{Name: "user_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Email"), Description: "email of the user"},
		},
	}
}

func listDeployment(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()

	deployments, err := client.DeploymentList(appName)
	if err != nil {
		return nil, err
	}
	for _, deployment := range deployments {
		d.StreamListItem(ctx, deployment)
	}
	return nil, nil
}

func getDeployment(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	quals := d.KeyColumnQuals

	id := quals["id"].GetStringValue()
	appName := quals["app_name"].GetStringValue()

	result, err := client.Deployment(appName, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
