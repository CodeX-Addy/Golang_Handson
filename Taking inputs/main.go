package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	hey := "Enter the rating for our pizza..";
	fmt.Println(hey)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating:", input)
}
