# Introduction to Go

> To run a Go program without compiling it, use the following command:

```bash
go run <filename>.go
```

> To compile a Go program, use the following command:

```bash
go build <filename>.go
```

## Basic Syntax

> Some of the basic keywords in Go are:

* ```:=``` is for declaration and assignment
* ```=``` is for assignment only
* ```var``` is for declaration only
* ```const``` is for constant declaration
* ```_``` is for blank identifier
* ```&``` is for address of operator
* ```*``` is for pointer dereference operator
* ```new``` is for pointer allocation
* ```make``` is for slice, map and channel allocation
* ```cap``` is for capacity of slice, map and channel
* `iota` is for enumeration (Incrementing integer constant)

## Data Types

* ```bool``` is for boolean
* ```string``` is for string
* ```int``` is for integer
* ```float``` is for floating point
* ```complex``` is for complex number
* ```byte``` is for alias of uint8
* ```rune``` is for alias of int32
* ```uint``` is for unsigned integer
* ```int8``` is for 8-bit signed integer
* ```int16``` is for 16-bit signed integer
* ```int32``` is for 32-bit signed integer
* ```int64``` is for 64-bit signed integer
* ```uint8``` is for 8-bit unsigned integer
* ```uint16``` is for 16-bit unsigned integer
* ```uint32``` is for 32-bit unsigned integer
* ```uint64``` is for 64-bit unsigned integer
* ```uintptr``` is for unsigned integer large enough to store the uninterpreted bits of a pointer value
* ```float32``` is for 32-bit floating point
* ```float64``` is for 64-bit floating point
* ```complex64``` is for 64-bit complex number
* ```complex128``` is for 128-bit complex number

### Int type
```json
uint8       unsigned  8-bit integers (0 to 255)
uint16      unsigned 16-bit integers (0 to 65535)
uint32      unsigned 32-bit integers (0 to 4294967295)
uint64      unsigned 64-bit integers (0 to 18446744073709551615)
int8        signed  8-bit integers (-128 to 127)
int16       signed 16-bit integers (-32768 to 32767)
int32       signed 32-bit integers (-2147483648 to 2147483647)
int64       signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

```

### Float type
```json
float32     IEEE-754 32-bit floating-point numbers
float64     IEEE-754 64-bit floating-point numbers
complex64   complex numbers with float32 real and imaginary parts
complex128  complex numbers with float64 real and imaginary parts
```

### Alias number types
```json
byte        alias for uint8
rune        alias for int32
```


> The preferred convention in Go is to use `MixedCaps` or `mixedCaps` rather than _underscores_ to write multi-word names