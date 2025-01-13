package main

import "fmt"

//Swap Function returning two variables
func swap(a, b int) (int , int){
	return b, a;
}

func main(){
	
	a, b:= 10, 20;

	fmt.Println("Before swapping:", a, b);
	a, b = swap(a, b);
	fmt.Println("After swapping:", a, b);

	//Anonymous func declaration
	add:= func(a, b int) int{
		return a+b;
	}

	fmt.Println("Sum of 2 and 3 is:", add(2,3));
}
