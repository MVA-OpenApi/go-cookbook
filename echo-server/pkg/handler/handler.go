package handler

import "github.com/labstack/echo/v4"

const (
	StoreBadRequest = "Invalid store id."
	StoreNotFound   = "No store found with given id."
	StoreDeleted    = "Store deleted successfully."
	StoreAdded      = "Store added successfully."

	ItemBadRequest = "Invalid item id."
	ItemNotFound   = "No item found with given id."
	ItemDeleted    = "Item deleted successfully."
	ItemAdded      = "Item added successfully."
	ItemUpdated    = "Item updated successfully."
)

func HandleRoutes(e *echo.Echo) {
	mainStoreRoutes(e)
	mainItemRoutes(e)
}
