# Table: scalingo_scm_repo_link

Link between your application and an scm integration.

The `scalingo_scm_repo_link` table can be used to query information about links to an scm and you must specify which application in the where or join clause using the `app_name` column.

## Examples

### Get scm repo link of an app

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

### Get scm repo link from all apps

```sql
select
  app.name,
  srl.owner,
  srl.repo
from
  scalingo_app app
left join
  scalingo_scm_repo_link srl
on
  app.name = srl.app_name;
```
