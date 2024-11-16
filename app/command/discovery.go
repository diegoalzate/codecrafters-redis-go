package command

import (
	"errors"
	"strings"

	resp "github.com/codecrafters-io/redis-starter-go/app/resp"
)

type Commander interface {
	run(m resp.Message) (string, error)
}

var commands = make(map[string]Commander)

func init() {
	commands["echo"] = Echo{}
	commands["ping"] = Ping{}
}

func RunCommand(m resp.Message) (string, error) {
	firstArg := strings.ToLower(m.Values[0].StringVal)

	fn := commands[firstArg]

	if fn == nil {
		return "", errors.New("unsupported command")
	}

	return fn.run(m)
}
