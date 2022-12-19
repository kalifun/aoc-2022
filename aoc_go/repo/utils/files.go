package utils

import (
	"log"
	"os"
	"path"

	"github.com/kalifun/aco-2022/entity/consts"
)

// ReadFile
func NewFileReader(fileName string) (*os.File, error) {
	filePath := findFile(fileName)
	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("can't open %s \n", filePath)
		return nil, err
	}
	return f, nil
}

func findFile(fileName string) string {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Panicln(currentPath)
	}
	return path.Join(currentPath, consts.DataDir, fileName)
}
