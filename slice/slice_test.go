package slice

import (
	"reflect"
	"strings"
	"testing"
)

func TestSafeSlice_Append(t *testing.T) {
	tests := []struct {
		name           string
		elements       []any
		expectedLen    int
		expectedValues []any
	}{
		{"AppendInt", []any{10, 20, 30}, 3, []any{10, 20, 30}},
		{"AppendString", []any{"test1", "test2", "test3"}, 3, []any{"test1", "test2", "test3"}},
		{"AppendFloat", []any{3.14, 2.71, 1.618}, 3, []any{3.14, 2.71, 1.618}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Check if the length of the slice is as expected
			if s.Len() != tt.expectedLen {
				t.Errorf("Expected length %d, got %d", tt.expectedLen, s.Len())
			}

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if s.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, s.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_GetInt(t *testing.T) {
	tests := []struct {
		name          string
		elements      []int
		index         int
		expectedValue int
	}{
		{"ValidIndex", []int{1, 2, 3}, 1, 2},
		{"InvalidIndex", []int{1, 2, 3}, 3, 0},
		{"EmptySlice", []int{}, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[int]()
			for _, element := range tt.elements {
				s.Append(element)
			}

			result := s.Get(tt.index)

			if !reflect.DeepEqual(result, tt.expectedValue) {
				t.Errorf("Expected value: %v, got value: %v", tt.expectedValue, result)
			}
		})
	}
}

func TestSafeSlice_GetString(t *testing.T) {
	tests := []struct {
		name          string
		elements      []string
		index         int
		expectedValue string
	}{
		{"ValidIndex", []string{"test1", "test2", "test3"}, 1, "test2"},
		{"InvalidIndex", []string{"test1", "test2", "test3"}, 3, ""},
		{"EmptySlice", []string{}, 0, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[string]()
			for _, element := range tt.elements {
				s.Append(element)
			}

			result := s.Get(tt.index)

			if !reflect.DeepEqual(result, tt.expectedValue) {
				t.Errorf("Expected value: %v, got value: %v", tt.expectedValue, result)
			}
		})
	}
}

func TestSafeSlice_Len(t *testing.T) {
	tests := []struct {
		name     string
		elements []any
		expected int
	}{
		{"EmptySlice", []any{}, 0},
		{"IntSlice", []any{[]int{1, 2, 3}, []int{4, 5, 6}}, 2},
		{"StringSlice", []any{[]string{"test1", "test2", "test3"}, []string{"test4", "test5", "test6"}}, 2},
		{"FloatSlice", []any{[]float64{3.14, 2.71, 1.618}, []float64{2.718, 1.414, 0.577}}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()
			for _, element := range tt.elements {
				s.Append(element)
			}

			result := s.Len()

			if result != tt.expected {
				t.Errorf("Expected length: %d, got length: %d", tt.expected, result)
			}
		})
	}
}

func TestSafeSlice_ExportInt(t *testing.T) {
	tests := []struct {
		name           string
		elements       []int
		expectedExport []int
	}{
		{"ExportInt", []int{10, 20, 30}, []int{10, 20, 30}},
		{"ExportEmptyInt", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[int]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Export the slice
			exportedSlice := s.Export()

			// Check if the exported slice has the same length as the original slice
			if len(exportedSlice) != s.Len() {
				t.Errorf("Expected exported slice length %d, got %d", s.Len(), len(exportedSlice))
			}

			// Check if the exported slice contains the same elements as the original slice
			for i, expectedValue := range tt.expectedExport {
				if exportedSlice[i] != expectedValue {
					t.Errorf("Expected exported value %v, got %v", expectedValue, exportedSlice[i])
				}
			}
		})
	}
}

func TestSafeSlice_ExportString(t *testing.T) {
	tests := []struct {
		name           string
		elements       []string
		expectedExport []string
	}{
		{"ExportString", []string{"test1", "test2", "test3"}, []string{"test1", "test2", "test3"}},
		{"ExportEmptyString", []string{}, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SafeSlice[string]

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Export the slice
			exportedSlice := s.Export()

			// Check if the exported slice has the same length as the original slice
			if len(exportedSlice) != s.Len() {
				t.Errorf("Expected exported slice length %d, got %d", s.Len(), len(exportedSlice))
			}

			// Check if the exported slice contains the same elements as the original slice
			for i, expectedValue := range tt.expectedExport {
				if exportedSlice[i] != expectedValue {
					t.Errorf("Expected exported value %v, got %v", expectedValue, exportedSlice[i])
				}
			}
		})
	}
}

