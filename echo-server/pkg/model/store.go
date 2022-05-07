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

func GetStoreByID(id uuid.UUID) (*Store, error) {
	for _, store := range Stores {
		if store.Id == id {
			return &store, nil
		}
	}

	return nil, errors.New("no store found")
}

func GetItemByStoreAndID(store *Store, id uuid.UUID) (*Item, error) {
	for _, item := range store.Items {
		if item.Id == id {
			return &item, nil
		}
	}

	return nil, errors.New("no item found")
}

func AddItemToStore(id uuid.UUID, item Item) {
	for _, store := range Stores {
		if store.Id == id {
			store.Items = append(store.Items, item)
			return
		}
	}
}
