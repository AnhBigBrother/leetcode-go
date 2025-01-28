package main

import "container/heap"

type TaskItem struct {
	UserId   int
	TaskId   int
	Priority int
	Removed  bool
	index    int
}

type TaskItemArr []*TaskItem

func (pq TaskItemArr) Len() int { return len(pq) }

func (pq TaskItemArr) Less(i, j int) bool {
	if pq[i].Priority == pq[j].Priority {
		return pq[i].TaskId > pq[j].TaskId
	}
	return pq[i].Priority > pq[j].Priority
}

func (pq TaskItemArr) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *TaskItemArr) Push(x any) {
	n := len(*pq)
	item := x.(*TaskItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *TaskItemArr) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

type TaskManager struct {
	pq    TaskItemArr
	idMap map[int]*TaskItem
}

func TaskManagerConstructor(tasks [][]int) TaskManager {
	h := TaskItemArr{}
	idMap := map[int]*TaskItem{}
	for i, t := range tasks {
		h = append(h, &TaskItem{
			UserId:   t[0],
			TaskId:   t[1],
			Priority: t[2],
			index:    i,
		})
		idMap[t[1]] = h[i]
	}
	heap.Init(&h)
	return TaskManager{
		pq:    h,
		idMap: idMap,
	}
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	item := TaskItem{
		UserId:   userId,
		TaskId:   taskId,
		Priority: priority,
		index:    len(this.pq),
	}
	heap.Push(&this.pq, &item)
	this.idMap[taskId] = &item
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	this.idMap[taskId].Priority = newPriority
	heap.Fix(&this.pq, this.idMap[taskId].index)
}

func (this *TaskManager) Rmv(taskId int) {
	this.idMap[taskId].Removed = true
}

func (this *TaskManager) ExecTop() int {
	if this.pq.Len() == 0 {
		return -1
	}
	x := heap.Pop(&this.pq).(*TaskItem)
	for x.Removed && this.pq.Len() > 0 {
		x = heap.Pop(&this.pq).(*TaskItem)
		// this.idMap[x.TaskId] = nil
	}
	if x.Removed {
		return -1
	}
	return x.UserId
}
