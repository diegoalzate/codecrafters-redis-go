package command

import "github.com/codecrafters-io/redis-starter-go/app/store"

func RunPing(store *store.Store, args []string) (string, error) {
	return "+PONG\r\n", nil
}
