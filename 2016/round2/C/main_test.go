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
		{botToTop},
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
		{botToTop, botToTop, topToBot},
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
		{topToBot, topToBot},
		{topToBot, botToTop},
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

func TestImpossibleWhenFillingPaths(t *testing.T) {
	// Arrange
	courtiers := []int{1, 0, 5, 4, 3, 2}

	// Act
	_, err := solve(caseParams{courtiers, 1, 2})

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
		{topToBot, topToBot, topToBot, topToBot},
		{topToBot, botToTop, botToTop, botToTop},
		{topToBot, botToTop, topToBot, topToBot},
		{topToBot, botToTop, topToBot, botToTop},
	}, garden)
}

func TestWithEmptyCenter(t *testing.T) {
	// Arrange
	courtiers := []int{1, 0, 3, 2, 5, 4, 7, 6, 9, 8, 11, 10, 13, 12, 15, 14}

	// Act
	garden, err := solve(caseParams{courtiers, 4, 4})

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	checkOutput(t, [][]tile{
		{topToBot, botToTop, topToBot, botToTop},
		{botToTop, empty, empty, topToBot},
		{topToBot, empty, empty, botToTop},
		{botToTop, topToBot, botToTop, topToBot},
	}, garden)
}

func TestComplexImpossibleGarden(t *testing.T) {
	// Arrange
	courtiers := []int{15, 14, 13, 5, 6, 3, 4, 8, 7, 10, 9, 12, 11, 2, 1, 0}

	// Act
	garden, err := solve(caseParams{courtiers, 4, 4})

	// Assert
	if err == nil {
		t.Fatalf("Expected error! garden: %v", garden)
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
				t.Errorf("Incorrect value at (%v, %v): %v garden: %v", x, y, actualRow[y], actual)
			}
		}
	}
}
