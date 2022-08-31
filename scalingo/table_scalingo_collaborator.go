package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableScalingoCollaborator() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_collaborator",
		Description: "A collaborator is someone who have access to an application.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listCollaborator,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID identifying the collaborator."},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "ID of the application where the collaborator belong."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status of the invitation."},
			{Name: "username", Type: proto.ColumnType_STRING, Description: "Username of the collaborator."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email of the collaborator."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Description: "User id of the collaborator."},
		},
	}
}

func listCollaborator(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_collaborator.listCollaborator", "connection_error", err)
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()

	collaborators, err := client.CollaboratorsList(ctx, appName)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_collaborator.listCollaborator", err)
		return nil, err
	}
	for _, collaborator := range collaborators {
		d.StreamListItem(ctx, collaborator)
	}
	return nil, nil
}
