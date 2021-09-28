package main

type Subject interface {
	attach(Observer)
	detach(Observer)
	notify() string
}

type Observer interface{
	update(Subject) string
}

type ConcreteObserver struct {
}

func (observer *ConcreteObserver) update(subject Subject) string {
	return "Subject updated "
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

func (subject *ConcreteSubject) notify() string {
	response := ""
	for i := 0; i < len(subject.observers); i++ {
			response += subject.observers[i].update(subject)
	}
	return response
}

func newSubject() Subject {
	subject := new(ConcreteSubject)
	subject.observers = make([]Observer, 0)
	return subject
}
