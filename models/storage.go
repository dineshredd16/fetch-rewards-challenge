package models

import "sync"

type InMemoryStore struct {
	sync.Mutex
	Receipts map[string]Receipt
	Users map[string]User
}

var InMemoryDatabase = InMemoryStore{
	Receipts: make(map[string]Receipt),
	Users: make(map[string]User),
}