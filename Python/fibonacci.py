import math

def fibonacci(n):
    if(n == 0):
        return 0
    if(n == 1):
        return 1
    else:
        return fibonacci(n-1) + fibonacci(n-2)


def main():
    
    n = int(input("Enter a number: "))
    
    if n >=0:
        print(f"Fibonacci de {n} es: {fibonacci(n)}")
    else:
        print("Choose another number")


if __name__ == "__main__":
    main()
