package utils

// CompareFunc is a function used to compare two values
// If value A is greater than value B it should return 1
// If value A is less than value B it should return -1
// if value A is equal to equal value it should return 0
type CompareFunc func(a, b interface{}) int

// FilterFunc is a function used to filter out some values
// If it returns true, value will be included, otherwise it will be dropped
type FilterFunc func(a interface{}) bool

// MapFunc is a function used to map/transform some values
type MapFunc func(a interface{}) interface{}
