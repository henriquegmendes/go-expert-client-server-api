package helpers

import (
	"errors"
	"fmt"
	"strings"
)

const ContextDeadlineExceededErrorMessage = "context deadline exceeded"

func CheckContextTimedOutError(err error) bool {
	if err == nil {
		return false
	}

	return strings.Contains(err.Error(), ContextDeadlineExceededErrorMessage)
}

func CheckClientWithErrorStatusCode(statusCode int) bool {
	return statusCode >= 400 && statusCode <= 599
}

func ParseClientErrorResponse(statusCode int, body []byte) error {
	return errors.New(fmt.Sprintf("client error - status: %v - body - '%s'", statusCode, string(body)))
}
