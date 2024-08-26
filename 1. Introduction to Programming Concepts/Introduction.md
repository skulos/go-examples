# Introduction

### 1.1 What is Programming?
---
Programming is the process of writing instructions that a computer can follow to perform specific tasks. These instructions are written in a programming language, which is a set of rules and syntax that allow you to communicate with the computer.

#### Key Concepts:

- **Program:** A sequence of instructions that the computer executes.
- **Programming Language:** A formal language used to write programs. Examples include Go, Python, Java, etc.
- **Compiler vs. Interpreter:** Go is a compiled language, meaning the code you write is converted into machine code before it can run, making it faster than interpreted languages.

```go
package main

func main() {}
```

### 1.2 Writing Your First Go Program
---
A basic Go program has the following structure:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

#### Breakdown:

- **`package main`:** Defines the package to which this code belongs. The `main` package is special because it defines an executable program.
  
- **`import "fmt"`:** Imports the `fmt` package, which includes functions for formatted I/O (like printing to the console).
  
- **`func main() { ... }`:** The `main` function is the entry point of the program. The program starts executing from here.

#### Compile and Run:

- **`go run main.go`**: This command compiles and runs the Go program in a single step.
  
- **`go build -o main.exe main.go`**: This command compiles the program and generates an executable file that you can run later.

#### Exercise:

- Write a Go program that prints "Welcome to Go programming!" to the console.


### 1.3 Variables and Data Types
---
In Go, variables are used to store data, and data types define the type of data that can be stored.

#### Variables in Go:

- Explicit Typing:
  
```go
var name string = "Alice"
```
  
  
- Implicit Typing:
  
```go
age := 30
```


- Multiple Variables:
  
```go
var a, b, c int
a, b, c = 1, 2, 3
```


#### Data Types:

- **Integer Types:**
	- `int`: Default integer type.
	- `int8`, `int16`, `int32`, `int64`: Integers of different sizes.
     
- **Floating-Point Types:**
    - `float32`, `float64`: For numbers with decimals.
      
- **String Type:**
    - `string`: A sequence of characters.
      
- **Boolean Type:**
    - `bool`: Represents true or false.


#### Exercise:

- Declare a variable for your name, age, and whether you are a student or not. Print these values to the console.


### 1.4 Taking User Input
---
Go provides several ways to take user input, with the `fmt.Scan` and `fmt.Scanln` functions being the most common.

```go
package main

import "fmt"

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scan(&name)
    fmt.Println("Hello,", name)
}
```


**`fmt.Scan(&name)`**: Takes input from the user and stores it in the `name` variable.

#### Exercise:

- Write a program that asks the user for their age and prints a message saying "You are ***age*** years old."


### 1.5 Handling Different Data Types
---
When working with user input, you may need to handle different data types.

```go
package main

import "fmt"

func main() {
    var age int
    fmt.Print("Enter your age: ")
    fmt.Scan(&age)
    fmt.Println("You are", age, "years old.")
}
```


#### Handling Multiple Inputs:

```go
package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("Enter your name and age: ")
    fmt.Scan(&name, &age)
    fmt.Println("Your name is", name, "and you are", age, "years old.")
}
```


#### Exercise:

-Modify the previous program to ask for the user's name and age, then print a message that says, "Hello ***name***, you are ***age*** years old."


### 1.6 Simple Arithmetic Operations
---
Arithmetic operations allow you to perform basic mathematical calculations.

```go
package main

import "fmt"

func main() {
    x := 10
    y := 5

    fmt.Println("Addition:", x+y)
    fmt.Println("Subtraction:", x-y)
    fmt.Println("Multiplication:", x*y)
    fmt.Println("Division:", x/y)
}
```


#### Exercise:

- Write a program that asks the user to enter two numbers and then prints the result of adding, subtracting, multiplying, dividing, and finding the modulus of the two numbers.


### 1.7 Exercise
---

Write a Go program that

1. Takes the length and width of a rectangle from the user and calculates the area.
```go
area = l x b
```
  
 
2. Calculate the Perimeter of a Rectangle
```go
perimeter = 2 x (length + width)
```
   
   
3. Convert Celsius to Fahrenheit
```go
Fahrenheit = (Celsius x 9/5) + 32
```
   
   
4. Simple Interest Calculation
```go
Simple Interest = (Principal x Rate x Time) / 100
```
   
   
5. Calculate the Hypotenuse of a Triangle
```go
hypotenuse = sqrt(a^2 + b^2)   
```
   
   
6. Determine Even or Odd Number
