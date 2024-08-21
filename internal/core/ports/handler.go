package ports

import "github.com/labstack/echo/v4"

type TimesheetHandler interface {
	HandleCreateRequest(c echo.Context) error
}
