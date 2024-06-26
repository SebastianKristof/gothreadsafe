package safeset

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSet_NewSet(t *testing.T) {
	s := NewSet[int]()
	require.NotNil(t, s)
	require.NotNil(t, s.items)
	require.Equal(t, 0, s.Size())
}

func TestSet_NewSetWithValues(t *testing.T) {
	s := NewSetWithValues[int](1, 2, 3)
	require.NotNil(t, s)
	require.NotNil(t, s.items)
	require.Equal(t, 3, s.Size())
	require.True(t, s.Contains(1))
	require.True(t, s.Contains(2))
	require.True(t, s.Contains(3))
}

func TestSet_AddInt(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	require.True(t, s.Contains(1))
	require.Equal(t, 1, s.Size())

	s.Add(2)
	require.True(t, s.Contains(1))
	require.True(t, s.Contains(2))
	require.Equal(t, 2, s.Size())

	s.Add(3)
	require.True(t, s.Contains(1))
	require.True(t, s.Contains(2))
	require.True(t, s.Contains(3))
	require.Equal(t, 3, s.Size())
}

func TestSet_AddString(t *testing.T) {
	s := NewSet[string]()
	s.Add("one")
	require.True(t, s.Contains("one"))
	require.Equal(t, 1, s.Size())

	s.Add("two")
	require.True(t, s.Contains("one"))
	require.True(t, s.Contains("two"))
	require.Equal(t, 2, s.Size())

	s.Add("three")
	require.True(t, s.Contains("one"))
	require.True(t, s.Contains("two"))
	require.True(t, s.Contains("three"))
	require.Equal(t, 3, s.Size())
}

func TestSet_AddStruct(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	s := NewSet[person]()
	s.Add(person{"Alice", 25})
	require.True(t, s.Contains(person{"Alice", 25}))
	require.Equal(t, 1, s.Size())

	s.Add(person{"Bob", 30})
	require.True(t, s.Contains(person{"Alice", 25}))
	require.True(t, s.Contains(person{"Bob", 30}))
	require.Equal(t, 2, s.Size())

	s.Add(person{"Charlie", 35})
	require.True(t, s.Contains(person{"Alice", 25}))
	require.True(t, s.Contains(person{"Bob", 30}))
	require.True(t, s.Contains(person{"Charlie", 35}))
	require.Equal(t, 3, s.Size())
}

func TestSet_AddWithCheckInt(t *testing.T) {
	s := NewSet[int]()
	require.False(t, s.AddWithCheck(1))
	require.True(t, s.Contains(1))
	require.Equal(t, 1, s.Size())

	require.True(t, s.AddWithCheck(1))
	require.True(t, s.Contains(1))
	require.Equal(t, 1, s.Size())
}

func TestSet_RemoveInt(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	s.Remove(2)
	require.True(t, s.Contains(1))
	require.False(t, s.Contains(2))
	require.True(t, s.Contains(3))
	require.Equal(t, 2, s.Size())
}

func TestSet_RemoveString(t *testing.T) {
	s := NewSet[string]()
	s.Add("one")
	s.Add("two")
	s.Add("three")

	s.Remove("two")
	require.True(t, s.Contains("one"))
	require.False(t, s.Contains("two"))
	require.True(t, s.Contains("three"))
	require.Equal(t, 2, s.Size())
}

func TestSet_RemoveStruct(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	s := NewSet[person]()
	s.Add(person{"Alice", 25})
	s.Add(person{"Bob", 30})
	s.Add(person{"Charlie", 35})

	s.Remove(person{"Bob", 30})
	require.True(t, s.Contains(person{"Alice", 25}))
	require.False(t, s.Contains(person{"Bob", 30}))
	require.True(t, s.Contains(person{"Charlie", 35}))
	require.Equal(t, 2, s.Size())
}

func TestSet_ContainsInt(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	require.True(t, s.Contains(1))
	require.True(t, s.Contains(2))
	require.True(t, s.Contains(3))
	require.False(t, s.Contains(4))
}

func TestSet_ContainsString(t *testing.T) {
	s := NewSet[string]()
	s.Add("one")
	s.Add("two")
	s.Add("three")

	require.True(t, s.Contains("one"))
	require.True(t, s.Contains("two"))
	require.True(t, s.Contains("three"))
	require.False(t, s.Contains("four"))
}

func TestSet_ContainsStruct(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	s := NewSet[person]()
	s.Add(person{"Alice", 25})
	s.Add(person{"Bob", 30})
	s.Add(person{"Charlie", 35})

	require.True(t, s.Contains(person{"Alice", 25}))
	require.True(t, s.Contains(person{"Bob", 30}))
	require.True(t, s.Contains(person{"Charlie", 35}))
	require.False(t, s.Contains(person{"David", 40}))
	require.False(t, s.Contains(person{"Alice", 99}))
}

func TestSet_Size(t *testing.T) {
	s := NewSet[int]()
	require.Equal(t, 0, s.Size())

	s.Add(1)
	require.Equal(t, 1, s.Size())

	s.Add(2)
	require.Equal(t, 2, s.Size())

	s.Add(3)
	require.Equal(t, 3, s.Size())

	s.Remove(2)
	require.Equal(t, 2, s.Size())
}

