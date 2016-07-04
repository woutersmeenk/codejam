package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Tile rune

const (
	empty   Tile = ' '
	goLeft  Tile = '/'
	goRight Tile = '\\'
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bufio.NewReader(file))
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	numTestcases, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	for t := 0; t < numTestcases; t++ {
		courtiers, numRows, numColumns, err := parse(scanner)
		if err != nil {
			panic(err)
		}
		garden, err := solve(courtiers, numRows, numColumns)
		fmt.Printf("Case #%v:\n", t+1)
		if err != nil {
			fmt.Println("IMPOSSIBLE")
		}
		printOutput(garden)
	}
	if scanner.Err() != nil {
		panic(err)
	}
}

func parse(scanner *bufio.Scanner) (courtiers []int, numRows, numColumns int, err error) {
	scanner.Scan()
	numRows, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, 0, 0, err
	}
	scanner.Scan()
	numColumns, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, 0, 0, err
	}
	courtiers = make([]int, 2*(numRows+numColumns))
	for i := 0; i < numRows+numColumns; i++ {
		scanner.Scan()
		courtier1, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, 0, 0, err
		}
		scanner.Scan()
		courtier2, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, 0, 0, err
		}
		courtiers[courtier1-1] = courtier2 - 1
		courtiers[courtier2-1] = courtier1 - 1
	}
	return courtiers, numRows, numColumns, nil
}

func solve(courtiers []int, numRows, numColumns int) (garden [][]Tile, err error) {

	return nil, errors.New("bla")
}

func printOutput(garden [][]Tile) {

}
