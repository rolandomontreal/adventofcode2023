package main

import (
	"testing"
)

func Test_getHandTypePt2(t *testing.T) {
	testcases := []struct{
		name string
		hand string
		expectedValue handtype
	}{
		{
			name: "One pair 1",
			hand: "32T3K",
			expectedValue: handTypes[5],
		},
		{
			name: "Four of a kind 1 (with joker)",
			hand: "T55J5",
			expectedValue: handTypes[1],
		},
		{
			name: "Two pairs",
			hand: "KK677",
			expectedValue: handTypes[4],
		},
		{
			name: "Four of a kind 2 (with joker)",
			hand: "KTJJT",
			expectedValue: handTypes[1],
		},
		{
			name: "Four of a kind 3 (with joker)",
			hand: "QQQJA",
			expectedValue: handTypes[1],
		},
	};

	for _, tc := range testcases {
		result := getHandTypePt2(tc.hand)
		rv := result.value
		evv := tc.expectedValue.value
		if result.value != tc.expectedValue.value {
			t.Errorf("Test case '%s' failed, expected value: %d, actual value: %d", tc.name, evv, rv)
		}
	}
}