func TestSafeSlice_ExportFloat(t *testing.T) {
	tests := []struct {
		name           string
		elements       []float64
		expectedExport []float64
	}{
		{"ExportFloat", []float64{3.14, 2.71, 1.618}, []float64{3.14, 2.71, 1.618}},
		{"ExportEmptyFloat", []float64{}, []float64{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[float64]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Export the slice
			exportedSlice := s.Export()

			// Check if the exported slice has the same length as the original slice
			if len(exportedSlice) != s.Len() {
				t.Errorf("Expected exported slice length %d, got %d", s.Len(), len(exportedSlice))
			}

			// Check if the exported slice contains the same elements as the original slice
			for i, expectedValue := range tt.expectedExport {
				if exportedSlice[i] != expectedValue {
					t.Errorf("Expected exported value %v, got %v", expectedValue, exportedSlice[i])
				}
			}
		})
	}
}

func TestSafeSlice_ValuesInt(t *testing.T) {
	tests := []struct {
		name           string
		elements       []int
		expectedValues []int
	}{
		{"ValuesInt", []int{10, 20, 30}, []int{10, 20, 30}},
		{"ValuesEmptyInt", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[int]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Get the values
			values := s.Values()

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if values[i] != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, values[i])
				}
			}
		})
	}
}

func TestSafeSlice_ValuesString(t *testing.T) {
	tests := []struct {
		name           string
		elements       []string
		expectedValues []string
	}{
		{"ValuesString", []string{"test1", "test2", "test3"}, []string{"test1", "test2", "test3"}},
		{"ValuesEmptyString", []string{}, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[string]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Get the values
			values := s.Values()

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if values[i] != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, values[i])
				}
			}
		})
	}
}

func TestSafeSlice_ValuesFloat(t *testing.T) {
	tests := []struct {
		name           string
		elements       []float64
		expectedValues []float64
	}{
		{"ValuesFloat", []float64{3.14, 2.71, 1.618}, []float64{3.14, 2.71, 1.618}},
		{"ValuesEmptyFloat", []float64{}, []float64{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[float64]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Get the values
			values := s.Values()

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if values[i] != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, values[i])
				}
			}
		})
	}
}

func TestSafeSlice_Clear(t *testing.T) {
	tests := []struct {
		name     string
		elements []any
	}{
		{"ClearInt", []any{10, 20, 30}},
		{"ClearString", []any{"test1", "test2", "test3"}},
		{"ClearFloat", []any{3.14, 2.71, 1.618}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Clear the slice
			s.Clear()

			// Check if the slice is empty
			if s.Len() != 0 {
				t.Errorf("Expected length 0, got %d", s.Len())
			}
		})
	}
}

