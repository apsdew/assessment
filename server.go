package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/nozadewz/assessment/expenses"
)

var db *sql.DB

func main() {
	expenses.InitDB()

	// os.Setenv("PORT", "2565")
	// port := os.Getenv("PORT")
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

	g.POST("", expenses.CreateExpensesHandler)
	g.GET("/:id", expenses.GetExpensesHandler)
	// g.PUT("/:id", expenses.UpdateExpensesHandler)
	g.GET("", expenses.GetAllExpensesHandler)

	log.Fatal(e.Start(":2565"))
}
