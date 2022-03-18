package game

import (
	"testing"
)

var board Board

func TestBoardCreation(t *testing.T) {
	board = CreateBoard(3, 3)
	for row := 0; row < len(board.Cells); row++ {
		for column := 0; column < len(board.Cells[row]); column++ {
			if board.Get(row, column) != 0 {
				t.Logf("FAILURE: Coordinates [%v, %v] are not initialized\n", row, column)
				t.Fail()
			}
		}
	}
}

func TestCellSetting(t *testing.T) {
	var row, column int = 1, 1
	board.Set(row, column, Alive)
	if board.Get(row, column) != Alive {
		t.Logf("FAILURE: Coordinate [%v, %v] was not set to Alive (value 1)\n", row, column)
		t.Fail()
	}
	board.Set(row, column, Dead)
	if board.Get(row, column) != Dead {
		t.Logf("FAILURE: Coordinate [%v, %v] was not set to Dead (value 0)\n", row, column)
		t.Fail()
	}
}

func TestCoordinateValidation(t *testing.T) {
	maxRows := board.TotalRows
	maxColumns := board.TotalColumns

	// Check valid coordinates
	if result := board.IsCoordinateValid(0, 0); result != true {
		t.Logf("FAILURE: Coordinate [%v, %v] SHOULD be valid, but has been detected as %v\n", 0, 0, result)
		t.Fail()
	}

	if result := board.IsCoordinateValid(maxRows-1, maxColumns-1); result != true {
		t.Logf("FAILURE: Coordinate [%v, %v] SHOULD be valid, but has been detected as %v\n", maxRows, maxColumns, result)
		t.Fail()
	}

	if result := board.IsCoordinateValid(0, maxColumns-1); result != true {
		t.Logf("FAILURE: Coordinate [%v, %v] SHOULD be valid, but has been detected as %v\n", 0, maxColumns, result)
		t.Fail()
	}

	if result := board.IsCoordinateValid(maxRows-1, 0); result != true {
		t.Logf("FAILURE: Coordinate [%v, %v] SHOULD be valid, but has been detected as %v\n", maxRows, 0, result)
		t.Fail()
	}

	//--------------------------------------------------------------------------

	// Check WRONG coordenates
	if result := board.IsCoordinateValid(-1, 0); result != false {
		t.Logf("FAILURE: Coordinate [%v, %v] SHOULD be INvalid, but has been detected as %v\n", 0, 0, result)
		t.Fail()
	}

	if result := board.IsCoordinateValid(0, -1); result != false {
		t.Logf("FAILURE: Coordinate [%v, %v] SHOULD be INvalid, but has been detected as %v\n", maxRows, maxColumns, result)
		t.Fail()
	}

	if result := board.IsCoordinateValid(maxRows, 0); result != false {
		t.Logf("FAILURE: Coordinate [%v, %v] SHOULD be INvalid, but has been detected as %v\n", 0, maxColumns, result)
		t.Fail()
	}

	if result := board.IsCoordinateValid(0, maxColumns); result != false {
		t.Logf("FAILURE: Coordinate [%v, %v] SHOULD be INvalid, but has been detected as %v\n", maxRows, 0, result)
		t.Fail()
	}
}

func TestNeighbourCounting(t *testing.T) {
	// Initialize data
	board.Set(2, 1, Alive)
	board.Set(2, 2, Alive)

	if count := board.CountAliveNb(0, 0); count != 0 {
		t.Logf("FAILURE: Coordinate [%v, %v] SHOULD have Zero Alive neighbours, but %v were detected \n", 0, 0, count)
		t.Fail()
	}

	if count := board.CountAliveNb(1, 0); count != 1 {
		t.Logf("FAILURE: Coordinate [%v, %v] SHOULD have 1 Alive neighbours, but %v were detected \n", 1, 0, count)
		t.Fail()
	}

	if count := board.CountAliveNb(1, 1); count != 2 {
		t.Logf("FAILURE: Coordinate [%v, %v] SHOULD have 2 Alive neighbours, but %v were detected \n", 1, 1, count)
		t.Fail()
	}

	// Reset data
	board.Set(2, 1, Dead)
	board.Set(2, 2, Dead)
}

// Checks the game logic is properly implemented by creating a "Blinker",
// a well known oscillating pattern in the game
func TestBlinker(t *testing.T) {
	// Create the blinker
	board.Set(1, 0, Alive)
	board.Set(1, 1, Alive)
	board.Set(1, 2, Alive)

	// Perform the "tick" of the game
	board.RunTicks(1)

	// Check the blinker pattern is correct
	if board.Get(0, 1) != Alive ||
		board.Get(2, 1) != Alive ||
		board.Get(1, 0) != Dead ||
		board.Get(1, 2) != Dead {
		t.Logf("FAILURE: Game logic is not correct. Blinker patter does not work\n")
		t.Fail()
	}

	// Do another "tick"
	board.RunTicks(1)

	// Check again the blinker pattern
	if board.Get(0, 1) != Dead ||
		board.Get(2, 1) != Dead ||
		board.Get(1, 0) != Alive ||
		board.Get(1, 2) != Alive {
		t.Logf("FAILURE: Game logic is not correct. Blinker patter does not work\n")
		t.Fail()
	}
}
