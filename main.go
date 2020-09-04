package main

import (
	"gitlab.com/tsuchinaga/kabus-legs/infra/scheduler"
	"gitlab.com/tsuchinaga/kabus-legs/ui/cli"
)

func main() {
	go scheduler.Run()
	if err := cli.Run(); err != nil {
		panic(err)
	}
}
