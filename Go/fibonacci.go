package main

import "fmt"

func fibonacci(n int) int{
	if n == 0{
		return 1
	}
	if n == 1{
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Print("Enter a number: ")
	var n int
	fmt.Scanf("%d", &n)

	if n >= 0{
		fmt.Println("Fibonacci de", n, "es" , fibonacci(n))
	}else{
		fmt.Println("Choose another number")
	}
}