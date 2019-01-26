package hashmap

func (hm *HashMap) put(key, val interface{}) {
	hm.m[key] = val
}

func (hm *HashMap) get(key interface{}) (interface{}, bool) {
	val, ok := hm.m[key]

	return val, ok
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
