package expenses

import (
	"net/http"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nozadewz/assessment/request"
	"github.com/nozadewz/assessment/util"
	"github.com/stretchr/testify/assert"
)

func TestUpdateExpensesHandler(t *testing.T) {
	//Arrange
	id := "1"
	expected := Expenses{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 89,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	ctx, rec := request.GetRequest(http.MethodPut, request.GetUri("expenses"), util.ConvertToString(expected))
	ctx.SetParamNames("id")
	ctx.SetParamValues(id)

	db, mock, err := sqlmock.New()

	mock.ExpectExec(("UPDATE expenses set (.+)")).
		WithArgs(expected.Title, expected.Amount, expected.Note, `{"`+strings.Join(expected.Tags, `","`)+`"}`, string(id)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	h := handler{db}

	//Act
	err = h.UpdateExpensesHandler(ctx)
	if err != nil {
		t.Errorf("Test failed: %v", err)
	}

	actual := Expenses{}
	err = util.ConvertToStruct(rec, &actual)
	if err != nil {
		t.Errorf("Test Failed because: %v", err)
	}

	//Assert
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, actual)
	}
}
