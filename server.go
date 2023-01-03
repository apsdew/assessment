package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/nozadewz/assessment/expenses"
)

var db *sql.DB

func main() {
	expenses.InitDB()
	h := expenses.NewApplication(db)

	// os.Setenv("PORT", "2565")
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
