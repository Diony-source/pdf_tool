package services

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"os"
)

// SplitPDF splits a PDF into separate pages
func SplitPDF(inputFile, outputDir string) error {
	// Check if the input file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist")
	}

	// Create a default PDF configuration
	conf := model.NewDefaultConfiguration()

	// Split PDF file into pages
	// The third argument specifies the number of pages to process at a time (use -1 for all pages)
	return api.SplitFile(inputFile, outputDir, -1, conf)
}
