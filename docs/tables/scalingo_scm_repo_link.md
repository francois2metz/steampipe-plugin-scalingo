# Table: scalingo_scm_repo_link

Link your Scalingo application to an existing integration.

The `scalingo_scm_repo_link` table can be used to query information about links to an scm.

## Examples

### List all scm repo links

```sql
select
  scm_type
  owner,
  repo,
  branch
from
  scalingo_scm_repo_link
```

### Get scm integration of an app

```sql
select
  scm_type
  owner,
  repo,
  branch
from
  scalingo_scm_repo_link
where
  app_name='caresteouvert-api';
```

### Get all scm integration with auto deploy enabled

```sql
select
  scm_type,
  owner,
  repo,
  branch
from
  scalingo_scm_repo_link
where
  auto_deploy_enabled=true;
```
