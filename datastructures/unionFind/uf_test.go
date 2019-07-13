package unionFind

import "testing"
import "github.com/stretchr/testify/assert"

func TestUnionQuickFind(t *testing.T) {
	uf := NewUnionFind(10)
	assert.Equal(t, false, uf.QuickFind(2,3))
	uf.Union(1,3)
	uf.Union(4,5)
	uf.Union(5,1)
	assert.Equal(t, true, uf.QuickFind(5,3))
	assert.Equal(t, true, uf.QuickFind(4,1))
}

func TestQuickUnionFind(t *testing.T) {
	uf := NewUnionFind(10)
	assert.Equal(t, false, uf.Find(2,3))
	uf.QuickUnion(1,3)
	uf.QuickUnion(4,5)
	uf.QuickUnion(5,1)
	assert.Equal(t, true, uf.Find(5,3))
	assert.Equal(t, true, uf.Find(4,1))
}


