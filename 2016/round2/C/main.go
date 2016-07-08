package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type tile rune

const (
	empty    tile = 0
	botToTop tile = '/'
	topToBot tile = '\\'
)

type location struct {
	x, y int
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

type caseParams struct {
	courtiers           []int
	numRows, numColumns int
}

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
		params, err := parse(scanner)
		if err != nil {
			panic(err)
		}
		garden, err := solve(params)
		printOutput(garden, err, t+1)
	}
	if scanner.Err() != nil {
		panic(err)
	}
}

func parse(scanner *bufio.Scanner) (params caseParams, err error) {
	scanner.Scan()
	params.numRows, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return params, err
	}
	scanner.Scan()
	params.numColumns, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return params, err
	}
	params.courtiers = make([]int, 2*(params.numRows+params.numColumns))
	for i := 0; i < params.numRows+params.numColumns; i++ {
		scanner.Scan()
		courtier1, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return params, err
		}
		scanner.Scan()
		courtier2, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return params, err
		}
		params.courtiers[courtier1-1] = courtier2 - 1
		params.courtiers[courtier2-1] = courtier1 - 1
	}
	return params, nil
}

func solve(params caseParams) (garden [][]tile, err error) {
	garden = make([][]tile, params.numRows)
	for y := 0; y < params.numRows; y++ {
		garden[y] = make([]tile, params.numColumns)
	}
	err = solveForRange(0, len(params.courtiers)-1, params, garden)
	return garden, err
}

func solveForRange(start, end int, params caseParams, garden [][]tile) (err error) {
	lover := params.courtiers[start]
	if lover < start {
		// Skip. We already processed this one before
		return solveForRange(start+1, end, params, garden)
	}
	if lover > end {
		return fmt.Errorf("lover %v of courtier %v is outside of range [%v, %v]",
			lover, start, start, end)
	}
	if start+1 != lover {
		// Process the ones within this path
		if err := solveForRange(start+1, lover-1, params, garden); err != nil {
			return err
		}
	}
	startLoc, dir := determineCourtierLocationAndDirection(start, params)
	endLoc, _ := determineCourtierLocationAndDirection(lover, params)
	if err := fillInPath(startLoc, endLoc, dir, garden); err != nil {
		return err
	}
	if lover == end {
		return nil
	}
	return solveForRange(lover+1, end, params, garden)
}

func determineCourtierLocationAndDirection(courtier int, params caseParams) (loc location, dir direction) {
	if courtier < params.numColumns {
		return location{courtier, -1}, down
	}
	courtier -= params.numColumns
	if courtier < params.numRows {
		return location{params.numColumns, courtier}, left
	}
	courtier -= params.numRows
	if courtier < params.numColumns {
		return location{params.numColumns - 1 - courtier, params.numRows}, up
	}
	courtier -= params.numColumns
	return location{-1, params.numRows - 1 - courtier}, right
}

func fillInPath(startLoc, endLoc location, startDir direction, garden [][]tile) (err error) {
	dir := startDir
	loc := startLoc
	for {
		loc = determineNewLocation(loc, dir)
		if loc == endLoc {
			return nil
		}
		if !isInsideGarden(loc, garden) {
			return fmt.Errorf("Ended up at the wrong courtier!")
		}
		if garden[loc.y][loc.x] == empty {
			garden[loc.y][loc.x] = determineTile(dir)
		}
		dir = determineNewDirection(dir, garden[loc.y][loc.x])
	}
}

func determineNewLocation(loc location, dir direction) location {
	switch dir {
	case left:
		return location{loc.x - 1, loc.y}
	case right:
		return location{loc.x + 1, loc.y}
	case up:
		return location{loc.x, loc.y - 1}
	case down:
		return location{loc.x, loc.y + 1}
	default:
		return loc
	}
}

func isInsideGarden(loc location, garden [][]tile) bool {
	if loc.x < 0 || loc.y < 0 {
		return false
	}
	if loc.y >= len(garden) {
		return false
	}
	if loc.x >= len(garden[0]) {
		return false
	}
	return true
}

func determineNewDirection(dir direction, tile tile) direction {
	switch tile {
	case topToBot:
		switch dir {
		case left:
			return up
		case right:
			return down
		case up:
			return left
		case down:
			return right
		}
	case botToTop:
		switch dir {
		case left:
			return down
		case right:
			return up
		case up:
			return right
		case down:
			return left
		}
	}
	return dir
}

func determineTile(dir direction) tile {
	switch dir {
	case left:
		return botToTop
	case right:
		return botToTop
	case up:
		return topToBot
	case down:
		return topToBot
	}
	return empty
}

func printOutput(garden [][]tile, err error, caseNumber int) {
	fmt.Printf("Case #%v:\n", caseNumber)
	if err != nil {
		fmt.Println("IMPOSSIBLE")
	} else {

		for y := 0; y < len(garden); y++ {
			row := garden[y]
			for x := 0; x < len(row); x++ {
				if row[x] == empty {
					fmt.Print(string(topToBot))
				} else {
					fmt.Print(string(row[x]))
				}
			}
			fmt.Print("\n")
		}
	}
}
