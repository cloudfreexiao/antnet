package blackboard

// KSS ...
type KSS struct {
	Key   string
	Value *[]string
}

// SetStringSliceP ...
func (bb *BB) SetStringSliceP(key string, value *[]string) {
	bb.SetValue(key, value)
}

// SetStringSliceP ...
func SetStringSliceP(key string, value *[]string) {
	Singleton().SetStringSliceP(key, value)
}

// SetStringSlice ...
func (bb *BB) SetStringSlice(key string, value []string) {
	bb.SetValue(key, &value)
}

// SetStringSlice ...
func SetStringSlice(key string, value []string) {
	Singleton().SetStringSlice(key, value)
}

// StringSliceP ...
func (bb *BB) StringSliceP(key string) *[]string {
	i, kok := bb.Value(key)
	if !kok {
		return nil
	}
	value, ok := i.(*[]string)
	if !ok {
		return nil
	}
	return value
}

// StringSliceP ...
func StringSliceP(key string) *[]string {
	return Singleton().StringSliceP(key)
}

// AllStringSlice ...
func (bb *BB) AllStringSlice() []KSS {
	slice := make([]KSS, 0)
	for k, v := range bb.value {
		if ssv, ok := v.(*[]string); ok {
			slice = append(slice, KSS{k, ssv})
		}
	}
	return slice
}

// AllStringSlice ...
func AllStringSlice() []KSS {
	return Singleton().AllStringSlice()
}
