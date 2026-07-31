package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var distance int;
	if len(a) != len(b) {
		return 0, errors.New("inputs have different lengths")
	}
	for i := 0; i < min(len(a), len(b)); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
