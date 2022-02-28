package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
)

func tableScalingoScmRepoLink() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_scm_repo_link",
		Description: "A link between an application and an SCM.",
		List: &plugin.ListConfig{
			Hydrate: listScmRepoLink,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("app_name"),
			Hydrate:           getScmRepoLink,
			ShouldIgnoreError: isNotFoundError,
		},
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID identifying the link."},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "The application ID."},
			{Name: "linker_username", Type: proto.ColumnType_STRING, Transform: transform.FromField("Linker.Username"), Description: "The size of the one-off container."},
			{Name: "linker_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Linker.Email"), Description: "The size of the one-off container."},
			{Name: "linker_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Linker.ID"), Description: "The size of the one-off container."},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Repository owner name."},
			{Name: "repo", Type: proto.ColumnType_STRING, Description: "Repository name."},
			{Name: "branch", Type: proto.ColumnType_STRING, Description: "The branch used for auto deployment."},
			{Name: "scm_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("SCMType"), Description: "The integration type."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Creation date of the link."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Last time the link was updated."},
			{Name: "auto_deploy_enabled", Type: proto.ColumnType_BOOL, Description: "True if a new deployment is trigered when the linked branch is updated."},
			{Name: "auth_integration_uuid", Type: proto.ColumnType_STRING, Description: "ID of the scm integration linked for authentication."},
			{Name: "deploy_review_apps_enabled", Type: proto.ColumnType_BOOL, Description: "Activation of the review apps feature."},
			{Name: "delete_on_close_enabled", Type: proto.ColumnType_BOOL, Description: "Delete the review app when the pull request is closed."},
			{Name: "delete_stale_enabled", Type: proto.ColumnType_BOOL, Description: "Delete the review app when there is no activity on the pull request."},
			{Name: "hours_before_delete_on_close", Type: proto.ColumnType_INT, Description: "Time to wait before deleting a review app linked to a closed pull request (in hours)."},
			{Name: "hours_before_delete_stale", Type: proto.ColumnType_INT, Description: "Time to wait for activity on the pull request before deleting the review app (in hours)"},
			{Name: "last_auto_deploy_at", Type: proto.ColumnType_TIMESTAMP, Description: "Date of the last deployment triggered by this link"},
		},
	}
}

func listScmRepoLink(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	scmRepoLinks, err := client.SCMRepoLinkList()
	if err != nil {
		return nil, err
	}
	for _, scmRepoLink := range scmRepoLinks {
		d.StreamListItem(ctx, scmRepoLink)
	}
	return nil, nil
}

func getScmRepoLink(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()

	scmRepoLink, err := client.SCMRepoLinkShow(appName)
	if err != nil {
		return nil, err
	}
	return scmRepoLink, nil
}
