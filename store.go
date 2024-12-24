package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

type record struct {
	Key, URL string
}
type URLStore struct {
	urls     map[string]string
	mu       sync.RWMutex
	file     *os.File
	saveChan chan record
}

func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

func (s *URLStore) Set(key, value string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, present := s.urls[key]; present {
		return false
	}
	s.urls[key] = value
	return true
}

func (s *URLStore) Put(url string) string {
	var key string
	for {
		key = generateShortURL()
		if s.Set(key, url) {
			s.saveChan <- record{key, url}
			return key
		}
	}
	panic("URLStore Put: We shouldn't even be here")
}

func (s *URLStore) save(filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Println("Error closing file: ", err)
		}
	}(f)

	enc := json.NewEncoder(f)

	for {
		if err := enc.Encode(<-s.saveChan); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}

}

func (s *URLStore) load() error {
	dec := json.NewDecoder(s.file)

	for {
		var r record
		if err := dec.Decode(&r); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		s.Set(r.Key, r.URL)
	}
}

func NewURLStore(filename string) *URLStore {
	store := &URLStore{
		urls:     make(map[string]string),
		saveChan: make(chan record, 100),
	}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("URLStore: ", err)
	}
	store.file = f
	if err = store.load(); err != nil {
		log.Println("URLStore loading: ", err)
	}
	go store.save(filename)
	return store
}
