package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

//Matrix defines the structure
type Matrix struct {
	layer [][]bool
	width, height int
}

//ClearScreen when called
func ClearScreen() { 
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//IsAlive check if a given cell is alive
func (matrix *Matrix) IsAlive(x, y int) bool {
	return matrix.layer[(x+matrix.height)%matrix.height][(y+matrix.width)%matrix.width]
}

//countNeighbors returns total of cell neighbors
func countNeighbors(matrix Matrix, line, column int) int {
	neighbors := 0

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				// don't count self
				continue
			}
			if matrix.IsAlive(line+dx, column+dy) { 
				neighbors++
			}
		}
	}

	return neighbors
}

func copyLayer(matrix Matrix) [][]bool {
	newLayer := make([][]bool, len(matrix.layer))
	
	for i := range matrix.layer {
    newLayer[i] = make([]bool, len(matrix.layer[i]))
    copy(newLayer[i], matrix.layer[i])
	}

	return newLayer
}

//NextGen generate the Game of Life's next generation of cells
func NextGen(matrix Matrix) [][]bool {
	newLayer := copyLayer(matrix)

	for line := 0; line < matrix.height; line++ {
		for column := 0; column < matrix.width; column++ {
			neighbors := countNeighbors(matrix, line, column)

			if matrix.layer[line][column] {
				if neighbors < 2 || neighbors > 3 { 
					// loneliness || superpopulation
					newLayer[line][column] = false
				}
			} else {
				if neighbors == 3 {
					newLayer[line][column] = true //revives
				}
			}
		}
	}

	return newLayer
}

func (matrix *Matrix) String() string {
	var buffer bytes.Buffer

	for line := 0; line < matrix.height; line++ {
		for column := 0; column < matrix.width; column++ {
			if matrix.layer[line][column] {
				buffer.WriteString("*")
			} else {
				buffer.WriteString(" ")
			}
		}

		buffer.WriteString("\n")
	}

	return buffer.String()
}

//Init2dLayer create the matrix and fill layer with random values
func Init2dLayer(height, width int) [][]bool {
	matrix := make([][]bool, height)
	for i := range matrix {
		matrix[i] = make([]bool, width)
	}

	n := (width * height) / 2
	for i := 0; i < n; i++ {
		matrix[rand.Intn(height)][rand.Intn(width)] = true
	}

	return matrix
}

//InitLayer define the matrix format
func InitLayer(height, width int) *Matrix {
	return &Matrix {
		layer: Init2dLayer(height, width),
		height: height,
		width: width,
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	matrix := InitLayer(20, 80)

	for i := 0; i < 100; i++ {
		ClearScreen()
		matrix.layer = NextGen(*matrix)
		fmt.Print(matrix)
		time.Sleep(time.Second/10)
	}
}