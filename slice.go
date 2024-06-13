package slice

import "sync"

type SafeSlice[T any] struct {
	slice []T        //nolint:structcheck
	mu    sync.Mutex //nolint:structcheck
}

func (s *SafeSlice[T]) Append(x T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice = append(s.slice, x)
}

func (s *SafeSlice[T]) Get(i int) T {
	s.mu.Lock()
	defer s.mu.Unlock()
	if i < 0 || i >= len(s.slice) {
		var zero T
		return zero
	}
	return s.slice[i]
}

func (s *SafeSlice[T]) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.slice)
}

func (s *SafeSlice[T]) Export() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	exportedSlice := make([]T, len(s.slice))
	copy(exportedSlice, s.slice)
	return exportedSlice
}
