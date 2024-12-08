# MispASM

- I made an [example folder](https://github.com/antosmichael07/MispASM/tree/main/examples) with examples to help you

## Getting started

- You need the [Misp virtual machine](https://github.com/antosmichael07/MispASM/tree/main/mispasm-vm) `mispvm` and [Misp assembler](https://github.com/antosmichael07/MispASM/tree/main/mispassembler) `mispc`
- Optionally you can have the [Misp library builder](https://github.com/antosmichael07/MispASM/tree/main/mispasm-lib-builder) `misplib` for making libraries

### Your first program

- Start by making a main [function](#functions) like this :

```
global _start

_start
```

- The `global _start` says that the `_start` [function](#functions) is gonna be the main [function](#functions)
- Then store `"Hello World!"` in a [register](#registers)

```
global _start

_start
    mov s0 "Hello World!"
```

- Now print `"Hello World!"`

```
global _start

_start
    mov s0 "Hello World!"
    
    push s0
    call "std.print"
    pop s0
```

- Great you wrote your first program, if it didn't work, you might not have the std [library](#writing-a-library), you can get the code from the [lib folder](https://github.com/antosmichael07/MispASM/tree/main/lib) here on github and compile it manually

## What MispASM offers

### Types

- int, float, string, register, variable, constant
- To differentiate between numbers you can type :
```
i8-123
i16-123
i32-123
i64-123
u8-123
u16-123
u32-123
u64-123
f32-123.45
f64-123.45
```
- If you just type a number it is automatically `i32` and if you type a `.` in the number it is automatically `f32`
- Strings are typed : `"Hello World!"`

### Functions

- Functions start with `_`
- Declare main function at the start of the program : `global _start`
- Call a function : `fcal <function>`
- Functions are typed as strings in [arguments](#types)
- Call a function from a [library](#writing-a-library) : `call "std.print"`
- When calling a [library](#writing-a-library) from a program, the `.misplib` file has to be located in `../lib/<name>.misplib`
- You can pass [registers](#registers) to the [library](#writing-a-library) with [stack](#stack)

### Variables and constants

- Variables have to be set at the start of the program and can be changed with : `set <variable> <arg2>`
- Constants cannot be changed
- You need to declare them before global [function](#functions)

```
const:
    hello "Hello World!"

var:
    bye "Goodbye!"

global _start
```

- They have to be declared in this order, firstly constants and then variables, but you don't need to use both

### Instructions

- They have a maximum of 2 [arguments](#types)
- Syntax : `<instruction> <arg1> <arg2>`

### Registers

- Registers are typed : `li0`, the zero indicates that it is the first long integer
- There is 256 of every register
- Are used as [variables](#variables-and-constants), but they can perform operations, [variables](#variables-and-constants) can only be set
- Can be used as [arguments](#types) to [instructions](#instructions)
- Store data with : `mov <register> <arg2>`
- Return registers are used for storing data after an [instruction](#instructions) that performs a calulation, for example `add` sets the register : `rli` return long integer or other if you use different [type](#types) than `i32`

### Stack

- Stores data and passes them to [libraries](#writing-a-library)
- Only [registers](#registers) can be stored
- Store data with : `push <register>`
- Delete data with : `pop <register>`

## Writing a library

### What you need

- [Misp virtual machine](https://github.com/antosmichael07/MispASM/tree/main/mispasm-vm) `mispvm`
- [Misp assembler](https://github.com/antosmichael07/MispASM/tree/main/mispassembler) `mispc`
- [Misp library builder](https://github.com/antosmichael07/MispASM/tree/main/mispasm-lib-builder) `misplib`
- [Go programming language](https://go.dev/) `go`

### How to

- When [calling](#functions) a library from a program, the `.misplib` file has to be located in `../lib/<name>.misplib`
- Use the [Misp Library builder](https://github.com/antosmichael07/MispASM/tree/main/mispasm-lib-builder), it will generate a structure of files
    - `./build.bat` for building your library
    - `./lib/lib.go` is for decoding and encoding data between a program and your library, you don't need to worry about it
    - `./lib/<name>.go` is for your functions in the library, there is a map of functions where you write your functions
    - `./test/test.mispasm` the `.mispasm` program that is used to test your library
- The command [misplib](https://github.com/antosmichael07/MispASM/tree/main/mispasm-lib-builder) needs these arguments : `misplib <destination> <name>`

## Cheatsheet

|[Instruction](#instructions)|Description             |
|----------------------------|------------------------|
|`add`                       |Addition                |
|`sub`                       |Subtraction             |
|`fcal`                      |Function call           |
|`mul`                       |Mulitplication          |
|`div`                       |Division                |
|`call`                      |Call from library       |
|`push`                      |Push to the stack       |
|`pop`                       |Pop from the stack      |
|`mov`                       |Move to register        |
|`label`                     |Label                   |
|`jmp`                       |Jump to label           |
|`mod`                       |Modulo                  |
|`cmp`                       |Compare                 |
|`je`                        |Jump equal              |
|`jne`                       |Jump not equal          |
|`jg`                        |Jump greater            |
|`jge`                       |Jump greater or equal   |
|`jl`                        |Jump less               |
|`jle`                       |Jump less or equal      |
|`inc`                       |Increase                |
|`dec`                       |Decrease                |
|`ret`                       |Return                  |
|`def`                       |Defenestrate the program|
|`set`                       |Set a variable          |
|`and`                       |Logical AND             |
|`or`                        |Logical OR              |
|`xor`                       |Logical XOR             |
|`shr`                       |Shift bits right        |
|`shl`                       |Shift bits left         |
|`not`                       |Logical NOT             |
|`hlt`                       |Halt                    |

|[Register](#registers)|Description                            |
|----------------------|---------------------------------------|
|`bi`                  |Byte integer `i8`                      |
|`si`                  |Short integer `i16`                    |
|`li`                  |Long integer `i32`                     |
|`lli`                 |Long long integer `i64`                |
|`bui`                 |Byte unsigned integer `u8`             |
|`sui`                 |Short unsigned integer `u16`           |
|`lui`                 |Long unsigned integer `u32`            |
|`llui`                |Long long unsigned integer `u64`       |
|`lf`                  |Long float `f32`                       |
|`llf`                 |Long long float `f64`                  |
|`s`                   |String                                 |
|`rbi`                 |Return byte integer `i8`               |
|`rsi`                 |Return short integer `i16`             |
|`rli`                 |Return long integer `i32`              |
|`rlli`                |Return long long integer `i64`         |
|`rbui`                |Return byte unsigned integer `u8`      |
|`rsui`                |Return short unsigned integer `u16`    |
|`rlui`                |Return long unsigned integer `u32`     |
|`rllui`               |Return long long unsigned integer `u64`|
|`rlf`                 |Return long float `f32`                |
|`rllf`                |Return long long float `f64`           |
|`rs`                  |Return string                          |

|[Type](#types)|Description            |
|--------------|-----------------------|
|`i8 `         |8-bit integer          |
|`i16`         |16-bit integer         |
|`i32`         |32-bit integer         |
|`i64`         |64-bit integer         |
|`u8 `         |8-bit unsigned integer |
|`u16`         |16-bit unsigned integer|
|`u32`         |32-bit unsigned integer|
|`u64`         |64-bit unsigned integer|
|`f32`         |32-bit float           |
|`f64`         |64-bit float           |