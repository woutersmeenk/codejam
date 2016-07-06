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
	params, err := parse(data)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	if params.numColumns != 2 {
		t.Errorf("numColomns was %v expected 2", params.numColumns)
	}
	if params.numRows != 2 {
		t.Errorf("numRows was %v expected 2", params.numRows)
	}
	if len(params.courtiers) != 8 {
		t.Fatalf("courtiers length was %v expected 8", len(params.courtiers))
	}
	expectedCourtiers := []int{7, 2, 1, 4, 3, 6, 5, 0}
	for i := 0; i < len(params.courtiers); i++ {
		if params.courtiers[i] != expectedCourtiers[i] {
			t.Errorf("Incorrectly parsed as index %v, result was: %v", i, params.courtiers)
		}
	}
}

func TestExample1(t *testing.T) {
	// Arrange
	courtiers := []int{3, 2, 1, 0}

	// Act
	garden, err := solve(caseParams{courtiers, 1, 1})

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	checkOutput(t, [][]tile{
		{goLeft},
	}, garden)
}

func TestExample2(t *testing.T) {
	// Arrange
	courtiers := []int{7, 6, 3, 2, 5, 4, 1, 0}

	// Act
	garden, err := solve(caseParams{courtiers, 1, 3})

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	checkOutput(t, [][]tile{
		{goLeft, goLeft, goRight},
	}, garden)
}

func TestExample3(t *testing.T) {
	// Arrange
	courtiers := []int{7, 2, 1, 4, 3, 6, 5, 0}

	// Act
	garden, err := solve(caseParams{courtiers, 2, 2})

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	checkOutput(t, [][]tile{
		{goLeft, goLeft},
		{goRight, goLeft},
	}, garden)
}

func TestExample4(t *testing.T) {
	// Arrange
	courtiers := []int{2, 3, 0, 1}

	// Act
	_, err := solve(caseParams{courtiers, 1, 1})

	// Assert
	if err == nil {
		t.Fatal("Expected error!")
	}
}

func TestComplexGarden(t *testing.T) {
	// Arrange
	courtiers := []int{15, 14, 13, 4, 3, 6, 5, 8, 7, 10, 9, 12, 11, 2, 1, 0}

	// Act
	garden, err := solve(caseParams{courtiers, 4, 4})

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	checkOutput(t, [][]tile{
		{goLeft, goLeft, goLeft, goLeft},
		{goLeft, goLeft, goLeft, goLeft},
		{goLeft, goLeft, goLeft, goLeft},
		{goLeft, goLeft, goLeft, goLeft},
	}, garden)
}

func TestComplexImpossibleGarden(t *testing.T) {
	// Arrange
	courtiers := []int{15, 14, 13, 6, 5, 4, 3, 8, 7, 10, 9, 12, 11, 2, 1, 0}

	// Act
	_, err := solve(caseParams{courtiers, 4, 4})

	// Assert
	if err == nil {
		t.Fatal("Expected error!")
	}
}

func checkOutput(t *testing.T, expected [][]tile, actual [][]tile) {
	if len(expected) != len(actual) {
		t.Fatal("Incorrect length:", actual)
	}
	for x := 0; x < len(expected); x++ {
		expectedRow := expected[x]
		actualRow := actual[x]
		if len(expectedRow) != len(actualRow) {
			t.Fatal("Incorrect row length", actual)
		}
		for y := 0; y < len(expectedRow); y++ {
			if expectedRow[y] != actualRow[y] {
				t.Errorf("Incorrect value at (%v, %v): %v", x, y, actualRow[y])
			}
		}
	}
}
