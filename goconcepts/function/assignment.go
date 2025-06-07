package function

import (
	"errors"
	"go-program/goconcepts/function"
)

// calculate mean and varience
// x := slice of int {x1, x2,x3 ...}
// square(x) = x * x
// sum(f(x)) = f(x1) + f(x2) + f(x3) + ...
// mean(x) = sum(x) / n
// sum of residual(x) = sum(square(x- mean(x)))
// variance(x) = sum of residual(x) / n


func calculateVariance(numbers []int) (*float64, error) {
	if numbers == nil || len(numbers) == 0 {
		return nil, errors.New("Numbers can't be empty")
	}
	// square function
	squareFunc := func(num float64) float64 {
		return num * num
	}
	n :=  float64(len(numbers))
	// sum function 

	sumFunc := func(f func(float64) float64) float64{
		var sum float64
		for _, num :=  range numbers {
			sum += f(float64(num))
		}
		return sum
	}
	// mean function 

	meanFunc := sumFunc(func(f float64) float64 {
		return f
	})

	// sum of residual

	sumReidualFunc := sumFunc(func(f float64) float64 {
		return squareFunc(f - meanFunc)
	})

	variance := sumReidualFunc / n
	return &variance, nil
}