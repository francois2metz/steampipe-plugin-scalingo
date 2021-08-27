package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableScalingoAddon() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_addon",
		Description: "Get addons from a specific app",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listAddon,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"app_name", "id"}),
			Hydrate:    getAddon,
		},
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Hydrate: appNameQual, Transform: transform.FromValue(), Description: "The name of the app"},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "unique ID identifying the addon"},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "ID of the application which owns the addon"},
			{Name: "resource_id", Type: proto.ColumnType_STRING, Description: "resource reference"},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "current status of the addon"},
			{Name: "provider_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("AddonProvider.ID"), Description: "id of the provider"},
			{Name: "provider_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("AddonProvider.Name"), Description: "name of the provider"},
			{Name: "provider_logo_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("AddonProvider.LogoURL"), Description: "Logo url of the provider"},
			{Name: "plan_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Plan.ID"), Description: "id of the plan"},
			{Name: "plan_logo_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Plan.LogoURL"), Description: "logo url of the plan"},
			{Name: "plan_display_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Plan.DisplayName"), Description: "display name of the plan"},
			{Name: "plan_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Plan.Name"), Description: "name of the plan"},
			{Name: "plan_description", Type: proto.ColumnType_STRING, Transform: transform.FromField("Plan.Description"), Description: "description of the plan"},
			{Name: "plan_price", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Plan.Price"), Description: "price of the plan"},
			{Name: "plan_sku", Type: proto.ColumnType_STRING, Transform: transform.FromField("Plan.SKU"), Description: "SKU of the plan"},
		},
	}
}

func listAddon(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()

	addons, err := client.AddonsList(appName)
	if err != nil {
		return nil, err
	}
	for _, addon := range addons {
		d.StreamListItem(ctx, addon)
	}
	return nil, nil
}

func getAddon(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	quals := d.KeyColumnQuals

	id := quals["id"].GetStringValue()
	appName := quals["app_name"].GetStringValue()

	result, err := client.AddonShow(appName, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
