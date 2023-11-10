// Package mapset is a Set implementation using map.
package mapset

// MapSet is a Set implementation using map.
type MapSet[T comparable] map[T]struct{}

// Put add a new item to the MapSet.
func (ms MapSet[T]) Put(v T) {
	ms[v] = struct{}{}
}

// Len returns length of elements in MapSet.
func (ms MapSet[T]) Len() uint {
	return uint(len(ms))
}

// HasMember returns true if the member is in MapSet.
func (ms MapSet[T]) HasMember(v T) bool {
	_, ok := ms[v]
	return ok
}

// Intersection returns a new MapSet that contiains member that intersect
// between 2 MapSets.
func (ms MapSet[T]) Intersection(ms2 MapSet[T]) MapSet[T] {
	ms3 := make(MapSet[T])
	for v := range ms2 {
		if ms.HasMember(v) {
			ms3.Put(v)
		}
	}
	return ms3
}

// Union returns a new MapSet that combine all members between 2 MapSets.
func (ms MapSet[T]) Union(ms2 MapSet[T]) MapSet[T] {
	ms3 := make(MapSet[T])
	for v := range ms {
		ms3.Put(v)
	}
	for v := range ms2 {
		ms3.Put(v)
	}
	return ms3
}

// ToSlice converts a MapSet into un-ordered slice. 
func (ms MapSet[T]) ToSlice() (slice []T) {
	for v := range ms {
		slice = append(slice, v)
	}
	return
}

// NewFromSlice create a new MapSet from slice.
func NewFromSlice[T comparable](slice []T) (ms MapSet[T]) {
	ms = make(MapSet[T])
	for _, v := range slice {
		v := v
		ms.Put(v)
	}
	return
}
