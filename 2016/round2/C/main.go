package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type tile rune

const (
	empty   tile = 0
	goLeft  tile = '/'
	goRight tile = '\\'
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
	for x := 0; x < params.numRows; x++ {
		garden[x] = make([]tile, params.numColumns)
	}
	return garden, solveForRange(0, len(params.courtiers)-1, params, garden)
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
		fmt.Printf("1: %v %v\n", courtier, params)
		return location{courtier, -1}, down
	}
	if courtier < params.numColumns+params.numRows {
		fmt.Printf("2: %v %v\n", courtier, params)
		return location{params.numColumns, courtier - params.numColumns}, left
	}
	if courtier < params.numColumns*2+params.numRows {
		fmt.Printf("3: %v %v\n", courtier, params)
		return location{courtier - params.numColumns - params.numRows, params.numRows}, up
	}
	fmt.Printf("4: %v %v\n", courtier, params)
	return location{-1, courtier - params.numColumns*2 - params.numRows}, right
}

func fillInPath(startLoc, endLoc location, startDir direction, garden [][]tile) (err error) {
	fmt.Printf("%v %v %v\n", startLoc, endLoc, startDir)
	dir := startDir
	loc := startLoc
	for {
		loc = determineNewLocation(loc, dir)
		fmt.Printf("fill2: %v\n", loc)
		if loc == endLoc {
			return nil
		}
		fmt.Printf("fill2.1: %v\n", garden)
		if garden[loc.x][loc.y] == empty {
			garden[loc.x][loc.y] = determineTile(dir)
			fmt.Printf("fill3: %v\n", garden)
		}
		dir = determineNewDirection(dir, garden[loc.x][loc.y])
		fmt.Printf("fill4: %v\n", dir)
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

func determineNewDirection(dir direction, tile tile) direction {
	switch tile {
	case goRight:
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
	case goLeft:
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
		return goLeft
	case right:
		return goLeft
	case up:
		return goRight
	case down:
		return goRight
	}
	return empty
}

func printOutput(garden [][]tile) {

}
