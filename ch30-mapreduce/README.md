# CH30: Map-Reduce

## Constraints of this style:

1. Input data is divided into **blocks**.
2. A **map** function applies a given worker function to each block of data, potentially in parallel.
3. A **reduce** function takes the results of the many worker functions and recombines them into a coherent output.


## Further explanation

```
>> Problems must be parallelizable
>> Results must be combinable
```

The **MapReduce** is a model to process large datasets by dividing the work/data into smaller chunks/blocks that can be processed independently and in parallel, then combining the results. It's similar to the `fan-out/fan-in` pattern where in the `fan-out` stage you divide the data into blocks and assign each to worker function, and then you extract the information you need from the resulted outputs in the `fan-in` stage.

The three core components of this style are:
1. Partition - divide the work; following 'divide and conquer' approach.
2. Map - process each chunk independently, and possibly in parallel.
3. Reduce - combine all results and extracts a meaningful information.

As illustrated in the book, the `map` is a function that takes two arguments: a worker function and the data chunks to be processed. It applies the worker funtion to each chunk, either in a regular loop with sequential execution or in parallel by calling the function in a separate thread, and this is better.

`reduce` is a function that also takes two arguments: a worker function and the results set. The worker function has two parameters in the same type as the result. It combine them and returns the combined result. The `reduce` phase can be done either in sequential or in parallel.


## How I implemented the term frequence task with this style in Go?

#### The main flow:

- The data is divided into chunks; each chunk is a set of particular number of lines -> `partition`
- Each chunk is divided into a list of pairs with the word and a frequency of 1
- The pairs are filtered and stop words are removed
- The pairs lists are grouped together and the frequencies of a word are combined

I implemented two functions: `map` and `reduce`. As mentioned above, each take two arguemnts with a worker function being the first argument and data blocks as the second argument. `map` applies the worker function `split` to each chunk in a sequential manner; the parallel version is implemented in `ex30.3`. `split` converts each word into a `pair` which has the word itself with a frequency of 1. `reduce` applies the worker function `countWords` to the lists of pairs that resulted from `map` function. The `countWords` function takes two pairs lists and combines them into a single list.

The data is normalized before and filtered (removing stop words) while its being processed.
