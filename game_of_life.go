package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
	"os"
	"os/exec"
)

//Matrix defines the structure
type Matrix struct {
	layer [][]bool
	width, height int
}

func ClearScreen() { 
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//IsAlive check if a given cell is alive
// func (matrix *Matrix) IsAlive(x, y int) bool {
// 	return matrix.layer[(x+matrix.width)%matrix.width][(y+matrix.height)%matrix.height]
// }

func countNeighbors(matrix Matrix, linha, coluna int) int {
	count := 0

	if coluna > 0 {
		if matrix.layer[linha][coluna - 1] { count++ }

		if linha > 0 {
			if matrix.layer[linha - 1][coluna - 1] { count++ }
		}

		if linha < matrix.height - 1 {
			if matrix.layer[linha + 1][coluna - 1] { count++ }
		}
	}

	if coluna < matrix.width - 1 {
		if matrix.layer[linha][coluna + 1] { count++ }

		if linha > 0 {
			if matrix.layer[linha - 1][coluna + 1] { count++ }
		}

		if linha < matrix.height - 1 {
			if matrix.layer[linha + 1][coluna + 1] { count++ }
		}
	}

	if linha > 0 {
		if matrix.layer[linha - 1][coluna] { count++ }
	}

	if linha < matrix.height - 1 {
		if matrix.layer[linha + 1][coluna] { count++ }
	}

	return count
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

	for linha := 0; linha < matrix.height; linha++ {
		for coluna := 0; coluna < matrix.width; coluna++ {
			qtdNeighbors := countNeighbors(matrix, linha, coluna)
			
			if !matrix.layer[linha][coluna] && qtdNeighbors == 3 { // dead cell //pass
				newLayer[linha][coluna] = true
			} else if matrix.layer[linha][coluna] { // living cell
				if qtdNeighbors < 2 { // loneliness
					newLayer[linha][coluna] = false
				}
	
				if qtdNeighbors > 3 { // superpopulation
					newLayer[linha][coluna] = false
				}
			}
		}
	}

	return newLayer
}

func (matrix *Matrix) String() string {
	var buffer bytes.Buffer

	for linha := 0; linha < matrix.height; linha++ {
		for coluna := 0; coluna < matrix.width; coluna++ {
			if matrix.layer[linha][coluna] {
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
		time.Sleep(time.Second)
	}
}