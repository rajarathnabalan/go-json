package jsonnode

// KeyValue :
type KeyValue struct {
	Key   string
	Value interface{}
}

// Util :
type Util struct {
}

// AddToMap :
func AddToMap(mapValue map[string]interface{}, keyValues ...KeyValue) map[string]interface{} {
	for _, keyValue := range keyValues {
		mapValue[keyValue.Key] = keyValue.Value
	}
	return mapValue
}

// AddToMap :
func (u Util) AddToMap(mapValue map[string]interface{}, keyValues ...KeyValue) map[string]interface{} {
	return AddToMap(mapValue, keyValues...)
}

// RemoveFromMap :
func RemoveFromMap(mapValue map[string]interface{}, keys ...string) map[string]interface{} {
	for _, key := range keys {
		delete(mapValue, key)
	}
	return mapValue
}

// RemoveFromMap :
func (u Util) RemoveFromMap(mapValue map[string]interface{}, keys ...string) map[string]interface{} {
	return RemoveFromMap(mapValue, keys...)
}

// NewKeyValue :
func NewKeyValue(key string, value interface{}) KeyValue {
	newKeyValue := KeyValue{Key: key, Value: value}
	return newKeyValue
}

// NewKeyValue :
func (u Util) NewKeyValue(key string, value interface{}) KeyValue {
	return NewKeyValue(key, value)
}

// MakeMap :
func MakeMap() map[string]interface{} {
	return make(map[string]interface{}, 0)
}

// MakeMap :
func (u Util) MakeMap() map[string]interface{} {
	return MakeMap()
}

// MakeMapArray :
func MakeMapArray() []interface{} {
	return make([]interface{}, 0)
}

// MakeMapArray :
func (u Util) MakeMapArray() []interface{} {
	return MakeMapArray()
}
