package fss

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestLogsRead(t *testing.T) {
	start := time.Now()
	root := "/home/andrii/_data/"
	err := filepath.WalkDir(root, func(path string, dir fs.DirEntry, err error) error {
		if dir.Name() != "_data" {
			fmt.Printf(">>> parsing file: %s \n", dir.Name())
			data, err := os.ReadFile(root + dir.Name())

			lines := strings.Split(string(data), "\n")
			words := strings.Split(string(data), " ")

			vowels := [6]string{"a", "o", "u", "e", "y", "i"}
			var voweled int64
			for _, word := range words {
				for _, vowel := range vowels {
					if len(word) > 0 {
						if string(word[0]) == vowel {
							voweled += 1
						}
					}
				}
			}

			fmt.Printf(">> Lines: %d \n", len(lines))
			fmt.Printf(">> Words: %d \n", len(words))
			fmt.Printf(">> Voweled: %d \n", voweled)

			if err != nil {
				fmt.Printf("Error when reading file from dir: %s", err)
			}
			return nil
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error when reading dir %v", err)
	}
	end := time.Now()
	fmt.Printf(">>>>> END: %d", end.Sub(start).Milliseconds())
}
