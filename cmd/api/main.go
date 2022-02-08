package main

import (
	"fmt"

	"github.com/mauromamani/go-clean-architecture/internal/server"
)

func main() {
	fmt.Println("Run")

	s := server.NewServer()

	s.Run()
}
