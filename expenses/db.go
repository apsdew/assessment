package expenses

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	// os.Setenv("DATABASE_URL", "postgres://yodzpcdo:NFIDn_3NeuQ9LKmHW_NP7Q7JIx6OF7ZU@tiny.db.elephantsql.com/yodzpcdo")

	url := os.Getenv("DATABASE_URL")
	fmt.Println("url :", url)

	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	createTb := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`
	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}

}
