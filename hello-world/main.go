package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
)

type Person struct {
	name string
	age  int
}

func (p *Person) SetAge(age int) {
	p.age = age
}

func main() {
	var a float64 = 100_000
	fmt.Println(math.Log10(a))
	println("Hello, World!")

	var name string = "Hwal"
	code := 4455
	fmt.Println("wowowowowo!!!!!!!!!!!", name, code)

	//////////////////////////////////////

	var person Person = Person{
		name: "Hwal",
		age:  30,
	}
	fmt.Printf("%v\n", person)
	fmt.Printf("%#v\n", person)
	fmt.Printf("%+v\n", person)
	fmt.Printf("%T\n", person)

	//////////////////////////////////////

	fmt.Println("Person:", person.name, "is", person.age, "years old.")
	fmt.Printf("Person: %s is %d years old.\n", person.name, person.age)
	person.SetAge(person.age + 1)
	s := fmt.Sprintf("Person: %s is %d years old.\n", person.name, person.age)
	fmt.Print(s)

	cmd := exec.Command("echo", "Hello from exec.Command!")
	log.Println("Executing command:", cmd.String())
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing command:", err)
		log.Fatal("Command execution failed:", err)
	}
	log.Println("Command executed successfully.")

	//////////////////////////////////////
}
