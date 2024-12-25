package main

import (
	"log"
	"pdf_tool/handlers"
)

func main() {
	log.Println("Starting PDF Tool...")
	handlers.StartCLI()
}