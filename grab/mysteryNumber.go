package grab

import (
	"fmt"
)

type MysteryNumberTest struct {
	mysteryNumber      int
	highestGuessNumber int
	lowestGuessNumber  int
}

func MysteryNumber() *MysteryNumberTest {
	return &MysteryNumberTest{
		highestGuessNumber: 100,
		lowestGuessNumber:  1,
	}
}

func (m *MysteryNumberTest) setMysteryNumber(mysteryNumber int) {
	m.mysteryNumber = mysteryNumber
}

func (m *MysteryNumberTest) playMysteryNumber(userGuess int) string {
	fmt.Printf("userGusess %v mysteryNumber %v lowestGuessNumber %v highestGuessNumber %v \n",
		userGuess, m.mysteryNumber, m.lowestGuessNumber, m.highestGuessNumber)

	if userGuess == m.mysteryNumber {
		return fmt.Sprintf("%v", "You are the loser")
	}

	if userGuess <= m.lowestGuessNumber || userGuess >= m.highestGuessNumber {
		return fmt.Sprintf("%v", "Invalid")
	}

	if userGuess-m.lowestGuessNumber == m.mysteryNumber {
		return fmt.Sprintf("%v", "The next player is the loser")
	}

	if userGuess > m.highestGuessNumber/2 {
		if userGuess > m.mysteryNumber {
			m.highestGuessNumber = userGuess
		} else {
			m.lowestGuessNumber = userGuess
		}
	} else {
		m.lowestGuessNumber = userGuess
	}
	return fmt.Sprintf("%v to %v", m.lowestGuessNumber, m.highestGuessNumber)
}
