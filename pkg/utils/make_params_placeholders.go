package utils

import "fmt"

func MakeParamPlaceholders(count int) string {
	res := ""
	for i := 0; i < count; i++ {
		res += fmt.Sprintf("$%d", i+1)
		if i <= count-2 {
			res += ", "
		}
	}
	return res
}
