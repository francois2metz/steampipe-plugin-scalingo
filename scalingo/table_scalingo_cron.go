package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableScalingoCron() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_cron",
		Description: "A cron task is a command executed at a scheduled interval.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.SingleColumn("app_name"),
			Hydrate:           listCron,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},

			{Name: "command", Type: proto.ColumnType_STRING, Description: "The cron expression followed by the command."},
			{Name: "size", Type: proto.ColumnType_STRING, Description: "The size of the one-off container."},
			{Name: "last_execution_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date of the last execution."},
			{Name: "next_execution_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date of the next execution."},
		},
	}
}

func listCron(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_cron.listCron", "connection_error", err)
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()

	tasks, err := client.CronTasksGet(appName)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_cron.listCron", err)
		return nil, err
	}
	for _, job := range tasks.Jobs {
		d.StreamListItem(ctx, job)
	}
	return nil, nil
}
