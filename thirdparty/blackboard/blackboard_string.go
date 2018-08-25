package blackboard

// KS ...
type KS struct {
	Key   string
	Value *string
}

// SetStringP ...
func (bb *BB) SetStringP(key string, value *string) {
	bb.SetValue(key, value)
}

// SetStringP ...
func SetStringP(key string, value *string) {
	Singleton().SetStringP(key, value)
}

// SetString ...
func (bb *BB) SetString(key string, value string) {
	bb.SetValue(key, &value)
}

// SetString ...
func SetString(key string, value string) {
	Singleton().SetString(key, value)
}

// StringP ...
func (bb *BB) StringP(key string) *string {
	i, kok := bb.Value(key)
	if !kok {
		return nil
	}
	value, ok := i.(*string)
	if !ok {
		return nil
	}
	return value
}

// StringP ...
func StringP(key string) *string {
	return Singleton().StringP(key)
}

// AllString ...
func (bb *BB) AllString() []KS {
	slice := make([]KS, 0)
	for k := range bb.value {
		if sv := bb.StringP(k); sv != nil {
			slice = append(slice, KS{k, sv})
		}
	}
	return slice
}

// AllString ...
func AllString() []KS {
	return Singleton().AllString()
}
