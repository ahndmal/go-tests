package fss

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestFilesDocs(t *testing.T) {
	err := os.WriteFile("2.txt", []byte("TO ADD THIS LINE"), fs.ModeAppend)
	if err != nil {
		t.Fatal(err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %s\n", err)
	}
	err = filepath.WalkDir(fmt.Sprintf("%s/Documents", homeDir), func(path string, dir fs.DirEntry, err error) error {
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
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %s\n", err)
	}
	docsDir, err := os.ReadDir(fmt.Sprintf("%s/Documents", homeDir))
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
