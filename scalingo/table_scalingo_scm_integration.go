package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableScalingoScmIntegration() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_scm_integration",
		Description: "A link between an application and an SCM.",
		List: &plugin.ListConfig{
			Hydrate: listScmIntegration,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("id"),
			Hydrate:           getScmIntegration,
			ShouldIgnoreError: isNotFoundError,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID identifying the SCM integration."},
			{Name: "scm_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("SCMType"), Description: "SCM type (github, gitlab, github-enterprise or gitlab-self-hosted)"},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL where the SCM platform is hosted"},
			{Name: "uid", Type: proto.ColumnType_STRING, Transform: transform.FromField("Uid"), Description: "User ID provided by the SCM platform."},
			{Name: "username", Type: proto.ColumnType_STRING, Description: "Username provided by the SCM platform."},
			{Name: "avatar_url", Type: proto.ColumnType_STRING, Description: "User avatar URL provided by the SCM platform."},
			{Name: "profile_url", Type: proto.ColumnType_STRING, Description: "User profile URL provided by the SCM platform."},
		},
	}
}

func listScmIntegration(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_scm_integration.listScmIntegration", "connection_error", err)
		return nil, err
	}

	scmIntegrations, err := client.SCMIntegrationsList()
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
	id := d.KeyColumnQuals["id"].GetStringValue()

	scmIntegration, err := client.SCMIntegrationsShow(id)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_scm_integration.getScmIntegration", err)
		return nil, err
	}
	return scmIntegration, nil
}
