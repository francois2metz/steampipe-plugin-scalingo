package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableScalingoNotificationPlatform() *plugin.Table {
	return &plugin.Table{
		Name:              "scalingo_notification_platform",
		Description:       "Scalingo lets you use different platforms to send your notifications.",
		GetMatrixItemFunc: BuildRegionList,
		List: &plugin.ListConfig{
			Hydrate: listNotificationPlatform,
		},
		Columns: []*plugin.Column{
			{Name: "region", Type: proto.ColumnType_STRING, Transform: transform.FromMatrixItem(matrixKeyRegion), Description: "The region associated to this notification platform."},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID identifying the notification platform."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the notification platform."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Human readable name for this notification platform."},
			{Name: "logo_url", Type: proto.ColumnType_STRING, Description: "URL to a logo for this notification platform."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the platform."},
			{Name: "available_event_ids", Type: proto.ColumnType_JSON, Description: "List of event IDs accepted by this platform.", Transform: transform.FromField("AvailableEventIDs")},
		},
	}
}

func listNotificationPlatform(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_notification_platform.listNotificationPlatform", "connection_error", err)
		return nil, err
	}
	notification_platforms, err := client.NotificationPlatformsList(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_notification_platform.listNotificationPlatform", err)
		return nil, err
	}
	for _, notification_platform := range notification_platforms {
		d.StreamListItem(ctx, notification_platform)
	}

	return nil, nil
}
