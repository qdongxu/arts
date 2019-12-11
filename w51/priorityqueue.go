package w51

import (
	"sync"
)

const (
	initCapacity = 8
)

type (
	Data struct {
		Id    int32
		Value interface{}
	}

	PriorityQueue struct {
		array []*Data
		size  int // point to the next pos after the last element

		lock      *sync.Mutex
		emptyCond *sync.Cond
	}
)

func MustNewPriorityQueue() *PriorityQueue {
	rel := &PriorityQueue{array: make([]*Data, initCapacity), lock: &sync.Mutex{}}
	rel.emptyCond = sync.NewCond(rel.lock)
	return rel
}

func (q *PriorityQueue) Take() *Data {
	q.lock.Lock()
	defer q.lock.Unlock()

	for q.size <= 0 {
		q.emptyCond.Wait()
	}

	rel := q.array[0]
	if q.size > 1 {
		q.array[0] = q.array[q.size-1]
		q.array[q.size-1] = nil
	} else {
		q.array[0] = nil
	}

	q.size--

	q.rebuildHeapFromTop(0)

	return rel
}

func (q *PriorityQueue) Push(data *Data) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.size >= len(q.array) {
		newarray := make([]*Data, len(q.array)*2)
		for i, d := range q.array {
			newarray[i] = d
		}
		q.array = newarray
	}

	q.array[q.size] = data
	q.size++

	q.rebuildHeapFromBottom()

	q.emptyCond.Broadcast()
}

func (q *PriorityQueue) rebuildHeapFromBottom() {

	for i, top := ((q.size-1)-1)/2, false; i >= 0 && !top; i = (i - 1) / 2 {
		if i == 0 {
			top = true
		}
		changed := q.heapify(i, q.size)
		if !changed {
			break
		}
	}
}

func (q *PriorityQueue) rebuildHeapFromTop(pos int) {
	q.heapify(0, q.size)

}

func (q *PriorityQueue) heapify(top int, len int) (changed bool) {
	left := 2*top + 1
	right := 2*top + 2
	bigPos := top

	if left < len && q.array[left].Id > q.array[bigPos].Id {
		bigPos = left
	}
	if right < len && q.array[right].Id > q.array[bigPos].Id {
		bigPos = right
	}

	if bigPos != top {
		tmp := q.array[top]
		q.array[top] = q.array[bigPos]
		q.array[bigPos] = tmp
		q.heapify(bigPos, len)
	}

	changed = bigPos != top
	return changed
}
