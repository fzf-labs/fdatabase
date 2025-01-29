# 数据库备份

```sh
fdatabase sqldump -d postgres -s "host=localhost user=postgres password=123456 dbname=fdatabase port=5432 sslmode=disable TimeZone=Asia/Shanghai" -o "./orm/example/sql"
```

# 数据库恢复

```sh
fdatabase sqltopb -d postgres -s "host=localhost user=postgres password=123456 dbname=fdatabase port=5432 sslmode=disable TimeZone=Asia/Shanghai" -o "./orm/example/pb"
```

# 数据库代码生成器

```sh
fdatabase ormgen -d postgres -s "host=localhost user=postgres password=123456 dbname=fdatabase port=5432 sslmode=disable TimeZone=Asia/Shanghai" -o "./orm/example/gorm"
```
