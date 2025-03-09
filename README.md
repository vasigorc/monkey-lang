# Writing an Interpreter in go

This repository walks as my walkthrough the [book](https://www.amazon.ca/dp/3982016118) with the same title, which
I used attending the course at the company that I work for, [Improving Ottawa](https://www.improving.com/locations/ottawa/?region=ca).

[![Book Cover](images/book-cover.jpg)](https://interpreterbook.com/)

## Usage

Unlike in the book, the root module includes the name of the project, followed 
by `monkey`, which is the name of the language that we are implementing. To run
tests, we have to include `monkey` before the module under test. For example, 
to run the tests for the lexer, we would run:

```go
go test -v ./monkey/lexer
=== RUN   TestNextToken
--- PASS: TestNextToken (0.00s)
PASS
ok  	waiig_vasile/monkey/lexer	0.003s
```