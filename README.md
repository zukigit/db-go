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
    MAXCONS := 2 // Max Connections for connection pool. Set zero if you dun wanna use it

    // Init mysql database. Don't need to call again.
    dbutil.Init_mysql(
        DBHOST, DBUSER, DBPASSWORD, DBNAME, DBPORT, MAXCONS)
    
    // Take database connection
    db, err := dbutil.GetCon()
    if err != nil {
        log.Fatal(err)
    }

    // Will release a taken connection.
    defer db.ReleaseCon()
}
```

## Performing a Transaction

```go
package main

import (
    "fmt"
    dbutil "github.com/zukigit/db-go/dbutil"
)

func doTest() {
    // Take database connection
    db, err := dbutil.GetCon()
    if err != nil {
        log.Fatal(err)
    }

    // Transaction begins
    if err = db.Begin(); err != nil {
        log.Fatal(err)
    }

    // Your transactional operations here

    // Commit Transaction
    if err = db.Commit(); err != nil {
        log.Fatal(err)
    }
}
```

## Example Program

```go
package main

import (
	"fmt"
	"time"

	"github.com/zukigit/db-go/dbutil"
)

func doTest() {
    var err error

    db, err := dbutil.GetCon()
    if err != nil {
        fmt.Printf("Error in connecting Database. Err: %s\n", err.Error())
        return
    }

    if err = db.Begin(); err != nil {
        fmt.Printf("Query get failed, error: %s\n", err.Error())
        return
    }

    effected_rows, err := db.Execute("insert into hosts (hostid, description) values(%d, '%s');",
        7075, "")
    if err != nil {
        fmt.Printf("Query get failed, error: %s\n", err.Error())
        return
    } else {
        fmt.Println("effected_rows:", effected_rows)
    }

    // if err = db.Commit(); err != nil {
    // 	fmt.Printf("Query get failed, error: %s\n", err.Error())
    // 	return
    // }

    if err = db.Rollback(); err != nil {
        fmt.Printf("Query get failed, error: %s\n", err.Error())
        return
    }

    rows, err := db.Select("select hostid from hosts where hostid = %d", 7073)
    if err != nil {
        fmt.Printf("Error in select. Err: %s\n", err.Error())
        return
    }
    for _, row := range rows {
        fmt.Println(row)
    }

    fmt.Println("task succeeded!!!!")
    db.ReleaseCon()
}

func main() {

    DBHOST := "" // Default: localhost
    DBUSER := "database_user"
    DBPASSWORD := "database_password"
    DBNAME := "database_name"
    DBPORT := 3306 // Default for mysql 3306
    MAXCONS := 2 // Max Connections for connection pool. Set zero if you dun wanna use it

    // Init mysql database. Don't need to call again.
    dbutil.Init_mysql(
        DBHOST, DBUSER, DBPASSWORD, DBNAME, DBPORT, MAXCONS)

    go doTest()
    time.Sleep(1 * time.Second)
    go doTest()
    time.Sleep(1 * time.Second)
    go doTest()
    time.Sleep(1 * time.Second)

    dbutil.Close()
}
```

## License

This project is licensed under the MPL-2.0 License. See the [LICENSE](https://github.com/zukigit/db-go?tab=MPL-2.0-1-ov-file) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue on GitHub.

## Issues

For any issues, please open an issue on GitHub: [db-go Issues](https://github.com/zukigit/db-go/issues)