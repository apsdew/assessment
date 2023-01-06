package request

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetRequest(method, url string, body string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, url, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	return c, rec
}

func GetUri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func GetAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if len(c.Request().Header["Authorization"]) > 0 {
			if c.Request().Header["Authorization"][0] == "November 10, 2009" {
				c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
				return next(c)
			}
		}
		return c.JSON(http.StatusUnauthorized, "You are not authorized!")
	}
}
