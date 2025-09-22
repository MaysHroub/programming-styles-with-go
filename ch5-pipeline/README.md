# CH5: Pipeline style

## Constraints of this style:

1. Larger problem is decomposed using functionl abstraction. Functions take input, and produce output.
2. No shared state between functions.
3. The larger problem is solved by composing functions one after the other, in pipeline.


## Further explanation

The pipeline style comes from the theory of 'mathematical functions' and hence restricts functions to behave like 'math functions' that take one input and produce one output, with no side effects. So, in the 'pure' form of this style, the world outside the boxed functions doesn't exist, except in two cases: as a source of input and as an output receiver.

A function is said to have side effects if it does things other that just compute and return; it modifies a state or has an interaction with the outside world, such as modifying the value of a non-local variable or an argument, reading/writing data, etc.

There's no shared state between functions and they must be idempotent, which means calling a function multiple times has the same effect as calling it once, given the same input of course.

A multi-argument function is transformed to a sequence of higer-order functions with one argument. This process is called currying.

A higer-order function is simply a function that takes function as input and/or give function as an output.


## How I implemented the term frequence task with this style in Go?

The task is divided into sub-tasks of which each function will handle. The sub-tasks are:

1. Read the file
2. Normalize the file's content
3. Convert it to list of words
4. Remove stop words
5. Count freqeuncies
6. Sort
7. Print the result


**Note:**

1 -> input: filepath, output: file content as string

2 -> input: file content, output: normalized string

3 -> input: normalized file content, output: list of words

4 -> input: list of words, output: list of words with no stop words

5 -> input: list of words with no stop words, output: list of pairs each with word and frequence

6 -> input: list of pairs, output: sorted list of pairs

7 -> input: sorted list of pairs, output: nothing (it just prints the result on the console)

Like an actual pipeline!

**Anyway**

I created a function for each task; each function takes exactly one input and gives one output. I also had to define a pair struct so I can store them as slice of pairs to be able to sort them.

The cool thing here is that you can use standard/external libraries, unlike the monolithic style.

Exercise 5.2 of this chapter requires you to pass the file name, of which removeStopWords function will read stop words from, along with the words list. Since multi-argument function is not allowed in this style, you have to use currying. So, I created a function that takes the file name and returns a **function** that processes these stop words.
