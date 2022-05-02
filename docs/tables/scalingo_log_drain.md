# Table: scalingo_log_drain

A log drain is a remote URL where scalingo push logs from an application.

The `scalingo_log_drain` table can be used to query information about log drains, and you must specify which application in the where or join clause using the `app_name` column.

## Examples

### List log drains of an application

```sql
select
  url
from
  scalingo_log_drain
where
  app_name='caresteouvert-api';
```
