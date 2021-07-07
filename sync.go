package u

import (
	"sync"
)

// MutexMap manages a pool of mutexes that can be get by key.
// MutexMap is thread-safe.
type MutexMap struct {
	self    sync.Mutex
	entries map[string]*sync.RWMutex
}

// Lock locks a mutex by key, and returns a callback for unlocking unlock.
// Lock will automatically create a new mutex for new keys.
func (mm *MutexMap) Lock(key string) func() {
	mm.self.Lock()
	if mm.entries == nil {
		mm.entries = make(map[string]*sync.RWMutex)
	}
	if _, found := mm.entries[key]; !found {
		mm.entries[key] = &sync.RWMutex{}
	}

	entry := mm.entries[key]
	entry.Lock()
	mm.self.Unlock()
	return entry.Unlock
}

// RLock locks a mutex by key for reading, and returns a callback for unlocking unlock.
// RLock will automatically create a new mutex for new keys.
func (mm *MutexMap) RLock(key string) func() {
	mm.self.Lock()
	if mm.entries == nil {
		mm.entries = make(map[string]*sync.RWMutex)
	}
	if _, found := mm.entries[key]; !found {
		mm.entries[key] = &sync.RWMutex{}
	}

	entry := mm.entries[key]
	entry.RLock()
	mm.self.Unlock()
	return entry.RUnlock
}
