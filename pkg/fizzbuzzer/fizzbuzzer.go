package fizzbuzzer

// FizzBuzzer provides an interface for the FizzBuzz implementation
type FizzBuzzer interface {
	// Converts an in to a FizzBuzz string
	String(number int) string
}
