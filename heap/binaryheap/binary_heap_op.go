package binaryheap

func (bh *BinaryHeap) push(values ...interface{}) bool {
	for _, value := range values {
		bh.list.PushBack(value)
		bh.percolateUp(bh.list.Len() - 1)
	}

	return true
}

func (bh *BinaryHeap) pop() (interface{}, bool) {
	if bh.list.Len() == 0 {
		return nil, false
	}

	first, _ := bh.list.At(0)
	last, _ := bh.list.Remove(bh.list.Len() - 1)

	bh.list.Insert(0, last)
	bh.percolateDown(0)

	return first, true
}

func (bh *BinaryHeap) peek() (interface{}, bool) {
	return bh.list.At(0)
}

func (bh *BinaryHeap) clear() {
	bh.list.Clear()
}

func (bh *BinaryHeap) len() int {
	return bh.list.Len()
}

func (bh *BinaryHeap) percolateUp(idx int) {
	x, _ := bh.list.At(idx)

	for idx > -1 {
		pidx := (idx - 1) / 2
		p, _ := bh.list.At(pidx)

		if bh.compare(p, x) == bh.variant {
			bh.list.Swap(idx, pidx)

			idx = pidx
			continue
		}
		break
	}
}

func (bh *BinaryHeap) percolateDown(idx int) {
	x, _ := bh.list.At(idx)

	for idx < bh.list.Len() {
		c := idx*2 + 1
		if lchv, ok := bh.list.At(c); ok {

			if rchv, ok := bh.list.At(c + 1); ok && bh.compare(lchv, rchv) == bh.variant {
				c = c + 1
			}

			if mchv, _ := bh.list.At(c); bh.compare(x, mchv) == bh.variant {
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
