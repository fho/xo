// Command xo generates code from database schemas and custom queries. Works
// with PostgreSQL, MySQL, Microsoft SQL Server, Oracle Database, and SQLite3.
package main

//go:generate ./gen.sh models
//go:generate go generate ./internal

import (
	"context"
	"fmt"
	"os"

	// drivers
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/sijms/go-ora/v2"

	// templates
	_ "github.com/fho/xo/templates/createdbtpl"
	_ "github.com/fho/xo/templates/dottpl"
	_ "github.com/fho/xo/templates/gotpl"
	_ "github.com/fho/xo/templates/jsontpl"
	_ "github.com/fho/xo/templates/yamltpl"

	"github.com/fho/xo/cmd"
	"github.com/fho/xo/internal"
	"github.com/fho/xo/templates"
)

// version is the app version.
var version = "0.0.0-dev"

func main() {
	ctx := context.WithValue(context.Background(), templates.SymbolsKey, internal.Symbols)
	if err := cmd.Run(ctx, "xo", version); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
