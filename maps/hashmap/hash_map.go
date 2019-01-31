package hashmap

import "sync"

// HashMap ..
type HashMap struct {
	m    map[interface{}]interface{}
	mtx  sync.RWMutex
	sync bool
}

// New returns a new empty hashmap
func New() *HashMap {
	return &HashMap{
		m:    make(map[interface{}]interface{}),
		sync: true,
	}
}

// Set a value in hashmap
func (hm *HashMap) Set(key, val interface{}) {
	hm.lock()
	defer hm.unlock()

	hm.set(key, val)
}

// Get finds value by key and returns it, if found, otherwise it returns nil
func (hm *HashMap) Get(key interface{}) (interface{}, bool) {
	hm.rlock()
	defer hm.runlock()

	return hm.get(key)
}

// Remove a value from hashmap
func (hm *HashMap) Remove(key interface{}) bool {
	hm.lock()
	defer hm.unlock()

	return hm.remove(key)
}

// Keys returns list of keys present in hashmap
func (hm *HashMap) Keys() []interface{} {
	hm.rlock()
	defer hm.runlock()

	return hm.keys()
}

// Values returns list of values present in hashmap
func (hm *HashMap) Values() []interface{} {
	hm.rlock()
	defer hm.runlock()

	return hm.values()
}

// Contains return whether given key in exists in hashmap
func (hm *HashMap) Contains(key interface{}) bool {
	hm.rlock()
	defer hm.runlock()

	return hm.contains(key)
}

// Len returns total number of entries in hashmap
func (hm *HashMap) Len() int {
	hm.rlock()
	defer hm.runlock()

	return hm.len()
}

// Clear removes all the entries from hashmap
func (hm *HashMap) Clear() {
	hm.lock()
	defer hm.unlock()

	hm.clear()
}

func (hm *HashMap) lock() {
	if hm.sync {
		hm.mtx.Lock()
	}
}

func (hm *HashMap) unlock() {
	if hm.sync {
		hm.mtx.Unlock()
	}
}

func (hm *HashMap) rlock() {
	if hm.sync {
		hm.mtx.RLock()
	}
}

func (hm *HashMap) runlock() {
	if hm.sync {
		hm.mtx.RUnlock()
	}
}
