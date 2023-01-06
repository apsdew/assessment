package expenses

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nozadewz/assessment/util"
	"github.com/stretchr/testify/assert"
)

func TestItCreateExpensesHandler(t *testing.T) {

	e := echo.New()
	go func(e *echo.Echo) {
		db, err := sql.Open("postgres", "postgres://yodzpcdo:NFIDn_3NeuQ9LKmHW_NP7Q7JIx6OF7ZU@tiny.db.elephantsql.com/yodzpcdo")
		if err != nil {
			log.Fatal(err)
		}

		h := NewApplication(db)

		e.POST("/expenses", h.CreateExpensesHandler)
		e.Start(fmt.Sprintf(":%d", serverPort))
	}(e)
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", serverPort), 30*time.Second)
		if err != nil {
			log.Println(err)
		}
		if conn != nil {
			conn.Close()
			break
		}
	}

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

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/expenses", serverPort), strings.NewReader(util.ConvertToString(newsMockRows)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "November 10, 2009")
	req.SetBasicAuth("expenses", "9999")
	client := http.Client{}

	//Act
	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	resp.Body.Close()

	var actual Expenses
	err = json.Unmarshal(byteBody, &actual)
	if err != nil {
		t.Errorf("Test Failed because: %v", err)
	}

	//Assert
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = e.Shutdown(ctx)
	assert.NoError(t, err)
}
