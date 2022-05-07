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

func getStoreIndex(c echo.Context) (int, error) {
	storeId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return -1, c.String(http.StatusBadRequest, StoreBadRequest)
	}

	storeIndex, err := model.FindStoreIndex(storeId)
	if err != nil {
		return -1, c.String(http.StatusNotFound, StoreNotFound)
	}

	return storeIndex, nil
}

func getAllItems(c echo.Context) error {
	storeIndex, err := getStoreIndex(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.GetStoreByIndex(storeIndex).Items)
}

func getItemByID(c echo.Context) error {
	storeIndex, err := getStoreIndex(c)
	if err != nil {
		return err
	}

	itemId, err := uuid.Parse(c.Param("item_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, ItemBadRequest)
	}

	item, err := model.GetItemByStoreAndID(storeIndex, itemId)
	if err != nil {
		return c.String(http.StatusNotFound, ItemNotFound)
	}

	return c.JSON(http.StatusOK, item)
}

func postItem(c echo.Context) error {
	storeIndex, err := getStoreIndex(c)
	if err != nil {
		return err
	}

	newItem := new(model.Item)
	bErr := c.Bind(newItem)
	newItem.Id = uuid.New()

	if bErr != nil {
		return c.String(http.StatusBadRequest, bErr.Error())
	}

	model.AddItemToStore(storeIndex, *newItem)

	return c.String(http.StatusOK, ItemAdded)
}

func putItemByID(c echo.Context) error {
	storeIndex, err := getStoreIndex(c)
	if err != nil {
		return err
	}

	itemId, err := uuid.Parse(c.Param("item_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, ItemBadRequest)
	}

	for i, item := range model.GetStoreByIndex(storeIndex).Items {
		if item.Id == itemId {
			newItem := new(model.Item)
			bErr := c.Bind(newItem)
			if bErr != nil {
				return c.String(http.StatusBadRequest, bErr.Error())
			}

			model.UpdateItemInStore(storeIndex, i, *newItem)

			return c.String(http.StatusOK, ItemUpdated)
		}
	}

	return c.String(http.StatusNotFound, ItemNotFound)
}

func deleteItemByID(c echo.Context) error {
	storeIndex, err := getStoreIndex(c)
	if err != nil {
		return err
	}

	itemId, err := uuid.Parse(c.Param("item_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, ItemBadRequest)
	}

	for i, item := range model.GetStoreByIndex(storeIndex).Items {
		if item.Id == itemId {
			model.DeleteItemFromStore(storeIndex, i)
			return c.String(http.StatusOK, ItemDeleted)
		}
	}

	return c.String(http.StatusNotFound, ItemNotFound)
}
