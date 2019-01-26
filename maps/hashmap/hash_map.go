package hashmap

import "sync"

// HashMap ..
type HashMap struct {
	m map[interface{}]interface{}
	sync.RWMutex
}

// New returns a new empty hashmap
func New() *HashMap {
	return &HashMap{
		m: make(map[interface{}]interface{}),
	}
}

// Put a value in hashmap
func (hm *HashMap) Put(key, val interface{}) {
	hm.Lock()
	defer hm.Unlock()

	hm.put(key, val)
}

// Get finds value by key and returns it, if found, otherwise it returns nil
func (hm *HashMap) Get(key interface{}) (interface{}, bool) {
	hm.RLock()
	defer hm.RUnlock()

	return hm.get(key)
}

// Contains return whether given key in exists in hashmap
func (hm *HashMap) Contains(key interface{}) bool {
	hm.RLock()
	defer hm.RUnlock()

	return hm.contains(key)
}

// Delete removes a value from hashmap
func (hm *HashMap) Delete(key interface{}) bool {
	hm.Lock()
	defer hm.Unlock()

	return hm.delete(key)
}

// Length returns total number of entries in hashmap
func (hm *HashMap) Length() int {
	hm.RLock()
	defer hm.RUnlock()

	return hm.length()
}

// Clear removes all the entries from hashmap
func (hm *HashMap) Clear() {
	hm.Lock()
	defer hm.Unlock()

	hm.clear()
}

// Keys returns list of keys present in hashmap
func (hm *HashMap) Keys() []interface{} {
	hm.RLock()
	defer hm.RUnlock()

	return hm.keys()
}

// Values returns list of values present in hashmap
func (hm *HashMap) Values() []interface{} {
	hm.RLock()
	defer hm.RUnlock()

	return hm.values()
}