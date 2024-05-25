package main

import (
	"encoding/json"
	"fmt"
	"go-tests/fss"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	// start
	UrlPingerAsync2()
}

func GitLanguages() {
	repos := fss.ReadJsonGT()
	jsonLangs := ""
	//var commonData map[string]int
	javaLines := 0
	goLines := 0
	pythonLines := 0
	cLines := 0
	htmlLines := 0
	jsLines := 0
	groovyLines := 0
	kotlinLines := 0

	log.Println("[] \033[34m Starting ...\033[32m")

	for _, repo := range repos {
		url := fmt.Sprintf("https://api.github.com/repos/AndriiMaliuta/%s/languages", repo.Name)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Printf("Error creating request for url: %v", err)
		}

		req.Header.Add("Authorization", "Bearer "+os.Getenv("GIT_TOKEN"))

		resp, err2 := client().Do(req)
		if err2 != nil {
			return
		}
		langsByte, err3 := ioutil.ReadAll(resp.Body)
		if err3 != nil {
			return
		}
		var data map[string]int
		err4 := json.Unmarshal(langsByte, &data)
		if err4 != nil {
			log.Panicln(err4)
		}

		for a := 0; a < len(data); a++ {
			if _, found := data["Go"]; found {
				goLines += data["Go"]
			} else if _, found := data["Java"]; found {
				javaLines += data["Java"]
			} else if _, found := data["HTML"]; found {
				htmlLines += data["HTML"]
			} else if _, found := data["Python"]; found {
				pythonLines += data["Python"]
			} else if _, found := data["C"]; found {
				cLines += data["C"]
			} else if _, found := data["JavaScript"]; found {
				jsLines += data["JavaScript"]
			} else if _, found := data["Groovy"]; found {
				groovyLines += data["Groovy"]
			} else if _, found := data["Kotlin"]; found {
				kotlinLines += data["Kotlin"]
			} else if _, found := data["HTML"]; found {
				htmlLines += data["HTML"]
			}
			log.Println(data)
		}

		jsonLangs += string(langsByte) + ","
		//log.Println(data)
		/*
			{
			  "Go": 3842,
			  "Makefile": 183
			}
		*/
		//var mapData map[string]int
		//mapData, ok := data.(map[string]int)
		//if !ok {
		//	log.Println("error when mapping data to map")
		//}
	}
	log.Printf("GO: %d", goLines)
	log.Printf("Java: %d", javaLines)
	log.Printf("C: %d", cLines)
	log.Printf("Pythond: %d", pythonLines)
	log.Printf("Groovy: %d", groovyLines)
	log.Printf("JS: %d", jsLines)
	log.Printf("Kotlin: %d", kotlinLines)
	log.Printf("HTML: %d", htmlLines)
}

func UrlPingerAsync2() {
	urls := []string{
		"https://us-central1-andmal-bot.cloudfunctions.net/gcp-java-perf-test",        // java
		"https://us-central1-silver-adapter-307718.cloudfunctions.net/go-test-parser", // Go
		"https://us-west2-silver-adapter-307718.cloudfunctions.net/node-perf-db",      // node
		"https://us-central1-andmal-bot.cloudfunctions.net/python-perf-test",          // python
		"https://us-central1-andmal-bot.cloudfunctions.net/dotnet2",                   // dotnet
	}

	//OneUrlReq(urls[0], 20)
	//OneUrlReq(urls[1], 20)
	//OneUrlReq(urls[2], 100)
	//OneUrlReq(urls[3], 100)
	//OneUrlReq(urls[4], 20)

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			OneUrlReq(url, 30)
			wg.Done()
		}(url)
	}
	wg.Wait() // block

	//time.Sleep(20000)

	//fmt.Println("Wait group init")
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go MyFunc(&wg)
	//wg.Wait()
	//fmt.Println("Finished Main")
}

func OneUrlReq(url string, times int) {
	//url := urls[0]
	for i := 0; i < times; i++ {
		response, err := http.Get(url)
		if err != nil {
			return
		}
		status := response.StatusCode
		fmt.Printf(" -- Req %d for %s DONE. Status: %d \n", i, url, status)
	}
}

func AllLinks(urls []string) {
	for _, url := range urls {
		fmt.Printf(" -- link is %s \n", url)
		for i := 0; i < 30; i++ {
			http.Get(url)
			fmt.Printf(" -- Req %d for %s is DONE \n", i, url)
		}
	}
}

func WaitEx2() {
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(50)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		//for i := 0; i < 5; i++ {
		go func() {
			response, err := http.Get("https://example.com")
			if err != nil {
				log.Println(err)
			}
			fmt.Println(response.Body)
			defer wg.Done()
			//time.Sleep(200)
		}()
		//}
	}

	//wg.Wait()
}

func client() *http.Client {
	client := http.Client{Timeout: 8 * time.Second}
	return &client
}
