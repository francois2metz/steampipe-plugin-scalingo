# Table: scalingo_container

The `scalingo_container` table can be used to query information about containers, and you must specify which application in the where or join clause using the `app_name` column.

## Examples

### List containers of an application

```sql
select
  name,
  amount
from
  scalingo_container
where
  app_name='caresteouvert-api';
```