func TestSet_Clear(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	require.Equal(t, 3, s.Size())

	s.Clear()
	require.Equal(t, 0, s.Size())
}

func TestSet_ToSlice(t *testing.T) {
	s := NewSet[int]()
	s.Add(3)
	s.Add(1)
	s.Add(2)

	slice := s.ToSlice()
	require.ElementsMatch(t, []int{1, 2, 3}, slice)
}

func TestSet_Union(t *testing.T) {
	s1 := NewSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := NewSet[int]()
	s2.Add(3)
	s2.Add(4)
	s2.Add(5)

	union := s1.Union(s2)
	require.ElementsMatch(t, []int{1, 2, 3, 4, 5}, union.ToSlice())
}

func TestSet_UnionEmpty(t *testing.T) {
	s1 := NewSet[int]()
	s2 := NewSet[int]()

	union := s1.Union(s2)
	require.Equal(t, 0, union.Size())
}

func TestSet_UnionSame(t *testing.T) {
	s1 := NewSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	union := s1.Union(s1)
	require.ElementsMatch(t, []int{1, 2, 3}, union.ToSlice())
}

func TestSet_UnionEqual(t *testing.T) {
	s1 := NewSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := NewSet[int]()
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)

	union := s1.Union(s2)
	require.ElementsMatch(t, []int{1, 2, 3}, union.ToSlice())
	require.True(t, s1.Equal(union))
	require.True(t, union.Equal(s1))
}

func TestSet_Intersection(t *testing.T) {
	type testCase struct {
		name         string
		set1         *Set[int]
		set2         *Set[int]
		expected     []int
		expectedSize int
	}

	set1 := NewSetWithValues[int](1, 2, 3)
	set2 := NewSetWithValues[int](3, 4, 5)

	testCases := []testCase{
		{
			name:         "Intersection of non-empty sets",
			set1:         set1,
			set2:         set2,
			expected:     []int{3},
			expectedSize: 1,
		},
		{
			name:         "Intersection of empty sets",
			set1:         NewSet[int](),
			set2:         NewSet[int](),
			expected:     []int{},
			expectedSize: 0,
		},
		{
			name:         "Intersection of same set",
			set1:         set1,
			set2:         set1,
			expected:     []int{1, 2, 3},
			expectedSize: 3,
		},
		{
			name:         "Intersection of equal sets",
			set1:         set1,
			set2:         set1.Clone(),
			expected:     []int{1, 2, 3},
			expectedSize: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			intersection := tc.set1.Intersection(tc.set2)
			require.ElementsMatch(t, tc.expected, intersection.ToSlice())
			require.Equal(t, tc.expectedSize, intersection.Size())
		})
	}
}

func TestSet_Difference(t *testing.T) {
	type testCase struct {
		name         string
		set1         *Set[int]
		set2         *Set[int]
		expected     []int
		expectedSize int
	}

	set1 := NewSetWithValues[int](1, 2, 3)
	set2 := NewSetWithValues[int](3, 4, 5)

	testCases := []testCase{
		{
			name:         "Difference of non-empty sets",
			set1:         set1,
			set2:         set2,
			expected:     []int{1, 2},
			expectedSize: 2,
		},
		{
			name:         "Difference of non-empty sets (reversed)",
			set1:         set2,
			set2:         set1,
			expected:     []int{4, 5},
			expectedSize: 2,
		},
		{
			name:         "Difference of empty sets",
			set1:         NewSet[int](),
			set2:         NewSet[int](),
			expected:     []int{},
			expectedSize: 0,
		},
		{
			name:         "Difference of same set",
			set1:         set1,
			set2:         set1,
			expected:     []int{},
			expectedSize: 0,
		},
		{
			name:         "Difference of equal sets",
			set1:         set1,
			set2:         set1.Clone(),
			expected:     []int{},
			expectedSize: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			difference := tc.set1.Difference(tc.set2)
			require.ElementsMatch(t, tc.expected, difference.ToSlice())
			require.Equal(t, tc.expectedSize, difference.Size())
		})
	}
}

func TestSet_SymmetricDifference(t *testing.T) {
	type testCase struct {
		name         string
		set1         *Set[int]
		set2         *Set[int]
		expected     []int
		expectedSize int
	}

	set1 := NewSetWithValues[int](1, 2, 3)
	set2 := NewSetWithValues[int](3, 4, 5)

	testCases := []testCase{
		{
			name:         "Symmetric difference of non-empty sets",
			set1:         set1,
			set2:         set2,
			expected:     []int{1, 2, 4, 5},
			expectedSize: 4,
		},
		{
			name:         "Symmetric difference of non-empty sets (reversed)",
			set1:         set2,
			set2:         set1,
			expected:     []int{1, 2, 4, 5},
			expectedSize: 4,
		},
		{
			name:         "Symmetric difference of empty sets",
			set1:         NewSet[int](),
			set2:         NewSet[int](),
			expected:     []int{},
			expectedSize: 0,
		},
		{
			name:         "Symmetric difference of same set",
			set1:         set1,
			set2:         set1,
			expected:     []int{},
			expectedSize: 0,
		},
		{
			name:         "Symmetric difference of equal sets",
			set1:         set1,
			set2:         set1.Clone(),
			expected:     []int{},
			expectedSize: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			symmetricDifference := tc.set1.SymmetricDifference(tc.set2)
			require.ElementsMatch(t, tc.expected, symmetricDifference.ToSlice())
			require.Equal(t, tc.expectedSize, symmetricDifference.Size())
		})
	}
}

