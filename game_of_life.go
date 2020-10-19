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

//countNeighbors returns total of cell neighbors
func countNeighbors(matrix Matrix, line, column int) int {
	neighbors := 0

	if column > 0 {
		if matrix.layer[line][column - 1] { neighbors++ }

		if line > 0 {
			if matrix.layer[line - 1][column - 1] { neighbors++ }
		}

		if line < matrix.height - 1 {
			if matrix.layer[line + 1][column - 1] { neighbors++ }
		}
	}

	if column < matrix.width - 1 {
		if matrix.layer[line][column + 1] { neighbors++ }

		if line > 0 {
			if matrix.layer[line - 1][column + 1] { neighbors++ }
		}

		if line < matrix.height - 1 {
			if matrix.layer[line + 1][column + 1] { neighbors++ }
		}
	}

	if line > 0 {
		if matrix.layer[line - 1][column] { neighbors++ }
	}

	if line < matrix.height - 1 {
		if matrix.layer[line + 1][column] { neighbors++ }
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
			
			if !matrix.layer[line][column] && neighbors == 3 { // dead cell //pass
				newLayer[line][column] = true
			} else if matrix.layer[line][column] { // living cell
				if neighbors < 2 { // loneliness
					newLayer[line][column] = false
				}
	
				if neighbors > 3 { // superpopulation
					newLayer[line][column] = false
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
func Init2dLayer(width, height int) [][]bool {
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
		layer: Init2dLayer(width, height),
		width: width,
		height: height,
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	matrix := InitLayer(10, 40)

	for i := 0; i < 100; i++ {
		ClearScreen()
		fmt.Print(matrix)
		matrix.layer = NextGen(*matrix)
		time.Sleep(time.Second/24)
	}
}