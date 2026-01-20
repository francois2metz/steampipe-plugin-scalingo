package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableScalingoAlert() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_alert",
		Description: "Alert based on an application metric.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listAlert,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"app_name", "id"}),
			Hydrate:    getAlert,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID, starts with “al-“."},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "ID of the application this alert applies to."},
			{Name: "container_type", Type: proto.ColumnType_STRING, Description: "Container type concerned by the alert."},
			{Name: "disabled", Type: proto.ColumnType_BOOL, Description: "Is the alert disabled."},
			{Name: "metric", Type: proto.ColumnType_STRING, Description: "Metric name this alert is about."},
			{Name: "limit", Type: proto.ColumnType_DOUBLE, Description: "Threshold to activate the alert."},
			{Name: "send_when_below", Type: proto.ColumnType_BOOL, Description: "Will the alert be sent when the value goes above or below the limit."},
			{Name: "duration_before_trigger", Type: proto.ColumnType_INT, Description: "Alert is triggered if the value is above the limit for the specified duration activated."},
			{Name: "remind_every", Type: proto.ColumnType_STRING, Description: "Send the alert at regular interval when activated."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Creation date of the alert."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Last time the alert has been updated."},
			{Name: "notifiers", Type: proto.ColumnType_JSON, Description: "Notifiers used by the alert."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Various data."},
		},
	}
}

func listAlert(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_alert.listAlert", "connection_error", err)
		return nil, err
	}
	appName := d.EqualsQuals["app_name"].GetStringValue()

	alerts, err := client.AlertsList(ctx, appName)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_alert.listAlert", err)
		return nil, err
	}
	for _, alert := range alerts {
		d.StreamListItem(ctx, alert)
	}
	return nil, nil
}

func getAlert(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_alert.getAlert", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals

	id := quals["id"].GetStringValue()
	appName := quals["app_name"].GetStringValue()

	result, err := client.AlertShow(ctx, appName, id)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_alert.getAlert", err)
		return nil, err
	}
	return result, nil
}
