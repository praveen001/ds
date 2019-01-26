package binaryheap

func (bh *BinaryHeap) push(values ...interface{}) bool {
	for _, value := range values {
		bh.list.Append(value)
		bh.percolateUp(bh.list.Length() - 1)
	}

	return true
}

func (bh *BinaryHeap) pop() (interface{}, bool) {
	if bh.list.Length() == 0 {
		return nil, false
	}

	last, _ := bh.list.Remove(bh.list.Length() - 1)
	first, _ := bh.list.Get(0)
	bh.list.Set(0, last)
	bh.percolateDown(0)

	return first, true
}

func (bh *BinaryHeap) peek() (interface{}, bool) {
	return bh.list.Get(0)
}

func (bh *BinaryHeap) clear() {
	bh.list.Clear()
}

func (bh *BinaryHeap) length() int {
	return bh.list.Length()
}

func (bh *BinaryHeap) percolateUp(idx int) {
	x, _ := bh.list.Get(idx)
	p, _ := bh.list.Get(idx / 2)

	for bh.compare(p, x) == 1 {
		bh.list.Swap(idx, idx/2)

		idx = idx / 2
		p, _ = bh.list.Get(idx / 2)
	}
}

func (bh *BinaryHeap) percolateDown(idx int) {
	x, _ := bh.list.Get(idx)

	for idx < bh.list.Length() {
		c := idx*2 + 1
		if lchv, ok := bh.list.Get(c); ok {

			if rchv, ok := bh.list.Get(c + 1); ok && bh.compare(lchv, rchv) == 1 {
				c = c + 1
			}

			if mchv, _ := bh.list.Get(c); bh.compare(x, mchv) == 1 {
				bh.list.Swap(idx, c)
			} else {
				break
			}

		} else {
			break
		}
		idx = c
	}
}

func (bh *BinaryHeap) string() string {
	return bh.list.String()
}
