# Data Structures and Algorithms in Go

i. [Go Basics](#go-basics)

ii. [Go Tips and Tricks](#go-tips-and-tricks)

iii. [Misc Notes](#misc)

#

## Go Basics

#### Loops

`for` is the only loop available
```go
  // standard for loop
  for i := 1; i < 10; i++ {}

  // while loop
  for ; i < 10; {}

  // can omit semicolons if there's only a condition
  for i < 10 {}

  // while (true) ~ can omit the condition
  for {}
```

#### Pointers

a `variable` represents the memory address of some value

a `pointer` is a value that represents the address of another variable

```go
  // `a` stores the value `13`
  // `a` is a memory address, and `13` is the value held at that address
  // `a` represents the value `13`  ~  `&a` represents the address of `13`
  a := 13

  // `b` stores the reference address of `a` ~ some wild hex looking number
  // if `a` represents a value, `&a` represents the address of that value
  b := &a

  // manipulate the value stored at `a` by dereferencing `b`
  // `a` is now equal to 14
  *b++
  // however, manipulating `c` will not manipulate the value stored at `a`
  // `a` remains equal to 13  ~  `c` is now equal to 14
  c := *b
  c++
```

## Go Tips and Tricks


## Misc

GOPATH via Powershell

    cd $Env:gopath

`Polymorphism` - using a single symbol for multiple different types

`&` returns a memory address
`*` returns the value at the address ?