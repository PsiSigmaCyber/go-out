package main

import "fmt"

type email struct {
	address string
}

type SMS struct {
	number string
}

func (e email) Send() { //   A receiver function utliizes a struct
	fmt.Println("Sending mail", e.address)
}
func (s SMS) Send() { //same
	fmt.Println("sending sms", s.number)
}

type Notifier interface { //interface was created that utlizees the reciever function?
	Send()
}

// func TriggerAlert(n Notifier) { //function to call the interface?
// 	n.Send()

//}
func main() { // two alerts (filling the boxes) and basically running the above three things?
	// alert1 := email{ /*address: "ceo@tech.com"*/ }
	// alert2 := SMS{number: "12345"}
	// TriggerAlert(alert1)
	// TriggerAlert(alert2)
	alert1 := email{address: "ceo@tech.com"}
	alert2 := SMS{number: "12345"}
	alert3 := email{address: "dev@tech.com"}

	// THE MAGIC: We create a slice of the *Interface*, not the structs.
	// The Bouncer lets both types into the exact same queue!
	queue := []Notifier{alert1, alert2, alert3}

	// A single loop processes completely different data types automatically.
	for _, item := range queue {
		item.Send()
	}
}
