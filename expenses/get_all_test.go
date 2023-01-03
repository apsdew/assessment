package expenses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllExpensesHandler(t *testing.T) {
	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/expenses", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	testcase := Expenses{
		ID:     0,
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	newsMockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).AddRow(ResString(testcase))

	fmt.Println(newsMockRows)

	db, mock, err := sqlmock.New()
	mock.ExpectQuery("SELECT (.+) FROM expenses ORDER BY id").WillReturnRows(newsMockRows)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	h := handler{db}
	c := e.NewContext(req, rec)
	expected := "[{\"ID\":1,\"Title\":\"strawberry smoothie\",\"Amount\":79,\"Note\":\"night market promotion discount 10 bath\",\"Tags\":[\"food\",\"beverage\"]}]"

	// Act
	err = h.CreateExpensesHandler(c)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
	}
}

func ResString(reqStruct interface{}) string {
	if reqStruct == nil {
		return ""
	}
	result, _ := json.Marshal(&reqStruct)
	return string(result)
}

func ResStruct(res *httptest.ResponseRecorder, result interface{}) error {
	return json.Unmarshal([]byte(res.Body.Bytes()), &result)
}
