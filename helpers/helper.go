package helpers

import (
    "fmt"

    "github.com/labstack/echo/v4"
)

// ConvertUserID safely converts userID from interface{} to uint
func ConvertUserID(c echo.Context, key string) (uint, error) {
    userIDInterface := c.Get(key)
    switch id := userIDInterface.(type) {
    case int:
        if id >= 0 {
            return uint(id), nil
        }
        return 0, fmt.Errorf("negative user ID not allowed")
    case uint:
        return id, nil
    default:
        return 0, fmt.Errorf("invalid user ID type")
    }
}
