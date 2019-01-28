package avltree

import (
	"testing"

	"github.com/praveen001/ds/utils"
)

func TestNew(t *testing.T) {
	at := New(utils.IntComparator)

	if at == nil {
		t.Errorf("Expected new tree to be created")
	}
}

func TestSet(t *testing.T) {
	at := New(utils.IntComparator)

	at.Set(10, 10)

	if l := at.Length(); l != 1 {
		t.Errorf("Expected %v Got %v", 1, l)
	}

	at.Set(10, 10)
	at.Set(20, 20)
	if l := at.Length(); l != 2 {
		t.Errorf("Expected %v Got %v", 2, l)
	}
}

func TestGet(t *testing.T) {
	at := New(utils.IntComparator)

	at.Set(1, 10)
	at.Set(2, 20)

	if val, ok := at.Get(1); val != 10 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 10, true, val, ok)
	}

	if val, ok := at.Get(3); val != nil || ok {
		t.Errorf("Expected %v, %v Got %v, %v", nil, false, val, ok)
	}
}

func TestRemove(t *testing.T) {
	at := New(utils.IntComparator)

	at.Set(10, 10)
	at.Set(20, 20)

	if ok := at.Remove(10); !ok {
		t.Errorf("Expected %v Got %v", true, ok)
	}

	if ok := at.Remove(10); ok {
		t.Errorf("Expected %v Got %v", false, ok)
	}
}

func TestHeight(t *testing.T) {
	at := New(utils.IntComparator)

	if h := at.Height(); h != 0 {
		t.Errorf("Expected %v Got %v", 0, h)
	}

	at.Set(10, 10)
	at.Set(20, 20)
	at.Set(30, 30)

	if h := at.Height(); h != 2 {
		t.Errorf("Expected %v Got %v", 2, h)
	}
}

func TestMin(t *testing.T) {
	at := New(utils.IntComparator)

	if m, ok := at.Min(); m != nil || ok {
		t.Errorf("Expected %v, %v Got %v, %v", nil, false, m, ok)
	}

	at.Set(100, 100)
	at.Set(10, 10)
	at.Set(20, 20)
	at.Set(5, 5)

	if m, ok := at.Min(); m.value != 5 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 5, true, m.value, ok)
	}
}

func TestMax(t *testing.T) {
	at := New(utils.IntComparator)

	if m, ok := at.Max(); m != nil || ok {
		t.Errorf("Expected %v, %v Got %v, %v", nil, false, m, ok)
	}

	at.Set(5, 5)
	at.Set(10, 10)
	at.Set(100, 100)
	at.Set(20, 20)

	if m, ok := at.Max(); m.value != 100 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 100, true, m.value, ok)
	}
}

func TestContains(t *testing.T) {
	at := New(utils.IntComparator)

	if ok := at.Contains(1); ok {
		t.Errorf("Expected %v Got %v", false, ok)
	}

	at.Set(5, 5)
	at.Set(10, 10)
	at.Set(100, 100)
	at.Set(20, 20)

	if ok := at.Contains(20); !ok {
		t.Errorf("Expected %v Got %v", true, ok)
	}

	at.Remove(20)

	if ok := at.Contains(20); ok {
		t.Errorf("Expected %v Got %v", false, ok)
	}
}

func TestLength(t *testing.T) {
	at := New(utils.IntComparator)

	if l := at.Length(); l != 0 {
		t.Errorf("Expected %v Got %v", 0, l)
	}

	at.Set(5, 5)
	at.Set(10, 10)
	at.Set(100, 100)
	at.Set(20, 20)

	if l := at.Length(); l != 4 {
		t.Errorf("Expected %v Got %v", 4, l)
	}

	at.Remove(10)

	if l := at.Length(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}
}

