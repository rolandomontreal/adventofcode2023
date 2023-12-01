package main

import "testing"

func Test_findFirstDigitsPt2FromLeft(t *testing.T) {
	testCases := []struct{
		s string
		expected string
	} {
		{
			s: "two1nine",
			expected: "2",
		},
		{
			s: "eightwothree",
			expected: "8",
		},
		{
			s: "abcone2threexyz",
			expected: "1",
		},
		{
			s: "xtwone3four",
			expected: "2",
		},
		{
			s: "4nineeightseven2",
			expected: "4",
		},
		{
			s: "zoneight234",
			expected: "1",
		},
		{
			s: "7pqrstsixteen",
			expected: "7",
		},
	}
	for _, tc := range testCases {
		result := findFirstDigitsPt2FromLeft(tc.s)
		if (result != tc.expected) {
			t.Errorf("'%s' failed, expected: '%s', actual result: '%s'.", tc.s, tc.expected, result)
		}
	}
}

func Test_findFirstDigitsPt2FromRight(t *testing.T) {
	testCases := []struct{
		s string
		expected string
	} {
		{
			s: "two1nine",
			expected: "9",
		},
		{
			s: "eightwothree",
			expected: "3",
		},
		{
			s: "abcone2threexyz",
			expected: "3",
		},
		{
			s: "xtwone3four",
			expected: "4",
		},
		{
			s: "4nineeightseven2",
			expected: "2",
		},
		{
			s: "zoneight234",
			expected: "4",
		},
		{
			s: "7pqrstsixteen",
			expected: "6",
		},
	}
	for _, tc := range testCases {
		result := findFirstDigitsPt2FromRight(tc.s)
		if (result != tc.expected) {
			t.Errorf("'%s' failed, expected: '%s', actual result: '%s'.", tc.s, tc.expected, result)
		}
	}
}

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