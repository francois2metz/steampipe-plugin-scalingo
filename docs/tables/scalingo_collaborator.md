# Table: scalingo_collaborator

A collaborator can access a scalingo application.

The `scalingo_collaborator` table can be used to query information about app collaborators, and you must specify which application in the where or join clause using the `app_name` column.

## Examples

### List collaborators of an application

```sql
select
  status,
  email
from
  scalingo_collaborator
where
  app_name='caresteouvert-api';
```

### Get all collaborators from all apps

```sql
select
  distinct(c.email) as email
from
  scalingo_collaborator as c
inner join
  scalingo_app as a
on a.name = c.app_name
```
