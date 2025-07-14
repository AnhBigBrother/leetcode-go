package main

type SnapshotArrayItem struct {
	value int
	snap  int
}

type SnapshotArray struct {
	arr        []*([]*SnapshotArrayItem)
	snap_count int
}

func Constructor(length int) SnapshotArray {
	arr := []*([]*SnapshotArrayItem){}
	for range length {
		arr = append(arr, &([]*SnapshotArrayItem{{
			value: 0,
			snap:  0,
		}}))
	}
	return SnapshotArray{
		arr:        arr,
		snap_count: 0,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	if index >= len(this.arr) {
		return
	}
	vals := this.arr[index]
	if (*vals)[len(*vals)-1].snap == this.snap_count {
		(*vals)[len(*vals)-1].value = val
	} else {
		*vals = append(*vals, &SnapshotArrayItem{
			value: val,
			snap:  this.snap_count,
		})
	}
}

func (this *SnapshotArray) Snap() int {
	this.snap_count++
	return this.snap_count - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	if snap_id >= this.snap_count {
		return -1
	}
	vals := *(this.arr[index])
	l, r := 0, len(vals)-1
	for l < r {
		m := (l + r + 1) / 2
		if vals[m].snap > snap_id {
			r = m - 1
		} else {
			l = m
		}
	}
	return vals[l].value
}
