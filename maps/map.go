package maps

// Map ..
type Map interface {
	// Set a value in map
	Set(key, val interface{})

	// Get finds value by key and returns it, if found, otherwise it returns nil
	Get(key interface{}) (interface{}, bool)

	// Remove a value from map
	Remove(key interface{}) bool

	// Keys returns list of keys present in map
	Keys() []interface{}

	// Values returns list of values present in map
	Values() []interface{}

	// Contains return whether given key exists in map
	Contains(key interface{}) bool

	// Length returns total number of entries in map
	Length() int

	// Clear removes all the entries from map
	Clear()
}
