package utils

import "fmt"

func MakeParamsPlaceholdersByStartIndex(count, startIndex int) string {
	res := ""
	actualLegth := startIndex + count
	for i := startIndex; i <= actualLegth; i++ {
		res += fmt.Sprintf("$%d", i)
		if i <= actualLegth-2 {
			res += ", "
		}
	}
	return res
}
