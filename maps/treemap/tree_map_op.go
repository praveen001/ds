package treemap

import (
	"github.com/praveen001/ds/ds"
)

func (tm *TreeMap) set(key, val interface{}) bool {
	return tm.m.Set(key, val)
}

func (tm *TreeMap) get(key interface{}) (interface{}, bool) {
	return tm.m.Get(key)
}

func (tm *TreeMap) remove(key interface{}) bool {
	return tm.m.Remove(key)
}

func (tm *TreeMap) keys() ds.List {
	return tm.m.Keys()
}

func (tm *TreeMap) values() ds.List {
	return tm.m.Values()
}

func (tm *TreeMap) contains(key interface{}) bool {
	return tm.m.Contains(key)
}

func (tm *TreeMap) len() int {
	return tm.m.Len()
}

func (tm *TreeMap) clear() {
	tm.m.Clear()
}
