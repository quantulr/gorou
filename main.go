package main

import (
	"github.com/quantulr/gorou/core"
)

func main() {
	r := core.CreateRoutes()
	r.Run(":3000")
}
