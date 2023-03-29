# go-basics

Code from the course [Go: The Complete Developer's Guide (Golang)](https://www.udemy.com/course/go-the-complete-developers-guide/)

# Notes

## Types

* bool
* string
* int (int8  int16  int32  int64)
* float (float32 float64)
* byte (alias for uint8)
* struct
* array & slice
* map
* chan (channel)

Strings are commonly represented as a slice of bytes, also called byte slice (`[]byte`).

To convert (cast) a variable from one type to the other, we use the syntax `desiredType(currentValue)`. Example:
```go
// Convert from string to byte slice
b := []byte("String to byte slice")

// Convert from byte slice to string
s := string(b)
```

Map is a key-value pair type.
```go
colors := map[string]string{
    "red":   "#ff0000",
    "green": "#4bf745", // Note the trailing comma
}
colors["white"] = "#ffffff"

delete(colors, "red")

color := colors["green"] // Get value: color == "#4bf745"
color = colors["black"]  // Not found: color == "" (zero value)

_, found := colors["green"] // found == true
_, found = colors["black"]  // found == false
```


Use the keyword `type` to create a custom variable type:
```go
type Person struct {
	Name string
	Age  int
}

type myInt int32
```

## Variable declaration

```go
// Declare and assign in two steps
var s string
s = "This is a string"

// Declare and assign in the same step
s := "This is a string" // assumes the type of the assigned value

// Examples of the make function
c := make(chan int)

i1 := make([]int) // []
i2 := make([]int, 10) // [0 0 0 0 0 0 0 0 0 0]
i3 := make([]int, 0, 10) // [] (maximum capacity is 10 elements)

m := make(map[string]int)
```

Note for the slice declarations above, from the documentation: `The capacity of the slice is equal to its length. A second integer argument may be provided to specify a different capacity; it must be no smaller than the length. For example, make([]int, 0, 10) allocates an underlying array of size 10 and returns a slice of length 0 and capacity 10 that is backed by this underlying array.`

## Conditionals

If / else syntax:
```go
if someCondition && someOtherCondition {
    // do something
} else if someOtherCondition || someOtherOtherCondition {
    // do something different
} else {
    // do something else
}
```

Switch case syntax:
```go
switch myVar {
case 1:
    // do something
case 2:
    // do something
default:
    // do something
}
```
Note: differently from other languages, we don't use break statements here. Go will stop checking for matching patterns after the first successful match.

## Loops

C-like loop syntax:
```go
for i := 0; i < 10; i++ {
    println(i)
}
```

Ranging over a slice:
```go
for i, e := range []string{"a", "b", "c"} {
    println(i, e)
}
```

Ranging over a map
```go
for k, v := range map[string]string{"a": "b", "c": "d"} {
    println(k, v)
}
```

Infinite loop:
```go
for  {
    // do something
}
```

Using a loop to capture values from a channel:
```go
c := make(chan int)

go func(ch chan int) {
    for i := 0; i < 10; i++ {
        time.Sleep(1 * time.Second)
        ch <- i
    }
    close(ch)
}(c)

for n := range c {
    fmt.Println(n)
}
```

## Interfaces

In the example below, the Person type can be passed as an argument to the `sayMyName` function that expects to receive a `hasName` interface. Person belongs to this interface because it implements the `getName` function with the exact same defined signature.

Any other type that implements the `getName` function will also be accepted as argument in the `sayMyName` function, and this is why interfaces are useful for code reusability.

```go
package main

import "fmt"

type hasName interface {
	getName() string
}

type Person struct {
	name string
}

func main() {
	var p Person
	p.name = "John"
	sayMyName(&p)
}

func (p *Person) getName() string {
	return p.name
}

func sayMyName(o hasName) {
	fmt.Println("My name is", o.getName())
}

```

Interfaces can also reference other interfaces.
```go
type DoThisAndThat interface {
	DoThis
    DoThat
}
```

## Go routines

Go routines are useful for concurrent (not necessarily paralel) task execution. It also leverages channels to send and receive data to and from other parts of the code. In the example below, the probability is that the number 2 will be displayed before the number 1, even though the order of the function calls are 1 and 2. This happens because when the `printOne` functions hits a blocking code (in this case a sleep function), the program will continue with whatever else can be run while waiting for the completion of this line. The execution of this function resumes later.

We use the `go` keyword to runa function in a separated go routine.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go printOne(c)
	go printTwo(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func printOne(c chan int) {
	time.Sleep(1 * time.Second)
	c <- 1
}

func printTwo(c chan int) {
	c <- 2
}
```

## JSON

Example of a code that marshal / unmarshal JSON from / to Go types.

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	PostalCode    string `json:"postal_code"`
}

type Person struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Age       int     `json:"age"`
	Address   Address `json:"address"`
}

func main() {

	myJson := `[{
		"first_name": "John",
		"last_name": "Smith",
		"age": 25,
		"address": {
			"street_address": "21 2nd Street",
			"city": "New York",
			"state": "NY",
			"postal_code": "10021"
		}
	},
	{
		"first_name": "John",
		"last_name": "Smith",
		"age": 25,
		"address": {
			"street_address": "21 2nd Street",
			"city": "New York",
			"state": "NY",
			"postal_code": "10021"
			}
	}]`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		panic(err)
	}
	fmt.Println(unmarshalled)

	marshalled, err := json.Marshal(&unmarshalled)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshalled))
}

```
## Tests

To run tests in Go, create a `_test.go` file next to the file with the code to be tested. The names need to match.

A basic structure of a test looks like this:

```go
package main

import "testing"

func TestSomething(t *testing.T) {
	got, err := funcOne()

	// Check for errors
	if err != nil {
		t.Errorf("funcCall() returned error: %v", err)
	}

	// Check for expected value
	if got != expected {
		t.Errorf("funcCall() = %v; want %v", got, expected)
	}
}
```

To run multiple test cases against a function, we can use a slice of values for test cases. Example is below:
```go
package main

import "testing"

var tests = []struct {
	x, y, expected int
}{
	{1, 2, 3},
	{2, 3, 5},
	{3, 4, 7},
}

func TestSum(t *testing.T) {
	for _, tt := range tests {
		if sum(tt.x, tt.y) != tt.expected {
			t.Errorf("Expected %d, got %d", tt.expected, sum(tt.x, tt.y))
		}
	}
}
```

One liner assignation and check:
```go
if got, want := sumExample(1,2), 3; got != want {
	t.Errorf("sumExample func returned %d; want %d", got, want)
}
```