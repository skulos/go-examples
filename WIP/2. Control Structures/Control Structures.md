## 2. Control Structures
---
Understand how to use conditional statements to control the flow of code execution.

- [[#2.1 If Statements]]
- [[#2.2 Nested and Complex If Statements]]
- [[#2.3 Loops in Go]]
- [[#2.4 Functions with Return Values]]
- [[#2.5 Exercise]]


### 2.1 If Statements
---
An `if` statement executes a block of code if a condition is *true*. An `else` statement can be used to execute a block of code if the condition is *false*.

```go
package main

import "fmt"

func main() {
    var age int
    fmt.Print("Enter your age: ")
    fmt.Scan(&age)

    if age < 18 {
        fmt.Println("You are a minor.")
    } else if age < 21 {
        fmt.Println("You are a young adult.")
    } else {
        fmt.Println("You are an adult.")
    }
}
```

#### Explanation:

- **`if age < 18`**: Checks if the age is less than 18.
- **`else if age < 21`**: Checks if the age is between 18 and 20.
- **`else`**: Catches all other cases.

#### Exercise:

- Write a program that checks if a number is positive, negative, or zero.


### 2.2 Nested and Complex If Statements
---
You can nest `if` statements inside other `if` statements to handle multiple layers of conditions.

```go
package main

import "fmt"

func main() {
    var age int
    var isStudent bool

    fmt.Print("Enter your age: ")
    fmt.Scan(&age)
    fmt.Print("Are you a student (true/false)? ")
    fmt.Scan(&isStudent)

    if age < 18 {
        if isStudent {
            fmt.Println("You are a minor student.")
        } else {
            fmt.Println("You are a minor and not a student.")
        }
    } else if age < 21 {
        if isStudent {
            fmt.Println("You are a young adult student.")
        } else {
            fmt.Println("You are a young adult and not a student.")
        }
    } else {
        if isStudent {
            fmt.Println("You are an adult student.")
        } else {
            fmt.Println("You are an adult and not a student.")
        }
    }
}
```


#### Logical Operators:

Combine multiple conditions using logical operators (`&&` for AND, `||` for OR).

```go
package main

import "fmt"

func main() {
    var age int
    var hasTicket bool

    fmt.Print("Enter your age: ")
    fmt.Scan(&age)
    fmt.Print("Do you have a ticket (true/false)? ")
    fmt.Scan(&hasTicket)

    if age >= 18 && hasTicket {
        fmt.Println("You are allowed entry.")
    } else {
        fmt.Println("You are not allowed entry.")
    }
}
```


#### Exercise:

- Write a program that checks if a user can enter a club based on their age and membership status. The user must be at least 21 years old or have a membership to enter.



### 2.3 Loops in Go
---

#### For Loop

The `for` loop is the only loop construct in Go. It can be used in several ways:

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        fmt.Println("Count:", i)
    }
}
```


#### For Loop with Conditional

You can also use `for` loops with just a condition (similar to a `while` loop in other languages).

```go
package main

import "fmt"

func main() {
    var i int = 1
    for i <= 5 {
        fmt.Println("Count:", i)
        i++
    }
}
```


#### Range Loops

The `range` keyword can be used to iterate over slices, arrays, maps, and strings.

```go
package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    
    for i, num := range numbers {
        fmt.Println("Index:", i, "Value:", num)
    }
}
```


#### Exercise:

- Write a program that prints all even numbers from 1 to 20 using a `for` loop.


### 2.4 Functions with Return Values
---
Define and use functions that return values.

#### Function Syntax

Functions can return values using the `return` statement. Functions must declare their return types.

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    sum := add(3, 4)
    fmt.Println("Sum:", sum)
}
```


#### Multiple Return Values

Functions can return multiple values.

```go
package main

import "fmt"

func divide(a int, b int) (int, int) {
    return a / b, a % b
}

func main() {
    quotient, remainder := divide(10, 3)
    fmt.Println("Quotient:", quotient)
    fmt.Println("Remainder:", remainder)
}
```


#### Exercise:

- Write a function that calculates the maximum of two numbers and returns the result. Use this function in your `main` program.


### 2.5 Exercise
---

##### **Exercise 1:**

- Write a program that prints numbers from 1 to 100. For multiples of 3, print "Fizz"; for multiples of 5, print "Buzz"; and for numbers that are multiples of both 3 and 5, print "FizzBuzz".

##### **Exercise 2:**

- Write a program that takes a number from the user and prints all factors of that number.

##### **Exercise 3:**

- Write a program that uses a `for` loop to calculate and print the factorial of a number provided by the user.

