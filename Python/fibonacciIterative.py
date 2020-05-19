import math

def fibonacciIterative(n):
    if(n == 0):
        return 0
    if(n == 1):
        return 1
    first = 0
    second = 1
    for i in range(1,n):
        tmp = first + second
        first = second
        second = tmp
    return second


def main():
    
    n = int(input("Enter a number: "))
    
    if n >=0:
        print(f"Fibonacci de {n} es: {fibonacciIterative(n)}")
    else:
        print("Choose another number")


if __name__ == "__main__":
    main()