func TestClear(t *testing.T) {
	at := New(utils.IntComparator)

	if l := at.Length(); l != 0 {
		t.Errorf("Expected %v Got %v", 0, l)
	}

	at.Set(5, 5)
	at.Set(10, 10)
	at.Set(100, 100)
	at.Set(20, 20)

	if l := at.Length(); l != 4 {
		t.Errorf("Expected %v Got %v", 4, l)
	}

	at.Clear()

	if l := at.Length(); l != 0 {
		t.Errorf("Expected %v Got %v", 0, l)
	}

	if str := at.InOrder(); str.String() != "[]" {
		t.Errorf("Expected %v Got %v", "[]", str.String())
	}
}

func TestInOrder(t *testing.T) {
	at := New(utils.IntComparator)

	if str := at.InOrder(); str.String() != "[]" {
		t.Errorf("Expected %v Got %v", "[]", str.String())
	}

	at.Set(5, 5)
	at.Set(10, 10)
	at.Set(100, 100)
	at.Set(20, 20)

	if str := at.InOrder(); str.String() != "[5 10 20 100]" {
		t.Errorf("Expected %v Got %v", "[5 10 20 100]", str.String())
	}
}

func TestPreOrder(t *testing.T) {
	at := New(utils.IntComparator)

	if str := at.PreOrder(); str.String() != "[]" {
		t.Errorf("Expected %v Got %v", "[]", str.String())
	}

	at.Set(5, 5)
	at.Set(10, 10)
	at.Set(7, 7)
	at.Set(100, 100)
	at.Set(20, 20)

	if str := at.PreOrder(); str.String() != "[7 5 20 10 100]" {
		t.Errorf("Expected %v Got %v", "[7 5 20 10 100]", str.String())
	}
}

func TestPostOrder(t *testing.T) {
	at := New(utils.IntComparator)

	if str := at.PostOrder(); str.String() != "[]" {
		t.Errorf("Expected %v Got %v", "[]", str.String())
	}

	at.Set(5, 5)
	at.Set(10, 10)
	at.Set(7, 7)
	at.Set(100, 100)
	at.Set(20, 20)

	if str := at.PostOrder(); str.String() != "[5 10 100 20 7]" {
		t.Errorf("Expected %v Got %v", "[7 20 100 10 5]", str.String())
	}
}

func TestLeftRotate(t *testing.T) {
	at := New(utils.IntComparator)

	/*
		Left rotation in root - Insert 10, 50, 100

					10														50
				 		\													 /	\
						50						===>				10	100
							\
							100
	*/
	at.Set(10, 10)
	at.Set(50, 50)
	at.Set(100, 100)

	if h := at.Height(); h != 2 {
		t.Errorf("Expected %v Got %v", 2, h)
	}

	if l := at.Length(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}

	if str := at.InOrder().String(); str != "[10 50 100]" {
		t.Errorf("Expected %v Got %v", "[10 50 100]", str)
	}
	at.Clear()

	/*
		Left Rotation in subtrees - Insert 10, 5, 100, 200, 300

				10														10
			 /  \													 /  \
			5		100						===>				5		200
						\													 /	 \
						200												100		300
							\
							300
	*/
	at.Set(10, 10)
	at.Set(5, 5)
	at.Set(100, 100)
	at.Set(200, 200)
	at.Set(300, 300)

	if h := at.Height(); h != 3 {
		t.Errorf("Expected %v Got %v", 3, h)
	}

	if l := at.Length(); l != 5 {
		t.Errorf("Expected %v Got %v", 5, l)
	}

	if str := at.InOrder().String(); str != "[5 10 100 200 300]" {
		t.Errorf("Expected %v Got %v", "[5 10 100 200 300]", str)
	}
	at.Clear()

	/* Left Rotation on Deletion - Insert 100, 200, 50, 20, Delete 200
			100																	100										50
		 /   \															 /										 /	\
		50		200      == delete 20 ==>			50	      		===>		20	100
	 /																	 /
	20																	20
	*/

	at.Set(100, 100)
	at.Set(200, 200)
	at.Set(50, 50)
	at.Set(20, 20)
	at.Remove(200)

	if h := at.Height(); h != 2 {
		t.Errorf("Expected %v Got %v", 3, h)
	}

	if l := at.Length(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}

	if str := at.InOrder().String(); str != "[20 50 100]" {
		t.Errorf("Expected %v Got %v", "[20 50 100]", str)
	}
}

