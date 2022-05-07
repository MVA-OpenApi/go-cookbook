package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Item struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float32   `json:"price"`
}

type Store struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Items []Item    `json:"items"`
}

var stores = []Store{}

func main() {
	e := echo.New()

	stores = append(stores, Store{Name: "Amazon", Id: uuid.New()})
	stores = append(stores, Store{Name: "Saturn", Id: uuid.New()})

	e.GET("/store", func(c echo.Context) error {
		return c.JSON(http.StatusOK, stores)
	})

	e.GET("/store/:id", func(c echo.Context) error {
		storeId, err := uuid.Parse(c.Param("id"))

		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid store id.")
		}

		for _, store := range stores {
			if store.Id == storeId {
				return c.JSON(http.StatusOK, store)
			}
		}

		return c.String(http.StatusNotFound, "No store found with given id.")
	})

	e.POST("/store", func(c echo.Context) error {
		newStore := new(Store)
		err := c.Bind(newStore)
		newStore.Id = uuid.New()
		newStore.Items = []Item{}

		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		stores = append(stores, *newStore)

		return c.String(http.StatusOK, "Store added successfully.")
	})

	e.DELETE("/store/:id", func(c echo.Context) error {
		storeId, err := uuid.Parse(c.Param("id"))

		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid store id.")
		}

		for i, store := range stores {
			if store.Id == storeId {
				stores = append(stores[:i], stores[i+1:]...)
				return c.String(http.StatusOK, "Store deleted successfully.")
			}
		}

		return c.String(http.StatusNotFound, "No store found with given id.")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
