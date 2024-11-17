package command

func RunPing(args []string) (string, error) {
	return "+PONG\r\n", nil
}
