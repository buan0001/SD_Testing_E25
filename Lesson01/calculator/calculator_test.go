package main

import (
	"testing"
	"math"
)

func TestSumHappyPath(t *testing.T) {
    testcases := []struct {
        x, y , want float64
    }{
		{1, 1, 2},
		{-10, 4.32, -5.68},
		{1, 1, 2},
    }
    for _, tc := range testcases {
        rev := Sum(tc.x, tc.y)
        if rev != tc.want {
                t.Errorf("Reverse: %v, want %v", rev, tc.want)
        }
    }
}

func FuzzSum(f *testing.F) {
	testcases := []struct {
        x, y , want float64
    }{
		{1, 1, 2},
		{-10, 4.32, -5.68},
		{1, 1, 2},
    }
    for _, tc := range testcases {
        f.Add(tc.x, tc.y)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, x float64, y float64) {
        result := Sum(x, y)
		expected := x + y
        if expected != result {
            t.Errorf("expected: %v, result: %v", expected, result)
        }
    })
}

func TestSubtract(t *testing.T) {
    testcases := []struct {
        x, y , want float64
    }{
		{2, 1, 1},
		{-10, 4.32, -14.32},
		{1, 2, -1},
    }
    for _, tc := range testcases {
        rev := Subtract(tc.x, tc.y)
        if rev != tc.want {
                t.Errorf("Reverse: %v, want %v", rev, tc.want)
        }
    }
}

func TestMultiplpy(t *testing.T) {
    testcases := []struct {
        x, y , want float64
    }{
		{1, 1, 1},
		{-10, 4.32, -43.2},
		{0, 0, 0},
		{0, 234123, 0},
    }
    for _, tc := range testcases {
        rev := Multiply(tc.x, tc.y)
        if rev != tc.want {
                t.Errorf("Reverse: %v, want %v", rev, tc.want)
        }
    }
}

func TestDivide(t *testing.T) {
    testcases := []struct {
        x, y , want float64
    }{
		{1, 1, 1},
		{-10, 4.32, -10/4.32},
		{-10, 0, math.Inf(-1)},
		{-12, 4, -3},
    }
    for _, tc := range testcases {
        rev := Divide(tc.x, tc.y)
        if rev != tc.want {
                t.Errorf("Reverse: %v, want %v", rev, tc.want)
        }
    }
}
