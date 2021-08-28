# Table: scalingo_addon

An addons is a database or other things attached to an application.

The `scalingo_addon` table can be used to query information about an addon, and you must specify which application in the where or join clause using the `app_name` column.

## Examples

### List addons of an application

```sql
select
  provider_name,
  status
from
  scalingo_addon
where
  app_name='caresteouvert-api';
```

### Get the price of all addons of a given application

```sql
select
  sum(plan_price)
from
  scalingo_addon
where
  app_name='caresteouvert-api';
```
