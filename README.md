# Programming Styles With Go

[*Exercises in Programming Style*](https://www.goodreads.com/book/show/23012704-exercises-in-programming-style) is a book that illustrates different styles in programming. The author implements the same task across all chapters, each in a different style: **The Term Frequency** task, in which it calculates the frequencies of words in a given file and displays the top 25 most frequent words after removing the stop words. The author writes the task's code in python, but I have implemented it, and the other exercises, in Go.

The last exercise of each chapter requires implementing one of the tasks presented in the *Prologue* in the corresponding style. The task I chose is **Word-Index**, in which you record the pages the word occured in and displays the top 25 word after sorting them in an alphabetical order. You specify the pages and how many lines each one should have. Also, if the word occurs more than 100 time, it's removed.

Oh, and I also added some unit tests in **ch5-pipeline** package.


## Styles/Chapters Implemented:

- CH3 - Monolithic
- CH5 - Pipeline
- CH10 - Things
- CH24 - Quarantine
- CH25 - Persistent Tables
- CH28 - Actors
- CH30 - Map-Reduce

Each style, along with its exercises, is implemented in a separate package. 

A `README.md` file is added in each package and explains some information about the style, its constraints, and how I implemented the term frequency task in Go.


## Prerequisites
- Go 1.19+ 
- C compiler (gcc/clang) for SQLite3 driver compilation


## Installation

### 1. If you have Go in your machine
```sh
go install github.com/MaysHroub/programming-styles-with-go@latest
```

### 2. If you have Git installed in your machine, clone the repo
```sh
git clone https://github.com/MaysHroub/programming-styles-with-go.git
```


## How to Run

Honestly, I don't know why you would this. It will give you the same output regardless of the style you choose :)

*Anywaaaaay,* if you still insist..

To run a specific exercise, include the package name, the exercise directory, and `main.go`. For example:

```sh
go run ch3-monolithic/ex3.1/main.go
```

Run this from the root directory of the project, otherwise it won't work.


## Developer's Note

This repository is for educational purposes - it's designed to be studied and referenced, not executed as a runnable application. (Well, it's not that anyone would study it)

And.. NO, I won't add any documentation.




