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
