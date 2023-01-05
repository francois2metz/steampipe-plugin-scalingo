package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableScalingoNotifier() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_notifier",
		Description: "Notifications are app events sent to a custom HTTP endpoint.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listNotifier,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID identifying the notifier."},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "ID of the application which owns the notifier."},
			{Name: "active", Type: proto.ColumnType_BOOL, Description: "Is the notifier active or not."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the notifier."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Notifier type."},
			{Name: "send_all_events", Type: proto.ColumnType_BOOL, Description: "Should the notifier accept all events."},
			{Name: "send_all_alerts", Type: proto.ColumnType_BOOL, Description: "Should the notifier accept all alerts."},
			{Name: "selected_event_ids", Type: proto.ColumnType_JSON, Description: "List of events accepted by this notifier.", Transform: transform.FromField("SelectedEventIDs")},
			{Name: "type_data", Type: proto.ColumnType_JSON, Description: "Notification platform dependant additional data"},
			{Name: "platform_id", Type: proto.ColumnType_STRING, Description: "Notification platform used by this notifer."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Date of creation of the notifier."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When was the notifier updated."},
		},
	}
}

func listNotifier(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_notifier.listNotifier", "connection_error", err)
		return nil, err
	}
	appName := d.EqualsQuals["app_name"].GetStringValue()

	notifiers, err := client.NotifiersList(ctx, appName)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_notifier.listNotifier", err)
		return nil, err
	}
	for _, notifier := range notifiers {
		d.StreamListItem(ctx, notifier)
	}
	return nil, nil
}
