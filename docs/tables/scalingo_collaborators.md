# Table: scalingo_collaborators

A collaborator can access a scalingo application.

The `scalingo_collaborators` table can be used to query information about app collaborators, and you must specify which application in the where or join clause using the `app_name` column.

## Examples

### List collaborators of an application

```sql
select
  status,
  email
from
  scalingo_collaborators
where
  app_name='caresteouvert-api';
```

### Get all collaborators from all apps

```sql
select
  distinct(c.email) as email
from
  scalingo_collaborators as c
inner join
  scalingo_apps as a
on a.name = c.app_name
```
