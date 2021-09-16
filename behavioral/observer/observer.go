package main

import "fmt"

type Subject interface {
	attach(Observer)
	detach(Observer)
	notify()
}

type Observer interface{
	update(Subject)
}

type ConcreteObserver struct {
}

func (observer *ConcreteObserver) update(subject Subject) {
	fmt.Println("Subject:", subject, "updated")
}

type ConcreteSubject struct {
	observers []Observer
}

func (subject *ConcreteSubject) attach(observer Observer){
	subject.observers = append(subject.observers, observer)
}

func (subject *ConcreteSubject) detach(observer Observer){
	var index int
	for i := 0; i < len(subject.observers); i++ {
			if(observer == subject.observers[i]){
				index = i
			}
	}

	remainingObservers := make([]Observer, 0)
	remainingObservers = append(remainingObservers, subject.observers[:index]...)
	subject.observers = append(remainingObservers, subject.observers[index+1:]...)
}

func (subject *ConcreteSubject) notify() {
	for i := 0; i < len(subject.observers); i++ {
			subject.observers[i].update(subject)
	}
}

func newSubject() Subject {
	subject := new(ConcreteSubject)
	subject.observers = make([]Observer, 0)
	return subject
}

func main()  {
	observer1 := new(ConcreteObserver)
	observer2 := new(ConcreteObserver)

	subject := newSubject()
	
	fmt.Println("Attaching 1")
	subject.attach(observer1)
	fmt.Println("Notifying")
	subject.notify()

	fmt.Println("Attaching 2")
	subject.attach(observer2)
	fmt.Println("Notifying")
	subject.notify()
}