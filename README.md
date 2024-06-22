# db-go

`db-go` is a Go library designed to simplify database interactions. It provides support for database connection pooling, transactions, and dynamic data selection, making it easier to work with various databases in your Go applications.

## Features

- **Connection Pooling**: Efficiently manage and reuse database connections.
- **Transactions**: Handle database transactions with ease.
- **Dynamic Data Selection**: Simplify dynamic queries and data retrieval.

## Installation

To install `db-go`, use the following command:

```bash
go get github.com/zukigit/db-go
```

## Usage

Importing the Package

```go
import dbutil "github.com/zukigit/db-go/dbutil"
```

## Connecting to a Database

Here's an example of how to establish a connection to a database:

```go
package main

import (
    "fmt"
    dbutil "github.com/zukigit/db-go/dbutil"
)

func main() {

    DBHOST := "" // Default: localhost
    DBUSER := "database_user"
    DBPASSWORD := "database_password"
    DBNAME := "database_name"
    DBPORT := 3306 // Default for mysql 3306
    MAXCON := 2 // Max Con Count.

    // Init mysql database. Don't need to call again.
    dbutil.Init_mysql(
        DBHOST, DBUSER, DBPASSWORD, DBNAME, DBPORT, MAXCONNECTIONS)
    
    // Take database connection
    db, err := dbutil.GetConnection()
    if err != nil {
        log.Fatal(err)
    }

    // Will release taken connection.
    defer db.ReleaseCon()
}
```