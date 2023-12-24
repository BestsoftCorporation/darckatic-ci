package main

import (
	"darkatic-ci/cmd"

	"darkatic-ci/internal/api"
)

func main() {

	api.StartServer()
	cmd.Execute()
}
