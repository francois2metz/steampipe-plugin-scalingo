# Table: scalingo_database_user

List of users of a database

You **must** specify `app_name` and `addon_id` in a where clause in order to use this table.

## Examples

### List users of a database

```sql
select
  name,
  read_only
from
  scalingo_database_user
where
  app_name='caresteouvert-api'
  and addon_id='ad-0c33a92f-000-000-000-0000000';
```

### List readonly users

```sql
select
  name,
  read_only
from
  scalingo_database_user
where
  app_name='caresteouvert-api'
  and addon_id='ad-0c33a92f-000-000-000-0000000'
  and read_only;
```
