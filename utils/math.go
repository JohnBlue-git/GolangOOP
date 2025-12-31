package utils

import "math"

// Round rounds a float64 to the specified number of decimal places
func Round(value float64, decimals int) float64 {
	multiplier := math.Pow(10, float64(decimals))
	return math.Round(value*multiplier) / multiplier
}

// Max returns the maximum of two float64 values
func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum of two float64 values
func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// Average calculates the average of a slice of float64 values
func Average(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

// Sum calculates the sum of a slice of float64 values
func Sum(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum
}

// isPositive is an unexported (private) function
// Only accessible within the utils package
func isPositive(value float64) bool {
	return value > 0
}

// ValidatePositive checks if a value is positive
// This exported function uses the private isPositive function
func ValidatePositive(value float64) bool {
	return isPositive(value)
}
