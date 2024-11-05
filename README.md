# FDatabase

FDatabase 是一个 Go 语言数据库工具集合，提供了对 MySQL、PostgreSQL 等数据库的便捷操作，以及代码生成等功能。

## 功能特性

- 支持 MySQL 和 PostgreSQL 数据库
- 基于 GORM 的数据库操作封装
- 数据库连接健康检查
- 数据库状态监控
- 数据库结构导出工具
- 代码生成器(支持生成 DAO、Model、Repository、Proto 文件)
- 缓存支持(Redis、RocksDB)

## 安装

要安装 FDatabase 库，你可以使用以下命令：

```bash
go get github.com/fzf-labs/fdatabase
```