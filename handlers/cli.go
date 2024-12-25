package handlers

import (
	"fmt"
	"log"
	"pdf_tool/services"
)

// StartCLI starts the command-line interface
func StartCLI() {
	for {
		fmt.Println("\nPDF Tool")
		fmt.Println("1. Merge PDFs")
		fmt.Println("2. Split PDF")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Merge PDFs
			var inputFiles []string
			var outputFile string
			fmt.Print("Enter input PDF files (comma-separated): ")
			var input string
			fmt.Scanln(&input)
			inputFiles = append(inputFiles, input)

			fmt.Print("Enter output file name: ")
			fmt.Scanln(&outputFile)

			err := services.MergePDFs(inputFiles, outputFile)
			if err != nil {
				log.Printf("Error merging PDFs: %v\n", err)
			} else {
				fmt.Println("PDFs merged successfully!")
			}

		case 2:
			// Split PDF
			var inputFile string
			var outputDir string
			fmt.Print("Enter input PDF file: ")
			fmt.Scanln(&inputFile)

			fmt.Print("Enter output directory: ")
			fmt.Scanln(&outputDir)

			err := services.SplitPDF(inputFile, outputDir)
			if err != nil {
				log.Printf("Error splitting PDF: %v\n", err)
			} else {
				fmt.Println("PDF split successfully!")
			}

		case 3:
			// Exit
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
