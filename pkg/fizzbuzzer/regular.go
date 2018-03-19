package fizzbuzzer

// Regular FizzBuzz implementation
func NewFizzBuzzer() (FizzBuzzer, error) {
	return NewAdvancedFizzBuzzer(Match{Number: 3, Term: "fizz"}, Match{Number: 5, Term: "buzz"})
}
