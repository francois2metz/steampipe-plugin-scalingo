# Table: scalingo_environment

The `scalingo_environment` table can be used to query information about environment variables, and you must specify which application in the where or join clause using the `app_name` column.

## Examples

### List environment variables of an application

```sql
select
  name,
  value
from
  scalingo_environment
where
  app_name='caresteouvert-api';
```

## Find apps with a specific environment variable value

```sql
select
  a.name as app_name,
  e.name as env_name
from
  scalingo_app as a
join
  scalingo_environment as e
on
 a.name = e.app_name
where
  e.value = 'my variable value'
```
