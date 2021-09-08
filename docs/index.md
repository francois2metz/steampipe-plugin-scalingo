---
organization: francois2metz
category: ["public cloud"]
brand_color: "#45AFE4"
display_name: "Scalingo"
short_name: "scalingo"
description: "Steampipe plugin for querying apps, addons and more from Scalingo."
og_description: "Query Scalingo with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/francois2metz/scalingo-social-graphic.png"
icon_url: "/images/plugins/francois2metz/scalingo.svg"
---

# Scalingo + Steampipe

[Scalingo](https://scalingo.com/) provides on-demand cloud computing platforms and APIs to authenticated customers on a metered pay-as-you-go basis.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  name,
  region,
  url
from
  scalingo_app
```

```
+--------------------+----------------+------------------------------------------------------+
| name               | region         | url                                                  |
+--------------------+----------------+------------------------------------------------------+
| caresteouvert-map  | osc-fr1        | https://carestouvert.fr                              |
| caresteouvert-api  | osc-secnum-fr1 | https://caresteouvert-api.osc-secnum-fr1.scalingo.io |
+--------------------+----------------+------------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/francois2metz/scalingo/tables)**

## Get started

### Install

Download and install the latest Scalingo plugin:

```bash
steampipe plugin install francois2metz/scalingo
```

### Configuration

Installing the latest scalingo plugin will create a config file (`~/.steampipe/config/scalingo.spc`) with a single connection named `scalingo`:

```hcl
connection "scalingo" {
  plugin = "francois2metz/scalingo"

  # The API Endpoint (default is https://api.osc-fr1.scalingo.com)
  # endpoint = "https://api.osc-fr1.scalingo.com"

  # API token for your scalingo instance (required).
  #token = "tk-us-0000-0000-000000000-000000000000000"
}
```

You can also use environment variables:

- `SCALINGO_ENDPOINT` the base url for the API endpoint of the region (ex: https://api.osc-fr1.scalingo.com)
- `SCALINGO_TOKEN` for the API token (ex: tk-us-00000-0000-000)

## Get Involved

* Open source: https://github.com/francois2metz/steampipe-plugin-scalingo

## Multi-Account Connections

You may create multiple scalingo connections:
```hcl
connection "scalingo_osc {
  plugin   = "francois2metz/scalingo"
  endpoint = "https://api.osc-fr1.scalingo.com"
  token    = "tk-us-00000-0000-000"
}

connection "scalingo_secnum {
  plugin   = "francois2metz/scalingo"
  endpoint = "https://api.osc-secnum-fr1.scalingo.com"
  token    = "tk-us-00000-0000-000"
}
```

Each connection is implemented as a distinct [Postgres schema](https://www.postgresql.org/docs/current/ddl-schemas.html).  As such, you can use qualified table names to query a specific connection:

```sql
select * from scalingo_osc.scalingo_app
```

You can multi-account connections by using an [**aggregator** connection](https://steampipe.io/docs/using-steampipe/managing-connections#using-aggregators).  Aggregators allow you to query data from multiple connections for a plugin as if they are a single connection:

```
connection "scalingo_all {
  plugin      = "francois2metz/scalingo"
  type        = "aggregator"
  connections = ["scalingo_osc", "scalingo_secnum"]
}
```

Querying tables from this connection will return results from the `scalingo_osc`, and `scalingo_secnum` connections:
```sql
select * from scalingo_all.scalingo_app
```

Steampipe supports the `*` wildcard in the connection names.  For example, to aggregate all the Scalingo plugin connections whose names begin with `scalingo_`:

```hcl
connection "scalingo_all" {
  type        = "aggregator"
  plugin      = "francois2metz/scalingo"
  connections = ["scalingo_*"]
}
```
