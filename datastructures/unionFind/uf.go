package unionFind

type UF struct {
	id []int
	sz []int
}

func NewUnionFind(n int) *UF {
	id := make([]int, n)
	sz := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
		sz[i] = 1
	}

	return &UF{id: id, sz: sz}
}

func (u *UF) Union(p, q int) {
	pid := u.id[p]
	qid := u.id[q]
	for i, val := range u.id {
		if val == pid {
			u.id[i] = qid
		}
	}
}

func (u *UF) QuickFind(p, q int) bool {
	return u.id[p] == u.id[q]
}

func (u *UF) root(p int) int {
	i := p
	for u.id[i] != i {
		//u.id[i] = u.id[id[i]] //path compression
		i = u.id[i]
	}
	return i
}

func (u *UF) Find(p, q int) bool {
	return u.root(p) == u.root(q)
}

func (u *UF) QuickUnion(p, q int) {
	pRoot := u.root(p)
	qRoot := u.root(q)
	pSize := u.sz[p]
	qSize := u.sz[q]
	if qSize < pSize {
		u.id[qRoot] = pRoot
		u.sz[pRoot] += u.sz[qRoot]
	} else {
		u.id[pRoot] = qRoot
		u.sz[qRoot] += u.sz[pRoot]
	}
}
