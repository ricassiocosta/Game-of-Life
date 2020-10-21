package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

//Matrix defines the structure
type Matrix struct {
	layer [][]bool
	width, height int
}

func init(){
	rand.Seed(time.Now().UTC().UnixNano())
}

//IsAlive check if a given cell is alive
func (matrix *Matrix) IsAlive(x, y int) bool {
	if x < 0 || y < 0 || x == matrix.height || y == matrix.width { 
		// return false if cell is out of matrix
		return false
	}
	return matrix.layer[x][y]
}

//CountNeighbors returns total of cell neighbors
func (matrix *Matrix) CountNeighbors(line, column int) int {
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

//NextGen generate the Game of Life's next generation of cells
func (matrix *Matrix) NextGen() bool {
	newLayer := copyLayer(matrix)
	var hasNextGen bool

	for line := 0; line < matrix.height; line++ {
		for column := 0; column < matrix.width; column++ {
			neighbors := matrix.CountNeighbors(line, column)

			if matrix.layer[line][column] {
				if neighbors < 2 || neighbors > 3 { 
					// loneliness || superpopulation
					newLayer[line][column] = false
					hasNextGen = true
				}
			} else {
				if neighbors == 3 {
					newLayer[line][column] = true //revives
					hasNextGen = true
				}
			}
		}
	}

	matrix.layer = newLayer

	return hasNextGen
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

//ClearScreen when called
func ClearScreen() { 
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func copyLayer(matrix *Matrix) [][]bool {
	newLayer := make([][]bool, len(matrix.layer))
	
	for i := range matrix.layer {
    newLayer[i] = make([]bool, len(matrix.layer[i]))
    copy(newLayer[i], matrix.layer[i])
	}

	return newLayer
}

//Init2dLayer create the matrix and fill layer with random values
func Init2dLayer(height, width int) [][]bool {
	matrix := make([][]bool, height)
	for i := range matrix {
		matrix[i] = make([]bool, width)
	}

	n	 := (width * height) / 2
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
	matrix := InitLayer(20, 80)
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, os.Interrupt)

	go func() {
		for hasNextGen := true; hasNextGen; {
			ClearScreen()
			hasNextGen = matrix.NextGen()
			fmt.Print(matrix)
			time.Sleep(time.Second/10)
		}
		
		signals <- os.Interrupt
	}()

	<-signals
	fmt.Println("Exiting...")
}