package services

import (
	"fmt"
	"pdf_tool/utils"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

// MergePDFs merges multiple PDF files into one
func MergePDFs(inputFiles []string, outputFile string) error {
	var fullPaths []string

	// Find the full paths of the input files
	for _, file := range inputFiles {
		fullPath, err := utils.FindFileInSystem("C:\\", file) // Search from C:\ (root directory)
		if err != nil {
			return fmt.Errorf("file %s not found: %v", file, err)
		}
		fullPaths = append(fullPaths, fullPath)
	}

	// Create a default PDF configuration
	conf := model.NewDefaultConfiguration()

	// Merge PDF files
	return api.MergeCreateFile(fullPaths, outputFile, true, conf)
}
