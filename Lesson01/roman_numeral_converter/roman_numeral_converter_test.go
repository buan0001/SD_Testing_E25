package roman_converter

import (
	"testing"
)

// In Roman numbers:
// I          1
// V          5
// X         10
// L         50
// C        100
// D        500
// M       1000
//     Larger values preceding smaller or equal ones keep their value (e.g., MDCCCLXVII = 1000 + 500 + (100 * 3) + 50 + 10 + 5 + (1 * 2) = 1867)
//     When a value precedes a larger one, it means subtraction (e.g, XCIV = (100 – 10) + (5 – 1) = 94)
//     I, X, C, M can be repeated up to 3 times consecutively (e.g., 4 is IV, not IIII)
//     V, L, D cannot be repeated
//     Only I, X, C can be used as subtractive numerals
//     The largest value that can be represented is 3999 (MMMCMXCIX)

func TestCorrectConversions(t *testing.T) {
	testcases := map[string]struct {
		roman_str string
		want      int
	}{
		"mix_of_upper_and_lower": {"mDvI", 1506},
		"CXXV":                   {"CXXV", 125},
		"VIX":                    {"VIX", 14},
		"all lower case":         {"mdiv", 1504},
		"all upper case":         {"XCD", 390},
		"empty string":           {"", 0},
	}
	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			rev, _ := Roman_to_decimal(test.roman_str)
			if rev != test.want {
				t.Errorf("Result: %v, expected %v", rev, test.want)
			}
		})
	}

}
func TestCorrectErrorHandling(t *testing.T) {
	testcases := map[string]struct {
		roman_str string
		want      int
	}{
		"using V as subtraction":                    {"IVC", -1},
		"using D as subtraction":                    {"dm", -1},
		"using L as subtraction":                    {"LM", -1},
		"going above 3999":                          {"MMMM", -1},
		"using same char to subtract and add":       {"ivi", -1},
		"using incorrect upper case characters":     {"MMMA", -1},
		"using incorrect lower case characters":     {"avix", -1},
		"using I to subtract 3 times":               {"iiiv", -1},
		"using the same character 4 times in a row": {"Dxxxx", -1},
		"repeating V":                               {"xvv", -1},
		"repeating L":                               {"DXLL", -1},
		"repeating D":                               {"MCDD", -1},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			rev, _ := Roman_to_decimal(tc.roman_str)
			if rev != tc.want {
				t.Errorf("Result: %v, expected %v", rev, tc.want)
			}
		})

	}
}

// func FuzzSum(f *testing.F) {
// 	testcases := []struct {
//         x, y , want float64
//     }{
// 		{1, 1, 2},
// 		{-10, 4.32, -5.68},
// 		{1, 1, 2},
//     }
//     for _, tc := range testcases {
//         f.Add(tc.x, tc.y)  // Use f.Add to provide a seed corpus
//     }
//     f.Fuzz(func(t *testing.T, x float64, y float64) {
//         result := Sum(x, y)
// 		expected := x + y
//         if expected != result {
//             t.Errorf("expected: %v, result: %v", expected, result)
//         }
//     })
// }
