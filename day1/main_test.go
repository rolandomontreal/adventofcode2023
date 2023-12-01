package main

import "testing"

func Test_findFirstDigit(t *testing.T) {
	testCases := []struct{
		s string
		expected1 string
		expected2 string
	}{
		{
			s: "1abc2",
			expected1: "1",
			expected2: "2",
		},
		{
			s: "pqr3stu8vwx",
			expected1: "3",
			expected2: "8",
		},
		{
			s: "a1b2c3d4e5f",
			expected1: "1",
			expected2: "5",
		},
		{
			s: "treb7uchet",
			expected1: "7",
			expected2: "7",
		},
	}

	for _, tc := range testCases {
		result1 := findFirstDigit(tc.s, true)
		if (result1 != tc.expected1) {
			t.Errorf("'%s' failed, expected: '%s', actual result: '%s'.", tc.s, tc.expected1, result1)
		}
		result2 := findFirstDigit(tc.s, false)
		if (result2 != tc.expected2) {
			t.Errorf("'%s' failed, expected: '%s', actual result: '%s'.", tc.s, tc.expected2, result2)
		}
	}
}