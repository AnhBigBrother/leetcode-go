package main

type AllOne struct {
	count  map[string]int
	bucket map[int]map[string]bool
}

func AllOneConstructor() AllOne {
	return AllOne{
		count:  make(map[string]int),
		bucket: make(map[int]map[string]bool),
	}
}

func (this *AllOne) Inc(key string) {
	if this.count[key] == 0 {
		this.count[key] = 1
		if this.bucket[1] == nil {
			this.bucket[1] = make(map[string]bool)
		}
		this.bucket[1][key] = true
		return
	}
	x := this.count[key]
	delete(this.bucket[x], key)
	if len(this.bucket[x]) == 0 {
		delete(this.bucket, x)
	}
	this.count[key] = x + 1
	if this.bucket[x+1] == nil {
		this.bucket[x+1] = make(map[string]bool)
	}
	this.bucket[x+1][key] = true
}

func (this *AllOne) Dec(key string) {
	x := this.count[key]
	delete(this.bucket[x], key)
	if len(this.bucket[x]) == 0 {
		delete(this.bucket, x)
	}
	this.count[key] = x - 1
	if this.count[key] == 0 {
		return
	}
	if this.bucket[x-1] == nil {
		this.bucket[x-1] = make(map[string]bool)
	}
	this.bucket[x-1][key] = true
}

func (this *AllOne) GetMaxKey() string {
	max_count := 0
	for k := range this.bucket {
		max_count = max(max_count, k)
	}
	max_bucket := this.bucket[max_count]
	for k := range max_bucket {
		return k
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	min_count := 50001
	for k := range this.bucket {
		min_count = min(min_count, k)
	}
	min_bucket := this.bucket[min_count]
	for k := range min_bucket {
		return k
	}
	return ""
}
