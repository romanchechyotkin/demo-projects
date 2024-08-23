package statistics

import (
	"context"
	"sync"
	"time"
)

type (
	Writer interface {
		Insert(rows Rows) error
	}

	Manager struct {
		writer        Writer
		flushInterval time.Duration
		ctx           context.Context
		cancel        context.CancelFunc
		mu            sync.Mutex
		rows          Rows
	}
)

const defaultCapacity = 1000

func NewManager() *Manager {
	ctx, cancel := context.WithCancel(context.Background())

	return &Manager{
		writer:        nil,
		flushInterval: 0,
		ctx:           ctx,
		cancel:        cancel,
		rows:          NewRows(),
	}
}

func NewRows() Rows {
	return make(Rows, defaultCapacity)
}
