package utils

type DisjointSet map[int]int

func NewDisjointSet() DisjointSet {
	return DisjointSet{}
}

func (ds DisjointSet) Find(x int) int {
	parent, ok := ds[x]
	if !ok || parent == x {
		return x
	}
	return ds.Find(parent)
}

func (ds DisjointSet) Union(x, y int) {
	px := ds.Find(x)
	py := ds.Find(y)
	if px < py {
		ds[py] = px
	} else {
		ds[px] = py
	}
}

func (ds DisjointSet) Check(x, y int) bool {
	return ds.Find(x) == ds.Find(y)
}
