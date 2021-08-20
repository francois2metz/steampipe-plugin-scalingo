package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableScalingoCollaborators() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_collaborators",
		Description: "Get collaborators from a specific app",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listCollaborator,
		},
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Hydrate: appNameQual, Transform: transform.FromValue(), Description: "The name of the app"},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "unique ID identifying the addon"},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "ID of the application where the collaborator belong"},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "status of the invitation"},
			{Name: "username", Type: proto.ColumnType_STRING, Description: "username of the collaborator"},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "email of the collaborator"},
			{Name: "user_id", Type: proto.ColumnType_STRING, Description: "user id of the collaborator"},
		},
	}
}

func listCollaborator(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()

	collaborators, err := client.CollaboratorsList(appName)
	if err != nil {
		return nil, err
	}
	for _, collaborator := range collaborators {
		d.StreamListItem(ctx, collaborator)
	}
	return nil, nil
}
