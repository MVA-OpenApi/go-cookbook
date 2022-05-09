package model

import (
	"errors"

	"github.com/google/uuid"
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

var Stores = []Store{}

func InitDB() {
	Stores = append(Stores, Store{Name: "Amazon", Id: uuid.New()})
	Stores[0].Items = append(Stores[0].Items, Item{Name: "A", Price: 10.0, Id: uuid.New()}, Item{Name: "B", Price: 69.0, Id: uuid.New()})
	Stores = append(Stores, Store{Name: "Saturn", Id: uuid.New()})
	Stores[1].Items = append(Stores[1].Items, Item{Name: "C", Price: 42.0, Id: uuid.New()})
}

func FindStoreIndex(id uuid.UUID) (int, error) {
	for i, store := range Stores {
		if store.Id == id {
			return i, nil
		}
	}

	return -1, errors.New("no store found")
}

func GetStoreByIndex(index int) *Store {
	return &Stores[index]
}

func GetItemByStoreAndID(index int, id uuid.UUID) (*Item, error) {
	for _, item := range Stores[index].Items {
		if item.Id == id {
			return &item, nil
		}
	}

	return nil, errors.New("no item found")
}

func AddItemToStore(index int, item Item) {
	Stores[index].Items = append(Stores[index].Items, item)
}

func UpdateItemInStore(storeIndex int, itemIndex int, item Item) {
	if item.Name != "" {
		Stores[storeIndex].Items[itemIndex].Name = item.Name
	}

	if item.Price > 0 {
		Stores[storeIndex].Items[itemIndex].Price = item.Price
	}
}

func DeleteItemFromStore(storeIndex int, itemIndex int) {
	Stores[storeIndex].Items = append(Stores[storeIndex].Items[:itemIndex], Stores[storeIndex].Items[itemIndex+1:]...)
}
