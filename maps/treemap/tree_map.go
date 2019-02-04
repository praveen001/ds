package treemap

import (
	"sync"

	"github.com/praveen001/ds/ds"

	"github.com/praveen001/ds/utils"

	"github.com/praveen001/ds/tree/redblacktree"
)

// TreeMap ..
type TreeMap struct {
	m    ds.Tree
	mtx  sync.RWMutex
	sync bool
}

// New returns a new empty hashmap
func New(c utils.CompareFunc) *TreeMap {
	return &TreeMap{
		m:    redblacktree.New(c),
		sync: true,
	}
}

// Set a value in hashmap
func (tm *TreeMap) Set(key, val interface{}) bool {
	tm.lock()
	defer tm.unlock()

	return tm.set(key, val)
}

// Get finds value by key and returns it, if found, otherwise it returns nil
func (tm *TreeMap) Get(key interface{}) (interface{}, bool) {
	tm.rlock()
	defer tm.runlock()

	return tm.get(key)
}

// Remove a value from hashmap
func (tm *TreeMap) Remove(key interface{}) bool {
	tm.lock()
	defer tm.unlock()

	return tm.remove(key)
}

// Keys returns list of keys present in hashmap
func (tm *TreeMap) Keys() []interface{} {
	tm.rlock()
	defer tm.runlock()

	return tm.keys()
}

// Values returns list of values present in hashmap
func (tm *TreeMap) Values() []interface{} {
	tm.rlock()
	defer tm.runlock()

	return tm.values()
}

// Contains return whether given key in exists in hashmap
func (tm *TreeMap) Contains(key interface{}) bool {
	tm.rlock()
	defer tm.runlock()

	return tm.contains(key)
}

// Len returns total number of entries in hashmap
func (tm *TreeMap) Len() int {
	tm.rlock()
	defer tm.runlock()

	return tm.len()
}

// Clear removes all the entries from hashmap
func (tm *TreeMap) Clear() {
	tm.lock()
	defer tm.unlock()

	tm.clear()
}

func (tm *TreeMap) lock() {
	if tm.sync {
		tm.mtx.Lock()
	}
}

func (tm *TreeMap) unlock() {
	if tm.sync {
		tm.mtx.Unlock()
	}
}

func (tm *TreeMap) rlock() {
	if tm.sync {
		tm.mtx.RLock()
	}
}

func (tm *TreeMap) runlock() {
	if tm.sync {
		tm.mtx.RUnlock()
	}
}
