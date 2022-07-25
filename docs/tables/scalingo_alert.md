# Table: scalingo_alert

List alerts associated to an app.

The `scalingo_alert` table can be used to query information about alert, and you must specify which application in the where or join clause using the `app_name` column.

## Examples

### List alerts of an application

```sql
select
  id,
  disabled,
  metric
from
  scalingo_alert
where
  app_name='caresteouvert-api';
```

### List alerts which do not send notifications when triggered

```sql
select
  id,
  disabled,
  container_type,
  metric
from
  scalingo_alert
where
  app_name='caresteouvert-api' and
  not send_when_below;
```

### List alerts which are disabled

```sql
select
  id,
  container_type,
  metric
from
  scalingo_alert
where
  app_name='caresteouvert-api' and
  disabled;
```
