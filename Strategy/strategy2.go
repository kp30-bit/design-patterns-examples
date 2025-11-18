package strategy

import (
	"fmt"
)

type NotificationSystem interface {
	send(message string)
}

type Email struct{}
type SMS struct{}
type Push struct{}

func (n *Email) send(message string) {
	fmt.Printf("\nMessage via email : %v\n", message)
}
func (n *SMS) send(message string) {
	fmt.Printf("\nMessage via sms : %v\n", message)
}
func (n *Push) send(message string) {
	fmt.Printf("\nMessage via push : %v\n", message)
}

type NotificationWrapper struct {
	Notifier NotificationSystem
}

func (n *NotificationWrapper) SetNotifier(notifier NotificationSystem) {
	n.Notifier = notifier
}

func (n *NotificationWrapper) Notify(message string) {
	n.Notifier.send(message)
}
