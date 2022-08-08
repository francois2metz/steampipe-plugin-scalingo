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

- **Breaking change**: The *scalingo_container* has been rename to *scalingo_container_type*
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
