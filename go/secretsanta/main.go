package main

import (
	"log"
	"math/rand"
	"time"
)

var santaList = make([]Pair, 0)

func main() {
	names := [10]string{"Ally", "Gus",
		"Roger", "Danil", "Isabel", "Pete",
		"Rob", "Renee", "Shannon", "Greg"}
	permutations := make([]Pair, 0)
	shuffledPermutations := make([]Pair, 0)

	// create all permutations
	for i := 0; i < len(names); i++ {
		for j := i + 1; j < len(names); j++ {
			forwardPair := Pair{santa: names[i], receiver: names[j]}
			backwardPair := Pair{santa: names[len(names)-(i+1)],
				receiver: names[len(names)-(j+1)]}
			permutations = append(permutations, forwardPair, backwardPair)
		}
	}

	// randomly shuffle list of permutations
	rand.Seed(time.Now().UnixNano())
	for len(permutations) > 0 {
		random := rand.Intn(len(permutations))
		shuffledPermutations = append(shuffledPermutations, permutations[random])
		permutations = append(permutations[:random], permutations[random+1:]...)
	}

	// iterate over permutations until full
	for _, pair := range shuffledPermutations {
		if len(santaList) == 2*len(names) {
			break
		}
		if isValidPair(pair) {
			santaList = append(santaList, pair)
		}
	}

	// print out results
	for _, pair := range santaList {
		log.Printf("%s --> %s\n", pair.santa, pair.receiver)
	}
}

// determine whether or not a giver or receiver in this pair has already been used
func isValidPair(pair Pair) bool {
	for _, santaPair := range santaList {
		if santaPair.santa == pair.santa ||
			santaPair.receiver == pair.receiver {
			return false
		}
	}
	return true
}

type Pair struct {
	santa    string
	receiver string
}
