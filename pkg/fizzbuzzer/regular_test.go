package fizzbuzzer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzzer(t *testing.T) {
	fizzbuzzer, err := NewFizzBuzzer()
	assert.Nil(t, err)

	assert.Equal(t, "1", fizzbuzzer.String(1))
	assert.Equal(t, "2", fizzbuzzer.String(2))
	assert.Equal(t, "fizz", fizzbuzzer.String(3))
	assert.Equal(t, "4", fizzbuzzer.String(4))
	assert.Equal(t, "buzz", fizzbuzzer.String(5))
	assert.Equal(t, "fizz", fizzbuzzer.String(6))
	assert.Equal(t, "7", fizzbuzzer.String(7))
	assert.Equal(t, "8", fizzbuzzer.String(8))
	assert.Equal(t, "fizz", fizzbuzzer.String(9))
	assert.Equal(t, "buzz", fizzbuzzer.String(10))
	assert.Equal(t, "11", fizzbuzzer.String(11))
	assert.Equal(t, "fizz", fizzbuzzer.String(12))
	assert.Equal(t, "13", fizzbuzzer.String(13))
	assert.Equal(t, "14", fizzbuzzer.String(14))
	assert.Equal(t, "fizzbuzz", fizzbuzzer.String(15))
}
