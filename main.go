package main

import (
	"structs/zoo"
)

func main() {
	cage1 := zoo.cage{openingState: false}
	a1 := zoo.Animal{name: "Sloth", catchingState: false, cage1}
	a2 := zoo.Animal{name: "Lion", catchingState: false, cage1}
	a3 := zoo.Animal{name: "Tiger", catchingState: false, cage1}
	a4 := zoo.Animal{name: "Wolf", catchingState: false, cage1}
	a5 := zoo.Animal{name: "Panthera", catchingState: false, cage1}
	a6 := zoo.Animal{name: "Capybara", catchingState: false, cage1}
	zoo.say(zoo.run(a1))
}
