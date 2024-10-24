package main

import "fmt"

const Login string = "abcdefgh"; //This is a public variable, in golang, we generally define the public variable name starts with capital letter.

func main() {

	//String type
	var user string = "Aditya";
	fmt.Println(user);
	fmt.Printf("The type of the variable user is: %T\n", user);

	//Boolean type
	var check bool = true;
	fmt.Println(check);
	fmt.Printf("The type of the variable check is: %T\n", check);

	//Integer type
	//We can simply use int as well
	var num uint8 = 255; //Highest permissable value for uint8
	fmt.Println(num);
	fmt.Printf("The type of the variable num is: %T\n", num);

	//Floating type
	var floatNum float32 = 255.45265525;
	fmt.Println(floatNum);
	fmt.Printf("The type of the variable floatNum is: %T\n", floatNum);

	//Some aliases
	var number int;
	fmt.Println(number); // It will give 0
	fmt.Printf("The type of the variable number is: %T\n", number);

	//Other types of declaring the variables
	//Implicit type
	var num1 = 2;
	fmt.Println(num1);

	//Other type
	name1 := "esoteric";
	fmt.Println(name1);

	fmt.Println(Login);
	fmt.Printf("The type of the login variable is: %T\n", Login);

}