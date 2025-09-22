# CH25: Persistent Tables

## Constraints of this style:

1. The data exists beyond the execution of programs that use it, and is meant to be used by many different programs.
2. The data is stored in a way that makes it easier/faster to explore.
3. The problem is solved by issuing queries over the data.


## Further explanation

This style is about storing data in databases. That's it.

The goal is store data in a way that makes it easy to retrieve and reuse later by current program or other programs. Instead of considering data as something you consume once, you store the data into a structured database. And, if you want to *mine* any information about the data, you can write sql queries to do that.


## How I implemented the term frequence task with this style in Go?

I created three tables: documents, words, and chars, and wrote sql queries to insert records to each table and to retrieve the words-frequences sorted in descending order according to the frequency. I could inline these sql queries in my Go code, but I separated them in a different package and used sqlc to generate code from these queries for the sake of keeping things clean and because I don't know how to actually execute sql queries directly... Uhm.

I first read the file and extracted the words from it (after normalizing them of course). Then I store these words in the database.

However, executing the code on the original input file takes very long time to finish; it was so slow and I had to wait for like 20+ minutes so I can see the output, but I didn't.