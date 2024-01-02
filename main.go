package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const (
	normalMode DisplayMode = iota
	heightMode
	defaultDisplayMode = normalMode
	displayGrid        = false
	tileCharacterWidth = 3
)

type stringProcessor func(tile Tile) string
type DisplayMode int
type Tile struct {
	X                int
	Y                int
	Altitude         int
	Terrain          string
	DisplayCharacter string
}
type World struct {
	Height  int
	Width   int
	Tiles   [][]Tile
	tilePad int
}

var modeProcessors = map[DisplayMode]stringProcessor{
	normalMode: normalModeProcessor,
	heightMode: heightModeProcessor,
}

func main() {
	mainWorld := World{}
	mainWorld.Initialize(20, 20)
	mainWorld.Display("normal")
}

func (w *World) Initialize(height, width int) {
	w.Height = height
	w.Width = width
	w.tilePad = 3
	for y := 1; y <= w.Height; y++ {
		currentRow := []Tile{}
		for x := 1; x <= w.Width; x++ {
			currentTile := Tile{
				X:                x,
				Y:                y,
				Terrain:          "grass",
				Altitude:         rand.Intn(50),
				DisplayCharacter: ".",
			}
			currentRow = append(currentRow, currentTile)
		}
		w.Tiles = append(w.Tiles, currentRow)
	}
}

func (w *World) Display(mode string) {
	for _, row := range w.Tiles {
		fmt.Println(w.getRowDisplayString(row, defaultDisplayMode))
	}
}

func (w *World) getRowDisplayString(row []Tile, mode DisplayMode) string {
	rowString := ""
	if displayGrid {
		rowString += "|"
	}
	processor, exists := modeProcessors[mode]
	if !exists {
		return ""
	}
	for _, tile := range row {
		rowString += processor(tile)
		if displayGrid {
			rowString += "|"
		}
	}
	return rowString
}

func normalModeProcessor(tile Tile) string {
	// Implementation for normal mode
	displayString := ""
	for i := 0; i < tileCharacterWidth; i++ {
		displayString += tile.DisplayCharacter
	}
	return displayString
}

func heightModeProcessor(tile Tile) string {
	displayString := ""
	necessaryPrependingZeros := tileCharacterWidth - len(strconv.Itoa(tile.Altitude))
	for i := 0; i < necessaryPrependingZeros; i++ {
		displayString += "0"
	}
	displayString += strconv.Itoa(tile.Altitude)
	return displayString
}

func (w *World) Update() {

}
