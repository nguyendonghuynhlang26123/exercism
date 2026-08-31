package luhn
import (
    "strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	value := 0
	id_len := len(id)
	if id_len <= 1 {
		return false
	}

	var s string;
	for i := id_len - 1; i >= 0; i-- {
		r := id[i]
		if r == ' ' {
			continue
		}

		v, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}

		reverse_location := id_len - i
		if reverse_location % 2 == 0 {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}

		s = strconv.Itoa(v) + s
		value += v
	}
	return value % 10 == 0
}
