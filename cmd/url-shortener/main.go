package main

import (
	"fmt"
	"url-shortener/cmd/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger

	// TODO: init storage: sqlite

	// TODO: init router

	// TODO: run server
}
