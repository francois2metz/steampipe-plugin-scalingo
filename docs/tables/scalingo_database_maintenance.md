# Table: scalingo_database_maintenance

List maintenance of databases.

You **must** specify `app_name` and `addon_id` in a where clause in order to use this table.

## Examples

### List maintenance of a database

```sql
select
  type,
  status,
  started_at
from
  scalingo_database_maintenance
where
  app_name='caresteouvert-api'
  and addon_id='ad-0c33a92f-000-000-000-0000000';
```

### List scheduled maintenance of a database

```sql
select
  type,
  status,
  started_at
from
  scalingo_database_maintenance
where
  app_name='caresteouvert-api'
  and addon_id='ad-0c33a92f-000-000-000-0000000'
  and status='scheduled';
```

### Get one maintenance

```sql
select
  type,
  status,
  started_at
from
  scalingo_database_maintenance
where
  app_name='caresteouvert-api'
  and addon_id='ad-0c33a92f-000-000-000-0000000'
  and id='60a4ab3d406c12000eed29e7';
```
