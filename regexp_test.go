package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestReadLogFile(t *testing.T) {
	var regExp = regexp.MustCompile("^(\\d+\\.\\d+\\.\\d+\\.\\d+) (.*) (.*) (\\[.*\\]) (\".*\") (\".*\") (\".*\")$")

	logData, err := os.ReadFile("C:\\logs\\jira\\_prod\\2025\\2025-03-08\\access_log.2025-03-08")
	if err != nil {
		return
	}

	lines := strings.Split(string(logData), "\n")

	//var m runtime.MemStats
	//runtime.ReadMemStats(&m)
	//fmt.Printf("Alloc = %s", formatBytes(m.Alloc))
	//fmt.Printf("\tTotalAlloc = %s", formatBytes(m.TotalAlloc))
	//fmt.Printf("\tSys = %s", formatBytes(m.Sys))
	//fmt.Printf("\tNumGC = %v\n", m.NumGC)

	userNames := make(map[string]struct{})

	for _, line := range lines {
		matchString := regExp.MatchString(line)

		if matchString {
			submatch := regExp.FindAllSubmatch([]byte(line), -1)

			//ipData := string(submatch[0][ipIdx])
			//reqId := string(submatch[0][reqIdx])
			//username := string(submatch[0][userNameIdx])
			//timestmp := string(submatch[0][timeIdx])
			//reqData := string(submatch[0][requestIdx])
			//clientData := string(submatch[0][clientIdx])
			//sessionData := string(submatch[0][sessionIdx])

			//for i := 0; i < 7; i++ {
			//	dataPart := submatch[0][i]
			//	println(string(dataPart))
			//}

			username := string(submatch[0][3])
			if _, exists := userNames[username]; exists {
				//fmt.Println("")
			} else {
				userNames[username] = struct{}{}
			}

		}
	}

	for key, _ := range userNames {
		println(key)
	}

}

func ReadLineByLine() {
	filePath := "my_file.txt"

	// 1. Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	// 2. Ensure the file is closed at the end of the function
	defer file.Close()

	// 3. Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	fmt.Println("File Content (line by line):")
	// 4. Loop through the file line by line
	for scanner.Scan() {
		// scanner.Text() returns the current line as a string
		fmt.Println(scanner.Text())
	}

	// 5. Check for any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
	}
}

func TestParseLogs(t *testing.T) {

	//
	var regExp = regexp.MustCompile("^(\\d+\\.\\d+\\.\\d+\\.\\d+) (.*) (.*) (\\[.*\\]) (\".*\") (\".*\") (\".*\")$")

	data := "127.0.0.1 158x6583999x2 a.navadiya [08/Mar/2025:02:38:51 -0800] \"GET /rest/internal/2.0/client- HTTP/1.0\" 200 18 6 \"https://jira.ontrq.com/secure/Dashboard.jspa?selectPageId=83502\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36\" \"1sx1d6m\""

	matchString := regExp.MatchString(data)

	if matchString {
		submatch := regExp.FindAllSubmatch([]byte(data), -1)

		//ipIdx := 0
		//reqIdx := 1
		//userNameIdx := 2
		//timeIdx := 3
		//requestIdx := 4
		//clientIdx := 5
		//sessionIdx := 6
		//
		//ipData := string(submatch[0][ipIdx])
		//reqId := string(submatch[0][reqIdx])
		//username := string(submatch[0][userNameIdx])
		//timestmp := string(submatch[0][timeIdx])
		//reqData := string(submatch[0][requestIdx])
		//clientData := string(submatch[0][clientIdx])
		//sessionData := string(submatch[0][sessionIdx])

		for i := 0; i < 7; i++ {
			dataPart := submatch[0][i]
			println(string(dataPart))
		}

	}
}
