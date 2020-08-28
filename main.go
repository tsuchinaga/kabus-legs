package main

import "gitlab.com/tsuchinaga/kabus-legs/ui/cli"

func main() {
	if err := cli.Run(); err != nil {
		panic(err)
	}
}
