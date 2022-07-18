# Table: scalingo_container

The `scalingo_container` table can be used to query information about container, and you must specify which application in the where or join clause using the `app_name` column.

## Examples

### List container of an application

```sql
select
  label,
  type,
  state
from
  scalingo_container
where
  app_name='caresteouvert-api';
```

### Count the number of containers running

```sql
select
  count(*)
from
  scalingo_container
where
  app_name='caresteouvert-api';
```

### Get one-off containers running

```sql
select
  label,
  state
from
  scalingo_container
where
  app_name='caresteouvert-api' and type='one-off';
```
