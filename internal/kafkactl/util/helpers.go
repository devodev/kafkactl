package util

import (
	"fmt"
	"strings"
)

func KeyValueParse(sep string, values []string) (map[string]string, error) {
	keyPairs := make(map[string]string, len(values))
	for _, value := range values {
		keyPair := strings.Split(value, sep)
		if len(keyPair) != 2 {
			return nil, fmt.Errorf("invalid format (key%svalue): %s", sep, value)
		}
		keyPairs[keyPair[0]] = keyPair[1]
	}
	return keyPairs, nil
}

func KeyValueDeleteParse(sep string, delChar string, values []string) (map[string]string, error) {
	keyPairs := make(map[string]string, len(values))
	for _, value := range values {
		keyPair := strings.Split(value, sep)
		keyPairLen := len(keyPair)

		if keyPairLen == 0 {
			continue
		}
		if keyPairLen > 2 {
			return nil, fmt.Errorf("invalid format (key%svalue): %s", sep, value)
		}

		key := keyPair[0]
		if key == "" {
			return nil, fmt.Errorf("invalid format: key empty")
		}

		if len(keyPair) == 1 {
			if !strings.HasSuffix(key, delChar) {
				return nil, fmt.Errorf("invalid format: %s", keyPair[0])
			}
			key = key[:len(key)-1]
		}

		value := ""
		if len(keyPair) == 2 {
			value = keyPair[1]
		}

		keyPairs[key] = value
	}
	return keyPairs, nil
}

type CLIError struct {
	Action       string
	ResourceType string
	Err          error
}

func MakeCLIError(action string, resourceType string, err error) CLIError {
	return CLIError{action, resourceType, err}
}

func (e CLIError) Error() string {
	return fmt.Sprintf("could not %s %s: %s", e.Action, e.ResourceType, e.Err)
}

type SerializationError struct{ error }

func MakeSerializationError(err error) SerializationError {
	return SerializationError{err}
}

func (e SerializationError) Error() string {
	return fmt.Sprintf("could not serialize: %s", e.error)
}
