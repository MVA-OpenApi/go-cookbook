package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Item struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type Store struct {
	Name  string `json:"name"`
	Items []Item `json:"items"`
}

var stores = []Store{}

func main() {
	e := echo.New()
	stores = append(stores, Store{Name: "Amazon"})
	stores = append(stores, Store{Name: "Saturn"})

	e.GET("/store", func(c echo.Context) error {
		return c.JSON(http.StatusOK, stores)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
