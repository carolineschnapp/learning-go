# My notes on my journey to learn go

![Alt text](https://monosnap.com/image/iihcQXz7takpv0wR235yRbfnPEdHOV.png)

## Every Go program is made up of packages

Programs start running in package “main”—which must contain a `main` function:

```go
package main

// Imported packages here, whatever we need.
import (
  "fmt"
  "math/rand"
)

func main() {
  fmt.Println("My favorite number is", rand.Intn(10))
}
```

Packages are referred to using the last name in their import path, so `fmt` and `rand`.
`fmt` and `rand` are the names of the packages. The "math/rand" package comprises files that begin with the statement `package rand`.

> `math` is a sort of namespace, and the name of a folder?

Yes.

Not only that, each package needs to have its own folder, even if it contains only one file.

Locally, custom packages must be added to `$GOPATH/src` to be found.

```shell
➜ echo $GOPATH
/Users/carolineschnapp/Code/go
```

### Factored import statement

```go
package main

// That's called a factored import statement:
import (
  "fmt"
  "math"
  "math/rand"
)

func main() {
  fmt.Println("My favorite number is", rand.Intn(10))
  fmt.Printf("Now you have %g problems.", math.Sqrt(7))
}
```

## Semicolons

You could end/separate each statement with a semicolon, like in Ruby, but it is frowned upon—unless statements are on the same line, for example when using `for`.

## The fmt package

Package **fmt** implements formatted I/O.

* `Println` formats using the default formats for its operands and writes to **standard output**. Spaces are always added between operands and a newline is appended.
* `Printf` formats according to a format specifier and writes to **standard output**.

Verbs to use with Printf:

* `%f` floating-point number, e.g. `fmt.Printf("%.2f", 8/7)`
* `%d` integer, base 10
* `%t` boolean, outputs the word true or false
* `%s` string
* `%q` a double-quoted string safely escaped with Go syntax

## Exported names have to begin with a capital letter

In Go, a name is exported from a package only if it begins with a capital letter.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

### That's why Println, Printf, Sqrt, and Intn begin with a capital letter

## Type à la Typescript

```go
package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println("The sum of 42 and 13 is", add(42, 13)) // same as
  fmt.Printf("The sum of 42 and 13 is %d", add(42, 13))
}
```

When two or more consecutive named function parameters share a type, you can omit the type from all but the last:

```go
package main

import "fmt"

func add(x, y int) int {
  return x + y
}

func main() {
  fmt.Println("The sum of 42 and 13 is", add(42, 13)) // same as
  fmt.Printf("The sum of 42 and 13 is %d", add(42, 13))
}
```

## A function can return any number of results

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
  return y, x
}

