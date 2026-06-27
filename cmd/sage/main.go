package main

import (
	"github.com/allenbiji/clone-sage/internal/cli"
	_ "github.com/allenbiji/clone-sage/internal/checks" // register all check types via init()
)

func main() {
	cli.Execute()
}

