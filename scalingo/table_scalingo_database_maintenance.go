package scalingo

import (
	"context"

	"github.com/Scalingo/go-scalingo/v9"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableScalingoDatabaseMaintenance() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_database_maintenance",
		Description: "Maintenance operations on a database.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"app_name", "addon_id"}),
			Hydrate:    listDatabaseMaintenance,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isAddonTokenError,
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"app_name", "addon_id", "id"}),
			Hydrate:    getDatabaseMaintenance,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "app_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("app_name"),
				Description: "Name of the app.",
			},
			{
				Name:        "addon_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("addon_id"),
				Description: "ID of the addon.",
			},
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Unique ID identifying the maintenance.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "Status of the maintenance (scheduled, notified, queued, cancelled, running, failed, done).",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "Type of the maintenance (no-op, failing).",
			},
			{
				Name:        "started_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Start date of the maintenance.",
			},
			{
				Name:        "ended_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "End date of the maintenance.",
			},
		},
	}
}

func listDatabaseMaintenance(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("scalingo_database_maintenance.listDatabaseMaintenance", "connection_error", err)
		return nil, err
	}
	appName := d.EqualsQuals["app_name"].GetStringValue()
	addon := d.EqualsQuals["addon_id"].GetStringValue()
	opts := scalingo.PaginationOpts{Page: 1, PerPage: 50}

	for {
		maintenances, pagination, err := client.DatabaseListMaintenance(ctx, appName, addon, opts)
		if err != nil {
			plugin.Logger(ctx).Error("scalingo_database_maintenance.listDatabaseMaintenance", err)
			return nil, err
		}

		for _, maintenance := range maintenances {
			d.StreamListItem(ctx, maintenance)
		}
		if pagination.NextPage == 0 {
			break
		}
		opts.Page = pagination.NextPage
	}
	return nil, nil
}

func getDatabaseMaintenance(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("scalingo_database_maintenance.getDatabaseMaintenance", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals

	id := quals["id"].GetStringValue()
	appName := quals["app_name"].GetStringValue()
	addon := quals["addon_id"].GetStringValue()

	result, err := client.DatabaseShowMaintenance(ctx, appName, addon, id)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_database_maintenance.getDatabaseMaintenance", err)
		return nil, err
	}
	return result, nil
}
