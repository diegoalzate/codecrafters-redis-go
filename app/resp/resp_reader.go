package resp

import (
	"bufio"
	"errors"
	"io"
)

type Message struct {
	Typ       Flag
	Integer   int
	StringVal string
	Values    []Message
}

func RespRead(i *bufio.Reader) (Message, error) {
	var m Message

	// read flag
	flag, err := readFlag(i)

	if err != nil {
		return m, err
	}

	// parse message
	switch flag {
	case SIMPLE_STRINGS:
		return readSimpleString(i)
	case INTEGER:
		return readInteger(i)
	case BULK_STRING:
		return readBulkString(i)
	case ARRAY:
		return readArray(i)
	default:
		return m, errors.New("unsupported function in resp")
	}
}

func readFlag(i *bufio.Reader) (Flag, error) {
	typ, err := i.ReadByte()

	if err != nil {
		return "", err
	}

	return NewFlag(typ)
}

func readSimpleString(i *bufio.Reader) (Message, error) {
	// pre checks check
	bytes, err := i.ReadBytes('\n')

	if err != nil {
		return Message{}, err
	}

	if trim := len(bytes) - 2; trim < 0 {
		return Message{}, errors.New("unexpected end of file")
	} else {
		bytes = bytes[:trim]
	}

	return Message{
		Typ:       SIMPLE_STRINGS,
		StringVal: string(bytes),
	}, nil
}

func readInteger(i *bufio.Reader) (Message, error) {
	val, err := readI(i)

	if err != nil {
		return Message{}, err
	}

	return Message{
		Typ:     INTEGER,
		Integer: val,
	}, nil
}

func readBulkString(i *bufio.Reader) (Message, error) {
	length, err := readI(i)

	if err != nil {
		return Message{}, err
	}

	out := make([]byte, length)

	// read bulk string by reading the length
	if _, err := io.ReadFull(i, out); err != nil {
		return Message{}, nil
	}

	if _, err = i.Discard(2); err != nil {
		return Message{}, nil
	}

	return Message{
		Typ:       BULK_STRING,
		StringVal: string(out),
	}, nil
}

func readI(i *bufio.Reader) (int, error) {
	bytes, err := i.ReadBytes('\n')

	if err != nil {
		return 0, err
	}

	if len(bytes) < 3 {
		return 0, errors.New("unexpected end of file")
	}

	var out int

	for _, b := range bytes[:len(bytes)-2] {
		byteVal := int(b - '0')

		if byteVal < 0 || byteVal > 9 {
			return 0, errors.New("unexpected end of file")
		}

		out = 10*out + byteVal
	}

	return out, nil
}

func readArray(i *bufio.Reader) (Message, error) {
	// read counter and then read the next command one by on
	counter, err := readI(i)

	if err != nil {
		return Message{}, err
	}

	out := Message{
		Typ:    ARRAY,
		Values: []Message{},
	}

	for range counter {
		message, err := RespRead(i)

		if err != nil {
			return out, err
		}

		out.Values = append(out.Values, message)
	}
	return out, nil
}
