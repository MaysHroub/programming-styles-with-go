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


## Project's Structure

<details>
<summary>Click to expand</summary>
  
```
.
├── ch10-things
│   ├── ex10.1
│   │   ├── main.go
│   │   └── thing
│   │       ├── data_storage_mngr.go
│   │       ├── stopwords_mngr.go
│   │       ├── word_freq_contr.go
│   │       └── word_freq_mngr.go
│   ├── ex10.2
│   │   ├── main.go
│   │   └── thing
│   │       ├── data_storage_mngr.go
│   │       ├── informer.go
│   │       ├── stopwords_mngr.go
│   │       ├── word_freq_contr.go
│   │       └── word_freq_mngr.go
│   ├── ex10.3
│   │   ├── main.go
│   │   └── thing
│   │       ├── data_processor.go
│   │       ├── data_reader.go
│   │       └── token_processor.go
│   ├── ex10.5
│   │   ├── main.go
│   │   └── thing
│   │       ├── data_manager.go
│   │       ├── page_processor.go
│   │       └── word_index_mngr.go
│   └── README.md
├── ch24-quarantine
│   ├── ex24.1
│   │   ├── main.go
│   │   └── quarantine.go
│   ├── ex24.3
│   │   ├── main.go
│   │   └── quarantine.go
│   ├── ex24.5
│   │   ├── main.go
│   │   └── quarantine.go
│   └── README.md
├── ch25-persistent-tables
│   ├── ex25.1
│   │   └── main.go
│   ├── ex25.2
│   │   ├── dbio
│   │   │   ├── reader.go
│   │   │   └── writer.go
│   │   └── main.go
│   ├── ex25.4
│   │   ├── dbio
│   │   │   ├── reader.go
│   │   │   └── writer.go
│   │   └── main.go
│   ├── ex25.5
│   │   └── main.go
│   ├── internal
│   │   └── database
│   │       ├── db.go
│   │       ├── models.go
│   │       ├── word_freq_data.sql.go
│   │       └── word_index_data.sql.go
│   ├── README.md
│   ├── sql
│   │   ├── queries
│   │   │   ├── word_freq_data.sql
│   │   │   └── word_index_data.sql
│   │   └── schema
│   │       ├── 001_docs.sql
│   │       ├── 002_words.sql
│   │       ├── 003_stopwords.sql
│   │       ├── 004_pages.sql
│   │       └── testdb.db
│   └── sqlc.yaml
├── ch28-actors
│   ├── ex28.1
│   │   ├── actor
│   │   │   ├── actors_interface.go
│   │   │   ├── data_storage_manager.go
│   │   │   ├── stop_words_manager.go
│   │   │   ├── word_freq_controller.go
│   │   │   └── word_freq_manager.go
│   │   └── main.go
│   ├── ex28.2
│   │   ├── actor
│   │   │   ├── actors_interface.go
│   │   │   ├── data_storage_manager.go
│   │   │   ├── word_freq_controller.go
│   │   │   └── word_freq_manager.go
│   │   └── main.go
│   ├── ex28.4
│   │   ├── actor
│   │   │   ├── actor_interface.go
│   │   │   ├── data_manager.go
│   │   │   ├── page_manager.go
│   │   │   └── word_index_contr.go
│   │   └── main.go
│   └── README.md
├── ch30-mapreduce
│   ├── ex30.1
│   │   └── main.go
│   ├── ex30.2
│   │   └── main.go
│   ├── ex30.3
│   │   └── main.go
│   ├── ex30.4
│   │   └── main.go
│   └── README.md
├── ch3-monolithic
│   ├── ex3.1
│   │   └── main.go
│   ├── ex3.2
│   │   └── main.go
│   ├── ex3.3
│   │   └── main.go
│   ├── exercises.md
│   └── README.md
├── ch5-pipeline
│   ├── ex5.1
│   │   ├── func_test.go
│   │   └── main.go
│   ├── ex5.2
│   │   └── main.go
│   ├── ex5.4
│   │   ├── func_test.go
│   │   └── main.go
│   └── README.md
├── config
│   └── config.go
├── files
│   ├── input.txt
│   ├── lightweightinput.txt
│   ├── repetativewords.txt
│   ├── stopwords.txt
│   └── test.txt
├── go.mod
├── go.sum
├── README.md
└── structure.txt

48 directories, 91 files
```

This is the output of running `tree` command; pretty neat, right?
</details>

## Prerequisites
- Go 1.19+ 
- C compiler (gcc/clang) for SQLite3 driver compilation


## Installation

Make sure Git is installed, and clone the repo:
```sh
git clone https://github.com/MaysHroub/programming-styles-with-go.git

cd programming-styles-with-go/
```


## How to Run

Honestly, I don't know why you would do this. It will give you the same output regardless of the style you choose :)

*Anywaaaaay,* 

To run a specific exercise, include the package name and the exercise directory. For example:

```sh
go run ./ch3-monolithic/ex3.1/
```

Run this from the root directory of the project, otherwise it won't work.

When running files from `ch24-quarantine` package, you need to provide the file name as an input after running the command:
```
./files/input.txt
```


## General Output
This what the output should be for the **term frequency** task given `input.txt` file:

<details>
<summary>Click to expand</summary>

```
mr  -  786
elizabeth  -  635
very  -  488
darcy  -  418
such  -  395
mrs  -  343
much  -  329
more  -  327
bennet  -  323
bingley  -  306
jane  -  295
miss  -  283
one  -  275
know  -  239
before  -  229
herself  -  227
though  -  226
well  -  224
never  -  220
sister  -  218
soon  -  216
think  -  211
now  -  209
time  -  203
good  -  201
```
</details>


And for the **word index** task:

<details>
<summary>Click to expand</summary>
  
```
word: abatement
pages: [72]

word: abhorrence
pages: [81 118 125 199 226 231]

word: abhorrent
pages: [209]

word: abide
pages: [130 240]

word: abiding
pages: [132]

word: abilities
pages: [51 52 78 115 128 145]

word: able
pages: [12 24 40 55 60 62 63 66 71 73 78 80 88 93 95 96 106 107 113 116 129 132 134 137 139 140 146 154 164 166 170 172 175 176 180 184 186 191 196 197 199 200 203 214 217 224 225 233 238]

word: ablution
pages: [88]

word: abode
pages: [41 46 80 90 96 132 197]

word: abominable
pages: [21 34 50 89 119]

word: abominably
pages: [32 98 203 225]

word: abominate
pages: [199 224]

word: abound
pages: [73]

word: above
pages: [5 21 113 134 146 151 157 159 160 161 164 166 175 179 194 198 210 215]

word: abroad
pages: [145 147 176 217]

word: abrupt
pages: [152]

word: abruptly
pages: [27 115]

word: abruptness
pages: [148 149]

word: absence
pages: [37 39 45 54 55 65 72 73 77 78 81 93 111 129 146 148 154 155 169 175 180 214]

word: absent
pages: [20 149 170 173]

word: absolute
pages: [55 171 191 232]

word: absolutely
pages: [10 16 21 67 68 92 109 124 125 128 142 152 168 183 197 203 225 229]

word: absurd
pages: [42 121 128 224 228]

word: absurdities
pages: [94 164]

word: absurdity
pages: [141]
```
</details>

## Developer's Note

This repository is for educational purposes - it's designed to be studied and referenced, not executed as a runnable application. (Well, it's not that anyone would study it...)

And.. NO, I won't add any documentation.




