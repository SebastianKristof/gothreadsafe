package safeslice

import (
	"sort"
	"sync"
)

// SafeSlice is a thread-safe implementation of a slice.
type SafeSlice[T any] struct {
	slice []T        //nolint:structcheck
	mu    sync.Mutex //nolint:structcheck
}

// NewSafeSlice creates a new SafeSlice.
func NewSafeSlice[T any]() *SafeSlice[T] {
	return &SafeSlice[T]{}
}

// NewSafeSliceFromSlice creates a new SafeSlice from the specified slice.
func NewSafeSliceFromSlice[T any](slice []T) *SafeSlice[T] {
	s := NewSafeSlice[T]()
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice = append(s.slice, slice...)

	return s
}

// Append appends an element to the SafeSlice.
func (s *SafeSlice[T]) Append(x T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice = append(s.slice, x)
}

// Get returns the element at the specified index in the SafeSlice.
// If the index is out of range, it returns the zero value of the element type.
func (s *SafeSlice[T]) Get(i int) T {
	s.mu.Lock()
	defer s.mu.Unlock()
	if i < 0 || i >= len(s.slice) {
		var zero T
		return zero
	}
	return s.slice[i]
}

// Len returns the length of the SafeSlice.
func (s *SafeSlice[T]) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.slice)
}

// Export returns a new slice containing a copy of the elements in the SafeSlice.
func (s *SafeSlice[T]) Export() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	exportedSlice := make([]T, len(s.slice))
	copy(exportedSlice, s.slice)
	return exportedSlice
}

// Values returns the elements in the SafeSlice.
func (s *SafeSlice[T]) Values() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.slice
}

// Range ?

// Clear removes all elements from the SafeSlice.
func (s *SafeSlice[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice = s.slice[:0]
}

// Swap swaps the elements at the specified indices in the SafeSlice.
func (s *SafeSlice[T]) Swap(i, j int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}

// Set sets the element at the specified index in the SafeSlice.
// If the index is out of range, it does nothing.
func (s *SafeSlice[T]) Set(i int, x T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if i >= 0 && i < len(s.slice) {
		s.slice[i] = x
	}
}

// Insert inserts the element at the specified index in the SafeSlice.
// If the index is out of range, it appends the element to the SafeSlice.
func (s *SafeSlice[T]) Insert(i int, x T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if i < 0 || i >= len(s.slice) {
		s.slice = append(s.slice, x)
		return
	}
	s.slice = append(s.slice[:i+1], s.slice[i:]...)
	s.slice[i] = x
}

// InsertMany inserts the elements at the specified index in the SafeSlice.
// If the index is out of range, it appends the elements to the SafeSlice.
func (s *SafeSlice[T]) InsertMany(i int, elements []T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if i < 0 || i >= len(s.slice) {
		s.slice = append(s.slice, elements...)
		return
	}
	s.slice = append(s.slice[:i], append(elements, s.slice[i:]...)...)
}

// Pop removes and returns the last element from the SafeSlice.
// If the SafeSlice is empty, it returns the zero value of the element type.
func (s *SafeSlice[T]) Pop() T {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.slice) == 0 {
		var zero T
		return zero
	}
	x := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return x
}

// PopFront removes and returns the first element from the SafeSlice.
// If the SafeSlice is empty, it returns the zero value of the element type.
func (s *SafeSlice[T]) PopFront() T {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.slice) == 0 {
		var zero T
		return zero
	}
	x := s.slice[0]
	s.slice = s.slice[1:]
	return x
}

// Push adds an element to the end of the SafeSlice.
// It is the same as Append.
func (s *SafeSlice[T]) Push(x T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice = append(s.slice, x)
}

// PushFront adds an element to the beginning of the SafeSlice.
func (s *SafeSlice[T]) PushFront(x T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice = append([]T{x}, s.slice...)
}

// Reverse reverses the elements in the SafeSlice.
func (s *SafeSlice[T]) Reverse() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, j := 0, len(s.slice)-1; i < j; i, j = i+1, j-1 {
		s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
	}
}

