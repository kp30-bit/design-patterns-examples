package decorator

import (
	"fmt"
	"log"
)

type Notifier interface {
	Send(msg string)
}

type EmailNotifier struct {
}

func (e *EmailNotifier) Send(msg string) {
	fmt.Printf("\n\n This is the message : %v\n", msg)
}

type LogNotifier struct {
	Notifier Notifier
}

func (ln *LogNotifier) Send(msg string) {
	ln.Notifier.Send(msg)

	log.Print("\n\nNotification Logged!")

}
