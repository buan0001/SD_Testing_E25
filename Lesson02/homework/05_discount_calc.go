package lesson2

import (
    "fmt"
)

func GetDiscountMultiplier(paymentAmount float64) (float64, error) {
    if paymentAmount < 0 {
        return 1, fmt.Errorf("negative payment amount registered")
    } else if paymentAmount <= 300 {
        return 1, nil
    } else if paymentAmount <= 800 {
        return 0.95, nil
    }
    return 0.9, nil
}