# CH3: Monolithic style

As I have seen, this style is fairly simple and doesn't require a lot of thinking. (Actually, it doesn't require thinking at all, apart from thinking about the logic itself, of course)

The constraints of this style are:

1. No named abstractions.
2. No, or little use of libraries.

But, what does each one mean?


### 1. "No named abstractions"

"Named abstractions" are functions, classes, or any modular chunck having a name. The monolithic style prevents you from using any of these and you just put all your logic implementation in one single (and probably giant) block. You don't divide the big task into smaller tasks, like what we do in functional programming where we divide the logic into separated functions. You can think of it as the complete opposite of "separation of concerns"; you don't separate anything, all your 'concerns' are in one block.

As shown in the python code snippet provided in the chapter, the author wrote every step to implement the task in a single block: retrieving stop words, reading the file line by line, finding the beginning and the end of each word, calculating its frequency, reordering, and printing the top 25 most frequent word; all in one block.


### 2. "No, or little use of libraries"

It means that you don't use any libraries or dependecies to implement your task, unless it's very necessary. You write everything by yourself. However, it's fine if you use code that is within the standard library of the programming language you use. You basically use if statements, loops, and few external functions.

In the code provided, the author didn't use a split function to split the line read. Instead, she manually find the start and end of each word and process them. She also didn't use a sort function, she wrote a code to sort them based on frequency of each word.


## How I implemented this style in Go?

#### Short answer:

I just put everything in the main function :)


#### Long answer:

First I defined two global variables: a map to store the stop words and a slice of pairs which I defined before; each pair has word and frequency.

I used a map instead of a slice to store stop words because I want to quickly check if a word is a stop word. I defined the value of the map to be an empty struct becuase I only care about the keys stored in the map, which are the stop words, and empty structs take zero memory! and this is a very neat feature in Go :)

Then, I implemented each step in the task in the main function; no other functions are defined.

I used functions like `ReadFile` and `Open` from the `os` package becuase I couldn't interact with files without them. I had to also define a scanner from `bufio` package so I can read the file line by line, otherwise I would use `ReadFile` function and load the whole file into memory which isn't something very efficient. And, of course, I had to use functions from the `unicode` package to check characters so I can decide the start and end of words.

The implementation is basically the same as the python code defined in the chapter, but I did a bit of cleaning to git rid of those evil nested if-statements.


## So, when should you use this style?

If you want to get kicked out of your job or, if you are an intern, want to get scolded in the project discussion, then gladly use it ;)

What I mean is, **DO NOT EVER THINK ABOUT APPLYING THIS STYLE IN YOUR PRODUCTION CODE**, never.

It has many issues as you can see. It produces code that is hard to read, debug, fix, and maintain; it's a spaghetti code, and you know how bad a spaghetti code is.


## Then, if this style is that bad, why did the author talk about it?

For educational purposes. She wants us to see how code in the "good old times" used to be; when people were using languages like Fortran and COBOL to write their code; before the appearence of functional and OO programming. 

However, you may use it for very small scripts or you just want to implement a quick prototype. But other than that, don't.