package main

import "testing"

func TestIterator(t *testing.T){
	list := newList()
	list.append(0)
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(4)

	t.Run("Iterator", func(t *testing.T) {
		listIterator := newListIterator(list)

		answer := [5]int{0, 1, 2, 3, 4}

		index := 0
		for listIterator.first(); !listIterator.isDone(); listIterator.next() {
			if listIterator.currentItem() != answer[index] {
				t.Errorf("Actual: %d Expected: %d",  listIterator.currentItem(), answer[index])
			}
			index++
		}
	})

	t.Run("Reverse Iterator", func(t *testing.T) {
		reverseListIterator := newReverseListIterator(list)

		answer := [5]int{4, 3, 2, 1, 0}

		index := 0
		for reverseListIterator.first(); !reverseListIterator.isDone(); reverseListIterator.next() {
			if reverseListIterator.currentItem() != answer[index] {
				t.Errorf("Actual: %d Expected: %d",  reverseListIterator.currentItem(), answer[index])
			}
			index++
		}
	})
}