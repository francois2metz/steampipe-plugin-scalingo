# Table: scalingo_notifier

List notifier associated to an app.

The `scalingo_notifier` table can be used to query information about notifier, and you must specify which application in the where or join clause using the `app_name` column.

## Examples

### List notifiers of an application

```sql
select
  name,
  type
from
  scalingo_notifier
where
  app_name='caresteouvert-api';
```

### List active notifiers of an application

```sql
select
  name,
  type
from
  scalingo_notifier
where
  app_name='caresteouvert-api' and active;
```

### List notifiers of an application posting to slack

```sql
select
  name,
  type
from
  scalingo_notifier
where
  app_name='caresteouvert-api' and type='slack';
```

### List events type associated to notifiers

```sql
select
  n.name,
  e.name
from
  scalingo_notifier as n,
  jsonb_array_elements_text(n.selected_event_ids) as event_id,
  scalingo_event_type as e
where
  e.id = event_id and n.app_name='caresteouvert-api';
```
