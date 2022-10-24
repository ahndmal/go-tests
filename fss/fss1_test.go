package fss

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFilesDocs(t *testing.T) {
	ioutil.WriteFile("2.txt", []byte("TO ADD THIS LINE"), fs.ModeAppend)
	err := filepath.WalkDir("/home/malandr/Documents", func(path string, dir fs.DirEntry, err error) error {
		fmt.Printf("Path is %s", path)
		println(dir)
		return err
	})
	if err != nil {
		return
	}

	file, err := os.Create("1.txt")
	if err != nil {
		return
	}
	file.WriteString("Added")
	file.Write([]byte("Second"))
	file.WriteAt([]byte("NEW"), 5)
}

func TestReadFs(t *testing.T) {
	//read file
	//fs.ReadFile(fs.SubFS(), "1.txt")
	docsDir, err := os.ReadDir("/home/malandr/docs")
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