func TestSafeSlice_Swap(t *testing.T) {
	tests := []struct {
		name           string
		elements       []any
		index1         int
		index2         int
		expectedValues []any
	}{
		{"SwapInt", []any{10, 20, 30}, 0, 2, []any{30, 20, 10}},
		{"SwapString", []any{"test1", "test2", "test3"}, 0, 1, []any{"test2", "test1", "test3"}},
		{"SwapFloat", []any{3.14, 2.71, 1.618}, 1, 2, []any{3.14, 1.618, 2.71}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Swap the values
			s.Swap(tt.index1, tt.index2)

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if s.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, s.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_SetInt(t *testing.T) {
	tests := []struct {
		name           string
		elements       []int
		index          int
		newValue       int
		expectedValues []int
	}{
		{"SetInt", []int{10, 20, 30}, 1, 25, []int{10, 25, 30}},
		{"SetIntInvalidIndex", []int{10, 20, 30}, 3, 25, []int{10, 20, 30}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[int]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Set the value
			s.Set(tt.index, tt.newValue)

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if s.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, s.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_SetString(t *testing.T) {
	tests := []struct {
		name           string
		elements       []string
		index          int
		newValue       string
		expectedValues []string
	}{
		{"SetString", []string{"test1", "test2", "test3"}, 1, "newTest", []string{"test1", "newTest", "test3"}},
		{"SetStringInvalidIndex", []string{"test1", "test2", "test3"}, 3, "newTest", []string{"test1", "test2", "test3"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[string]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Set the value
			s.Set(tt.index, tt.newValue)

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if s.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, s.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_SetFloat(t *testing.T) {
	tests := []struct {
		name           string
		elements       []float64
		index          int
		newValue       float64
		expectedValues []float64
	}{
		{"SetFloat", []float64{3.14, 2.71, 1.618}, 1, 2.718, []float64{3.14, 2.718, 1.618}},
		{"SetFloatInvalidIndex", []float64{3.14, 2.71, 1.618}, 3, 2.718, []float64{3.14, 2.71, 1.618}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[float64]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Set the value
			s.Set(tt.index, tt.newValue)

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if s.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, s.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_SetStruct(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}

	tests := []struct {
		name           string
		elements       []testStruct
		index          int
		newValue       testStruct
		expectedValues []testStruct
	}{
		{"SetStruct", []testStruct{{"Alice", 25}, {"Bob", 30}, {"Charlie", 35}}, 1, testStruct{"David", 40}, []testStruct{{"Alice", 25}, {"David", 40}, {"Charlie", 35}}},
		{"SetStructInvalidIndex", []testStruct{{"Alice", 25}, {"Bob", 30}, {"Charlie", 35}}, 3, testStruct{"David", 40}, []testStruct{{"Alice", 25}, {"Bob", 30}, {"Charlie", 35}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[testStruct]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Set the value
			s.Set(tt.index, tt.newValue)

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if s.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, s.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_Insert(t *testing.T) {
	tests := []struct {
		name           string
		elements       []any
		index          int
		newValue       any
		expectedValues []any
	}{
		{"InsertInt", []any{10, 20, 30}, 1, 25, []any{10, 25, 20, 30}},
		{"InsertIntInvalidIndex", []any{10, 20, 30}, 4, 25, []any{10, 20, 30}},
		{"InsertString", []any{"test1", "test2", "test3"}, 1, "newTest", []any{"test1", "newTest", "test2", "test3"}},
		{"InsertStringInvalidIndex", []any{"test1", "test2", "test3"}, 4, "newTest", []any{"test1", "test2", "test3"}},
		{"InsertFloat", []any{3.14, 2.71, 1.618}, 1, 2.718, []any{3.14, 2.718, 2.71, 1.618}},
		{"InsertFloatInvalidIndex", []any{3.14, 2.71, 1.618}, 4, 2.718, []any{3.14, 2.71, 1.618}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Insert the value
			s.Insert(tt.index, tt.newValue)

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if s.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, s.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_InsertMany(t *testing.T) {
	tests := []struct {
		name           string
		elements       []any
		index          int
		newValues      []any
		expectedValues []any
	}{
		{"InsertManyInt", []any{10, 20, 30}, 1, []any{25, 35}, []any{10, 25, 35, 20, 30}},
		{"InsertManyIntInvalidIndex", []any{10, 20, 30}, 4, []any{25, 35}, []any{10, 20, 30}},
		{"InsertManyString", []any{"test1", "test2", "test3"}, 1, []any{"newTest1", "newTest2"}, []any{"test1", "newTest1", "newTest2", "test2", "test3"}},
		{"InsertManyStringInvalidIndex", []any{"test1", "test2", "test3"}, 4, []any{"newTest1", "newTest2"}, []any{"test1", "test2", "test3"}},
		{"InsertManyFloat", []any{3.14, 2.71, 1.618}, 1, []any{2.718, 1.414}, []any{3.14, 2.718, 1.414, 2.71, 1.618}},
		{"InsertManyFloatInvalidIndex", []any{3.14, 2.71, 1.618}, 4, []any{2.718, 1.414}, []any{3.14, 2.71, 1.618}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Insert the values
			s.InsertMany(tt.index, tt.newValues)

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if s.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, s.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_Pop(t *testing.T) {
	tests := []struct {
		name           string
		elements       []any
		expectedValues []any
	}{
		{"PopInt", []any{10, 20, 30}, []any{10, 20}},
		{"PopString", []any{"test1", "test2", "test3"}, []any{"test1", "test2"}},
		{"PopFloat", []any{3.14, 2.71, 1.618}, []any{3.14, 2.71}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Pop the value
			s.Pop()

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if s.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, s.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_PopFront(t *testing.T) {
	tests := []struct {
		name           string
		elements       []any
		expectedValues []any
	}{
		{"PopFrontInt", []any{10, 20, 30}, []any{20, 30}},
		{"PopFrontString", []any{"test1", "test2", "test3"}, []any{"test2", "test3"}},
		{"PopFrontFloat", []any{3.14, 2.71, 1.618}, []any{2.71, 1.618}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Pop the value
			s.PopFront()

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if s.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, s.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_Push(t *testing.T) {
	tests := []struct {
		name           string
		elements       []any
		newValue       any
		expectedValues []any
	}{
		{"PushInt", []any{10, 20, 30}, 40, []any{10, 20, 30, 40}},
		{"PushString", []any{"test1", "test2", "test3"}, "test4", []any{"test1", "test2", "test3", "test4"}},
		{"PushFloat", []any{3.14, 2.71, 1.618}, 1.414, []any{3.14, 2.71, 1.618, 1.414}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Push the value
			s.Push(tt.newValue)

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if s.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, s.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_PushFront(t *testing.T) {
	tests := []struct {
		name           string
		elements       []any
		newValue       any
		expectedValues []any
	}{
		{"PushFrontInt", []any{10, 20, 30}, 5, []any{5, 10, 20, 30}},
		{"PushFrontString", []any{"test1", "test2", "test3"}, "test0", []any{"test0", "test1", "test2", "test3"}},
		{"PushFrontFloat", []any{3.14, 2.71, 1.618}, 0.577, []any{0.577, 3.14, 2.71, 1.618}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Push the value
			s.PushFront(tt.newValue)

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if s.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, s.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_Reverse(t *testing.T) {
	tests := []struct {
		name           string
		elements       []any
		expectedValues []any
	}{
		{"ReverseInt", []any{10, 20, 30}, []any{30, 20, 10}},
		{"ReverseString", []any{"test1", "test2", "test3"}, []any{"test3", "test2", "test1"}},
		{"ReverseFloat", []any{3.14, 2.71, 1.618}, []any{1.618, 2.71, 3.14}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Reverse the slice
			s.Reverse()

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if s.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, s.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_Map(t *testing.T) {
	tests := []struct {
		name           string
		elements       []any
		fn             func(any) any
		expectedValues []any
	}{
		{"MapInt", []any{10, 20, 30}, func(x any) any { return x.(int) * 2 }, []any{20, 40, 60}},
		{"MapString", []any{"test1", "test2", "test3"}, func(x any) any { return x.(string) + "!" }, []any{"test1!", "test2!", "test3!"}},
		{"MapFloat", []any{3.14, 2.71, 1.618}, func(x any) any { return x.(float64) * 2 }, []any{6.28, 5.42, 3.236}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Map the values
			result := s.Map(tt.fn)

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if result.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, result.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_ForEach(t *testing.T) {
	tests := []struct {
		name           string
		elements       []any
		fn             func(any) any
		expectedValues []any
	}{
		{"ForEachInt", []any{10, 20, 30}, func(x any) any { return x.(int) * 2 }, []any{20, 40, 60}},
		{"ForEachString", []any{"test1", "test2", "test3"}, func(x any) any { return x.(string) + "!" }, []any{"test1!", "test2!", "test3!"}},
		{"ForEachFloat", []any{3.14, 2.71, 1.618}, func(x any) any { return x.(float64) * 2 }, []any{6.28, 5.42, 3.236}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			//Apply the function to each element
			s.ForEach(tt.fn)

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if s.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, s.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_Filter(t *testing.T) {
	tests := []struct {
		name           string
		elements       []any
		fn             func(any) bool
		expectedValues []any
	}{
		{"FilterInt", []any{10, 20, 30}, func(x any) bool { return x.(int) > 15 }, []any{20, 30}},
		{"FilterString", []any{"test1", "test2", "test3"}, func(x any) bool { return strings.HasSuffix(x.(string), "2") }, []any{"test2"}},
		{"FilterFloat", []any{3.14, 2.71, 1.618}, func(x any) bool { return x.(float64) < 3 }, []any{2.71, 1.618}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Filter the values
			result := s.Filter(tt.fn)

			// Check if the values are correct
			for i, expectedValue := range tt.expectedValues {
				if result.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, result.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_Reduce(t *testing.T) {
	tests := []struct {
		name          string
		elements      []any
		fn            func(any, any) any
		expectedValue any
	}{
		{"ReduceInt", []any{10, 20, 30}, func(a, b any) any { return a.(int) + b.(int) }, 60},
		{"ReduceString", []any{"test1", "test2", "test3"}, func(a, b any) any { return a.(string) + b.(string) }, "test1test2test3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()

			// Append values to the slice
			for _, element := range tt.elements {
				s.Append(element)
			}

			// Reduce the values
			result := s.Reduce(tt.fn)

			// Check if the values are correct
			if result != tt.expectedValue {
				t.Errorf("Expected value %v, got %v", tt.expectedValue, result)
			}
		})
	}
}

func TestSafeSlice_All(t *testing.T) {
	tests := []struct {
		name     string
		elements []any
		fn       func(any) bool
		expected bool
	}{
		{"AllIntTrue", []any{1, 2, 3}, func(x any) bool { return x.(int) > 0 }, true},
		{"AllIntFalse", []any{1, 2, -3}, func(x any) bool { return x.(int) > 0 }, false},
		{"AllStringTrue", []any{"test1", "test2", "test3"}, func(x any) bool { return len(x.(string)) > 0 }, true},
		{"AllStringFalse", []any{"test1", "", "test3"}, func(x any) bool { return len(x.(string)) > 0 }, false},
		{"AllFloatTrue", []any{3.14, 2.71, 1.618}, func(x any) bool { return x.(float64) > 0 }, true},
		{"AllFloatFalse", []any{3.14, -2.71, 1.618}, func(x any) bool { return x.(float64) > 0 }, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()
			for _, e := range tt.elements {
				s.Append(e)
			}
			result := s.All(tt.fn)
			if result != tt.expected {
				t.Errorf("Expected All() to return %v, but got %v", tt.expected, result)
			}
		})
	}
}

func TestSafeSlice_Any(t *testing.T) {
	tests := []struct {
		name     string
		elements []any
		fn       func(any) bool
		expected bool
	}{
		{"AnyIntTrue", []any{1, 2, 3}, func(x any) bool { return x.(int) > 2 }, true},
		{"AnyIntFalse", []any{1, 2, 3}, func(x any) bool { return x.(int) > 3 }, false},
		{"AnyStringTrue", []any{"test1", "test2", "test3"}, func(x any) bool { return strings.Contains(x.(string), "2") }, true},
		{"AnyStringFalse", []any{"test1", "test2", "test3"}, func(x any) bool { return strings.Contains(x.(string), "4") }, false},
		{"AnyFloatTrue", []any{3.14, 2.71, 1.618}, func(x any) bool { return x.(float64) < 2 }, true},
		{"AnyFloatFalse", []any{3.14, 2.71, 1.618}, func(x any) bool { return x.(float64) < 1 }, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()
			for _, e := range tt.elements {
				s.Append(e)
			}
			result := s.Any(tt.fn)
			if result != tt.expected {
				t.Errorf("Expected Any() to return %v, but got %v", tt.expected, result)
			}
		})
	}
}

func TestSafeSlice_Find(t *testing.T) {
	tests := []struct {
		name     string
		elements []any
		fn       func(any) bool
		expected any
	}{
		{"FindInt", []any{1, 2, 3}, func(x any) bool { return x.(int) > 1 }, 2},
		{"FindString", []any{"test1", "test2", "test3"}, func(x any) bool { return strings.Contains(x.(string), "2") }, "test2"},
		{"FindFloat", []any{3.14, 2.71, 1.618}, func(x any) bool { return x.(float64) < 2 }, 1.618},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()
			for _, e := range tt.elements {
				s.Append(e)
			}
			result := s.Find(tt.fn)
			if result != tt.expected {
				t.Errorf("Expected Find() to return %v, but got %v", tt.expected, result)
			}
		})
	}
}

func TestSafeSlice_FindIndex(t *testing.T) {
	tests := []struct {
		name     string
		elements []any
		fn       func(any) bool
		expected int
	}{
		{"FindIndexInt", []any{1, 2, 3}, func(x any) bool { return x.(int) > 1 }, 1},
		{"FindIndexString", []any{"test1", "test2", "test3"}, func(x any) bool { return strings.Contains(x.(string), "2") }, 1},
		{"FindIndexFloat", []any{3.14, 2.71, 1.618}, func(x any) bool { return x.(float64) < 2 }, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()
			for _, e := range tt.elements {
				s.Append(e)
			}
			result := s.FindIndex(tt.fn)
			if result != tt.expected {
				t.Errorf("Expected FindIndex() to return %v, but got %v", tt.expected, result)
			}
		})
	}
}

func TestSafeSlice_FindLast(t *testing.T) {
	tests := []struct {
		name     string
		elements []any
		fn       func(any) bool
		expected any
	}{
		{"FindLastInt", []any{1, 2, 3}, func(x any) bool { return x.(int) > 1 }, 3},
		{"FindLastString", []any{"test1", "test2", "test3"}, func(x any) bool { return strings.Contains(x.(string), "2") }, "test2"},
		{"FindLastFloat", []any{3.14, 2.71, 1.618}, func(x any) bool { return x.(float64) < 2 }, 1.618},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()
			for _, e := range tt.elements {
				s.Append(e)
			}
			result := s.FindLast(tt.fn)
			if result != tt.expected {
				t.Errorf("Expected FindLast() to return %v, but got %v", tt.expected, result)
			}
		})
	}
}

func TestSafeSlice_FindLastIndex(t *testing.T) {
	tests := []struct {
		name     string
		elements []any
		fn       func(any) bool
		expected int
	}{
		{"FindLastIndexInt", []any{1, 2, 3}, func(x any) bool { return x.(int) > 1 }, 2},
		{"FindLastIndexString", []any{"test1", "test2", "test3"}, func(x any) bool { return strings.Contains(x.(string), "2") }, 1},
		{"FindLastIndexFloat", []any{3.14, 2.71, 1.618}, func(x any) bool { return x.(float64) < 2 }, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()
			for _, e := range tt.elements {
				s.Append(e)
			}
			result := s.FindLastIndex(tt.fn)
			if result != tt.expected {
				t.Errorf("Expected FindLastIndex() to return %v, but got %v", tt.expected, result)
			}
		})
	}
}

func TestSafeSlice_Remove(t *testing.T) {
	tests := []struct {
		name           string
		elements       []any
		fn             func(any) bool
		expectedValues []any
	}{
		{"RemoveInt", []any{1, 2, 3}, func(x any) bool { return x.(int) > 1 }, []any{1}},
		{"RemoveString", []any{"test1", "test2", "test3"}, func(x any) bool { return strings.Contains(x.(string), "2") }, []any{"test1", "test3"}},
		{"RemoveFloat", []any{3.14, 2.71, 1.618}, func(x any) bool { return x.(float64) < 2 }, []any{3.14, 2.71}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()
			for _, e := range tt.elements {
				s.Append(e)
			}
			s.Remove(tt.fn)
			for i, expectedValue := range tt.expectedValues {
				if s.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, s.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_RemoveAt(t *testing.T) {
	tests := []struct {
		name           string
		elements       []any
		index          int
		expectedValues []any
	}{
		{"RemoveAtInt", []any{1, 2, 3}, 1, []any{1, 3}},
		{"RemoveAtString", []any{"test1", "test2", "test3"}, 2, []any{"test1", "test2"}},
		{"RemoveAtFloat", []any{3.14, 2.71, 1.618}, 0, []any{2.71, 1.618}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()
			for _, e := range tt.elements {
				s.Append(e)
			}
			s.RemoveAt(tt.index)
			for i, expectedValue := range tt.expectedValues {
				if s.Get(i) != expectedValue {
					t.Errorf("Expected value %v, got %v", expectedValue, s.Get(i))
				}
			}
		})
	}
}

func TestSafeSlice_SplitByFilter(t *testing.T) {
	tests := []struct {
		name                 string
		elements             []any
		fn                   func(any) bool
		expectedSatisfied    *SafeSlice[any]
		expectedNotSatisfied *SafeSlice[any]
	}{
		{"SplitInt", []any{1, 2, 3, 4, 5}, func(x any) bool { return x.(int)%2 == 0 },
			NewSafeSliceFromSlice([]any{2, 4}),
			NewSafeSliceFromSlice([]any{1, 3, 5}),
		},
		{"SplitString", []any{"test1", "test2", "test3", "test4", "test5"}, func(x any) bool { return strings.HasSuffix(x.(string), "2") },
			NewSafeSliceFromSlice([]any{"test2"}),
			NewSafeSliceFromSlice([]any{"test1", "test3", "test4", "test5"}),
		},
		{"SplitFloat", []any{1.1, 2.2, 3.3, 4.4, 5.5}, func(x any) bool { return int(x.(float64))%2 == 0 },
			NewSafeSliceFromSlice([]any{2.2, 4.4}),
			NewSafeSliceFromSlice([]any{1.1, 3.3, 5.5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()
			for _, e := range tt.elements {
				s.Append(e)
			}
			satisfied, notSatisfied := s.SplitByFilter(tt.fn)

			for i, el := range satisfied.Values() {
				if el != tt.expectedSatisfied.Get(i) {
					t.Errorf("Expected value %v, got %v", tt.expectedSatisfied.Get(i), el)
				}
			}

			for i, el := range notSatisfied.Values() {
				if el != tt.expectedNotSatisfied.Get(i) {
					t.Errorf("Expected value %v, got %v", tt.expectedNotSatisfied.Get(i), el)
				}
			}
		})
	}
}

func TestSafeSlice_SplitAtIndex(t *testing.T) {
	tests := []struct {
		name          string
		elements      []any
		index         int
		expectedLeft  *SafeSlice[any]
		expectedRight *SafeSlice[any]
	}{
		{"SplitAtIndexInt", []any{1, 2, 3, 4, 5}, 2,
			NewSafeSliceFromSlice([]any{1, 2}),
			NewSafeSliceFromSlice([]any{3, 4, 5}),
		},
		{"SplitAtIndexString", []any{"test1", "test2", "test3", "test4", "test5"}, 3,
			NewSafeSliceFromSlice([]any{"test1", "test2", "test3"}),
			NewSafeSliceFromSlice([]any{"test4", "test5"}),
		},
		{"SplitAtIndexFloat", []any{1.1, 2.2, 3.3, 4.4, 5.5}, 4,
			NewSafeSliceFromSlice([]any{1.1, 2.2, 3.3, 4.4}),
			NewSafeSliceFromSlice([]any{5.5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()
			for _, e := range tt.elements {
				s.Append(e)
			}
			left, right := s.SplitAtIndex(tt.index)

			for i, el := range left.Values() {
				if el != tt.expectedLeft.Get(i) {
					t.Errorf("Expected value %v, got %v", tt.expectedLeft.Get(i), el)
				}
			}

			for i, el := range right.Values() {
				if el != tt.expectedRight.Get(i) {
					t.Errorf("Expected value %v, got %v", tt.expectedRight.Get(i), el)
				}
			}
		})
	}
}

func TestSafeSlice_SortBy(t *testing.T) {
	tests := []struct {
		name           string
		elements       []any
		fn             func(any, any) bool
		expectedValues []any
	}{
		{"SortByInt", []any{3, 1, 2}, func(a, b any) bool { return a.(int) < b.(int) }, []any{1, 2, 3}},
		{"SortByString", []any{"c", "b", "a"}, func(a, b any) bool { return a.(string) < b.(string) }, []any{"a", "b", "c"}},
		{"SortByFloat", []any{3.14, 1.618, 2.718}, func(a, b any) bool { return a.(float64) < b.(float64) }, []any{1.618, 2.718, 3.14}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSlice[any]()
			for _, e := range tt.elements {
				s.Append(e)
			}

			s.SortBy(tt.fn)

			// compare s and tt.expectedValues
			if s.Len() != len(tt.expectedValues) {
				t.Errorf("Length of sorted slice is incorrect")
			}

			for i, v := range s.Values() {
				if v != tt.expectedValues[i] {
					t.Errorf("Value at index %d is incorrect", i)
				}
			}
		})
	}
}

func TestSafeSlice_Copy(t *testing.T) {
	tests := []struct {
		name     string
		elements []any
	}{
		{"CopyInt", []any{10, 20, 30}},
		{"CopyString", []any{"test1", "test2", "test3"}},
		{"CopyFloat", []any{3.14, 2.71, 1.618}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSafeSliceFromSlice(tt.elements)
			c := s.Copy()

			// Check if the copied slice has the same values as the original slice
			if !reflect.DeepEqual(c.Values(), s.Values()) {
				t.Errorf("Copy() = %v, want %v", c.Values(), s.Values())
			}

			// Modify the copied slice and check if the original slice remains unchanged
			c.Push(100)
			if reflect.DeepEqual(s.Values(), c.Values()) {
				t.Errorf("Original slice was modified after copying")
			}
		})
	}
}
