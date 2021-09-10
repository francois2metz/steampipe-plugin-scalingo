# Table: scalingo_app_event

List events associated to an app.

The `scalingo_app_event` table can be used to query information about an app event, and you must specify which application in the where or join clause using the `app_name` column.

By default the plugin request the last 1000 events. If you want more, you can use a `limit` parameter in your sql query to increase it (or decrease if needed).

## Examples

### List all events of an application

```sql
select
  created_at
  type,
  user_username
from
  scalingo_app_event
where
  app_name='caresteouvert-api';
```

### Filter by type edit_variables

```sql
select
  created_at,
  type,
  user_username
from
  scalingo_app_event
where
  app_name='caresteouvert-api'
  and type='edit_variables';
```
