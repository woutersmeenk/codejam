package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	// Arrange
	data := bufio.NewScanner(strings.NewReader("2 2\r\n8 1 4 5 2 3 7 6"))
	data.Split(bufio.ScanWords)

	// Act
	courtiers, numRows, numColumns, err := parse(data)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	if numColumns != 2 {
		t.Errorf("numColomns was %v expected 2", numColumns)
	}
	if numRows != 2 {
		t.Errorf("numRows was %v expected 2", numRows)
	}
	if len(courtiers) != 8 {
		t.Fatalf("courtiers length was %v expected 8", len(courtiers))
	}
	expectedCourtiers := []int{7, 2, 1, 4, 3, 6, 5, 0}
	for i := 0; i < len(courtiers); i++ {
		if courtiers[i] != expectedCourtiers[i] {
			t.Errorf("Incorrectly parsed as index %v, result was: %v", i, courtiers)
		}
	}
}

func TestExample1(t *testing.T) {
	// Arrange
	courtiers := []int{3, 2, 1, 0}

	// Act
	garden, err := solve(courtiers, 1, 1)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	checkOutput(t, [][]Tile{[]Tile{goLeft}}, garden)
}

func TestExample2(t *testing.T) {
	// Arrange
	courtiers := []int{7, 6, 3, 2, 5, 4, 1, 0}

	// Act
	garden, err := solve(courtiers, 1, 3)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	checkOutput(t, [][]Tile{[]Tile{goLeft, goLeft, goRight}}, garden)
}

func TestExample3(t *testing.T) {
	// Arrange
	courtiers := []int{7, 2, 1, 4, 3, 6, 5, 0}

	// Act
	garden, err := solve(courtiers, 2, 2)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	checkOutput(t, [][]Tile{[]Tile{goLeft, goLeft}, []Tile{goRight, goLeft}}, garden)
}

func TestExample4(t *testing.T) {
	// Arrange
	courtiers := []int{2, 3, 0, 1}

	// Act
	garden, err := solve(courtiers, 1, 1)

	// Assert
	if err == nil {
		t.Fatal(err)
	}
	if garden != nil {
		t.Fatal(err)
	}
}

func checkOutput(t *testing.T, expected [][]Tile, actual [][]Tile) {
	if len(expected) != len(actual) {
		t.Fatal("Incorrect length:", actual)
	}
	for x := 0; x <= len(expected); x++ {
		expectedRow := expected[x]
		actualRow := actual[x]
		if len(expectedRow) != len(actualRow) {
			t.Fatal("Incorrect row length", actual)
		}
		for y := 0; y <= len(expectedRow); y++ {
			if expectedRow[y] != actualRow[y] {
				t.Errorf("Incorrect value at (%v, %v): %v", x, y, actualRow[y])
			}
		}
	}
}