func TestRightRotate(t *testing.T) {
	at := New(utils.IntComparator)

	/*
		Right rotation in root - Insert 100, 50, 10

					100													50
				 /													 /	\
				50						===>				  10	 100
			 /
			10
	*/
	at.Set(100, 100)
	at.Set(50, 50)
	at.Set(10, 10)

	if h := at.Height(); h != 2 {
		t.Errorf("Expected %v Got %v", 2, h)
	}

	if l := at.Length(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}

	if str := at.InOrder().String(); str != "[10 50 100]" {
		t.Errorf("Expected %v Got %v", "[10 50 100]", str)
	}
	at.Clear()

	/*
			Left Rotation in subtrees - Insert 200, 100, 300, 10, 5

					200															200
				 /   \													 /  \
				100	 300						===>				10	 300
			  /													 		 /  \
			10														 	5   100
			/
		 5
	*/
	at.Set(200, 200)
	at.Set(100, 100)
	at.Set(300, 300)
	at.Set(10, 10)
	at.Set(5, 5)

	if h := at.Height(); h != 3 {
		t.Errorf("Expected %v Got %v", 3, h)
	}

	if l := at.Length(); l != 5 {
		t.Errorf("Expected %v Got %v", 5, l)
	}

	if str := at.InOrder().String(); str != "[5 10 100 200 300]" {
		t.Errorf("Expected %v Got %v", "[5 10 100 200 300]", str)
	}
	at.Clear()

	/* Right Rotation on Deletion - Insert 100, 200, 50, 300, Delete 50

			100																	100											200
		 /   \															 		\								 		 /	\
		50		200      == delete 50 ==>				   200   		===>		 100	 300
	 					\																	 \
						300																	300
	*/

	at.Set(100, 100)
	at.Set(200, 200)
	at.Set(50, 50)
	at.Set(300, 300)
	at.Remove(50)

	if h := at.Height(); h != 2 {
		t.Errorf("Expected %v Got %v", 2, h)
	}

	if l := at.Length(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}

	if str := at.InOrder().String(); str != "[100 200 300]" {
		t.Errorf("Expected %v Got %v", "[100 200 300]", str)
	}
}

func TestRightLeftRotate(t *testing.T) {
	at := New(utils.IntComparator)

	/*
		Right-Left rotation in root - Insert 10, 50, 40

					10														40
				 		\													 /	\
						50						===>				10	50
						/
					 40
	*/
	at.Set(10, 10)
	at.Set(50, 50)
	at.Set(40, 40)

	if h := at.Height(); h != 2 {
		t.Errorf("Expected %v Got %v", 2, h)
	}

	if l := at.Length(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}

	if str := at.InOrder().String(); str != "[10 40 50]" {
		t.Errorf("Expected %v Got %v", "[10 40 50]", str)
	}
	at.Clear()

	/*
		Right-Left Rotation in subtrees - Insert 10, 5, 100, 200, 150

				10														10
			 /  \													 /  \
			5		100						===>				5		200
						\													 /	 \
						200												100		150
						/
					150
	*/
	at.Set(10, 10)
	at.Set(5, 5)
	at.Set(100, 100)
	at.Set(200, 200)
	at.Set(150, 150)

	if h := at.Height(); h != 3 {
		t.Errorf("Expected %v Got %v", 3, h)
	}

	if l := at.Length(); l != 5 {
		t.Errorf("Expected %v Got %v", 5, l)
	}

	if str := at.InOrder().String(); str != "[5 10 100 150 200]" {
		t.Errorf("Expected %v Got %v", "[5 10 100 150 200]", str)
	}
	at.Clear()

	/*
			Right-Left Rotation in subtrees on deletion  - Insert 10, 5, 100, 200, 90, 1, 150, Delete 90

				 10																		 	 10																		   10																	 10
				/  \													 					/  \																	  /  \																/  \
			 5		100						==Delete 90=>				 5		100								===>						 5		100	   			=======>					5		150
			/		 /	\													 			/			 	\				right rotate at 200			/				\    left rotate at 100		 /		/	 \
		 1		90	200														 1		 		200														 1				150												1		100		200
							/																				/																					\
						150																			150																					200
	*/
	at.Set(10, 10)
	at.Set(5, 5)
	at.Set(100, 100)
	at.Set(200, 200)
	at.Set(90, 90)
	at.Set(1, 1)
	at.Set(150, 150)
	at.Remove(90)

	if h := at.Height(); h != 3 {
		t.Errorf("Expected %v Got %v", 3, h)
	}

	if l := at.Length(); l != 6 {
		t.Errorf("Expected %v Got %v", 6, l)
	}

	if str := at.InOrder().String(); str != "[1 5 10 100 150 200]" {
		t.Errorf("Expected %v Got %v", "[1 5 10 100 150 200]", str)
	}
}

