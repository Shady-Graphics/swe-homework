package compare

import (
	"testing"
)

func TestCompareVersions(t *testing.T) {
	testCases := []struct {
		version1 string
		version2 string
		expected int
	}{
		{"0.1", "1.1", -1},
		{"1.1", "1.2", -1},
		{"1.2", "1.2.9.9.9.9", -1},
		{"1.2.9.9.9.9", "1.3", -1},
		{"1.3", "1.3.4", -1},
		{"1.3.4", "1.10", -1},
		{"1.10", "1.10", 0},
		{"2.0", "1.11", 1},
		{"2.0", "2", 0},
		{"3.5.2", "3.5.2", 0},
		{"3.5.2", "3.5.1", 1},
	}

	for _, testCase := range testCases {
		result := Versions(testCase.version1, testCase.version2)
		if result != testCase.expected {
			t.Errorf("Comparison failed for versions %s and %s. Expected: %d, Got: %d",
				testCase.version1, testCase.version2, testCase.expected, result)
		}
	}
}