func TestSet_IsSubsetOf(t *testing.T) {
	type testCase struct {
		name     string
		set1     *Set[int]
		set2     *Set[int]
		expected bool
	}

	set1 := NewSetWithValues[int](1, 2, 3)
	set2 := NewSetWithValues[int](1, 2, 3, 4, 5)

	testCases := []testCase{
		{
			name:     "Subset of non-empty sets",
			set1:     set1,
			set2:     set2,
			expected: true,
		},
		{
			name:     "Subset of non-empty sets (reversed)",
			set1:     set2,
			set2:     set1,
			expected: false,
		},
		{
			name:     "Subset of empty sets",
			set1:     NewSet[int](),
			set2:     NewSet[int](),
			expected: true,
		},
		{
			name:     "Subset of same set",
			set1:     set1,
			set2:     set1,
			expected: true,
		},
		{
			name:     "Subset of equal sets",
			set1:     set1,
			set2:     set1.Clone(),
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.set1.IsSubsetOf(tc.set2))
		})
	}
}

func TestSet_IsSupersetOf(t *testing.T) {
	type testCase struct {
		name     string
		set1     *Set[int]
		set2     *Set[int]
		expected bool
	}

	set1 := NewSetWithValues[int](1, 2, 3, 4, 5)
	set2 := NewSetWithValues[int](1, 2, 3)

	testCases := []testCase{
		{
			name:     "Superset of non-empty sets",
			set1:     set1,
			set2:     set2,
			expected: true,
		},
		{
			name:     "Superset of non-empty sets (reversed)",
			set1:     set2,
			set2:     set1,
			expected: false,
		},
		{
			name:     "Superset of empty sets",
			set1:     NewSet[int](),
			set2:     NewSet[int](),
			expected: true,
		},
		{
			name:     "Superset of same set",
			set1:     set1,
			set2:     set1,
			expected: true,
		},
		{
			name:     "Superset of equal sets",
			set1:     set1,
			set2:     set1.Clone(),
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.set1.IsSupersetOf(tc.set2))
		})
	}
}

func TestSet_Equal(t *testing.T) {
	type testCase struct {
		name     string
		set1     *Set[int]
		set2     *Set[int]
		expected bool
	}

	set1 := NewSetWithValues[int](1, 2, 3)
	set2 := NewSetWithValues[int](1, 2, 3)

	testCases := []testCase{
		{
			name:     "Equal sets",
			set1:     set1,
			set2:     set2,
			expected: true,
		},
		{
			name:     "Equal sets (reversed)",
			set1:     set2,
			set2:     set1,
			expected: true,
		},
		{
			name:     "Equal empty sets",
			set1:     NewSet[int](),
			set2:     NewSet[int](),
			expected: true,
		},
		{
			name:     "Equal same set",
			set1:     set1,
			set2:     set1,
			expected: true,
		},
		{
			name:     "Equal equal sets",
			set1:     set1,
			set2:     set1.Clone(),
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.set1.Equal(tc.set2))
		})
	}
}

func TestSet_Clone(t *testing.T) {
	s := NewSetWithValues[int](1, 2, 3)
	clone := s.Clone()

	require.True(t, s.Equal(clone))
	require.True(t, clone.Equal(s))
}

func TestSet_StringInt(t *testing.T) {
	s := NewSetWithValues[int](1, 2, 3)

	require.Contains(t, s.String(), "1")
	require.Contains(t, s.String(), "2")
	require.Contains(t, s.String(), "3")
}

func TestSet_StringEmpty(t *testing.T) {
	s := NewSet[int]()
	require.Equal(t, "{}", s.String())
}

func TestSet_StringString(t *testing.T) {
	s := NewSetWithValues[string]("one", "two", "three")

	require.Contains(t, s.String(), "one")
	require.Contains(t, s.String(), "two")
	require.Contains(t, s.String(), "three")
}

func TestSet_StringStruct(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	s := NewSet[person]()
	s.Add(person{"Alice", 25})
	s.Add(person{"Bob", 30})
	s.Add(person{"Charlie", 35})

	require.Contains(t, s.String(), "{Alice 25}")
	require.Contains(t, s.String(), "{Bob 30}")
	require.Contains(t, s.String(), "{Charlie 35}")
}

func TestSet_IsEmpty(t *testing.T) {
	s := NewSet[int]()
	require.True(t, s.IsEmpty())

	s.Add(1)
	require.False(t, s.IsEmpty())
}

func TestSet_IsNotEmpty(t *testing.T) {
	s := NewSet[int]()
	require.False(t, !s.IsEmpty())

	s.Add(1)
	require.True(t, !s.IsEmpty())
}
