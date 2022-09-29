package scalingo

import (
	"context"

	"github.com/Scalingo/go-scalingo/v5"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableScalingoInvoice() *plugin.Table {
	return &plugin.Table{
		Name:             "scalingo_invoice",
		Description:      "An invoice associed to the account.",
		DefaultTransform: transform.FromGo(),
		List: &plugin.ListConfig{
			Hydrate: listInvoice,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getInvoice,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID of the invoice."},
			{Name: "total_price", Type: proto.ColumnType_INT, Description: "price of this invoice (cents)."},
			{Name: "total_price_with_vat", Type: proto.ColumnType_INT, Description: "Price of this invoice including VAT (cents)."},
			{Name: "pdf_url", Type: proto.ColumnType_STRING, Description: "URL to download the PDF invoice."},
			{Name: "invoice_number", Type: proto.ColumnType_STRING, Description: "The invoice number."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "The state of this invoice (new, paid or failed)."},
			{Name: "vat_rate", Type: proto.ColumnType_INT, Description: "The VAT rate applied to this invoice (in ‰)."},
		},
	}
}

func listInvoice(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_invoice.listInvoice", "connection_error", err)
		return nil, err
	}

	opts := scalingo.PaginationOpts{Page: 1, PerPage: 100}

	if d.QueryContext.Limit != nil && *d.QueryContext.Limit < int64(opts.PerPage) {
		opts.PerPage = int(*d.QueryContext.Limit)
	}

	for {
		invoices, pagination, err := client.InvoicesList(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("scalingo_invoice.listInvoice", err)
			return nil, err
		}
		for _, invoice := range invoices {
			d.StreamListItem(ctx, invoice)
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

func getInvoice(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_invoice.getInvoice", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetStringValue()
	result, err := client.InvoiceShow(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_invoice.getInvoice", err)
		return nil, err
	}
	return result, nil
}
