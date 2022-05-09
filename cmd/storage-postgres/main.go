package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

type PostgreSQL struct {
	host string
	port int
	db   string
}

var databases []PostgreSQL = []PostgreSQL{
	PostgreSQL{host: "postgres-1", port: 5432, db: "customers"},
	PostgreSQL{host: "postgres-1", port: 5432, db: "bookings"},
	PostgreSQL{host: "postgres-2", port: 5432, db: "flights"},
}

func main() {
	pp := &cli.App{
		Name:        "storage-postgres",
		Description: "manipulate with organization PostgreSQL databases",
		Commands: []*cli.Command{
			&cli.Command{
				Name:   "list",
				Action: list,
			},
			&cli.Command{
				Name:      "get",
				ArgsUsage: "<db>",
				Action:    get,
			},
		},
	}
	pp.Run(os.Args)
}

func list(ctx *cli.Context) error {
	fmt.Printf("organization.com PostgreSQL databases\n")
	for _, db := range databases {
		fmt.Printf("%s\t%s\t%d\n", db.host, db.db, db.port)
	}
	return nil
}

func get(ctx *cli.Context) error {
	dbToFind := ctx.Args().First()
	for _, db := range databases {
		if db.db == dbToFind {
			fmt.Printf("host: %s\n", db.host)
			fmt.Printf("db:   %s\n", db.db)
			fmt.Printf("port: %d\n", db.port)
			return nil
		}
	}

	fmt.Printf("unknown db\n")
	return nil
}
