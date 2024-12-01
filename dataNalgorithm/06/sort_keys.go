package main

import (
	"fmt"
	"sort"
)

type Mass float64
type Miles float64

type Thing struct {
	name          string
	mass          Mass
	distance      Miles
	meltingpoint  int
	freezingpoint int
}

type ByFactor func(Thing1 *Thing, Thing2 *Thing) bool

type ThingSorter struct {
	Things   []Thing
	byFactor func(Thing1 *Thing, Thing2 *Thing) bool
}

func (byFactor ByFactor) Sort(Things []Thing) {
	var sortedThings *ThingSorter
	sortedThings = &ThingSorter{
		Things:   Things,
		byFactor: byFactor,
	}
	sort.Sort(sortedThings)

}

func (ThingSorter *ThingSorter) Len() int {
	return len(ThingSorter.Things)
}

func (ThingSorter *ThingSorter) Swap(i int, j int) {
	ThingSorter.Things[i], ThingSorter.Things[j] = ThingSorter.Things[j], ThingSorter.Things[i]
}

func (ThingSorter *ThingSorter) Less(i int, j int) bool {
	return ThingSorter.byFactor(&ThingSorter.Things[i], &ThingSorter.Things[j])
}

func main() {
	var Things = []Thing{
		{"IronRod", 0.055, 0.4, 3000, -180},
		{"SteelChair", 0.815, 0.7, 4000, -209},
		{"CopperBowl", 1.0, 1.0, 60, -30},
		{"BrassPot", 0.107, 1.5, 10000, -456},
	}

	var name func(*Thing, *Thing) bool
	name = func(Thing1 *Thing, Thing2 *Thing) bool {
		return Thing1.name < Thing2.name
	}

	var mass func(*Thing, *Thing) bool
	mass = func(thing1 *Thing, thing2 *Thing) bool {
		return thing1.mass < thing2.mass
	}

	var distance func(*Thing, *Thing) bool
	distance = func(thing1 *Thing, thing2 *Thing) bool {
		return thing1.distance < thing2.distance
	}

	var decreasingDistance func(*Thing, *Thing) bool
	decreasingDistance = func(p1, p2 *Thing) bool {
		return distance(p2, p1)
	}

	ByFactor(name).Sort(Things)
	fmt.Println("By name:", Things)
	ByFactor(mass).Sort(Things)
	fmt.Println("By mass:", Things)
	ByFactor(distance).Sort(Things)
	fmt.Println("By distance:", Things)
	ByFactor(decreasingDistance).Sort(Things)
	fmt.Println("By decreasing distance:", Things)
}
