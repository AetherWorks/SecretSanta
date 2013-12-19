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

The Bash implementation uses a list of email addresses, representing the Secret Santa participants. To generalise the approach, imagine asking all participants to stand in a (randomly ordered) circle and buy a present for the person on their left.

To do this in bash there are essentially there are 5 steps:

  * shuffle email addresses
  * read in the emails from the file into an array
  * save the name of the first person
  * iterate over the array, 
     * each person buys for the next person
  * the last person buys for the first person

#### To run 

Requires bash 4.0, as the `readarray` command is relatively new!

    > ./santa.sh

This script will output the list selections to the terminal, but will also attempt to use the `mail` command to email each person their recipient. To make it a surprise, simply comment out the line in the `actionSanta` function that `echo`'s to the terminal.  


### Rust

#### Running
* Install rust (http://www.rust-lang.org/)
* Execute `rust run secret-santa.rs`

#### Algorithm

* Create stack of names (N)
* Create stack of shuffled names (N')
* While N is not empty
 * Pop values off N (Nx) and N' (Nx')
 * Compare Nx and Nx'
   * If same
      * Push values Nx and Nx' onto N and N'
      * Is size of Nx equal to 1
         * Re-run - no solution can be found (we're stuck with one person buying for themselves)
      * Shuffle N'
      * Loop
  * Add pair of names to picked pairs (P)
* Return P
* 


### Python

#### Running
* Install pyhton (http://www.python.org/download/)
* Run
Run:
    > python secretSanta.py


#### Algorithm

The python program is relatively simple and short as the python language provides some great build-in functions. A collections `deque` list contains the names that is first shuffled to produce a random permutation of the names. Since the arrangement of people is unknown we can simply assign a secret santa to one person by selecting the neighbour to the left or right. This can be realized by copying the original list into a new list and using the `rotate` function to shift each name by 1 to the left. 
