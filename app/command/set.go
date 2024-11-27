package command

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/store"
)

func RunSet(store *store.Store, args []string) (string, error) {
	if len(args) < 2 {
		fmt.Println(fmt.Printf("expected at least 2 args, received: %v", args))
		return "", errors.New("expected at least 2 args")
	}

	key := args[0]
	val := args[1]

	if len(args) > 2 {
		px := strings.ToLower(args[2])

		if px == "px" {
			ttlStr := args[3]
			ttlVal, err := strconv.Atoi(ttlStr)

			if err != nil {
				return "", err
			}

			ttl := time.Duration(ttlVal) * time.Millisecond

			store.SetWithExpiry(key, val, ttl)

			return resp.OK, nil
		}

		return "", errors.New("failed to parse expiry")
	}

	store.Set(key, val)

	return resp.OK, nil
}
