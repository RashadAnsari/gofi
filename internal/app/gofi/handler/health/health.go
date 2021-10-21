package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func (h Handler) Healthz(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
