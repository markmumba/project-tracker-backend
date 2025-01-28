package custommiddleware

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/auth"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		cookie, err := c.Cookie("token")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &auth.JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
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
