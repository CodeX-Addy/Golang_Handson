//Using pointers in functions

package main

import "fmt"

func change(x* int){
	*x = 20;
}

func main(){
	x := 10;
	fmt.Println("Before:", x);

	change(&x);

	fmt.Println("After:", x);
}