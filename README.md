# Golang

## Task 1 : User Form 

### Problem Statement

Create user form using html and css.

### Running Code
1. Clone this repo
3. cd Task1
4. start live server on index.html

## Task 2 : Table

### Problem Statement

Create Table Using Html and css.

### Running Code
1. Clone this repo
3. cd Task2
4. start live server on index.html

## Task 3 : Sudoku Solver

### Problem Statement

Given a filled 9x9 grid, your task is to check whether the given sudoku is solvable or not. The input for your program will be a 9x9 matrix where empty cells are represented by 0, and the filled cells are represented by their respective digits.

### Instructions 

1. Create a Golang program that takes the input Sudoku grid and returns whether the given sudoku is solvable or not.
2. If it fails then also return the row and column for which it fails
3. Test your solution with different Sudoku puzzles to validate its correctness.

### Solution

This Problem can be solved using recursion and backtracking. Below are the steps that we can follow to find the solution of above problem.

1. Check if given sudoku is valid. If Yes then proceed to step 2 else return "Invalid Sudoku Provided" <br>
2. For Finding If Given Sudoku is solvable or not, we actually need to solve the sudoku. Hence lets try solving the sudoku by entering values from 1 to 9 in sudoku cell one by one and doing trial and error mechanism using backtracking. If All trials gone wrong then we will return false else there is always a chance to find a right solved sudoku using given algorithm. Below is the Algorithm for same.
3. Iterate sudoku cells from top left to bottom right in row by row fashion.( for eg. 00, 01,..,08,10,11,....and till end of the sudoku cell which is 88 ) ( Here 00 describes row number 0 and column number 0 of sudoku matrix )
4. If Given Cell is not Empty ( Value is already been provided by puzzle maker ) THEN go on to check next cells without doing anything. Also if given cell is last then we got answer and hence return true.
5. If Given Cell is Empty THEN try all possible values from 1 to 9 inclusive. For Any Given Value, Check If we can insert it to sudoku -> If Yes then insert it and check if it produces a valid sudoku or not. -> If it produces valid sudoku then immediatly return true as we got our answer. ( we only want single solution ha ha). If it doesn't produces valid sudoku then backtrack from this value and try next value till 9 and if nothing works then return false.
6. Hence In this way we can check if given sudoku is solvable or not.

### Running Code
1. Clone this repo
2. Go to task3SudokuSolver Branch
3. cd Task3
4. go run main.go


## Task 4 : Merging JSON Data

### Problem Statement

Given 3 files. user.json, contact.json and tech.json . Our task is to merge the data from all 3 files and give a desired output as final-output.json .

### Solution

- Read All 3 files and store it in structures using unmarshaling JSON.
- Then do simple join operation between all given data.
- Marshal given data into result.json file

### Running Code
1. Clone this repo
2. cd Task4
3. go run main.go

Final output can be seen in result.json file which will be created after running above steps.

## Task 5 : Working With net/http Package

### Problem Statement

Utilise Previously Made UserForm In Task1 to make a post request on server which is made using net/http package. And Store the data in postgresql database. Also Show All Users Data on separate page.

### Solution

- Create a http server using net/http package.
- Use FileServer for handling static data.
- Initialize Table in database.
- Write CRUD operations using sql package.
- Define Routes using HandleFunc function and serve html templates using html/template library.

### Running Code
1. Clone this repo
2. Configure Postgresql And Change URL of type const in main.go ( add your username and password )
3. Execute initialize.sql once in your database ( it will create table for us)
4. cd Task5
5. go run server.go

It would start our http server and now we can go to `localhost:8080` to view our fuctionality being served. 


