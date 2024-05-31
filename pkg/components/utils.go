package components

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// Replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func TemplRender(ctx echo.Context, statusCode int, t templ.Component) error {
	// other status codes already have their header already set
	if statusCode == http.StatusOK {
		ctx.Response().Writer.WriteHeader(statusCode)
	}
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}
