package handlers

import (
	"fmt"
	"log"
	"strings"
	"pdf_tool/services"
)

// StartCLI starts the command-line interface
func StartCLI() {
	for {
		fmt.Println("\nPDF Tool")
		fmt.Println("1. Merge PDFs")
		fmt.Println("2. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Merge PDFs
			var inputFiles string
			var outputFile string

			fmt.Print("Enter input PDF files (comma-separated): ")
			fmt.Scanln(&inputFiles)

			fmt.Print("Enter output file name: ")
			fmt.Scanln(&outputFile)

			files := strings.Split(inputFiles, ",")
			err := services.MergePDFs(files, outputFile)
			if err != nil {
				log.Printf("Error merging PDFs: %v\n", err)
			} else {
				fmt.Println("PDFs merged successfully!")
			}

		case 2:
			// Exit
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
