// Package game implements Conway's Game of Life
// https://en.wikipedia.org/wiki/Conway's_Game_of_Life
package game

import (
	"fmt"
)

type Board struct {
	Cells        [][]int8
	TotalRows    int
	TotalColumns int
}

// Cell status. Dead and alive are self explanatory. Resurrecting means
// a cell that in the next "tick" of the game is going to be Alive, and
// Dying means a cell that in the next "tick" will be Dead
const (
	Resurrecting = iota - 1
	Dead
	Alive
	Dying
)

// Returns a Board of Rows*Columns cells, all of them in Death state (value 0)
func CreateBoard(rows, columns int) (board Board) {
	board.Cells = make([][]int8, rows)

	for c := 0; c < columns; c++ {
		board.Cells[c] = make([]int8, columns)
	}

	board.TotalRows = rows
	board.TotalColumns = columns
	return
}

// Check the value of a cell
func (board *Board) Get(row, column int) (cellValue int8) {
	// @TODO check coordenates are valid, else, return an error
	cellValue = board.Cells[row][column]
	return
}

func (board *Board) Set(row, column int, cellStatus int8) {
	// @TODO check coordenates are valid, else, return an error
	board.Cells[row][column] = cellStatus
}

func (board *Board) RunTicks(ticks int) {
	for i := 0; i < ticks; i++ {
		board.Tick()
	}
}

func (board *Board) IsCoordinateValid(row, column int) (valid bool) {
	valid = false

	// Indexes go from 0 to Total<Rows/Columns> -1
	if row >= 0 &&
		row < board.TotalRows &&
		column >= 0 &&
		column < board.TotalColumns {
		valid = true
	}

	return
}

func (board *Board) Tick() {
	for row := 0; row < len(board.Cells); row++ {
		for column := 0; column < len(board.Cells[row]); column++ {

			aliveNb := board.CountAliveNb(row, column)

			if board.Get(row, column) == Alive {
				if (aliveNb < 2) || (aliveNb > 3) {
					board.Set(row, column, Dying)
				}
			} else { // Cell is dead
				if aliveNb == 3 {
					board.Set(row, column, Resurrecting)
				}
			}

		}
	}
	board.UpdateBoard()
}

// Check alive neighbours in the row+rowUpdate index, given that only the columns
// checked by the colIncrement value are considered
func (board *Board) checkAliveNbIn(row, rowUpdate, column, colIncrement int) (aliveNbInThis int) {
	rowToCheck := row + rowUpdate

	// For loop start and end index are fixed, because for a specific cell
	// we want to check from the "left" cell (column-1) to the "right" cell (column+1)
	for i := -1; i <= 1; i += colIncrement {
		columnToCheck := column + i
		if board.IsCoordinateValid(rowToCheck, columnToCheck) {
			if board.Get(rowToCheck, columnToCheck) > Dead {
				aliveNbInThis += 1
			}
		}
	}

	return
}

func (board *Board) CountAliveNb(row, column int) (aliveNb int) {
	aliveNb = 0

	// Check the three upper neighbours
	aliveNb += board.checkAliveNbIn(row, -1, column, 1)

	// Check left and right neighbours
	aliveNb += board.checkAliveNbIn(row, 0, column, 2)

	// Check the three bottom neighbours
	aliveNb += board.checkAliveNbIn(row, +1, column, 1)

	return
}

func (board *Board) UpdateBoard() {
	for row := 0; row < len(board.Cells); row++ {
		for column := 0; column < len(board.Cells[row]); column++ {
			if board.Get(row, column) == Resurrecting {
				board.Set(row, column, Alive)
			} else if board.Get(row, column) == Dying {
				board.Set(row, column, Dead)
			}
		}
	}
}

func (board *Board) PrintData() {
	for i := 0; i < len(board.Cells); i++ {
		fmt.Println(board.Cells[i])
	}
}
