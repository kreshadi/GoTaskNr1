package driver

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	pgUrl, err := pq.ParseURL("postgres://ainmslxk:6lkM-fzYQwgeQO_ZUHceeyei9ws2F9Tk@kandula.db.elephantsql.com:5432/ainmslxk")
	//in the case that we have error to connect to db
	if err != nil {
		log.Fatal(err)
	}

	//first parameter is the driver name: "postgres", and the second one contains inf about our datasource
		db, err = sql.Open("postgres", pgUrl)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("db:", db)

	err = db.Ping()

	return db
}
