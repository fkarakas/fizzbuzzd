package fizzbuzzer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdvancedFizzBuzzer(t *testing.T) {
	fizzbuzzer, err := NewAdvancedFizzBuzzer(Match{Number: 2, Term: "bob"}, Match{Number: 3, Term: "dylan"})
	assert.Nil(t, err)

	assert.Equal(t, "1", fizzbuzzer.String(1))
	assert.Equal(t, "bob", fizzbuzzer.String(2))
	assert.Equal(t, "dylan", fizzbuzzer.String(3))
	assert.Equal(t, "bob", fizzbuzzer.String(4))
	assert.Equal(t, "5", fizzbuzzer.String(5))
	assert.Equal(t, "bobdylan", fizzbuzzer.String(6))
	assert.Equal(t, "7", fizzbuzzer.String(7))
	assert.Equal(t, "bob", fizzbuzzer.String(8))
	assert.Equal(t, "dylan", fizzbuzzer.String(9))
	assert.Equal(t, "bob", fizzbuzzer.String(10))
	assert.Equal(t, "11", fizzbuzzer.String(11))
	assert.Equal(t, "bobdylan", fizzbuzzer.String(12))
	assert.Equal(t, "13", fizzbuzzer.String(13))
	assert.Equal(t, "bob", fizzbuzzer.String(14))
	assert.Equal(t, "dylan", fizzbuzzer.String(15))
}

func TestAdvancedFizzBuzzerBadParameter(t *testing.T) {
	fizzbuzzer, err := NewAdvancedFizzBuzzer(Match{Number: 0, Term: "bob"}, Match{Number: 3, Term: "dylan"})
	assert.NotNil(t, err)
	assert.Nil(t, fizzbuzzer)
}
