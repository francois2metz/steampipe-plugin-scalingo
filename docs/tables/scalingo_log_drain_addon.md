# Table: scalingo_log_drain_addon

A log drain is a remote URL where scalingo push logs from an addon.

The `scalingo_log_drain_addon` table can be used to query information about log drains, and you must specify which application in the where or join clause using the `app_name` and the `id` column.

## Examples

### List log drains of an addon

```sql
select
  url
from
  scalingo_log_drain_addon
where
  app_name='caresteouvert-api'
  and id='60a4ab3d406c12000eed29e7';
```
