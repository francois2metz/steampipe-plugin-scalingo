# Table: scalingo_invoice

List the invoices associated to your account.

## Examples

### List invoices

```sql
select
  id,
  total_price
from
  scalingo_invoice;
```

### Get unpaid invoices

```sql
select
  id,
  total_price
from
  scalingo_invoice
where
  state != 'paid';
```

### Get price in € per years

```sql
select
  year,
  round(sum(total_price) / 100, 2) as total_price,
  round(sum(total_price_with_vat) / 100, 2) as total_price_with_vat
from
  (
    select
      date_part('year', billing_month) as year,
      total_price,
      total_price_with_vat
    from
      scalingo_invoice
  ) as invoices_as_year
group by
  year;
```
