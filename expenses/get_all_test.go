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

func TestGetAllExpensesHandler(t *testing.T) {
	//Arrange
	expecteds := []Expenses{
		{
			ID:     1,
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"food", "beverage"},
		},
		{
			ID:     2,
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"food", "beverage"},
		},
	}

	ctx, res := request.GetRequest(http.MethodGet, request.GetUri("expenses"), "")

	newsMockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"})
	for _, i := range expecteds {
		newsMockRows.AddRow(i.ID, i.Title, i.Amount, i.Note, pq.Array(i.Tags))
	}

	db, mock, err := sqlmock.New()
	mock.ExpectQuery("SELECT (.+) FROM expenses ORDER BY id").WillReturnRows(newsMockRows)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	h := handler{db}

	//Act
	err = h.GetAllExpensesHandler(ctx)

	actuals := []Expenses{}
	err = util.ConvertToStruct(res, &actuals)
	if err != nil {
		t.Errorf("Test Failed because: %v", err)
	}

	//Assert
	assert.NoError(t, err)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expecteds, actuals)
	}
}