func main() {
  a, b := swap("hello", "world")
  fmt.Println(a, b)
}
```

### Introducing `:=`

The `:=` notation serves both as a declaration and as initialization.

```go
foo := "bar"
```

is equivalent to:

```go
var foo = "bar"
```

### Rappel

`fmt.Println` formats using the default formats for its operands and writes to **standard output**. Spaces are always added between operands and a newline is appended.

## Return values may be named

If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as they can harm readability in longer functions.

All the below are equivalent:

```go
func split(sum int) (int, int) {
  var x = sum * 4 / 9
  var y = sum - x
  return x, y
}
```

```go
func split(sum int) (int, int) {
  // using the `:=` shortcut
  x := sum * 4 / 9
  y := sum - x
  return x, y
}
```

```go
func split(sum int) (x, y int) {
  // it's like we have var x and var y here,
  // by naming the returned values
  x = sum * 4 / 9
  y = sum - x
  return x, y
}
```

```go
func split(sum int) (x, y int) {
  // it's like we have var x and var y here,
  // by naming the returned values
  x = sum * 4 / 9
  y = sum - x
  return // naked return, thanks to the named return values
}
```

### Careful if you have only one named result

If you have a single result, and you name it, and supply a type for it, the name and type must be wrapped in parenthesis.

This is valid:

```go
func do() int {
  stuff := 4
  return stuff
}
```

This is not valid because “undefined: stuff”:

```go
func do() stuff {
  stuff = 4
  return
}
```

This is not valid because “missing function body” and
“syntax error: unexpected int after top level declaration”:

```go
func do() stuff int {
  stuff = 4
  return
}
```

But after we addd the parenthesis, all is good:

```go
func do() (stuff int) {
  stuff = 4
  return
}
```

## Types

The types we have seen so far:

* `int`
* `string`

Strings must be wrapped in double quotes, not single quotes.

### Single quotes are used for rune literals

**UTF-8 is a variable width character encoding capable of encoding all 1,112,064 valid code points in Unicode using 1 _to_ 4 8bit bytes.** (Read again until that sinks in.)

4 * 8 bit = 32 bits, so we need an int32 to store an encoded UTF-8 character.

`a` is a rune literal. A rune is a type, and it is an alias for int32. It occupies 32bit and is meant to represent a Unicode code point. As an analogy, the english characters set encoded in 'ASCII' has 128 code points. Thus is able to fit inside a byte (8bit). We need up to 4 bytes to store a Unicode character. The rune literal 'a' is actuality the number 97.

```go
var rune = 'a'
fmt.Printf("This is a rune: %T %v %s\n", rune, rune, rune)
// => This is a rune: int32 97 %!s(int32=97)
fmt.Printf("Converted to a string: %s\n", string(rune))
// => Converted to a string: a
```

## Variables

The `var` statement declares a list of variables; as in function argument lists, the type is last.

A `var` statement can be at package or function level.

```go
package main

import "fmt"

var c, python, java bool

func main() {
  var i int
  fmt.Println(i, c, python, java)
}
```

This outputs:
​
> 0 false false false

So the default value for `int` is `0`, and the default value for `bool` is `false`.

## Variables with initializers

A `var` declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

```go
package main

import "fmt"

var c, python, java = false, false, "no!"

func main() {
  var i, j = 4, "hello"
  fmt.Println(i, j, c, python, java)
}
```

This outputs:
​
> 4 hello false false no!

## Short variable declarations

**Inside a function**, the `:=` short assignment statement can be used in place of a `var` declaration **with implicit type because you use an initializer**.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the `:=` construct is not available.

### Can you use `:=` outside of a `func`

Nope.

### Can you declare the type with `:=`

Nope. You use an initializer! No need!

### Recap

The `:=` notation serves both as a declaration and as initialization.

```go
foo := "bar"
```

is equivalent to:

```go
var foo = "bar"
```

### Example

Before:

```go
package main

import "fmt"

func main() {
  var i, j int = 1, 2
  var k int = 3
  var c, python, java = true, false, "no!"

  fmt.Println(i, j, k, c, python, java)
}
```

> 1 2 3 true false no!

After:

```go
package main

import "fmt"

func main() {
  i, j := 1, 2
  k := 3
  c, python, java := true, false, "no!"

  fmt.Println(i, j, k, c, python, java)
}
```

> 1 2 3 true false no!

## Types, so far

The types we have seen so far:

* `int`
* `string`
* `bool`

## All the types

* `bool`
* `string`
* `int` `int8` `int16` `int32` `int64` `uint` `uint8` `uint16` `uint32` `uint64` `uintptr`
* `byte` alias for `uint8`
* `rune` alias for `int32`, represents a Unicode code point
* `float32` `float64`
* `complex64` `complex128`

## Variable declarations may be "factored" into blocks, as with import statements

```go
package main

import (
  "fmt"
  "math/cmplx"
)

