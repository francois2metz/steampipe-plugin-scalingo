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
