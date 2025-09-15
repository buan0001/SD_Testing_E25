package driverslicense

import (
	"reflect"
	"testing"
)

// Test case R1
func TestPointsLowMistakesLow(t *testing.T) {
	testcases := map[string]struct {
		points, mistakes int
	}{
		"0_pts_0_mistakes":  {0, 0},
		"1_pts_1_mistakes":  {1, 1},
		"83_pts_2_mistakes": {83, 2},
	}
	expected := map[string]bool{
		"failTheoretical": true,
		"pass":            false,
		"failPractical":   false,
		"takeMoreLessons": false,
	}
	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			results, _ := OutcomeOfExamResults(test.points, test.mistakes)
			if !reflect.DeepEqual(results, expected) {
				t.Errorf("expected %v, got: %v", expected, results)
			}
		})
	}
}

// Test case R2
func TestPointsHighMistakesHigh(t *testing.T) {
	testcases := map[string]struct {
		points, mistakes int
	}{
		"85_pts_3_mistakes":  {85, 3},
		"86_pts_4_mistakes":  {86, 4},
		"99_pts_5_mistakes":  {99, 5},
		"100_pts_6_mistakes": {100, 6},
	}
	expected := map[string]bool{
		"failTheoretical": false,
		"pass":            false,
		"failPractical":   true,
		"takeMoreLessons": false,
	}
	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			results, _ := OutcomeOfExamResults(test.points, test.mistakes)
			if !reflect.DeepEqual(results, expected) {
				t.Errorf("expected %v, got: %v", expected, results)
			}
		})
	}
}

// Test case R3
func TestPointsHighMistakesLow(t *testing.T) {
	testcases := map[string]struct {
		points, mistakes int
	}{
		"85_pts_0_mistakes":  {85, 0},
		"86_pts_1_mistakes":  {86, 1},
		"99_pts_2_mistakes":  {99, 2},
		"100_pts_2_mistakes": {100, 2},
	}
	expected := map[string]bool{
		"failTheoretical": false,
		"pass":            true,
		"failPractical":   false,
		"takeMoreLessons": false,
	}
	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			results, _ := OutcomeOfExamResults(test.points, test.mistakes)
			if !reflect.DeepEqual(results, expected) {
				t.Errorf("expected %v, got: %v", expected, results)
			}
		})
	}
}

// Test case R4
func TestPointsLowMistakesHigh(t *testing.T) {
	testcases := map[string]struct {
		points, mistakes int
	}{
		"0_pts_3_mistakes":  {0, 3},
		"1_pts_4_mistakes":  {1, 4},
		"83_pts_4_mistakes": {83, 4},
		"84_pts_4_mistakes": {84, 4},
	}
	expected := map[string]bool{
		"failTheoretical": true,
		"pass":            false,
		"failPractical":   true,
		"takeMoreLessons": true,
	}
	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			results, _ := OutcomeOfExamResults(test.points, test.mistakes)
			if !reflect.DeepEqual(results, expected) {
				t.Errorf("expected %v, got: %v", expected, results)
			}
		})
	}
}

func TestErrors(t *testing.T) {
	testcases := map[string]struct {
		points, mistakes int
	}{
		"negative_both-2":   {-2, -2},
		"negative_both-1":   {-1, -1},
		"negative_points":   {-2, 4},
		"negative_mistakes": {99, -1},
		"too_high_points":   {101, 6},
		"too_high_points+1": {102, 6},
	}
	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			_, err := OutcomeOfExamResults(test.points, test.mistakes)
			if err == nil {
				t.Errorf("expected error from %v, got nil", test)
			}
		})
	}
}
