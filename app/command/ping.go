package command

import (
	resp "github.com/codecrafters-io/redis-starter-go/app/resp"
)

type Ping struct{}

func (p Ping) run(m resp.Message) (string, error) {
	return "+PONG\r\n", nil
}
