package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

func TestRange2(t *testing.T) {
	today := time.Now()
	twoHoursAgo := today.Add(time.Duration(-2) * time.Hour)
	println(twoHoursAgo.String())
	//twoHoursAgo.Sub(time.Duration(22) * time.Hour)
	timer := time.NewTimer(time.Second * 10)
	t1 := timer.C
	println(t1)
	//println(time.Location.String())
	println(time.Duration.Minutes(3))
}

func TestDates1(t *testing.T) {
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
