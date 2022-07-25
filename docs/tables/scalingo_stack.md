# Table: scalingo_stack

List base docker image where applications can be executed in.

## Examples

### List stacks

```sql
select
  id,
  name,
  created_at
from
  scalingo_stack;
```

### Get default stack

```sql
select
  id,
  name
from
  scalingo_stack
where
  "default";
```
