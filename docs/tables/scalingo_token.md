# Table: scalingo_key

List the API tokens associated to your account.

## Examples

### List tokens

```sql
select
  id,
  name,
  last_used_at
from
  scalingo_token;
```

### List tokens unused since 2 months

```sql
select
  id,
  name,
  last_used_at
from
  scalingo_token
where
  last_used_at < now() - interval '60 days';
```
