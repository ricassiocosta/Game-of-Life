package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

//Matrix defines the structure
type Matrix struct {
	layer [][]bool
	width, height   int
}

//IsAlive check if a given cell is alive
func (matrix *Matrix) IsAlive(x, y int) bool {
	return matrix.layer[(x+matrix.width)%matrix.width][(y+matrix.height)%matrix.height]
}

//NextGen generate the Game of Life's next generation of cells
func (matrix *Matrix) NextGen(x, y int) bool {
	
}

//Populate the matrix 
func (matrix *Matrix) Populate() {
	
}

func (matrix *Matrix) String() string {
	var buffer bytes.Buffer
	for y := 0; y < matrix.height; y++ {
		for x := 0; x < matrix.width; x++ {
			if matrix.IsAlive(x, y) {
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
	// init
	matrix := make([][]bool, width)
	for i := range matrix {
		matrix[i] = make([]bool, height)
	}
	// randomize
	n := (width * height) / 2
	for i := 0; i < n; i++ {
		matrix[rand.Intn(width)][rand.Intn(height)] = true
	}
	return matrix
}

//InitLayer define the matrix format
func InitLayer(width, height int) *Matrix {
	return &Matrix {
		layer: Init2dLayer(width, height),
		width:  width,
		height:  height,
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	matrix := InitLayer(40, 10)

	for i := 0; i < 3600; i++ {
		matrix.Populate()
		fmt.Print(matrix)
		time.Sleep(time.Second / 60)
	}
}
