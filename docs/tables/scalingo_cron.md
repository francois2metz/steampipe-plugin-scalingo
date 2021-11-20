# Table: scalingo_cron

A cron task is a command executed at a scheduled interval.

The `scalingo_cron` table can be used to query information about cron tasks, and you must specify which application in the where or join clause using the `app_name` column.

## Examples

### List cron tasks of an application

```sql
select
  command
from
  scalingo_cron
where
  app_name='caresteouvert-api';
```
