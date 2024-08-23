package main

import (
	"log"

	"mistakes/13_package_naming/stringset"
	"mistakes/13_package_naming/utils"
)

func main() {
	set := utils.NewStringSet("qwerty", "wasd", "amen")
	log.Println(utils.SortStringSet(set))

	// package name would be self-explanatory
	set = stringset.New("qwerty", "wasd", "amen", "111", "666")
	log.Println(stringset.Sort(set))
}
