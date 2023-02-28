# Table: scalingo_database_backup

List of backup of a database associated to an app.

The `scalingo_database_backup` table can be used to query information about a backup of addon database, and you must specify which application and addon in the where or join clause using the `app_name` and the `addon_id` columns.

## Examples

### Get Backup info

```sql
select
  created_at,
  size,
  status
from
  scalingo_database_backup
where
  app_name='caresteouvert-api' and addon_id='ad-0c33a92f-000-000-000-0000000';
```

### List backup from all addons of an application

```sql
with apps_and_addons as (
  select
    ad.id as addon_id,
    ad.app_name as app_name
  from
    scalingo_app app
  join
    scalingo_addon ad
  on
    ad.app_name = app.name
  order by
    app.id
)
select
  created_at,
  size,
  status
from
  scalingo_database_backup bk
join
  apps_and_addons db
on
  db.addon_id = bk.addon_id and db.app_name = bk.app_name;
```
