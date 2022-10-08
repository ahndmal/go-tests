package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

func TestDates1(t *testing.T) {
	//
	december := time.December
	year, month, day := time.Now().Date()
	println(december)
	println(year, month, day)

	dir, err := ioutil.ReadDir("/home/malandr/Documents")
	if err != nil {
		log.Panicln(err)
	}
	var count int
	for _, fi := range dir {
		file, err := ioutil.ReadFile("/home/malandr/Documents/" + fi.Name())
		if err != nil {
			log.Panicln(err)
		}
		count += 1
		println(string(file))
	}
	println("================")
	println(fmt.Sprintf(">>>> Parsed %d files", count))
}
