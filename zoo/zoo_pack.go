package zoo

import "fmt"

type zooPlace struct {
	zookeeper
	Animal
	cage
	openingState bool
}

type zookeeper struct {
	animals          [5]Animal
	numberOfCatching int
}

type Animal struct {
	name          string
	catchingState bool
	cage
}

type cage struct {
	openingState bool
}

func run(animal *Animal) string {
	animal.catchingState := false
	say := animal.name + " running!"
	return say
}

func putCatchedAnimal(animal *Animal) string {
	animal.catchingState := true
	say := animal.name + " catched!"
	return say
}

func getCatchedAnimal(keeper *zookeeper, animal *Animal) string {
	keeper.numberOfCatching++
	keeper.animals[keeper.numberOfCatching-1] := animal.name
	say := animal.name + " catched!"
	return say
}

func say(phrase string) {
	fmt.Printf(phrase)
}
