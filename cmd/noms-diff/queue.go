package main

import (
	"container/list"

	"github.com/attic-labs/noms/go/types"
)

type diffInfo struct {
	path types.Path
	key  types.Value
	v1   types.Value
	v2   types.Value
}

type diffQueue struct {
	*list.List
}

func (q *diffQueue) PopFront() (diffInfo, bool) {
	el := q.Front()
	if el == nil {
		return diffInfo{}, false
	}
	q.Remove(el)
	return el.Value.(diffInfo), true
}

func NewDiffQueue() *diffQueue {
	return &diffQueue{list.New()}
}
