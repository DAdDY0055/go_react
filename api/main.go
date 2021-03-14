package main

import (
	// "net/http"

	"api/model"

	"github.com/DAdDY0055/go_react/api/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err = gorm.Open("mysql", "user:zaquser@tcp(localhost:3360)/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("[main] db open error")
	}
	defer db.Close()

	db.AutoMigrate(&model.Todo{})

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
