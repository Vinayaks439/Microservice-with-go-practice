# Data Types

Data types define possible values that variable can have and operations that can be performed. Go supports basic data types (like: int, string, bool etc) and complex data types (like: array, maps , pointers , structures etc)

Basic Data types are self explanatory hence lets dive into complex data types.

# Pointers

Pointer is a data type that holds the memory address of a value. The ***default value*** for pointer is ***nil***, decalration is as follows:

`var pi *int`

This is an integer pointer. We can create pointer to any type of value.

## We have two pointer operations.

### The 1st operator ***'&'*** is used to find the address of variable. For example:

`var pi = &i`

Now, pointer pi will have value of address of variable i.

### The 2nd operator **'*'** can give us access to value pointer points to. We can use this operator to change the pointed value. For example:

`*pi = 5`

Integer Pointer will now point to value 5.

## Code example.

```go
package main
import "fmt"

func main() {
    var i int = 8
    var pi *int
    pi = &i
    *pi = 5
    fmt.Println(i)
}
```

# Struct

Struct is a complex data type and represents the collection of fields. We use ***type*** keyword to declare ***struct***.

## Code example:

```go
package main
import "fmt"

type Circle struct {
    X, Y, R int
}

func main() {
    c1 := Circle{2,3,5}
    c2 := Circle{Y: 1, R :5}
    c3 := Circle{}

    fmt.Println(c1)
    fmt.Println(c2)
    fmt.Println(c3)
}
```

When we create new struct we don't need to specify values for all fields. For all fields without value, default value will be assigned. In the example above, the 1st circle (c1) is created with all values specified. The 2nd circle (c2) is created with values for Y and R, 0 is assigned as default for X. We use field name : (colon) field value notation to assign value to specific field. The 3rd circle (c3) is created with no value assigned to any field. All fields will have default values for their respective types.

# Arrays

Array is a complex data type and represents collection of same type of elements. Each element is identified by array index. The 1st element of array will have index 0.

### Declaring an Array

`var arr [5]int`

[5] here denotes the length of the array.

`arr := [5]int{4, 7, 2, 9, 1}`

In Go, array length is part of array type. So when size is once defined, cannot be changed. This may sound like a problem, but Go has an elegant solution for working with arrays.

# Slices

In order to fix array length problem, slices are introduced in Go. Slices do not store any data, we can describe them as pointers to *"underlying"* arrays. 

We can create slices with built-in function **make()**. This function receives three arguments: **array type, length and capacity.**

## Code example: 

```go
package main
import "fmt"

func main(){
    s1 := make([]int, 5)
    s2 := make([]int , 0 , 5)

    s1 = append(s1, 1)
    s2 = append(s2, 2)

    fmt.Println(len(s1), cap(s1))
    fmt.Println(len(s2), cap(s2))
}
```

Here, we have two slices created with **make()** function s1 and s2. The 1st slice (s1) is created without capacity argument, so by default its capacity will be the same as length. This call of make function will create array of five 0 values (default value for int)

The second slice (s2) will be created with length 0 and capacity 5. In this case, make will allocate space for 5 int elements but no real values will be stored.

When we call **append()** function for each of these slices, results will be different. The first slice (s1) will append value 1, but slice is not big enough so additional space will be allocated. Usually, allocated space will be doubled.

Now our slice s1 will contain 6 elements (5 zeros and 1) and have capacity of 10 (double the original). First print function will display values 6 and 10 for length and capacity of slice (s1)

Second slice (s2), originally, does not contain any elements. After call of **append()** function, s2 will contains one element and the slice is big enough so there is no need for additional space allocation.

The second print function will display values 1 and 5 for length and capacity of slice s2.

If we know exact max capacity for our slice, we can create slice without capacity argument in **make()** function and reference elements of slice by indexes. We should be very careful with **append()** function to avoid any unnecessary allocation of memory. If possible, we must try to avoid any additional allocations, because they can affect performance of our programs.

# Maps

Maps are complex data types that store key, value pairs. Using a key we can find the value stored in that map. Each key can appear only once in a map. The default value for map is **nil**. Keys can't be added in nil map.

`m := make(map[int]string)`

## Code example:

```go
package main
import "fmt"

func main(){
    m := make(map[int]string)

    m[1] = "Monday"
    m[3] = "Wednesday"

    value, ok := m[2]

    fmt.Println(value, ok)
}
```