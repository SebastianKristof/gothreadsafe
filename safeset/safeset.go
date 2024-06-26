package safeset

import (
	"fmt"
	"sync"
)

// Set represents a thread-safe set of elements of type T
type Set[T comparable] struct {
	mu    sync.RWMutex
	items map[T]struct{}
}

// NewSet creates and returns a new Set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]struct{}),
	}
}

// NewSetWithValues creates and returns a new Set with the given values
func NewSetWithValues[T comparable](values ...T) *Set[T] {
	s := NewSet[T]()
	for _, value := range values {
		s.Add(value)
	}
	return s
}

// Add adds an element to the set
func (s *Set[T]) Add(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[item] = struct{}{}
}

// AddWithCheck adds an element to the set and returns true if the element was already in the set
func (s *Set[T]) AddWithCheck(item T) (existed bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, existed = s.items[item]
	s.items[item] = struct{}{}
	return existed
}

// Remove removes an element from the set
func (s *Set[T]) Remove(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.items, item)
}

// Contains checks if an element is in the set
func (s *Set[T]) Contains(item T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.items[item]
	return exists
}

// Size returns the number of elements in the set
func (s *Set[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.items)
}

// Clear removes all elements from the set
func (s *Set[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = make(map[T]struct{})
}

// IsEmpty returns true if the set is empty
func (s *Set[T]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.items) == 0
}

// ToSlice returns a slice containing all elements in the set
func (s *Set[T]) ToSlice() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	slice := make([]T, 0, len(s.items))
	for item := range s.items {
		slice = append(slice, item)
	}
	return slice
}

// Union returns a new set that is the union of s and other
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	s.mu.RLock()
	defer s.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()

	unionSet := NewSet[T]()
	for item := range s.items {
		unionSet.Add(item)
	}
	for item := range other.items {
		unionSet.Add(item)
	}
	return unionSet
}

// Intersection returns a new set that is the intersection of s and other
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	s.mu.RLock()
	defer s.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()

	intersectionSet := NewSet[T]()
	for item := range s.items {
		if other.Contains(item) {
			intersectionSet.Add(item)
		}
	}
	return intersectionSet
}

// Difference returns a new set that is the difference of s and other
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	s.mu.RLock()
	defer s.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()

	differenceSet := NewSet[T]()
	for item := range s.items {
		if !other.Contains(item) {
			differenceSet.Add(item)
		}
	}
	return differenceSet
}

// IsSubsetOf returns true if s is a subset of other
func (s *Set[T]) IsSubsetOf(other *Set[T]) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()

	for item := range s.items {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// IsSupersetOf returns true if s is a superset of other
func (s *Set[T]) IsSupersetOf(other *Set[T]) bool {
	return other.IsSubsetOf(s)
}

// Equal returns true if s and other contain the same elements
func (s *Set[T]) Equal(other *Set[T]) bool {
	return s.IsSubsetOf(other) && s.IsSupersetOf(other)
}

// Clone returns a new set with the same elements as s
func (s *Set[T]) Clone() *Set[T] {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cloneSet := NewSet[T]()
	for item := range s.items {
		cloneSet.Add(item)
	}
	return cloneSet
}

// String returns a string representation of the set
func (s *Set[T]) String() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	str := "{"
	i := 0
	for item := range s.items {
		if i > 0 {
			str += ", "
		}
		str += fmt.Sprintf("%v", item)
		i++
	}
	str += "}"
	return str
}

// SymmetricDifference returns a new set that is the symmetric difference (XOR) of s and other
func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	s.mu.RLock()
	defer s.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()

	xorSet := NewSet[T]()
	for item := range s.items {
		if !other.Contains(item) {
			xorSet.Add(item)
		}
	}
	for item := range other.items {
		if !s.Contains(item) {
			xorSet.Add(item)
		}
	}
	return xorSet
}
