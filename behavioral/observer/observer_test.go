package main

import "testing"

func TestObserver(t *testing.T){
	observer1 := new(ConcreteObserver)
	observer2 := new(ConcreteObserver)

	subject := newSubject()

	subject.attach(observer1)
	actualObserver1 := subject.notify()

	subject.attach(observer2)
	actualObserver12 := subject.notify()

	expectObserver1 := "Subject updated "
	expectObserver12 := "Subject updated Subject updated "

	if actualObserver1 != expectObserver1 {
		t.Errorf("Actual: %s Expected: %s",  actualObserver1, expectObserver1)
	}
	
	if actualObserver12 != expectObserver12 {
		t.Errorf("Actual: %s Expected: %s",  actualObserver12, expectObserver12)
	}
}