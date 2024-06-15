package safemap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSafeMap(t *testing.T) {
	sm := NewSafeMap[int, string]()
	require.NotNil(t, sm)
	require.Empty(t, sm.GetMap())
}

func TestNewSafeMapFromMap(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	sm := NewSafeMapFromMap(m)
	require.NotNil(t, sm)
	require.Equal(t, m, sm.GetMap())
}

func TestNewSafeMapFromKeysValues(t *testing.T) {
	keys := []int{1, 2, 3}
	values := []string{"one", "two", "three"}
	sm, err := NewSafeMapFromKeysValues[int, string](keys, values)
	require.NoError(t, err)
	require.NotNil(t, sm)
	require.Equal(t, len(keys), sm.Len())
	for i := 0; i < len(keys); i++ {
		result := sm.Get(keys[i])
		require.True(t, result.Found)
		require.Equal(t, values[i], result.Value)
	}
}

func TestNewSafeMapFromKeyValuePairs(t *testing.T) {
	keysValues := []interface{}{
		1, "one", 2, "two", 3, "three",
	}
	sm, err := NewSafeMapFromKeyValuePairs[int, string](keysValues)
	require.NoError(t, err)
	require.NotNil(t, sm)
	require.Equal(t, len(keysValues)/2, sm.Len())
	for i := 0; i < len(keysValues); i += 2 {
		key := keysValues[i].(int)
		value := keysValues[i+1].(string)
		result := sm.Get(key)
		require.True(t, result.Found)
		require.Equal(t, value, result.Value)
	}
}

func TestSafeMap_Set(t *testing.T) {
	sm := NewSafeMap[int, string]()
	sm.Set(1, "one")
	require.Equal(t, 1, sm.Len())
	require.Equal(t, "one", sm.Get(1).Value)
}

func TestSafeMap_SetNX(t *testing.T) {
	sm := NewSafeMap[int, string]()
	sm.SetNX(1, "one")
	require.Equal(t, 1, sm.Len())
	sm.SetNX(1, "two")
	require.Equal(t, 1, sm.Len())
	require.Equal(t, "one", sm.Get(1).Value)
}

func TestSafeMap_Get(t *testing.T) {
	sm := NewSafeMap[int, string]()
	sm.Set(1, "one")
	result := sm.Get(1)
	require.True(t, result.Found)
	require.Equal(t, "one", result.Value)
}

func TestSafeMap_Delete(t *testing.T) {
	sm := NewSafeMap[int, string]()
	sm.Set(1, "one")
	require.Equal(t, 1, sm.Len())
	sm.Delete(1)
	require.Equal(t, 0, sm.Len())
}

func TestSafeMap_GetMap(t *testing.T) {
	sm := NewSafeMap[int, string]()
	m := sm.GetMap()
	require.NotNil(t, m)
	require.Empty(t, m)
}

func TestSafeMap_Len(t *testing.T) {
	sm := NewSafeMap[int, string]()
	sm.Set(1, "one")
	require.Equal(t, 1, sm.Len())
}

func TestSafeMap_Keys(t *testing.T) {
	sm := NewSafeMap[int, string]()
	sm.Set(1, "one")
	sm.Set(2, "two")
	sm.Set(3, "three")
	keys := sm.GetKeys()
	require.Len(t, keys, 3)
	require.Contains(t, keys, 1)
	require.Contains(t, keys, 2)
	require.Contains(t, keys, 3)
}

func TestSafeMap_Values(t *testing.T) {
	sm := NewSafeMap[int, string]()
	sm.Set(1, "one")
	sm.Set(2, "two")
	sm.Set(3, "three")
	values := sm.GetValues()
	require.Len(t, values, 3)
	require.Contains(t, values, "one")
	require.Contains(t, values, "two")
	require.Contains(t, values, "three")
}

func TestSafeMap_GetKeyValuePairs(t *testing.T) {
	sm := NewSafeMap[int, string]()
	sm.Set(1, "one")
	sm.Set(2, "two")
	sm.Set(3, "three")
	keysValues := sm.GetKeyValuePairs()
	require.Len(t, keysValues, 6)
	require.Contains(t, keysValues, 1)
	require.Contains(t, keysValues, "one")
	require.Contains(t, keysValues, 2)
	require.Contains(t, keysValues, "two")
	require.Contains(t, keysValues, 3)
	require.Contains(t, keysValues, "three")

	newSM, err := NewSafeMapFromKeyValuePairs[int, string](keysValues)
	require.NoError(t, err)
	require.Equal(t, sm.GetMap(), newSM.GetMap())
}

func TestSafeMap_GetKeysValues(t *testing.T) {
	sm := NewSafeMap[int, string]()
	sm.Set(1, "one")
	sm.Set(2, "two")
	sm.Set(3, "three")
	keys, values := sm.GetKeysValues()
	require.Len(t, keys, 3)
	require.Contains(t, keys, 1)
	require.Contains(t, keys, 2)
	require.Contains(t, keys, 3)
	require.Len(t, values, 3)
	require.Contains(t, values, "one")
	require.Contains(t, values, "two")
	require.Contains(t, values, "three")
}

func TestSafeMap_Copy(t *testing.T) {
	sm := NewSafeMap[int, string]()
	sm.Set(1, "one")
	sm.Set(2, "two")
	sm.Set(3, "three")
	newSM := sm.Copy()
	require.Equal(t, sm.GetMap(), newSM.GetMap())
}

func TestSafeMap_Clear(t *testing.T) {
	sm := NewSafeMap[int, string]()
	sm.Set(1, "one")
	sm.Set(2, "two")
	sm.Set(3, "three")
	require.NotEmpty(t, sm.GetMap())
	sm.Clear()
	require.Empty(t, sm.GetMap())
}

func TestSafeMap_IsEmpty(t *testing.T) {
	sm := NewSafeMap[int, string]()
	require.True(t, sm.IsEmpty())
	sm.Set(1, "one")
	require.False(t, sm.IsEmpty())
}

func TestSafeMap_String(t *testing.T) {
	sm := NewSafeMap[int, string]()
	sm.Set(1, "one")
	sm.Set(2, "two")
	sm.Set(3, "three")
	require.NotEmpty(t, sm.String())
	require.Contains(t, sm.String(), "1:one")
	require.Contains(t, sm.String(), "2:two")
	require.Contains(t, sm.String(), "3:three")
	require.Equal(t, "map[1:one 2:two 3:three]", sm.String())
}

func TestSafeMap_Export(t *testing.T) {
	sm := NewSafeMap[int, string]()
	sm.Set(1, "one")
	sm.Set(2, "two")
	sm.Set(3, "three")
	expM := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	m := sm.Export()
	require.Equal(t, expM, m)
}
