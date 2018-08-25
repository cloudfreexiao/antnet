package blackboard

// KB ...
type KB struct {
	Key   string
	Value *bool
}

// SetBoolP ...
func (bb *BB) SetBoolP(key string, value *bool) {
	bb.SetValue(key, value)
}

// SetBoolP ...
func SetBoolP(key string, value *bool) {
	Singleton().SetBoolP(key, value)
}

// SetBool ...
func (bb *BB) SetBool(key string, value bool) {
	bb.SetValue(key, &value)
}

// SetBool ...
func SetBool(key string, value bool) {
	Singleton().SetBool(key, value)
}

// BoolP ...
func (bb *BB) BoolP(key string) *bool {
	i, kok := bb.Value(key)
	if !kok {
		return nil
	}
	value, ok := i.(*bool)
	if !ok {
		return nil
	}
	return value
}

// BoolP ...
func BoolP(key string) *bool {
	return Singleton().BoolP(key)
}

// AllBool ...
func (bb *BB) AllBool() []KB {
	slice := make([]KB, 0)
	for k := range bb.value {
		if bv := bb.BoolP(k); bv != nil {
			slice = append(slice, KB{k, bv})
		}
	}
	return slice
}

// AllBool ...
func AllBool() []KB {
	return Singleton().AllBool()
}
