package main

import (
	"fmt"
	"log"

	"github.com/chengaro/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("It works"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetById(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("It restored", restoredFile)
}
