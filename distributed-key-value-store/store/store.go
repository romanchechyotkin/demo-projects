package store

import (
	"errors"
	"log"
	"sync"

	"kv-store/logger"
)

var ErrorKeyExists = errors.New("record due to this key exists")
var ErrorNoSuchKey = errors.New("such key does not exist")

type Store interface {
	AddRecord(key, value string) error
	GetRecord(key string) (string, error)
	DeleteRecord(key string) error
	InitLog() error
	WritePostLog(key, value string)
	WriteDeleteLog(key string)
}

type store struct {
	m map[string]string
	mux sync.RWMutex
	logger logger.TransactionLogger
}

func NewStore(l logger.TransactionLogger) Store {
	s := &store{
		m: make(map[string]string, 0),
		logger: l,
	}
	return s
}

func (s *store) AddRecord(key, value string) error {
	if _, ok := s.m[key]; ok {
		return ErrorKeyExists
	}
	s.mux.Lock()
	s.m[key] = value
	s.mux.Unlock()

	log.Printf("add key %s value %s", key, value)
	return nil
}

func (s *store) GetRecord(key string) (string, error) {
	s.mux.RLock()
	value, ok := s.m[key]
	s.mux.RUnlock()
	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func (s *store) DeleteRecord(key string) error {
	if _, ok := s.m[key]; !ok {
		return ErrorNoSuchKey
	}
	s.mux.Lock()
	delete(s.m, key)
	s.mux.Unlock()
	
	log.Println("deleted key", key)
	return nil 
}

func (s *store) InitLog() error {
	s.logger.RemoveRedundantLogs()
	events, errs := s.logger.ReadEvents()
	e, ok := logger.Event{}, true

	var err error
	for ok && err == nil {
		select {
		case err, ok = <-errs:
		case e, ok = <- events:
			switch e.EventType {
			case logger.EventPost:
				s.AddRecord(e.Key, e.Value)
			case logger.EventDelete:
				s.DeleteRecord(e.Key)
			}
		}	
	}

	s.logger.Run()
	return err
}

func (s *store) WritePostLog(key, value string) {
	s.logger.WritePost(key, value)
}

func (s *store) WriteDeleteLog(key string) {
	s.logger.WriteDelete(key)
}