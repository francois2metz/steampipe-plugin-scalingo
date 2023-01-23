# Table: scalingo_database

List databases associated to an app.

The `scalingo_database` table can be used to query information about a addon database, and you must specify which application in the where or join clause using the `app_name` column and the `addon_id`.

## Examples

### Get Database info

```sql
select
  *
from
  scalingo_database
where
  app_name='caresteouvert-api' and addon_id='ad-0c33a92f-000-000-000-0000000';
```

### List database from all addons of an application

```sql
select
  db.app_name,
  db.status,
  db.encryption_at_rest
from
  scalingo_addon ad
inner join
  scalingo_database db
on
  ad.id = db.addon_id and ad.app_name = db.app_name
where
  ad.app_name='caresteouvert-api';
```
