package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	// default tests
	result := Add(1, 2)
	assert.Equal(t, result, 3, "they should be equal")
	assert.NotEqual(t, result, 4, "they should not be equal")
}

func TestAddOneShort(t *testing.T) {
	// shorthand notation, advantages: one less "t"
	assert := assert.New(t)

	result := Add(1, 2)
	assert.Equal(result, 3, "they should be equal")
	assert.NotEqual(result, 4, "they should not be equal")
}

func TestAddOneRequire(t *testing.T) {
	// Exits tests if failed

	require.Equal(t, Add(1, 2), 3, "they have to be equal")
	// Print formatted string
	require.Greaterf(t, Add(1, 2), 2, "%v has to be greater than 2", Add(1, 2))
}

func TestAddOneGroup(t *testing.T) {
	// Run multiple grouped tests

	testCases := []struct {
		name     string
		x        int
		y        int
		expected int
	}{
		{name: "Zeros", x: 0, y: 0, expected: 0},
		{name: "Ones", x: 0, y: 1, expected: 1},
		{name: "Negative", x: 0, y: -1, expected: -1},
	}

	for _, tc := range testCases {
		actual := Add(tc.x, tc.y)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestAddOneSubTests(t *testing.T) {
	// Run multiple tests as subtests, advantages: Failed state show actual line number and name

	testCases := []struct {
		name     string
		x        int
		y        int
		expected int
	}{
		{name: "Zeros", x: 0, y: 0, expected: 0},
		{name: "Ones", x: 0, y: 1, expected: 1},
		{name: "Negative", x: 0, y: -1, expected: -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Add(tc.x, tc.y)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
