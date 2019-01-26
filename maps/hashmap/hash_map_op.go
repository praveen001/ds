package hashmap

func (hm *HashMap) put(key, val interface{}) {
	hm.m[key] = val
}

func (hm *HashMap) get(key interface{}) (interface{}, bool) {
	val, ok := hm.m[key]

	return val, ok
}

func (hm *HashMap) contains(key interface{}) bool {
	_, ok := hm.m[key]

	return ok
}

func (hm *HashMap) delete(key interface{}) bool {
	_, ok := hm.m[key]
	delete(hm.m, key)

	return ok
}

func (hm *HashMap) length() int {
	return len(hm.m)
}

func (hm *HashMap) clear() {
	hm.m = make(map[interface{}]interface{})
}

func (hm *HashMap) keys() []interface{} {
	ks := make([]interface{}, hm.length())

	i := 0
	for k := range hm.m {
		ks[i] = k
		i++
	}

	return ks
}

func (hm *HashMap) values() []interface{} {
	vals := make([]interface{}, hm.length())

	i := 0
	for _, v := range hm.m {
		vals[i] = v
		i++
	}

	return vals
}
