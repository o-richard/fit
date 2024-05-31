package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/o-richard/fit/pkg/components"
	"github.com/o-richard/fit/pkg/db"
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
		var message string
		var he *echo.HTTPError
		if errors.As(err, &he) {
			code = he.Code
			message, _ = he.Message.(string)
		}

		switch code {
		case http.StatusNotFound:
			_ = components.TemplRender(c, code, components.Error404())
		case http.StatusInternalServerError:
			_ = components.TemplRender(c, code, components.Error500(message))
		default:
			_ = components.TemplRender(c, code, components.ErrorCustom(http.StatusText(code), code))
		}
	}

	e.Static("/", "assets/static")
	e.Static("/media", "assets/images")

	e.GET("/", func(c echo.Context) error {
		return components.TemplRender(c, http.StatusOK, components.InsertHealthEntry())
	})
	e.POST("/entry/new", insertHealthEntry)

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

func insertHealthEntry(c echo.Context) error {
	entry := healthEntry{
		title: c.FormValue("title"), content: c.FormValue("content"), entryType: c.FormValue("entryType"),
		timezone: c.FormValue("timezone"), startedAt: c.FormValue("startedAt"), endedAt: c.FormValue("endedAt"),
	}
	if !entry.validate() {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}
	images := form.File["images"]
	imageNames, isAppError, err := parseImageMediaFiles(images)
	if err != nil && isAppError {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	if err != nil && !isAppError {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	dbHealthEntry := db.HealthEntry{
		Type: entry.entryType, Title: entry.title, Content: entry.content,
		Images: strings.Join(imageNames, ","), StartedAt: entry.startedAtTime, EndedAt: entry.endedAtTime,
	}
	appdb, _ := db.NewDB()
	if err := appdb.InsertHealthEntries(true, []db.HealthEntry{dbHealthEntry}); err != nil {
		for i := range imageNames {
			_ = os.Remove(imageNames[i])
		}
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: fmt.Sprintf("the entry could not be inserted to the databse, %v", err)}
	}

	return c.String(http.StatusOK, "Success")
}
