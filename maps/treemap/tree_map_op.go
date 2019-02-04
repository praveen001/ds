package treemap

func (tm *TreeMap) set(key, val interface{}) bool {
	return tm.m.Set(key, val)
}

func (tm *TreeMap) get(key interface{}) (interface{}, bool) {
	return tm.m.Get(key)
}

func (tm *TreeMap) remove(key interface{}) bool {
	return tm.m.Remove(key)
}

// TODO
func (tm *TreeMap) keys() []interface{} {
	return []interface{}{}
}

func (tm *TreeMap) values() []interface{} {
	return tm.m.InOrder().Values()
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
