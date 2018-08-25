package blackboard

// BB ...
type BB struct {
	value map[string]interface{}
}

// NewBlackboard ...
func NewBlackboard() *BB {
	return &BB{make(map[string]interface{}, 0)}
}

var (
	bb = NewBlackboard()
)

// Singleton ...
func Singleton() *BB {
	return bb
}

// SetValue ...
func (bb *BB) SetValue(key string, value interface{}) {
	bb.value[key] = value
}

// SetValue ...
func SetValue(key string, value interface{}) {
	Singleton().SetValue(key, value)
}

// Value ...
func (bb *BB) Value(key string) (v interface{}, ok bool) {
	v, ok = bb.value[key]
	return
}

// Value ...
func Value(key string) (interface{}, bool) {
	return Singleton().Value(key)
}
