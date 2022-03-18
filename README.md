# conwaysgol: Conway's Game of Life implemented in Go

Basic usage:
```
package main

import (
    "fmt"
    conways "github.com/amcajal/conwaysgol/game"
)

func main() {
    // Create 5x5 board, all cells initialized to Dead state
    board := conways.CreateBoard(5, 5)
    
    // Create Blinker. Each "Set" established coordinates (row, column), and the value of the Cell
    board.Set(2,1,conways.Alive); board.Set(2,2,conways.Alive); board.Set(2,3,conways.Alive)
    
    board.PrintData(); fmt.Printf("\n\n\n");
    
    // Perform the "tick" operation
    board.RunTicks(1)
    board.PrintData(); fmt.Printf("\n\n\n");
    board.RunTicks(1)
    board.PrintData(); fmt.Printf("\n\n\n");
}
```

The previous code simulates the "Blinker". Output should be as follows (0 indicates a dead cell, 1 an alive one)
```
[0 0 0 0 0]
[0 0 0 0 0]
[0 1 1 1 0]
[0 0 0 0 0]
[0 0 0 0 0]



[0 0 0 0 0]
[0 0 1 0 0]
[0 0 1 0 0]
[0 0 1 0 0]
[0 0 0 0 0]



[0 0 0 0 0]
[0 0 0 0 0]
[0 1 1 1 0]
[0 0 0 0 0]
[0 0 0 0 0]
```
