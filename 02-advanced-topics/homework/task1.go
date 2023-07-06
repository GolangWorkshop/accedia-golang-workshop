package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Shape interface {
	surface() float64
	volume() float64
}

type Cube struct {
	side float64
}

func (c *Cube) surface() float64 {
	return 6 * c.side * c.side
}

func (c *Cube) volume() float64 {
	return math.Pow(c.side, 3)
}

func newCube(side float64) *Cube {
	return &Cube{
		side: side,
	}
}

type Sphere struct {
	radius float64
}

func (s *Sphere) surface() float64 {
	return 4 * math.Pi * math.Pow(s.radius, 2)
}

func (s *Sphere) volume() float64 {
	return 4 * math.Pi * math.Pow(s.radius, 3) / 3
}

func newSphere(radius float64) *Sphere {
	return &Sphere{
		radius: radius,
	}
}

type Pyramid struct {
	edge float64
}

func (p *Pyramid) surface() float64 {
	return math.Sqrt(3) * math.Pow(p.edge, 2)
}

func (p *Pyramid) volume() float64 {
	return math.Pow(p.edge, 3) / (6 * math.Sqrt(2))
}

func newPyramid(edge float64) *Pyramid {
	return &Pyramid{
		edge: edge,
	}
}

func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("Could not read line.")
	}

	// remove \n from input string
	line = strings.Replace(line, "\n", "", -1)

	if len(line) == 0 {
		return "", errors.New("No input provided.")
	}

	return line, nil
}

func main() {
	line, err := readLine()
	if err != nil {
		panic(err)
	}

	parts := strings.Split(line, " ")

	numberOfWines, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	if numberOfWines*2 != len(parts[1:]) {
		fmt.Println(len(parts))
		panic("Incorrect number of wines provided.")
	}

	totalSurfaceArea := 0.0
	totalVolume := 0.0

	for i := 1; i < len(parts); i += 2 {
		shape := parts[i]
		x := parts[i+1]

		parsedX, err := strconv.ParseFloat(x, 64)
		if err != nil {
			fmt.Printf("Could not parse shape %s's size %s", shape, x)
		}

		var actualShape Shape
		switch shape {
		case "cube":
			actualShape = newCube(parsedX)
		case "sphere":
			actualShape = newSphere(parsedX)
		case "pyramid":
			actualShape = newPyramid(parsedX)
		default:
			fmt.Printf("Unhandled shape %s encountered.\n", shape)
			continue
		}

		totalSurfaceArea += actualShape.surface()
		totalVolume += actualShape.volume()
	}

	fmt.Printf("%.2f %.2f\n", totalSurfaceArea, totalVolume)
}
