connection "scalingo" {
    plugin = "francois2metz/scalingo"

    # API token for your scalingo instance (required).
    # Get it on: https://dashboard.scalingo.com/account/tokens
    # This can also be set via the `SCALINGO_TOKEN` environment variable.
    # token = "tk-us-0000-0000-000000000-000000000000000"

    # Regions
    # By default the scalingo plugin will only use the osc-fr1 region
    # regions = ["osc-fr1", "osc-secnum-fr1"]
}
