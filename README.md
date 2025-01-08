# FDatabase

[![Go Report Card](https://goreportcard.com/badge/github.com/fzf-labs/fdatabase)](https://goreportcard.com/report/github.com/fzf-labs/fdatabase)
[![GoDoc](https://godoc.org/github.com/fzf-labs/fdatabase?status.svg)](https://godoc.org/github.com/fzf-labs/fdatabase)
[![License](https://img.shields.io/github/license/fzf-labs/fdatabase.svg)](https://github.com/fzf-labs/fdatabase/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/fzf-labs/fdatabase)](https://github.com/fzf-labs/fdatabase/blob/main/go.mod)

FDatabase 是一个 Go 语言数据库工具集合，提供了对 MySQL、PostgreSQL 等数据库的便捷操作，以及代码生成等功能。

## ✨ 功能特性

- 🔌 支持 MySQL 和 PostgreSQL 数据库
- 🛠 基于 GORM 的数据库操作封装
- 🔍 数据库连接健康检查
- 📊 数据库状态监控
- 📁 数据库结构导出工具
- ⚡️ 代码生成器(支持生成 DAO、Model、Repository、Proto 文件)
- 📦 缓存支持(Redis、RocksDB)

## 📦 安装

```bash
go get github.com/fzf-labs/fdatabase
```

## 🚀 快速开始

### MySQL 示例

```go
package main

import (
    "github.com/fzf-labs/fdatabase/orm"
)

func main() {
    db, err := orm.NewGormMySQLClient(&orm.GormMySQLClientConfig{
        DataSourceName:  "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local",
        MaxIdleConn:    10,
        MaxOpenConn:    100,
        ShowLog:        true,
    })
    if err != nil {
        panic(err)
    }
    // 使用 db 进行数据库操作...
}
```

### PostgreSQL 示例

```go
package main

import (
    "github.com/fzf-labs/fdatabase/orm"
)

func main() {
    db, err := orm.NewGormPostgresClient(&orm.GormPostgresClientConfig{
        DataSourceName:  "host=localhost port=5432 user=postgres password=123456 dbname=test sslmode=disable",
        MaxIdleConn:    10,
        MaxOpenConn:    100,
        ShowLog:        true,
    })
    if err != nil {
        panic(err)
    }
    // 使用 db 进行数据库操作...
}
```

## 📚 使用文档

### 代码生成

```go
// 生成 DAO、Model、Repository 代码
db.NewGenerationDB(db,
    "./example/postgres/",
    WithOutRepo(),
    WithDBNameOpts(DBNameOpts()),
    WithTables([]string{"users", "orders"}),
).Do()

// 生成 Proto 文件
db.NewGenerationPB(db,
    "./api/proto/",
    "myapp.v1",
    "myapp/api/proto/v1;v1",
).Do()
```

### 数据库导出

```go
// 导出数据库结构
make sqldump -d "user:password@tcp(localhost:3306)/dbname" -o "./doc/sql"
```

## 🤝 贡献

欢迎提交 issue 和 Pull Request。

## 📄 开源协议

本项目采用 [MIT 许可证](LICENSE)。