// Map applies the function to each element in the SafeSlice and returns a new SafeSlice.
func (s *SafeSlice[T]) Map(fn func(T) T) *SafeSlice[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := NewSafeSlice[T]()
	for _, e := range s.slice {
		result.Append(fn(e))
	}
	return result
}

// ForEach applies the function to each element in the SafeSlice in place.
func (s *SafeSlice[T]) ForEach(fn func(T) T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, e := range s.slice {
		s.slice[i] = fn(e)
	}
}

// Filter applies the function to each element in the SafeSlice and returns a new SafeSlice containing the elements for which the function returns true.
func (s *SafeSlice[T]) Filter(fn func(T) bool) *SafeSlice[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := NewSafeSlice[T]()
	for _, e := range s.slice {
		if fn(e) {
			result.Append(e)
		}
	}
	return result
}

// Reduce applies the function to each element in the SafeSlice and returns the accumulated value.
func (s *SafeSlice[T]) Reduce(fn func(T, T) T) T {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.slice) == 0 {
		var zero T
		return zero
	}
	result := s.slice[0]
	for i := 1; i < len(s.slice); i++ {
		result = fn(result, s.slice[i])
	}
	return result
}

// All checks if all elements in the SafeSlice satisfy the predicate.
func (s *SafeSlice[T]) All(fn func(T) bool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, e := range s.slice {
		if !fn(e) {
			return false
		}
	}
	return true
}

// Any checks if any element in the SafeSlice satisfies the predicate.
func (s *SafeSlice[T]) Any(fn func(T) bool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, e := range s.slice {
		if fn(e) {
			return true
		}
	}
	return false
}

// Find returns the first element in the SafeSlice that satisfies the predicate.
// If no element satisfies the predicate, it returns the zero value of the element type.
func (s *SafeSlice[T]) Find(fn func(T) bool) T {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, e := range s.slice {
		if fn(e) {
			return e
		}
	}
	var zero T
	return zero
}

// FindIndex returns the index of the first element in the SafeSlice that satisfies the predicate.
// If no element satisfies the predicate, it returns -1.
func (s *SafeSlice[T]) FindIndex(fn func(T) bool) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, e := range s.slice {
		if fn(e) {
			return i
		}
	}
	return -1
}

// FindLast returns the last element in the SafeSlice that satisfies the predicate.
// If no element satisfies the predicate, it returns the zero value of the element type.
func (s *SafeSlice[T]) FindLast(fn func(T) bool) T {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := len(s.slice) - 1; i >= 0; i-- {
		if fn(s.slice[i]) {
			return s.slice[i]
		}
	}
	var zero T
	return zero
}

// FindLastIndex returns the index of the last element in the SafeSlice that satisfies the predicate.
// If no element satisfies the predicate, it returns -1.
func (s *SafeSlice[T]) FindLastIndex(fn func(T) bool) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := len(s.slice) - 1; i >= 0; i-- {
		if fn(s.slice[i]) {
			return i
		}
	}
	return -1
}

// Remove removes all elements in the SafeSlice that satisfy the predicate.
func (s *SafeSlice[T]) Remove(fn func(T) bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var result []T
	for _, e := range s.slice {
		if !fn(e) {
			result = append(result, e)
		}
	}
	s.slice = result
}

// RemoveAt removes the element at the specified index in the SafeSlice.
func (s *SafeSlice[T]) RemoveAt(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if i < 0 || i >= len(s.slice) {
		return
	}
	s.slice = append(s.slice[:i], s.slice[i+1:]...)
}

// SplitByFilter splits the SafeSlice into two SafeSlices based on the predicate.
func (s *SafeSlice[T]) SplitByFilter(fn func(T) bool) (*SafeSlice[T], *SafeSlice[T]) {
	s.mu.Lock()
	defer s.mu.Unlock()
	left := NewSafeSlice[T]()
	right := NewSafeSlice[T]()
	for _, e := range s.slice {
		if fn(e) {
			left.Append(e)
		} else {
			right.Append(e)
		}
	}
	return left, right
}

