package middlewares

import (
	"fmt"
	"log"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func JwtAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accessToken := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
			if accessToken == "" {
				log.Println("Authorization header is missing.")
				return echo.NewHTTPError(echo.ErrUnauthorized.Code, map[string]interface{}{
					"status":      "Unauthorized",
					"status_code": echo.ErrUnauthorized.Code,
					"message":     "Authorization header is missing.",
					"result":      nil,
				})
			}

			// secretKey := os.Getenv("JWT_SECRET_KEY")
			secretKey := "d_secret"

			if secretKey == "" {
				log.Println("JWT_SECRET_KEY environment variable is not set.")
				return echo.NewHTTPError(echo.ErrInternalServerError.Code, map[string]interface{}{
					"status":      "Internal Server Error",
					"status_code": echo.ErrInternalServerError.Code,
					"message":     "Server configuration error.",
					"result":      nil,
				})
			}

			token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(secretKey), nil
			})
			if err != nil {
				log.Println("Error parsing token:", err)
				return echo.NewHTTPError(echo.ErrUnauthorized.Code, map[string]interface{}{
					"status":      "Unauthorized",
					"status_code": echo.ErrUnauthorized.Code,
					"message":     "Invalid token.",
					"result":      nil,
				})
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Set("user_id", claims["user_id"])
				c.Set("email", claims["email"])
				return next(c)
			}
			return echo.NewHTTPError(echo.ErrUnauthorized.Code, map[string]interface{}{
				"status":      "Unauthorized",
				"status_code": echo.ErrUnauthorized.Code,
				"message":     "Invalid token claims.",
				"result":      nil,
			})
		}
	}
}
