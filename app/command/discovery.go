package command

import (
	"errors"
	"fmt"
	"strings"

	resp "github.com/codecrafters-io/redis-starter-go/app/resp"
)

type Command struct {
	funcName string
	fn       func(args []string) (string, error)
}

var commands = make(map[string]Command)

func init() {
	commands["echo"] = Command{
		funcName: "ECHO",
		fn:       RunEcho,
	}
	commands["ping"] = Command{
		funcName: "PING",
		fn:       RunPing,
	}
}

func findCommand(m resp.Message) (Command, error) {
	if m.Typ != resp.ARRAY {
		return Command{}, fmt.Errorf("unsupported resp typ: %v", string(m.Typ))
	}

	firstArg := strings.ToLower(m.Values[0].StringVal)

	if firstArg == "" {
		return Command{}, errors.New("no first arg")
	}

	foundCommand, exists := commands[firstArg]

	if !exists {
		return Command{}, errors.New("unsupported command")
	}

	return foundCommand, nil
}

func RunCommand(m resp.Message) (string, error) {
	cmd, err := findCommand(m)

	if err != nil {
		return "", err
	}

	args := []string{}

	for _, arg := range m.Values {
		args = append(args, arg.StringVal)
	}

	return cmd.fn(args)

}
