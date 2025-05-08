package scalingo

import (
	"context"

	"github.com/Scalingo/go-scalingo/v8"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableScalingoUserEvent() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_user_event",
		Description: "A user event is generated automatically according to your, other, or the plaform action .",
		List: &plugin.ListConfig{
			Hydrate: listUserEvent,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID identifying the event."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Creation date of the event."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the event."},
			{Name: "type_data", Type: proto.ColumnType_JSON, Description: "Data of the event."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.ID"), Description: "Unique id of the user."},
			{Name: "user_username", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Username"), Description: "Username of the user."},
			{Name: "user_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Email"), Description: "Email of the user."},
		},
	}
}

func listUserEvent(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_user_event.listUserEvent", "connection_error", err)
		return nil, err
	}
	opts := scalingo.PaginationOpts{Page: 1, PerPage: 100}

	if d.QueryContext.Limit != nil && *d.QueryContext.Limit < int64(opts.PerPage) {
		opts.PerPage = int(*d.QueryContext.Limit)
	}

	for {
		events, pagination, err := client.UserEventsList(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("scalingo_user_event.listUserEvent", err)
			return nil, err
		}
		for _, event := range events {
			d.StreamListItem(ctx, event)
		}
		if pagination.NextPage == 0 {
			break
		}
		opts.Page = pagination.NextPage
		if d.RowsRemaining(ctx) <= 0 {
			break
		}
	}
	return nil, nil
}
