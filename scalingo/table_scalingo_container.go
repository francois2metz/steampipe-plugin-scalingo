package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableScalingoContainer() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_container",
		Description: "Container from an application.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listContainer,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the container."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Creation date of the container."},
			{Name: "deleted_at", Type: proto.ColumnType_TIMESTAMP, Description: "Deletion date of the container."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of container (web, worker, etc.)."},
			{Name: "type_index", Type: proto.ColumnType_INT, Description: "Index number of the container for the given type."},
			{Name: "label", Type: proto.ColumnType_STRING, Description: "Label of the container."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "State of the container."},
			{Name: "container_size_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("ContainerSize.ID"), Description: "ID of the container size."},
			{Name: "container_size_sku", Type: proto.ColumnType_STRING, Transform: transform.FromField("ContainerSize.SKU"), Description: "Stock Keeping Unit aka Products catalog ID."},
			{Name: "container_size_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("ContainerSize.Name"), Description: "Name of the size, used as parameter in operations."},
			{Name: "container_size_human_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("ContainerSize.HumanName"), Description: "Display name of the type."},
			{Name: "container_size_human_cpu", Type: proto.ColumnType_STRING, Transform: transform.FromField("ContainerSize.HumanCPU"), Description: "Human representation of the CPU priority."},
			{Name: "container_size_memory", Type: proto.ColumnType_INT, Transform: transform.FromField("ContainerSize.Memory"), Description: "RAM allocated to the containers in bytes."},
			{Name: "container_size_hourly_price", Type: proto.ColumnType_INT, Transform: transform.FromField("ContainerSize.HourlyPrice"), Description: "Price per hour of this container size in cents."},
			{Name: "container_size_thirtydays_price", Type: proto.ColumnType_INT, Transform: transform.FromField("ContainerSize.ThirtydaysPrice"), Description: "Price for 30 days in cents."},
			{Name: "container_size_ordinal", Type: proto.ColumnType_INT, Transform: transform.FromField("ContainerSize.Ordinal"), Description: "Sorting index to display a list of sizes."},
		},
	}
}

func listContainer(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_container.listContainer", "connection_error", err)
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()

	containers, err := client.AppsContainersPs(ctx, appName)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_container.listContainer", err)
		return nil, err
	}
	for _, container := range containers {
		d.StreamListItem(ctx, container)
	}
	return nil, nil
}
