package main

import "fmt"

type List struct{
	size int
	simpleArray []interface{}
}

func newList() *List {
	return &List{size: 0, simpleArray: make([]interface{}, 0)}
}

func (list *List) count() int {
	return list.size
}

func (list *List) get(index int) interface{} {
	return list.simpleArray[index]
}

func (list *List) append(data interface{}) {
	list.size++
	list.simpleArray = append(list.simpleArray, data)
}

type Iterator interface {
	first()
	next()
	isDone() bool
	currentItem() interface{}
}

type ListIterator struct {
	list *List
	currentIndex int
}

func newListIterator(list *List) *ListIterator {
	return &ListIterator{list: list, currentIndex: 0}
}

func (iterator *ListIterator) first() {
	iterator.currentIndex = 0
}

func (iterator *ListIterator) next() {
	iterator.currentIndex++
}

func (iterator *ListIterator) isDone() bool {
	return iterator.currentIndex >= iterator.list.count()
}

func (iterator *ListIterator) currentItem() interface{} {
	if(iterator.isDone()) {
		fmt.Println("ERROR, OUT OF BOUNDS")
	}
	return iterator.list.get(iterator.currentIndex)
}

type ReverseListIterator struct {
	list *List
	currentIndex int
}

func newReverseListIterator(list *List) *ReverseListIterator {
	return &ReverseListIterator{list: list, currentIndex: list.count() - 1}
}

func (iterator *ReverseListIterator) first() {
	iterator.currentIndex = iterator.list.count() - 1
}

func (iterator *ReverseListIterator) next() {
	iterator.currentIndex--
}

func (iterator *ReverseListIterator) isDone() bool {
	return iterator.currentIndex < 0
}

func (iterator *ReverseListIterator) currentItem() interface{} {
	if(iterator.isDone()) {
		fmt.Println("ERROR, OUT OF BOUNDS")
	}
	return iterator.list.get(iterator.currentIndex)
}

func main() {
	list := newList()
	list.append(0)
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(4)

	listIterator := newListIterator(list)
	reverseListIterator := newReverseListIterator(list)

	for listIterator.first(); !listIterator.isDone(); listIterator.next() {
		fmt.Println(listIterator.currentItem())
	}

	for reverseListIterator.first(); !reverseListIterator.isDone(); reverseListIterator.next() {
		fmt.Println(reverseListIterator.currentItem())
	}
}