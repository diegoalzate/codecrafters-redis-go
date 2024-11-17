package command

import (
	"errors"
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/store"
)

func RunSet(store *store.Store, args []string) (string, error) {
	if len(args) != 2 {
		fmt.Println(fmt.Printf("expected 2 args, received: %v", args))
		return "", errors.New("expected 2 args")
	}

	key := args[0]
	val := args[1]

	store.Set(key, val)

	return resp.OK, nil
}
