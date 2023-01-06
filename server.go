package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/nozadewz/assessment/expenses"
)

var db *sql.DB

func main() {
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}
	h := expenses.NewApplication(db)

	os.Setenv("PORT", "2565")
	serverPort := ":" + os.Getenv("PORT")
	e := echo.New()

	e.Use(middleware.BasicAuth(func(user, pass string, c echo.Context) (bool, error) {
		if user == "expenses" || pass == "9999" {
			return true, nil
		}
		return false, nil
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	g := e.Group("/expenses")
	g.POST("", h.CreateExpensesHandler)
	g.GET("/:id", h.GetExpensesHandler)
	g.PUT("/:id", h.UpdateExpensesHandler)
	g.GET("", h.GetAllExpensesHandler)

	log.Fatal(e.Start(serverPort))
}

func InitDB() (*sql.DB, error) {
	os.Setenv("DATABASE_URL", "postgres://yodzpcdo:NFIDn_3NeuQ9LKmHW_NP7Q7JIx6OF7ZU@tiny.db.elephantsql.com/yodzpcdo")

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

	return db, nil
}
