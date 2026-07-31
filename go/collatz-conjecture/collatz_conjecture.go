package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("Invalid input")
	}
	if n == 1 {
		return 0, nil
	}
	if n % 2 == 0 {
		value, error := CollatzConjecture(n / 2)
		return value + 1, error
	} else {
		value, error := CollatzConjecture(n * 3 + 1)
		return value + 1, error
	}
}
