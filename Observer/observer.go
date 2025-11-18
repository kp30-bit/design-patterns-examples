package observer

import "fmt"

type Observer interface {
	Notify(Notification string)
}

type Subject interface {
	RegisterObserver(Observer)
	DeleteObserver(Observer)
	NotifyObserver()
}

//Observers

type Subscriber1 struct{}
type Subscriber2 struct{}

func (s *Subscriber1) Notify(notification string) {
	fmt.Printf("\n\nNotify notification to observer 1: %v\n", notification)
}
func (s *Subscriber2) Notify(notification string) {
	fmt.Printf("\nNotify notification to observer 2: %v\n", notification)
}

// Subject struct
type Channel struct {
	Observer     []Observer
	Notification string
}

func (s *Channel) RegisterObserver(o Observer) {
	s.Observer = append(s.Observer, o)
}

func (s *Channel) DeleteObserver(o Observer) {
	for i, v := range s.Observer {
		if v == o {
			s.Observer = append(s.Observer[:i], s.Observer[i+1:]...)
		}
	}
}

func (s *Channel) SetNotification(notification string) {
	s.Notification = notification
}
