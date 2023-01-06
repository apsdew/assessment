package expenses

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (h *handler) GetAllExpensesHandler(c echo.Context) error {
	rows, err := h.DB.Query("SELECT * FROM expenses ORDER BY id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "Not found your expenses"})
	}

	defer rows.Close()

	res := []Expenses{}
	e := Expenses{}

	for rows.Next() {
		err := rows.Scan(&e.ID, &e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: "can't find expenses please contact admin:" + err.Error()})
		}
		res = append(res, e)
	}

	return c.JSON(http.StatusOK, res)
}
