package storage

import (
	"slices"
	"sync"
)

type Request struct {
	ID    int      `json:"id"`
	Links []string `json:"links"`
}

var (
	requests []Request
	nextID   int
	mu       sync.RWMutex
)

func AddRequest(links []string) int {
	mu.Lock()
	defer mu.Unlock()
	id := nextID
	nextID++
	requests = append(requests, Request{ID: id, Links: slices.Clone(links)}) // безопасное копирование
	return id
}

func GetRequestByID(id int) (Request, bool) {
	mu.RLock()
	defer mu.RUnlock()
	if id < 0 || id >= len(requests) {
		return Request{}, false
	}
	req := requests[id]
	req.Links = slices.Clone(req.Links) // защищаем оригинальные данные
	return req, true
}

func GetAllRequests() []Request {
	mu.RLock()
	defer mu.RUnlock()
	return slices.Clone(requests) // безопасное копирование
}
