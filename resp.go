package main

var Redis = make(map[string]string)

func ReadSet(parts []string) string {
	Redis[parts[4]] = parts[6]
	return "ok"
}

func ReadGet(key string) string {
	return Redis[key]
}

func ReadDel() {

}
