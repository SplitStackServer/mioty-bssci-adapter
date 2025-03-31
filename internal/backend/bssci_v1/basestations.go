package bssci_v1

import (
	"mioty-bssci-adapter/internal/backend/events"
	"mioty-bssci-adapter/internal/common"
	"sync"

	"github.com/pkg/errors"
)

type basestations struct {
	sync.RWMutex
	basestations          map[common.EUI64]*connection
	subscribeEventHandler func(events.Subscribe)
}

func (b *basestations) get(eui common.EUI64) (*connection, error) {
	b.RLock()
	defer b.RUnlock()

	gw, ok := b.basestations[eui]
	if !ok {
		return gw, errors.New("basestation does not exist")
	}
	return gw, nil
}

func (b *basestations) set(eui common.EUI64, c *connection) error {
	b.Lock()
	defer b.Unlock()

	b.basestations[eui] = c

	if b.subscribeEventHandler != nil {
		b.subscribeEventHandler(events.Subscribe{Subscribe: true, BasestationEui: eui})
	}
	return nil
}

func (b *basestations) remove(eui common.EUI64) error {
	b.Lock()
	defer b.Unlock()

	if b.subscribeEventHandler != nil {
		b.subscribeEventHandler(events.Subscribe{Subscribe: false, BasestationEui: eui})
	}

	delete(b.basestations, eui)
	return nil
}
