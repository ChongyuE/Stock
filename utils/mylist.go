package utils

import (
	"container/list"
	"sync"
)

type Mylist struct {
	ml *list.List
	sm sync.Mutex
}

func (l *Mylist) Init() *Mylist {
	l.ml = list.New()
	return l
}
func NewMylist() *Mylist {
	return new(Mylist).Init()
}

func (l *Mylist) PushBack(v interface{}) {
	l.sm.Lock()
	l.ml.PushBack(v)
	l.sm.Unlock()
}
func (l *Mylist) Back() (re *list.Element) {
	if l.ml.Len() == 0 {
		return nil
	} else {
		a := l.ml.Back()
		return a
	}
}
func (l *Mylist) Remove(re *list.Element) {
	l.ml.Remove(re)
}

func (l *Mylist) GetFrontRemove() (re *list.Element) {
	l.sm.Lock()
	if l.ml.Len() == 0 {
		l.sm.Unlock()
		return nil
	} else {
		a := l.ml.Front()
		l.ml.Remove(a)
		l.sm.Unlock()
		return a
	}
}
func (l *Mylist) Len() int {
	l.sm.Lock()
	i := l.ml.Len()
	l.sm.Unlock()
	return i
}
