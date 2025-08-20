package main

import "fmt"

func main(){
	// If else statement
	num := 10
	if num%2==0{
		fmt.Println("Even")
	}else{
		fmt.Println("Odd")
	}

	// For Loop
	for i:=1; i<=10; i++{
		fmt.Println(i);
	}

	//Switch case
	a := 7
	
	switch a{
	case 1:
		fmt.Println("One");
	case 7:
		fmt.Println("Six");
	default:
		fmt.Println("Another Number");
	}

}
