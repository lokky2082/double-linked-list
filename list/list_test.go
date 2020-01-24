package list

import (
	"testing"
)
func TestLen(t *testing.T) {
	testList := &List{}
	Item1 := &Item {
		Value: 1,
	}
	Item2 := &Item {
		Value: 2,
	}
  if testList.Len() > 0 {
		t.Errorf("list should be empty")
  }
	testList.Append(Item1)
	testList.Append(Item2)
	if size := testList.Len(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}
}
func TestAppend(t *testing.T) {
	testList := &List{}
	Item1 := &Item {
		Value: 1,
	}
	Item2 := &Item {
		Value: 2,
	}
  if testList.Len() > 0 {
		t.Errorf("list should be empty")
  }
	testList.Append(Item1)
	if size := testList.Len(); size != 1 {
		t.Errorf("wrong count, expected 1 and got %d", size)
	}
	testList.Append(Item2)
	if tail:= testList.Tail(); tail != Item2 {
		t.Errorf("wrong position, expected %+v and got %+v", Item2, tail)
	}
	if size := testList.Len(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}
}
func TestPrepend(t *testing.T) {
	testList := &List{}
	Item1 := &Item {
		Value: 1,
	}
	Item2 := &Item {
		Value: 2,
	}
	testList.Prepend(Item1)
	if size := testList.Len(); size != 1 {
		t.Errorf("wrong count, expected 1 and got %d", size)
	}
	testList.Prepend(Item2)
	if head:= testList.Head(); head != Item2 {
		t.Errorf("wrong position, expected %+v and got %+v", Item2, head)
	}
	if size := testList.Len(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}
}
func TestDelete(t *testing.T) {
	testList := &List{}
	Item1 := &Item {
		Value: 1,
	}
	Item2 := &Item {
		Value: 2,
	}
	testList.Append(Item1)
	testList.Append(Item2)
	if tail:= testList.Tail(); tail != Item2 {
		t.Errorf("wrong position, expected %+v and got %+v", Item2, tail)
	}
	testList.Delete(Item2)
	if tail:= testList.Tail(); tail != Item1 {
		t.Errorf("wrong position, expected %+v and got %+v", Item1, tail)
	}
	if size := testList.Len(); size != 1 {
		t.Errorf("wrong count, expected 1 and got %d", size)
	}
}
func TestNextPrev(t *testing.T) {
	testList := &List{}
	Item1 := &Item {
		Value: 1,
	}
	Item2 := &Item {
		Value: 2,
	}
	Item3 := &Item {
		Value: 3,
	}
	testList.Append(Item1)
	testList.Append(Item2)
	testList.Append(Item3)
	if next:= Item2.Next(); next != Item3 {
		t.Errorf("wrong next node, expected %+v and got %+v", Item3, next)
	}
	if prev:= Item2.Prev(); prev != Item1 {
		t.Errorf("wrong prev node, expected %+v and got %+v", Item1, prev)
	}
}
func TestValue(t *testing.T) {
	testList := &List{}
	Item1 := &Item {
		Value: 1,
	}
	testList.Append(Item1)
	if val:= Item1.getValue(); val != 1 {
		t.Errorf("wrong value of item, expected 1 and got %d",  val)
	}
}