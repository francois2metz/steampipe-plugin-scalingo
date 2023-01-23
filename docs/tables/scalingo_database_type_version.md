# Table: scalingo_database_type_version

Get the version of a database.

The `scalingo_database_type_version` table can be used to query information about a version of a database, and you must specify which application in the where or join clause using the `app_name` column and the `addon_id` as well as the version `id`.

## Examples

### Get Database version

```sql
select
  *
from
  scalingo_database_type_version
where
  app_name='caresteouvert-api'
  and addon_id='ad-0c33a92f-000-000-000-0000000'
  and id='60a4ab3d406c12000eed29e7';
```

### Get all Database versions of an application

```sql
with database as (
  select
    db.app_name,
    ad.id as addon_id,
    db.version_id
  from
    scalingo_addon ad
  inner join
    scalingo_database db
  on
    ad.id = db.addon_id and ad.app_name = db.app_name
  where
    ad.app_name='caresteouvert-api'
  order by
    db.id
)
select
  *
from
  database db
join
  scalingo_database_type_version tv on (tv.id = db.version_id and tv.addon_id = db.addon_id and tv.app_name = db.app_name);
```

### Check if an upgrade exist for your database

```sql
select
  case when next_upgrade_id is null
    then false
    else true
  end as toupdate
from
  scalingo_database_type_version
where
  app_name='caresteouvert-api'
  and addon_id='ad-0c33a92f-000-000-000-0000000'
  and id='60a4ab3d406c12000eed29e7';
```
