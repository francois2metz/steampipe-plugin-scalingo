## v0.18.0 [2023-10-13]

_What's new?_

- Update go-scalingo to v6.7.2
- Update steampipe sdk to 5.6.2
- Update to go 1.21

## v0.17.0 [2023-07-18]

_What's new?_

- Update go-scalingo to v6.7.1
- Update steampipe sdk to 5.5.0
- Add `maintenance_window_*` columns to the database table

## v0.16.1 [2023-04-19]

_What's new?_

- Update go-scalingo to v6.5.0

## v0.16.0 [2023-03-16]

_What's new?_

- Add columns `started_at` and `method` to `scalingo_database_backup` table
- Add a column `stack_base_image` to `scalingo_deployment` table
- Update steampipe sdk to 5.3.0
- Update go-scalingo to v6.4.0

## v0.15.0 [2023-02-28]

_What's new?_

- **Breaking change**: The *scalingo_scm_repo_link* table has been splitted to not mix different behavior of the Scalingo API. The *scalingo_scm_repo_link* table is now only used to get the repo link from a specific application with the where clause `app_name`. The *scalingo_user_scm_repo_link* table now is used to list all repository links associated to your account and there is no where clause to use to make it works.
- Update the *scalingo_database_backup* table doc.

## v0.14.0 [2023-02-23]

_What's new?_

- Add `scalingo_database_backup` table. Thanks @emeric-martineau.

## v0.13.1 [2023-02-20]

_What's new?_

- Add columns `items` and `detailed_items` to `scalingo_invoice` table

## v0.13.0 [2023-02-17]

_What's new?_

- Add a column `automatic_creation_from_forks_allowed` to `scalingo_scm_repo_link` table
- Add a column `image_size` to `scalingo_deployment` table
- Add a column `hds_resource` to `scalingo_app` table
- Update steampipe sdk to 5.1.4
- Update go-scalingo to v6.3.0

## v0.12.1 [2023-01-25]

_What's new?_

- Disable query timeout from the scalingo sdk
- Update steampipe sdk to 5.1.2
- Update go-scalingo to v6.1.0
- Update doc

## v0.12.0 [2022-11-25]

_What's new?_

- Query all regions on `scalingo_stack` table
- Add a column `region` on `scalingo_stack` table
- Update go-scalingo to v6.0.1

## v0.11.0 [2022-11-10]

_What's new?_

- Add `flags`, `limits` and `data_access_consent_*` columns to the `scalingo_app` table
- Update steampipe sdk to 4.1.8

## v0.10.0 [2022-10-18]

_What's new?_

- Add `type_data` to the `scalingo_alert` table
- Fix the description to the `scalingo_autoscaler`

## v0.9.0 [2022-09-29]

_What's new?_

- Update steampipe sdk to 4.1.7
- Update go-scalingo to v6
- Add `scalingo_invoice` table
- Add `created_at`, `updated_at` and `metadata` to the `scalingo_alert` table
- Add `created_at`, `owner_id`, `owner_username` and `owner_email` to the `scalingo_scm_integration` table
- Add `provider_hds_available`, `provider_short_description`, `provider_description`, `plan_on_demand`, `plan_disabled`, `plan_disabled_alternative_plan_id` and `plan_hds_available` to the `scalingo_addon` table
- Remove `plan_logo_url` of the `scalingo_addon` table
- Add `url` to the `scalingo_scm_repo_link` table

## v0.8.2 [2022-09-05]

_What's new?_

- Add `deprecated_at` column to the `scalingo_stack` table

## v0.8.1 [2022-09-05]

_What's new?_

- Update steampipe sdk to 4.1.6

## v0.8.0 [2022-09-01]

_What's new?_

- Update to go 1.19
- Update steampipe sdk to 4.1.5
- Update go-scalingo to v5

## v0.7.1 [2022-08-08]

_What's new?_

- Add an example on `scalingo_database_type_version` database
- Fix some columns description
- Add a column `region` on `scalingo_event_type` and `scalingo_notification_platform` tables

## v0.7.0 [2022-07-26]

_What's new?_

- New tables added

  - scalingo_autoscaler
  - scalingo_notification_platform

- Update steampipe sdk to 3.3.2
- Fix the scalingo_database_type_version name and error
- Add some examples on scalingo_alert
- Comment the scalingo token on the default config

## v0.6.0 [2022-07-24]

_What's new?_

- New tables added

  - scalingo_alert
  - scalingo_notifier
  - scalingo_event_type

## v0.5.0 [2022-07-18]

_What's new?_

- **Breaking change**: The *scalingo_container* table has been renamed to *scalingo_container_type*
- Add *scalingo_container* table

## v0.4.0 [2022-05-31]

_What's new?_

- Add scalingo_stack table

## v0.3.2 [2022-05-30]

_What's new?_

- Update steampipe sdk to 3.2
- Ignore unauthorized errors

## v0.3.1 [2022-05-22]

_What's new?_

- Fix the documentation
- Add error logs on all API calls

## v0.3.0 [2022-05-15]

_What's new?_

- Add scalingo_scm_integration table

## v0.2.0 [2022-05-02]

_What's new?_

- Update steampipe sdk to 3.1.0
- Update to go 1.18
- Build ARM64 binaries

## v0.1.3 [2022-05-02]

_What's new?_

- Add scalingo_log_drain table
- Add scalingo_log_drain_addon table

## v0.1.2 [2022-04-27]

_What's new?_

- Fetch all entries for scalingo_scm_repo_link table

## v0.1.1 [2022-04-02]

_What's new?_

- Add scalingo_user_event table
- Add scalingo_database_type_version table

## v0.1.0 [2022-03-14]

_What's new?_

- Add multi-region support per connection. This add a new `regions` configuration. The old `region` parameter is still supported for now, but the `SCALINGO_REGION` environmnent variable is removed.
- The `endpoint` configuration support has been removed

## v0.0.9 [2022-03-02]

_What's new?_

- Add scalingo_scm_repo_link table
- Update steampipe sdk to 2.0.3
- Remove deprecated configuration about endpoint

## v0.0.8 [2022-01-28]

_What's new?_

- Add scalingo_database table
- Update steampipe sdk to 1.8.3

## v0.0.7 [2021-11-24]

_What's new?_

- Add scalingo_cron table
- Update date column types to use timestamp instead of datetime
- Update steampipe sdk to 1.8.2

## v0.0.6 [2021-11-16]

_What's new?_

- A new region config parameter has been added (set osc-fr1 or osc-secnum-fr1)
- The endpoint config parameter has been deprecated

## v0.0.5 [2021-11-08]

_What's new?_

- Remove the default limit in the scalingo_app_event

## v0.0.4 [2021-10-18]

_What's new?_

- Handle 404 when the app_name cannot be found

## v0.0.3 [2021-09-14]

_What's new?_

- Added the .goreleaser.yml file

## v0.0.2 [2021-09-14]

_What's new?_

- New tables added

  - scalingo_app_event
  - scalingo_collaborator
  - scalingo_container
  - scalingo_deployment
  - scalingo_domain
  - scalingo_environment
  - scalingo_key
  - scalingo_region
  - scalingo_token
