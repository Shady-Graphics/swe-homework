package compare

import (
	"strconv"
	"strings"
)

// Versions compares two version strings and returns:
//
//	1 if version1 is greater than version2
//	-1 if version1 is less than version2
//	0 if version1 is equal to version2
func Versions(version1, version2 string) int {
	levels1 := strings.Split(version1, ".")
	levels2 := strings.Split(version2, ".")

	len1 := len(levels1)
	len2 := len(levels2)
	maxLen := max(len1, len2)

	for i := 0; i < maxLen; i++ {
		num1 := 0
		if i < len1 {
			num1, _ = strconv.Atoi(levels1[i])
		}

		num2 := 0
		if i < len2 {
			num2, _ = strconv.Atoi(levels2[i])
		}

		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
