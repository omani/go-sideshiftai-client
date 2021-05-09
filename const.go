package sideshiftai

import (
	"fmt"
)

// APIBaseAddress is the base URL of the sideshift.ai API
const APIBaseAddress = "https://sideshift.ai/api/"

// APIVersion is the API version identifier (currently v3)
const APIVersion = "v1"

// ErrorCode is a sideshift.ai error code table
type ErrorCode int

// APIError represents an error message by the sideshift.ai API
type APIError struct {
	Err struct {
		Message string `json:"message"`
	} `json:"error"`
}

func (we *APIError) Error() string {
	return fmt.Sprintf("API Error: (%v) - See https://documenter.getpostman.com/view/6895769/TWDZGvjd for help!", we.Err.Message)
}
