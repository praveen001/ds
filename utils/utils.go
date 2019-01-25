package utils

// Comparator is a function used to compare two values
// If value A is greater than value B it should return 1
// If value A is less than value B it should return -1
// if value A is equal to equal value it should return 0
type Comparator func(a, b interface{}) int
