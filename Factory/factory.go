package factory

import "fmt"

// Used to create objects without specifying the type or class of the objects explicitly
// You basically call a function which creates the object of the exact class based on the input or request context
// Helpful for abstracting the exact types of server implementation from clients/callers
// Open to Extension - Easy to add

type Notifier interface {
	Send(to, message string)
}

type EmailNotification struct{}

type WhatsappNotification struct{}

type SMSNotification struct{}

func (n *EmailNotification) Send(to, message string) {
	fmt.Printf("Destination : %v\t Message: %v\n", to, message)
}
func (n *SMSNotification) Send(to, message string) {
	fmt.Printf("Destination : %v\t Message: %v\n", to, message)
}
func (n *WhatsappNotification) Send(to, message string) {
	fmt.Printf("Destination : %v\t Message: %v\n", to, message)
}

func NotificationFactory(kind string) Notifier {
	if kind == "sms" {
		return &SMSNotification{}
	} else if kind == "whatsapp" {
		return &WhatsappNotification{}
	} else if kind == "email" {
		return &EmailNotification{}
	}
	return nil
}
