package main

import (
	"fmt"

	decorator "payment.go/Decorator"
	factory "payment.go/Factory"
	observer "payment.go/Observer"
	singleton "payment.go/Singleton"
	strategy "payment.go/Strategy"
)

func main() {
	//Example main for Singleton

	fmt.Printf("\n------ SINGLETON EXAMPLE -------")
	logger1 := singleton.GetLogger()
	logger2 := singleton.GetLogger()

	logger1.Log("\nHi this is logger 1")
	logger2.Log("\nHi this is logger 2")

	if logger1 == logger2 {
		fmt.Printf("Print if logger 1 is same as logger 2")
	} else {
		fmt.Print("logger 1 is not same as logger 2")
	}
	fmt.Printf("\n\n\n------ FACTORY EXAMPLE -------")

	notifier := factory.NotificationFactory("sms")
	notifier.Send("abc@email.com", "Hi this is a notification factory")

	fmt.Printf("\n\n\n------ STRATEGY EXAMPLE -------")

	planner := strategy.TravelWrapper{}
	planner.SetStrategy(&strategy.Car{}) // use & since Travel expects pointer receiver
	planner.Travel("Home", "School")

	// planner.Travel("Home", "School")

	notifier2 := strategy.NotificationWrapper{}
	notifier2.SetNotifier(&strategy.SMS{})

	notifier2.Notify("Hi!!")

	fmt.Printf("\n\n\n------ Observer EXAMPLE -------")

	subject1 := observer.Channel{}
	observer1 := observer.Subscriber1{}
	observer2 := observer.Subscriber2{}
	subject1.RegisterObserver(&observer1)
	subject1.RegisterObserver(&observer2)

	subject1.SetNotification("Sample notification")

	fmt.Printf("\n\n\n------ Decorator EXAMPLE -------")

	var mycoffee decorator.Beverage = decorator.Coffee{}

	fmt.Printf(mycoffee.Description()+"  at Rs %v", mycoffee.Cost())

	mycoffee = decorator.BeverageWithMilk{Beverage: mycoffee}
	fmt.Printf(mycoffee.Description()+"  at Rs %v", mycoffee.Cost())

}
