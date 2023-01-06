package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/nozadewz/assessment/expenses"
	"github.com/nozadewz/assessment/request"
)

var db *sql.DB

func main() {
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	handler(db, *e)

	serverPort := ":" + os.Getenv("PORT")
	go func() {
		if err := e.Start(serverPort); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	graceFulShutDown(e)
}

func InitDB() (*sql.DB, error) {

	url := os.Getenv("DATABASE_URL")

	var err error
	db, err = sql.Open("postgres", url)
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

func handler(db *sql.DB, e echo.Echo) {

	h := expenses.NewApplication(db)

	// ### If you want to use the Basicauth ###
	// e.Use(middleware.BasicAuth(func(user, pass string, c echo.Context) (bool, error) {
	// 	if user == "expenses" || pass == "9999" {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))

	g := e.Group("/expenses")
	g.Use(middleware.Logger())
	g.Use(middleware.Recover())
	g.Use(request.GetAuth)

	g.POST("", h.CreateExpensesHandler)
	g.GET("/:id", h.GetExpensesHandler)
	g.PUT("/:id", h.UpdateExpensesHandler)
	g.GET("", h.GetAllExpensesHandler)
}

func graceFulShutDown(e *echo.Echo) {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
