package generators

import (
	"fmt"
	"os"
	"strings"
)

// (fileName, separator)

type fileGenerator struct {
	data []string
	pos  int
}

func (fg *fileGenerator) Next() bool {
	if fg.pos >= len(fg.data) {
		return false
	}
	fg.pos++
	return true
}

func (fg *fileGenerator) Val() Element {
	return fg.data[fg.pos-1]
}

func NewFileGenerator(fileName string, separator string) (*fileGenerator, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Printf("fileGenerator reading error %v", err)
		return &fileGenerator{}, err
	}

	parts := strings.Split(string(data), separator)

	return &fileGenerator{
		data: parts,
		pos:  0,
	}, nil
}
