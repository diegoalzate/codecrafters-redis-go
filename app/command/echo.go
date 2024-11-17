package command

import (
	resp "github.com/codecrafters-io/redis-starter-go/app/resp"
)

func RunEcho(args []string) (string, error) {
	msg := ""

	for _, word := range args {
		msg += word
	}

	echoMessage := resp.Message{
		Typ:       resp.BULK_STRING,
		StringVal: msg,
	}

	return echoMessage.String()
}
