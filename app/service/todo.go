package service

import (
	"sync"

	"stencil-go/app/model"
)

type Todo interface {
	Get(owner string) []model.Item
	Save(owner string, newItems []model.Item) error
}

type MemoryTodo struct {
	// key = session id, value the list of todo items that this session id has.
	items map[string][]model.Item
	// protected by locker for concurrent access.
	mu sync.RWMutex
}

func NewMemoryTodo() *MemoryTodo {
	return &MemoryTodo{
		items: make(map[string][]model.Item, 0),
	}
}

func (s *MemoryTodo) Get(sessionOwner string) []model.Item {
	s.mu.RLock()
	items := s.items[sessionOwner]
	s.mu.RUnlock()

	return items
}

func (s *MemoryTodo) Save(sessionOwner string, newItems []model.Item) error {
	var prevID int64
	for i := range newItems {
		if newItems[i].ID == 0 {
			newItems[i].ID = prevID
			prevID++
		}
	}

	s.mu.Lock()
	s.items[sessionOwner] = newItems
	s.mu.Unlock()
	return nil
}
