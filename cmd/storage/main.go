package main

import (
	"fmt"
	"log"
)
import "github.com/dima_osin/storage/internal/storage"

func main() {
	storage := storage.NewStorage()

	file, error := storage.Upload("test.txt", []byte("Hello, World!"))

	if error != nil {
		log.Fatal(error)
	}

	findedFile, error := storage.GetById(file.ID)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Storage:", storage)
	fmt.Println("\n")
	fmt.Println("Finded file:", findedFile.String())
}
