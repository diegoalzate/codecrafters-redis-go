package command

import (
	resp "github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/store"
)

func RunEcho(store *store.Store, args []string) (string, error) {
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
