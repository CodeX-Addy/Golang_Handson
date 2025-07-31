package main

import "fmt"

// Defining struct
type Person struct{
	Name string
	Age int
}

// Method for struct
func (p Person) Greet(){
	fmt.Printf("Hey My Name is %s and I am %d years old", p.Name, p.Age);
}

func main(){
	p := Person{Name: "Aditya", Age: 22};
	p.Greet();
}
