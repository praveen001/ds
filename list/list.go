package list

// List ..
type List interface {
	Append(values ...interface{}) *List
	Prepend(values ...interface{}) *List
	Get(index int) (interface{}, error)
	Set(index int, value interface{}) error
	Remove(index int) (interface{}, error)
	Contains(value int) bool
	IndexOf(value int) int
	Size() int
	WithInRange(index int) bool
	String() string
}
