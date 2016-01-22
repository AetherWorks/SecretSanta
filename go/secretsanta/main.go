package main

import (
	"log"
	"math/rand"
	"time"
)

var santaList = make([]Pair, 0)
var santaMap = make(map[string]string)

func main() {
	names := []string{"Ally", "Gus",
		"Roger", "Danil", "Isabel", "Pete",
		"Rob", "Renee", "Shannon", "Greg"}
	permutations := make([]Pair, 0)
	shuffledPermutations := make([]Pair, 0)

	// randomly shuffle list of permutations
	rand.Seed(time.Now().UnixNano())
	for len(permutations) > 0 {
		random := rand.Intn(len(permutations))
		shuffledPermutations = append(shuffledPermutations, permutations[random])
		permutations = append(permutations[:random], permutations[random+1:]...)
	}

	giftReceivers := []string{"Ally", "Gus",
                "Roger", "Danil", "Isabel", "Pete",
                "Rob", "Renee", "Shannon", "Greg"}

	giftGivers := []string{"Pete", "Isabel", "Ally", "Gus",
                "Roger", "Danil", "Shannon",
                "Rob", "Renee", "Greg"}

	log.Printf("givers [%v]", giftGivers)
	log.Printf("receivers [%v]", giftReceivers)

	// create all permutations
	for i := 0; i < len(giftGivers); i++ {
		for j := 0; j < len(giftReceivers); j++ {
			if len(santaList) == 2*len(names) {
				break
			}
			pair := Pair{giftGiver: giftGivers[i], giftReceiver: giftReceivers[j]}
			if isValidPair(pair, len(names)) {
				santaList = append(santaList, pair)
				santaMap[pair.giftGiver] = pair.giftReceiver
				break
			}
		}
	}

	// print out results
	for _, pair := range santaList {
		log.Printf("%s --> %s\n", pair.giftGiver, pair.giftReceiver)
	}
}

// shuffle the list of people
func shuffleList(names []string) []string {
	rand.Seed(time.Now().UnixNano())
	namesToShuffle := make([]string, len(names))
	copy(namesToShuffle, names) 
	log.Printf("shuffling [%v]", namesToShuffle)
	shuffledNames := make([]string, 0, len(namesToShuffle))
	for len(namesToShuffle) > 0 {
		random := rand.Intn(len(namesToShuffle))
		shuffledNames = append(shuffledNames, namesToShuffle[random])
		namesToShuffle = append(namesToShuffle[:random], namesToShuffle[random+1:]...)
	}
	return shuffledNames
}

// determine whether or not a giver or receiver in this pair has already been used
func isValidPair(pair Pair, numPeople int) bool {
	for _, santaPair := range santaList {
		if santaPair.giftGiver == pair.giftGiver ||
			santaPair.giftReceiver == pair.giftReceiver {
			return false
		}
	}
	return !validPairCreatesCycles(pair, numPeople)
}

func validPairCreatesCycles(pair Pair, numPeople int) bool {
	allGivers := getAllGivers()
	currentCycleSize := 1
	sumOfCycleSizes := 0
	cycleStart := pair.giftGiver
	mostRecentReceiver := pair.giftReceiver
	for {
		log.Printf("Starting cycle search with giver [%s]", mostRecentReceiver)
		if mostRecentReceiver == cycleStart {
			sumOfCycleSizes+= currentCycleSize
			log.Printf("Found cycle of size [%d]", currentCycleSize)
			currentCycleSize = 1
			if len(allGivers) == 0 {
				break
			}
			for giver, _ := range allGivers {
				cycleStart = giver
				mostRecentReceiver = santaMap[giver]
				log.Printf("End of cycle previous giver [%s] next giver [%s]", giver, mostRecentReceiver)
				delete(allGivers, giver)
				break
			}
		}
		receiver, exists := santaMap[mostRecentReceiver]
		if (!exists) {
			log.Printf("Found incomplete cycle for giver [%s]", mostRecentReceiver)
			// there's still an incomplete cycle
			return false
		}
		delete(allGivers, mostRecentReceiver)
		mostRecentReceiver = receiver
		currentCycleSize++
	}
	log.Printf("Sum of all cycles is [%d]", sumOfCycleSizes)
	return sumOfCycleSizes == numPeople - 1
}

func getAllGivers() map[string]interface{} {
	allGivers := make(map[string]interface{})
	for key, _ := range santaMap {
		allGivers[key] = struct{}{}
	}
	return allGivers
}
		

type Pair struct {
	giftGiver    string
	giftReceiver string
}
