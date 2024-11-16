package resp

import "errors"

func (m Message) String() (string, error) {
	if m.Typ == BULK_STRING {
		return writeBulkString(m), nil
	}

	return "", errors.New("unsupported message type to string")
}

func writeBulkString(m Message) string {
	out := m.Typ.String(len(m.StringVal))
	out += m.StringVal + "\r\n"
	return out
}
