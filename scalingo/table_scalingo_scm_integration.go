package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableScalingoScmIntegration() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_scm_integration",
		Description: "A link between your account and an SCM.",
		List: &plugin.ListConfig{
			Hydrate: listScmIntegration,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getScmIntegration,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID identifying the SCM integration."},
			{Name: "scm_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("SCMType"), Description: "SCM type (github, gitlab, github-enterprise or gitlab-self-hosted)"},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL where the SCM platform is hosted"},
			{Name: "uid", Type: proto.ColumnType_STRING, Description: "User ID provided by the SCM platform."},
			{Name: "username", Type: proto.ColumnType_STRING, Description: "Username provided by the SCM platform."},
			{Name: "avatar_url", Type: proto.ColumnType_STRING, Description: "User avatar URL provided by the SCM platform."},
			{Name: "profile_url", Type: proto.ColumnType_STRING, Description: "User profile URL provided by the SCM platform."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "creation date of the SCM integration."},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Owner.ID"), Description: "Unique id of the owner."},
			{Name: "owner_username", Type: proto.ColumnType_STRING, Transform: transform.FromField("Owner.Username"), Description: "Username of the owner."},
			{Name: "owner_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Owner.Email"), Description: "Email of the owner."},
		},
	}
}

func listScmIntegration(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_scm_integration.listScmIntegration", "connection_error", err)
		return nil, err
	}

	scmIntegrations, err := client.SCMIntegrationsList(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_scm_integration.listScmIntegration", err)
		return nil, err
	}
	for _, scmIntegration := range scmIntegrations {
		d.StreamListItem(ctx, scmIntegration)
	}

	return nil, nil
}

func getScmIntegration(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_scm_integration.getScmIntegration", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetStringValue()

	scmIntegration, err := client.SCMIntegrationsShow(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_scm_integration.getScmIntegration", err)
		return nil, err
	}
	return scmIntegration, nil
}
