package RESP

import (
	"errors"
	"strconv"
	"strings"
)

var Sep = "\r\n"

func SerializeCommand(command string) ([]byte, error) {
	SerializedCommand := ""
	command = strings.TrimSpace(command)
	args := strings.Split(command, " ")
	if len(args) > 3 {
		return []byte(""), errors.New("too many arguments")
	}
	for _, arg := range args {
		SerializedCommand += Sep + "$" + strconv.Itoa(len(arg)) + Sep + arg
	}
	return []byte("*" + strconv.Itoa(len(args)) + SerializedCommand + Sep), nil
}
