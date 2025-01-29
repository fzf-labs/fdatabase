# FDatabase

[![Go Report Card](https://goreportcard.com/badge/github.com/fzf-labs/fdatabase)](https://goreportcard.com/report/github.com/fzf-labs/fdatabase)
[![GoDoc](https://godoc.org/github.com/fzf-labs/fdatabase?status.svg)](https://godoc.org/github.com/fzf-labs/fdatabase)
[![License](https://img.shields.io/github/license/fzf-labs/fdatabase.svg)](https://github.com/fzf-labs/fdatabase/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/fzf-labs/fdatabase)](https://github.com/fzf-labs/fdatabase/blob/main/go.mod)

FDatabase 是一个 Go 语言数据库工具集合，提供了对 MySQL、PostgreSQL 等数据库的便捷操作，以及代码生成等功能。

## ✨ 功能特性

- 🛠 基于 GORM 的数据库操作封装
- 🔌 支持 MySQL 和 PostgreSQL 数据库
- 📦 查询缓存支持
- 📁 代码生成器(支持生成 dao、model、repo 文件)
- 📁 数据表结构导出工具
- 📁 数据表生成proto文件工具
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

## 📚 工具使用文档
### 工具安装
```
go install github.com/fzf-labs/fdatabase/orm/cmd/fdatabase
```
### 代码生成

```go
fdatabse ormgen -d "user:password@tcp(localhost:3306)/dbname" -o "./doc/gorm"
```

### 数据库导出

```go
// 导出数据库结构
fdatabse sqldump -d "user:password@tcp(localhost:3306)/dbname" -o "./doc/sql"
```

### 数据库表生成proto文件

```go
fdatabse sqltopb -d "user:password@tcp(localhost:3306)/dbname" -o "./doc/pb"
```

## 🤝 贡献

欢迎提交 issue 和 Pull Request。

## 📄 开源协议

本项目采用 [MIT 许可证](LICENSE)。