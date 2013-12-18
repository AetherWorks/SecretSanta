SecretSanta
===========

Pick your language, pick your secret santa. 

Several different Implementations of secret santa algorithms, each in a different language.

## How to Run

### Erlang 

The only pre-requisite is the erlang compiler. To run in the terminal:

    > erl
    > c(santa).
    > santa:select(["EmployeeA", "EmployeeB", "EmployeeC", "EmployeeD"]).
  
This function returns an array of tuples matching the secret santa giver to the peson they have to give a gift to. For example:

    > santa:select(["Angus", "Greg", "Mike", "Isabel", "Lewis", "Shannon", "Ally", "Rob"]).
    [{"Angus","Shannon"}, {"Shannon","Lewis"}, {"Lewis","Greg"}, {"Greg","Mike"}, {"Mike","Rob"}, 
    {"Rob","Ally"}, {"Ally","Isabel"}, {"Isabel","Angus"}]

### Bash

Requires bash 4.0, as the `readarray` command is relatively new!

    > ./santa.sh

This script will output the list selections to the terminal, but will also attempt to use the `mail` command to email each person their recipient. To make it a surprise, simply comment out the line in the `actionSanta` function that `echo`'s to the terminal.  
