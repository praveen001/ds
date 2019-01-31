package hashmap

func (hm *HashMap) set(key, val interface{}) bool {
	hm.m[key] = val
	return true
}

func (hm *HashMap) get(key interface{}) (interface{}, bool) {
	val, ok := hm.m[key]

	return val, ok
}

func (hm *HashMap) remove(key interface{}) (interface{}, bool) {
	v, ok := hm.m[key]
	delete(hm.m, key)

	return v, ok
}

func (hm *HashMap) keys() []interface{} {
	ks := make([]interface{}, hm.len())

	i := 0
	for k := range hm.m {
		ks[i] = k
		i++
	}

	return ks
}

func (hm *HashMap) values() []interface{} {
	vals := make([]interface{}, hm.len())

	i := 0
	for _, v := range hm.m {
		vals[i] = v
		i++
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
