package daryheap

func (dh *DAryHeap) push(values ...interface{}) bool {
	for _, value := range values {
		dh.list.PushBack(value)
		dh.percolateUp(dh.list.Len() - 1)
	}

	return true
}

func (dh *DAryHeap) pop() (interface{}, bool) {
	if dh.list.Len() == 0 {
		return nil, false
	}

	first, _ := dh.list.Get(0)
	last, _ := dh.list.Remove(dh.list.Len() - 1)

	dh.list.Set(0, last)
	dh.percolateDown(0)

	return first, true
}

func (dh *DAryHeap) peek() (interface{}, bool) {
	return dh.list.Get(0)
}

func (dh *DAryHeap) clear() {
	dh.list.Clear()
}

func (dh *DAryHeap) len() int {
	return dh.list.Len()
}

func (dh *DAryHeap) percolateUp(idx int) {
	x, _ := dh.list.Get(idx)

	for idx > -1 {
		pidx := (idx - 1) / dh.d
		p, _ := dh.list.Get(pidx)

		if dh.compare(p, x) == dh.variant {
			dh.list.Swap(idx, pidx)

			// Move up
			idx = pidx
			continue
		}
		break
	}
}

func (dh *DAryHeap) percolateDown(idx int) {
	x, _ := dh.list.Get(idx)

	for idx < dh.list.Len() {
		c := idx*dh.d + 1

		maxIdx := c
		max, ok := dh.list.Get(c)
		if !ok {
			return
		}

		for i := 1; i < dh.d; i++ {
			if chv, ok := dh.list.Get(c); ok {
				if dh.compare(max, chv) == dh.variant {
					max = chv
					maxIdx = c
				}
			} else {
				break
			}
			c++
		}

		if dh.compare(x, max) == dh.variant {
			dh.list.Swap(idx, maxIdx)
		}
		idx = c
	}
}
