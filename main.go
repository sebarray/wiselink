package main

import (
	"fmt"

	"github.com/sebarray/wiselink/config"
	"github.com/sebarray/wiselink/handler"
)

func init() {
	config.Loadenv()
}

func main() {
	fmt.Println("Starting Wiselink")

	handler.Manager()
}
