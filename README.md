SecretSanta
===========

Pick your language, pick your secret santa. 

Several different Implementations of secret santa algorithms, each in a different language.

## Contributions Welcome

If you want to write your own secret santa implementation in another language, we're a very welcoming bunch. Send as a pull request and we'll add it!

## Solutions

These implementations solve the secret santa problem in of two ways.

#### Algorithm A:

1. Randomly shuffle the list of participants
2. Each participant receives a gift from the participant that precedes them in the shuffled list
3. Each participant buys a gift for the participant that follows them in the shuffled list

#### Algorithm B:

1. Copy the participants into two lists: buyers and receivers
2. Randomly shuffle receivers; do nothing to buyers
3. Check the value at every position in buyers and make sure that the corresponding value at the same position in receivers is not the same (the buyer is not buying for him/herself)
4. Go back to step (2) if the check in step (3) failed; otherwise we’re done!

Algorithm A isn't able to generate every possible valid combination of assignments, because it cannot produce loops where person A has to buy for person B, and person B has to buy for person A.

Algorithm B is able to generate every valid combination, but will often have to shuffle multiple times to find a correct combination.

## Solution-specific Notes

### Erlang 

This solution contains implementations of both algorithms. `select_rotate` is an implementation of Algorithm A, while `select_shuffle` is an implementation of Algorithm B.

#### To Run

The only pre-requisite is the erlang compiler. To run in the terminal:

    > erl
    > c(santa).
    > santa:select_rotate(["EmployeeA", "EmployeeB", "EmployeeC", "EmployeeD"]).

Or to run the shuffled solution:

    > santa:select_shuffle(["EmployeeA", "EmployeeB", "EmployeeC", "EmployeeD"]).

This function returns an array of tuples matching the secret santa giver to the peson they have to give a gift to. For example:

    > santa:select_rotate(["Angus", "Greg", "Mike", "Isabel", "Lewis", "Shannon", "Ally", "Rob"]).
    [{"Angus","Shannon"}, {"Shannon","Lewis"}, {"Lewis","Greg"}, {"Greg","Mike"}, {"Mike","Rob"}, 
    {"Rob","Ally"}, {"Ally","Isabel"}, {"Isabel","Angus"}]

### Bash

The Bash implementation uses a list of email addresses, representing the Secret Santa participants. It implements Algorithm A.

#### To Run 

Requires bash 4.0, as the `readarray` command is relatively new!

    > ./santa.sh

This script will output the list selections to the terminal, but will also attempt to use the `mail` command to email each person their recipient. To make it a surprise, simply comment out the line in the `actionSanta` function that `echo`'s to the terminal.  

### Rust

This solution implements a variant on Algorithm B. It works sequentially picking pairs and validating them. It backtracks and picks again if we pick somebody for themselves. If we reach a situation where we are left with only one person picking for themselves we restart.

#### To Run 

First install rust (http://www.rust-lang.org/), then run with:

    rust run secret-santa.rs

### Python

This solution implements Algorithm A.

#### To Run 

First install python (http://www.python.org/download/), then run with:

    > python secretSanta.py 

### Java

This solution implements Algorithm B.

#### To Run

First install Java JDK (http://www.oracle.com/technetwork/java/javase/downloads/jdk7-downloads-1880260.html), then run with:

    > javac SecretSanta.java
    > java SecretSanta [space-separated list of names]
