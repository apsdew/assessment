package expenses

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (h *handler) GetExpensesHandler(c echo.Context) error {
	id := c.Param("id")
	ex := Expenses{}

	row := h.DB.QueryRow("SELECT * FROM expenses WHERE id = $1", id)
	err := row.Scan(&ex.ID, &ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Tags))

	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "expenses not found"})
	case nil:
		return c.JSON(http.StatusOK, ex)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't find expenses please contact admin:" + err.Error()})
	}
}
