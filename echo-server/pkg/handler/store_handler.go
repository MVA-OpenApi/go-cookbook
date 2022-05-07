package handler

import (
	"echo-server/pkg/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func mainStoreRoutes(e *echo.Echo) {
	s := e.Group("/store")

	s.GET("", getAllStores)
	s.POST("", postStore)
	s.GET("/:id", getStoreByID)
	s.DELETE("/:id", deleteStoreByID)
}

func getAllStores(c echo.Context) error {
	return c.JSON(http.StatusOK, model.Stores)
}

func getStoreByID(c echo.Context) error {
	storeId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return c.String(http.StatusBadRequest, StoreBadRequest)
	}

	for _, store := range model.Stores {
		if store.Id == storeId {
			return c.JSON(http.StatusOK, store)
		}
	}

	return c.String(http.StatusNotFound, StoreNotFound)
}

func postStore(c echo.Context) error {
	newStore := new(model.Store)
	err := c.Bind(newStore)
	newStore.Id = uuid.New()
	newStore.Items = []model.Item{}

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	model.Stores = append(model.Stores, *newStore)

	return c.String(http.StatusOK, StoreAdded)
}

func deleteStoreByID(c echo.Context) error {
	storeId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return c.String(http.StatusBadRequest, StoreBadRequest)
	}

	for i, store := range model.Stores {
		if store.Id == storeId {
			model.Stores = append(model.Stores[:i], model.Stores[i+1:]...)
			return c.String(http.StatusOK, StoreDeleted)
		}
	}

	return c.String(http.StatusNotFound, StoreNotFound)
}
