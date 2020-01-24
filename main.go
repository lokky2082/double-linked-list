package main

import (
	"fmt"
	"github.com/lokky2082/double-linked-list/list"
)
func main() {
	testList := &list.List{}
	Item1 := &list.Item {
		Value: 1,
	}
	Item2 := &list.Item {
		Value: 2,
	}
	Item3 := &list.Item {
		Value: 3,
	}
	Item4 := &list.Item {
		Value: 4,
	}
	testList.Append(Item1)
	testList.Append(Item2)
	testList.Append(Item3)
	testList.Append(Item4)
	testList.Display()
	testList.Delete(Item2)
	testList.Display()
	testList.Delete(Item1)
	testList.Display()
	testList.Append(Item1)
	fmt.Println(testList.Tail().Value)
	fmt.Println(testList.Head().Value)
	fmt.Println(Item1.Value)
	fmt.Println(Item1.Next())
	fmt.Println(Item1.Prev())
	testList.Display()
}