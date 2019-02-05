package hashmap

import (
	"github.com/praveen001/ds/ds"
	"github.com/praveen001/ds/list/linkedlist"
)

func (hm *HashMap) set(key, val interface{}) bool {
	hm.m[key] = val
	return true
}

func (hm *HashMap) get(key interface{}) (interface{}, bool) {
	val, ok := hm.m[key]

	return val, ok
}

func (hm *HashMap) remove(key interface{}) bool {
	_, ok := hm.m[key]
	delete(hm.m, key)

	return ok
}

func (hm *HashMap) keys() ds.List {
	ks := linkedlist.New()

	for k := range hm.m {
		ks.PushBack(k)
	}

	return ks
}

func (hm *HashMap) values() ds.List {
	vals := linkedlist.New()

	for _, v := range hm.m {
		vals.PushBack(v)
	}

	return vals
}

func (hm *HashMap) contains(key interface{}) bool {
	_, ok := hm.m[key]

	return ok
}

func (hm *HashMap) len() int {
	return len(hm.m)
}

func (hm *HashMap) clear() {
	hm.m = make(map[interface{}]interface{})
}
