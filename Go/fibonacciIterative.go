package main

import "fmt"

func fibonacciIterative(n int) int {
	if n == 0{
		return 1
	}
	if n == 1{
		return 1
	}
	first, second, tmp := 0,1,0
	for i:=0; i < n; i++{
		tmp = first + second
		first = second
		second = tmp
	}
	return second
}

func main() {
	fmt.Print("Enter a number: ")
	var n int
	fmt.Scanf("%d", &n)

	if n >= 0{
		fmt.Println("Fibonacci de", n, "es" , fibonacciIterative(n))
	}else{
		fmt.Println("Choose another number")
	}
}