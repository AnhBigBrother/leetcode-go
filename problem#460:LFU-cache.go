package main

type LFUCache struct {
	freq    map[int]int
	bucket  map[int]map[int]int
	max_len int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		freq:    make(map[int]int),
		bucket:  make(map[int]map[int]int),
		max_len: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if this.freq[key] == 0 {
		return -1
	}
	f := this.freq[key]
	res := this.bucket[f][key]
	delete(this.bucket[f], key)
	if len(this.bucket[f]) == 0 {
		delete(this.bucket, f)
	}
	this.freq[key] = f + 1
	if this.bucket[f+1] == nil {
		this.bucket[f+1] = make(map[int]int)
	}
	this.bucket[f+1][key] = res
	return res
}

func (this *LFUCache) Put(key int, value int) {
	if len(this.freq) < this.max_len {
		if this.freq[key] == 0 {
			this.freq[key] = 1
			if this.bucket[1] == nil {
				this.bucket[1] = make(map[int]int)
			}
			this.bucket[1][key] = value
		} else {
			f := this.freq[key]
			delete(this.bucket[f], key)
			if len(this.bucket[f]) == 0 {
				delete(this.bucket, f)
			}
			this.freq[key] = f + 1
			if this.bucket[f+1] == nil {
				this.bucket[f+1] = make(map[int]int)
			}
			this.bucket[f+1][key] = value
		}
	} else {
		if this.freq[key] == 0 {
			this.freq[key] = 1
			if this.bucket[1] == nil {
				min_freq := 50001
				for k := range this.bucket {
					min_freq = min(min_freq, k)
				}
				for k := range this.bucket[min_freq] {
					delete(this.bucket[min_freq], k)
					delete(this.freq, k)
					break
				}
				if len(this.bucket[min_freq]) == 0 {
					delete(this.bucket, min_freq)
				}
				this.bucket[1] = make(map[int]int)
			} else {
				for k := range this.bucket[1] {
					delete(this.bucket[1], k)
					delete(this.freq, k)
					break
				}
			}
			this.bucket[1][key] = value
		} else {
			f := this.freq[key]
			delete(this.bucket[f], key)
			if len(this.bucket[f]) == 0 {
				delete(this.bucket, f)
			}
			this.freq[key] = f + 1
			if this.bucket[f+1] == nil {
				this.bucket[f+1] = make(map[int]int)
			}
			this.bucket[f+1][key] = value
		}
	}
}
