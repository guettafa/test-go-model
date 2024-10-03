package main

import (
	"fmt"
	model "myproject/models"
)

type Person struct {
	name string
	lastName string
	age uint8
}

func (people Person) GiveDescription() string {
	return fmt.Sprintf("My name is %s and am %d", people.name, people.age)
}

func main() {
	u := model.NewUser("Arslane", "Guettaf", 19)
	fmt.Println(u)

	u.SetLastName("Bob")
	fmt.Println(u)

	// p1 := Person{"Myname", "MyLastname", 19}
	// p1.GiveDescription()
	// fmt.Println(p1.GiveDescription())

	// numbers := []uint{1,2,3,4,5}
	// for i := range numbers {
	// 	fmt.Println(i) // print each index and number from the array numbers
	// }	


	// for i := 0; i <= 4; i++ {
	// 	fmt.Println(i)
	// }
	
	// var somePtr *int
	// if somePtr == nil {
	// 	fmt.Print(somePtr, "This is a null ptr") // This is a null ptr
	// }

	// greeting := "Hello there!"
  	// Brainwash(&greeting)
	// fmt.Println("greeting is now:", greeting)
}
