package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	length    int
	firstElem *ListItem
	lastElem  *ListItem
}

func (l list) Len() int {
	return l.length
}

func (l list) Front() *ListItem {
	return l.firstElem
}

func (l list) Back() *ListItem {
	return l.lastElem
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.length == 0 {
		l.length++
		li := &ListItem{v, nil, nil}
		l.firstElem = li
		l.lastElem = li
		return li
	}
	l.length++
	li := &ListItem{v, l.firstElem, nil}
	l.firstElem.Prev = li
	l.firstElem = li
	return li
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.length == 0 {
		l.length++
		li := &ListItem{v, nil, nil}
		l.firstElem = li
		l.lastElem = li
		return li
	}
	l.length++
	li := &ListItem{v, nil, l.lastElem}
	l.lastElem.Next = li
	l.lastElem = li
	return li
}
func (l *list) Remove(i *ListItem) {
	l.length--
	if i.Prev == nil {
		l.firstElem = i.Next
		i.Next.Prev = i.Prev
		return
	} else if i.Next == nil {
		l.lastElem = i.Prev
		i.Prev.Next = i.Next
		return
	}
	i.Next.Prev = i.Prev
	i.Prev.Next = i.Next
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	} else if i.Next == nil {
		i.Prev.Next = i.Next
		l.lastElem = i.Prev
		i.Next = l.firstElem
		i.Prev = nil
		l.firstElem.Prev = i
		l.firstElem = i
		return
	}
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
	i.Prev = nil
	i.Next = l.firstElem
	l.firstElem.Prev = i
	l.firstElem = i
}

func NewList() List {
	return new(list)
}
