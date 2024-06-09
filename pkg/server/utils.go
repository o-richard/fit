package server

import (
	"fmt"
	"io"
	"math/rand/v2"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/o-richard/fit/pkg/db"
)

const (
	// Expected layout as sent by the browser
	datetimeLayout = "2006-01-02T15:04"
)

type healthEntry struct {
	title         string
	content       string
	entryType     string
	startedAt     string
	endedAt       string
	timezone      string
	startedAtTime time.Time
	endedAtTime   time.Time
}

func (h *healthEntry) validate() bool {
	h.title = strings.TrimSpace(h.title)
	h.content = strings.TrimSpace(h.content)

	if h.content == "" {
		return false
	}
	switch h.entryType {
	case db.Sleep, db.Nutrition, db.Activity, db.Health:
	default:
		return false
	}

	location, err := time.LoadLocation(h.timezone)
	if err != nil {
		return false
	}
	startedAtTime, err := time.ParseInLocation(datetimeLayout, h.startedAt, location)
	if err != nil {
		return false
	}
	endedAtTime, err := time.ParseInLocation(datetimeLayout, h.endedAt, location)
	if err != nil {
		return false
	}
	if endedAtTime.Compare(startedAtTime) < 0 {
		return false
	}

	h.startedAtTime = startedAtTime
	h.endedAtTime = endedAtTime
	return true
}

// returns image file names & whether the error is by the application (true) or by user (false)
func parseImageMediaFiles(images []*multipart.FileHeader) ([]string, bool, error) {
	var maxUploadSize int64 = 2 * 1024 * 1024

	imageNames := make([]string, 0, len(images))
	imageCleanup := func(names []string, isAppError bool, msg string) ([]string, bool, error) {
		for i := range names {
			_ = os.Remove(names[i])
		}
		return nil, isAppError, fmt.Errorf(msg)
	}

	for _, image := range images {
		if image.Size > maxUploadSize {
			return imageCleanup(imageNames, false, "Maximum image size exceeded")
		}
		imageType := image.Header.Get(echo.HeaderContentType)
		if imageType != "image/jpeg" && imageType != "image/png" {
			return imageCleanup(imageNames, false, "Image should be .jpg OR .png")
		}

		src, err := image.Open()
		if err != nil {
			return imageCleanup(imageNames, true, "Unable to open image")
		}
		defer src.Close()

		// the likelihood of name collision is minimal
		newFileName := fmt.Sprintf("assets/media/%d_%v.%v", time.Now().Unix(), rand.IntN(1_000_000), filepath.Ext(image.Filename))
		dst, err := os.Create(newFileName)
		if err != nil {
			return imageCleanup(imageNames, true, "Unable to open image destination")
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return imageCleanup(imageNames, true, "Unable to move image to destination")
		}
		imageNames = append(imageNames, newFileName)
	}
	return imageNames, false, nil
}
