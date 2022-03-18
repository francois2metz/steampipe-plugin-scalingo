# Table: scalingo_user_event

List events associated to your account.

The `scalingo_user_event` table can be used to query information about any events associated to your account.

## Examples

### List all events

```sql
select
  created_at
  type,
  user_username
from
  scalingo_user_event
```
