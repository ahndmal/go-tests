package fss

import (
	"fmt"
	"os"
	"testing"
)

func TestFilesDocs(t *testing.T) {
	file, err := os.Create("1.txt")
	if err != nil {
		return
	}
	file.WriteString("Added")
}

func TestReadFs(t *testing.T) {
	//read file
	//fs.ReadFile(fs.SubFS(), "1.txt")
	docsDir, err := os.ReadDir("/home/andrii/docs")
	if err != nil {
		return
	}
	println(len(docsDir)) // files nums

	fileBts, err := os.ReadFile(fmt.Sprintf("/home/andrii/docs/%s", docsDir[1].Name()))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(fileBts))

	for i, item := range docsDir {
		println("item {}", i)
		fmt.Println(item.Name())

	}
}
