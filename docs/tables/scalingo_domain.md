# Table: scalingo_domain

List domains associated to an app.

The `scalingo_domain` table can be used to query information about domains, and you must specify which application in the where or join clause using the `app_name` column.

## Examples

### List domains of an application

```sql
select
  name,
  ssl,
  canonical
from
  scalingo_domain
where
  app_name='caresteouvert-api';
```
