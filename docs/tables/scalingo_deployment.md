# Table: scalingo_deployment

List deployments associated to an app.

The `scalingo_deployment` table can be used to query information about deployments, and you must specify which application in the where or join clause using the `app_name` column.

## Examples

### List all deployments of an application

```sql
select
  created_at,
  status,
  user_username
from
  scalingo_deployment
where
  app_name='caresteouvert-api';
```

### Get average deploy duration

```sql
select
  mean(duration)
from
  scalingo_deployment
where
  app_name='caresteouvert-api';
```
