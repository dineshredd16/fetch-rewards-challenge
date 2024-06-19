package models

import "sync"

type InMemoryStore struct {
	sync.Mutex
	Receipts map[string]Receipt
}

var InMemoryDatabase = InMemoryStore{
	Receipts: make(map[string]Receipt),
}