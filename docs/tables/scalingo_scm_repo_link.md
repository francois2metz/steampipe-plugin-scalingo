# Table: scalingo_scm_repo_link

Link your Scalingo application to an existing integration.

The `scalingo_scm_repo_link` table can be used to query information about links to an scm and you must specify which application in the where or join clause using the `app_name` column.

## Examples

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
