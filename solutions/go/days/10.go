package days

import (
	"fmt"
	"strings"
)

type Day10 struct{}

type Screen struct {
	Items [][]string
}

func (s Screen) readTile(X int, Y int) string {
	return s.Items[Y][X]
}

func (s Screen) getNextTileIndex(currentTile string, currentX int,
	currentY int, prevX int,
	prevY int) (nextX int, nextY int) {

	if currentTile == "|" {
		if prevY < currentY {
			return currentX, currentY + 1
		} else {
			return currentX, currentY - 1
		}
	}
	if currentTile == "-" {
		if prevX < currentX {
			return currentX + 1, currentY
		} else {
			return currentX - 1, currentY
		}
	}
	if currentTile == "J" {
		if prevY == currentY {
			return currentX - 1, currentY
		} else {
			return currentX, currentY - 1
		}
	}
	if currentTile == "L" {
		if prevY == currentY {
			return currentX + 1, currentY
		} else {
			return currentX, currentY - 1
		}
	}
	if currentTile == "F" {
		if prevY == currentY {
			return currentX + 1, currentY
		} else {
			return currentX, currentY + 1
		}
	}
	if currentTile == "7" {
		if prevY == currentY {
			return currentX - 1, currentY
		} else {
			return currentX, currentY + 1
		}
	}

	return nextX, nextY
}

func (d Day10) Part1(input []string) string {
	chars := make([][]string, len(input))
	for i := range input {
		fmt.Println(input[i])
		chars[i] = strings.Split(input[i], "")
	}

	screen := Screen{Items: chars}

	currentTile := "|"
	prevX := 0
	prevY := 2
	currentX := 0
	currentY := 3
	count := 1

	for count < 10 {
		X, Y := screen.getNextTileIndex(currentTile, currentX, currentY, prevX, prevY)
		fmt.Printf("\nX: %v, Y: %v", X, Y)
		prevX = currentX
		prevY = currentY
		currentX = X
		currentY = Y
		currentTile = screen.readTile(currentX, currentY)
		count++
		fmt.Printf("\ncurrentTile: %s, currentX: %v, currentY: %v, prevX: %v, "+
			"prevY: %v\n", currentTile, currentX, currentY, prevX, prevY)
	}

	return fmt.Sprintf("%d", count)
}

func (d Day10) Part2(input []string) string {
	return ""
}

func NewDay10() Day10 {
	return Day10{}
}
