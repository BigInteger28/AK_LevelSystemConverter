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
	var eloLvl2 uint64 = 875
	var level, nodes, elo, upperelo, lowerelo uint64
	var colors = []string {
		"White", "Brown", "Grey", "Taupe", "Olive", "Black", "Yellow", "Green", "Lime", "Mint", "Teal_Green", "Blue", "Dark_Blue", "Red", "Crimson", "Dark_Red", "Ochre", "Orange", "Bright_Lavender", "Purple", "Turquoise", "Cyan", "Pink", "Magenta", "Salmon",
	}
	var leagues = []string {
		"WOOD II","WOOD I","IRON II","IRON I","BRONZE III","BRONZE II","BRONZE I","SILVER III","SILVER II","SILVER I","GOLD III","GOLD II","GOLD I","PLATINUM III","PLATINUM II","PLATINUM I","DIAMOND III","DIAMOND II","DIAMOND I","EMERALD III","EMERALD II","EMERALD I","MASTER III","MASTER II","MASTER I","GRAND MASTER III","GRAND MASTER II","GRAND MASTER I","SUPER GRANDMASTER III","SUPER GRANDMASTER II","SUPER GRANDMASTER I","CHALLENGER III","CHALLENGER II","CHALLENGER I","TOP DOG III","TOP DOG II","TOP DOG I","STAR III","STAR II","STAR I","SUPER STAR III","SUPER STAR II","SUPER STAR I","WORLD TOPPER III","WORLD TOPPER II","WORLD TOPPER I","ELITE III","ELITE II","ELITE I","SUPER ELITE III","SUPER ELITE II","SUPER ELITE I","LEGEND III","LEGEND II","LEGEND I","SUPER LEGEND III","SUPER LEGEND II","SUPER LEGEND I","CHAMPION III","CHAMPION II","CHAMPION I","SUPER LEAGUE III","SUPER LEAGUE II","SUPER LEAGUE I","GOAT III","GOAT II","GOAT I",
	}
	var leagueS string
	for {
		input := getUserInput("\nLevel/Nodes/ELO/tabel tot level (..l / ..n / ..e / ..t): ")
		if input[len(input)-1:] == "l" {
			level, _ = strconv.ParseUint(input[0:len(input)-1], 10, 64)
			nodes = getNodes(level)
			if level > 1 {
				upperelo = ((level-1) * elomultiplier) + eloLvl2
				elo = ((level-2) * elomultiplier) + eloLvl2
			} else {
				upperelo = eloLvl2
				elo = 0
			}		
			if level-1 < 67 {
				leagueS = leagues[level-1]
			}
			fmt.Println("Level", input[:len(input)-1], " Tier", ((level-1)/25)+1, colors[(level-1)%25], " | ", leagueS, " --> ", nodes, "nodes")
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
			if level-1 < 67 {
				leagueS = leagues[level-1]
			}
			lowernodes := getNodes(level)
			uppernodes := getNodes(level + 1)
			if level > 1 {
				lowerelo = ((level - 2) * elomultiplier) + eloLvl2
				elo = lowerelo + ((elomultiplier * (nodes - lowernodes)) / (uppernodes - lowernodes))
			} else {
				lowerelo = 0
				elo = lowerelo + ((eloLvl2 * (nodes - lowernodes)) / (uppernodes - lowernodes))
			}
			fmt.Println(input[:len(input)-1], "Nodes --> Level", level, "Tier", ((level-1)/25)+1, colors[(level-1)%25], " | ", leagueS)
			fmt.Println(elo, "ELO")
		} else if input[len(input)-1:] == "e" {
			elo, _ = strconv.ParseUint(input[0:len(input)-1], 10, 64)
			if elo < eloLvl2 {
				level = 1
			} else {
				level = ((elo - eloLvl2) / elomultiplier) + 2
			}
			if level-1 < 67 {
				leagueS = leagues[level-1]
			}
			lowernodes := getNodes(level)
			uppernodes := getNodes(level+1)
			if level > 1 {
				lowerelo = ((level - 2) * elomultiplier) + eloLvl2
				nodes = lowernodes + (((uppernodes - lowernodes) * (elo - lowerelo)) / elomultiplier)
			} else {
				lowerelo = 0
				nodes = lowernodes + (((uppernodes - lowernodes) * (elo - lowerelo)) / eloLvl2)
			}
			fmt.Println(input[:len(input)-1], "ELO --> Level", level, "Tier", ((level-1)/25)+1, colors[(level-1)%25], " | ", leagueS)
			fmt.Println(nodes, "Nodes")
		} else if input[len(input)-1:] == "t" {
			totlevel, _ := strconv.ParseUint(input[0:len(input)-1], 10, 64)
			var i uint64 = 1
			for ; i <= totlevel; i++ {
				nodes = getNodes(i)
				upperelo := (i * elomultiplier) + eloLvl2
				elo = ((i - 1) * elomultiplier) + eloLvl2
				if i-1 < 67 {
					leagueS = leagues[i-1]
				}
				fmt.Println("\nLevel", i, " Tier", ((i-1)/25)+1, colors[(i-1)%25], " | ", leagueS, " --> ", nodes, "nodes")
				fmt.Println(elo, "ELO -", upperelo-1, "ELO")
			}
		}
	}
}
