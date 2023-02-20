# Table: scalingo_user_scm_repo_link

List all the SCM integrations of your account.

## Examples

### List all scm repo links associated to your account

```sql
select
  scm_type
  owner,
  repo,
  branch
from
  scalingo_scm_repo_link;
```

### Get all scm integration from your account with auto deploy enabled

```sql
select
  scm_type,
  owner,
  repo,
  branch
from
  scalingo_scm_repo_link
where
  auto_deploy_enabled=true;
```
