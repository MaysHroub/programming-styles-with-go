# CH24: Quarantine style

## Constraints of this style:

1. Core program functions have no side effects of any kind, including IO.
2. All IO actions must be contained in computation sequences that are clearly separated from the pure functions.
3. All sequences that have IO must be called from the main program.


## Further explanation

The main constraint in this style is that **core functions cannot do IO**. Any IO must be quarantined in high-order functions and executed only in `main`.

The reason we separate IO from core functions is to keep them 'pure'. IO makes functions 'impure' as it may make the function not produce the same output for given input if called multiple times. Impure functions are harder to test, reason about, and predict. So, we need to keep IO isolated and explicited.

**So, how it works?**

- First-order functions must be **pure** (always produce the same result for the same input).
- IO functions (like reading files) are wrapped in higher-order functions. Doing so makes first-order function be 'pure' since any call to any of them will give the same output (their inner function).
- These IO functions are stored in a chain and executed **later** in main.

You also need to implement a quarantine class the does lazy evaluation of the chain of functions; it first stores them, and then call them in main.


## How I implemented the term frequence task with this style in Go?

It's similar to the implementation of the pipeline style, but this time it uses high-order functions to quarantine IO and implements a quarantine struct that stores functions through its `bind` method and execute them in main via its `execute` method.

It was a bit tricky to implement the quarantine struct and the functions since Go is a strictly typed language and types are defined in compile time. So, I had to use `any` as the type of parameters and return values for the functions so that I can use them with the `quarantine` struct.


## Note

The code of this style takes the input file path as an argument when running the file, so pass `./files/input.txt` when running the file from root directory.