package utils

type DisjointSet []int

func (ds DisjointSet) Find(x int) int {
	if ds[x] != x {
		return ds.Find(ds[x])
	}
	return x
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

func NewDisjointSet(n int) DisjointSet {
	ds := DisjointSet{}
	for i := 0; i <= n; i++ {
		ds = append(ds, i)
	}
	return ds
}
