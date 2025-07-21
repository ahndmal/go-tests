package fss

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

const DirName = "/home/andrii/Documents/"

func TestReadDocs(t *testing.T) {
	docsDir, err := os.ReadDir(DirName)
	if err != nil {
		return
	}
	println(len(docsDir)) // files amount

	start := time.Now()
	var wg sync.WaitGroup

	for i := 1; i <= 400; i++ {
		wg.Add(1)
		go func(num int) {
			fileBts, err := os.ReadFile(fmt.Sprintf("%s%s", DirName, docsDir[num].Name()))
			if err != nil {
				fmt.Println(err)
			}
			content := string(fileBts)
			fmt.Println(content)
		}(i)
	}
	wg.Wait()

	//for _, entry := range docsDir {
	//	if !entry.IsDir() {
	//		fileBts, err := os.ReadFile(fmt.Sprintf("/home/andrii/docs/%s", entry.Name()))
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		content := string(fileBts)
	//		fmt.Println(content)
	//
	//	}
	//}
	fmt.Println(time.Now().Sub(start))
}
