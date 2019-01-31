package ds

// Map ..
type Map interface {
	// Set a value in map
	Set(k, v interface{}) (ok bool)

	// Get finds value by key and returns it, if found, otherwise it returns nil
	Get(k interface{}) (v interface{}, ok bool)

	// Remove a value from map
	Remove(k interface{}) (v interface{}, ok bool)

	// Keys returns list of keys present in map
	Keys() (ks []interface{})

	// Values returns list of values present in map
	Values() (ks []interface{})

	// Contains return whether given key exists in map
	Contains(k interface{}) (ok bool)

	// Len returns total number of entries in map
	Len() (l int)

	// Clear removes all the entries from map
	Clear()
}
