package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type polar struct {
	radius float64
	o      float64
}

type cartesian struct {
	x float64
	y float64
}

const RESULT string = "Polar: radius=%.02f angle=%.02f degrees -- Cartesian: x=%.02f y=%.02f\n"

var prompt string = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}

func createSolver(question chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			polarCoord := <-question
			o := polarCoord.o * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(o)
			y := polarCoord.radius * math.Sin(o)
			answers <- cartesian{x, y}
		}
	}()
	return answers
}

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Printf("Radius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = line[:len(line)-1]
		if numbers := strings.Fields(line); len(numbers) == 2 {
			polars, err := floatForStrings(numbers)
			if err != nil {
				fmt.Fprintln(os.Stderr, "invalid number")
				continue
			}
			questions <- polar{polars[0], polars[1]}
			coord := <-answers
			fmt.Printf(RESULT, polars[0], polars[1], coord.x, coord.y)
		} else {
			fmt.Fprintln(os.Stderr, "invalid input")
		}
	}
	fmt.Println()
}

func floatForStrings(numbers []string) ([]float64, error) {
	var floats []float64
	for _, number := range numbers {
		if x, err := strconv.ParseFloat(number, 64); err != nil {
			return nil, err
		} else {
			floats = append(floats, x)
		}
	}
	return floats, nil
}
