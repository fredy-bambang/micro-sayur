package adapter

import (
	"net/http"
	"strings"
	"user-service/config"
	"user-service/internal/adapter/handler/response"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type MiddlewareAdapterInterface interface {
	CheckToken() echo.MiddlewareFunc
}

type middlewareAdapter struct {
	cfg *config.Config
}

func (m *middlewareAdapter) CheckToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			respErr := response.DefaultResponse{}
			redisConn := config.NewRedisClient()
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				log.Errorf("[MiddlewareAdapter-1] CheckToken: %s", "Missing or Invalid Token")
				respErr.Message = "Missing or Invalid Token"
				respErr.Data = nil
				return c.JSON(http.StatusUnauthorized, respErr)
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			// fmt.Println("tokenString====" + tokenString + "\n")
			getSession, err := redisConn.Get(c.Request().Context(), tokenString).Result()
			if err != nil {
				log.Errorf("[MiddlewareAdapter-2] CheckToken: %v", err)
				respErr.Message = "Invalid Token"
				respErr.Data = nil
				return c.JSON(http.StatusUnauthorized, respErr)
			}

			if len(getSession) == 0 {
				log.Errorf("[MiddlewareAdapter-3] Session: %v", getSession)
				respErr.Message = "Session Not Found"
				respErr.Data = nil
				return c.JSON(http.StatusUnauthorized, respErr)
			}

			c.Set("user", getSession)
			return next(c)
		}
	}
}

func NewMiddlewareAdapter(cfg *config.Config) MiddlewareAdapterInterface {
	return &middlewareAdapter{
		cfg: cfg,
	}
}
