package RESP

import (
	"bytes"
	"io"
	"strings"
)

var db = make(map[string]string)

func DeserializeCommand(reader io.Reader) string {
	buf := make([]byte, 1024)
	n, err := reader.Read(buf)
	if err != nil {
		return "Client disconnected"
	}
	idx := bytes.Index(buf[:n], []byte("\r\n"))
	if idx == -1 {
		return "invalid command"
	}
	parts := bytes.Split(buf, []byte("\r\n"))
	switch strings.ToUpper(string(parts[2])) {
	case "SET":
		db[string(parts[4])] = string(parts[6])
		return "+Ok"
	case "GET":
		val, ok := db[string(parts[4])]
		if !ok {
			return "No such key"
		}
		return val
	case "DEL":
		delete(db, string(parts[4]))
		return "-Ok"
	}
	return "invalid command"
}
