package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/o-richard/fit/pkg/components"
)

func StartServer(port string) {
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		var he *echo.HTTPError
		if errors.As(err, &he) {
			code = he.Code
		}

		switch code {
		case http.StatusNotFound:
			_ = components.TemplRender(c, code, components.Error404())
		case http.StatusInternalServerError:
			_ = components.TemplRender(c, code, components.Error500())
		default:
			_ = components.TemplRender(c, code, components.ErrorCustom(http.StatusText(code), code))
		}
	}

	e.Static("/", "assets/static")
	e.Static("/media", "assets/images")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Index Page")
	})

	go func() {
		if err := e.Start(fmt.Sprintf(":%v", port)); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				fmt.Println(" shuttting down server ... ")
			} else {
				shutdownCh <- syscall.SIGINT
				fmt.Println(" error while starting server, ", err)
			}
		}
	}()
	<-shutdownCh
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		fmt.Println(" error while shutting down server, ", err)
	}
}
