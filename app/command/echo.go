package command

import (
	"strings"

	resp "github.com/codecrafters-io/redis-starter-go/app/resp"
)

type Echo struct{}

func (e Echo) run(m resp.Message) (string, error) {
	secondArg := strings.ToLower(m.Values[1].StringVal)

	echoMessage := resp.Message{
		Typ:       resp.BULK_STRING,
		StringVal: secondArg,
	}

	return echoMessage.String()
}
