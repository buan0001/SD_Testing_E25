package driverslicense

import (
	"fmt"
)

func OutcomeOfExamResults(theoret_pts, mistakes int) (map[string]bool, error) {
	if theoret_pts < 0 || theoret_pts > 100 {
		return nil, fmt.Errorf("theoretical points out of bound (0-100): %d", theoret_pts)
	}
	if mistakes < 0 {
		return nil, fmt.Errorf("negative mistakes: %d", mistakes)
	}

	results := make(map[string]bool, 4)
	results["failTheoretical"] = theoret_pts < 85
	results["failPractical"] = mistakes > 2
	results["takeMoreLessons"] = results["failPractical"] && results["failTheoretical"]
	results["pass"] = !results["failPractical"] && !results["failTheoretical"]
	return results, nil
}
