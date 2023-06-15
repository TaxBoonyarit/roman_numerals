package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanFunction(t *testing.T) {
	t.Run("Error input null", func(t *testing.T) {
		result, err := romanNumerals("")
		assert.Error(t, err, "invalid value character min : 1 and max : 15")
		assert.Equal(t, result, 0)
	})

	t.Run("Error not is roman", func(t *testing.T) {
		result, err := romanNumerals("A")
		assert.EqualError(t, err, "invalid value : A is not roman number")
		assert.Equal(t, result, 0)
	})

	t.Run("Example input = III ", func(t *testing.T) {
		result, err := romanNumerals("III")
		assert.NoError(t, err)
		assert.Equal(t, result, 3)
	})

	t.Run("Example input = LVIII ", func(t *testing.T) {
		result, err := romanNumerals("LVIII")
		assert.NoError(t, err)
		assert.Equal(t, result, 58)
	})

	t.Run("Example input = MCMXCIV ", func(t *testing.T) {
		result, err := romanNumerals("MCMXCIV")
		assert.NoError(t, err)
		assert.Equal(t, result, 1994)
	})
}
