package client

import "fmt"

type ClientError struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func (c ClientError) Error() string {
	return fmt.Sprintf("[%d] %s", c.ErrorCode, c.Message)
}
