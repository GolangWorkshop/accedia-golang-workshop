# Go Workshop

# Table of contents

- [**OOP**](#oop)
  - [Interfaces](#interfaces)
  - [Generics](#generics)
  - [OOP Concepts in Go](#oop-concepts-in-go)
    - [Abstraction](#abstraction)
    - [Encapsulation](#encapsulation)
    - [Inheritance](#inheritance)
    - [Polymorphism](#polymorphism)
- [**Concurrency**](#concurrency)
  - [Goroutines](#goroutines)
  - [Context](#context)
  - [Channels](#channels)
    - [Buffering](#buffering)
    - [Synchronization](#synchronization)
    - [Wait Groups](#wait-groups)
    - [Channel axioms](#channel-axioms)
- [**Errors**](#errors)
  - [Handling errors](#handling-errors)
  - [Panic and Recover](#panic-and-recover)
  - [Custom error types](#custom-error-types)
  - [Wrapping errors](#wrapping-errors)
  - [Unwrap error chains](#unwrap-error-chains)
- [**IO (streams)**](#io-streams)
  - [Reading](#reading)
    - [Reading from a file](#reading-from-a-file)
    - [Reading the response from a network request](#reading-the-response-from-a-network-request)
  - [Writing](#writing)
    - [Writing to a file](#writing-to-a-file)
    - [Complex example](#complex-example)
- [**CGO**](#cgo)
- [**Testing**](#testing)
  - [Single-case test](#single-case-test)
  - [Multiple-cases test](#multiple-cases-test)
  - [Running tests in parallel](#running-tests-in-parallel)
- [**Extra readings**](#extra-readings)
  - [Online Guides and Articles](#online-guides-and-articles)
  - [Books](#books)
  - [Videos](#videos)

# OOP

## Interfaces

Interfaces in Golang are named collections of method signatures. They are implemented implicitly, that is, as long as the struct implements all the methods of an interface, that interface is implemented:

```go
type Human struct {
  Name string
  age int
}

type Speaker interface {
  Speak(words string)
}

func (h *Human) Speak(words string) {
  fmt.Printf("%s says \"%s\"", h.Name, words)
}

func main() {
  var s Speaker = &Human{"Stamat", 40}
	s.Speak("Hello")
}
```

Since `Human` implements all the methods that the `Speaker` interface requires, any variable holding a pointer to a `Human` struct can be considered a `Speaker`

- To ensure [encapsulation](#encapsulation) it's common to export only the interface and a construction function:

```go
func NewSpeaker(Name string, Age int) (&Speaker, error) {
  h := Human{Name, Age}
  return *h             // implicitly cast to Speaker
}
```

## Generics

Generics, as the name implies, are a way to create generic function signatures in golang. They are available after v1.18:

```go
type MyNumber interface {
  int64 | float64
}

func Sum[T MyNumber] (numbers T[]) T {
  var s T
  for _; val := range numbers {
    s += val
  }

  return s
}
```

In the example the `Sum` function can work with both `int64` and `float64` values, without the need to define concrete implementations.

List of already defined constraints is available [here](https://pkg.go.dev/golang.org/x/exp/constraints).

Structs also support generics:

```go

type UserID int64

type Generic interface {
  ~int64 | ~float64       // ~ means alias, so UserID can also be used
}

type Data[T Generic] struct {
  Field: string
  OtherField: T
}

// ...

d := Data[int64]{Field: "test", OtherField: 12}
```

## OOP Concepts in Go

Go is a [post-OOP programming language](#go-overview), but object-oriented patterns are used for structuring a program in a clear and understandable way

### Abstraction

Abstraction (handle complexity by hiding unnecessary details) is achieved by implementing interfaces and hiding the complex logic under them.

### Encapsulation

Encapsulation (restricting access to data) is mainly achieved by using [scopes](#scopes) and [closures](#closures)

### Inheritance

Go does not have inheritance. Instead, using composition, base structs are included into a derived struct via an anonymous field. Including an anonymous field in a struct is also known as **embedding**:

```go
//base class
type Discount struct {
    percent float32
}

type PremiumDiscount struct{
    Discount //embedding
    additional float32
}
```

Here, the methods of the embedded struct can be called directly. The initialization of embedded structs is done via composite literals. The embedded type serves as the field name.

```go
func (d *Discount) Calculate(originalPrice float64) float64 {
  return originalPrice - originalPrice*float64(d.percent)/100
}

func main() {
  price := 100.00
  d := PremiumDiscount{Discount: Discount{percent: 25}, additional: 10}
  fmt.Println(d.Calculate(price))
}
```

Values of the embedded class can be referenced directly in the deriveds functions:

```go
func (pd *PremiumDiscount) CalculateAdditional(originalPrice float64) float64 {
  // percent belongs to embedded Discount struct
  // additional belongs to PremiumDiscount struct
  return originalPrice - originalPrice*float64(pd.percent+pd.additional)/100
}
```

Additionally, values of the embedded struct can be accessed by the full path using the base struct name

```go
func (pd *PremiumDiscount) CalculateAdditional(originalPrice float64) float64 {
  //percent uses full path
  return originalPrice - originalPrice*float64(pd.Discount.percent+pd.additional)/100
}
```

A struct can have many embedded fields

```go
type Swimmable struct {
}

type Flyable struct {
}

type Walkable struct {
}

type Duck struct {
  Swimmable
  Flyable
  Walkable
}
```

### Polymorphism

Polymorphism (the ability of a something to be displayed in more than one form) in OOP languages has two variations:

1. Runtime Polymorphism: a function call to the overridden method is resolved at runtime. This is possible in Golang and is achieved through [interfaces](#interfaces). Once an interface implements a type, the functionality defined within it is open to any values of that type

```go
type Country interface {
  getCapital() string
}

// a structure that has the required by the interface field
type Bulgaria struct {
  Capital string
}

// implementing methods of interface 'Country'
// for structure 'Bulgaria'
func (bulgaria Bulgaria) getCapital() string {
  return bulgaria.Capital
}

// a second structure that has the required by the interface field
type Germany struct {
  Capital string
}

// implementing methods of interface 'Country'
// for structure 'Germany'
func (germany Germany) getCapital() string {
  return germany.Capital
}

func main() {

  // creating an instance of interface 'Country'
  var ICountry Country

  bulgaria := Bulgaria{Capital: "Sofia"}
  germany := Germany{Capital: "Berlin"}

  // assigning object 'bulgaria' to 'ICountry'
  // and invoking getCapital()
  ICountry = bulgaria
  fmt.Println(ICountry.getCapital()) // this prints Sofia in the console

  // assigning object 'germany' to 'ICountry'
  // and invoking getDevelopedBy()
  ICountry = germany
  fmt.Println(ICountry.getCapital()) // this prints Berlin in the console
}
```

2. Compile time (static) Polymorphism: function or operator overloading; the compiler knows which function is going to be used
   This is **not possible** in Go since there is no method or operator overloading (although there are workarounds as using [variadic function](#variadic-parameters))

## Concurrency

### Goroutines

A goroutine is a lightweight thread of execution managed by the runtime. They can run in parallel if there are resources available. That is why Golang is considered `concurrent` but not always `parallel`. For more info check this talk by Rob Pike: [Concurrency is not Parallelism](https://www.youtube.com/watch?v=oV9rvDllKEg)

`go f(a, b)` starts a new goroutine that runs `f`

```go
// just a function (which can be later started as a goroutine)
func doStuff(s string) {
}

func main() {
    // using a named function in a goroutine
    go doStuff("Hello")
}
```

Characteristics:

- Cheaper than threads
- Stored in the stack and the size of the stack can grow and shrink, according to the requirement of the program (in threads, the size of the stack is fixed)
- Communicate using channels and these channels are specially designed to prevent race conditions when accessing shared memory using Goroutines
- Usage:

  - Suppose a program has one thread, and that thread has many Goroutines associated with it
  - If any of Goroutine blocks the thread due to resource requirement then all the remaining Goroutines will assign to a newly created OS thread
  - All these details are hidden from the programmers

- Anonymous goroutines - an anonymous function that runs on a separate goroutine from which it was invoked

```go
go func(name string) {
      fmt.Println("Welcome to", name)
   }("our workshop")
```

### Context

Contexts in Golang are objects that carry information about deadlines, cancellation signals (ctrl+c), and other request-scoped values across API boundaries and between processes.

They are most often used to signal to spanwed goroutines that something has happened to the root program.

`context.Context` interface, which defines a set of methods for working with a context

- Methods:

  - `Deadline()` - returns the time at which the context will be canceled, if any
  - `Done()` - returns a channel, that is closed when the context is canceled
  - `Err()` - returns an error indicating why the context was canceled, if any
  - `Value(key interface{})` - returns the value associated with key in the context, or `nil` if no value is associated with key
  - `context.WithCancel(parent Context)` - returns a new context and a cancel function that can be used to cancel the context.

    - When the cancel function is called, the context's `Done()` channel will be closed
    - Any goroutines or API calls, that are listening on that channel, will receive a cancellation signal
    - A regular check on the `ctx` cancelation is required

    ```go
    func main() {
        ctx, cancel := context.WithCancel(context.Background())
        go func() {
            time.Sleep(5 * time.Second)
            cancel()
        }()
        err := doSomething(ctx)
        if err != nil {
          fmt.Println(err)
        }
    }
    ```

  - `context.WithDeadline(parent Context, deadline time.Time)` - function returns a new context and a cancel function that can be used to cancel the context after the specified deadline has passed
  - `context.WithTimeout(parent Context, timeout time.Duration)`

    - Similar to `WithDeadline`, but the deadline is calculated as the current time plus the specified timeout
    - Will return error if the function has not finished in the time being passed

    ```go
    func main() {
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        err := doSomething(ctx)
        if err != nil {
          fmt.Println(err)
    }
    ```

  - `context.WithTimeout(parent Context, timeout time.Duration)`
    - Similar to `WithDeadline`, but it takes a duration, instead of an absolute time
    - Returns a copy of the parent context with a deadline set to the current time plus the specified duration
    - Returns error if the function has not finished in the time being passed
  - `context.WithValue(parent Context, key, val interface{})` - function returns a new context with the specified key-value pair added to its context. This allows request-scoped values to be propagated across API boundaries.

  ```go
  func doSomething(ctx context.Context) {
      if val, ok := ctx.Value("key").(string); ok {
          fmt.Println(val)
      }

  func main() {
        ctx := context.WithValue(context.Background(), "key", "value")
        doSomething(ctx)
  }
  ```

### Channels

Channels are the communication pipes that connect concurrent goroutines. They provide a way for a goroutine to send and receive data to other goroutines. Here are some common operations on channels:

- Create a channel:

  ```go
  messages := make(chan string)
  ```

- Send a value into the channel:

  ```go
  go func() { messages <- "ping" }()
  ```

- Receive the value (in other goroutine or the same one)

  ```go
  msg := <-messages
  ```

- Close a channel

  - No more values will be sent on it
  - It’s possible to close a non-empty channel but still have the remaining values be received

  ```go
  close(messages)
  ```

- Range over channels

  ```go
  func main() {
      queue := make(chan string, 2)

      queue <- "one"
      queue <- "two"

      close(queue)

      for elem := range queue {
          fmt.Println(elem)
      }
  }
  ```

- Wait on multiple channel operations with `select`

  ```go
  func main() {
      c1 := make(chan string)
      c2 := make(chan string)

      //ach channel receives a value after some amount of time, to simulate blocking operations executing in concurrent goroutines.
      go func() {
          time.Sleep(1 * time.Second)
          c1 <- "one"
      }()
      go func() {
          time.Sleep(2 * time.Second)
          c2 <- "two"
      }()

      //await both of these values simultaneously, printing each one as it arrives.
      for i := 0; i < 2; i++ {
          select {
          case msg1 := <-c1:
              fmt.Println("received", msg1)
          case msg2 := <-c2:
              fmt.Println("received", msg2)
          }
      }
  }
  ```

#### Buffering

By default channels are unbuffered, meaning that for a message to be sent there needs to be someone on the other end to read it, or in other words, if there is no one to read the message, the send operation will block. This can be omitted by creating a buffered channel:

```go
messages := make(chan string, 2)

messages <- "buffered"
messages <- "channel"

fmt.Println(<-messages)
fmt.Println(<-messages)
```

#### Synchronization

Using channels, we can force a goroutine to wait for another one:

```go
func worker(done chan bool) {
    fmt.Print("workshop...")
    time.Sleep(time.Second)
    fmt.Println("is going")

    //send a value to notify that we’re done.
    done <- true
}

func main() {
    //make goroutine, giving it the channel to notify on.
    done := make(chan bool, 1)
    go worker(done)
    //Block until we receive a notification from the worker on the channel.
    <-done
    //If we remove it the main goroutine will finish before even starting worker goroutine

}
```

-Mutual exclusion lock (mutex)

```go
var m sync.Mutex
```

- Used to provide a locking mechanism
- Makes sure only one goroutine can access a variable at a time to avoid conflicts
- Available methods: `Lock` and `Unlock`
  - Calling `m.Lock` will “lock” the mutex
  - If any other goroutine calls `m.Lock`, it will block the thread until `m.Unlock` is called
  - If a goroutine calls `m.Lock` before its first read/write access to the relevant data, and calls `m.Unlock` after its last, it is guaranteed that between this period, the goroutine will have exclusive access to the data

```go
func main() {
	n := 0
	var m sync.Mutex

	// now, both goroutines call m.Lock() before accessing `n`
	// and call m.Unlock once they are done
	go func() {
		m.Lock()

		// hold a lock for the entire scope of a function
		defer m.Unlock()

		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
	}()

	go func() {
		m.Lock()
		n++
		m.Unlock()
	}()

	time.Sleep(time.Second)
}
```

#### Wait Groups

Used to wait for multiple goroutines to finish. The Goroutine that's called has the responsibility to signal that the work is done

```go
wg := sync.WaitGroup{}
```

Methods:

- `wg.Add(1)`
  - Whenever called, it increases the counter by the parameter specified
  - Should not used more than onece
  - If decreasing `wg` in goroutine and it's not 0, this will cause a deadlock
- `wg.Done()` - decreases `wg` by 1
- `wg.Wait()` - blocks the main goroutine and it will be released until the counter is 0

```go
import (
    "fmt"
    "sync"
)

func main() {

   //declare a variable wg and instantiate a new sync.WaitGroup{}
    wg := sync.WaitGroup{}

    //call Add(1) before attempting to execute our go print().
    wg.Add(1)

    //use Done() function once the print task is completed, see print()
    go print(&wg)

    //block main goroutine until print() is done
    wg.Wait()

    fmt.Println("hello from main")
}

func print(wg *sync.WaitGroup) {
    fmt.Println("hello from goroutine")
    wg.Done()
}
```

#### Channel axioms

1. **A send to a nil channel blocks forever**
   - Deadlock, because the zero value for an uninitalised channel is `nil`

```go
var c chan string
c <- "Hello, World!" // deadlock
```

2. **A receive from a nil channel blocks forever**

```go
var c chan string
fmt.Println(<-c) // deadlock
```

3. **A send to a closed channel panics**

```go
var c = make(chan string, 1)
c <- "Hello, World!"
close(c)
c <- "Hello, Panic!" // panic
```

4. **A receive from a closed channel returns the zero value immediately**
   - Once a channel is closed, and all values drained from its buffer, the channel will always return zero values immediately

```go
var c = make(chan int, 2)
c <- 1
c <- 2
close(c)
for i := 0; i < 3; i++ {
fmt.Printf("%d ", <-c) // 1 2 0
}
```

## Errors

There is no exception handling per se, instead, errors are represented as values of the `error` type, which is a predeclared interface type. Error is anything that implements the `Error()` method, defined by the interface.

```go
type error interface {
    Error() string
}
```

### Handling errors

```go
result, err := someFunction()
if err != nil {
    // handle the error
}

```

### Panic and Recover

Typically, `panic` is used when something has gone unexpectedly wrong or some integral part of the program is not able to operate. Calling the `panic` function will stop the ordinary flow of control and wil enter the panicking mode, which will eventually print the error passed to the `panic` function and stop the program.

The `recover` function regains control of a panicking goroutine. Here are a few important characteristics:

- Useful only in inside deferred functions
- During normal execution, a call to recover will return `nil` and have no other effect
- If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution

Example:

```go
package main

import "fmt"

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}
```

In the example above:

- The function `g` takes the int `i`
- Panics if `i` is greater than 3
- Otherwise it calls itself with the argument `i + 1`
- The function `f` defers a function that calls recover and prints the recovered value (if it is non-nil).

- Output:

```go
Calling g.
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking!
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
Recovered in f 4
Returned normally from f.
```

### Custom error types

Golang's standart library exports a constructor function `errors.New()` that returns an error.

```go
package main

import (
    "errors"
    "fmt"
)
// Main function
func main() {
    err := errors.New("Sample Error")
    if err != nil {
        fmt.Print(err)
    }
}
```

### Wrapping errors

`errors.Wrap()` allows us to wrap errors and provide additional context

```go
import "errors"

func someFunction() error {
    err := doSomething()
    if err != nil {
        return errors.Wrap(err, "someFunction failed")
    }
    return nil
}

func main() {
    err := someFunction()
    if err != nil {
        // prints "someFunction failed: original error message"
        fmt.Println(err.Error())
    }
}
```

In the example above:

- `someFunction` calls `doSomething` and uses `errors.Wrap` to wrap any errors returned
- The `main` function checks for errors returned by `someFunction` and calls `err.Error` which returns the error message

### Unwrap error chains

`errors.Is()` allows the detection of a specific error

```go
package main

import (
    "errors"
    "fmt"
)

const badInput = "abc"

var ErrBadInput = errors.New("bad input")

func validateInput(input string) error {
    if input == badInput {
        return fmt.Errorf("validateInput: %w", ErrBadInput)
    }
    return nil
}

func main() {
    input := badInput

    err := validateInput(input)
    if errors.Is(err, ErrBadInput) {
        fmt.Println("bad input error")
    }
}
```

- `errors.As()`
  - Checks if any error in the chain of wrapped errors matches the target
  - Difference:
    - `As` checks whether the error has a specific type
    - `Is` examines if it is a particular error object
  - Because `As` considers the whole chain of errors, it should be preferable to the type assertion

```go
package main

import (
    "errors"
    "fmt"
)

const badInput = "abc"

type BadInputError struct {
    input string
}

func (e *BadInputError) Error() string {
    return fmt.Sprintf("bad input: %s", e.input)
}

func validateInput(input string) error {
    if input == badInput {
        return fmt.Errorf("validateInput: %w", &BadInputError{input: input})
    }
    return nil
}

func main() {
    input := badInput

    err := validateInput(input)
    var badInputErr *BadInputError
    if errors.As(err, &badInputErr) {
        fmt.Printf("bad input error occured: %s\n", badInputErr)
    }
}
```

## IO (streams)

IO operations in golang are done using 2 core interfaces that the io package exposes: `io.Reader` - something that we can read from & `io.Writer` - something that we can write to. There are many other useful interfaces provided by the standard library, but these two are at the core of it all.

Typing `go doc io.Reader` in a terminal yields the following result:

```go
package io // import "io"
type Reader interface {
	Read(p []byte) (n int, err error)
}
    Reader is the interface that wraps the basic Read method.
    Read reads up to len(p) bytes into p. It returns the number of bytes read (0
    <= n <= len(p)) and any error encountered.
    ...
```

Typing `go doc io.Writer` in a terminal yields the following result:

`go doc io.Writer`

```go
package io // import "io"
type Writer interface {
	Write(p []byte) (n int, err error)
}
    Writer is the interface that wraps the basic Write method.
    Write writes len(p) bytes from p to the underlying data stream. It returns
    the number of bytes written from p (0 <= n <= len(p)) and any error
    encountered that caused the write to stop early.
```

### Reading

#### Reading from a file

```go
package main

import (
  "fmt"
  "io"
  "os"
)

func main() {
  f, err := os.Open("README.md") // opens a file for read only
  if err != nil {
    panic(err) // handle better in prod
  }
  defer f.Close() // do you remember what defer does?

  buf := make([]byte, 32) // 32b buffer, will store the data we read

  for {
    n, err := f.Read(buf) // fills the buffer with as much data as can be read
    if err == io.EOF {
      fmt.Println("EOF reached")
      break
    }
    // handle any other error we might get
    if err != nil {
      fmt.Println(err)
      break
    }
    if n > 0 {
      fmt.Println(string(buf[:n])) // print the bytes from 0 to n casted as string
    }
  }
}
```

#### Reading the response from a network request

```go
package main

import (
  "fmt"
  "io"
  "net"
)

func main() {
  conn, err := net.Dial("tcp", "google.com:80")
  if err != nil {
    panic(err) // handle better in prod
  }
  defer conn.Close()

  conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))

  buf := make([]byte, 100)

  for {
    n, err := conn.Read(buf) // fills the buffer with as much data as can be read
    if err == io.EOF {
      fmt.Println("EOF reached")
      break
    }
    // handle any other error we might get
    if err != nil {
      fmt.Println(err)
      break
    }
    if n > 0 {
      fmt.Println(string(buf[:n])) // print the bytes from 0 to n casted as string
    }
  }
}
```

The code above can be abstracted easily to handle both examples and that is because of the `io.Reader` interface.

### Writing

#### Writing to a file

```go
func writeToFile() {
	f, err := os.OpenFile("file.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600) // write only, create if missing, truncate existing content
	if err != nil {
		panic(err)
	}

	defer f.Close() // do you remember what defer does?

	f.Write([]byte("write this data into the file\n"))
}
```

#### Complex example

Here is a more complex task that will make use of both readers and writers:

- download a gzipped file
- display how many mb of data we have downloaded
- uncompress the data
- write the uncompressed data to a file

```go
type counter struct {
  total uint64
}

// Write
func (c *counter) Write(b []byte) (int, error) {
  c.total += uint64(len(b)) // 32kb at a time
  progress := float64(c.total) / (1024 * 1024)
  fmt.Printf("\rDownloading %f MB...", progress)
  return len(b), nil
}

// big gzipped file
// http://storage.googleapis.com/books/ngrams/books/googlebooks-eng-all-5gram-20120701-0.gz

func completeExample() {
  res, err := http.Get("http://storage.googleapis.com/books/ngrams/books/googlebooks-eng-all-5gram-20120701-0.gz")
  if err != nil {
    panic(err)
  }
  // download the file into our local fs
  local, err := os.OpenFile("download-5gram.txt", os.O_CREATE|os.O_WRONLY, 0600)
  if err != nil {
    panic(err)
  }
  defer local.Close()

  dec, err := gzip.NewReader(res.Body)
  if err != nil {
    panic(err)
  }
  // copy res.Body into local file
  if _, err := io.Copy(local, io.TeeReader(dec, &counter{})); err != nil {
    panic(err)
  }
}
```

For more information, check the [`tee` linux command](https://www.geeksforgeeks.org/tee-command-linux-example/) and use the `go doc io.TeeReader` to read more about this particular implementation

## CGO

CGO is a tool that allows Go programs to call C functions and use C libraries.

- Usage:
  - when there is a need to interface with existing C code
  - when performance is critical and certain operations can be implemented more efficiently in C

Common examples:

- Calling a C function:

  ```go
  package main

  // #include <stdio.h>
  // void hello() {
  //     printf("Hello, C!\n");
  // }
  import "C"

  func main() {
      C.hello()
  }
  ```

- Passing arguments to a C function:

  ```go
  //Pass a string argument to a C function.
  package main

  //declare a C function named print_message() that takes a char* argument and prints it to the console.
  // #include <stdio.h>
  // void print_message(char* message) {
  //     printf("%s\n", message);
  // }

  import "C"

  import "unsafe"

  func main() {
      //create a Go string and convert it to a C string using C.CString()
      message := "Hello, C!"
      cmessage := C.CString(message)

      //Pass the C string to the C function and free the memory allocated for it using C.free().
      defer C.free(unsafe.Pointer(cmessage))

      C.print_message(cmessage)
  }
  ```

- Returning a value from a C function:

  ```go
  package main
  // return a value from a C function
  // #include <stdio.h>
  // int square(int x) {
  //     return x * x;
  // }
  import "C"

  import "fmt"

  func main() {
      x := 5
      result := int(C.square(C.int(x)))
      fmt.Printf("The square of %d is %d\n", x, result)
  }
  ```

## Testing

Testing in Golang is done by using the `testing` package supplied by the standard library.

Tests are ran by typing `go test` in a directory. Go will then look for all the files with the name signature `<file_name>\_test.go`.

A test is a simple function that needs to abide by a few rules:

- Signature: `func TestXxxx(t *testing.T)`
- Function name convention: starts with the word `Test`, followed by a suffix whose first word is capitalized, eg `TestFib`
- Accepts a `*testing.T` object as a parameter

Tests can be written in the same package, but it's also a common approach to define a new package for tests so that you can't import variables and functions that are private to the package itself. This is called black-box testing.

For example, given the following function:

```go
package main

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
```

### Single-case test

```go
package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	seqNo := 1
	expected := 1
	actual := fibonacci(seqNo)

	if expected != actual {
		t.Errorf("fibonacci(%d) = %d, got %d", seqNo, expected, actual)
	}
}
```

### Multiple-cases test

With tests being essentially functions that can fail, it's possible to abstract logic away and have very concise tests that do a lot of checks:

```go
func TestFibMultiple(t *testing.T) {
	cases := []struct{ seqNo, expected int }{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{5, 5},
		{10, 55},
		{20, 6765},
		{30, 832040},
		{40, 102334155},
		{45, 1134903170},
		// {45, 1134903170}, // will take about double the time
	}
	for _, tc := range cases {
		actual := fibonacci(tc.seqNo)

		if tc.expected != actual {
			t.Errorf("fibonacci(%d) = %d, got %d", tc.seqNo, tc.expected, actual)
		}
	}
}
```

### Running tests in parallel

While the previous approach does save developer time, it certainly can start to become sluggish as the code base grows. Golang allows tests to be ran in parallel using the `Parallel()` method

This test would take about 24 seconds if ran in sequence, but in parallel that time is reduced to around 6s:

```go
func TestFibMultipleParallel(t *testing.T) {
	cases := []struct{ seqNo, expected int }{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{5, 5},
		{10, 55},
		{20, 6765},
		{30, 832040},
		{40, 102334155},
		{45, 1134903170},
		{45, 1134903170},
		{45, 1134903170},
		{45, 1134903170},
	}
	for n, tc := range cases {
		testCase := tc

		t.Run(fmt.Sprintf("test case %d", n), func(t *testing.T) {
			t.Parallel()
			actual := fibonacci(testCase.seqNo)

			if testCase.expected != actual {
				t.Errorf("fibonacci(%d) = %d, got %d", testCase.seqNo, testCase.expected, actual)
			}
		})
	}
}
```

## Extra readings

### Online Guides and Articles

- [Official documentation](https://go.dev/) : Official documentation, definitely worth checking out the blog section
- [A tour of Go](https://go.dev/tour/welcome/1) : Online sandtool and tutorial by the Go developers explaining the main concepts
- [Go by example](https://gobyexample.com/) : List of code snippets illustrating a wide variety of topics
- [Take your first steps with Go](https://learn.microsoft.com/en-us/training/paths/go-first-steps/) : Introduction to Go by Microsoft
- [Go 101](https://go101.org/) : An online collection of articles, examples and quizzes in Go, can be downloaded as an ebook as well
- [Learn Go in 100 Lines](https://fireship.io/lessons/learn-go-in-100-lines/) : Learn Go in 100 lines
- [Practical Go Lessons](https://www.practical-go-lessons.com/) : Practical Go lessons with examples. Exists also as a book

### Books

- [How to code in Go](https://www.digitalocean.com/community/books/how-to-code-in-go-ebook) : extensive tutorial with several examples. Exists also as an online resource or free ebook (in epub and pdf)
- [For the Love of Go](https://bitfieldconsulting.com/books/love) : A book by John Arundel, which covers the fundamentals of Go and has an online video course
- [Go Style Guide](https://google.github.io/styleguide/go/index) : Article about best practices when writing in Go

### Videos

- [Go Path](https://app.pluralsight.com/paths/skill/go-core-language) : Series of extensive courses in Go in Pluralsight
- [Golang Dojo](https://www.youtube.com/@GolangDojo) : YouTube channel that offers many Go specific tutorials
- [NerdCademy](https://www.youtube.com/@NerdCademyDev) : Another YouTube channel dedicated to Go
- [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs) : Google I/O 2012 lecture on Go Concurrency Patterns
- [Learn Go in one video](https://www.youtube.com/watch?v=YzLrWHZa-Kc) : A course in Go