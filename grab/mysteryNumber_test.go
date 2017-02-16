package grab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Given the unit test, please write a simple application to make the test pass.
// This is a game call mystery number. Whoever guess and get the mystery number is the loser.
// The program will take in the guessing number and return certain status.
// Status could be the reduced numbers range, invalid, you are the loser, the next player is the loser
// Refer to the unit test for more info

func TestPlayMysteryNumber(t *testing.T) {
	mysteryNumber := 50
	unit := MysteryNumber()
	unit.setMysteryNumber(mysteryNumber)

	assert.Equal(t, "Invalid", unit.playMysteryNumber(1))
	assert.Equal(t, "2 to 100", unit.playMysteryNumber(2))
	assert.Equal(t, "10 to 100", unit.playMysteryNumber(10))
	assert.Equal(t, "40 to 100", unit.playMysteryNumber(40))
	assert.Equal(t, "40 to 80", unit.playMysteryNumber(80))
	assert.Equal(t, "40 to 56", unit.playMysteryNumber(56))
	assert.Equal(t, "Invalid", unit.playMysteryNumber(90))
	assert.Equal(t, "Invalid", unit.playMysteryNumber(20))
	assert.Equal(t, "40 to 51", unit.playMysteryNumber(51))
	assert.Equal(t, "49 to 51", unit.playMysteryNumber(49))
	assert.Equal(t, "You are the loser", unit.playMysteryNumber(50))

	mysteryNumber = 1
	unit = MysteryNumber()
	unit.setMysteryNumber(mysteryNumber)

	assert.Equal(t, "Invalid", unit.playMysteryNumber(101))
	assert.Equal(t, "Invalid", unit.playMysteryNumber(100))
	assert.Equal(t, "1 to 99", unit.playMysteryNumber(99))
	assert.Equal(t, "The next player is the loser", unit.playMysteryNumber(2))

	mysteryNumber = 50
	unit = MysteryNumber()
	unit.setMysteryNumber(mysteryNumber)

	assert.Equal(t, "You are the loser", unit.playMysteryNumber(50))
}
