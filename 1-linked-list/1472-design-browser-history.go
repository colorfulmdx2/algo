package linked_list

type BrowserHistoryItem struct {
	Url  string
	Next *BrowserHistoryItem
	Prev *BrowserHistoryItem
}

type BrowserHistory struct {
	Cur *BrowserHistoryItem
}

func Constructor(homepage string) BrowserHistory {
	item := &BrowserHistoryItem{homepage, nil, nil}
	return BrowserHistory{item}
}

func (r *BrowserHistory) Visit(url string) {
	newPage := &BrowserHistoryItem{url, nil, r.Cur}
	r.Cur.Next = newPage
	r.Cur = newPage
}

func (r *BrowserHistory) Back(steps int) string {
	for steps > 0 && r.Cur.Prev != nil {
		r.Cur = r.Cur.Prev
		steps--
	}
	return r.Cur.Url
}

func (r *BrowserHistory) Forward(steps int) string {
	for steps > 0 && r.Cur.Next != nil {
		r.Cur = r.Cur.Next
		steps--
	}
	return r.Cur.Url
}
