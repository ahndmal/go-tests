package fss

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

func TestReadDocs(t *testing.T) {
	docsDir, err := os.ReadDir("/home/andrii/docs")
	if err != nil {
		return
	}
	println(len(docsDir)) // files nums

	start := time.Now()
	var wg sync.WaitGroup

	for i := 2; i < 40; i++ {
		wg.Add(1)
		go func(num int) {
			fileBts, err := os.ReadFile(fmt.Sprintf("/home/andrii/docs/%s", docsDir[num].Name()))
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
