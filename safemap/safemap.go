package safemap

import (
	"errors"
	"fmt"
	"sync"
)

// SafeMap is a thread-safe map.
type SafeMap[K comparable, V any] struct {
	sync.RWMutex
	m map[K]V
}

// NewSafeMap creates a new SafeMap.
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	sm := new(SafeMap[K, V])
	sm.m = make(map[K]V)
	return sm
}

// NewSafeMapFromMap creates a new SafeMap from a map.
func NewSafeMapFromMap[K comparable, V any](m map[K]V) *SafeMap[K, V] {
	sm := new(SafeMap[K, V])
	sm.m = make(map[K]V)
	for k, v := range sm.m {
		sm.m[k] = v
	}
	return sm
}

// NewSafeMapFromKeysValues creates a new SafeMap from keys and values.
func NewSafeMapFromKeysValues[K comparable, V any](keys []K, values []V) (*SafeMap[K, V], error) {
	sm := new(SafeMap[K, V])
	sm.m = make(map[K]V)
	if len(keys) != len(values) {
		return sm, errors.New("keys and values must have the same length")
	}

	for i := 0; i < len(keys); i++ {
		sm.m[keys[i]] = values[i]
	}

	return sm, nil
}

// NewSafeMapFromKeyValuePairs creates a new SafeMap from key-value pairs.
func NewSafeMapFromKeyValuePairs[K comparable, V any](keysValues []any) (*SafeMap[K, V], error) {
	sm := new(SafeMap[K, V])
	sm.m = make(map[K]V)
	// check if the length of keysValues is even
	if len(keysValues)%2 != 0 {
		return sm, errors.New("keysValues must have an even length")
	}

	for i := 0; i < len(keysValues); i += 2 {
		key, ok := keysValues[i].(K)
		if !ok {
			return sm, fmt.Errorf("key must be comparable: %v is of type: %T", keysValues[i], keysValues[i])
		}
		value, ok := keysValues[i+1].(V)
		if !ok {
			return sm, fmt.Errorf("problem with value: %v of type: %T", keysValues[i+1], keysValues[i+1])
		}

		sm.m[key] = value
	}

	return sm, nil
}

// Get returns the value associated with the key.
func (sm *SafeMap[K, V]) Get(k K) (V, bool) {
	sm.RLock()
	defer sm.RUnlock()
	var val V

	if val, ok := sm.m[k]; ok {
		return val, true
	}

	return val, false
}

// Set sets the value associated with the key.
func (sm *SafeMap[K, V]) Set(k K, v V) {
	sm.Lock()
	defer sm.Unlock()
	sm.m[k] = v
}

// SetNX sets the value associated with the key if the key does not exist.
func (sm *SafeMap[K, V]) SetNX(k K, v V) bool {
	sm.Lock()
	defer sm.Unlock()
	if _, ok := sm.m[k]; !ok {
		sm.m[k] = v
		return true
	}
	return false
}

// Delete deletes the key-value pair associated with the key.
func (sm *SafeMap[K, V]) Delete(k K) {
	sm.Lock()
	defer sm.Unlock()
	delete(sm.m, k)
}

// Pop deletes the key-value pair associated with the key and returns the value.
func (sm *SafeMap[K, V]) Pop(k K) (V, bool) {
	sm.Lock()
	defer sm.Unlock()
	v, ok := sm.m[k]
	delete(sm.m, k)
	return v, ok
}

// Len returns the number of key-value pairs.
func (sm *SafeMap[K, V]) Len() int {
	return len(sm.m)
}

// IsEmpty returns true if the map is empty.
func (sm *SafeMap[K, V]) IsEmpty() bool {
	return len(sm.m) == 0
}

// Clear deletes all key-value pairs.
func (sm *SafeMap[K, V]) Clear() {
	sm.Lock()
	defer sm.Unlock()
	sm.m = make(map[K]V)
}

// GetMap returns the underlying map.
// Attention: the returned map is not thread-safe.
func (sm *SafeMap[K, V]) GetMap() map[K]V {
	sm.RLock()
	defer sm.RUnlock()
	return sm.m
}

// GetKeys returns the keys of the map as a slice.
func (sm *SafeMap[K, V]) GetKeys() []K {
	sm.RLock()
	defer sm.RUnlock()
	keys := make([]K, 0, len(sm.m))
	for k := range sm.m {
		keys = append(keys, k)
	}
	return keys
}

// GetValues returns the values of the map as a slice.
func (sm *SafeMap[K, V]) GetValues() []V {
	sm.RLock()
	defer sm.RUnlock()
	values := make([]V, 0, len(sm.m))
	for _, v := range sm.m {
		values = append(values, v)
	}
	return values
}

// GetKeyValuePairs returns the key-value pairs of the map as a slice.
func (sm *SafeMap[K, V]) GetKeyValuePairs() []any {
	sm.RLock()
	defer sm.RUnlock()
	keysValues := make([]any, 0, len(sm.m)*2)
	for k, v := range sm.m {
		keysValues = append(keysValues, k)
		keysValues = append(keysValues, v)
	}
	return keysValues
}

// GetKeysValues returns the keys and values of the map as separate slices.
// The order of the keys and values is the same.
func (sm *SafeMap[K, V]) GetKeysValues() ([]K, []V) {
	sm.RLock()
	defer sm.RUnlock()
	keys := make([]K, 0, len(sm.m))
	values := make([]V, 0, len(sm.m))
	for k, v := range sm.m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

// Copy returns a new SafeMap with the same key-value pairs.
func (sm *SafeMap[K, V]) Copy() *SafeMap[K, V] {
	sm.RLock()
	defer sm.RUnlock()
	newSm := NewSafeMap[K, V]()
	for k, v := range sm.m {
		newSm.Set(k, v)
	}
	return newSm
}

// Export returns a new map with the same key-value pairs as the SafeMap.
func (sm *SafeMap[K, V]) Export() map[K]V {
	sm.RLock()
	defer sm.RUnlock()
	m := make(map[K]V)
	for k, v := range sm.m {
		m[k] = v
	}
	return m
}
