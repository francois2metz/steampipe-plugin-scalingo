package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableScalingoDomain() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_domain",
		Description: "A domain name associated to an application.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.SingleColumn("app_name"),
			Hydrate:           listDomain,
			ShouldIgnoreError: isNotFoundError,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.AllColumns([]string{"app_name", "id"}),
			Hydrate:           getDomain,
			ShouldIgnoreError: isNotFoundError,
		},
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID identifying the domain."},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "ID of the application where the domain belong."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Hostname your want to associate with the app."},
			{Name: "tls_cert", Type: proto.ColumnType_STRING, Transform: transform.FromField("TLSCert"), Description: "Subject of the submitted certificate."},
			{Name: "tls_key", Type: proto.ColumnType_STRING, Transform: transform.FromField("TLSKey"), Description: "Private key type and length."},
			{Name: "ssl", Type: proto.ColumnType_BOOL, Transform: transform.FromField("SSL"), Description: "Flag if SSL with a custom certificate is enabled."},
			{Name: "validity", Type: proto.ColumnType_DATETIME, Description: "Once a certificate has been submitted, display the validity of it."},
			{Name: "canonical", Type: proto.ColumnType_BOOL, Description: "The domain is the canonical domain of this application."},
			{Name: "lets_encrypt", Type: proto.ColumnType_BOOL, Description: "The domain is using a Let’s Encrypt certificate."},
			{Name: "lets_encrypt_status", Type: proto.ColumnType_STRING, Description: "Let’s Encrypt certificate generation status."},
			{Name: "ssl_status", Type: proto.ColumnType_STRING, Description: "SSL certificate status (pending, success, error)."},
			{Name: "acme_dns_fqdn", Type: proto.ColumnType_STRING, Transform: transform.FromField("AcmeDNSFqdn"), Description: "ACME DNS-01 TXT entry FQDN."},
			{Name: "acme_dns_value", Type: proto.ColumnType_STRING, Transform: transform.FromField("AcmeDNSValue"), Description: "ACME DNS-01 TXT entry value."},
			{Name: "acme_dns_error", Type: proto.ColumnType_STRING, Transform: transform.FromField("AcmeDNSError"), Description: "ACME DNS-01 error."},
		},
	}
}

func listDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()

	domains, err := client.DomainsList(appName)
	if err != nil {
		return nil, err
	}
	for _, domain := range domains {
		d.StreamListItem(ctx, domain)
	}
	return nil, nil
}

func getDomain(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	quals := d.KeyColumnQuals

	id := quals["id"].GetStringValue()
	appName := quals["app_name"].GetStringValue()

	result, err := client.DomainsShow(appName, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
