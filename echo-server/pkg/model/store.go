package model

import "github.com/google/uuid"

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
	Stores = append(Stores, Store{Name: "Saturn", Id: uuid.New()})
}