// SplitAtIndex splits the SafeSlice into two SafeSlices at the specified index.
func (s *SafeSlice[T]) SplitAtIndex(i int) (*SafeSlice[T], *SafeSlice[T]) {
	s.mu.Lock()
	defer s.mu.Unlock()
	left := NewSafeSlice[T]()
	right := NewSafeSlice[T]()
	for j, e := range s.slice {
		if j < i {
			left.Append(e)
		} else {
			right.Append(e)
		}
	}
	return left, right
}

// SortBy sorts the SafeSlice in place using the specified comparison function.
func (s *SafeSlice[T]) SortBy(less func(T, T) bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sort.Slice(s.slice, func(i, j int) bool {
		return less(s.slice[i], s.slice[j])
	})
}

// Copy returns a new SafeSlice containing a copy of the elements in the SafeSlice.
func (s *SafeSlice[T]) Copy() *SafeSlice[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := NewSafeSlice[T]()
	result.slice = make([]T, len(s.slice))
	copy(result.slice, s.slice)
	return result
}

// comparable

type SafeSliceComparable[T comparable] struct {
	slice []T        //nolint:structcheck
	mu    sync.Mutex //nolint:structcheck
}

func NewSafeSliceComparable[T comparable]() *SafeSliceComparable[T] {
	return &SafeSliceComparable[T]{}
}

// Contains checks if the SafeSlice contains the specified element.
func (s *SafeSliceComparable[T]) Contains(x T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, e := range s.slice {
		if e == x {
			return true
		}
	}
	return false
}

// Remove removes the first occurrence of the specified element from the SafeSlice.
func (s *SafeSliceComparable[T]) Remove(x T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, e := range s.slice {
		if e == x {
			s.slice = append(s.slice[:i], s.slice[i+1:]...)
			break
		}
	}
}

// RemoveAll removes all occurrences of the specified element from the SafeSlice.
func (s *SafeSliceComparable[T]) RemoveAll(x T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var result []T
	for _, e := range s.slice {
		if e != x {
			result = append(result, e)
		}
	}
	s.slice = result
}

// Equal checks if the SafeSlice is equal to the specified SafeSlice.
func (s *SafeSliceComparable[T]) Equal(other *SafeSlice[T]) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	other.mu.Lock()
	defer other.mu.Unlock()
	if len(s.slice) != len(other.slice) {
		return false
	}
	for i, e := range s.slice {
		if e != other.slice[i] {
			return false
		}
	}
	return true
}

// EqualSlice checks if the SafeSlice is equal to the specified slice.
func (s *SafeSliceComparable[T]) EqualSlice(other []T) bool {
	return s.EqualFunc(other, func(a, b T) bool {
		return a == b
	})
}

// EqualValues checks if the SafeSlice is equal to the specified values.
func (s *SafeSliceComparable[T]) EqualValues(values ...T) bool {
	return s.EqualSlice(values)
}

// EqualSafeSlice checks if the SafeSlice is equal to the specified SafeSlice.
func (s *SafeSliceComparable[T]) EqualSafeSlice(other *SafeSlice[T]) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	other.mu.Lock()
	defer other.mu.Unlock()
	return s.Equal(other)
}

// EqualSafeSliceFunc checks if the SafeSlice is equal to the specified SafeSlice using the specified comparison function.
func (s *SafeSliceComparable[T]) EqualSafeSliceFunc(other *SafeSlice[T], fn func(T, T) bool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	other.mu.Lock()
	defer other.mu.Unlock()
	return s.EqualFunc(other.slice, fn)
}

// EqualFunc checks if the SafeSlice is equal to the specified slice using the specified comparison function.
func (s *SafeSliceComparable[T]) EqualFunc(other []T, fn func(T, T) bool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.slice) != len(other) {
		return false
	}
	for i, e := range s.slice {
		if !fn(e, other[i]) {
			return false
		}
	}
	return true
}
