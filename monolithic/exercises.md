# CH3 Monolithic - exercises' solutions (some of them)

## Ex3.1
**"Another language. Implement the example program in another language, but preserve the style."**

Already solved.


## Ex3.2
**"Readlines. The example program reads one line at a time from the file. Modify it so that it first loads the entire file into memory (with readlines()), and then iterates over the lines in memory. Is this better or worse practice? Why?"**

I replaced `os.Open` with `os.ReadFile` for `input.txt` file; this reads the whole file and load it in the variable `inputFileContent` as slice of bytes. I converted the slice to a string and iterated through each character in that string.

#### Is this better or worse practice, any why?

It depends. 

It's generally fine to read the whole file and load it into memory for small to medium-sized files. This approach reduces I/O calls and can make processing faster, and preferrable if you want to process the file multiple times. However, it's not recommended for large files as it can lead to high memory consumption and this results an out-of-memory error. Reading the file line by line is generally safer and keeps memory usage low regardless of the file size. That said, choosing either approach depends on the file size and the use case; it's ok to load the whole file into memory small to medium-sized ones, and read line by line for larger one.


## Ex3.3

**"Two loops. In lines #36–42 the example program potentially reorders, at every detected word, the list of word frequencies, so that it is always ordered by decreasing value of frequency. Within the monolithic style, modify the program so that the reordering is done in a separate loop at the end, before the word frequencies are printed out on the screen. What are the pros and cons of doing that?"**

I did separate the sorting logic and put it outside the loop and it did improve the performance, like dramatically.

I honestly don't see any cons with this method. Sorting the words separately is better than sorting them while processing the words. And it has better time complexity.

Assume there are N words in total and M unique words. In most cases N >> M. Sorting the words within the main loop results with O(N.M) time complexity, because in **the worst case**, with each iteration (aka when processing each word) you need to shift the word M shifts after updating its frequency, resulting in O(N.M). Sorting outside the loop, in the other hand, results in O(M^2) time complexity in the worst case. This also allows you to use other sorting algorithms, like quick sort which gives you a better time complexity of O(M log M).
