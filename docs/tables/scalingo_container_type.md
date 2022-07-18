# Table: scalingo_container_type

The `scalingo_container_type` table can be used to query information about container types, and you must specify which application in the where or join clause using the `app_name` column.

## Examples

### List container types of an application

```sql
select
  name,
  amount
from
  scalingo_container_type
where
  app_name='caresteouvert-api';
```
