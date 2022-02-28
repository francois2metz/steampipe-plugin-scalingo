package scalingo

import (
	"context"

	"github.com/Scalingo/go-scalingo/v4"
	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
)

func tableScalingoDeployment() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_deployment",
		Description: "A deployment represent a new version deployed of an application.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.SingleColumn("app_name"),
			Hydrate:           listDeployment,
			ShouldIgnoreError: isNotFoundError,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.AllColumns([]string{"app_name", "id"}),
			Hydrate:           getDeployment,
			ShouldIgnoreError: isNotFoundError,
		},
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID identifying the event."},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "ID of the application where the event belong."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Creation date of the event."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status of the deployment."},
			{Name: "git_ref", Type: proto.ColumnType_STRING, Description: "Git SHA."},
			{Name: "image", Type: proto.ColumnType_STRING, Description: "Link to the resulting image."},
			{Name: "registry", Type: proto.ColumnType_STRING, Description: "Name of the registry."},
			{Name: "duration", Type: proto.ColumnType_INT, Description: "Duration of the deployment."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.ID"), Description: "Unique id of the user."},
			{Name: "user_username", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Username"), Description: "Username of the user."},
			{Name: "user_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Email"), Description: "Email of the user."},
		},
	}
}

func listDeployment(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()
	opts := scalingo.PaginationOpts{Page: 1, PerPage: 50}

	for {
		deployments, pagination, err := client.DeploymentListWithPagination(appName, opts)
		if err != nil {
			return nil, err
		}
		for _, deployment := range deployments {
			d.StreamListItem(ctx, deployment)
		}
		if pagination.NextPage == 0 {
			break
		}
		opts.Page = pagination.NextPage
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
