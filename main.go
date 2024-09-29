package main

import (
	"fmt"
	"strconv"
)

var levelNodes = [17]uint64{0, 5, 15, 25, 35, 45, 55, 65, 75, 85, 95, 110, 125, 140, 160, 180, 200}

func getUserInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}

func getNodes(level uint64) uint64 {
	if level < 18 {
		return levelNodes[level-1]
	}
	if level < 38 {
		return 200 + ((level - 17) * 40)
	}
	return 1000 + ((level - 37) * 50)
}

func main() {
	var elomultiplier uint64 = 75
	var elostart uint64 = 800
	var level, nodes, elo uint64
	var colors = []string{
		"White", "Grey", "Yellow", "Ochre Yellow", "Salmon", "Orange", "Lime", "Mint", "Green", "Teal Green", "Cyan", "Blue", "Dark_Blue", "Pink", "Magenta", "Bright Lavender", "Purple", "Indigo", "Olive", "Taupe", "Brown", "Red", "Crimson", "Dark_Red", "Black",
	}
	for {
		input := getUserInput("\nLevel/Nodes/ELO/tabel tot level (..l / ..n / ..e / ..t): ")
		if input[len(input)-1:] == "l" {
			level, _ = strconv.ParseUint(input[0:len(input)-1], 10, 64)
			nodes = getNodes(level)
			upperelo := (level * elomultiplier) + elostart
			elo = ((level - 1) * elomultiplier) + elostart
			fmt.Println("Level", input[:len(input)-1], " Tier", ((level-1)/25)+1, colors[(level-1)%25], " --> ", nodes, "nodes")
			fmt.Println(elo, "ELO -", upperelo-1, "ELO")
		} else if input[len(input)-1:] == "n" {
			nodes, _ = strconv.ParseUint(input[0:len(input)-1], 10, 64)
			if nodes <= 200 {
				var i uint64 = 17
				for ; i > 0; i-- {
					if nodes >= levelNodes[i-1] {
						level = i
						break
					}
				}
			} else if nodes < 1000 {
				level = 17 + ((nodes - 200) / 40)
			} else {
				level = 37 + ((nodes - 1000) / 50)
			}
			lowernodes := getNodes(level)
			uppernodes := getNodes(level + 1)
			lowerelo := ((level - 1) * elomultiplier) + elostart
			elo = lowerelo + ((elomultiplier * (nodes - lowernodes)) / (uppernodes - lowernodes))
			fmt.Println(input[:len(input)-1], "Nodes --> Level", level, "Tier", ((level-1)/25)+1, colors[(level-1)%25])
			fmt.Println(elo, "ELO")
		} else if input[len(input)-1:] == "e" {
			elo, _ = strconv.ParseUint(input[0:len(input)-1], 10, 64)
			if elo <= elostart {
				level = 1
			} else {
				level = ((elo - elostart) / elomultiplier) + 1
			}
			lowernodes := getNodes(level)
			uppernodes := getNodes(level + 1)
			lowerelo := ((level - 1) * elomultiplier) + elostart
			nodes = lowernodes + (((uppernodes - lowernodes) * (elo - lowerelo)) / elomultiplier)
			fmt.Println(input[:len(input)-1], "ELO --> Level", level, "Tier", ((level-1)/25)+1, colors[(level-1)%25])
			fmt.Println(nodes, "Nodes")
		} else if input[len(input)-1:] == "t" {
			totlevel, _ := strconv.ParseUint(input[0:len(input)-1], 10, 64)
			var i uint64 = 1
			for ; i <= totlevel; i++ {
				nodes = getNodes(i)
				upperelo := (i * elomultiplier) + elostart
				elo = ((i - 1) * elomultiplier) + elostart
				fmt.Println("\nLevel", i, " Tier", ((i-1)/25)+1, colors[(i-1)%25], " --> ", nodes, "nodes")
				fmt.Println(elo, "ELO -", upperelo-1, "ELO")
			}
		}
	}
}
