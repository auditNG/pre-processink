package transform

import (
	"errors"
	"strconv"
	"strings"
)

type TransformHelper struct{}

func NewTransformHelper() TransformHelper {
	return TransformHelper{}
}

func (t TransformHelper) GetStringValue(message string, key string) (string, error) {
	data := message
	start := 0
	end := 0
	if start = strings.Index(data, key); start < 0 {
		return "", errors.New("Error parsing")
	}

	// Progress the start point beyond the = sign
	start += len(key)
	if end = strings.IndexByte(data[start:], spaceChar); end < 0 {
		// There was no ending space, maybe the syscall id is at the end of the line
		end = len(data) - start
	}

	retval := data[start : start+end]
	return retval, nil
}

func (t TransformHelper) GetIntValue(message string, key string) (int, error) {
	val, err := t.GetStringValue(message, key)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(val)
}
