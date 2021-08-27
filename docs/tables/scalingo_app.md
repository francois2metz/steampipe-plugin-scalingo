# Table: scalingo_app

An application is the base object of the scalingo API.

The `scalingo_app` table can be used to query information about any applications.

## Examples

### List applications

```sql
select
  name,
  region,
  url
from
  scalingo_app;
```

### Get stopped applications

```sql
select
    name
from
  scalingo_app
where
    status = 'stopped';
```
