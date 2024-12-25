package services

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"os"
)

// MergePDFs merges multiple PDF files into one
func MergePDFs(inputFiles []string, outputFile string) error {
	// Check if the output file already exists
	if _, err := os.Stat(outputFile); err == nil {
		return fmt.Errorf("output file already exists")
	}

	// Create a default PDF configuration
	conf := model.NewDefaultConfiguration()

	// Merge PDF files
	// The third argument is 'true' to overwrite the output file if it exists
	return api.MergeCreateFile(inputFiles, outputFile, true, conf)
}
