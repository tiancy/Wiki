```sql
  select * from table where name = ANY(regexp_split_to_array('tom,mary,jack', ','))
```
