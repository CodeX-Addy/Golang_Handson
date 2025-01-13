package main

import "fmt"

//Swap Function returning two variables
func swap(a, b int) (int , int){
	return b, a;
}

func main(){
	
	a, b:= 10, 20;

	fmt.Println("Before swapping: ", a, b);
	a, b = swap(a, b);
	fmt.Println("After swapping: ", a, b);

}
