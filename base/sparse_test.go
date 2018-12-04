package base

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSparseIdSet(t *testing.T) {
	// Create a ID set
	set := MakeSparseIdSet()
	assert.Equal(t, set.Len(), 0)
	// Add IDs
	set.Add(1)
	set.Add(2)
	set.Add(4)
	set.Add(8)
	assert.Equal(t, 4, set.Len())
	assert.Equal(t, 0, set.ToDenseId(1))
	assert.Equal(t, 1, set.ToDenseId(2))
	assert.Equal(t, 2, set.ToDenseId(4))
	assert.Equal(t, 3, set.ToDenseId(8))
	assert.Equal(t, NotId, set.ToDenseId(1000))
	assert.Equal(t, 1, set.ToSparseId(0))
	assert.Equal(t, 2, set.ToSparseId(1))
	assert.Equal(t, 4, set.ToSparseId(2))
	assert.Equal(t, 8, set.ToSparseId(3))
}

func TestSparseVector(t *testing.T) {
	vec := NewSparseVector()
	// Add new items
	vec.Add(2, 1)
	vec.Add(0, 0)
	vec.Add(8, 3)
	vec.Add(4, 2)
	assert.Equal(t, []int{2, 0, 8, 4}, vec.Indices)
	assert.Equal(t, []float64{1, 0, 3, 2}, vec.Values)
	// Sort indices
	sort.Sort(vec)
	assert.Equal(t, []int{0, 2, 4, 8}, vec.Indices)
	assert.Equal(t, []float64{0, 1, 2, 3}, vec.Values)
}

func TestSparseVector_ForIntersection(t *testing.T) {
	a := NewSparseVector()
	a.Add(2, 1)
	a.Add(1, 0)
	a.Add(8, 3)
	a.Add(4, 2)
	b := NewSparseVector()
	b.Add(16, 2)
	b.Add(1, 0)
	b.Add(64, 3)
	b.Add(4, 1)
	intersectIndex := make([]int, 0)
	intersectA := make([]float64, 0)
	intersectB := make([]float64, 0)
	a.ForIntersection(b, func(index int, a, b float64) {
		intersectIndex = append(intersectIndex, index)
		intersectA = append(intersectA, a)
		intersectB = append(intersectB, b)
	})
	assert.Equal(t, []int{1, 4}, intersectIndex)
	assert.Equal(t, []float64{0, 2}, intersectA)
	assert.Equal(t, []float64{0, 1}, intersectB)
}

func TestAdjacentVector(t *testing.T) {

}
