# GORM - Working with Databases

## Installation

```bash
go get gorm.io/gorm

go get gorm.io/driver/<database>
```

## Setting up connection

- Set data source name (dsn). Only the `dbName` is a must-have
    ```golang
    var dsn := "user:pass@tcp(127.0.0.1:3306)/dbName?charset=utf8mb4"
    ```

- To create a connection,
    ```go
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    ```

- To create tables, drop tables and more we use the `Migrator()` interface
    ```go
    db.Migrator().CreateTable(&structure)
    ```

- To insert values and error handle,
    ```go
    if err=db.Create(&Struct{val:"foo"}).Error; err != nil {
        // handle err
    }
    ```
