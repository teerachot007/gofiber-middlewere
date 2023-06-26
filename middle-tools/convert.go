package middletools

import "strconv"

func StringToInt(s string, v int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return v
	}
	return i
}