var (
  ToBe   bool       = false
  MaxInt uint64     = 1<<64 - 1
  z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
  fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
  fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
  fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

> Type: bool Value: false
> Type: uint64 Value: 18446744073709551615
> Type: complex128 Value: (2+3i)

We just learned two new verbs to use with `fmt.Printf`:

* `%v` is the value in a default format
* `%T` is a Go-syntax representation of the type of the value

Also:

* `%%` is a literal percent sign; consumes no value

### Just use `int` if you need an integer

When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

### Just use `float64` if you need a float

Go uses that type by default:

```go
fmt.Printf("Type: %T Value: %v\n", 6.5, 6.5)
// Type: float64 Value: 6.5
```

## Zero values

Variables declared without an explicit initial value are given a zero value.

The zero value is:

* `0` for int,
* `0` for float64,
* `false` for the bool type, and
* `""` (the empty string) for string

```go
package main

import "fmt"

func main() {
  var i int
  var f float64
  var b bool
  var s string
  fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```

> 0 0 false ""

We just learned one new verb to use with `fmt.Printf`:

* `%q` is a a double-quoted string safely escaped with Go syntax

## Type conversions

Type conversions are very handy in a statically-typed language.

Here is how you do it:

```go
var i int = 42
# => 42
var f float64 = float64(i)
# => 42
var s string = string(i)
# => "*"
```

Or using `:=`:

```go
i := 42
f := float64(i)
s := string(i)
```

## Type inference

A type is inferred from the initializer:

```go
package main

import "fmt"

func main() {
  v := 3.142
  fmt.Printf("v is of type %T\n", v)
}
```

> v is of type float64

## Constants

Constants are declared like variables, but with the `const` keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.

```go
package main

import "fmt"

const Pi = 3.14

func main() {
  const World = "世界"
  fmt.Println("Hello", World)
  fmt.Println("Happy", Pi, "Day")

  const Truth = true
  fmt.Println("Go rules?", Truth)
}
```

> Hello 世界
> Happy 3.14 Day
> Go rules? true

## For loop

The basic for loop has three components separated by semicolons:

* the init statement: executed before the first iteration (optional)
* the condition expression: evaluated **before** every iteration
* the post statement: executed at the end of every iteration (optional)

### The if condition and switch statement also have an init statement

The variables declared in the init statement have their scope limited to the for loop — and in the case of `if` or `switch` all the logical branches.

```go
package main

import "fmt"

func main() {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}
```

Just like in `if` statements, parenthesis are optional (even frowned upon!), but `{}` are obligatory.

Without init and post statements, the `for` loop acts like a while loop:

```go
package main

import "fmt"

func main() {
  sum := 1
  for sum < 100 {
    sum += sum
  }
  fmt.Println(sum)
}
// => 128
```

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed like so:

```go
package main

import "fmt"

func main() {
  for {
    fmt.Println("hello")
  }
}
// hello
// hello
// hello
// hello
// hello
// hello
// hello
// hello
// hello
// hello
// hello
// signal: interrupt
```

CTRL-C will stop the infinite loop.

## Back to `fmt`, what if you want to format and _return_ a string

* `fmt.Sprint` formats using the default formats for its operands and **returns** the resulting string. Spaces are added between operands when neither is a string. It behaves a bit like the I/O equivalent `fmt.Println`. However, with `fmt.Println`, spaces are _always_ added between operands and a newline is appended, and nothing is returned. The result is output to standard output.

* `fmt.Sprintf` formats according to a format specifier and returns the resulting string. It is the I/O equivalent of `fmt.Printf`, but returns the resulting string.

## Avant-goût of error handling

```go
package hamming

import "fmt"

// Calculates the Hamming difference between two DNA strands.
func Distance(strandA, strandB string) (distance int, err error) {
  if len(strandA) != len(strandB) {
    err = fmt.Errorf("Strands %s and %s don't have the same length", strandA, strandB)
    return
  }
  for i := 0; i < len(strandA); i++ {
    if strandA[i] != strandB[i] {
      distance++
    }
  }
  return
}
```

In the above code, `fmt.Errorf` formats according to a format specifier and returns an error with the string as message.

If you don't need to format anything, you can just use `errors.New`. For this, you need to import the `errors` package.

```go
package hamming

import "errors"

// Calculates the Hamming difference between two DNA strands.
func Distance(strandA, strandB string) (distance int, err error) {
  if len(strandA) != len(strandB) {
    err = errors.New("Strands don't have the same length")
    return
  }
  for i := 0; i < len(strandA); i++ {
    if strandA[i] != strandB[i] {
      distance++
    }
  }
  return
}
```

## if

Like `for`, the `if` statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if. Variables declared inside an if short statement are also available inside any of the else blocks.

```go
package main

import (
  "fmt"
  "math"
)

func pow(x, n, limit float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  // can't use v here, though
  return lim
}

func main() {
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )
}

```

## Switch

A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in other languages is provided automatically in Go.

Also, like in Ruby, Go's switch cases need not be constants, and the values involved need not be integers.

Because it's the shorthand for an `if else if ...` structure, it can too start with a **init statement**.

```go
package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("Linux.")
  default:
    // freebsd, openbsd,
    // plan9, windows...
    fmt.Printf("%s.", os)
  }
}
// => Go runs on OS X.
```

### Switch with no condition

Switch without a condition is the same as `switch true`.

This construct can be a clean way to write long if-then-else chains.

```go
package main

import (
  "fmt"
  "time"
)

func main() {
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("Good morning!")
  case t.Hour() < 17:
    fmt.Println("Good afternoon.")
  default:
    fmt.Println("Good evening.")
  }
}
```

## Defer

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```go
package main

import "fmt"

func main() {
  defer fmt.Println("world")
  fmt.Println("hello")
}
// hello
// world
```

## Stacking defers

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in **last-in-first-out** order.

```go
package main

import "fmt"

func main() {
  fmt.Println("counting")

  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }

  fmt.Println("done")
}
// Counting
// done
// 9
// 8
// 7
// 6
// 5
// 4
// 3
// 2
// 1
// 0
```

## Pointers

(You're got this. Remember C++? Easy.)

```go
package main

import "fmt"

func main() {
  i, j := 42, 2701

  p := &i         // p points to i
  fmt.Println(*p) // read i through the pointer: 42
  *p = 21         // set i through the pointer
  fmt.Println(i)  // see the new value of i: 21

  p = &j         // point to j (was pointing to i)
  *p = *p / 37   // divide j by 37 through the pointer
  fmt.Println(j) // see the new value of j: 73
}
```

## How to create sort-of objects

A sort of object with properties and methods can be created using `type` with `struct`, like so:

```go
package main

import "fmt"

type Vertex struct {
  X, Y int
  Z string
}

func main() {
  v := Vertex {1, 2, "hello"}
  v.X = 4
  fmt.Println(v.X)
  // => 4
  p := &v
  p.X = 1e9
  fmt.Println(v)
  // => {1000000000 2 hello}
  v1 := Vertex{1, 2, "shit"} // has type Vertex
  v2 := Vertex{X: 1} // Y:0 and Z:"" are implicit
  v3 := Vertex{} // X:0 and Y:0 and Z:"" are implicit
  p = &Vertex{1, 2, "bye"} // has type *Vertex
  fmt.Println(v1, p, v2, v3)
  // => {1 2 shit} &{1 2 bye} {1 0 } {0 0 }
}
```

### Passing a struct by value makes a copy

The method that receives the struct only modifies a copy:

```go
package main

import "fmt"

type Person struct {
  firstName string
  lastName  string
}

func changeName(p Person) {
  p.firstName = "Caro"
}

func main() {
  person := Person{
    firstName: "Caroline",
    lastName: "Schnapp",
  }

  changeName(person)

  fmt.Println(person)
  // => {Caroline Schnapp}
}
```

To modify the original, you need to pass a pointer to the struct:

```go
package main

import "fmt"

type Person struct {
  firstName string
  lastName  string
}

func changeName(p *Person) {
  p.firstName = "Caro"
}

func main() {
  person := Person{
    firstName: "Caroline",
    lastName: "Schnapp",
  }

  changeName(&person)

  fmt.Println(person)
  // => {Caro Schnapp}
}
```

![Alt text](https://monosnap.com/image/7WVXZFlPZsi9kDl2eIQ2NtrCzZh8co.png)

## Arrays

An array has a fixed size, and all elements must be of the same type.

Declared without initialization:

```go
var myArray[5]int
fmt.Println(myArray)
// => [0 0 0 0 0]
```

Declared with initialization:

```go
myOtherArray := [5]int{5, 6, 7, 8, 9}
fmt.Println(myOtherArray)
// => [5 6 7 8 9]
```

You can change the value at an index:

```go
var a[2]string
a[0] = "Hello"
a[1] = "World"
fmt.Println(a[0], a[1])
// => Hello World
fmt.Println(a)
// => [Hello World]
```

## Slices

Slices references subsets of arrays. They are a “view into an array” says Go Tour.

Their type is `[]T`.

They need an array when initialized, or assigned to, unless they are “slice literals”:

```go
package main

import "fmt"

func main() {
  myArray := [6]int{2, 3, 5, 7, 11, 13}

  var slice []int = myArray[1:4]
  fmt.Println(slice)
  // => [3 5 7]
  slice[1] = 8
  fmt.Println(slice)
  // => [3 8 7]
  slice = myArray[2:3]
  fmt.Println(slice)
  // => [8]
  slice[0] = 78
  fmt.Println(slice)
  // => [78]
  fmt.Println(myArray)
  // => [2 3 78 7 11 13]
}
```

![Alt text](https://monosnap.com/image/bzzyrpN0odyEQNSzQCm9rucev69bWZ.png)

End index is **not** included.

Start and end indices are 0-based, of course.

Modify the slice content, and it will modify the referenced array.

### Slices are like **references to arrays**

A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

### Slice literals versus arrays versus count-it-for-me-I-am-lazy arrays

The type of a slice is `[]T`.

The type of an array is `[n]T`.

```go
a := []int{2, 3, 5, 7, 11, 13} // This is a slice literal.
fmt.Printf("Type: %T Value: %v\n", a, a)

b := [6]int{2, 3, 5, 7, 11, 13} // This is an array.
fmt.Printf("Type: %T Value: %v\n", b, b)

c := [...]int{2, 3, 5, 7, 11, 13} // This is an array where we were too lazy to count.
fmt.Printf("Type: %T Value: %v\n", c, c)
```

![Alt text](https://monosnap.com/image/7AQTpa75yndkDgZ31tn4VNiXpo8GS0.png)

### What is a slice literal

A slice literal is like an array literal without the length. It builds an array _and_ returns a slice that references it.

### Slice default bounds

When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

All those are equivalent:

```go
a[0:10]
a[:10]
a[0:]
a[:]
```

### Slicing a slice

⚠️ If you slice a slice, you get less and less elements:

```go
func main() {
  s := []int{2, 3, 5, 7, 11, 13} // Slice literal

  s = s[1:4]
  fmt.Println(s)
  // => [3 5 7]

  s = s[:2]
  fmt.Println(s)
  // => [3 5]

  s = s[1:]
  fmt.Println(s)
  // => [5]
}
```

## I need true randomness

```go
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  // What we will use:
  // - func Seed(seed int64)
  // - func Intn(n int) int
  // - func Now() Time
  // - func (t Time) Unix() int64
  var (
    now time.Time
    unixTime int64
    randNumber int
  )
  now = time.Now()
  unixTime = now.Unix() // Number of seconds since Thursday, January 1st, 1970.
  rand.Seed(unixTime) // To get a different number every time we run the program.
  randNumber = rand.Intn(100) + 1
  fmt.Printf("Random number between 1 and 100 is %d.\n", randNumber)
  // => Random number between 1 and 100 is 48.
}
```

## User input

```go
package main

import (
  "bufio"
  "fmt"
  "log"
  "math/rand"
  "os"
  "strconv"
  "strings"
  "time"
)

func main() {
  // What we will use:
  // - strings: func TrimSpace(s string) string
  // - strconv: func Atoi(s string) (int, error)
  // - io: to access the variable os.Stdin
  // - bufio: func NewReader(rd io.Reader) *Reader
  // - bufio: func (b *Reader) ReadString(delim byte) (string, error)
  // - math/rand: func Seed(seed int64)
  // - math/rand: func Intn(n int) int
  // - time: func Now() Time
  // - time: func (t Time) Unix() int64
  randNumber := randomNumber()
  // fmt.Printf("Random number between 1 and 100 is %d.\n", randNumber)
  // => Random number between 1 and 100 is 48.
  reader := bufio.NewReader(os.Stdin)
  for i := 0; i < 10; i++ {
    guess := guess(reader)
    if guess == randNumber {
      fmt.Println("You have guessed correctly!")
      return
    } else if guess < randNumber {
      fmt.Println("Higher baby!")
    } else {
      fmt.Println("Lower baby!")
    }
  }
  fmt.Println("You have ran out of guesses. Good day to you, Sir!")
}

func randomNumber() (randNumber int) {
  now := time.Now()
  unixTime := now.Unix() // Number of seconds since Thursday, January 1st, 1970.
  rand.Seed(unixTime) // To get a different number every time we run the program.
  randNumber = rand.Intn(100) + 1
  return
}

func guess(reader *bufio.Reader) (guess int) {
  fmt.Println("Can you guess which number it is?")
  value, err := reader.ReadString('\n')
  if err != nil {
  log.Fatal(err)
  }
  trimmedString := strings.TrimSpace(value)
  guess, err = strconv.Atoi(trimmedString)
  if err != nil {
  log.Fatal("You must enter a number!")
  }
  return
}
```

## First and foremost, a slice is a reference to an array

That's its first usefulness. It points to an array.

```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  // An array
  colors := [3]string{"red", "green", "blue"}
  // Printing it
  fmt.Println(colors)
  // => [red green blue]
  // Printing a slice of it
  fmt.Println(colors[:])
  // => [red green blue]
  // Passing slice to method, in order to pass array by reference.
  upperCaseFirstLetter(colors[:])
  // Printing it again
  fmt.Println(colors)
  // => [Red Green Blue]
}

func upperCaseFirstLetter(colors []string) {
  for i := 0; i < len(colors); i++ {
    colors[i] = strings.Title(colors[i])
  }
}
```

The only time you're dealing with an array is when you create it with a size.

Everything else is a slice.

## The make function

The make function allocates a zeroed array and returns a slice that refers to that array.

```go
a := make([]int, 5) // len(a)=5
fmt.Println(a)
// => [0 0 0 0 0]

b := make([]string, 3) // len(a)=3
fmt.Println(b)
// => [   ]
```

## Writing to a file

Writing "Hey Caro" to a file:

```go
func heyCaroToFile() (err error) {
  var file *os.File
  file, err = os.Create("/Users/carolineschnapp/caro.txt")
  if err != nil {
    return
  }

  defer file.Close()

  _, err = io.WriteString(file, "Hey Caro")
  return
}
```

## Reading from a file

Reading each line in a file to a slice, and returning the slice:

```go
package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

func main() {
  var lines, err = readLinesfromFile("/Users/carolineschnapp/caro.txt")
  if err != nil {
    log.Fatal(err)
  }
  for i := 0; i < len(lines); i++ {
    fmt.Printf("Line %d: %s\n", i, lines[i])
  }
}

func readLinesfromFile(filepath string) (lines []string, err error) {
  var file *os.File
  file, err = os.Open(filepath)
  if err != nil {
    return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  if scanner.Err() != nil {
    err = scanner.Err()
  }
  return
}
```

## When using append, you need to look at the returned value

The `append` function works with slices, to append an element at the end of the slice.

It does not mutate the slice, so you need to assign the returned value to the slice.

I know, it's weird: [golang append() evaluated but not used](https://stackoverflow.com/a/45005304/9459399).

Typical append usage is:

```go
a = append(a, x)
```

Example:

```go
scanner := bufio.NewScanner(file)
for scanner.Scan() {
  lines = append(lines, scanner.Text())
}
```

It is a _variadic_ function because it may be invoked with 1 or more elements to append:

```go
a = append(a, x, y, z)
```

If the capacity of a is not large enough to fit the additional values, append allocates a new, sufficiently large underlying array that fits both the existing slice elements and the additional values. Otherwise, append re-uses the underlying array.

## Reading and writing to files best stackoverflow answer

https://stackoverflow.com/a/9739903/9459399

## The bufio package

The package bufio implements buffered I/O.

Buffer say what? [Read this excellent article](https://medium.com/golangspec/introduction-to-bufio-package-in-golang-ad7d1877f762)
