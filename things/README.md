# CH10: Things style

## Constraints of this style:

1. The larger problem is decomposed into things that make sense for the problem domain.
2. Each thing is a **capsule** of data that exposes procedures to the rest of the world.
3. Data is *never* accessed directly, only through these procedures
4. Capsules can reappropriate procedures defined in other capsules.


## Further explanation

The problem is divided into capsules where each one contains *procedures* and data. The data is hidden and is only accessed through exposed procedures, and this achieves **encapsulation**. These capsules are called objects (data + method) as we know in OOP, where you can access object's data through its methods.

As you can see, this style is not strictly for OOP languages; you can implement this style using classes, but this is not necessary. The core idea of this style is to have procedures sharing (or operating on) hidden data. So, a non-OOP language like Go can implement this style using structs for example. All what you want is creating objects/instances regardless of whether it's through classes or through other ways.

Objects can be swapped as long as their interfaces remain the same, and we've seen this feature in OOP: Polymorphism. 

You can also apply inheritance in this style.


## How I implemented the term frequence task with this style in Go?

I divided the problem into structs where each one is concerned with a specific functionality. 

`DataStorageManager` reads data from the file input and has a method called `Words` that returns the file's content a list of words.

`StopWordsManager` retrieves and stores stop words, and has a method called `IsStopWord` that checks if a given word is a stop word or not. 

`WordFreqManager` stores the frequences of each word after counting them, and has a method that returns a sorted list of pairs.

`WordFrequencyController` is like the driver class where you assemble all your code in. It create an instanc of each defined manager struct and executes the code in its `Run` method.

It's basically the same implementation as the code snippet provided in the chapter, but in *Go-ish* style.