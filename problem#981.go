package main

type TimeMap_Item981 struct {
	value     string
	timestamp int
}

type TimeMap struct {
	TMap map[string]([]*TimeMap_Item981)
	Size int
}

func NewTimeMap() TimeMap {
	return TimeMap{
		TMap: map[string]([]*TimeMap_Item981){},
		Size: 0,
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.TMap[key] == nil {
		this.TMap[key] = []*TimeMap_Item981{}
	}
	this.TMap[key] = append(this.TMap[key], &TimeMap_Item981{
		value:     value,
		timestamp: timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	key_data := this.TMap[key]
	if key_data == nil {
		return ""
	}
	l, r := 0, len(key_data)-1
	for l < r {
		m := (l + r + 1) / 2
		if key_data[m].timestamp > timestamp {
			r = m - 1
		} else {
			l = m
		}
	}
	if key_data[l].timestamp > timestamp {
		return ""
	}
	return key_data[l].value
}
