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

func TestCreateExpensesHandler(t *testing.T) {
	//Arrange
	newsMockRows := Expenses{
		ID:     0,
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	expected := Expenses{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	ctx, res := request.GetRequest(http.MethodPost, request.GetUri("expenses"), util.ConvertToString(newsMockRows))
	db, mock, err := sqlmock.New()
	mock.ExpectQuery("INSERT INTO expenses (.+) RETURNING id").
		WithArgs(newsMockRows.Title, newsMockRows.Amount, newsMockRows.Note, `{"`+strings.Join(newsMockRows.Tags, `","`)+`"}`).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	h := handler{db}

	//Act
	err = h.CreateExpensesHandler(ctx)
	if err != nil {
		t.Errorf("Test failed: %v", err)
	}

	actual := Expenses{}
	err = util.ConvertToStruct(res, &actual)
	if err != nil {
		t.Errorf("Test Failed because: %v", err)
	}

	//Assert
	assert.Equal(t, http.StatusCreated, res.Code)
	assert.Equal(t, expected, actual)
}
