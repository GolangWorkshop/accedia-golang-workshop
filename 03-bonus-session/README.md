# Go Workshop

# Table of contents
- [**Context**](#context)
- [**CGO**](#cgo)

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
