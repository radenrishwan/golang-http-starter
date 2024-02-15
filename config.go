package starter

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var (
	dbUrl = flag.String("DB_URL", "", "set database url (you can use env variable DB_URL)")
)

func NewDatabase(context context.Context) *pgx.Conn {
	flag.Parse()

	if *dbUrl == "" {
		// get from env
		*dbUrl = os.Getenv("DB_URL")
		if *dbUrl == "" {
			log.Fatalln("You need to set database url using flag or env variable DB_URL")
		}
	}

	conn, err := pgx.Connect(context, *dbUrl)
	if err != nil {
		log.Fatalln("can't connect to database, make sure your database is running and the url is correct. Error:", err)
	}

	return conn
}
