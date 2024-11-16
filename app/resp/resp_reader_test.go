package resp

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected Message
		wantErr  bool
	}{
		{
			name:  "simple string",
			input: "+hello\r\n",
			expected: Message{
				Typ:       SIMPLE_STRINGS,
				StringVal: "hello",
			},
		},
		{
			name:  "integer",
			input: ":123\r\n",
			expected: Message{
				Typ:     INTEGER,
				Integer: 123,
			},
		},
		{
			name:  "bulk string",
			input: "$5\r\nhello\r\n",
			expected: Message{
				Typ:       BULK_STRING,
				StringVal: "hello",
			},
		},
		{
			name:  "array",
			input: "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n",
			expected: Message{
				Typ: ARRAY,
				Values: []Message{
					{
						Typ:       BULK_STRING,
						StringVal: "hello",
					},
					{
						Typ:       BULK_STRING,
						StringVal: "world",
					},
				},
			},
		},
		{
			name:    "invalid flag",
			input:   "invalid\r\n",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			p := bufio.NewReader(strings.NewReader(tc.input))

			got, err := RespRead(p)

			if (err != nil) != tc.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("Parse() = %v, want %v", got, tc.expected)
			}
		})
	}
}
