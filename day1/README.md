# Day One
### 1. Getting Started with Go

#### Hello World
As any good programming tutorial, we will start with a hello world program to get a hang of what are the minimum requirements for running a program.

```go
package main
```
The first line is crucial for having an executable program.

This indicates that the current file is the one that will be executed. We will come back to this concept later on. Just know that without this your program won't run. In fact, go ahead and delete that line for now.

The next line

#### Running a Program

To run the program, you will need to open the terminal and navigate to the folder where the .go file is located.
Now run the command 

```bash
go run hello.go
```
You will get an error saying that it was expecting a package. This is because we removed the first line we mentioned in the previous section. 

Go ahead and write it back and run the program again. This time around you should see the message ```Hello World``` in the terminal.

### 2. Go Data Types

#### Numbers


#### Strings

#### Booleans

### 3. Variables and Constants

#### Variables

#### Constants

# Tasks

## Loan Payment Calculator

We will build a calculator for estimating the monthly payments of a loan.

This estimation is based on the [Loan Amortization](https://www.investopedia.com/terms/a/amortized_loan.asp) principle. Where given a **principal loan amount (P)**, an **interest rate (r)**, the **number of yearly payments (n)**, and the **loan term (t)** we can estimate the **monthly payment amount (p)** and the **total amount paid with interest (I)** with the following formula:

![p](https://github.com/ddkdl/power-academy/)
![I](https://github.com/ddkdl/power-academy/)


* Introduce repl.it [https://repl.it]
* Hello World (To get introduced to Go and have a first working program)
* (Student Loan Calculator?) Some Calculator (To use variables and constants) (Movie Score Aggregator?)
* Word Play?