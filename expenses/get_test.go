package expenses

import (
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/nozadewz/assessment/request"
	"github.com/nozadewz/assessment/util"
	"github.com/stretchr/testify/assert"
)

func TestGetExpensesHandler(t *testing.T) {
	//Arrange
	ex := Expenses{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	c, rec := request.GetRequest(http.MethodGet, request.GetUri("expenses"), "")
	c.SetParamNames("id")
	c.SetParamValues("1")

	newsMockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow(ex.ID, ex.Title, ex.Amount, ex.Note, pq.Array(ex.Tags))

	db, mock, err := sqlmock.New()
	mock.ExpectQuery("SELECT (.+) FROM expenses").WillReturnRows(newsMockRows)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	h := handler{db}

	//Act
	err = h.GetExpensesHandler(c)

	actual := Expenses{}
	err = util.ConvertToStruct(rec, &actual)
	if err != nil {
		t.Errorf("Test Failed because: %v", err)
	}

	//Assert
	assert.NoError(t, err)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, ex, actual)
	}
}
