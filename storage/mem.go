package storage

import (
	"github.com/reddec/state-machine/machine"
	"os"
	"sync"
)

type memItem struct {
	state machine.State
	data  []byte
}

type MemStorage struct {
	lock sync.RWMutex
	data map[string]memItem
}

func (ms *MemStorage) Append(ctx *machine.StateContext, state machine.State, e error) error {
	cp := make([]byte, len(ctx.Storage))
	copy(cp, ctx.Storage)
	item := memItem{
		state: state,
		data:  cp,
	}
	ms.lock.Lock()
	defer ms.lock.Unlock()
	if ms.data == nil {
		ms.data = make(map[string]memItem)
	}
	ms.data[ctx.ID] = item
	return nil
}

func (ms *MemStorage) Last(id string) (*machine.IncompleteStateContext, error) {
	if ms.data == nil {
		return nil, os.ErrNotExist
	}
	ms.lock.RLock()
	defer ms.lock.RUnlock()
	item, ok := ms.data[id]
	if !ok {
		return nil, os.ErrNotExist
	}
	cp := make([]byte, len(item.data))
	copy(cp, item.data)
	return &machine.IncompleteStateContext{ID: id, Current: item.state, Storage: cp}, nil
}
