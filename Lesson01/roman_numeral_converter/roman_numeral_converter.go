package roman_converter

import (
	"fmt"
)


func Roman_to_decimal(roman_num string) (int, error) {
	valid_characters := map[rune]map[string]int{
		'I': {"value": 1, "repeatable": 1, "subtractive":1},
		'V': {"value": 5, "repeatable": 0, "subtractive":0},
		'X': {"value": 10, "repeatable": 1, "subtractive":1},
		'L': {"value": 50, "repeatable": 0, "subtractive":0},
		'C': {"value": 100, "repeatable": 1, "subtractive":1},
		'D': {"value": 500, "repeatable": 0, "subtractive":0},
		'M': {"value": 1000, "repeatable": 1, "subtractive":0},
	}

	var prev_char rune
	repeats := 1
	sum := 0
	just_subtracted_with := 0
	for _, char := range roman_num {
		if char >= 97 && char <= 122 {
			char -= 32 // Convert from lowercase letters to UPPERCASE (the difference is always 32, following the ascii-table)
		}
		entry, ok := valid_characters[char]
		if !ok {
			return -1, fmt.Errorf("roman_to_decimal: invalid character %c", char)
		}
		val := entry["value"]
		prev_val :=  valid_characters[prev_char]["value"]
		is_subtractive, exists := valid_characters[prev_char]["subtractive"]
		if !exists {
			is_subtractive = 1 // On first run, valid_characters[prev_char] wont exist, but it shouldn't result in error.
			// Could also use "first run flag", which would be more explicit
		}
		switch {
			case val == prev_val:
				switch {
					case entry["repeatable"] == 0:
						return -1, fmt.Errorf("roman_to_decimal: %c cannot be repeated twice in a row", char)
					case repeats == 3:
						return -1, fmt.Errorf("roman_to_decimal: %c cannot be repeated four times in a row", char)
					default:
						repeats++
						sum += val
				}
			case val > prev_val:
				if repeats > 2 {
					return -1, fmt.Errorf("roman_to_decimal: %c cannot be repeated three times in a row when subtracting", prev_char)
					} else if is_subtractive == 0 {
					return -1, fmt.Errorf("roman_to_decimal: cannon use %c for subtraction", prev_char)
				}
				sum += val - (prev_val * repeats * 2) // Since we already added the numbers that we now wish to subtract, we subtract each number twice
				repeats = 1
				just_subtracted_with = prev_val
				
			case just_subtracted_with == val:
				return -1, fmt.Errorf("roman_to_decimal: cannot add %c since it was just used for subtraction", char)
			default:
				repeats = 1
				just_subtracted_with = 0
				sum += val
			}
		if sum > 3999 {
			return -1, fmt.Errorf("roman_to_decimal: the highest representable number is MMMCMXCIX (3999)")
		}
		prev_char = char
	}

	return sum, nil
}