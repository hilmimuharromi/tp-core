package middleware

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func makeLogEntry(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}

func MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		makeLogEntry(c).Info("incoming request")
		return next(c)
	}
}

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	makeLogEntry(c).Error(report.Message)
	res := map[string]interface{}{
		"code":    report.Code,
		"status":  false,
		"message": report.Message.(string),
	}
	c.JSON(report.Code, res)
}