# Table: scalingo_event_type

Event types represent the definition of the events which are published by the service. They are used by Notification Platforms and by Notifiers to configure which types of events will be sent to the configured destinations.

## Examples

### List type of events

```sql
select
  id,
  name,
  display_name
from
  scalingo_event_type;
```