func TestLeftRightRotate(t *testing.T) {
	at := New(utils.IntComparator)

	/*
		Left-Right rotation in root - Insert 100, 50, 60

					100													60
				 /													 /	\
				50						===>				  50	 100
			 		\
					60
	*/
	at.Set(100, 100)
	at.Set(50, 50)
	at.Set(60, 60)

	if h := at.Height(); h != 2 {
		t.Errorf("Expected %v Got %v", 2, h)
	}

	if l := at.Length(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}

	if str := at.InOrder().String(); str != "[50 60 100]" {
		t.Errorf("Expected %v Got %v", "[50 60 100]", str)
	}
	at.Clear()

	/*
			Left Rotation in subtrees - Insert 200, 100, 300, 10, 60

					200															200
				 /   \													 /  \
				100	 300						===>				60	 300
			  /													 		 /  \
			10														 	10   100
				\
		 		60
	*/
	at.Set(200, 200)
	at.Set(100, 100)
	at.Set(300, 300)
	at.Set(10, 10)
	at.Set(60, 60)

	if h := at.Height(); h != 3 {
		t.Errorf("Expected %v Got %v", 3, h)
	}

	if l := at.Length(); l != 5 {
		t.Errorf("Expected %v Got %v", 5, l)
	}

	if str := at.InOrder().String(); str != "[10 60 100 200 300]" {
		t.Errorf("Expected %v Got %v", "[10 60 100 200 300]", str)
	}
	at.Clear()

	/*
			Right-Left Rotation in subtrees on Deletion - Insert 200, 100, 300, 10, 150, 400, 60, Delete 150

					 200															200																	200																		200
				 /   	\													 		/  \															 /	 \																 /		\
				100	 	300						===>					100	 	300					======>						100		300				========>							60		300
			  /	\			\					Delete 150		 /  			\		 Left rotate at 10			/				\			Right Rotate at 100		 /  \			\
			10	 150	 400										10 			  400												60				400													10  100		400
				\																	\																	/
		 		60																60															10
	*/
	at.Set(200, 200)
	at.Set(100, 100)
	at.Set(300, 300)
	at.Set(10, 10)
	at.Set(150, 150)
	at.Set(400, 400)
	at.Set(60, 60)
	at.Remove(150)

	if h := at.Height(); h != 3 {
		t.Errorf("Expected %v Got %v", 3, h)
	}

	if l := at.Length(); l != 6 {
		t.Errorf("Expected %v Got %v", 6, l)
	}

	if str := at.InOrder().String(); str != "[10 60 100 200 300 400]" {
		t.Errorf("Expected %v Got %v", "[10 60 100 200 300 400]", str)
	}
}
