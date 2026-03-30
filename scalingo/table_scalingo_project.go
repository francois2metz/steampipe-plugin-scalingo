package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableScalingoProject() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_project",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listProject,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getProject,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique id of the project."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the project."},
			{Name: "default", Type: proto.ColumnType_BOOL, Description: "Is this the default project."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Creation date of the project."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Last time the project has been updated."},
		},
	}
}

func listProject(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_project.listProject", "connection_error", err)
		return nil, err
	}
	projects, err := client.ProjectsList(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_project.listProject", err)
		return nil, err
	}
	for _, project := range projects {
		d.StreamListItem(ctx, project)
	}
	return nil, nil
}

func getProject(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_project.getProject", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetStringValue()
	result, err := client.ProjectGet(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_project.getProject", err)
		return nil, err
	}
	return result, nil
}
