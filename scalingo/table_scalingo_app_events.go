package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
	"github.com/Scalingo/go-scalingo/v4"
)

func tableScalingoAppEvents() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_app_events",
		Description: "Get events from a specific app",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listAppEvent,
		},
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Description: "The name of the app"},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "unique ID identifying the event"},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "ID of the application where the event belong"},
			{Name: "created_at", Type: proto.ColumnType_DATETIME, Description: "creation date of the event"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "type of the event"},
			{Name: "user_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.ID"), Description: "unique id of the user"},
			{Name: "user_username", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Username"), Description: "username of the user"},
			{Name: "user_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Email"), Description: "email of the user"},
		},
	}
}

func listAppEvent(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()
	opts := scalingo.PaginationOpts{ Page: 1, PerPage: 100 }

	for {
		events, pagination, err := client.EventsList(appName, opts)
		if err != nil {
			return nil, err
		}
		for _, event := range events {
			d.StreamListItem(ctx, event)
		}
		if pagination.NextPage == 0 {
			break
		}
		opts.Page = pagination.NextPage

	}
	return nil, nil
}
