package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableScalingoStack() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_stack",
		Description: "Stack is the base docker image where applications are executed in.",
		List: &plugin.ListConfig{
			Hydrate: listStack,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID of the stack."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Stack display name."},
			{Name: "base_image", Type: proto.ColumnType_STRING, Description: "Docker image used to build your app."},
			{Name: "default", Type: proto.ColumnType_BOOL, Description: "Is this the default stack for new app."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Creation date of the stack."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Human readable description of the stack."},
		},
	}
}

func listStack(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_stack.listStack", "connection_error", err)
		return nil, err
	}
	stacks, err := client.StacksList(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_stack.listStack", err)
		return nil, err
	}
	for _, stack := range stacks {
		d.StreamListItem(ctx, stack)
	}

	return nil, nil
}
