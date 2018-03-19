package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fkarakas/fizzbuzzd/pkg/fizzbuzzer"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetArrayParameters is used by GetArray route for streaming json array of fizzbuzz strings
type GetArrayParameters struct {
	number1 int
	number2 int
	term1   string
	term2   string
	limit   int
}

// parseGetArrayParameters retrieves parameters for GetArray route and makes some checks
func parseGetArrayParameters(c *gin.Context) (*GetArrayParameters, error) {
	result := GetArrayParameters{}

	number1, err := strconv.Atoi(c.Param("number1"))
	if err != nil {
		return nil, err
	}

	number2, err := strconv.Atoi(c.Param("number2"))
	if err != nil {
		return nil, err
	}

	if number1 == 0 || number2 == 0 {
		return nil, errors.New("Zero value not allowed")
	}

	result.number1 = number1
	result.number2 = number2

	result.term1 = c.Param("term1")
	result.term2 = c.Param("term2")

	result.limit = 100

	limit := c.Query("limit")
	if limit != "" {
		num, err := strconv.Atoi(limit)
		if err != nil {
			return nil, err
		}
		result.limit = num
	}

	return &result, nil
}

func badRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Bad request : %v", err)})
}

// Write to the server output
func write(c *gin.Context, s string) {
	if _, err := c.Writer.Write([]byte(s)); err != nil {
		panic(err)
	}
}

// GetArray streams a json array of generated strings
func GetArray(c *gin.Context) {
	params, err := parseGetArrayParameters(c)
	if err != nil {
		badRequest(c, err)
		return
	}

	fizzbuzzer, err := fizzbuzzer.NewAdvancedFizzBuzzer(fizzbuzzer.Match{Number: params.number1, Term: params.term1},
		fizzbuzzer.Match{Number: params.number2, Term: params.term2})
	if err != nil {
		badRequest(c, err)
		return
	}

	e := json.NewEncoder(c.Writer)

	// If you really want to stream json
	write(c, "{ \"result\": [")

	for i := 1; i <= params.limit; i++ {
		e.Encode(fizzbuzzer.String(i))

		if i < params.limit {
			write(c, ",")
		}
	}

	write(c, "]}")

	return
}
