package gogenericssyncmap

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRWSyncMap(t *testing.T) {
	var rwmap, rwmap2 RWSyncMap[int, string]
	rwmap.Init()
	rwmap2.Init()

	rwmap.Store(1, "one")
	require.Equal(t, 1, rwmap.Len())
	val, ok := rwmap.Load(1)
	require.Equal(t, "one", val)
	require.Equal(t, true, ok)

	rwmap2.Store(2, "two")
	rwmap2.Store(3, "three")
	require.Equal(t, 2, rwmap2.Len())
	val, ok = rwmap2.Load(2)
	require.Equal(t, "two", val)
	require.Equal(t, true, ok)
}

// BenchmarkRWSyncMap-8    18423238                59.31 ns/op            0 B/op          0 allocs/op
func BenchmarkRWSyncMap(b *testing.B) {
	b.ReportAllocs()

	var (
		m         = RWSyncMap[string, int]{}
		v, ok any = 0, false
	)

	m.Init()

	const key = "a"
	for i := 0; i < b.N; i++ {
		m.Store(key, 1)
		v, ok = m.Load(key)
		m.Delete(key)
	}

	require.True(b, ok.(bool))
	require.EqualValues(b, 1, v)
}

func BenchmarkSyncMap(b *testing.B) {
	b.ReportAllocs()
	var (
		m         = sync.Map{}
		v, ok any = 0, false
	)

	const key = "a"

	for i := 0; i < b.N; i++ {
		m.Store(key, 1)
		v, ok = m.Load(key)
		m.Delete(key)
	}

	require.True(b, ok.(bool))
	require.EqualValues(b, 1, v)
}
