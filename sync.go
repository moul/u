package u

import (
	"sync"
)

// MutexMap manages a pool of mutexes that can be get by key.
// MutexMap is thread-safe.
type MutexMap struct {
	self    sync.Mutex
	entries map[string]*sync.Mutex
}

// Lock locks a mutex by key, and returns a callback for unlocking unlock.
// Lock will automatically create a new mutex for new keys.
func (mm *MutexMap) Lock(key string) func() {
	mm.self.Lock()
	if mm.entries == nil {
		mm.entries = make(map[string]*sync.Mutex)
	}
	if _, found := mm.entries[key]; !found {
		mm.entries[key] = &sync.Mutex{}
	}

	entry := mm.entries[key]
	entry.Lock()
	mm.self.Unlock()
	return entry.Unlock
}
