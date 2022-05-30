package scalingo

import (
	"context"

	"github.com/Scalingo/go-scalingo/v4"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableScalingoAppEvent() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_app_event",
		Description: "An application event is generated automatically according to your, other or plaform action on an application.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listAppEvent,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Description: "Name of the app."},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID identifying the event."},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "ID of the application where the event belong."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Creation date of the event."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the event."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.ID"), Description: "Unique id of the user."},
			{Name: "user_username", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Username"), Description: "Username of the user."},
			{Name: "user_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Email"), Description: "Email of the user."},
		},
	}
}

func listAppEvent(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_app_event.listAppEvent", "connection_error", err)
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()
	opts := scalingo.PaginationOpts{Page: 1, PerPage: 100}

	if d.QueryContext.Limit != nil && *d.QueryContext.Limit < int64(opts.PerPage) {
		opts.PerPage = int(*d.QueryContext.Limit)
	}

	for {
		events, pagination, err := client.EventsList(appName, opts)
		if err != nil {
			plugin.Logger(ctx).Error("scalingo_app_event.listAppEvent", err)
			return nil, err
		}
		for _, event := range events {
			d.StreamListItem(ctx, event)
		}
		if pagination.NextPage == 0 {
			break
		}
		opts.Page = pagination.NextPage
		if d.QueryStatus.RowsRemaining(ctx) <= 0 {
			break
		}
	}
	return nil, nil
}
