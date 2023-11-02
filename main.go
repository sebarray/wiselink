package main

import (
	"github.com/sebarray/wiselink/config"
	"github.com/sebarray/wiselink/handler"
)

func init() {
	config.Loadenv()
}

func main() {


	handler.Manager()
}
