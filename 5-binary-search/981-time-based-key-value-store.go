package __binary_search

type TimeMap struct {
	store map[string][]*StoreElement
}

type StoreElement struct {
	Timestamp int
	Value     string
}

func Constructor() TimeMap {
	store := make(map[string][]*StoreElement)

	return TimeMap{store: store}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	el := &StoreElement{
		Timestamp: timestamp,
		Value:     value,
	}

	this.store[key] = append(this.store[key], el)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	slice := this.store[key]

	l, r := 0, len(slice)-1

	for l <= r {
		mid := l + (r-l)/2

		el := slice[mid]

		if el.Timestamp == timestamp {
			return el.Value
		}

		if el.Timestamp < timestamp {
			l = mid + 1
		}

		if el.Timestamp > timestamp {
			r = mid - 1
		}
	}

	if r >= 0 {
		return slice[r].Value
	}

	return ""
}
