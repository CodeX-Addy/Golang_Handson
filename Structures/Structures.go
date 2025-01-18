package main

import "fmt"

//Struct defining way
type Person struct{
	Name string
	Age int
}

//Method for struct
func (p Person) Greet(){
	fmt.Printf("Hello my name is %s and age is %d", p.Name,p.Age);
}

func main(){
	p := Person{Name:"Aditya", Age:22};
	p.Greet();
}