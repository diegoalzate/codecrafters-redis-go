package command

import (
	"errors"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/store"
)

func RunGet(store *store.Store, args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New("GET: expected 1 arg")
	}

	key := args[0]

	val, exists := store.Get(key)

	if !exists {
		return resp.NULL, nil
	}

	response := resp.Message{
		Typ:       resp.BULK_STRING,
		StringVal: val,
	}

	return response.String()
}
