# CH28: Actors

## Constraints of this style:

1. The larger problem is decomposed into things that make sense for the problem domain.
2. Each thing has a queue meant for other things to place messages in it.
3. Each thing is a capsule of data that exposes only its ability to receive messages via the queue.
4. Each thing has its own thread of execution independent of the others.


## Further explanation

The actors style solves your problem by creating **things** called *actors* that communicate with each other through messages. Each actor is designed to perform a specific task and has a message queue, like a mailbox, that other actors can drop messages into this queue. Messages are lined in the queue to be processed by the actor. Other actors don't directly access each other, hence, each actor is encapsulated; the only way to communicate is by sending messages. Each actor has its own thread and thus multiple actors can work together simultaneously **(concurrently)**; they don't wait for each other to finish. Of course, if there are no messages in the queue, the actor is blocked and doesn't do anything until a new message arrives.

This style is more *complicated* and involves other concepts like **supervisors** which manages particular actors. 

This style is well-suited for several real-world applications, including distributed systems, real-time applications, and multiplayer games.

So, it is perfect for systems where you need independent, concurrent processing with reliable message delivery between components.

`Akka` is one of the modern applications of this style.


## How I implemented the term frequence task with this style in Go?

The cool thing about Go is that **concurrency** is one the features that makes Go outsands among other languages; it is optimized and robust.

The program consists of four actors: `DataStorageManager`, `StopWordManager`, `WordFrequencyManager`, and a `Controller`. So four goroutines (or threads) along with the main thread.

The main flow of the program is like this:

```
- Main thread -> sends an 'init' message to DataStorageManager and StopWordManager.

- Upon receiving the 'init' message:
  - DataStorageManager -> reads the file given its name in the message.
  - StopWordManager -> loads the stop words list.

- Main thread -> sends 'run' message to the Controller to start executing.

- Upon receiving the 'run' message:
  - Controller -> sends 'process-words' message to DataStorageManager to start the process.

- Upon receiving the 'process-words' message:
  - DataStorageManager -> iterates through each word and sends it as a message with 'filter' to StopWordManager to check if the word is a stop word or not.

- Upon receiving the 'filter' message:
  - StopWordManager -> checks if the word received is a stop word or not, if it is, it sends it to WordFrequencyManager to store it and its count.

- When all words are counted, WordFrequencyManager generates the top 25 words and sends it to the Controller to print the result.

- Once the result is displayed, all actors shut down...
```

Each actor is a struct and has its own *mailbox* which is a **channel** that receives data of the pre-defined type `Message`. A `Message` is basically a slice of `any` type. I used `any` because we need to send several different types through messages. 

Each actor implements the `Actor` interface which has two functions: `Run` and `SendToMailbox`. 

`Run` creates a loop that iterates through the channel and dispatches each message; the loop, and hence the goroutine, is blocked until new messages arrive. 

`SendToMailbox` takes a receiver, which is an actor, and a message, and adds the message to the mailbox, aka the channel, of the receiver.

To manage these four goroutines together, I used a `WaitGroup` which blocks the main thread until all goroutines finish their jobs.