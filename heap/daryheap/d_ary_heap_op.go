package daryheap

func (dh *DAryHeap) push(values ...interface{}) bool {
	for _, value := range values {
		dh.list.Append(value)
		dh.percolateUp(dh.list.Length() - 1)
	}

	return true
}

func (dh *DAryHeap) pop() (interface{}, bool) {
	if dh.list.Length() == 0 {
		return nil, false
	}

	last, _ := dh.list.Remove(dh.list.Length() - 1)
	first, _ := dh.list.Get(0)
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

func (dh *DAryHeap) length() int {
	return dh.list.Length()
}

func (dh *DAryHeap) percolateUp(idx int) {
	x, _ := dh.list.Get(idx)
	p, _ := dh.list.Get(idx / dh.d)

	for dh.compare(p, x) == dh.variant {
		dh.list.Swap(idx, idx/dh.d)

		idx = idx / dh.d
		p, _ = dh.list.Get(idx)
	}
}

func (dh *DAryHeap) percolateDown(idx int) {
	x, _ := dh.list.Get(idx)

	for idx < dh.list.Length() {
		c := idx*dh.d + 1

		maxIdx := c
		max, ok := dh.list.Get(c)
		if !ok {
			return
		}

		for i := 1; i < dh.d; i++ {
			if chv, ok := dh.list.Get(c); ok {
				if dh.compare(chv, max) == dh.variant {
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
