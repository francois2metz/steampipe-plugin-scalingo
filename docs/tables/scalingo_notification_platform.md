# Table: scalingo_notification_platform

List notification platforms that can be used on notifier.

## Examples

### List notification platforms

```sql
select
  id,
  name
from
  scalingo_notification_platform;
```

## Get notification platform used by a notifier

```sql
select
  np.name,
  np.display_name,
  n.name
from
  scalingo_notifier n
join
  scalingo_notification_platform np on np.id = n.platform_id
where
  n.app_name='caresteouvert-api';
```
