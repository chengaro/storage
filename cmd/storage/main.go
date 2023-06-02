package main

import (
	"fmt"
	"github.com/chengaro/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	fmt.Println("It works", st)
}
