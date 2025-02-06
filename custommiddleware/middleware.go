package custommiddleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/auth"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
		}

		// Check if the header follows "Bearer <token>" format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "invalid token format"})
		}

		tokenStr := parts[1]
		token, err := jwt.ParseWithClaims(tokenStr, &auth.JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
		}

		if claims, ok := token.Claims.(*auth.JwtCustomClaims); ok && token.Valid {
			c.Set("userId", claims.UserId)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}
}
