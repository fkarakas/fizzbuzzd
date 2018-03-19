package fizzbuzzer

import (
	"errors"
	"fmt"
)

// Used by advanced FizzBuzz implementation for input tuples
type Match struct {
	Number int
	Term   string
}

// Contructor with a variadic list of Match parameter
func NewAdvancedFizzBuzzer(matchs ...Match) (FizzBuzzer, error) {
	return newAdvancedFizzBuzzer(matchs)
}

func newAdvancedFizzBuzzer(matchs []Match) (*advancedFizzBuzzer, error) {

	for _, match := range matchs {
		if match.Number == 0 {
			return nil, errors.New("Illegal parameter")
		}
	}

	return &advancedFizzBuzzer{
		matchs: matchs,
	}, nil
}

// Internal struct holding list of Matchs
type advancedFizzBuzzer struct {
	matchs []Match
}

func (a *advancedFizzBuzzer) String(number int) string {
	result := ""

	for _, match := range a.matchs {
		if number%match.Number == 0 {
			result += match.Term
		}
	}

	if result == "" {
		result = fmt.Sprint(number)
	}

	return result
}
