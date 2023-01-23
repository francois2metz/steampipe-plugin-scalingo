# Table: scalingo_scm_integration

SCM Integrations represent a link between your account and an SCM platform like github.com.

## Examples

### List all scm integrations

```sql
select
  scm_type
  username
from
  scalingo_scm_integration;
```

### List scm integrations from GitHub

```sql
select
  *
from
  scalingo_scm_integration
where
  scm_type='github';
```
