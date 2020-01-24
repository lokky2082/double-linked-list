package list

import (
	"fmt"
)
// Item struct
type Item struct {
	Value interface{}
	next, prev *Item
}
// List struct
type List struct {
	length int
	head, tail *Item
}
// Append adds *Item to the end of List
func(list *List) Append(n *Item) {
  if list.length == 0 {
		list.head = n
		list.tail = n
		list.length++
		return
	}
	prevTail := list.tail
	list.tail = n
	list.tail.prev = prevTail
	list.tail.next = nil
	prevTail.next = n
	list.length++
}
// Prepend adds *Item to the start of List
func(list *List) Prepend(n *Item) {
  if list.length == 0 {
		list.head = n
		list.tail = n
		list.length++
		return
	}
	prevHead := list.head
	list.head = n
	list.head.next = prevHead
	list.head.prev = nil
	prevHead.prev = n
	list.length++
	fmt.Println(list.length)
}
// Delete removes a Item from the List
func(list *List) Delete(item *Item) {
  if list.length == 0 {
		return
	}
  if list.head == item {
		list.head = item.next
		list.head.prev = nil
		list.length--
		return
	}
	if list.tail == item {
		list.tail = item.prev
		list.tail.next = nil
		list.length--
		return
	}
	prevItem := item.prev
	nextItem := item.next
	prevItem.next = item.next
	nextItem.prev = item.prev
	list.length--
}
// Len return length of List
func(list *List) Len() int {
  return list.length
}
// Head return head of List
func(list *List) Head() *Item {
  return list.head
}
// Tail return tail of List
func(list *List) Tail() *Item {
  return list.tail
}

// Display Items values from the List
func (list *List) Display() {
	head := list.head
	for head != nil {
			fmt.Printf("%+v ->", head.Value)
			head = head.next
	}
	fmt.Println()
}
// Next return next node of item
func(item *Item) Next() *Item {
  return item.next
}
// Prev return prev node of item
func(item *Item) Prev() *Item {
  return item.prev
}
