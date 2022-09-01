package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableScalingoAutoscaler() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_autoscaler",
		Description: "",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listAutoscaler,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"app_name", "id"}),
			Hydrate:    getAutoscaler,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID, starts with “au-“."},
			{Name: "container_type", Type: proto.ColumnType_STRING, Description: "Container type affected by the autoscaling."},
			{Name: "metric", Type: proto.ColumnType_STRING, Description: "Metric name this autoscaler is about."},
			{Name: "target", Type: proto.ColumnType_STRING, Description: "Metric value the autoscaler aims to reach."},
			{Name: "min_containers", Type: proto.ColumnType_INT, Description: "Lower limit of containers."},
			{Name: "max_containers", Type: proto.ColumnType_INT, Description: "Upper limit of containers."},
			{Name: "disabled", Type: proto.ColumnType_BOOL, Description: "Is the autoscaler disabled."},
		},
	}
}

func listAutoscaler(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_autoscaler.listAutoscaler", "connection_error", err)
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()

	autoscalers, err := client.AutoscalersList(ctx, appName)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_autoscaler.listAutoscaler", err)
		return nil, err
	}
	for _, autoscaler := range autoscalers {
		d.StreamListItem(ctx, autoscaler)
	}
	return nil, nil
}

func getAutoscaler(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_autoscaler.getAutoscaler", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals

	id := quals["id"].GetStringValue()
	appName := quals["app_name"].GetStringValue()

	result, err := client.AutoscalerShow(ctx, appName, id)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_autoscaler.getAutoscaler", err)
		return nil, err
	}
	return result, nil
}
