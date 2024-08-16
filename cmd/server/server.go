package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thitiphongD/my-echo/infrastructure/database"
)

func main() {

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Failed to get DB from GORM: %v", err)
		}
		sqlDB.Close()
	}()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
