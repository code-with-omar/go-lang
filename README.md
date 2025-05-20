## Data types

### Basic Data types

### 1.Numeric Types

- Integer
  - `Signed`: int, int8, int16, int32, int64
  - `Unsigned`: uint, uint8 (alias byte), uint16, uint32, uint64, uintptr
- `Floating-point`: float32, float64
- `Complex numbers`: complex64, complex128
- `Boolean`
  bool (values: true or false)
- String
  - string (immutable sequence of bytes; UTF-8 encoded)

## Variable Declarations

### 1 `var` with type Inference

```golang
var city = "Comilla" // Inference as string
var temparature = 30.5 // Inference as float64
```

### 2. Short Declaration (`:=`)

```golang
message := "Hello, GO!"
count:=10
```

- Note: ⚠️ := can only be used inside functions.

### 3. Declaring Multiple Variables

```golang
var x,y,z int
var p,q,r int =5,6,7
var (
    title string = "Developer"
    salary float64 = 5000.00
    active bool = true
)
```

### 4. Zero Values

```golang
var a int
var b string
var c bool
```

## Conditions

A condition can be either true or false.
Go supports the usual comparison operators from mathematics:

- < Less than

- <= Less than or equal

- `> ` Greater than

- `> =` Greater than or equal

- == Equal to

- != Not equal to

Additionally, Go supports common logical operators:

- && — Logical AND

- || — Logical OR

- ! — Logical NOT

We can use these operators — alone or combined — to create conditions for decision-making.
Example

```golang
x > y
x != y
(x > y) && (y > z)
(x == y) || z

```

### Go Conditional Statements

Go provides the following conditional statements:

- `if` — Specifies a block of code to run if a condition is true.

- `else` — Specifies a block of code to run if the same condition is false.

- `else if` — Specifies a new condition to test if the first condition is false.

- `switch` — Allows selection among multiple possible code blocks.

Examples:

1. if-else

   ```go
   package main

   import "fmt"

   func main() {
       score := 85
       if score >= 90 {
   	    fmt.Println("Grade: A")
       } else if score >= 80 {
   	    fmt.Println("Grade: B")
       } else {
   	    fmt.Println("Grade: C or lower")
       }
   }
   ```

2. Switch

   ```go
   package main
   import "fmt"

   func main() {
       age := 12
       switch {
       case age < 13:
   	    fmt.Println("Child")
       case age < 18:
   	    fmt.Println("Teen")
       default:
   	    fmt.Println("Adult")
       }
   ```

## Go Functions

- A function is a block of statements that can be used repeatedly in a program.

- A function will not execute automatically when a page loads.

- A function will be executed by a call to the function.

Example

```golang
package main

import "fmt"

// function without return
func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println(sum)
}
// function with return
func calculateSum(n1 int, n2 int) int {
	return n1 + n2
}
func main() {
	num1 := 5
	num2 := 10
	add(num1, num2)
	sum := calculateSum(num1,num2)
	fmt.Println(sum)
}
```
