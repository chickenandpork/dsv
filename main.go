// A quick utility to validate a backing datastore and to allow independent development to extend to datastores we don't currently support
package main

import (
	"flag"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	dia := flag.String("dialect", "mysql", `dialect for connection to test ("sqlite3", "mysql", etc)`)

	// connect strings vary form DB to DB:
	// dialect:
	//   - aws-rds:  fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true", dbUser, authToken, dbEndpoint, dbName)
	//   - postgres:  "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	//   - mariadb: "username:password@protocol(address)/dbname?param=value"
	//   - mssql:  "sqlserver://username:password@host/instance?param1=value&param2=value"
	//   - mysql: "username:password@protocol(address)/dbname?param=value"
	//   - sqlite: "file:test.db?cache=shared&mode=memory"
	//   - sqlite3: "file:test.db?cache=shared&mode=memory"
	//              ":memory:" (in-ram test db)
	connect := flag.String("connect", "scott:tiger@/thedatabase?charset=utf8&parseTime=True&loc=Local", "connection string")

	flag.Parse()

	// fixups
	dialect := *dia
	switch dialect {
	case "aws-rds":
		dialect = "mysql"
	case "sqlite":
		dialect = "sqlite3"
	case "mariadb":
		dialect = "mysql"
	}

	
		

	db, err := gorm.Open(dialect, *connect)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	defer db.Close()

	// ... I promise there will be more meat here

	fmt.Println("Everything runs correctly")
}
