# Table: scalingo_autoscaler

List autoscaler associated of an app.

The `scalingo_autoscaler` table can be used to query information about autoscaler, and you must specify which application in the where or join clause using the `app_name` column.

## Examples

### List autoscaler of an application

```sql
select
  id,
  metric,
  target
from
  scalingo_autoscaler
where
  app_name='caresteouvert-api';
```

### List disabled autoscaler

```sql
select
  id,
  metric,
  target
from
  scalingo_autoscaler
where
  app_name='caresteouvert-api' and
  disabled;
```

### List autoscaler for a specific container type

```sql
select
  id,
  metric,
  target
from
  scalingo_autoscaler
where
  app_name='caresteouvert-api' and
  container_type='web';
```
