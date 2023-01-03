package expenses

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (h *handler) UpdateExpensesHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	e := Expenses{}
	err = c.Bind(&e)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "invalid data"})
	}

	if e.Title == "" || e.Amount == 0 || e.Note == "" || e.Tags == nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "invalid data"})
	}

	result, err := db.Exec(`
	UPDATE expenses
	set title=$1,amount=$2,note=$3,tags=$4
	WHERE id=$5;
	`, e.Title, e.Amount, e.Note, pq.Array(e.Tags), id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	rows, err := result.RowsAffected()

	if rows != 1 {
		return c.JSON(http.StatusInternalServerError, Err{Message: "data not found"})
	}

	e.ID = id
	return c.JSON(http.StatusOK, e)
}
