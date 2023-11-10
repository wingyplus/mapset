package mapset

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSet(t *testing.T) {
	ms := NewFromSlice([]string{"apple", "orange", "orange"})
	slice := ms.ToSlice()
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	assert.Equal(t, []string{"apple", "orange"}, slice)
}

func TestMapSet_Put(t *testing.T) {
	ms := NewFromSlice([]string{})
	ms.Put("apple")
	ms.Put("orange")
	ms.Put("orange")
	slice := ms.ToSlice()
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	assert.Equal(t, []string{"apple", "orange"}, slice)
}

func TestMapSet_Len(t *testing.T) {
	ms := NewFromSlice([]string{"apple", "orange", "orange"})
	assert.Equal(t, uint(2), ms.Len())

	ms = NewFromSlice([]string{})
	ms.Put("apple")
	ms.Put("orange")
	ms.Put("orange")
	assert.Equal(t, uint(2), ms.Len())
}

func TestMapSet_HasMember(t *testing.T) {
	ms := NewFromSlice([]string{"apple", "orange", "orange"})
	assert.True(t, ms.HasMember("apple"))
	assert.True(t, ms.HasMember("orange"))
	assert.False(t, ms.HasMember("mango"))
}

func TestMapSet_Intersection(t *testing.T) {
	ms := NewFromSlice([]string{"apple", "orange"})
	ms2 := NewFromSlice([]string{"apple"})
	assert.Equal(t, []string{"apple"}, ms.Intersection(ms2).ToSlice())
}

func TestMapSet_Union(t *testing.T) {
	ms := NewFromSlice([]string{"apple", "orange"})
	ms2 := NewFromSlice([]string{"apple", "mango"})
	slice := ms.Union(ms2).ToSlice()
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	assert.Equal(t, []string{"apple", "mango", "orange"}, slice)
}
