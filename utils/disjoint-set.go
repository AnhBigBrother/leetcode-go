package utils

type DisjointSet []int

func NewDisjointSet(n int) DisjointSet {
	ds := []int{}
	for i := range n {
		ds = append(ds, i)
	}
	return ds
}

func (ds DisjointSet) Find(x int) int {
	parent := ds[x]
	if parent == x {
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
