package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableScalingoEventType() *plugin.Table {
	return &plugin.Table{
		Name:          "scalingo_event_type",
		Description:   "The list of event types to get IDs to create Notifiers.",
		GetMatrixItem: BuildRegionList,
		List: &plugin.ListConfig{
			Hydrate: listEventType,
		},
		Columns: []*plugin.Column{
			{Name: "region", Type: proto.ColumnType_STRING, Transform: transform.FromMatrixItem(matrixKeyRegion), Description: "The region associated to this event."},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID of event type."},
			{Name: "category_id", Type: proto.ColumnType_STRING, Description: "Category ID of the type."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Camel case name of the type."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Fancy name of the type."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description these events are produced."},
		},
	}
}

func listEventType(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_event_type.listEventType", "connection_error", err)
		return nil, err
	}
	event_types, err := client.EventTypesList(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_event_type.listEventType", err)
		return nil, err
	}
	for _, event_type := range event_types {
		d.StreamListItem(ctx, event_type)
	}

	return nil, nil
}
