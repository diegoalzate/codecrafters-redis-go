package command

import (
	"errors"
	"fmt"
	"strings"

	resp "github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/store"
)

type Command struct {
	funcName string
	fn       func(store *store.Store, args []string) (string, error)
}

type CommandHandler struct {
	store    *store.Store
	commands map[string]Command
}

func NewCommandHandler(store *store.Store) CommandHandler {
	ch := CommandHandler{
		commands: make(map[string]Command),
		store:    store,
	}

	ch.registerCommands()

	return ch
}

func (ch *CommandHandler) registerCommands() {
	ch.commands["echo"] = Command{
		funcName: "ECHO",
		fn:       RunEcho,
	}
	ch.commands["ping"] = Command{
		funcName: "PING",
		fn:       RunPing,
	}
	ch.commands["get"] = Command{
		funcName: "GET",
		fn:       RunGet,
	}
	ch.commands["set"] = Command{
		funcName: "SET",
		fn:       RunSet,
	}
}

func (ch *CommandHandler) findCommand(m resp.Message) (Command, error) {
	if m.Typ != resp.ARRAY {
		return Command{}, fmt.Errorf("unsupported resp typ: %v", string(m.Typ))
	}

	firstArg := strings.ToLower(m.Values[0].StringVal)

	if firstArg == "" {
		return Command{}, errors.New("no first arg")
	}

	foundCommand, exists := ch.commands[firstArg]

	if !exists {
		return Command{}, errors.New("unsupported command")
	}

	return foundCommand, nil
}

func (ch *CommandHandler) RunCommand(m resp.Message) (string, error) {
	cmd, err := ch.findCommand(m)

	if err != nil {
		return "", err
	}

	args := []string{}

	for _, arg := range m.Values[1:] {
		args = append(args, arg.StringVal)
	}

	return cmd.fn(ch.store, args)
}
