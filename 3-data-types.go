package main

// Go has quite a large number of primitive data type which seem to be an
// aspect it inherited from C. This may be helpful to program memory efficient
// data structures. You have boolean, string, byte, various numeric, and derived
// data types. The numeric type fall into integer, float, and complex data types.
// You can declare integers with different word lengths, i.e. 16, 32, and 64 bit.
// For floats you can choose between 32 and 64 bit and for complex numbers between
// 64 and 128 bits. The latter are represented internally by a float32 or float64
// number and and imaginary number parts using the remainder.

import (
    "fmt"
)

func main() {
    var x float64 = 20.0
    // simplified variable declaration with type inference.
    // Let's see what type it is later
    y := 42

    fmt.Println(x)
    fmt.Println(y)
    fmt.Printf("y is of type %T\n", y)

    // here is another variant of variable declaration with type inference.
    // Note, multiple variables of different types can be declared in one
    // go using type inference, i.e., by providing their initial values.
    var a, b, c = 3, 4, "foo"
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Printf("a is of type %T\n", a)
    fmt.Printf("b is of type %T\n", b)
    fmt.Printf("c is of type %T\n", c)

    // The compiler does strict type checking. So, we need to use cast operators
    // to cast between different types. Surprisingly, if there is an overflow
    // for a number type coversion, no error will be thrown. Seems like some
    // of the bad things have also been derived from ancient C. See also
    // https://github.com/golang/go/issues/19624
    var n1 float64 = 748375347509347587348957348920.0
    var n2 float32 = float32(n1)
    var n3 uint16 = uint16(n1)
    fmt.Println(n1)
    fmt.Println(n2)
    fmt.Println(n3)
}