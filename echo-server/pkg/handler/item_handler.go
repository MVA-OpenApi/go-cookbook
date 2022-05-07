package handler

import (
	"echo-server/pkg/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func mainItemRoutes(e *echo.Echo) {
	i := e.Group("/store/:id/item")

	i.GET("", getAllItems)
	i.GET("/:item_id", getItemByID)
	i.POST("", postItem)
	i.PUT("/:item_id", putItemByID)
	i.DELETE("/:item_id", deleteItemByID)
}

func getStore(c echo.Context) (*model.Store, error) {
	storeId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return nil, c.String(http.StatusBadRequest, StoreBadRequest)
	}

	store, err := model.GetStoreByID(storeId)
	if err != nil {
		return nil, c.String(http.StatusNotFound, StoreNotFound)
	}

	return store, nil
}

func getAllItems(c echo.Context) error {
	store, err := getStore(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, store.Items)
}

func getItemByID(c echo.Context) error {
	store, err := getStore(c)
	if err != nil {
		return err
	}

	itemId, err := uuid.Parse(c.Param("item_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, ItemBadRequest)
	}

	item, err := model.GetItemByStoreAndID(store, itemId)
	if err != nil {
		return c.String(http.StatusNotFound, ItemNotFound)
	}

	return c.JSON(http.StatusOK, item)
}

func postItem(c echo.Context) error {
	store, err := getStore(c)
	if err != nil {
		return err
	}

	newItem := new(model.Item)
	bErr := c.Bind(newItem)
	newItem.Id = uuid.New()

	if bErr != nil {
		return c.String(http.StatusBadRequest, bErr.Error())
	}

	store.Items = append(store.Items, *newItem)

	return c.String(http.StatusOK, ItemAdded)
}

func putItemByID(c echo.Context) error {
	store, err := getStore(c)
	if err != nil {
		return err
	}

	itemId, err := uuid.Parse(c.Param("item_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, ItemBadRequest)
	}

	for i, item := range store.Items {
		if item.Id == itemId {
			newItem := new(model.Item)
			bErr := c.Bind(newItem)
			if bErr != nil {
				return c.String(http.StatusBadRequest, bErr.Error())
			}

			if newItem.Name != "" {
				store.Items[i].Name = newItem.Name
			}

			if newItem.Price > 0 {
				store.Items[i].Price = newItem.Price
			}

			return c.String(http.StatusOK, ItemUpdated)
		}
	}

	return c.String(http.StatusNotFound, ItemNotFound)
}

func deleteItemByID(c echo.Context) error {
	store, err := getStore(c)
	if err != nil {
		return err
	}

	itemId, err := uuid.Parse(c.Param("item_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, ItemBadRequest)
	}

	for i, item := range store.Items {
		if item.Id == itemId {
			store.Items = append(store.Items[:i], store.Items[i+1:]...)
			return c.String(http.StatusOK, ItemDeleted)
		}
	}

	return c.String(http.StatusNotFound, ItemNotFound)
}
