# Go Lang

### Why Go?

Go is modern, concurrent, statically typed, compiled language. It is designed to be efficient, fast and easy to work with. It is also open source and backed by Google.

### Use case

- Creating robust, scalable, concurrent backend services
- Writing CLI tools and APIs

### Important features

- Routine and Channels
- Garbage collection
- Not a strict object oriented language but has structs and have composition instead of inheritance
- No exceptions or generics instead has error handling and interfaces respectively
- Can use pointers, but no pointer arithmetic

**To run a Go program without compiling, use the following command**:

``` shell
go run <filename>.go
```

**To compile a Go program, use the following command**:

```shell
go build <filename>.go
```

### Basic Syntax

- `:=` for declaration and assignment while `=` is for assignment only.
- `&` and `*` are for address of operator and pointer dereference operator respectively.
- `new` is for pointer allocation while `make` is for slice, map and channel allocation.
- `cap` is for capacity of slice, map and channel.

### Data Types

- `bool` for boolean
- `string` for string
- `int` for integer
- `float` for floating point
- `complex` for complex number
- `byte` for alias of uint8 { 0 to 255 }
- `rune` for alias of int32 { -2147483648 to 2147483647 }
- `uint` for unsigned integer
- `uintptr` is for unsigned integer large enough to store the uninterpreted bits of a pointer value

Along with these there are other types like `int8`, `int16`, `int32`, `int64`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64`, `complex64`, `complex128` etc.

### Implicit and explicit type inference

```go
var name dtype = value
name := value
```

**Functions**

```go
func func_name(param1 type1, param2 type2) (return_type1, return_type2) {
    // function body
}
```

```go
func CalculateArea(width, height float64) float64 {
    return width * height
}
```

**Methods**

```go

func (receiver_name receiver_type) func_name(param1 type1, param2 type2) (return_type1, return_type2) {
    // function body
}
```

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method associated with Rectangle type
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

**Arrays**
    
```go
var myArray [size]data_type            //Declaration

myArray := [size]data_type{value1, value2, ...}    //Declaration and Initialization

myArray := [...]data_type{value1, value2, ...}    //Declaration and Initialization with a literal
```
In go each array with a specific size is considered as a different type. It is designed like this to avoid the overhead of checking the size of the array at runtime. Arrays are value types { _when passed to a function, the copy of the array is passed not the pointer to the array_ }.

**Slices**
    
```go
var mySlice []data_type            //Declaration

mySlice := []data_type{value1, value2, ...}    //Declaration and Initialization

mySlice := make([]data_type, size)    //Declaration and Initialization with a literal
```

Slices are reference types { _when passed to a function, the pointer to the slice is passed not the copy of the slice_ }. Slices are not thread safe. Slices are ordered collections. 

**Maps**

```go
var myMap map[key_data_type] value_data_type            // Declaration
myMap = make(map[key_data_type]value_data_type)        // Initialization
myMap2 := map[key_data_type]value_data_type{"a": 1}    // Declaration and Initialization with a literal
```

In go maps are reference types { when passed to a function, the pointer to the map is passed not the copy of the map }. Maps are not thread safe. Maps are unordered collections. Slices, maps, and other maps cannot be used as map keys. 

**Structs**

```go
type struct_name struct {

    field1 data_type
    field2 data_type
    ...
}
```

Structs are a way to bundle attributes and methods together. Structs can be passed to functions and can be embedded in other structs. They can implement interfaces also. Structs sort of servers the purpose of classes in other languages.

**Interface**

```go

type interface_name interface {
    method_name1() return_type1
    method_name2() return_type2
    ...
}
```

Interfaces are a way to define behavior. They are not concrete types. They are implemented implicitly. A type implements an interface by implementing its methods. A type can implement multiple interfaces. Interfaces can be used as fields in structs. Interfaces can be used as function parameters, function return types, type of a variable, type of a map key, type of a slice.

**Error Handling**

```go
func func_name(param1 type1, param2 type2) (return_type1, return_type2, error) {
    // function body
}
```

```go
func main() {
    result, err := func_name(param1, param2)
    if err != nil {
        // handle error
    }
    // use result
}
```

Go handles errors differently than other languages. It uses the `error` interface to represent errors. In go we have both run time and compile time errors. We call the run time error as panic. We can use the `panic` function to throw a panic. We can use the `recover` function to recover from a panic. We can use the `defer` keyword to defer the execution of a function until the surrounding function returns. `panic` is similar to exceptions in other programming languages. When a panic occurs, it stops the regular flow of the program and begins unwinding the stack, which means it starts to exit functions and defer statements (if any) in reverse order until the program terminates. Panicking should be reserved for truly exceptional situations where we don't know what to do with the current exception, and generally in all other cases we use error handling.

```go
 // panic
func main() {
    fmt.Println("Starting the program")
    panic("A severe error occurred: stopping the program!")
    fmt.Println("Ending the program") // won't be executed
}
```

`defer` is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. It is sort of like a finally block in other languages. We usually place defer statements at the beginning of a function. When the surrounding function returns, the deferred function calls are executed even in the case of a run time panic. It is typically used for cleanup or resource release. { We don't want any resources to be leaked, even after the program panics. }

```go
// defer
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
```

```go
// recover
func main() {
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    panic("PANIC")
}
```

**Pointers and references**

```go
var myInt int = 42
var myIntPointer *int = &myInt
fmt.Println(*myIntPointer) // 42
```

Pointers and dereferencing in go is similar to other languages, but go does not support pointer arithmetic.

**Go routines and Concurrency in Go**

Go is famous for its superior concurrency support. Go routines are lightweight threads. They are not OS threads. They are managed by the go runtime. They are multiplexed onto OS threads. They are cheap to create and destroy and are not preemptive. They are cooperative. Unlike python we not limited by the GIL. We can run as many go routines as we want. Although the routines may not always run in parallel, but they could. We can use the `go` keyword to create a go routine. We can use the `sync` package to synchronize go routines.

**Channels**

Go routines becomes truly powerful when we use them with channels. Channels are a way to communicate between go routines. They are used to synchronize go routines.


## Go Project Structure

```shell
my-golang-project/
│
├── cmd/
│   └── myapp/
│       └── main.go
│
├── pkg/
│   ├── utils/
│   │   ├── utility1.go
│   │   └── utility2.go
│   │
│   └── ...
│
├── internal/
│   ├── app/
│   │   ├── myapp.go
│   │   └── ...
│   │
│   └── ...
│
├── api/
│   ├── http/
│   │   ├── server.go
│   │   └── ...
│   │
│   └── ...
│
├── web/
│   ├── static/
│   │   ├── css/
│   │   ├── js/
│   │   └── ...
│   │
│   ├── templates/
│   │   ├── index.html
│   │   └── ...
│   │
│   └── ...
│
├── configs/
│   ├── config.yaml
│   └── ...
│
├── scripts/
│   ├── deployment.sh
│   └── ...
│
├── tests/
│   ├── unit/
│   ├── integration/
│   └── ...
│
├── .gitignore
├── go.mod
├── go.sum
├── README.md
```
