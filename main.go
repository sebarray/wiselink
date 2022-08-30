package main

import (
	"github.com/sebarray/wiselink/config"
	"github.com/sebarray/wiselink/handler"
)

func main() {
	config.Loadenv()
	handler.Manager()
}
