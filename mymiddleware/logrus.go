package mymiddleware

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func Logrus() echo.MiddlewareFunc {

	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		ForceColors : false,
		DisableColors: true,
		DisableTimestamp : false,
		FullTimestamp : true,
	}

	/* ... logger 초기화 */
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logEntry := logrus.NewEntry(logger)
			logEntry.WithFields(logrus.Fields{
				"method": c.Request().Method,
				"url":    c.Request().URL,
				"ip":     c.RealIP(),
			}).Info("access")

			// logEntry를 Context에 저장. 추후 다른곳에서 로깅이 필요할때 사용
			/*
			req := c.Request()
			c.SetRequest(req.WithContext(
				context.WithValue(
					req.Context(),
					"LOG",
					logEntry,
				),
			))*/

			return next(c)
		}
	}
}
