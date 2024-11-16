package resp

import (
	"errors"
	"strconv"
)

type Flag byte

// REQUEST THINGS
const (
	SIMPLE_STRINGS Flag = '+'
	SIMPLE_ERRORS  Flag = '-'
	INTEGER        Flag = ':'
	BULK_STRING    Flag = '$'
	ARRAY          Flag = '*'
)

func NewFlag(input byte) (Flag, error) {
	sym := Flag(input)

	switch sym {
	case SIMPLE_STRINGS, SIMPLE_ERRORS, INTEGER, BULK_STRING, ARRAY:
		return sym, nil
	default:
		return 0, errors.New("unsupported symbol")
	}
}

func (f Flag) SimpleString() (string, error) {
	if f == ARRAY || f == BULK_STRING {
		return "", errors.New("use String() for this flag")
	}

	out := string(f) + "\r\n"

	return out, nil
}

func (f Flag) String(count int) string {
	countCh := strconv.Itoa(count)
	out := string(f) + countCh + "\r\n"
	return out
}
