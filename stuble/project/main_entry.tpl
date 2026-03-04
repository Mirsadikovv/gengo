package main

import "{{ .ModPath }}/src"

func main() {
	var env src.Env
	src.Exec(&env)
}
