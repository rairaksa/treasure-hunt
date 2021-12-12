package main

import (
	"fmt"
	"time"
)

type coordinate struct {
	y, x int
}

var grid[6][8]string // store maps of treasure hunt
var current_coordinate[2]int // store current coordinate, updated on every step
var last_top_coordinate[2]int // store last top movement, use when right direction has no move
var last_right_coordinate[2]int // store last right movement, use when bottom direction has no move
var last_bottom_coordinate[2]int // unnecessary
var start_direction string // store first move direction
var treasure_coordinate = []coordinate{} // store treasure coordinate


func main() {
	initiate()
	print()
	find(start_direction)
	print()
	fmt.Print("Treasure Coordinate : ")
	fmt.Println(treasure_coordinate)
}

// initiate config
func initiate() {
	// 1. init grid
	// set all coordinate with obstacle ("#")
	for x, b := range grid {
		for y, _ := range b {
			grid[x][y] = "#"
		}
	}
	// set clear path
	grid[1][1] = "."
	grid[1][2] = "."
	grid[1][3] = "."
	grid[1][4] = "."
	grid[1][5] = "."
	grid[1][6] = "."
	grid[2][1] = "."
	grid[2][5] = "."
	grid[2][6] = "."
	grid[3][1] = "."
	grid[3][2] = "."
	grid[3][3] = "."
	grid[3][5] = "."
	grid[4][3] = "."
	grid[4][4] = "."
	grid[4][5] = "."
	grid[4][6] = "."

	// 2. init current position
	grid[4][1] = "X"
	current_coordinate[0] = 4
	current_coordinate[1] = 1

	// 3. init start direction
	start_direction = "top"
}

// recursive function to find all possible move
func find(direction string) {
	// print change of grid on each step
	print()
	// add delay to see movement of treasure hunter
    time.Sleep(1 * time.Second)

	if(check_direction(direction)) {
		set_coordinate(direction)
		store_last_step_coordinate(direction)
		mark_current_coordinate()
		if(direction == "top") {
			find("right")
		} else if(direction == "right") {
			find("bottom")
		} else if(direction == "bottom") {
			// mark coordinate with string "$"
			grid[current_coordinate[0]][current_coordinate[1]] = "$"

			// store treasury coordinate
			treasure_coordinate = append(treasure_coordinate, coordinate{current_coordinate[0], current_coordinate[1]})

			find("bottom")
		}
	} else {
		if(direction == "top") {
			// stop
		} else if(direction == "right") {
			// back to last top coordinate
			current_coordinate[0] = last_top_coordinate[0]
			current_coordinate[1] = last_top_coordinate[1]
			find("top")
		} else if(direction == "bottom") {
			current_coordinate[0] = last_right_coordinate[0]
			current_coordinate[1] = last_right_coordinate[1]
			find("right")
			// back to last right coordinate
		}
	}
}

// function to check move is possible or not
func check_direction(direction string) bool {
	value := false

	if(direction == "top" && grid[current_coordinate[0] - 1][current_coordinate[1]] == ".") {
		value = true
	} else if(direction == "right" && grid[current_coordinate[0]][current_coordinate[1] + 1] == ".") {
		value = true
	} else if(direction == "bottom" && grid[current_coordinate[0] + 1 ][current_coordinate[1]] == ".") {
		value = true
	}

	return value
}

// function to set current coordinate on each move
func set_coordinate(direction string) {
	if(direction == "top") {
		current_coordinate[0] = current_coordinate[0] - 1; // y axis go top
	} else if(direction == "right") {
		current_coordinate[1] = current_coordinate[1] + 1; // x axis go right
	} else if(direction == "bottom") {
		current_coordinate[0] = current_coordinate[0] + 1; // y axis go bottom
	}
}

// store last step on each direction, used to previous direction when current direction is not possible (eg. when bottom direction stuck, current_coordintae should back to last step on right direction)
func store_last_step_coordinate(direction string) {
	if(direction == "top") {
		last_top_coordinate[0] = current_coordinate[0]
		last_top_coordinate[1] = current_coordinate[1]
	} else if(direction == "right") {
		last_right_coordinate[0] = current_coordinate[0]
		last_right_coordinate[1] = current_coordinate[1]
	} else if(direction == "bottom") {
		last_bottom_coordinate[0] = current_coordinate[0]
		last_bottom_coordinate[1] = current_coordinate[1]
	}
}

// additional function to update current coordinate to o (mean the coordinate has been explored)
func mark_current_coordinate() {
	grid[current_coordinate[0]][current_coordinate[1]] = "O"
}

// function to print current state of map/grid
func print() {
	// init all coordinate with obstacle ("#")
	for y, b := range grid {
		for x, _ := range b {
			fmt.Print(grid[y][x])
		}
		fmt.Println()
	}

	fmt.Println() // give margin
}