package main

import "fmt"

func main(){
	// Array Declaration
	arr := [5] int {1,2,3,4,5};
	fmt.Println(arr);

	//Slice 
	slice := []int{1,2,3,4,5};
	slice = append(slice, 6);
	fmt.Println("Sliced:", slice);

	//Map
	scores := map[string]int{"A":10, "B":20};
	scores["C"] = 30;
	fmt.Println("Scores:", scores);
}