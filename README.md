# String Challenge

[![Build Status](https://travis-ci.org/crerwin/stringchallenge.svg?branch=master)](https://travis-ci.org/crerwin/stringchallenge)

## Execution
You'll need Go installed obviously.  
1. Clone into your `%GOPATH%/src` or `%GOPATH\src` on Windows.
```
cd %GOPATH%/src
git clone https://github.com/crerwin/stringchallenge.git
```
2. Build the code
```
cd %GOPATH%
go install github.com/crerwin/stringchallenge
```
3. Execute the code, using quotes around the input
```
bin/stringchallenge "(input,that(looks,like),this)"
```

Here's sample output:
```
~/code/go $ bin/stringchallenge "(id,created,employee(id,firstname,employeeType(id),lastname),location)"
Non-alphabetized:

id
created
employee
- id
- firstname
- employeeType
-- id
- lastname
location

Alphabetized:

created
employee
- employeeType
-- id
- firstname
- id
- lastname
id
location
~/code/go $
```

## Overview
This is an over-engineered solution to the following challenge:
```
Problem to Solve
Convert the string:
"(id,created,employee(id,firstname,employeeType(id),lastname),location)"

to the following output

id
created
employee
- id
- firstname
- employeeType
-- id
- lastname
location

Bonus (output in alphabetical order):

created
employee
- employeeType
-- id
- firstname
- id
- lastname
id
location
```

## Assumptions
In writing this code, I've made a few assumptions.

### Support lots of different input
I'm assuming we'd want to be flexible here, so we'll work with any valid input.

### What constitutes valid input
- Input must begin with `(`, end with `)`, and have no spaces (note, the original input was given with a space, but I assumed it was a pasting issue).
- All opening parentheses must have a corresponding closing parentheses and vice versa.  The first open parentheses shouldn't close until the end.

### This data format is for the exercise only
The data would fit better into JSON, and then could be marshalled/unmarshalled, stored in a NoSQL database directly, etc.  Also our output kind of looks like YAML which would also be better.

## Process
I generally wrote tests first, and then wrote code to pass the tests.  The Item code has 100% test coverage.  If the cli did more than one thing, I'd use https://github.com/mitchellh/cli.
