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
)

func StartServer(port string) {
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

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
