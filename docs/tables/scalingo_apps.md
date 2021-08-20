# Table: scalingo_apps

An application is the base object of the scalingo API.

The `scalingo_apps` table can be used to query information about any applications.

## Examples

### List applications

```sql
select
  name,
  region,
  url
from
  scalingo_apps;
```

### Get stopped applications

```sql
select
    name
from
  scalingo_apps
where
    status = 'stopped';
```
